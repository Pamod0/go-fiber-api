package apihandlers

import (
	"GoFiberAPI/dbConfig"
	"GoFiberAPI/dto"
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

func Router(app *fiber.App) {

	// Connect to MongoDB and select the collection
	client := dbConfig.ConnectToMongoDB()
	userCollection = client.Database("testdb").Collection("users")

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API is running!")
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Health endpoint is working!")
	})

	// Route to get all users
	app.Get("/users", func(c *fiber.Ctx) error {
		var users []dto.User
		cursor, err := userCollection.Find(context.Background(), bson.D{})
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		if err = cursor.All(context.Background(), &users); err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.JSON(users)
	})

	// Route to get a specific user by ID
	app.Get("/users/:id", func(c *fiber.Ctx) error {
		idParam := c.Params("id")
		objID, err := primitive.ObjectIDFromHex(idParam)
		if err != nil {
			return c.Status(400).SendString("Invalid user ID")
		}

		var user dto.User
		err = userCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&user)
		if err != nil {
			return c.Status(404).SendString("User not found")
		}

		return c.JSON(user)
	})

	// Route to create a new user
	app.Post("/users", func(c *fiber.Ctx) error {
		var newUser dto.User

		// Parse the request body into the newUser struct
		if err := c.BodyParser(&newUser); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		// Insert the new user into the MongoDB collection
		newUser.ID = primitive.NewObjectID()
		_, err := userCollection.InsertOne(context.Background(), newUser)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.Status(201).JSON(newUser)
	})
}
