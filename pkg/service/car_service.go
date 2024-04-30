package service

import (
	"example.com/go/models"
	"example.com/go/pkg/database"
)

type CarService struct {
	carRepository database.CarRepository
}

func NewCarService(carRepository database.CarRepository) *CarService {
	return &CarService{carRepository: carRepository}
}

func (cs *CarService) GetCarById(carID int) (*models.Car, error) {
	return cs.carRepository.GetCarById(carID)
}

func (cs *CarService) CreateNewCar(newCar *models.Car) (*models.Car, error) {
	return cs.carRepository.CreateNewCar(newCar)
}

func (cs *CarService) GetAllCar() ([]*models.Car, error) {
	return cs.carRepository.GetAllCar()
}

func (cs *CarService) FindFreeCar() ([]*models.Car, error) {
	return cs.carRepository.FindFreeCars()
}
