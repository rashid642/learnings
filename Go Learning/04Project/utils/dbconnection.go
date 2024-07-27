package utils

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Collection *mongo.Collection

func InitMongoDB(uri string) {
    clientOptions := options.Client().ApplyURI(uri)
    Client, _ = mongo.Connect(context.TODO(), clientOptions)
    Collection = Client.Database("goProject").Collection("people")
	fmt.Println("MongoDB connected")
}