package models

import "gorm.io/gorm"

type CargoType struct {
	gorm.Model
	TypeName    string  `json:"typeName"`
	Description string  `json:"descriptionType"`
	PriceCoeff  float64 `json:"PriceCoeff"`
}
