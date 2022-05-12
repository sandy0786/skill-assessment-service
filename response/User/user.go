package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserResponse struct {
	ID       primitive.ObjectID `json:"_id"`
	Username string             `json:"username"`
	Email    string             `json:"email"`
	Role     string             `json:"role"`
}
