package service

import (
	"context"
	"log"
	"net/http"
	"time"

	configuration "github.com/sandy0786/skill-assessment-service/configuration"
	categoryDao "github.com/sandy0786/skill-assessment-service/dao/category"
	categoryDocument "github.com/sandy0786/skill-assessment-service/documents/category"
	categoryRequest "github.com/sandy0786/skill-assessment-service/request/category"
	categoryResponse "github.com/sandy0786/skill-assessment-service/response/category"
	successResponse "github.com/sandy0786/skill-assessment-service/response/success"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/jinzhu/copier"
)

type CategoryService interface {
	// GetServiceStatus(context.Context) (string, error)
	AddCategory(context.Context, categoryRequest.CategoryRequest) (successResponse.SuccessResponse, error)
	// AddMultipleQuestions(context.Context, []categoryRequest.CategoryRequest) (successResponse.SuccessResponse, error)
	GetAllCategories(context.Context) ([]categoryResponse.CategoryResponse, error)
	// GetEmployeeById(context.Context, int64) (employeeResponse.EmployeeResponse, error)
}

// service for druid
type categoryService struct {
	questionServiceConfig configuration.ConfigurationInterface
	dao                   categoryDao.CategoryDAO
}

func NewCategoryService(c configuration.ConfigurationInterface, dao categoryDao.CategoryDAO) *categoryService {
	return &categoryService{
		questionServiceConfig: c,
		dao:                   dao,
	}
}

func (t *categoryService) AddCategory(_ context.Context, categoryRequest categoryRequest.CategoryRequest) (successResponse.SuccessResponse, error) {
	log.Println("Inside Add Category")
	var category categoryDocument.Category
	copier.Copy(&category, &categoryRequest)
	category.CreatedAt = time.Now().UTC()
	category.UpdatedAt = time.Now().UTC()
	category.ID = primitive.NewObjectID()
	categoryCreated, err := t.dao.Save(&category)
	if categoryCreated {
		return successResponse.SuccessResponse{
			Status:    http.StatusOK,
			Message:   "Category submitted successfully",
			TimeStamp: time.Now().UTC().String(),
		}, err
	}
	return successResponse.SuccessResponse{}, err
}

// func (t *categoryService) AddMultipleQuestions(_ context.Context, questionsRequests []questionRequest.QuestionRequest) (successResponse.SuccessResponse, error) {
// 	log.Println("Inside Add MultipleQuestions")
// 	var questions []questionDocument.Question
// 	for _, questionsRequest := range questionsRequests {
// 		var question questionDocument.Question
// 		copier.Copy(&question, &questionsRequest)
// 		question.CreatedAt = time.Now().UTC()
// 		question.UpdatedAt = time.Now().UTC()
// 		question.ID = primitive.NewObjectID()
// 		questions = append(questions, question)
// 	}

// 	questionsCreated, err := t.dao.SaveAll(questions)

// 	if questionsCreated {
// 		return successResponse.SuccessResponse{
// 			Status:    http.StatusOK,
// 			Message:   "Questions submitted successfully",
// 			TimeStamp: time.Now().UTC().String(),
// 		}, err
// 	}
// 	return successResponse.SuccessResponse{}, err
// }

func (t *categoryService) GetAllCategories(_ context.Context) ([]categoryResponse.CategoryResponse, error) {
	log.Println("Inside GetAllCategories")
	var categoryResponses []categoryResponse.CategoryResponse
	categorys, err := t.dao.FindAll()
	copier.Copy(&categoryResponses, &categorys)
	return categoryResponses, err
}
