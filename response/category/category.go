package category

import "go.mongodb.org/mongo-driver/bson/primitive"

type CategoryResponse struct {
	ID             primitive.ObjectID `json:"_id"`
	CategoryName   string             `json:"categoryName"`
	CollectionName string             `json:"collectionName"`
	Author         string             `json:"author"`
}
