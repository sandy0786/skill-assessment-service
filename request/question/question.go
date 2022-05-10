package question

type QuestionRequest struct {
	Question     string   `validate:"required"`
	Options      []string `validate:"required"`
	Answer       string   `validate:"required"`
	QuestionType string   `validate:"required"`
	Author       string   `validate:"required"`
}
