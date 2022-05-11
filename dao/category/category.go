package category

import (
	"log"

	Database "github.com/sandy0786/skill-assessment-service/database"
	categoryDocument "github.com/sandy0786/skill-assessment-service/documents/category"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryDAO interface {
	Save(*categoryDocument.Category) (bool, error)
	// SaveAll([]questionDocument.Question) (bool, error)
	// FindById(int64) (userModel.User, error)
	FindAll() ([]categoryDocument.Category, error)
}

type categoryDAOImpl struct {
	db              Database.DatabaseInterface
	collectionName  string
	mongoCollection *mongo.Collection
}

func NewCategoryDAO(db Database.DatabaseInterface, collectionName string) *categoryDAOImpl {
	return &categoryDAOImpl{
		db:              db,
		collectionName:  collectionName,
		mongoCollection: db.GetMongoDbObject().Collection(collectionName),
	}
}

func (q *categoryDAOImpl) Save(category *categoryDocument.Category) (bool, error) {
	log.Println("Save category")
	_, err := q.mongoCollection.InsertOne(q.db.GetMongoDbContext(), category)
	if err != nil {
		return false, err
	}
	return true, err
}

// func (q *categoryDAOImpl) SaveAll(questions []questionDocument.Question) (bool, error) {
// 	log.Println("Save all questions")
// 	tempQuestions := []interface{}{}
// 	for _, question := range questions {
// 		tempQuestions = append(tempQuestions, question)
// 	}
// 	_, err := q.mongoCollection.InsertMany(q.db.GetMongoDbContext(), tempQuestions)

// 	if err != nil {
// 		log.Println("err >> ", err)
// 		return false, err
// 	}
// 	return true, err
// }

// func (e *userDAOImpl) FindById(idint64) (model.Employee, error) {
// 	log.Println("FindById employee : ", id)
// 	var employeemodel.Employee
// 	// db := e.db.GtDbObject().Find(&employee, id)
//	db := e.db.GetDbObject().Model(&employee).Preload("Address").Find(&employee, id)
// 	return employe, db.Error
// }

func (q *categoryDAOImpl) FindAll() ([]categoryDocument.Category, error) {
	log.Println("FindAll Questions")
	var categories []categoryDocument.Category
	cursor, err := q.mongoCollection.Find(q.db.GetMongoDbContext(), bson.M{})
	if err != nil {
		return categories, err
	}
	if err = cursor.All(q.db.GetMongoDbContext(), &categories); err != nil {
		return categories, err
	}
	return categories, err
}
