package category

// Category Request
// swagger:model
type CategoryRequest struct {
	// the category name
	// required: true
	// min length: 2
	// example: go
	Category string `json:"category" validate:"required,min=2"`
	// Author of this category
	// required: true
	// min length: 5
	// example: admin
	Author string `json:"author" validate:"required,min=5"`
}

// Category Request
// swagger:parameters CategoryRequest
type CategoryRequestSwagger struct {
	// in:body
	Body CategoryRequest
}

// swagger:parameters GetAllCategoriesRequest
type AllUsersRequestSwagger struct {
	// Provide page number
	// in: query
	// required: true
	// example: 1
	Page int `json:"page"`
	// Provide page size
	// in: query
	// required: true
	// example: 10
	PageSize int `json:"pageSize"`
	// Provide search literal
	// in: query
	// required: true
	// example: admin
	Search string `json:"search"`
	// provide order by asc|desc
	// in: query
	// required: false
	// example: asc
	OrderBy string `json:"orderBy"`
}
