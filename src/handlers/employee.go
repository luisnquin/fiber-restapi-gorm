package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/luisnquin/fiber-restapi-gorm/models"
)

func GetAllEmployees(c *fiber.Ctx) error {
	return c.SendString("Hi")
}

func GetAllEmployesByCompanyCode(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	
	return c.SendStatus(id)
}

func AddOneEmployee(c *fiber.Ctx) error {
	c.Accepts("application/json")

	company := new(models.Company)
	err := c.BodyParser(company)
	if err != nil {
		return err
	}

	return c.JSON(company)
}

func ModifyOneEmployee(c *fiber.Ctx) error {
	c.Accepts("application/json")
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusOK + id)
}

func UpdateOneEmployee(c *fiber.Ctx) error {
	c.Accepts("application/json")
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK + id)
}

func RemoveOneEmployee(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	return c.SendStatus(200 + id)
}
