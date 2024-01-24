package database

import (
	"example.com/go/models"
	"gorm.io/gorm"
)

type DriverRepository interface {
	GetDriverById(driverID int) (*models.Driver, error)
	GetAllDriver() ([]*models.Driver, error)
	CreateNewDriver(newDriver *models.Driver) (*models.Driver, error)
}

type DriverRepositoryImpl struct {
	db *gorm.DB
}

func NewDriverRepository(db *gorm.DB) *DriverRepositoryImpl {
	return &DriverRepositoryImpl{
		db: db,
	}
}

func (dr *DriverRepositoryImpl) GetDriverById(driverID int) (*models.Driver, error) {
	var driver models.Driver
	if err := dr.db.Preload("Car").First(&driver, driverID).Error; err != nil {
		return nil, err
	}

	return &driver, nil
}

func (dr *DriverRepositoryImpl) GetAllDriver() ([]*models.Driver, error) {
	var driverRecord []*models.Driver
	if err := dr.db.Preload("Car").Find(&driverRecord).Error; err != nil {
		return nil, err
	}

	return driverRecord, nil
}

func (dr *DriverRepositoryImpl) CreateNewDriver(newDriver *models.Driver) (*models.Driver, error) {
	if err := dr.db.Create(newDriver).Error; err != nil {
		return nil, err
	}

	if err := dr.db.Preload("Car").First(newDriver, newDriver.ID).Error; err != nil {
		return nil, err
	}

	return newDriver, nil
}
