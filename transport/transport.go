package transport

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	errors "github.com/sandy0786/skill-assessment-service/errors"
	userRequest "github.com/sandy0786/skill-assessment-service/request/user"

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
	finalResponse.Status = `ok`
	return json.NewEncoder(w).Encode(finalResponse)
}

//EncodeAddEmployeeResponse - decodes status GET request
func DecodeAddEmployeeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	log.Println("transport:DecodeAddEmployeeRequest")
	var uRequest userRequest.UserRequest
	err := json.NewDecoder(r.Body).Decode(&uRequest)
	err = Validate.Struct(uRequest)
	log.Println("aa >> ", err)
	log.Println("path >> ", r.URL.Path)

	return uRequest, err
}

// EncodeAddEmployeeResponse - encodes status service response
func EncodeAddEmployeeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	log.Println("transport:EncodeAddEmployeeResponse")
	return json.NewEncoder(w).Encode(response)
}

func DecodeGetAllEmpRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	log.Println("transport:DecodeGetAllEmpRequest")
	return r, nil
}

// EncodeGetAllEmpResponse - encodes status service response
func EncodeGetAllEmpResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	log.Println("transport:EncodeGetAllEmpResponse")
	return json.NewEncoder(w).Encode(response)
}

func DecodeGetEmpByIdRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	log.Println("transport:DecodeGetEmpByIdRequest")
	vars := mux.Vars(r)
	empId, ok := vars["id"]
	if !ok {
		log.Println("id is missing in parameters")
	}
	return empId, nil
}

//ErrorEncoder will encode error to our format
func ErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	log.Println("transport:ErrorEncoder: Inside ErrorEncoder: ")
	var globalError errors.GlobalError
	if _, ok := err.(validator.ValidationErrors); ok {
		log.Println("err ... ", err.Error())
		var message string
		if strings.Contains(err.Error(), ".Age") {
			message = "Age should be between 20 and 60"
		}
		globalError = errors.GlobalError{
			Result:    nil,
			TimeStamp: time.Now().UTC().String(),
			Status:    400,
			Message:   message,
		}
	}
	// finalResponse := map[string]string{"a": "b"}
	json.NewEncoder(w).Encode(globalError)
}
