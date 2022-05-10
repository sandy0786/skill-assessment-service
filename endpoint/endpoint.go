package endpoint

import (
	"context"
	"log"

	questionRequest "github.com/sandy0786/skill-assessment-service/request/question"
	userRequest "github.com/sandy0786/skill-assessment-service/request/user"
	service "github.com/sandy0786/skill-assessment-service/service"
	transport "github.com/sandy0786/skill-assessment-service/transport"

	"github.com/go-kit/kit/endpoint"
)

//Endpoints are exposed here
//Contains endpoints that map request from client to our service
type Endpoints struct {
	StatusEndpoint      endpoint.Endpoint
	AddUserEndpoint     endpoint.Endpoint
	GetAllUsersEndpoint endpoint.Endpoint
	// GetEmployeeByIdEndpoint endpoint.Endpoint
	AddQuestionEndpoint     endpoint.Endpoint
	GetAllQuestionsEndpoint endpoint.Endpoint
}

//MakeStatusEndpoint returns response
func MakeStatusEndpoint(srv service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		log.Println("MakeStatusEndpoint")
		responseFromService, err := srv.GetServiceStatus(ctx)
		if err != nil {
			return &transport.StatusResponse{Status: responseFromService}, err
		}
		return &transport.StatusResponse{Status: responseFromService}, nil
	}
}

//MakeAddUserEndpoint returns response
func MakeAddUserEndpoint(srv service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		log.Println("MakeAddUserEndpoint")
		userRequest := request.(userRequest.UserRequest)
		responseFromService, err := srv.AddUser(ctx, userRequest)
		return responseFromService, err
	}
}

// MakeGetAllUsersEndpoint returns response
func MakeGetAllUsersEndpoint(srv service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		log.Println("MakeGetAllUsersEndpoint")
		responseFromService, err := srv.GetAllUsers(ctx)
		return responseFromService, err
	}
}

// //MakeGetEmployeeByIdEndpoint returns response
// func MakeGetEmployeeByIdEndpoint(srv service.Service) endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (interface{}, error) {
// 		log.Println("MakeGetEmployeeByIdEndpoint")
// 		id, err := strconv.ParseInt(request.(string), 10, 64)
// 		responseFromService, err := srv.GetEmployeeById(ctx, id)
// 		return responseFromService, err
// 	}
// }

//MakeAddQuestionEndpoint returns response
func MakeAddQuestionEndpoint(srv service.QuestionService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		log.Println("MakeAddQuestionEndpoint")
		questionRequest := request.(questionRequest.QuestionRequest)
		responseFromService, err := srv.AddQuestion(ctx, questionRequest)
		return responseFromService, err
	}
}

// MakeGetAllQuestionsEndpoint returns response
func MakeGetAllQuestionsEndpoint(srv service.QuestionService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		log.Println("MakeGetAllQuestionsEndpoint")
		responseFromService, err := srv.GetAllQuestions(ctx)
		return responseFromService, err
	}
}
