package auth

// login Request
// swagger:model
type LoginRequest struct {
	// the username for this user
	// required: true
	// example: admin
	Username string `json:"username" validate:"required,min=5"`
	// the password for this user
	// required: true
	// min length: 8
	// example: admin@123
	Password string `json:"password" validate:"required,min=8"`
}

// User Request
// swagger:parameters LoginRequest
type LoginRequestSwagger struct {
	// in:body
	Body LoginRequest
}

// swagger:parameters RefreshTokenRequest
type RefreshTokenRequestSwagger struct {
	// in:body
	Body interface{}
}
