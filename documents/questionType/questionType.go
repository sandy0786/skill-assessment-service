package questionType

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QuestionType struct {
	ID           primitive.ObjectID `bson:"_id"`
	CreatedAt    time.Time          `bson:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt"`
	QuestionType string             `bson:"questionType"`
}
