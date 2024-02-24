package service

import (
	"example.com/go/models"
	"example.com/go/pkg/database"
)

type OrderService struct {
	orderRepository database.OrderRepository
}

func NewOrderService(orderRepository database.OrderRepository) *OrderService {
	return &OrderService{
		orderRepository: orderRepository,
	}
}

func (os *OrderService) GetOrderById(orderID int) (*models.Order, error) {
	return os.orderRepository.GetOrderById(orderID)
}

func (os *OrderService) GetAllOrder() ([]*models.Order, error) {
	return os.orderRepository.GetAllOrder()
}

func (os *OrderService) CreateNewOrder(newOrder *models.Order) (*models.Order, error) {
	return os.orderRepository.CreateNewOrder(newOrder)
}

func (os *OrderService) GetOrdersByStatus(clientId int, status string) ([]*models.Order, error) {
	return os.orderRepository.GetOrdersByStatus(clientId, status)
}
