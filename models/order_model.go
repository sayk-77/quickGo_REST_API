package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ClientID           uint      `json:"clientId"`
	CargoTypeID        uint      `json:"cargoTypeId"`
	Recipient          string    `json:"recipient"`
	Status             string    `json:"status"`
	OrderDate          string    `json:"orderDate"`
	DestinationAddress string    `json:"destinationAddress"`
	SendingAddress     string    `json:"sendingAddress"`
	DeliveryDate       string    `json:"deliveryDate"`
	OrderPrice         int       `json:"orderPrice"`
	Client             Client    `gorm:"foreignKey:ClientID"`
	CargoType          CargoType `gorm:"foreignKey:CargoTypeID"`
}

type OrderConfirm struct {
	gorm.Model
	DeliveryDate string `json:"deliveryDate"`
	SendDate     string `json:"sendDate"`
	ArriveDate   string `json:"arriveDate"`
	DriverId     int    `json:"driverId"`
	CarId        int    `json:"carId"`
	ExpiryDate   string `json:"expiryDate"`
	ClientID     int    `json:"clientId"`
}
