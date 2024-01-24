package controllers

import (
	"strconv"

	"example.com/go/models"
	"example.com/go/pkg/service"
	"github.com/gofiber/fiber/v2"
)

type WaybillController struct {
	waybillService *service.WaybillService
}

func NewWaybillController(app *fiber.App, waybillService *service.WaybillService) *WaybillController {
	waybillController := &WaybillController{
		waybillService: waybillService,
	}

	app.Get("/waybill/all", waybillController.GetAllWaybill)
	app.Get("/waybill/:id", waybillController.GetWaybillById)
	app.Post("/waybill/add", waybillController.CreateNewWaybill)

	return waybillController
}

func (wc *WaybillController) GetWaybillById(c *fiber.Ctx) error {
	waybillID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	waybill, err := wc.waybillService.GetWaybillById(waybillID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(waybill)
}

func (wc *WaybillController) GetAllWaybill(c *fiber.Ctx) error {
	waybillRecord, err := wc.waybillService.GetAllWaybill()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(waybillRecord)
}

func (wc *WaybillController) CreateNewWaybill(c *fiber.Ctx) error {
	var newWaybill *models.WayBill
	if err := c.BodyParser(&newWaybill); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	createdWaybill, err := wc.waybillService.CreateNewWaybill(newWaybill)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(createdWaybill)
}
