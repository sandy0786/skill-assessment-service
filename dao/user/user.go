package user

import (
	"log"

	Database "github.com/sandy0786/skill-assessment-service/database"
	userDocument "github.com/sandy0786/skill-assessment-service/documents/user"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserDAO interface {
	Save(*userDocument.User) (userDocument.User, error)
	// FindById(int64) (userModel.User, error)
	// FindAll() ([]userModel.User, error)
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

func (u *userDAOImpl) Save(user *userDocument.User) (userDocument.User, error) {
	log.Println("save employee")
	test, err := u.mongoCollection.InsertOne(u.db.GetMongoDbContext(), user)
	log.Println("test : ", test)
	// db := e.db.GetDbObject().Create(employee)
	// log.Println("> ", db.RowsAffected)
	// db = e.db.GetDbObject().Save(employee)
	// log.Println("> ", db.Value)
	// var savedEmployee model.Employee
	// savedEmployee := db.Value.(*model.Employee)
	return *user, err
}

// func (e *userDAOImpl) FindById(id int64) (model.Employee, error) {
// 	log.Println("FindById employee : ", id)
// 	var employee model.Employee
// 	// db := e.db.GetDbObject().Find(&employee, id)
// 	db := e.db.GetDbObject().Model(&employee).Preload("Address").Find(&employee, id)
// 	return employee, db.Error
// }

// func (e *userDAOImpl) FindAll() ([]model.Employee, error) {
// 	log.Println("FindAll employees")
// 	var employees []model.Employee
// 	db := e.db.GetDbObject().Find(&employees)
// 	return employees, db.Error
// }
