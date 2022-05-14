package endpoint

import (
	"context"
	"log"

	questionDTO "github.com/sandy0786/skill-assessment-service/dto/question"
	categoryRequest "github.com/sandy0786/skill-assessment-service/request/category"
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
	AddQuestionEndpoint         endpoint.Endpoint
	AddMultipleQuestionEndpoint endpoint.Endpoint
	GetAllQuestionsEndpoint     endpoint.Endpoint
	AddCategoryEndpoint         endpoint.Endpoint
	GetAllCategoriesEndpoint    endpoint.Endpoint
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
		questionDto := request.(questionDTO.QuestionDTO)
		responseFromService, err := srv.AddQuestion(ctx, questionDto)
		return responseFromService, err
	}
}

//MakeAddMultipleQuestionsEndpoint returns response
func MakeAddMultipleQuestionsEndpoint(srv service.QuestionService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		log.Println("MakeAddMultipleQuestionsEndpoint")
		// questionsRequest := request.([]questionRequest.QuestionRequest)
		questionsDto := request.(questionDTO.QuestionsDTO)
		responseFromService, err := srv.AddMultipleQuestions(ctx, questionsDto)
		return responseFromService, err
	}
}

// MakeGetAllQuestionsEndpoint returns response
func MakeGetAllQuestionsEndpoint(srv service.QuestionService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		log.Println("MakeGetAllQuestionsEndpoint")
		category := request.(string)
		responseFromService, err := srv.GetAllQuestions(ctx, category)
		return responseFromService, err
	}
}

//MakeAddCategoryEndpoint returns response
func MakeAddCategoryEndpoint(srv service.CategoryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		log.Println("MakeAddCategoryEndpoint")
		categoryRequest := request.(categoryRequest.CategoryRequest)
		responseFromService, err := srv.AddCategory(ctx, categoryRequest)
		return responseFromService, err
	}
}

//MakeGetAllCategoriesEndpoint returns response
func MakeGetAllCategoriesEndpoint(srv service.CategoryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		log.Println("MakeGetAllCategoriesEndpoint")
		responseFromService, err := srv.GetAllCategories(ctx)
		return responseFromService, err
	}
}
