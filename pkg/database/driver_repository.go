package database

import (
	"example.com/go/models"
	"gorm.io/gorm"
)

type DriverRepository interface {
	GetDriverById(driverID int) (*models.Driver, error)
	GetAllDriver() ([]*models.Driver, error)
	CreateNewDriver(newDriver *models.Driver) (*models.Driver, error)
	UpdateStatusDriver(driverID int, status string) error
	UpdateDriver(driver *models.Driver) error
	DeleteDriverById(driverID int) error
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

func (dr *DriverRepositoryImpl) UpdateStatusDriver(driverID int, status string) error {
	driver, err := dr.GetDriverById(driverID)
	if err != nil {
		return err
	}

	driver.Status = status

	if err := dr.db.Save(&driver).Error; err != nil {
		return err
	}

	return nil
}

func (dr *DriverRepositoryImpl) UpdateDriver(driver *models.Driver) error {
	if err := dr.db.Save(&driver).Error; err != nil {
		return err
	}

	return nil
}

func (dr *DriverRepositoryImpl) DeleteDriverById(driverID int) error {
	if err := dr.db.Delete(&models.Driver{}, driverID).Error; err != nil {
		return err
	}

	return nil
}
