package controllers

import (
	"strconv"

	"example.com/go/models"
	servise "example.com/go/pkg/service"
	"example.com/go/tools"
	"github.com/gofiber/fiber/v2"
)

type CarController struct {
	carService *servise.CarService
}

func NewCarController(app *fiber.App, carService *servise.CarService) *CarController {
	carController := &CarController{
		carService: carService,
	}

	app.Get("/car/all", tools.AuthCheck, carController.GetAllCar)
	app.Get("/car/:id", tools.AuthCheck, carController.GetCarById)
	app.Post("/car/add", tools.AuthCheck, carController.CreateNewCar)

	return carController
}

func (cc *CarController) GetCarById(c *fiber.Ctx) error {
	carID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	car, err := cc.carService.GetCarById(carID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(car)
}

func (cc *CarController) CreateNewCar(c *fiber.Ctx) error {
	var newCar models.Car
	if err := c.BodyParser(&newCar); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	createdCar, err := cc.carService.CreateNewCar(&newCar)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(createdCar)
}

func (cc *CarController) GetAllCar(c *fiber.Ctx) error {
	carRecord, err := cc.carService.GetAllCar()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(carRecord)
}
