package apiHandlers

import (
	"GoFiberAPI/api"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Router(app *fiber.App) {
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())

	group := app.Group("/GoFiber/api")
	defaultGroup := app.Group("/")
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Static("/", "./docs/rapiDoc/build")
	DefaultMappings(defaultGroup)
	RouteMappings(group)
}

func RouteMappings(cg fiber.Router) {

	cg.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("API is running!")
	})
	cg.Post("/createUser", api.CreateUserApi)
	cg.Get("/swagger", api.SwaggerHandler)

}

func DefaultMappings(cg fiber.Router) {
	cg.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{"message": "GO Fiber API service is up and running", "version": "1.0"})
	})
}
