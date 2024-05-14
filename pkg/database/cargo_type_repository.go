package database

import (
	"example.com/go/models"
	"gorm.io/gorm"
)

type CargoTypeRepository interface {
	GetCargoTypeById(cargoTypeID int) (*models.CargoType, error)
	GetAllCargoType() ([]*models.CargoType, error)
	CreateNewCargoType(newCargoType *models.CargoType) (*models.CargoType, error)
	UpdateCargoType(updatedCargoType *models.CargoType) error
	DeleteCargoType(cargoTypeID int) error
}

type CargoTypeRepositoryImpl struct {
	db *gorm.DB
}

func NewCargoTypeRepository(db *gorm.DB) *CargoTypeRepositoryImpl {
	return &CargoTypeRepositoryImpl{
		db: db,
	}
}

func (cr *CargoTypeRepositoryImpl) GetCargoTypeById(cargoTypeID int) (*models.CargoType, error) {
	var cargoType models.CargoType
	if err := cr.db.First(&cargoType, cargoTypeID).Error; err != nil {
		return nil, err
	}

	return &cargoType, nil
}

func (cr *CargoTypeRepositoryImpl) GetAllCargoType() ([]*models.CargoType, error) {
	var cargoTypeRecord []*models.CargoType
	if err := cr.db.Find(&cargoTypeRecord).Error; err != nil {
		return nil, err
	}

	return cargoTypeRecord, nil
}

func (cr *CargoTypeRepositoryImpl) CreateNewCargoType(newCargoType *models.CargoType) (*models.CargoType, error) {
	if err := cr.db.Create(newCargoType).Error; err != nil {
		return nil, err
	}

	return newCargoType, nil
}

func (cr *CargoTypeRepositoryImpl) UpdateCargoType(updatedCargoType *models.CargoType) error {
	if err := cr.db.Save(updatedCargoType).Error; err != nil {
		return err
	}

	return nil
}

func (cr *CargoTypeRepositoryImpl) DeleteCargoType(cargoTypeID int) error {
	if err := cr.db.Delete(&models.CargoType{}, cargoTypeID).Error; err != nil {
		return err
	}
	return nil
}
