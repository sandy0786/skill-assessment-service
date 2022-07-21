package transport

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"reflect"
	"time"

	userDTO "github.com/sandy0786/skill-assessment-service/dto/user"
	globalErr "github.com/sandy0786/skill-assessment-service/errors"
	jwtP "github.com/sandy0786/skill-assessment-service/jwt"
	userRequest "github.com/sandy0786/skill-assessment-service/request/user"
	userResponse "github.com/sandy0786/skill-assessment-service/response/user"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

var Validate *validator.Validate

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

	// verify token
	_, tokenErr := jwtP.VerifyToken(r)
	if tokenErr != (globalErr.GlobalError{}) {
		return r, tokenErr
	}

	var uRequest userRequest.UserRequest
	decodeErr := json.NewDecoder(r.Body).Decode(&uRequest)
	if decodeErr != nil {
		var errorMessage string
		if reflect.TypeOf(decodeErr).String() == "*json.SyntaxError" {
			errorMessage = "Invalid request body"
		} else {
			errorMessage = "Request body parse error"
		}
		return uRequest, globalErr.GlobalError{
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
	return r, nil
}

// EncodeGetAllUsersResponse - encodes status service response
func EncodeGetAllUsersResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	log.Println("transport:EncodeGetAllUsersResponse")
	resp := response.(userResponse.UserResults)
	if resp.TotalRecords == 0 {
		// if no questions found return empty response with 404 status code
		w.WriteHeader(http.StatusOK)
		var responseMap = make(map[string]interface{})
		responseMap["data"] = []interface{}{}
		responseMap["totalRecords"] = resp.TotalRecords
		return json.NewEncoder(w).Encode(responseMap)
	}
	return json.NewEncoder(w).Encode(response)
}

//DecodeDeleteUserRequest - decodes status GET request
func DecodeDeleteUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	log.Println("transport:DecodeDeleteUserRequest")

	// verify token
	_, err := jwtP.VerifyToken(r)
	if err != (globalErr.GlobalError{}) {
		return r, err
	}

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

	// verify token
	_, err := jwtP.VerifyToken(r)
	if err != (globalErr.GlobalError{}) {
		return r, err
	}

	username := mux.Vars(r)["username"]
	if len(username) == 0 {
		return "", errors.New("Path variable 'username' not found")
	}
	var userDto userDTO.UserDTO
	err1 := json.NewDecoder(r.Body).Decode(&userDto)
	err1 = Validate.Struct(userDto)
	log.Println("aa >> ", err1)
	return userDto, nil
}

// EncodePasswordResetRequest
func EncodePasswordResetRequest(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	log.Println("transport:EncodeDeleteUserResponse")
	return json.NewEncoder(w).Encode(response)
}

//DecodeUserPasswordResetRequest
func DecodeUserPasswordResetRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	log.Println("transport:DecodeUserPasswordResetRequest")

	// verify token
	claims, err := jwtP.VerifyToken(r)
	if err != (globalErr.GlobalError{}) {
		return r, err
	}

	// username := claims.Username

	var userDto userDTO.UserDTO
	err1 := json.NewDecoder(r.Body).Decode(&userDto)
	userDto.Username = claims.Username
	err1 = Validate.Struct(userDto)
	log.Println("aa >> ", err1)
	return userDto, nil
}

//ErrorEncoder will encode error to our format
func ErrorEncoder(ctx context.Context, err1 error, w http.ResponseWriter) {
	log.Println("transport:ErrorEncoder: Inside ErrorEncoder: ")
	var globalError globalErr.GlobalError

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
		case "RoleId":
			message = "Provide valid roleId"
		}

		globalError = globalErr.GlobalError{
			TimeStamp: time.Now().UTC().String()[0:19],
			Status:    http.StatusBadRequest,
			Message:   message,
		}
	} else {
		globalError, ok = err1.(globalErr.GlobalError)
		if !ok {
			globalError = globalErr.GlobalError{
				TimeStamp: time.Now().UTC().String()[0:19],
				Status:    http.StatusInternalServerError,
				Message:   "Something went wrong. Please try again after sometime",
			}
		}
	}

	w.WriteHeader(globalError.Status)
	json.NewEncoder(w).Encode(globalError)
}
