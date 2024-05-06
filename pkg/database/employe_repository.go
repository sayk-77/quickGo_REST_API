package database

import (
	"errors"
	"example.com/go/models"
	"fmt"
	"gorm.io/gorm"
)

type EmployeeRepositoryImpl interface {
	AddEmployee(employee *models.Employee) (string, error)
	GetEmployeesByEmail(id string) (*models.Employee, error)
}

type EmployeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (r *EmployeeRepository) AddEmployee(employee *models.Employee) (string, error) {
	if err := r.db.Where("email = ?", employee.Email).First(&models.Employee{}).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return "", err
		}
		if err := r.db.Create(employee).Error; err != nil {
			return "", err
		}
		return "Успешно добавлен", nil
	}
	return "Данный email уже зарегистрирован", nil
}

func (r *EmployeeRepository) GetEmployeesByEmail(email string) (*models.Employee, error) {
	var employee models.Employee
	if err := r.db.Where("email =?", email).First(&employee).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Данные не верны")
		}
		return nil, err
	}
	return &employee, nil
}
