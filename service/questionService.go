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
	questionResponse "github.com/sandy0786/skill-assessment-service/response/question"
	successResponse "github.com/sandy0786/skill-assessment-service/response/success"

	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QuestionService interface {
	AddQuestion(context.Context, questionDTO.QuestionDTO) (successResponse.SuccessResponse, error)
	AddMultipleQuestions(context.Context, questionDTO.QuestionsDTO) (successResponse.SuccessResponse, error)
	GetAllQuestions(context.Context, string) ([]questionResponse.QuestionResponse, error)
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

func (t *questionService) AddQuestion(ctx context.Context, questionDto questionDTO.QuestionDTO) (successResponse.SuccessResponse, error) {
	log.Println("Inside Add question ")
	// var userResponse userResponse.UserSuccessResponse
	var question questionDocument.Question
	copier.Copy(&question, &questionDto.Question)
	question.CreatedAt = time.Now().UTC()
	question.UpdatedAt = time.Now().UTC()
	question.ID = primitive.NewObjectID()

	// set collection name
	// t.dao = questionDao.NewQuestionDAO(t.dao.GetDbObject(), questionDto.Category)
	t.dao = t.dao.GetDaoObject(questionDto.Category)

	// save data in the provided collection
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

func (t *questionService) AddMultipleQuestions(_ context.Context, questionsDto questionDTO.QuestionsDTO) (successResponse.SuccessResponse, error) {
	log.Println("Inside Add MultipleQuestions")
	t.dao = t.dao.GetDaoObject(questionsDto.Category)
	var questions []questionDocument.Question
	for _, questionsRequest := range questionsDto.Question {
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

func (t *questionService) GetAllQuestions(_ context.Context, category string) ([]questionResponse.QuestionResponse, error) {
	log.Println("Inside GetAllQuestions")
	t.dao = t.dao.GetDaoObject(category)
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
