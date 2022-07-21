package user

import (
	"log"
	"net/http"
	"strings"
	"time"

	Database "github.com/sandy0786/skill-assessment-service/database"
	userDocument "github.com/sandy0786/skill-assessment-service/documents/user"
	err "github.com/sandy0786/skill-assessment-service/errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserDAO interface {
	Save(*userDocument.User) (bool, error)
	FindAll(*int64, *int64, string, string, int) ([]userDocument.User, error)
	DeleteByUserName(string) (bool, error)
	RevokeByUserName(string) (bool, error)
	ResetUserPassword(string, string, string) (bool, error)
	GetCount(string) (int64, error)
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
	_, writeError := u.mongoCollection.InsertOne(u.db.GetMongoDbContext(), user)
	if writeError != nil {
		writeException, ok := writeError.(mongo.WriteException)
		// handle mongo errors
		if ok {
			var errMessage string
			switch writeException.WriteErrors[0].Code {
			case 11000: // duplicate error
				if strings.Contains(writeException.WriteErrors[0].Error(), "email") {
					errMessage = "Email already exists"
				} else if strings.Contains(writeException.WriteErrors[0].Error(), "username") {
					errMessage = "Username already exists"
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

func (u *userDAOImpl) FindAll(start *int64, length *int64, username string, sortBy string, orderBy int) ([]userDocument.User, error) {
	log.Println("FindAll Users")
	var users []userDocument.User

	sort := bson.M{}
	if sortBy != "" && orderBy != 0 {
		sort = bson.M{sortBy: orderBy}
	}

	opts := &options.FindOptions{
		Limit: length,
		Skip:  start,
		Sort:  sort,
	}
	cursor, err := u.mongoCollection.Find(u.db.GetMongoDbContext(), bson.M{"username": bson.M{"$regex": username}}, opts)
	if err != nil {
		return users, err
	}
	if err = cursor.All(u.db.GetMongoDbContext(), &users); err != nil {
		return users, err
	}
	return users, err
}

func (u *userDAOImpl) DeleteByUserName(username string) (bool, error) {
	log.Println("Delete user")
	// _, err := u.mongoCollection.DeleteOne(u.db.GetMongoDbContext(), bson.M{"username": username})
	_, err := u.mongoCollection.UpdateOne(u.db.GetMongoDbContext(), bson.M{"username": username},
		bson.D{
			{"$set", bson.D{{"active", false}}},
		},
	)

	if err != nil {
		return false, err
	}
	return true, err
}

func (u *userDAOImpl) RevokeByUserName(username string) (bool, error) {
	log.Println("Revoke user")
	_, err := u.mongoCollection.UpdateOne(u.db.GetMongoDbContext(), bson.M{"username": username},
		bson.D{
			{"$set", bson.D{{"active", true}}},
		},
	)
	if err != nil {
		return false, err
	}
	return true, err
}

func (u *userDAOImpl) ResetUserPassword(username string, oldpassword string, newPassword string) (bool, error) {
	log.Println("Reset user password")
	_, err := u.mongoCollection.UpdateOne(u.db.GetMongoDbContext(), bson.M{"username": username, "password": oldpassword},
		bson.D{
			{"$set", bson.D{{"password", newPassword}}},
		},
	)

	if err != nil {
		return false, err
	}
	return true, err
}

func (u *userDAOImpl) GetCount(username string) (int64, error) {
	log.Println("GetCount")
	totalCount, err := u.mongoCollection.CountDocuments(u.db.GetMongoDbContext(), bson.M{"username": bson.M{"$regex": username}})
	if err != nil {
		return totalCount, err
	}
	return totalCount, err
}
