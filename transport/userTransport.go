package transport

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	userDTO "github.com/sandy0786/skill-assessment-service/dto/user"
	err "github.com/sandy0786/skill-assessment-service/errors"
	jwtP "github.com/sandy0786/skill-assessment-service/jwt"
	userRequest "github.com/sandy0786/skill-assessment-service/request/user"
	userResponse "github.com/sandy0786/skill-assessment-service/response/user"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

var Validate *validator.Validate

type getDruidRequest struct {
	queryRequest http.Request
}

type StatusRequest struct{}
type StatusResponse struct {
	Status string `json:"status"`
}

//DecodeStatusRequest - decodes status GET request
func DecodeStatusRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	log.Println("transport:DecodeStatusRequest: Inside DecodeStatusRequest")
	return r, nil
}

// EncodeStatusResponse - encodes status service response
func EncodeStatusResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	log.Println("transport:EncodeStatusResponse: Inside EncodeStatusResponse")
	var finalResponse StatusResponse
	finalResponse.Status = `UP`
	return json.NewEncoder(w).Encode(finalResponse)
}

//DecodeAddUserRequest - decodes status GET request
func DecodeAddUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	log.Println("transport:DecodeAddUserRequest")
	var uRequest userRequest.UserRequest
	decodeErr := json.NewDecoder(r.Body).Decode(&uRequest)
	if decodeErr != nil {
		var errorMessage string
		if reflect.TypeOf(decodeErr).String() == "*json.SyntaxError" {
			errorMessage = "Invalid request body"
		} else {
			errorMessage = "Request body parse error"
		}
		return uRequest, err.GlobalError{
			TimeStamp: time.Now().UTC().String()[0:19],
			Status:    http.StatusBadRequest,
			Message:   errorMessage,
		}
	}
	validateErr := Validate.Struct(uRequest)
	return uRequest, validateErr
}

// EncodeAddUserResponse - encodes status service response
func EncodeAddUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	log.Println("transport:EncodeAddUserResponse")
	return json.NewEncoder(w).Encode(response)
}

func DecodeGetAllUsersRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	log.Println("transport:DecodeGetAllUsersRequest")

	if len(r.Header["Authorization"]) == 0 {
		return r, err.GlobalError{
			TimeStamp: time.Now().UTC().String()[0:19],
			Status:    http.StatusBadRequest,
			Message:   "Authorization header is missing",
		}
	}

	Bearertoken := r.Header["Authorization"][0]
	token := strings.Split(Bearertoken, "Bearer ")[1]

	// Initialize a new instance of `Claims`
	claims := &jwtP.Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err1 := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("t%682@#90"), nil
	})

	if err1 != nil {
		return r, err.GlobalError{
			TimeStamp: time.Now().UTC().String()[0:19],
			Status:    http.StatusUnauthorized,
			Message:   "User is not authorized",
		}
	}

	if !tkn.Valid {
		return r, err.GlobalError{
			TimeStamp: time.Now().UTC().String()[0:19],
			Status:    http.StatusForbidden,
			Message:   "User has no rights to access",
		}
	}

	return r, nil
}

// EncodeGetAllUsersResponse - encodes status service response
func EncodeGetAllUsersResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	log.Println("transport:EncodeGetAllUsersResponse")
	resp := response.([]userResponse.UserResponse)
	if len(resp) == 0 {
		// if no questions found return empty response with 404 status code
		w.WriteHeader(http.StatusNotFound)
		return json.NewEncoder(w).Encode([]interface{}{})
	}
	return json.NewEncoder(w).Encode(response)
}

// func DecodeGetEmpByIdRequest(ctx context.Context, r *http.Request) (interface{}, error) {
// 	log.Println("transport:DecodeGetEmpByIdRequest")
// 	vars := mux.Vars(r)
// 	empId, ok := vars["id"]
// 	if !ok {
// 		log.Println("id is missing in parameters")
// 	}
// 	return empId, nil
// }

//DecodeDeleteUserRequest - decodes status GET request
func DecodeDeleteUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	log.Println("transport:DecodeDeleteUserRequest")
	username := mux.Vars(r)["username"]
	if len(username) == 0 {
		return "", errors.New("Path variable 'username' not found")
	}
	return username, nil
}

// EncodeDeleteUserResponse - encodes status service response
func EncodeDeleteUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	log.Println("transport:EncodeDeleteUserResponse")
	return json.NewEncoder(w).Encode(response)
}

//DecodePasswordResetRequest
func DecodePasswordResetRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	log.Println("transport:DecodePasswordResetRequest")
	username := mux.Vars(r)["username"]
	if len(username) == 0 {
		return "", errors.New("Path variable 'username' not found")
	}
	var userDto userDTO.UserDTO
	err := json.NewDecoder(r.Body).Decode(&userDto)
	err = Validate.Struct(userDto)
	log.Println("aa >> ", err)
	return userDto, nil
}

// EncodePasswordResetRequest
func EncodePasswordResetRequest(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	log.Println("transport:EncodeDeleteUserResponse")
	return json.NewEncoder(w).Encode(response)
}

//ErrorEncoder will encode error to our format
func ErrorEncoder(ctx context.Context, err1 error, w http.ResponseWriter) {
	log.Println("transport:ErrorEncoder: Inside ErrorEncoder: ")
	var globalError err.GlobalError

	// Return proper validation error message
	if errorFields, ok := err1.(validator.ValidationErrors); ok {
		var message string

		switch errorFields[0].Field() {
		case "Username":
			message = "Username should not contain any special characters and should be atleast 5 characters"
		case "Password":
			message = "Password should have the length of atleast 8 characters"
		case "Email":
			message = "Invalid email"
		case "Role":
			message = "Provide valid role"
		}

		globalError = err.GlobalError{
			TimeStamp: time.Now().UTC().String()[0:19],
			Status:    http.StatusBadRequest,
			Message:   message,
		}
	} else {
		globalError, ok = err1.(err.GlobalError)
		if !ok {
			globalError = err.GlobalError{
				TimeStamp: time.Now().UTC().String()[0:19],
				Status:    http.StatusInternalServerError,
				Message:   "Something went wrong. Please try again after sometime",
			}
		}
	}

	w.WriteHeader(globalError.Status)
	json.NewEncoder(w).Encode(globalError)
}
