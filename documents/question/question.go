package question

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Question struct {
	ID           primitive.ObjectID `bson:"_id"`
	CreatedAt    time.Time          `bson:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt"`
	Question     string             `bson:"question"`
	Options      []string           `bson:"options"`
	Answer       string             `bson:"answer"`
	QuestionType string             `bson:"questionType"`
	Author       string             `bson:"author"`
	Deleted      bool               `bson:"deleted"`
	DeletedBy    string             `bson:"deletedBy"`
}
