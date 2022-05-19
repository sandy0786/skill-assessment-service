package user

import "go.mongodb.org/mongo-driver/bson/primitive"

// User Response
// swagger:response UserResponse
type UserResponse struct {
	ID       primitive.ObjectID `json:"_id"`
	Username string             `json:"username"`
	Email    string             `json:"email"`
	Role     string             `json:"role"`
}

// List of Users
// swagger:response UsersResponse
type UsersResponse struct {
	// list of Users
	// out: []UserResponse
	UsersResponse []UserResponse
}
