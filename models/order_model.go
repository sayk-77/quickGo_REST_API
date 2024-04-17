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
