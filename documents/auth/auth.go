package auth

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Username string             `bson:"username"`
	Password string             `bson:"password"`
	Email    string             `bson:"email,unique"`
	Role     primitive.ObjectID `bson:"role"`
	Active   bool               `bson:"active"`
}
