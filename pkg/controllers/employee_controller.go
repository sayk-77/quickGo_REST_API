package controllers

import (
	"example.com/go/models"
	"example.com/go/pkg/service"
	"github.com/gofiber/fiber/v2"
)

type EmployeeController struct {
	employeeService *service.EmployeeService
}

func NewEmployeeController(app *fiber.App, employeeService *service.EmployeeService) *EmployeeController {
	employeeController := &EmployeeController{
		employeeService: employeeService,
	}
	app.Post("/employee/create", employeeController.CreateEmployee)
	app.Post("/employee/login", employeeController.Login)

	return employeeController
}

func (ec *EmployeeController) CreateEmployee(c *fiber.Ctx) error {
	var employee *models.Employee
	err := c.BodyParser(&employee)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if employee == nil {
		return c.Status(400).JSON(fiber.Map{"error": "некорректные данные сотрудника"})
	}
	msg, err := ec.employeeService.Create(employee)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": msg})
}

func (ec *EmployeeController) Login(c *fiber.Ctx) error {
	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var loginRequest *LoginRequest
	if err := c.BodyParser(&loginRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	token, err := ec.employeeService.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"token": token})
}
