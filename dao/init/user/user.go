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

var AdminUser = userModel.User{
	ID:        primitive.ObjectID{},
	CreatedAt: time.Now().UTC(),
	UpdatedAt: time.Now().UTC(),
	Username:  "admin",
	Password:  "admin@123",
	Email:     "admin@admin.com",
	Role:      "admin",
	Active:    true,
}
