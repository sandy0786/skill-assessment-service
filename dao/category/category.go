package category

import (
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
	UpdateCategoryById(*categoryDocument.Category) (bool, error)
	FindAll(categoryDTO.Pagination) ([]categoryDocument.Category, error)
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

func (c *categoryDAOImpl) UpdateCategoryById(category *categoryDocument.Category) (bool, error) {
	log.Println("Update Category by id ")

	update := bson.M{"$set": bson.M{"category": category.Category, "updatedAt": category.UpdatedAt}}
	updateResult, UpdateErr := c.mongoCollection.UpdateOne(c.db.GetMongoDbContext(), bson.M{"_id": category.ID}, update)

	if UpdateErr != nil {
		writeException, ok := UpdateErr.(mongo.WriteException)
		// handle mongo errors
		if ok {
			var errMessage string
			switch writeException.WriteErrors[0].Code {
			case 11000: // duplicate error
				if strings.Contains(writeException.WriteErrors[0].Error(), "category") {
					errMessage = "Category already exists "
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
	} else if updateResult.MatchedCount == 0 {
		return false, err.GlobalError{
			Message:   "No matching category found",
			TimeStamp: time.Now().UTC().String()[0:19],
			Status:    http.StatusNotFound,
		}
	}

	return true, nil
}

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

func (u *categoryDAOImpl) GetCount(category string) (int64, error) {
	log.Println("GetCount")
	totalCount, err := u.mongoCollection.CountDocuments(u.db.GetMongoDbContext(), bson.M{"category": bson.M{"$regex": category}})
	if err != nil {
		return totalCount, err
	}
	return totalCount, err
}
