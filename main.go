package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Use(func(c *fiber.Ctx) {
		c.SendStatus(404) // => 404 "Not Found"
	})

	app.Listen(":3000")
}