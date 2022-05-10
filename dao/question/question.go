package question

import (
	"log"

	Database "github.com/sandy0786/skill-assessment-service/database"
	questionDocument "github.com/sandy0786/skill-assessment-service/documents/question"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type QuestionDAO interface {
	Save(*questionDocument.Question) (bool, error)
	// FindById(int64) (userModel.User, error)
	FindAll() ([]questionDocument.Question, error)
}

type questionDAOImpl struct {
	db              Database.DatabaseInterface
	collectionName  string
	mongoCollection *mongo.Collection
}

func NewQuestionDAO(db Database.DatabaseInterface, collectionName string) *questionDAOImpl {
	return &questionDAOImpl{
		db:              db,
		collectionName:  collectionName,
		mongoCollection: db.GetMongoDbObject().Collection(collectionName),
	}
}

func (q *questionDAOImpl) Save(user *questionDocument.Question) (bool, error) {
	log.Println("Save questions")
	_, err := q.mongoCollection.InsertOne(q.db.GetMongoDbContext(), user)
	if err != nil {
		return false, err
	}
	return true, err
}

// func (e *userDAOImpl) FindById(idint64) (model.Employee, error) {
// 	log.Println("FindById employee : ", id)
// 	var employeemodel.Employee
// 	// db := e.db.GtDbObject().Find(&employee, id)
//	db := e.db.GetDbObject().Model(&employee).Preload("Address").Find(&employee, id)
// 	return employe, db.Error
// }

func (q *questionDAOImpl) FindAll() ([]questionDocument.Question, error) {
	log.Println("FindAll Questions")
	var questions []questionDocument.Question
	cursor, err := q.mongoCollection.Find(q.db.GetMongoDbContext(), bson.M{})
	if err != nil {
		log.Println("error , ", err)
		return questions, err
	}
	if err = cursor.All(q.db.GetMongoDbContext(), &questions); err != nil {
		log.Println("error , ", err)
		return questions, err
	}
	return questions, err
}
