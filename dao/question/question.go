package question

import (
	"context"
	"log"

	Database "github.com/sandy0786/skill-assessment-service/database"
	questionDocument "github.com/sandy0786/skill-assessment-service/documents/question"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type QuestionDAO interface {
	Save(*questionDocument.Question) (bool, error)
	SaveAll([]questionDocument.Question) (bool, error)
	// FindById(int64) (userModel.User, error)
	FindAll() ([]questionDocument.Question, error)
	SetCollectionName(string)
	GetCollectionObject(string) *mongo.Collection
	GetDbObject() Database.DatabaseInterface
	CreateCollection(string) *questionDAOImpl
	GetDaoObject(string) *questionDAOImpl
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

func (q *questionDAOImpl) Save(question *questionDocument.Question) (bool, error) {
	log.Println("Save questions")
	_, err := q.mongoCollection.InsertOne(q.db.GetMongoDbContext(), question)
	if err != nil {
		return false, err
	}
	return true, err
}

func (q *questionDAOImpl) SaveAll(questions []questionDocument.Question) (bool, error) {
	log.Println("Save all questions")
	tempQuestions := []interface{}{}
	for _, question := range questions {
		tempQuestions = append(tempQuestions, question)
	}
	_, err := q.mongoCollection.InsertMany(q.db.GetMongoDbContext(), tempQuestions)

	if err != nil {
		log.Println("err >> ", err)
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

func (q *questionDAOImpl) SetCollectionName(collectionName string) {
	q.collectionName = collectionName
	q.mongoCollection = q.mongoCollection.Database().Collection(collectionName)
}

func (q *questionDAOImpl) GetCollectionObject(collectionName string) *mongo.Collection {
	q.collectionName = collectionName
	// err := q.mongoCollection.Database().CreateCollection(collectionName)
	return q.mongoCollection.Database().Collection(collectionName)
}

func (q *questionDAOImpl) GetDbObject() Database.DatabaseInterface {
	return q.db
}

func (q *questionDAOImpl) CreateCollection(collectionName string) *questionDAOImpl {
	ctx := context.TODO()
	err := q.db.GetMongoDbObject().CreateCollection(ctx, collectionName)
	if err != nil {
		log.Println("Error while creating collection : ", err)
	}
	q.mongoCollection = q.mongoCollection.Database().Collection(collectionName)
	return q
}

func (q *questionDAOImpl) GetDaoObject(collectionName string) *questionDAOImpl {
	if q.collectionName != collectionName {
		return &questionDAOImpl{
			db:              q.db,
			collectionName:  collectionName,
			mongoCollection: q.db.GetMongoDbObject().Collection(collectionName),
		}
	} else {
		return q
	}

}
