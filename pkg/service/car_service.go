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

func (cs *CarService) UpdateCar(newCar *models.Car) error {
	car, err := cs.carRepository.GetCarById(int(newCar.ID))
	if err != nil {
		return err
	}

	car.Brand = newCar.Brand
	car.CarModel = newCar.CarModel
	car.Year = newCar.Year
	car.Color = newCar.Color
	car.Mileage = newCar.Mileage
	car.TechnicalStatus = newCar.TechnicalStatus
	car.ImageUrl = newCar.ImageUrl

	if err := cs.carRepository.SaveCar(car); err != nil {
		return err
	}
	return nil
}

func (cs *CarService) DeleteCar(carID int) error {
	return cs.carRepository.DeleteCar(carID)
}
