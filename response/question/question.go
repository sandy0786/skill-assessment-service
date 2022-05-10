package question

type QuestionResponse struct {
	Question     string   `json:"username"`
	Options      []string `json:"options"`
	Answer       string   `json:"answer"`
	QuestionType string   `json:"questionType"`
	Author       string   `json:"author"`
}

// Questions success created
type QuestionSuccessResponse struct {
	TimeStamp string `json:"timestamp"`
	Status    int    `json:"status"`
	Message   string `json:"message"`
}
