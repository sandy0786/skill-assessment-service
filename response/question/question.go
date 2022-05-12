package question

import "go.mongodb.org/mongo-driver/bson/primitive"

type QuestionResponse struct {
	ID           primitive.ObjectID `json:"_id"`
	Question     string             `json:"username"`
	Options      []string           `json:"options"`
	Answer       string             `json:"answer"`
	QuestionType string             `json:"questionType"`
	Author       string             `json:"author"`
	Deleted      bool               `json:"deleted"`
	DeletedBy    string             `json:"deletedBy"`
}
