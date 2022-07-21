package category

// List of Categories
// swagger:model
type CategoryResponse struct {
	ID string `json:"id"`
	// example: go
	Category string `json:"category"`
	// example: 1658228764
	CreatedAt int64 `json:"createdAt"`
	// example: 1658228764
	UpdatedAt int64 `json:"updatedAt"`
}

type CategoryResults struct {
	Data []CategoryResponse `json:"data"`
	// example: 1
	TotalRecords int64 `json:"totalRecords"`
}

// swagger:response CategoriesResponse
type CategoriesResponse struct {
	// in: body
	Body CategoryResults
}
