package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	conn "github.com/luisnquin/fiber-restapi-gorm/db"
	"github.com/luisnquin/fiber-restapi-gorm/models"
)

func GetAllCompanies(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")

	companies := []models.Company{}
	conn.DB.Model(models.Company{}).Find(&companies)

	return c.JSON(companies)
}

func GetOneCompany(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	uid := uint(id)
	company := models.Company{}
	company.ID = uid

	err = conn.DB.Where(company.ID).First(&company).Error
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(company)
}

func AddOneCompany(c *fiber.Ctx) error {
	c.Accepts("application/json")

	company := models.Company{}
	err := c.BodyParser(&company)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err = conn.DB.Create(&company).Error
	if err != nil {
		return c.SendStatus(fiber.StatusNotImplemented)
	}

	return c.SendStatus(fiber.StatusCreated)
}

func ModifyOneCompany(c *fiber.Ctx) error {
	c.Accepts("application/json")

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	uid := uint(id)
	request := models.Company{}
	err = c.BodyParser(&request)
	if err != nil {
		return err
	}

	company := models.Company{}
	company.ID = uid
	err = conn.DB.Where(company.ID).First(&company).Error
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	switch {
	case request.Name != "":
		company.Name = request.Name
		conn.DB.Save(&company)
		return c.SendStatus(fiber.StatusAccepted)

	case request.FundationYear != 0:
		company.FundationYear = request.FundationYear
		conn.DB.Save(&company)
		return c.SendStatus(fiber.StatusAccepted)
		
	default:
		return c.SendStatus(fiber.StatusBadRequest)
	}
}

func UpdateOneCompany(c *fiber.Ctx) error {
	c.Accepts("application/json")
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	request := new(models.Company)
	err = c.BodyParser(&request)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	uid := uint(id)
	company := models.Company{}
	company.ID = uid

	err = conn.DB.Where(company.ID).First(&company).Error
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if request.Name != "" {
		company.Name = request.Name
	}
	if request.FundationYear != 0 {
		company.FundationYear = request.FundationYear
	}

	conn.DB.Save(&company)

	return c.SendStatus(fiber.StatusAccepted)
}

func RemoveOneCompany(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	err = conn.DB.Table("companies").Where("id = ?", id).Delete(&models.Company{}).Error
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.SendStatus(fiber.StatusAccepted)
}
