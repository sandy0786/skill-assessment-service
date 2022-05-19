package transport

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	questionDTO "github.com/sandy0786/skill-assessment-service/dto/question"
	errors "github.com/sandy0786/skill-assessment-service/errors"
	questionRequest "github.com/sandy0786/skill-assessment-service/request/question"
	questionResponse "github.com/sandy0786/skill-assessment-service/response/question"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

// swagger:route GET /admin/company/ admin listCompany
// Get companies list
//
// security:
// - apiKey: []
// responses:
//  401: CommonError
//  200: GetCompanies
//DecodeAddQuestionRequest - decodes status GET request
func DecodeAddQuestionRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	log.Println("transport:DecodeAddUserRequest")
	var qRequest questionRequest.QuestionRequest
	err := json.NewDecoder(r.Body).Decode(&qRequest)
	err = Validate.Struct(qRequest)
	// log.Println("aa >> ", err)
	// log.Println("path >> ", r.URL.Path)

	category := mux.Vars(r)["category"]
	// pathSplit := strings.Split(req.URL.Path, "/")
	// dataSourceName := pathSplit[len(pathSplit)-1]
	// log.Println("category >> ", category)
	qstnDto := questionDTO.QuestionDTO{
		Category: category,
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

	category := mux.Vars(r)["category"]
	qstnsDto := questionDTO.QuestionsDTO{
		Category: category,
		Question: qRequests,
	}
	return qstnsDto, err
}

// EncodeAddMultipleQuestionsResponse - encodes status service response
func EncodeAddMultipleQuestionsResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	log.Println("transport:EncodeAddMultipleQuestionsResponse")
	return json.NewEncoder(w).Encode(response)
}

func DecodeGetAllQuestionsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	log.Println("transport:DecodeGetAllQuestionsRequest")
	// extract category name
	category := mux.Vars(r)["category"]
	return category, nil
}

// EncodeGetAllQuestionsResponse - encodes status service response
func EncodeGetAllQuestionsResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	log.Println("transport:EncodeGetAllQuestionsResponse")
	resp := response.([]questionResponse.QuestionResponse)
	if len(resp) == 0 {
		// if no questions found return empty response with 404 status code
		w.WriteHeader(http.StatusNotFound)
		return json.NewEncoder(w).Encode([]interface{}{})
	}
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
}

//QuestionErrorEncoder will encode error to our format
func QuestionErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	log.Println("transport:ErrorEncoder: Inside ErrorEncoder: ")
	var globalError errors.GlobalError
	if _, ok := err.(validator.ValidationErrors); ok {
		// log.Println("err ... ", err.Error())
		message := err.Error()

		globalError = errors.GlobalError{
			TimeStamp: time.Now().UTC().String()[0:19],
			Status:    http.StatusBadRequest,
			Message:   message,
		}
	}
	// finalResponse := map[string]string{"a": "b"}
	w.WriteHeader(globalError.Status)
	json.NewEncoder(w).Encode(globalError)
}
