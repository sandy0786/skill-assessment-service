package user

// import "go.mongodb.org/mongo-driver/bson/primitive"

// User Response
// swagger:model
type UserResponse struct {
	// example: admin
	Username string `json:"username"`
	// example: admin@provider.com
	Email string `json:"email"`
	// example: admin
	Role string `json:"role"`
}

// List of Users
// swagger:response UserResponse
type UsersResponse struct {
	// in: body
	Body []UserResponse
}
