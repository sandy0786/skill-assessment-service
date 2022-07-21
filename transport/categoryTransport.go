package transport

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	errors "github.com/sandy0786/skill-assessment-service/errors"
	categoryRequest "github.com/sandy0786/skill-assessment-service/request/category"
	categoryResponse "github.com/sandy0786/skill-assessment-service/response/category"

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
	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(response)
}

func DecodeGetAllCategoriesRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	log.Println("transport:DecodeGetAllCategoriesRequest")
	return r, nil
}

// EncodeGetAllCategoriesResponse - encodes status service response
func EncodeGetAllCategoriesResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	log.Println("transport:EncodeGetAllCategoriesResponse")
	resp := response.(categoryResponse.CategoryResults)

	if resp.TotalRecords == 0 {
		// if no questions found return empty response with 404 status code
		w.WriteHeader(http.StatusOK)
		var responseMap = make(map[string]interface{})
		responseMap["data"] = []interface{}{}
		responseMap["totalRecords"] = resp.TotalRecords
		return json.NewEncoder(w).Encode(responseMap)
	}
	return json.NewEncoder(w).Encode(resp)
}

//CategoryErrorEncoder will encode error to our format
func CategoryErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	log.Println("transport:ErrorEncoder: Inside CategoryErrorEncoder : ", err)
	var globalError errors.GlobalError

	// Return proper validation error message
	if errorFields, ok := err.(validator.ValidationErrors); ok {
		var message string
		switch errorFields[0].Field() {
		case "Category":
			message = "'category' should not contain any special characters and should be atleast 2 characters"
		case "Author":
			message = "'author' should not contain any special characters and should be atleast 5 characters"
		}

		globalError = errors.GlobalError{
			TimeStamp: time.Now().UTC().String()[0:19],
			Status:    http.StatusBadRequest,
			Message:   message,
		}
	} else {
		globalError, ok = err.(errors.GlobalError)
		if !ok {
			globalError = errors.GlobalError{
				TimeStamp: time.Now().UTC().String()[0:19],
				Status:    http.StatusInternalServerError,
				Message:   "Something went wrong. Please try again after sometime",
			}
		}
	}

	w.WriteHeader(globalError.Status)
	json.NewEncoder(w).Encode(globalError)
}
