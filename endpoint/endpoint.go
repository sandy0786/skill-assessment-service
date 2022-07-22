package endpoint

import (
	"context"
	"log"
	"net/http"

	categoryDTO "github.com/sandy0786/skill-assessment-service/dto/category"
	questionDTO "github.com/sandy0786/skill-assessment-service/dto/question"
	userDTO "github.com/sandy0786/skill-assessment-service/dto/user"
	authRequest "github.com/sandy0786/skill-assessment-service/request/auth"
	categoryRequest "github.com/sandy0786/skill-assessment-service/request/category"
	userRequest "github.com/sandy0786/skill-assessment-service/request/user"
	service "github.com/sandy0786/skill-assessment-service/service"
	transport "github.com/sandy0786/skill-assessment-service/transport"

	"github.com/go-kit/kit/endpoint"
)

//Endpoints are exposed here
//Contains endpoints that map request from client to our service
type Endpoints struct {
	StatusEndpoint              endpoint.Endpoint
	AddUserEndpoint             endpoint.Endpoint
	DeleteUserEndpoint          endpoint.Endpoint
	GetAllUsersEndpoint         endpoint.Endpoint
	RevokeUserEndpoint          endpoint.Endpoint
	ResetPasswordEndpoint       endpoint.Endpoint
	LoginEndpoint               endpoint.Endpoint
	AddQuestionEndpoint         endpoint.Endpoint
	AddMultipleQuestionEndpoint endpoint.Endpoint
	GetAllQuestionsEndpoint     endpoint.Endpoint
	AddCategoryEndpoint         endpoint.Endpoint
	UpdateCategoryEndpoint      endpoint.Endpoint
	DeleteCategoryEndpoint      endpoint.Endpoint
	GetAllCategoriesEndpoint    endpoint.Endpoint
	RefreshTokenEndpoint        endpoint.Endpoint
	ResetUserPasswordEndpoint   endpoint.Endpoint
	GetAllRolesEndpoint         endpoint.Endpoint
	GetAllQuestionTypesEndpoint endpoint.Endpoint
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
		req := request.(*http.Request)
		responseFromService, err := srv.GetAllUsers(ctx, req)
		return responseFromService, err
	}
}

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

//MakeUpdateCategoryEndpoint returns response
func MakeUpdateCategoryEndpoint(srv service.CategoryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		log.Println("MakeUpdateCategoryEndpoint")
		categoryRequest := request.(categoryDTO.UpdateCategory)
		responseFromService, err := srv.UpdateCategory(ctx, categoryRequest)
		return responseFromService, err
	}
}

//MakeDeleteCategoryEndpoint returns response
func MakeDeleteCategoryEndpoint(srv service.CategoryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		log.Println("MakeDeleteCategoryEndpoint")
		categoryRequest := request.(categoryDTO.UpdateCategory)
		responseFromService, err := srv.DeleteCategory(ctx, categoryRequest)
		return responseFromService, err
	}
}

//MakeGetAllCategoriesEndpoint returns response
func MakeGetAllCategoriesEndpoint(srv service.CategoryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		log.Println("MakeGetAllCategoriesEndpoint")
		req := request.(*http.Request)
		responseFromService, err := srv.GetAllCategories(ctx, req)
		return responseFromService, err
	}
}

//MakeDeleteUserEndpoint returns response
func MakeDeleteUserEndpoint(srv service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		log.Println("MakeDeleteUserEndpoint")
		username := request.(string)
		responseFromService, err := srv.DeleteUserByUserName(ctx, username)
		return responseFromService, err
	}
}

//MakeRevokeUserEndpoint returns response
func MakeRevokeUserEndpoint(srv service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		log.Println("MakeRevokeUserEndpoint")
		username := request.(string)
		responseFromService, err := srv.RevokeUserByUserName(ctx, username)
		return responseFromService, err
	}
}

//MakeResetPasswordEndpoint returns response
func MakeResetPasswordEndpoint(srv service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		log.Println("MakeResetPasswordEndpoint")
		userDto := request.(userDTO.UserDTO)
		responseFromService, err := srv.ResetUserPassword(ctx, userDto)
		return responseFromService, err
	}
}

//MakeLoginEndpoint returns response
func MakeLoginEndpoint(srv service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		log.Println("MakeLoginEndpoint")
		authRequest := request.(authRequest.LoginRequest)
		responseFromService, err := srv.Login(ctx, authRequest)
		return responseFromService, err
	}
}

//MakeRefreshTokenEndpoint returns response
func MakeRefreshTokenEndpoint(srv service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		log.Println("MakeRefreshTokenEndpoint")
		token := request.(string)
		responseFromService, err := srv.Refresh(ctx, token)
		return responseFromService, err
	}
}

// MakeGetAllRolesEndpoint returns response
func MakeGetAllRolesEndpoint(srv service.RoleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		log.Println("MakeGetAllRolesEndpoint")
		responseFromService, err := srv.GetAllRoles(ctx)
		return responseFromService, err
	}
}

// MakeGetAllQuestionTypesEndpoint returns response
func MakeGetAllQuestionTypesEndpoint(srv service.QuestionTypeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		log.Println("MakeGetAllQuestionTypesEndpoint")
		responseFromService, err := srv.GetAllQuestionType(ctx)
		return responseFromService, err
	}
}
