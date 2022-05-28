package category

// Category Request
// swagger:model
type CategoryRequest struct {
	// the category name
	// required: true
	// min length: 2
	// example: go
	CategoryName string `json:"categoryName" validate:"required"`
	// the collection name
	// required: true
	// min length: 2
	// example: go
	CollectionName string `json:"collectionName" validate:"required"`
	// Author of this category
	// required: true
	// min length: 5
	// example: admin
	Author string `json:"author" validate:"required"`
}

// Category Request
// swagger:parameters CategoryRequest
type CategoryRequestSwagger struct {
	// in:body
	Body CategoryRequest
}
