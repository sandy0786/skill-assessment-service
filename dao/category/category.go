package category

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	Database "github.com/sandy0786/skill-assessment-service/database"
	categoryDocument "github.com/sandy0786/skill-assessment-service/documents/category"
	categoryDTO "github.com/sandy0786/skill-assessment-service/dto/category"
	err "github.com/sandy0786/skill-assessment-service/errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CategoryDAO interface {
	Save(*categoryDocument.Category) (bool, error)
	// SaveAll([]questionDocument.Question) (bool, error)
	// FindById(int64) (userModel.User, error)
	FindAll(categoryDTO.Pagination) ([]categoryDocument.Category, error)
	// CreateCollection(string) (*categoryDAOImpl, error)
	GetCount(string) (int64, error)
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
	_, writeError := q.mongoCollection.InsertOne(q.db.GetMongoDbContext(), category)
	if writeError != nil {
		writeException, ok := writeError.(mongo.WriteException)
		// handle mongo errors
		if ok {
			var errMessage string
			switch writeException.WriteErrors[0].Code {
			case 11000: // duplicate error
				if strings.Contains(writeException.WriteErrors[0].Error(), "category") {
					errMessage = "category already exists "
				} else {
					errMessage = writeException.WriteErrors[0].Error()
				}
				return false, err.GlobalError{
					TimeStamp: time.Now().UTC().String()[0:19],
					Status:    http.StatusConflict,
					Message:   errMessage,
				}
			}
		}
	}
	return true, nil
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

func (q *categoryDAOImpl) FindAll(pagination categoryDTO.Pagination) ([]categoryDocument.Category, error) {
	log.Println("FindAll Questions")
	var categories []categoryDocument.Category

	opts := &options.FindOptions{
		Limit: pagination.Length,
		Skip:  pagination.Start,
		Sort:  bson.M{"category": pagination.OrderBy},
	}

	cursor, err := q.mongoCollection.Find(q.db.GetMongoDbContext(), bson.M{"category": bson.M{"$regex": pagination.Search}}, opts)
	if err != nil {
		return categories, err
	}
	if err = cursor.All(q.db.GetMongoDbContext(), &categories); err != nil {
		return categories, err
	}
	return categories, err
}

func (c *categoryDAOImpl) CreateCollection(collectionName string) (*categoryDAOImpl, error) {
	ctx := context.TODO()
	err := c.db.GetMongoDbObject().CreateCollection(ctx, collectionName)
	if err != nil {
		log.Println("Error while creating collection : ", err)
		return nil, err
	}
	// c.mongoCollection = c.mongoCollection.Database().Collection(collectionName)
	return c, err
}

func (u *categoryDAOImpl) GetCount(category string) (int64, error) {
	log.Println("GetCount")
	totalCount, err := u.mongoCollection.CountDocuments(u.db.GetMongoDbContext(), bson.M{"category": bson.M{"$regex": category}})
	if err != nil {
		return totalCount, err
	}
	return totalCount, err
}
