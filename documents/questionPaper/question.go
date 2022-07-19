package questionPaper

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Question struct {
	Category string
	Question primitive.ObjectID
}

type QuestionPaper struct {
	ID          primitive.ObjectID `bson:"_id"`
	CreatedAt   time.Time          `bson:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt"`
	Name        string             `bson:"name"`
	Questions   []Question         `bson:"questions"`
	TotalMarks  float32            `bson:"totalMarks"`
	TotalScore  float32            `bson:"totalScore"`
	Description string             `bson:"description"`
	Author      string             `bson:"author"`
}
