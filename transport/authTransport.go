package transport

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	err "github.com/sandy0786/skill-assessment-service/errors"
	authRequest "github.com/sandy0786/skill-assessment-service/request/auth"

	"github.com/go-playground/validator"
)

//DecodeAuthRequest - decodes POST request
func DecodeAuthRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	log.Println("transport:DecodeAuthRequest")
	var authRequest authRequest.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&authRequest)
	err = Validate.Struct(authRequest)
	return authRequest, err
}

// EncodeAuthResponse - encodes status service response
func EncodeAuthResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	log.Println("transport:EncodeAuthResponse ")
	return json.NewEncoder(w).Encode(response)
}

//AuthErrorEncoder will encode error to our format
func AuthErrorEncoder(ctx context.Context, err1 error, w http.ResponseWriter) {
	log.Println("transport:ErrorEncoder: Inside AuthErrorEncoder ", err1)
	var globalError err.GlobalError
	if errorFields, ok := err1.(validator.ValidationErrors); ok {
		var message string

		switch errorFields[0].Field() {
		case "Username":
			message = "Username should not contain any special characters and should be atleast 5 characters"
		case "Password":
			message = "Password should have the length of atleast 8 characters"
		}

		globalError = err.GlobalError{
			TimeStamp: time.Now().UTC().String()[0:19],
			Status:    http.StatusBadRequest,
			Message:   message,
		}
	} else {
		globalError, ok = err1.(err.GlobalError)
		var message string
		var status int
		if err1.Error() == "mongo: no documents in result" {
			message = "Invalid username or password"
			status = http.StatusUnauthorized
		} else {
			message = "Something went wrong. Please try again after sometime"
			status = http.StatusInternalServerError
		}
		if !ok {
			globalError = err.GlobalError{
				TimeStamp: time.Now().UTC().String()[0:19],
				Status:    status,
				Message:   message,
			}
		}
	}

	w.WriteHeader(globalError.Status)
	json.NewEncoder(w).Encode(globalError)
}
