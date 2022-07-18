package service

import (
	"context"
	"log"

	configuration "github.com/sandy0786/skill-assessment-service/configuration"
	questionTypeDao "github.com/sandy0786/skill-assessment-service/dao/questionType"
	questionTypeResponse "github.com/sandy0786/skill-assessment-service/response/questionType"

	"github.com/jinzhu/copier"
)

type QuestionTypeService interface {
	GetAllQuestionType(context.Context) ([]questionTypeResponse.QuestionTypeResponse, error)
}

type questionTypeService struct {
	config configuration.ConfigurationInterface
	dao    questionTypeDao.QuestionTypeDAO
}

func NewQuestionTypeService(c configuration.ConfigurationInterface, dao questionTypeDao.QuestionTypeDAO) *questionTypeService {
	return &questionTypeService{
		config: c,
		dao:    dao,
	}
}

func (r *questionTypeService) GetAllQuestionType(context.Context) ([]questionTypeResponse.QuestionTypeResponse, error) {
	log.Println("Inside GetAllQuestionType")
	var questionTypesResponse []questionTypeResponse.QuestionTypeResponse
	questionTypes, err := r.dao.GetAllQuestionTypes()
	copier.Copy(&questionTypesResponse, &questionTypes)
	return questionTypesResponse, err
}
