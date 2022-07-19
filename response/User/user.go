package user

// import "go.mongodb.org/mongo-driver/bson/primitive"

// User Response
// swagger:model
type UserResponse struct {
	// example: admin
	Username string `json:"username"`
	// example: admin@provider.com
	Email string `json:"email"`
	// example: 62d64f2142dac7953ac4ff31
	Role string `json:"roleId"`
	// example: true
	Active bool `json:"active"`
}

// List of Users
// swagger:response UsersResponse
type UsersResponse struct {
	// in: body
	Body []UserResponse
}

// swagger:model
type PasswordResetSuccess struct {
	// example: 2022-05-20 16:59:05
	TimeStamp string `json:"timestamp"`
	// example: 200
	Status int `json:"status"`
	// example: Password reset success
	Message string `json:"message"`
}

// swagger:response ResetUserPasswordResponse
type UserPasswordResetResponse struct {
	// in: body
	Body PasswordResetSuccess
}
