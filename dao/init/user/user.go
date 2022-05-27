package user

import (
	"go.mongodb.org/mongo-driver/bson"
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

// var MongoIndexes []mongo.IndexModel
// MongoIndexes = (MongoIndexes, userIndexUsername)

// var Indexes = []mongo.IndexModel [
// 	userIndexUsername,
// 	userIndexPassword
// ]
