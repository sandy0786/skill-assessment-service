package init

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	category "github.com/sandy0786/skill-assessment-service/dao/init/category"
	questionType "github.com/sandy0786/skill-assessment-service/dao/init/questionType"
	role "github.com/sandy0786/skill-assessment-service/dao/init/role"
	user "github.com/sandy0786/skill-assessment-service/dao/init/user"
	database "github.com/sandy0786/skill-assessment-service/database"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoMetadata []struct {
	collection        string
	validatorFilePath string
	indexes           []mongo.IndexModel
	documents         []interface{}
}

func init() {
	role.AdminRole.ID = role.AdminObjId
	user.AdminUser.Role = role.AdminObjId
}

// var mongoMetadataObj MongoMetadata

func InitializeMongoMetadataObject() MongoMetadata {
	mongoMetadataObj := []struct {
		collection        string
		validatorFilePath string
		indexes           []mongo.IndexModel
		documents         []interface{}
	}{
		{
			collection:        role.RoleCollectionName,
			validatorFilePath: role.RoleValidatorFilePath,
			indexes: []mongo.IndexModel{
				role.RoleIndexRoleName,
			},
			documents: []interface{}{
				role.AdminRole,
				role.ManagerRole,
				role.GuestRole,
			},
		},
		{
			collection:        user.UserCollectionName,
			validatorFilePath: user.UserValidatorFilePath,
			indexes: []mongo.IndexModel{
				user.UserIndexEmail,
				user.UserIndexUsername,
			},
			documents: []interface{}{
				user.AdminUser,
			},
		},
		{
			collection:        category.CollectionName,
			validatorFilePath: category.ValidatorFilePath,
			indexes: []mongo.IndexModel{
				category.CategoryNameIndex,
				category.CollectionNameIndex,
			},
			documents: []interface{}{},
		},
		{
			collection:        questionType.QuestionTypeCollectionName,
			validatorFilePath: questionType.QuestionTypeValidatorFilePath,
			indexes: []mongo.IndexModel{
				questionType.QuestionTypeIndexRoleName,
			},
			documents: []interface{}{
				questionType.CheckboxQT,
				questionType.RadioQT,
			},
		},
	}

	return mongoMetadataObj
}

func InitMongoDBCollections(db database.DatabaseInterface, mongoMetadataObj MongoMetadata) {
	log.Println("InitMongoDBCollections")

	// Initialize mongo metadata
	// initializeMongoMetadataObject()

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
			_, _ = mongoCollection.InsertOne(ctx, document)
		}

	}
}
