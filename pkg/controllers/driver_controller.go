package controllers

import (
	"strconv"

	"example.com/go/models"
	"example.com/go/pkg/service"
	"github.com/gofiber/fiber/v2"
)

type DriverController struct {
	driverService *service.DriverService
}

func NewDriverController(app *fiber.App, driverService *service.DriverService) *DriverController {
	driverController := &DriverController{
		driverService: driverService,
	}

	app.Get("/driver/all", driverController.GetAllDriver)
	app.Get("/driver/free", driverController.GetFreeDriver)
	app.Get("/driver/:id", driverController.GetDriverById)
	app.Post("/driver/add", driverController.CreateNewDriver)

	return driverController
}

func (dc *DriverController) GetDriverById(c *fiber.Ctx) error {
	driverID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	driver, err := dc.driverService.GetDriverById(driverID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(driver)
}

func (dc *DriverController) GetAllDriver(c *fiber.Ctx) error {
	driverRecord, err := dc.driverService.GetAllDriver()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(driverRecord)
}

func (dc *DriverController) GetFreeDriver(c *fiber.Ctx) error {
	freeDrivers := make([]*models.Driver, 0)

	allDrivers, err := dc.driverService.GetAllDriver()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	for _, driver := range allDrivers {
		if driver.Status == "Свободен" {
			freeDrivers = append(freeDrivers, driver)
		}
	}

	return c.JSON(freeDrivers)

}

func (dc *DriverController) CreateNewDriver(c *fiber.Ctx) error {
	var newDriver models.Driver
	if err := c.BodyParser(&newDriver); err != nil {
		c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	createdDriver, err := dc.driverService.CreateNewDriver(&newDriver)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(createdDriver)
}
