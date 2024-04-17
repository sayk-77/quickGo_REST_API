package models

import "gorm.io/gorm"

type Car struct {
	gorm.Model
	Brand           string  `json:"brend"`
	CarModel        string  `json:"model"`
	Year            int     `json:"year"`
	Color           string  `json:"color"`
	Mileage         float32 `json:"mileage"`
	ImageUrl        string  `json:"imageUrl"`
	TechnicalStatus string  `json:"technicalStatus"`
}
