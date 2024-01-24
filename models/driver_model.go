package models

import "gorm.io/gorm"

type Driver struct {
	gorm.Model
	FirstName          string `json:"firstName"`
	LastName           string `json:"lastName"`
	LicenseNumber      string `json:"licenseNumber"`
	TransportationCert string `json:"transportationCert"`
	CarID              uint   `json:"carId"`
	Car                Car    `gorm:"foreignKey:CarID"`
}
