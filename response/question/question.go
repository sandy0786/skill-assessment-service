package question

type QuestionResponse struct {
	Question     string   `json:"username"`
	Options      []string `json:"options"`
	Answer       string   `json:"answer"`
	QuestionType string   `json:"questionType"`
	Author       string   `json:"author"`
}
