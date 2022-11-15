package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var ctx = context.TODO()

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017/")
	_, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
}
