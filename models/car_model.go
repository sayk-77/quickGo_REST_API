package models

import "gorm.io/gorm"

type Car struct {
	gorm.Model
	Brand           string `json:"brend"`
	CarModel        string `json:"model"`
	Year            int    `json:"yearOfIssue"`
	TechnicalStatus string `json:"technicalStatus"`
}
