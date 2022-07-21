package question

// swagger:model
type QuestionRequest struct {
	// Provide Question
	// required: true
	// example: How are you
	Question string `json:"question" validate:"required"`
	// provide options to choose
	// required: true
	// example: ["good", "excellent"]
	Options []string `json:"options" validate:"required"`
	// Provide correct answer
	// required: true
	// example: good
	Answer string `json:"answer" validate:"required"`
	// Provide Question type
	// required: true
	// example: checkbox/radio
	QuestionType string `json:"questionType" validate:"required"`
	// Provide marks for this question
	// required: true
	// example: 1.0
	Marks float32 `json:"marks" validate:"required"`
	// Provide username of the user who has created this question
	// required: true
	// example: admin
	Author string `json:"author" validate:"required"`
	// Is this question deleted?
	// example: false
	Deleted bool `json:"deleted"`
	// Provide username of the user who has deleted this question
	// example: admin
	DeletedBy string `json:"deletedBy"`
}

// swagger:parameters QuestionRequest
type QuestionRequestSwagger struct {
	// Provide category
	// in: path
	// required: true
	// example: go
	Category string
	// in:body
	// required: true
	Body QuestionRequest
}

// swagger:parameters QuestionsRequest
type QuestionsRequestSwagger struct {
	// in: path
	// required: true
	// example: go
	Category string
	// in:body
	// required: true
	Body []QuestionRequest
}

// swagger:parameters GetQuestionRequest
type GetQuestionRequestSwagger struct {
	// in: path
	// required: true
	// example: go
	Category string
}
