package questionType

import (
	"time"

	questionModel "github.com/sandy0786/skill-assessment-service/documents/questionType"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var QuestionTypeCollectionName = "questionType"
var QuestionTypeValidatorFilePath = "./dao/init/questionType/questionTypeValidator.json"

// index model
var QuestionTypeIndexRoleName = mongo.IndexModel{
	Keys: bson.M{
		"questionType": 1, // index in ascending order
	}, Options: options.Index().SetUnique(true),
}

var CheckboxQT = questionModel.QuestionType{
	ID:           primitive.NewObjectID(),
	CreatedAt:    time.Now().UTC(),
	UpdatedAt:    time.Now().UTC(),
	QuestionType: "checkbox",
}

var RadioQT = questionModel.QuestionType{
	ID:           primitive.NewObjectID(),
	CreatedAt:    time.Now().UTC(),
	UpdatedAt:    time.Now().UTC(),
	QuestionType: "radio",
}
