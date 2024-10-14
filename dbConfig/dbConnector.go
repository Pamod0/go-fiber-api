package dbconfig

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect to MongoDB
func ConnectToMongoDB() *mongo.Client {
	// uri := "mongodb://localhost:27017" // For local MongoDB
	uri := "mongodb+srv://pamod:29StecdLbCivz7nR@pamod.inscxbr.mongodb.net/?retryWrites=true&w=majority&appName=Pamod" // For MongoDB Atlas

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}
