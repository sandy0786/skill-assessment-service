package user

import (
	"time"

	userModel "github.com/sandy0786/skill-assessment-service/documents/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var UserCollectionName = "users"
var UserValidatorFilePath = "./dao/init/user/UsersValidator.json"

// index model
var UserIndexUsername = mongo.IndexModel{
	Keys: bson.M{
		"username": 1, // index in ascending order
	}, Options: options.Index().SetUnique(true),
}

// index model
var UserIndexEmail = mongo.IndexModel{
	Keys: bson.M{
		"email": 1, // index in ascending order
	}, Options: options.Index().SetUnique(true),
}

var objId primitive.ObjectID

func init() {
	objId, _ = primitive.ObjectIDFromHex("62d6435f333f27963c162e02")
}

var AdminUser = userModel.User{
	ID:        primitive.NewObjectID(),
	CreatedAt: time.Now().UTC(),
	UpdatedAt: time.Now().UTC(),
	Username:  "admin",
	Password:  "admin@123",
	Email:     "admin@admin.com",
	Role:      objId,
	// Role:   "admin",
	Active: true,
}
