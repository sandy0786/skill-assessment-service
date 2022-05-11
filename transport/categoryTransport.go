package transport

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	errors "github.com/sandy0786/skill-assessment-service/errors"
	categoryRequest "github.com/sandy0786/skill-assessment-service/request/category"

	"github.com/go-playground/validator"
)

//DecodeAddCategoryRequest - decodes POST request
func DecodeAddCategoryRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	log.Println("transport:DecodeAddCategoryRequest")
	var cRequest categoryRequest.CategoryRequest
	err := json.NewDecoder(r.Body).Decode(&cRequest)
	err = Validate.Struct(cRequest)
	return cRequest, err
}

// EncodeAddCategoryResponse - encodes status service response
func EncodeAddCategoryResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	log.Println("transport:EncodeAddCategoryResponse")
	return json.NewEncoder(w).Encode(response)
}

func DecodeGetAllCategoriesRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	log.Println("transport:DecodeGetAllCategoriesRequest")
	return r, nil
}

// EncodeGetAllCategoriesResponse - encodes status service response
func EncodeGetAllCategoriesResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	log.Println("transport:EncodeGetAllCategoriesResponse")
	return json.NewEncoder(w).Encode(response)
}

//CategoryErrorEncoder will encode error to our format
func CategoryErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	log.Println("transport:ErrorEncoder: Inside CategoryErrorEncoder ")
	var globalError errors.GlobalError
	if _, ok := err.(validator.ValidationErrors); ok {
		// log.Println("err ... ", err.Error())
		message := err.Error()
		if strings.Contains(err.Error(), ".Age") {
			message = "Age should be between 20 and 60"
		}
		globalError = errors.GlobalError{
			TimeStamp: time.Now().UTC().String(),
			Status:    400,
			Message:   message,
		}
	}
	// finalResponse := map[string]string{"a": "b"}
	w.WriteHeader(globalError.Status)
	json.NewEncoder(w).Encode(globalError)
}
