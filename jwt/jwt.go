package jwt

// https://www.sohamkamani.com/golang/jwt-authentication/

import (
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	err "github.com/sandy0786/skill-assessment-service/errors"
)

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func VerifyToken(r *http.Request) (*Claims, err.GlobalError) {
	if len(r.Header["Authorization"]) == 0 {
		return &Claims{}, err.GlobalError{
			TimeStamp: time.Now().UTC().String()[0:19],
			Status:    http.StatusBadRequest,
			Message:   "Authorization header is missing",
		}
	}

	Bearertoken := r.Header["Authorization"][0]

	if !strings.Contains(Bearertoken, "Bearer") {
		return &Claims{}, err.GlobalError{
			TimeStamp: time.Now().UTC().String()[0:19],
			Status:    http.StatusBadRequest,
			Message:   "Invalid token",
		}
	}

	token := strings.Split(Bearertoken, "Bearer ")[1]

	// Initialize a new instance of `Claims`
	claims := &Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err1 := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("t%682@#90"), nil
	})

	if err1 != nil {
		return &Claims{}, err.GlobalError{
			TimeStamp: time.Now().UTC().String()[0:19],
			Status:    http.StatusUnauthorized,
			Message:   "User is not authorized",
		}
	}

	if !tkn.Valid {
		return &Claims{}, err.GlobalError{
			TimeStamp: time.Now().UTC().String()[0:19],
			Status:    http.StatusForbidden,
			Message:   "User has no rights to access",
		}
	}

	return claims, err.GlobalError{}
}
