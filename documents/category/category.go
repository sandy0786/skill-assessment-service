package category

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	ID             primitive.ObjectID `bson:"_id"`
	CreatedAt      time.Time          `bson:"createdAt"`
	UpdatedAt      time.Time          `bson:"updatedAt"`
	CategoryName   string             `bson:"categoryName"`
	CollectionName string             `bson:"collectionName"`
	Author         string             `bson:"author"`
}
