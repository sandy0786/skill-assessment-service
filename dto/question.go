package question

import (
	questionRequest "github.com/sandy0786/skill-assessment-service/request/question"
)

type QuestionDTO struct {
	Category string
	Question questionRequest.QuestionRequest
}
