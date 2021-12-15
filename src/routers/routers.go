package routers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/luisnquin/fiber-restapi-gorm/handlers"
)

func Init() *fiber.App {
	app := fiber.New()
	api := app.Group("/api")
	v1 := api.Group("/v1")

	company := v1.Group("/companies")

	company.Get("", handlers.GetAllCompanies)
	company.Get("/:id", handlers.GetOneCompany)
	company.Post("", handlers.AddOneCompany)
	company.Patch("/:id", handlers.ModifyOneCompany)
	company.Put("/:id", handlers.UpdateOneCompany)
	company.Delete("/:id", handlers.RemoveOneCompany)

	employee := v1.Group("/employees")
	employee.Get("", handlers.GetAllEmployees)
	employee.Get("/:id", handlers.GetAnyEmployee)
	employee.Get("/group/:id", handlers.GetAllEmployeesByCompanyCode)
	employee.Post("", handlers.AddOneEmployee)
	employee.Patch("/:id", handlers.ModifyOneEmployee)
	employee.Put("/:id", handlers.UpdateOneEmployee)
	employee.Delete("/:id", handlers.RemoveOneEmployee)

	return app
}
