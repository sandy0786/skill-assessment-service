package question

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Question struct {
	ID           primitive.ObjectID `bson:"_id"`
	// ID           int64     `bson:"_id"`
	CreatedAt    time.Time `bson:"createdAt"`
	UpdatedAt    time.Time `bson:"updatedAt"`
	Question     string    `bson:"question"`
	Options      []string  `bson:"options"`
	Answer       string    `bson:"answer"`
	QuestionType string    `bson:"questionType"`
	Marks        float32       `bson:"marks"`
	Author       string    `bson:"author"`
	Deleted      bool      `bson:"deleted"`
	DeletedBy    string    `bson:"deletedBy"`
}
