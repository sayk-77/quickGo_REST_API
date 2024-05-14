package controllers

import (
	"strconv"

	"example.com/go/models"
	"example.com/go/pkg/service"
	"github.com/gofiber/fiber/v2"
)

type CargoTypeController struct {
	cargoTypeService *service.CargoTypeService
}

func NewCargoTypeController(app *fiber.App, cargoTypeService *service.CargoTypeService) *CargoTypeController {
	cargoTypeController := &CargoTypeController{
		cargoTypeService: cargoTypeService,
	}

	app.Get("/cargo_type/all", cargoTypeController.GetAllCargoType)
	app.Post("/cargo_type/update", cargoTypeController.UpdateCargoType)
	app.Get("/cargo_type/delete/:id", cargoTypeController.DeleteCargoType)
	app.Get("/cargo_type/:id", cargoTypeController.GetCargoTypeById)
	app.Post("/cargo_type/add", cargoTypeController.CreateNewCargoType)

	return cargoTypeController
}

func (ctc *CargoTypeController) GetCargoTypeById(c *fiber.Ctx) error {
	cargoTypeID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	cargoType, err := ctc.cargoTypeService.GetCargoTypeById(cargoTypeID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(cargoType)
}

func (ctc *CargoTypeController) GetAllCargoType(c *fiber.Ctx) error {
	cargoTypeRecord, err := ctc.cargoTypeService.GetAllTypeCargo()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(cargoTypeRecord)
}

func (ctx *CargoTypeController) CreateNewCargoType(c *fiber.Ctx) error {
	var newTypeCargo models.CargoType
	if err := c.BodyParser(&newTypeCargo); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	createdCargoType, err := ctx.cargoTypeService.CreateNewCargoType(&newTypeCargo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(createdCargoType)
}

func (ctx *CargoTypeController) UpdateCargoType(c *fiber.Ctx) error {
	var updatedTypeCargo models.CargoType
	if err := c.BodyParser(&updatedTypeCargo); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := ctx.cargoTypeService.UpdateCargoType(&updatedTypeCargo); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON("ok")
}

func (ctx *CargoTypeController) DeleteCargoType(c *fiber.Ctx) error {
	cargoTypeID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := ctx.cargoTypeService.DeleteCargoType(cargoTypeID); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON("ok")
}
