package role

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
	Role      string             `bson:"role"`
}
