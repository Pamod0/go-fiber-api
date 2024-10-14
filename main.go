package main

import (
	apihandlers "GoFiberAPI/apiHandlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize the Fiber app
	app := fiber.New()

	// Handle routes
	apihandlers.Router(app)

	// Start the server on port 3000
	app.Listen(":3000")
}
