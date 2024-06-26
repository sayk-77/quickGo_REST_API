package controllers

import (
	"strconv"

	"example.com/go/models"
	servise "example.com/go/pkg/service"
	"github.com/gofiber/fiber/v2"
)

type CarController struct {
	carService *servise.CarService
}

func NewCarController(app *fiber.App, carService *servise.CarService) *CarController {
	carController := &CarController{
		carService: carService,
	}

	app.Get("/car/all", carController.GetAllCar)
	app.Get("/car/free", carController.FindFreeCar)
	app.Get("/car/delete/:id", carController.DeleteCar)
	app.Get("/car/:id", carController.GetCarById)
	app.Post("/car/add", carController.CreateNewCar)
	app.Post("/car/edit", carController.EditCar)

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

func (cc *CarController) FindFreeCar(c *fiber.Ctx) error {
	carRecord, err := cc.carService.FindFreeCar()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(carRecord)
}

func (cc *CarController) EditCar(c *fiber.Ctx) error {
	var car models.Car
	if err := c.BodyParser(&car); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	if err := cc.carService.UpdateCar(&car); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON("ok")
}

func (cc *CarController) DeleteCar(c *fiber.Ctx) error {
	carID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	if err := cc.carService.DeleteCar(carID); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON("ok")
}
