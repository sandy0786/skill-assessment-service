package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Task struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	Text      string             `bson:"text"`
	completed bool               `bson:"completed"`
}

var collection *mongo.Collection
var ctx = context.TODO()

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("tasker").Collection("tasks")
}

func main() {
	log.Println("hello")
	task := &Task{
		ID:        primitive.NewObjectID(),
		Text:      "test",
		completed: true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	a, err := collection.InsertOne(ctx, task)
	if err != nil {
		log.Println("error while insertion : ", err)
	}
	log.Println("a > ", a)
}
