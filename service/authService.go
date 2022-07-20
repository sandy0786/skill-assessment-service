package service

import (
	"context"
	"log"
	"net/http"
	"time"

	configuration "github.com/sandy0786/skill-assessment-service/configuration"
	authDao "github.com/sandy0786/skill-assessment-service/dao/auth"
	roleDao "github.com/sandy0786/skill-assessment-service/dao/role"
	authDocument "github.com/sandy0786/skill-assessment-service/documents/auth"
	err "github.com/sandy0786/skill-assessment-service/errors"
	jwtP "github.com/sandy0786/skill-assessment-service/jwt"
	authRequest "github.com/sandy0786/skill-assessment-service/request/auth"
	authResponse "github.com/sandy0786/skill-assessment-service/response/auth"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/copier"
)

type AuthService interface {
	Login(context.Context, authRequest.LoginRequest) (authResponse.LoginResponse, error)
	Logout(context.Context) (authResponse.LoginResponse, error)
	Refresh(context.Context, string) (authResponse.LoginResponse, error)
}

// authentication and authorisation servvice
type authService struct {
	config  configuration.ConfigurationInterface
	dao     authDao.AuthDAO
	roleDao roleDao.RoleDAO
}

func NewAuthService(c configuration.ConfigurationInterface, dao authDao.AuthDAO, roleDao roleDao.RoleDAO) *authService {
	return &authService{
		config:  c,
		dao:     dao,
		roleDao: roleDao,
	}
}

func (a *authService) Login(ctx context.Context, userRequest authRequest.LoginRequest) (authResponse.LoginResponse, error) {
	log.Println("Inside user login")
	var authDoc authDocument.User
	copier.Copy(&authDoc, &userRequest)
	// user.Active = true
	user, validUser, err1 := a.dao.Login(&authDoc)
	if validUser {
		if !user.Active {
			return authResponse.LoginResponse{}, err.GlobalError{
				TimeStamp: time.Now().UTC().String()[0:19],
				Status:    http.StatusUnauthorized,
				Message:   "Account is disabled. Please contact administrator",
			}
		}
		// get role from roleId
		role, _ := a.roleDao.GetRoleById(user.Role)
		// Declare the expiration time of the token
		// here, we have kept it as 5 minutes
		expirationTime := time.Now().UTC().Add(time.Duration(a.config.GetConfigDetails().Jwt.ExpiryTime) * time.Minute)
		// Create the JWT claims, which includes the username and expiry time
		claims := &jwtP.Claims{
			Username: userRequest.Username,
			Role:     role,
			StandardClaims: jwt.StandardClaims{
				// In JWT, the expiry time is expressed as unix milliseconds
				ExpiresAt: expirationTime.Unix(),
			},
		}

		// Declare the token with the algorithm used for signing, and the claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Create the JWT string
		secret := []byte(a.config.GetConfigDetails().Jwt.Secret)
		tokenString, err := token.SignedString(secret)
		return authResponse.LoginResponse{
			Token: tokenString,
		}, err
	}

	return authResponse.LoginResponse{Token: ""}, err1
}

func (a *authService) Logout(_ context.Context) (authResponse.LoginResponse, error) {
	log.Println("Inside user logout")
	return authResponse.LoginResponse{}, nil
}

func (a *authService) Refresh(ctx context.Context, token string) (authResponse.LoginResponse, error) {
	log.Println("Inside refresh token")

	claims := &jwtP.Claims{}
	tkn, err1 := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.config.GetConfigDetails().Jwt.Secret), nil
	})
	if err1 != nil {
		return authResponse.LoginResponse{}, err.GlobalError{
			TimeStamp: time.Now().UTC().String()[0:19],
			Status:    http.StatusBadRequest,
			Message:   "Invalid token",
		}
	}
	if !tkn.Valid {
		return authResponse.LoginResponse{}, err.GlobalError{
			TimeStamp: time.Now().UTC().String()[0:19],
			Status:    http.StatusBadRequest,
			Message:   "Invalid token",
		}
	}

	// We ensure that a new token is not issued until enough time has elapsed
	// In this case, a new token will only be issued if the old token is within
	// 30 seconds of expiry. Otherwise, return a bad request status
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		return authResponse.LoginResponse{}, err.GlobalError{
			TimeStamp: time.Now().UTC().String()[0:19],
			Status:    http.StatusBadRequest,
			Message:   "Invalid token",
		}
	}

	// Now, create a new token for the current use, with a renewed expiration time
	expirationTime := time.Now().UTC().Add(time.Duration(a.config.GetConfigDetails().Jwt.ExpiryTime) * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	tempToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err1 := tempToken.SignedString([]byte(a.config.GetConfigDetails().Jwt.Secret))
	if err1 != nil {
		return authResponse.LoginResponse{}, err.GlobalError{
			TimeStamp: time.Now().UTC().String()[0:19],
			Status:    http.StatusInternalServerError,
			Message:   "Something went wrong",
		}
	}

	return authResponse.LoginResponse{Token: tokenString}, nil
}
