package question

// swagger:model
type QuestionResponse struct {
	// example: How are you?
	Question string `json:"question"`
	// example: ["good", "Excellent"]
	Options []string `json:"options"`
	// example: good
	Answer string `json:"answer"`
	// example: radio
	QuestionType string `json:"questionType"`
	// example: admin
	Author string `json:"author"`
	// example: false
	Deleted bool `json:"deleted"`
	// example: admin
	DeletedBy string `json:"deletedBy"`
}

// List of questions
// swagger:response QuestionsResponse
type QuestionsResponse struct {
	// in: body
	Body []QuestionResponse
}
