package controllers

import (
	"example.com/go/models"
	"example.com/go/pkg/service"
	"example.com/go/tools"
	"fmt"
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
	app.Get("/client/info", clientController.GetClientById)
	app.Post("/client/add", clientController.CreateNewClient)
	app.Post("/client/login", clientController.ClientLogin)
	app.Post("/client/update", clientController.ClientUpdateData)
	app.Post("/client/change-password", clientController.ClientChangePassword)

	return clientController
}

func (cc *ClientController) GetClientById(c *fiber.Ctx) error {

	var token string = c.Get("Authorization")

	decodeToken, err := tools.Decoder(token)
	if err != nil {
		return err
	}

	clientId := int(decodeToken["id"].(float64))

	client, err := cc.clientService.GetClientById(clientId)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	clientResponse := models.ClientResponse{
		ID:          client.ID,
		FirstName:   client.FirstName,
		LastName:    client.LastName,
		Address:     client.Address,
		Email:       client.Email,
		PhoneNumber: client.PhoneNumber,
	}

	return c.JSON(clientResponse)
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

func (cc *ClientController) ClientLogin(c *fiber.Ctx) error {
	var LoginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&LoginData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	token, err := cc.clientService.ClientLogin(LoginData.Email, LoginData.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"token": token})
}

func (cc *ClientController) ClientUpdateData(c *fiber.Ctx) error {
	var UpdateClientData models.ClientResponse
	if err := c.BodyParser(&UpdateClientData); err != nil {
		return err
	}

	if err := cc.clientService.ClientUpdateData(UpdateClientData); err != nil {
		return err
	}

	return nil
}

func (cc *ClientController) ClientChangePassword(c *fiber.Ctx) error {
	var clientData struct {
		CurrentPassword string
		NewPassword     string
	}

	var token string = c.Get("Authorization")
	decodeToken, err := tools.Decoder(token)
	if err != nil {
		return err
	}

	clientId := int(decodeToken["id"].(float64))

	if err := c.BodyParser(&clientData); err != nil {
		return fmt.Errorf("Не верные данные")
	}

	if err := cc.clientService.ClientChangePassword(clientData.CurrentPassword, clientData.NewPassword, clientId); err != nil {
		return err
	}

	return nil
}
