package database

import (
	"example.com/go/models"
	"gorm.io/gorm"
)

type CargoTypeRepository interface {
	GetCargoTypeById(cargoTypeID int) (*models.CargoType, error)
	GetAllCargoType() ([]*models.CargoType, error)
	CreateNewCargoType(newCargoType *models.CargoType) (*models.CargoType, error)
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
