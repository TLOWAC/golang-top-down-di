package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var mongodb *mongo.Client

func ConnectMongoDB() {
	connectOnce.Do(func(){

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

		// Set client options
		clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

		// Connect to MongoDB
		client, err := mongo.Connect(ctx, clientOptions)

		if err != nil {
			panic(err)
		}

		// Check the connection
		err = client.Ping(ctx, nil)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Connected to MongoDB!")

		mongodb = client
	})
}

func GetMongoDB() *mongo.Client {
	return mongodb
}