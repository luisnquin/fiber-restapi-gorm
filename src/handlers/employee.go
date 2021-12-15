package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	conn "github.com/luisnquin/fiber-restapi-gorm/db"
	"github.com/luisnquin/fiber-restapi-gorm/models"
)

func GetAllEmployees(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")

	employees := []models.Employee{}
	err := conn.DB.Model(models.Employee{}).Find(&employees).Error
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(employees)
}

func GetAnyEmployee(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	uid := uint(id)
	employee := models.Employee{}
	err = conn.DB.Where("id = ?", uid).Find(&employee).Error
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.JSON(employee)
}

func GetAllEmployeesByCompanyCode(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	company_employees := []models.EmployeesAndCompany{}
	err = conn.DB.Raw("SELECT employees.fullname, employees.age, employees.position, employees.company_code, companies.name, companies.fundation_year FROM employees INNER JOIN companies ON employees.company_code = companies.id AND employees.company_code = ?", id).Scan(&company_employees).Error
	if err != nil {
		c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(company_employees)
}

func AddOneEmployee(c *fiber.Ctx) error {
	c.Accepts("application/json")

	employee := models.Employee{}
	err := c.BodyParser(&employee)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err = conn.DB.Create(&employee).Error
	if err != nil {
		return c.SendStatus(fiber.StatusNotImplemented)
	}

	return c.SendStatus(fiber.StatusCreated)
}

func ModifyOneEmployee(c *fiber.Ctx) error {
	c.Accepts("application/json")

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	employee := models.Employee{}
	err = conn.DB.Where("id = ?", id).Find(&employee).Error
	if err != nil {
		c.SendStatus(fiber.StatusNotFound)
	}

	request := models.Employee{}
	err = c.BodyParser(&request)
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
	}

	switch {
	case request.Fullname != "":
		employee.Fullname = request.Fullname
		conn.DB.Save(&employee)
		return c.SendStatus(fiber.StatusAccepted)

	case request.Age != 0:
		employee.Age = request.Age
		conn.DB.Save(&employee)
		return c.SendStatus(fiber.StatusAccepted)

	case request.Position != "":
		employee.Position = request.Position
		conn.DB.Save(&employee)
		return c.SendStatus(fiber.StatusAccepted)

	case request.CompanyCode != 0:
		employee.CompanyCode = request.CompanyCode
		conn.DB.Save(&employee)
		return c.SendStatus(fiber.StatusAccepted)

	default:
		return c.SendStatus(fiber.StatusNotModified)
	}
}

func UpdateOneEmployee(c *fiber.Ctx) error {
	c.Accepts("application/json")

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	
	request := models.Employee{}
	err = c.BodyParser(&request)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	employee := models.Employee{}
	err = conn.DB.Where("id = ?", id).Find(&employee).Error
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	
	if request.Fullname != "" {
		employee.Fullname = request.Fullname
	}
	if request.Age != 0 {
		employee.Age = request.Age
	}
	if request.Position != "" {
		employee.Position = request.Position
	}
	if request.CompanyCode != 0 {
		employee.CompanyCode = request.CompanyCode
	}
	conn.DB.Save(&employee)
	return c.SendStatus(fiber.StatusAccepted)
}

func RemoveOneEmployee(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	err = conn.DB.Table("employees").Where("id = ?", id).Delete(&models.Employee{}).Error
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.SendStatus(fiber.StatusAccepted)
}
