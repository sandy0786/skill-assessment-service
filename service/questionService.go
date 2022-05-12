package service

import (
	"context"
	"log"
	"net/http"
	"time"

	configuration "github.com/sandy0786/skill-assessment-service/configuration"
	questionDao "github.com/sandy0786/skill-assessment-service/dao/question"
	questionDocument "github.com/sandy0786/skill-assessment-service/documents/question"
	questionDTO "github.com/sandy0786/skill-assessment-service/dto/question"
	questionRequest "github.com/sandy0786/skill-assessment-service/request/question"
	questionResponse "github.com/sandy0786/skill-assessment-service/response/question"
	successResponse "github.com/sandy0786/skill-assessment-service/response/success"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/jinzhu/copier"
)

type QuestionService interface {
	// GetServiceStatus(context.Context) (string, error)
	AddQuestion(context.Context, questionDTO.QuestionDTO) (successResponse.SuccessResponse, error)
	AddMultipleQuestions(context.Context, []questionRequest.QuestionRequest) (successResponse.SuccessResponse, error)
	GetAllQuestions(context.Context) ([]questionResponse.QuestionResponse, error)
	// GetEmployeeById(context.Context, int64) (employeeResponse.EmployeeResponse, error)
}

// service for druid
type questionService struct {
	questionServiceConfig configuration.ConfigurationInterface
	dao                   questionDao.QuestionDAO
}

func NewQuestionService(c configuration.ConfigurationInterface, dao questionDao.QuestionDAO) *questionService {
	return &questionService{
		questionServiceConfig: c,
		dao:                   dao,
	}
}

func (t *questionService) AddQuestion(_ context.Context, questionDto questionDTO.QuestionDTO) (successResponse.SuccessResponse, error) {
	log.Println("Inside Add question : ", questionDto.Category)
	// var userResponse userResponse.UserSuccessResponse
	var question questionDocument.Question
	copier.Copy(&question, &questionDto.questionRequest)
	question.CreatedAt = time.Now().UTC()
	question.UpdatedAt = time.Now().UTC()
	question.ID = primitive.NewObjectID()
	questionCreated, err := t.dao.Save(&question)
	// copier.Copy(&userResponse, &userRequest)
	if questionCreated {
		return successResponse.SuccessResponse{
			Status:    http.StatusOK,
			Message:   "Questions submitted successfully",
			TimeStamp: time.Now().UTC().String(),
		}, err
	}
	return successResponse.SuccessResponse{}, err
}

func (t *questionService) AddMultipleQuestions(_ context.Context, questionsRequests []questionRequest.QuestionRequest) (successResponse.SuccessResponse, error) {
	log.Println("Inside Add MultipleQuestions")
	var questions []questionDocument.Question
	for _, questionsRequest := range questionsRequests {
		var question questionDocument.Question
		copier.Copy(&question, &questionsRequest)
		question.CreatedAt = time.Now().UTC()
		question.UpdatedAt = time.Now().UTC()
		question.ID = primitive.NewObjectID()
		questions = append(questions, question)
	}

	questionsCreated, err := t.dao.SaveAll(questions)

	if questionsCreated {
		return successResponse.SuccessResponse{
			Status:    http.StatusOK,
			Message:   "Questions submitted successfully",
			TimeStamp: time.Now().UTC().String(),
		}, err
	}
	return successResponse.SuccessResponse{}, err
}

func (t *questionService) GetAllQuestions(_ context.Context) ([]questionResponse.QuestionResponse, error) {
	log.Println("Inside GetAllQuestions")
	var questionResponses []questionResponse.QuestionResponse
	questions, err := t.dao.FindAll()
	copier.Copy(&questionResponses, &questions)
	return questionResponses, err
}

// func (t *userService) GetEmployeeById(_ context.Context, id int64) (employeeResponse.EmployeeResponse, error) {
// 	log.Println("Inside GetEmployeeById : ", id)
// 	var employeeResponse employeeResponse.EmployeeResponse
// 	employee, err := t.dao.FindById(id)
// 	copier.Copy(&employeeResponse, &employee)
// 	return employeeResponse, err
// }
