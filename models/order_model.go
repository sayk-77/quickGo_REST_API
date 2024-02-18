package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ClientID    uint      `json:"clientId"`
	CargoTypeID uint      `json:"cargoTypeId"`
	Status      string    `json:"status"`
	OrderDate   string    `json:"orderDate"`
	Client      Client    `gorm:"foreignKey:ClientID"`
	CargoType   CargoType `gorm:"foreignKey:CargoTypeID"`
	OrderPrice  int       `json:"orderPrice"`
}
