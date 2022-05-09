package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseInterface interface {
	Connect() error
	Close() error
	GetMongoDbObject() *mongo.Database
	GetMongoDbContext() context.Context
}

type mongoDB struct {
	Client           *mongo.Client
	Db               *mongo.Database
	ctx              context.Context
	host             string
	port             string
	user             string
	password         string
	dbname           string
	connectionString string
}

func NewMongoObj(host string, port string, user string, password string, dbname string, connectionString string) *mongoDB {
	return &mongoDB{
		host:             host,
		port:             port,
		user:             user,
		password:         password,
		dbname:           dbname,
		connectionString: connectionString,
	}
}

func (m *mongoDB) Connect() error {
	log.Println("connecting to database")
	m.ctx = context.TODO()
	// connectionString := fmt.Sprintf("host= %s port = %s user = %s password = %s dbname = %s sslmode=disable", y.host, y.port, y.user, y.password, y.dbname)
	var err error
	// m.Collection, err = gorm.Open("postgres", connectionString)
	clientOptions := options.Client().ApplyURI(m.connectionString)
	m.Client, err = mongo.Connect(m.ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = m.Client.Ping(m.ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	m.Db = m.Client.Database(m.dbname)

	return nil
}

func (m *mongoDB) Close() error {
	log.Println("Closing database connection")
	return m.Client.Disconnect(m.ctx)
}

func (m *mongoDB) GetMongoDbObject() *mongo.Database {
	return m.Db
}

func (m *mongoDB) GetMongoDbContext() context.Context {
	return m.ctx
}
