package models

import "gorm.io/gorm"

type WayBill struct {
	gorm.Model
	DriverID               uint                   `json:"driverId"`
	CarID                  uint                   `json:"carId"`
	ContractID             uint                   `json:"contractId"`
	DepartureDate          string                 `json:"departureDate"`
	ReturnDate             string                 `json:"returnDate"`
	Driver                 Driver                 `gorm:"foreignKey:DriverID"`
	Car                    Car                    `gorm:"foreignKey:CarID"`
	TransportationContract TransportationContract `gorm:"foreignKey:ContractID"`
}
