package models

import "gorm.io/gorm"

type TransportationContract struct {
	gorm.Model
	OrderID      uint   `json:"orderId"`
	ClientID     uint   `json:"clientId"`
	CarID        uint   `json:"carId"`
	ContractDate string `json:"contractDate"`
	ExpiryDate   string `json:"expiryDate"`
	Client       Client `gorm:"foreignKey:ClientID"`
	Car          Car    `gorm:"foreignKey:CarID"`
	Order        Order  `gorm:"foreignKey:OrderID;association_autoupdate:false"`
}
