package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Address     string `json:"address"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

type ClientResponse struct {
	ID          uint   `json:"ID"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Address     string `json:"address"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}
