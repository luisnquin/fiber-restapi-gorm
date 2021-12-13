package routers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/luisnquin/fiber-restapi-gorm/handlers"
)

func Init() *fiber.App {
	app := fiber.New()
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/", handlers.GetAll)

	return app
}