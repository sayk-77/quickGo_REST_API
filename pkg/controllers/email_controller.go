package controllers

import (
	"example.com/go/models"
	"example.com/go/pkg/service"
	"github.com/gofiber/fiber/v2"
)

type EmailController struct {
	emailService *service.EmailService
}

func NewEmailController(app *fiber.App, emailService *service.EmailService) *EmailController {
	emailController := &EmailController{
		emailService: emailService,
	}
	app.Post("/send-mail", emailController.SendMail)

	return emailController
}

func (ec *EmailController) SendMail(c *fiber.Ctx) error {
	var email *models.Email
	if err := c.BodyParser(&email); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	if err := ec.emailService.SendMail(email); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return nil
}
