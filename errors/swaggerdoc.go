package errors

// swagger:model
type ConflictError struct {
	// example: 2022-05-20 16:59:05
	TimeStamp string `json:"timestamp"`
	// example: 409
	Status int `json:"status"`
	// example: Already exist
	Message string `json:"message"`
}

// swagger:model
type BadRequestError struct {
	// example: 2022-05-20 16:59:05
	TimeStamp string `json:"timestamp"`
	// example: 400
	Status int `json:"status"`
	// example: Invalid parameters
	Message string `json:"message"`
}

// swagger:model
type NotFoundError struct {
	// example: 2022-05-20 16:59:05
	TimeStamp string `json:"timestamp"`
	// example: 404
	Status int `json:"status"`
	// example: Not found
	Message string `json:"message"`
}

// swagger:model
type UnAuthorizedError struct {
	// example: 2022-05-20 16:59:05
	TimeStamp string `json:"timestamp"`
	// example: 400
	Status int `json:"status"`
	// example: Invalid username or password
	Message string `json:"message"`
}

// swagger:response InternalServerErrorResponse
type InternalErrorStruct struct {
	// in: body
	Body GlobalError
}

// swagger:response NotFoundEmptyErrorResponse
type NotFoundEmptyErrorStruct struct {
	// in: body
	// example: []
	Body []interface{}
}

// swagger:response BadRequestErrorResponse
type BadRequestErrorStruct struct {
	// in: body
	Body BadRequestError
}

// swagger:response ConflictErrorResponse
type ConflictErrorStruct struct {
	// in: body
	Body ConflictError
}

// swagger:response NotFoundErrorResponse
type NotFoundErrorStruct struct {
	// in: body
	Body NotFoundError
}

// swagger:response UnAuthorizedResponse
type UnAuthorizedErrorStruct struct {
	// in: body
	Body UnAuthorizedError
}
