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
	emailService       *EmailService
	clientRepository   database.ClientRepository
}

func NewOrderService(orderRepository database.OrderRepository, driverRepository database.DriverRepository, contractRepository database.TransportationContractRepository, waybillRepository database.WaybillRepository, emailService *EmailService, clientRepository database.ClientRepository) *OrderService {
	return &OrderService{
		orderRepository:    orderRepository,
		driverRepository:   driverRepository,
		contractRepository: contractRepository,
		waybillRepository:  waybillRepository,
		emailService:       emailService,
		clientRepository:   clientRepository,
	}
}

func (os *OrderService) GetOrderById(orderID int) (*models.Order, error) {
	return os.orderRepository.GetOrderById(orderID)
}

func (os *OrderService) GetAllOrder() ([]*models.Order, error) {
	return os.orderRepository.GetAllOrder()
}

func (os *OrderService) CreateNewOrder(newOrder *models.Order) (*models.Order, error) {
	client, err := os.clientRepository.GetClientById(int(newOrder.ClientID))
	if err != nil {
		return nil, err
	}

	order, err := os.orderRepository.CreateNewOrder(newOrder)
	if err != nil {
		return nil, err
	}

	err = os.emailService.SendOrderCreateMail(order, client.Email)
	if err != nil {
		return nil, err
	}

	return order, nil
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
	contract.ExpiryDate = order.DeliveryDate

	contractId, err := os.contractRepository.CreateNewTransportationContract(&contract)
	if err != nil {
		return err
	}

	var newContract *models.TransportationContract

	newContract, err = os.contractRepository.GetTransportationContractById(int(contractId))
	if err != nil {
		return err
	}

	var waybill models.WayBill
	waybill.CarID = uint(order.CarId)
	waybill.DriverID = uint(order.DriverId)
	waybill.DepartureDate = order.SendDate
	waybill.ReturnDate = order.ArriveDate
	waybill.ContractID = contractId

	waybillCreate, err := os.waybillRepository.CreateNewWaybill(&waybill)
	if err != nil {
		return err
	}

	var driverName string = waybillCreate.Driver.FirstName + " " + waybillCreate.Driver.LastName

	if err := os.emailService.SendConfirmOrderMail(int(order.ID), order.SendDate, order.DeliveryDate, driverName, newContract.Client.Email); err != nil {
		return err
	}

	return nil
}
