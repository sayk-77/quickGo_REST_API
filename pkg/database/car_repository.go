package database

import (
	"example.com/go/models"
	"gorm.io/gorm"
)

type CarRepository interface {
	GetCarById(carID int) (*models.Car, error)
	CreateNewCar(newCar *models.Car) (*models.Car, error)
	GetAllCar() ([]*models.Car, error)
	FindFreeCars() ([]*models.Car, error)
}

type CarRepositoryImpl struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) *CarRepositoryImpl {
	return &CarRepositoryImpl{
		db: db,
	}
}

func (cr *CarRepositoryImpl) GetCarById(carID int) (*models.Car, error) {
	var car models.Car
	if err := cr.db.First(&car, carID).Error; err != nil {
		return nil, err
	}

	return &car, nil
}

func (cr *CarRepositoryImpl) CreateNewCar(newCar *models.Car) (*models.Car, error) {
	if err := cr.db.Create(newCar).Error; err != nil {
		return nil, err
	}

	return newCar, nil
}

func (cr *CarRepositoryImpl) GetAllCar() ([]*models.Car, error) {
	var carRecord []*models.Car
	if err := cr.db.Find(&carRecord).Error; err != nil {
		return nil, err
	}
	return carRecord, nil
}

func (cr *CarRepositoryImpl) FindFreeCars() ([]*models.Car, error) {
	var cars []*models.Car
	if err := cr.db.
		Not("id IN (?)", cr.db.Table("drivers").Select("car_id").Distinct()).
		Find(&cars).Error; err != nil {
		return nil, err
	}
	return cars, nil
}
