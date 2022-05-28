package init

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	user "github.com/sandy0786/skill-assessment-service/dao/init/user"
	database "github.com/sandy0786/skill-assessment-service/database"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoMetadata []struct {
	collection        string
	validatorFilePath string
	indexes           []mongo.IndexModel
	documents         []interface{}
}

var mongoMetadataObj mongoMetadata

func initializeMongoMetadataObject() {
	mongoMetadataObj = []struct {
		collection        string
		validatorFilePath string
		indexes           []mongo.IndexModel
		documents         []interface{}
	}{
		{
			collection:        "users",
			validatorFilePath: user.UserValidatorFilePath,
			indexes: []mongo.IndexModel{
				user.UserIndexEmail,
				user.UserIndexUsername,
			},
			documents: []interface{}{
				user.AdminUser,
			},
		},
	}
}

func InitMongoDBCollections(db database.DatabaseInterface) {
	log.Println("InitMongoDBCollections")

	// Initialize mongo metadata
	initializeMongoMetadataObject()

	// Create new context object
	ctx := context.TODO()

	// loop through mongo metadata
	for _, metadata := range mongoMetadataObj {

		// read validator file
		readFile, readFileErr := ioutil.ReadFile(metadata.validatorFilePath)
		if readFileErr != nil {
			log.Println("Cannot read UsersValidator file ", readFileErr)
		}

		var validatorObj interface{}
		// Unmarshall to object
		_ = json.Unmarshal(readFile, &validatorObj)

		// Create mongo options
		validationLevel := "strict"
		collectionOptions := &options.CreateCollectionOptions{
			Validator:       validatorObj,
			ValidationLevel: &validationLevel,
		}

		// create collection with respective validator
		db.GetMongoDbObject().CreateCollection(ctx, metadata.collection, collectionOptions)

		// Get collection object
		mongoCollection := db.GetMongoDbObject().Collection(metadata.collection)

		// Create indexes
		for _, index := range metadata.indexes {
			mongoCollection.Indexes().CreateOne(ctx, index)
		}

		// Insert required documents
		for _, document := range metadata.documents {
			mongoCollection.InsertOne(ctx, document)
		}

	}
}
