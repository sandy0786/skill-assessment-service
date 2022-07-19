package auth

import (
	"log"

	Database "github.com/sandy0786/skill-assessment-service/database"
	authDocument "github.com/sandy0786/skill-assessment-service/documents/auth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthDAO interface {
	Login(*authDocument.User) (authDocument.User, bool, error)
}

type authDAOImpl struct {
	db              Database.DatabaseInterface
	collectionName  string
	mongoCollection *mongo.Collection
}

func NewAuthDAO(db Database.DatabaseInterface, collectionName string) *authDAOImpl {
	return &authDAOImpl{
		db:              db,
		collectionName:  collectionName,
		mongoCollection: db.GetMongoDbObject().Collection(collectionName),
	}
}

func (u *authDAOImpl) Login(user *authDocument.User) (authDocument.User, bool, error) {
	log.Println("user login")
	cursor := u.mongoCollection.FindOne(u.db.GetMongoDbContext(), bson.M{"username": user.Username, "password": user.Password})
	var validUser authDocument.User
	err := cursor.Decode(&validUser)
	if err != nil {
		// if err.Error() == "mongo: no documents in result"
		log.Println("error , ", err)
		return validUser, false, err
	}

	return validUser, true, nil
}
