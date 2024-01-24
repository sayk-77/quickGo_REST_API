package controllers

import (
	"strconv"

	"example.com/go/models"
	"example.com/go/pkg/service"
	"github.com/gofiber/fiber/v2"
)

type TransportationContractController struct {
	transportationContractService *service.TransportationContractService
}

func NewTransportationContractController(app *fiber.App, transportationContractService *service.TransportationContractService) *TransportationContractController {
	transportationContractController := &TransportationContractController{
		transportationContractService: transportationContractService,
	}

	app.Get("/contract/all", transportationContractController.GetAllTransportationContract)
	app.Get("/contract/:id", transportationContractController.GetTransportationContractById)
	app.Post("/contract/add", transportationContractController.CreateNewTransportationContract)

	return transportationContractController
}

func (tcc *TransportationContractController) GetTransportationContractById(c *fiber.Ctx) error {
	transportationContractID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	transportationContract, err := tcc.transportationContractService.GetTransportationContractById(transportationContractID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(transportationContract)
}

func (tcc *TransportationContractController) GetAllTransportationContract(c *fiber.Ctx) error {
	transportationContractRecord, err := tcc.transportationContractService.GetAllTransportationContract()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(transportationContractRecord)
}

func (tcc *TransportationContractController) CreateNewTransportationContract(c *fiber.Ctx) error {
	var transportationContract *models.TransportationContract
	if err := c.BodyParser(&transportationContract); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	createdTransportation, err := tcc.transportationContractService.CreateNewTransportationContract(transportationContract)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(createdTransportation)
}
