package user

import (
	"log"

	Database "github.com/sandy0786/skill-assessment-service/database"
	userDocument "github.com/sandy0786/skill-assessment-service/documents/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserDAO interface {
	Save(*userDocument.User) (bool, error)
	// FindById(int64) (userModel.User, error)
	FindAll() ([]userDocument.User, error)
}

type userDAOImpl struct {
	db              Database.DatabaseInterface
	collectionName  string
	mongoCollection *mongo.Collection
}

func NewUserDAO(db Database.DatabaseInterface, collectionName string) *userDAOImpl {
	return &userDAOImpl{
		db:              db,
		collectionName:  collectionName,
		mongoCollection: db.GetMongoDbObject().Collection(collectionName),
	}
}

func (u *userDAOImpl) Save(user *userDocument.User) (bool, error) {
	log.Println("save user")
	_, err := u.mongoCollection.InsertOne(u.db.GetMongoDbContext(), user)
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

func (u *userDAOImpl) FindAll() ([]userDocument.User, error) {
	log.Println("FindAll Users")
	var users []userDocument.User
	cursor, err := u.mongoCollection.Find(u.db.GetMongoDbContext(), bson.M{})
	if err != nil {
		log.Println("error , ", err)
		return users, err
	}
	if err = cursor.All(u.db.GetMongoDbContext(), &users); err != nil {
		log.Println("error , ", err)
		return users, err
	}
	return users, err
}
