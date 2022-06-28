package service

import (
	"context"
	"log"
	"time"

	configuration "github.com/sandy0786/skill-assessment-service/configuration"
	authDao "github.com/sandy0786/skill-assessment-service/dao/auth"
	authDocument "github.com/sandy0786/skill-assessment-service/documents/auth"
	jwtP "github.com/sandy0786/skill-assessment-service/jwt"
	authRequest "github.com/sandy0786/skill-assessment-service/request/auth"
	authResponse "github.com/sandy0786/skill-assessment-service/response/auth"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/copier"
)

type AuthService interface {
	Login(context.Context, authRequest.LoginRequest) (authResponse.LoginResponse, error)
	Logout(context.Context) (authResponse.LoginResponse, error)
}

// authentication and authorisation servvice
type authService struct {
	testServiceConfig configuration.ConfigurationInterface
	dao               authDao.AuthDAO
}

func NewAuthService(c configuration.ConfigurationInterface, dao authDao.AuthDAO) *authService {
	return &authService{
		testServiceConfig: c,
		dao:               dao,
	}
}

func (a *authService) Login(ctx context.Context, userRequest authRequest.LoginRequest) (authResponse.LoginResponse, error) {
	log.Println("Inside user login")
	var authDoc authDocument.User
	copier.Copy(&authDoc, &userRequest)
	// user.Active = true
	validUser, err := a.dao.Login(&authDoc)
	if validUser {
		// Declare the expiration time of the token
		// here, we have kept it as 5 minutes
		expirationTime := time.Now().Add(time.Duration(a.testServiceConfig.GetConfigDetails().Jwt.ExpiryTime) * time.Minute)
		// Create the JWT claims, which includes the username and expiry time
		claims := &jwtP.Claims{
			Username: userRequest.Username,
			Role:     "admin",
			StandardClaims: jwt.StandardClaims{
				// In JWT, the expiry time is expressed as unix milliseconds
				ExpiresAt: expirationTime.Unix(),
			},
		}

		// Declare the token with the algorithm used for signing, and the claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Create the JWT string
		secret := []byte(a.testServiceConfig.GetConfigDetails().Jwt.Secret)
		tokenString, err := token.SignedString(secret)
		return authResponse.LoginResponse{
			Token: tokenString,
		}, err
	}

	return authResponse.LoginResponse{Token: ""}, err
}

func (a *authService) Logout(_ context.Context) (authResponse.LoginResponse, error) {
	log.Println("Inside user logout")
	return authResponse.LoginResponse{}, nil
}
