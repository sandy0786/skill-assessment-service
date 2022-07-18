package transport

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	err "github.com/sandy0786/skill-assessment-service/errors"
	roleResponse "github.com/sandy0786/skill-assessment-service/response/role"

	"github.com/go-playground/validator"
)

//DecodeGetAllRolesRequest - decodes POST request
func DecodeGetAllRolesRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	log.Println("transport:DecodeGetAllRolesRequest")
	return r, nil
}

// EncodeGetAllRolesResponse - encodes status service response
func EncodeGetAllRolesResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	log.Println("transport:EncodeGetAllRolesResponse ")
	resp := response.([]roleResponse.Role)
	if len(resp) == 0 {
		// if no questions found return empty response with 404 status code
		w.WriteHeader(http.StatusNotFound)
		return json.NewEncoder(w).Encode([]interface{}{})
	}
	return json.NewEncoder(w).Encode(response)
}

//RoleErrorEncoder will encode error to our format
func RoleErrorEncoder(ctx context.Context, err1 error, w http.ResponseWriter) {
	log.Println("transport:ErrorEncoder: Inside RoleErrorEncoder ")
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
