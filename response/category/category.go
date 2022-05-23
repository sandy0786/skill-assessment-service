package category

// import "go.mongodb.org/mongo-driver/bson/primitive"

// swagger:model
type CategoryResponse struct {
	// example: go
	CategoryName string `json:"categoryName"`
	// example: go
	CollectionName string `json:"collectionName"`
	// example: admin
	Author string `json:"author"`
}

// List of Categories
// swagger:response CategoriesResponse
type CategoriesResponse struct {
	// in: body
	Body []CategoryResponse
}