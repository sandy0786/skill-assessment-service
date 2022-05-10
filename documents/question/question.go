package question

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Question struct {
	ID           primitive.ObjectID `bson:"_id"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
	Question     string             `bson:"question"`
	Options      []string           `bson:"options"`
	Answer       string             `bson:"answer"`
	QuestionType string             `bson:"questionType"`
	Author       string             `bson:"author"`
}
