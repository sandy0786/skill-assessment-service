package questionType

// swagger:model
type QuestionTypeResponse struct {
	// example: radio
	QuestionType string `json:"questionType"`
}

// List of questions types
// swagger:response QuestionTypeResponse
type QuestionsResponse struct {
	// in: body
	Body []QuestionTypeResponse
}
