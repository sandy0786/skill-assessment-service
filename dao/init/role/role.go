package role

import (
	"time"

	roleModel "github.com/sandy0786/skill-assessment-service/documents/role"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var RoleCollectionName = "role"
var RoleValidatorFilePath = "./dao/init/role/RolesValidator.json"

// index model
var RoleIndexRoleName = mongo.IndexModel{
	Keys: bson.M{
		"role": 1, // index in ascending order
	}, Options: options.Index().SetUnique(true),
}

var AdminObjId primitive.ObjectID

func init() {
	AdminObjId, _ = primitive.ObjectIDFromHex("62d6435f333f27963c162e01")
}

var AdminRole = roleModel.Role{
	ID:        AdminObjId,
	CreatedAt: time.Now().UTC(),
	UpdatedAt: time.Now().UTC(),
	Role:      "admin",
}

var ManagerRole = roleModel.Role{
	ID:        primitive.NewObjectID(),
	CreatedAt: time.Now().UTC(),
	UpdatedAt: time.Now().UTC(),
	Role:      "manager",
}

var GuestRole = roleModel.Role{
	ID:        primitive.NewObjectID(),
	CreatedAt: time.Now().UTC(),
	UpdatedAt: time.Now().UTC(),
	Role:      "guest",
}
