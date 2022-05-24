package transport

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	userDTO "github.com/sandy0786/skill-assessment-service/dto/user"
	err "github.com/sandy0786/skill-assessment-service/errors"
	userRequest "github.com/sandy0786/skill-assessment-service/request/user"
	userResponse "github.com/sandy0786/skill-assessment-service/response/user"

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
	err := json.NewDecoder(r.Body).Decode(&uRequest)
	err = Validate.Struct(uRequest)
	log.Println("aa >> ", err)
	log.Println("path >> ", r.URL.Path)

	return uRequest, err
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
	if _, ok := err1.(validator.ValidationErrors); ok {
		log.Println("err ... ", err1.Error())
		message := err1.Error()
		if strings.Contains(err1.Error(), ".Age") {
			message = "Age should be between 20 and 60"
		}
		globalError = err.GlobalError{
			TimeStamp: time.Now().UTC().String()[0:19],
			Status:    http.StatusBadRequest,
			Message:   message,
		}
	}
	w.WriteHeader(globalError.Status)
	json.NewEncoder(w).Encode(globalError)
}
