package controllers

import (
	"strconv"

	"example.com/go/models"
	"example.com/go/pkg/service"
	"github.com/gofiber/fiber/v2"
)

type ClientController struct {
	clientService *service.ClientService
}

func NewClientContorller(app *fiber.App, clientService *service.ClientService) *ClientController {
	clientController := &ClientController{
		clientService: clientService,
	}

	app.Get("/client/all", clientController.GetAllClient)
	app.Get("/client/:id", clientController.GetClientById)
	app.Post("/client/add", clientController.CreateNewClient)

	return clientController
}

func (cc *ClientController) GetClientById(c *fiber.Ctx) error {
	clientId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	client, err := cc.clientService.GetClientById(clientId)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(client)
}

func (cc *ClientController) GetAllClient(c *fiber.Ctx) error {
	carRecord, err := cc.clientService.GetAllClient()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(carRecord)
}

func (cc *ClientController) CreateNewClient(c *fiber.Ctx) error {
	var newClient models.Client
	if err := c.BodyParser(&newClient); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	createdClient, err := cc.clientService.CreateNewClient(&newClient)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(createdClient)
}
