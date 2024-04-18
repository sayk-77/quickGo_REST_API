package service

import (
	"example.com/go/models"
	"example.com/go/pkg/database"
	"time"
)

type OrderService struct {
	orderRepository    database.OrderRepository
	driverRepository   database.DriverRepository
	contractRepository database.TransportationContractRepository
	waybillRepository  database.WaybillRepository
}

func NewOrderService(orderRepository database.OrderRepository, driverRepository database.DriverRepository, contractRepository database.TransportationContractRepository, waybillRepository database.WaybillRepository) *OrderService {
	return &OrderService{
		orderRepository:    orderRepository,
		driverRepository:   driverRepository,
		contractRepository: contractRepository,
		waybillRepository:  waybillRepository,
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

func (os *OrderService) DeleteOrderById(orderId int) error {
	return os.orderRepository.DeleteOrderById(orderId)
}

func (os *OrderService) UpdateOrderStatus(order models.OrderConfirm, statusOrder string, statusDriver string) error {
	if err := os.orderRepository.UpdateOrderStatus(int(order.ID), statusOrder, order.DeliveryDate); err != nil {
		return err
	}

	if err := os.driverRepository.UpdateStatusDriver(int(order.DriverId), statusDriver); err != nil {
		return err
	}

	var contract models.TransportationContract
	contract.OrderID = order.ID
	contract.ClientID = uint(order.ClientID)
	contract.ContractDate = time.Now().Format("2006-01-02")
	contract.CarID = uint(order.CarId)
	contract.ExpiryDate = time.Now().Format("2006-01-02")

	contractId, err := os.contractRepository.CreateNewTransportationContract(&contract)
	if err != nil {
		return err
	}

	var waybill models.WayBill
	waybill.CarID = uint(order.CarId)
	waybill.DriverID = uint(order.DriverId)
	waybill.DepartureDate = order.SendDate
	waybill.ReturnDate = order.ArriveDate
	waybill.ContractID = contractId

	_, err = os.waybillRepository.CreateNewWaybill(&waybill)
	if err != nil {
		return err
	}

	return nil
}
