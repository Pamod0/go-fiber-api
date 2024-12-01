package router

import (
	"github.com/Pamod0/go-fiber-api/handler"
	"github.com/Pamod0/go-fiber-api/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	// Middleware
	api := app.Group("/api", middleware.AuthReq())

	// routes
	api.Get("/", handler.GetAllProducts)
	api.Get("/:id", handler.GetSingleProduct)
	api.Post("/", handler.CreateProduct)
	api.Delete("/:id", handler.DeleteProduct)

}