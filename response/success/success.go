package success

// Success response
// swagger:model
type SuccessResponse struct {
	// example: 2022-05-20 16:59:05
	TimeStamp string `json:"timestamp"`
	// example: 200
	Status int `json:"status"`
	// example: Data saved successfully
	Message string `json:"message"`
}

// Success response
// swagger:model
type UserDeleteSuccessResponse struct {
	// example: 2022-05-20 16:59:05
	TimeStamp string `json:"timestamp"`
	// example: 200
	Status int `json:"status"`
	// example: User Deleted successfully
	Message string `json:"message"`
}

// Success response
// swagger:model
type UserRevokeSuccessResponse struct {
	// example: 2022-05-20 16:59:05
	TimeStamp string `json:"timestamp"`
	// example: 200
	Status int `json:"status"`
	// example: User revoked successfully
	Message string `json:"message"`
}

// swagger:response HealthResponse
type Health struct {
	// in: body
	Body struct {
		// example: UP
		Status string `json:"status"`
	}
}

// swagger:response SuccessResponse
type UserSuccessResponse struct {
	// in: body
	Body SuccessResponse
}

// swagger:response UserDeleteSuccessResponse
type UserDeleteSuccessResp struct {
	// in: body
	Body UserDeleteSuccessResponse
}

// swagger:response UserRevokeSuccessResponse
type UserRevokeSuccessResp struct {
	// in: body
	Body UserRevokeSuccessResponse
}
