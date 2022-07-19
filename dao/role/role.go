package role

import (
	"context"
	"log"

	Database "github.com/sandy0786/skill-assessment-service/database"
	roleDocument "github.com/sandy0786/skill-assessment-service/documents/role"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoleDAO interface {
	GetAllRoles() ([]roleDocument.Role, error)
	CreateCollection(string) (*roleDAOImpl, error)
	GetRoleIdByRoleName(string) (primitive.ObjectID, error)
	GetRoleById(primitive.ObjectID) (string, error)
	ValidateRole(string) bool
}

type roleDAOImpl struct {
	db              Database.DatabaseInterface
	collectionName  string
	mongoCollection *mongo.Collection
}

type idStruct struct {
	ID   primitive.ObjectID `bson:"_id"`
	Role string             `bson:"role"`
}

func NewRoleDAO(db Database.DatabaseInterface, collectionName string) *roleDAOImpl {
	return &roleDAOImpl{
		db:              db,
		collectionName:  collectionName,
		mongoCollection: db.GetMongoDbObject().Collection(collectionName),
	}
}

func (q *roleDAOImpl) GetAllRoles() ([]roleDocument.Role, error) {
	log.Println("GetAll Roles ")
	var roles []roleDocument.Role
	cursor, err := q.mongoCollection.Find(q.db.GetMongoDbContext(), bson.M{"role": bson.M{"$ne": "admin"}})
	if err != nil {
		return roles, err
	}
	if err = cursor.All(q.db.GetMongoDbContext(), &roles); err != nil {
		return roles, err
	}
	return roles, err
}

func (c *roleDAOImpl) CreateCollection(collectionName string) (*roleDAOImpl, error) {
	ctx := context.TODO()
	err := c.db.GetMongoDbObject().CreateCollection(ctx, collectionName)
	if err != nil {
		log.Println("Error while creating collection : ", err)
		return nil, err
	}
	// c.mongoCollection = c.mongoCollection.Database().Collection(collectionName)
	return c, err
}

func (r *roleDAOImpl) GetRoleIdByRoleName(role string) (primitive.ObjectID, error) {
	log.Println("role >> ", role)
	// var roleId primitive.ObjectID
	var roleIdStruct idStruct
	r.mongoCollection.FindOne(r.db.GetMongoDbContext(), bson.M{"role": role}).Decode(&roleIdStruct)
	// if err != nil {
	// 	return roleId, err
	// }
	log.Println("result > ", roleIdStruct.ID)
	return roleIdStruct.ID, nil
}

func (r *roleDAOImpl) GetRoleById(roleId primitive.ObjectID) (string, error) {
	var roleIdStruct idStruct
	r.mongoCollection.FindOne(r.db.GetMongoDbContext(), bson.M{"_id": roleId}).Decode(&roleIdStruct)
	return roleIdStruct.Role, nil
}

func (r *roleDAOImpl) ValidateRole(roleId string) bool {
	log.Println("roleId > ", roleId)
	roleIdd, _ := primitive.ObjectIDFromHex(roleId)
	var roleIdStruct idStruct
	r.mongoCollection.FindOne(r.db.GetMongoDbContext(), bson.M{"_id": roleIdd}).Decode(&roleIdStruct)
	if roleIdStruct.ID == primitive.NilObjectID {
		return false
	}
	return true
}
