package database

import (
	"errors"
	"example.com/go/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	GetOrderById(orderID int) (*models.Order, error)
	GetAllOrder() ([]*models.Order, error)
	CreateNewOrder(newOrder *models.Order) (*models.Order, error)
	GetOrdersByStatus(clientId int, status string) ([]*models.Order, error)
	DeleteOrderById(orderId int) error
	UpdateOrderStatus(orderId int, status string, deliveryDate string) error
	CompleteOrder(orderId int) error
}

type OrderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepositoryImpl {
	return &OrderRepositoryImpl{
		db: db,
	}
}

func (or *OrderRepositoryImpl) GetOrderById(orderID int) (*models.Order, error) {
	var order models.Order
	if err := or.db.Preload("Client").Preload("CargoType").First(&order, orderID).Error; err != nil {
		return nil, err
	}

	return &order, nil
}

func (or *OrderRepositoryImpl) GetAllOrder() ([]*models.Order, error) {
	var orderRecord []*models.Order
	if err := or.db.Preload("Client").Preload("CargoType").Find(&orderRecord).Error; err != nil {
		return nil, err
	}

	return orderRecord, nil
}

func (or *OrderRepositoryImpl) CreateNewOrder(newOrder *models.Order) (*models.Order, error) {
	if err := or.db.Create(newOrder).Error; err != nil {
		return nil, err
	}

	createdOrder := &models.Order{}
	if err := or.db.Preload("Client").Preload("CargoType").First(createdOrder, newOrder.ID).Error; err != nil {
		return nil, err
	}

	return createdOrder, nil
}

func (or *OrderRepositoryImpl) GetOrdersByStatus(clientId int, status string) ([]*models.Order, error) {
	var ordersRecord []*models.Order
	if err := or.db.Preload("CargoType").Where("client_id = ? and status = ?", clientId, status).Find(&ordersRecord).Error; err != nil {
		return nil, err
	}

	return ordersRecord, nil
}

func (or *OrderRepositoryImpl) DeleteOrderById(orderId int) error {
	condition := or.db.Where("id = ?", orderId)

	result := condition.Delete(&models.Order{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("Запись не найдена")
	}

	return nil
}

func (or *OrderRepositoryImpl) UpdateOrderStatus(orderId int, status string, deliveryDate string) error {
	order, err := or.GetOrderById(orderId)
	if err != nil {
		return err
	}

	order.Status = status
	order.DeliveryDate = deliveryDate

	if err := or.db.Save(&order).Error; err != nil {
		return err
	}

	return nil
}

func (or *OrderRepositoryImpl) CompleteOrder(orderId int) error {
	var order models.Order
	if err := or.db.Preload("Client").Preload("CargoType").First(&order, orderId).Error; err != nil {
		return err
	}
	order.Status = "Завершен"
	if err := or.db.Save(&order).Error; err != nil {
		return err
	}
	var contract models.TransportationContract
	if err := or.db.Where("order_id = ?", orderId).First(&contract).Error; err != nil {
		return err
	}

	var waybill models.WayBill
	if err := or.db.Where("contract_id = ?", contract.ID).First(&waybill).Error; err != nil {
		return err
	}

	var driver models.Driver
	if err := or.db.Where("id =?", waybill.DriverID).First(&driver).Error; err != nil {
		return err
	}

	driver.Status = "Свободен"
	if err := or.db.Save(&driver).Error; err != nil {
		return err
	}

	if err := or.db.Delete(&contract).Error; err != nil {
		return err
	}

	if err := or.db.Delete(&waybill).Error; err != nil {
		return err
	}

	return nil
}
