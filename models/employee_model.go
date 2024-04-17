package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
