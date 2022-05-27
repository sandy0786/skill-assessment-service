package init

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	user "github.com/sandy0786/skill-assessment-service/dao/init/user"
	database "github.com/sandy0786/skill-assessment-service/database"
	userModel "github.com/sandy0786/skill-assessment-service/documents/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDBCollections(db database.DatabaseInterface) {
	log.Println("InitMongoDBCollections")
	ctx := context.TODO()
	readFile, readFileErr := ioutil.ReadFile(user.UserValidatorFilePath)
	if readFileErr != nil {
		log.Println("Cannot read UsersValidator file ", readFileErr)
	}

	var validatorObj interface{}
	_ = json.Unmarshal(readFile, &validatorObj)

	validationLevel := "strict"
	collectionOptions := &options.CreateCollectionOptions{
		Validator:       validatorObj,
		ValidationLevel: &validationLevel,
	}
	// create collection with respective validator
	db.GetMongoDbObject().CreateCollection(ctx, user.UserCollectionName, collectionOptions)
	mongoCollection := db.GetMongoDbObject().Collection(user.UserCollectionName)

	// Create indexes
	mongoCollection.Indexes().CreateOne(ctx, user.UserIndexEmail)
	mongoCollection.Indexes().CreateOne(ctx, user.UserIndexUsername)

	adminUser := userModel.User{
		ID:        primitive.ObjectID{},
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Username:  "admin",
		Password:  "admin@123",
		Email:     "admin@admin.com",
		Role:      "admin",
		Active:    true,
	}
	mongoCollection.InsertOne(ctx, adminUser)

}
