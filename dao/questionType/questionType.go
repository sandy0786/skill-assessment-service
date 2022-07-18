package questionType

import (
	"log"

	Database "github.com/sandy0786/skill-assessment-service/database"
	questionTypeDocument "github.com/sandy0786/skill-assessment-service/documents/questionType"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type QuestionTypeDAO interface {
	GetAllQuestionTypes() ([]questionTypeDocument.QuestionType, error)
}

type questionTypeDAOImpl struct {
	db              Database.DatabaseInterface
	collectionName  string
	mongoCollection *mongo.Collection
}

func NewQuestionTypeDAO(db Database.DatabaseInterface, collectionName string) *questionTypeDAOImpl {
	return &questionTypeDAOImpl{
		db:              db,
		collectionName:  collectionName,
		mongoCollection: db.GetMongoDbObject().Collection(collectionName),
	}
}

func (q *questionTypeDAOImpl) GetAllQuestionTypes() ([]questionTypeDocument.QuestionType, error) {
	log.Println("Get all question types")
	var questionTypes []questionTypeDocument.QuestionType
	cursor, err := q.mongoCollection.Find(q.db.GetMongoDbContext(), bson.M{})
	if err != nil {
		return questionTypes, err
	}
	if err = cursor.All(q.db.GetMongoDbContext(), &questionTypes); err != nil {
		return questionTypes, err
	}
	return questionTypes, err
}
