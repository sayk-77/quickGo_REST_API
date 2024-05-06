package controllers

import (
	"example.com/go/tools"
	"strconv"
	"time"

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
	app.Get("/order/complete/:id", orderController.CompleteOrder)
	app.Get("/order/:id", orderController.GetOrderById)
	app.Post("/order/status", orderController.GetOrdersByStatus)
	app.Post("/order/add", orderController.CreateNewOrder)
	app.Post("/order/confirm", orderController.OrderConfirm)
	app.Delete("/order/delete/:orderId", orderController.DeleteOrderByID)

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
	var token string = c.Get("Authorization")

	var clientId int
	var err error

	if token != "" {
		clientId, err = tools.Decoder(token)
		if err != nil {
			return err
		}
	} else {
		if err := c.BodyParser(&newOrder); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		clientId = int(newOrder.ClientID)
	}

	if err := c.BodyParser(&newOrder); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	newOrder.ClientID = uint(clientId)
	newOrder.Status = "Создан"
	newOrder.OrderDate = time.Now().Format("2006-01-02 15:04:05")

	createdOrder, err := oc.orderService.CreateNewOrder(newOrder)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(createdOrder)
}

func (oc *OrderController) GetOrdersByStatus(c *fiber.Ctx) error {
	var token string = c.Get("Authorization")
	clientId, err := tools.Decoder(token)
	if err != nil {
		return err
	}

	var orderStatus struct {
		Status string `json:"status"`
	}
	if err := c.BodyParser(&orderStatus); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	orders, err := oc.orderService.GetOrdersByStatus(clientId, orderStatus.Status)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(orders)
}

func (oc *OrderController) DeleteOrderByID(c *fiber.Ctx) error {
	orderId, err := strconv.Atoi(c.Params("orderId"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	if err := oc.orderService.DeleteOrderById(orderId); err != nil {
		if err.Error() == "Запись не найдена" {
			return c.Status(404).JSON(fiber.Map{"error": "Запись не найдена"})
		}
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Запись удалена"})
}

func (oc *OrderController) OrderConfirm(c *fiber.Ctx) error {
	var orderResponse models.OrderConfirm
	if err := c.BodyParser(&orderResponse); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	var statusOrder string = "В процессе"
	var statusDriver string = "Занят"

	if err := oc.orderService.UpdateOrderStatus(orderResponse, statusOrder, statusDriver); err != nil {
		return err
	}

	return c.JSON(0)
}

func (oc *OrderController) CompleteOrder(c *fiber.Ctx) error {
	orderId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	if err := oc.orderService.CompleteOrder(orderId); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(0)
}
