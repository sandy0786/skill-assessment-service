package question

type QuestionRequest struct {
	Question     string   `json:"question" validate:"required"`
	Options      []string `json:"options" validate:"required"`
	Answer       string   `json:"answer" validate:"required"`
	QuestionType string   `json:"questionType" validate:"required"`
	Author       string   `json:"author" validate:"required"`
	Deleted      bool     `json:"deleted" validate:"required"`
	DeletedBy    string   `json:"deletedBy" validate:"required"`
}
