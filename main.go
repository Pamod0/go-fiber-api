package main

import (
	"GoFiberAPI/apiHandlers"
	"GoFiberAPI/platform/authenticator"
	"log"

	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load the env vars: %v", err)
	}
	// Initialize the Fiber app
	app := fiber.New()

	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	// Handle routes
	apiHandlers.Router(app, auth)

	// Start the server on port 3000
	app.Listen(":3000")
}
