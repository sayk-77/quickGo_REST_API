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

func (ds *DriverService) EditDriver(editedDriver *models.Driver) error {
	var driver *models.Driver
	driver, err := ds.driverRepository.GetDriverById(int(editedDriver.ID))
	if err != nil {
		return err
	}

	driver.FirstName = editedDriver.FirstName
	driver.LastName = editedDriver.LastName
	driver.LicenseNumber = editedDriver.LicenseNumber
	driver.TransportationCert = editedDriver.TransportationCert
	driver.TransportationCertDate = editedDriver.TransportationCertDate
	driver.Status = editedDriver.Status

	if err := ds.driverRepository.UpdateDriver(driver); err != nil {
		return err
	}

	return nil
}

func (ds *DriverService) DeleteDriver(driverID int) error {
	return ds.driverRepository.DeleteDriverById(driverID)
}
