package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
	Username  string             `bson:"username"`
	Password  string             `bson:"password"`
	Email     string             `bson:"email,unique"`
	Role      string             `bson:"string"`
}
