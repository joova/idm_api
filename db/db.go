package db

import (
	"context"
	"log"

	"github.com/mongodb/mongo-go-driver/mongo"
)

var db *mongo.Database

// init connect to database
func init() {
	client, err := mongo.Connect(context.TODO(), "mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")

	db = client.Database("logika")
}

// Disconnect disconnect from database
func Disconnect() {
	client := db.Client()
	client.Disconnect(context.TODO())
	log.Println("Disconnected from MongoDB!")
}
