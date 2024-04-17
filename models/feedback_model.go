package models

import "gorm.io/gorm"

type Feedback struct {
	gorm.Model
	Name        string `json:"name"`
	NumberPhone string `json:"numberPhone"`
	Email       string `json:"email"`
	Question    string `json:"question"`
	Status      string `json:"status"`
}
