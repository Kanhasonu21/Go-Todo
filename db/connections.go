package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConnectMongo() {
	MONGO_DB_URL := os.Getenv("MONGO_DB_URL")
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(MONGO_DB_URL))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	// Optionally, ping the database to ensure the connection is established
	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	fmt.Println("Database Connected!")

}

func GetDB() *mongo.Client {
	return client
}

func DisconnectMongo() {
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatal("Failed to disconnect from MongoDB:", err)
	}
	fmt.Println("Database DisConnected!")

}
