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
	// Provide id of user who can create questions on this category
	// required: true
	// example: ["62d91ea59b5df5fa6df6ff0f"]
	Users []string `json:"users" validate:"required"`
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

// Update Category Request
// swagger:model
type UpdateCategoryRequest struct {
	// the category name
	// required: true
	// min length: 2
	// example: golang
	Category string `json:"category" validate:"required,min=2"`
}

// swagger:parameters UpdateCategoryRequestId
type UpdateCategoryRequestSwagger struct {
	// Provide categoryId
	// in: path
	// required: true
	// example: 62d91ea59b5df5fa6df6ff0f
	Id string `json:"id"`
	// in:body
	Body UpdateCategoryRequest
}

// swagger:parameters DeleteCategoryRequestId
type DeleteCategoryRequestSwagger struct {
	// Provide categoryId
	// in: path
	// required: true
	// example: 62d91ea59b5df5fa6df6ff0f
	Id string `json:"id"`
}
