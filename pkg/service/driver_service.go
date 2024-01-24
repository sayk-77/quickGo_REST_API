package service

import (
	"example.com/go/models"
	"example.com/go/pkg/database"
)

type DriverService struct {
	driverRepository database.DriverRepository
}

func NewDriverService(driverRepository database.DriverRepository) *DriverService {
	return &DriverService{
		driverRepository: driverRepository,
	}
}

func (ds *DriverService) GetDriverById(driverID int) (*models.Driver, error) {
	return ds.driverRepository.GetDriverById(driverID)
}

func (ds *DriverService) GetAllDriver() ([]*models.Driver, error) {
	return ds.driverRepository.GetAllDriver()
}

func (ds *DriverService) CreateNewDriver(newDriver *models.Driver) (*models.Driver, error) {
	return ds.driverRepository.CreateNewDriver(newDriver)
}
