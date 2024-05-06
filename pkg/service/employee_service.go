package service

import (
	"errors"
	"example.com/go/models"
	"example.com/go/pkg/database"
	"example.com/go/tools"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"os"
)

type EmployeeService struct {
	employeeRepository *database.EmployeeRepository
}

func NewEmployeeService(employeeRepository *database.EmployeeRepository) *EmployeeService {
	return &EmployeeService{
		employeeRepository: employeeRepository,
	}
}

func (e *EmployeeService) Create(employee *models.Employee) (string, error) {
	if employee == nil {
		return "", errors.New("некорректные данные сотрудника")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(employee.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	employee.Password = string(hashedPassword)

	msg, err := e.employeeRepository.AddEmployee(employee)
	if err != nil {
		return "", err
	}
	return msg, nil
}

func (e *EmployeeService) Login(email string, password string) (string, error) {
	employee, err := e.employeeRepository.GetEmployeesByEmail(email)
	if err != nil {
		return "", err
	}
	if employee == nil {
		return "", fmt.Errorf("Данные не верны")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(employee.Password), []byte(password)); err != nil {
		return "", fmt.Errorf("Данные не верны")
	}

	accessTokenSecretKey := []byte(os.Getenv("ACCESS_TOKEN_SECRET_KEY"))

	token, err := tools.GenerateTokenAdmin(employee.ID, employee.Role, accessTokenSecretKey)
	return token, err
}
