package category

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var CollectionName = "categories"
var ValidatorFilePath = "./dao/init/category/CategoryValidator.json"

// index model
var CategoryNameIndex = mongo.IndexModel{
	Keys: bson.M{
		"categoryName": 1, // index in ascending order
	}, Options: options.Index().SetUnique(true),
}

// index model
var CollectionNameIndex = mongo.IndexModel{
	Keys: bson.M{
		"collectionName": 1, // index in ascending order
	}, Options: options.Index().SetUnique(true),
}
