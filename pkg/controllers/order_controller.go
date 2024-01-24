package controllers

import (
	"strconv"

	"example.com/go/models"
	"example.com/go/pkg/service"
	"github.com/gofiber/fiber/v2"
)

type OrderController struct {
	orderService *service.OrderService
}

func NewOrderController(app *fiber.App, orderService *service.OrderService) *OrderController {
	orderController := &OrderController{
		orderService: orderService,
	}

	app.Get("/order/all", orderController.GetAllOrder)
	app.Get("/order/:id", orderController.GetOrderById)
	app.Post("/order/add", orderController.CreateNewOrder)

	return orderController
}

func (oc *OrderController) GetOrderById(c *fiber.Ctx) error {
	orderId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	order, err := oc.orderService.GetOrderById(orderId)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(order)
}

func (oc *OrderController) GetAllOrder(c *fiber.Ctx) error {
	orderRecord, err := oc.orderService.GetAllOrder()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(orderRecord)
}

func (oc *OrderController) CreateNewOrder(c *fiber.Ctx) error {
	var newOrder *models.Order
	if err := c.BodyParser(&newOrder); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	createdOrder, err := oc.orderService.CreateNewOrder(newOrder)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(createdOrder)
}
