package main

import (
	"log"

	"github.com/Pamod0/go-fiber-api/database"
	"github.com/Pamod0/go-fiber-api/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() { // entry point to our program

	// Connect to database
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New() // call the New() method - used to instantiate a new Fiber App

	app.Use(logger.New())

	router.SetupRoutes(app)

	app.Listen(":3000") // listen/Serve the new Fiber app on port 3000

}
