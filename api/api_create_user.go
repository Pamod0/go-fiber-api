package api

import (
	"GoFiberAPI/dto"
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoDB collection
var userCollection *mongo.Collection

func CreateUserApi(c *fiber.Ctx) error {
	newUser := dto.User{}

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
}
