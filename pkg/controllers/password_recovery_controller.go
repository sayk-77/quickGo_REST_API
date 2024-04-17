package controllers

import (
	"example.com/go/pkg/service"
	"github.com/gofiber/fiber/v2"
)

type PasswordRecoveryController struct {
	passwordRecoveryService *service.PasswordRecoveryService
}

func NewPasswordRecoveryController(app *fiber.App, passwordRecoveryService *service.PasswordRecoveryService) *PasswordRecoveryController {
	passwordRecoveryController := &PasswordRecoveryController{
		passwordRecoveryService: passwordRecoveryService,
	}

	app.Post("/password/email", passwordRecoveryController.CheckEmail)
	app.Post("/password/confirm", passwordRecoveryController.CheckCode)
	app.Post("/password/change", passwordRecoveryController.ChangePassword)

	return passwordRecoveryController
}

func (prc *PasswordRecoveryController) CheckEmail(c *fiber.Ctx) error {
	type Response struct {
		Email string `json:"email"`
	}

	response := new(Response)

	if err := c.BodyParser(response); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	code, err := prc.passwordRecoveryService.CheckEmail(response.Email)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"code": code})
}

func (prc *PasswordRecoveryController) CheckCode(c *fiber.Ctx) error {
	type Response struct {
		Email string `json:"email"`
		Code  int    `json:"code"`
	}

	response := new(Response)

	if err := c.BodyParser(response); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	answer, err := prc.passwordRecoveryService.CheckCode(response.Code, response.Email)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"answer": answer})
}

func (prc *PasswordRecoveryController) ChangePassword(c *fiber.Ctx) error {
	type Response struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	response := new(Response)

	if err := c.BodyParser(response); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	if err := prc.passwordRecoveryService.ChangePassword(response.Email, response.Password); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "success"})
}
