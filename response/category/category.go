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

// swagger:response UpdateCategorySuccessResponse
type UpdateCategorySuccessResponse struct {
	// in: body
	Body struct {
		// example: 2022-05-20 16:59:05
		TimeStamp string `json:"timestamp"`
		// example: 200
		Status int `json:"status"`
		// example: Category updated successfully
		Message string `json:"message"`
	}
}

// swagger:response UpdateCategoryNotFoundResponse
type UpdateCategoryNotFoundResponse struct {
	// in: body
	Body struct {
		// example: 2022-05-20 16:59:05
		TimeStamp string `json:"timestamp"`
		// example: 404
		Status int `json:"status"`
		// example: No matching category found
		Message string `json:"message"`
	}
}

// swagger:response UpdateCategoryConflictResponse
type UpdateCategoryConflictResponse struct {
	// in: body
	Body struct {
		// example: 2022-05-20 16:59:05
		TimeStamp string `json:"timestamp"`
		// example: 409
		Status int `json:"status"`
		// example: Category already exists
		Message string `json:"message"`
	}
}
