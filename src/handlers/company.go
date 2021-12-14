package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/luisnquin/fiber-restapi-gorm/models"
)

func GetAllCompanies(c *fiber.Ctx) error {
	return c.SendString("Hi")
}

func GetOneCompany(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	return c.SendStatus(id)
}

func AddOneCompany(c *fiber.Ctx) error {
	c.Accepts("application/json")

	company := new(models.Company)
	err := c.BodyParser(company)
	if err != nil {
		return err
	}

	return c.JSON(company)
}

func ModifyOneCompany(c *fiber.Ctx) error {
	c.Accepts("application/json")
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusOK + id)
}

func UpdateOneCompany(c *fiber.Ctx) error {
	c.Accepts("application/json")
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK + id)
}

func RemoveOneCompany(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	return c.SendStatus(200 + id)
}