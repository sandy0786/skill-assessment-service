package role

import (
	"context"
	"log"

	Database "github.com/sandy0786/skill-assessment-service/database"
	roleDocument "github.com/sandy0786/skill-assessment-service/documents/role"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoleDAO interface {
	GetAllRoles() ([]roleDocument.Role, error)
	CreateCollection(string) (*roleDAOImpl, error)
}

type roleDAOImpl struct {
	db              Database.DatabaseInterface
	collectionName  string
	mongoCollection *mongo.Collection
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
