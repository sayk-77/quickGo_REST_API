package database

import (
	"example.com/go/models"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.Car{},
		&models.CargoType{},
		&models.Client{},
		&models.Driver{},
		&models.Order{},
		&models.TransportationContract{},
		&models.WayBill{},
	)
}
