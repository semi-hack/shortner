package db

import (
	"context"
	"fmt"
	"log"

	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Database
)

func init() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/shortner")


	// Connect to MongoDB
	if clt, err := mongo.Connect(context.TODO(), clientOptions); err != nil {
		log.Fatalln(err)
	} else {
		if err := clt.Ping(context.TODO(), nil); err != nil {
			log.Fatalln(err)
		}
		client = clt.Database("rest")

	}
	//collection := client.Collection("users")

	fmt.Println("Connected to MongoDB!")
}

// Collection ...
type Collection struct {
	// collection is a connection to a MongoDB collection
	collection *mongo.Collection

	// database is a connection to the MongoDB database that houses the collection
	database *mongo.Database
}

// GetConnection returns a connection to the DB
func GetConnection() *mongo.Database {
	return client
}
