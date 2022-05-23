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
	// example: true
	Active bool `json:"active"`
}

// List of Users
// swagger:response UsersResponse
type UsersResponse struct {
	// in: body
	Body []UserResponse
}
