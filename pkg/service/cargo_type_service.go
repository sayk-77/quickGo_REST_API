package service

import (
	"example.com/go/models"
	"example.com/go/pkg/database"
)

type CargoTypeService struct {
	cargoTypeRepository database.CargoTypeRepository
}

func NewCargoTypeService(cargoTypeRepository database.CargoTypeRepository) *CargoTypeService {
	return &CargoTypeService{
		cargoTypeRepository: cargoTypeRepository,
	}
}

func (cts *CargoTypeService) GetCargoTypeById(cargoTypeId int) (*models.CargoType, error) {
	return cts.cargoTypeRepository.GetCargoTypeById(cargoTypeId)
}

func (cts *CargoTypeService) GetAllTypeCargo() ([]*models.CargoType, error) {
	return cts.cargoTypeRepository.GetAllCargoType()
}

func (cts *CargoTypeService) CreateNewCargoType(newCargoType *models.CargoType) (*models.CargoType, error) {
	return cts.cargoTypeRepository.CreateNewCargoType(newCargoType)
}

func (cts *CargoTypeService) UpdateCargoType(cargoType *models.CargoType) error {
	return cts.cargoTypeRepository.UpdateCargoType(cargoType)
}

func (cts *CargoTypeService) DeleteCargoType(cargoTypeId int) error {
	return cts.cargoTypeRepository.DeleteCargoType(cargoTypeId)
}
