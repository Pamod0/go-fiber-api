package main

import (
	"GoFiberAPI/apiHandlers"
	"GoFiberAPI/dbConfig"
	"GoFiberAPI/integrations"
	// "context"
	"fmt"
	"log"
	// "time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

// User struct to represent the MongoDB document
// type User struct {
// 	ID   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
// 	Name string             `json:"name"`
// 	Age  int                `json:"age"`
// }

// // MongoDB collection
// var userCollection *mongo.Collection

// // Connect to MongoDB
// func ConnectDB() *mongo.Client {
// 	// Replace with your MongoDB connection URI
// 	// uri := "mongodb://localhost:27017" // For local MongoDB
// 	uri := "mongodb+srv://pamod:29StecdLbCivz7nR@pamod.inscxbr.mongodb.net/?retryWrites=true&w=majority&appName=Pamod" // For MongoDB Atlas

// 	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	err = client.Connect(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Connected to MongoDB!")
// 	return client
// }

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Cannot load environment file")
	}
	integrations.SetEnvironmentVariables()
}

func main() {
	fmt.Println("Starting application")

	// Initialize the Fiber app
	app := fiber.New(fiber.Config{
		AppName:   "GoFiberAPI",
		BodyLimit: 4000 * 1024,
	})

	// Connect To Database
	dbConfig.ConnectMongoDB()

	//Remove Pre-Generated Outs
	// dbConfig.RemoveGeneratedOuts()

	// Define the API routes
	apiHandlers.Router(app)

	// Start the server
	log.Fatal(app.Listen(":8888"))

	// // Connect to MongoDB and select the collection
	// client := ConnectDB()
	// userCollection = client.Database("gofiberapi").Collection("users")

	// // Route to get all users
	// app.Get("/users", func(c *fiber.Ctx) error {
	// 	var users []User
	// 	cursor, err := userCollection.Find(context.Background(), bson.D{})
	// 	if err != nil {
	// 		return c.Status(500).SendString(err.Error())
	// 	}

	// 	if err = cursor.All(context.Background(), &users); err != nil {
	// 		return c.Status(500).SendString(err.Error())
	// 	}

	// 	return c.JSON(users)
	// })

	// // Route to get a specific user by ID
	// app.Get("/users/:id", func(c *fiber.Ctx) error {
	// 	idParam := c.Params("id")
	// 	objID, err := primitive.ObjectIDFromHex(idParam)
	// 	if err != nil {
	// 		return c.Status(400).SendString("Invalid user ID")
	// 	}

	// 	var user User
	// 	err = userCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&user)
	// 	if err != nil {
	// 		return c.Status(404).SendString("User not found")
	// 	}

	// 	return c.JSON(user)
	// })

	// // Route to create a new user
	// app.Post("/users", func(c *fiber.Ctx) error {
	// 	var newUser User

	// 	// Parse the request body into the newUser struct
	// 	if err := c.BodyParser(&newUser); err != nil {
	// 		return c.Status(400).SendString(err.Error())
	// 	}

	// 	// Insert the new user into the MongoDB collection
	// 	newUser.ID = primitive.NewObjectID()
	// 	_, err := userCollection.InsertOne(context.Background(), newUser)
	// 	if err != nil {
	// 		return c.Status(500).SendString(err.Error())
	// 	}

	// 	return c.Status(201).JSON(newUser)
	// })

	// // Start the server on port 3000
	// app.Listen(":3000")
}
