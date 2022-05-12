package transport

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	questionDTO "github.com/sandy0786/skill-assessment-service/dto/question"
	errors "github.com/sandy0786/skill-assessment-service/errors"
	questionRequest "github.com/sandy0786/skill-assessment-service/request/question"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

//DecodeAddQuestionRequest - decodes status GET request
func DecodeAddQuestionRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	log.Println("transport:DecodeAddUserRequest")
	var qRequest questionRequest.QuestionRequest
	err := json.NewDecoder(r.Body).Decode(&qRequest)
	err = Validate.Struct(qRequest)
	log.Println("aa >> ", err)
	log.Println("path >> ", r.URL.Path)

	category := mux.Vars(r)["category"]
	// pathSplit := strings.Split(req.URL.Path, "/")
	// dataSourceName := pathSplit[len(pathSplit)-1]
	// log.Println("category >> ", category)
	qstnDto := questionDTO.QuestionDTO{
		category: category,
		Question: qRequest,
	}
	return qstnDto, err
}

// EncodeAddQuestionResponse - encodes status service response
func EncodeAddQuestionResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	log.Println("transport:EncodeAddQuestionResponse")
	return json.NewEncoder(w).Encode(response)
}

//DecodeAddMutlipleQuestionsRequest - decodes status POST request
func DecodeAddMutlipleQuestionsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	log.Println("transport:DecodeAddMutlipleQuestionsRequest")
	var qRequests []questionRequest.QuestionRequest
	err := json.NewDecoder(r.Body).Decode(&qRequests)
	for _, qRequest := range qRequests {
		err = Validate.Struct(qRequest)
	}
	return qRequests, err
}

// EncodeAddMultipleQuestionsResponse - encodes status service response
func EncodeAddMultipleQuestionsResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	log.Println("transport:EncodeAddMultipleQuestionsResponse")
	return json.NewEncoder(w).Encode(response)
}

func DecodeGetAllQuestionsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	log.Println("transport:DecodeGetAllQuestionsRequest")
	return r, nil
}

// EncodeGetAllQuestionsResponse - encodes status service response
func EncodeGetAllQuestionsResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	log.Println("transport:EncodeGetAllQuestionsResponse")
	return json.NewEncoder(w).Encode(response)
}

//QuestionErrorEncoder will encode error to our format
func QuestionErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	log.Println("transport:ErrorEncoder: Inside ErrorEncoder: ")
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
