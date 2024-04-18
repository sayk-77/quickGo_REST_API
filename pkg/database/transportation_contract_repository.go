package database

import (
	"example.com/go/models"
	"gorm.io/gorm"
)

type TransportationContractRepository interface {
	GetTransportationContractById(transportationContractID int) (*models.TransportationContract, error)
	GetAllTransportationContract() ([]*models.TransportationContract, error)
	CreateNewTransportationContract(newTransportationContract *models.TransportationContract) (uint, error)
}

type TransportationContractImpl struct {
	db *gorm.DB
}

func NewTransportationContractRepository(db *gorm.DB) *TransportationContractImpl {
	return &TransportationContractImpl{
		db: db,
	}
}

func (tcr *TransportationContractImpl) GetTransportationContractById(transportationContractId int) (*models.TransportationContract, error) {
	var transportationContract *models.TransportationContract
	if err := tcr.db.Preload("Car").Preload("Client").Preload("Order").First(&transportationContract, transportationContractId).Error; err != nil {
		return nil, err
	}

	if err := tcr.db.Preload("Client").Preload("CargoType").First(&transportationContract.Order, transportationContract.OrderID).Error; err != nil {
		return nil, err
	}

	return transportationContract, nil
}

func (tcr *TransportationContractImpl) GetAllTransportationContract() ([]*models.TransportationContract, error) {
	var recordTransportationContract []*models.TransportationContract
	if err := tcr.db.Preload("Car").Preload("Client").Preload("Order").Find(&recordTransportationContract).Error; err != nil {
		return nil, err
	}

	for _, contract := range recordTransportationContract {
		if err := tcr.db.Preload("Client").Preload("CargoType").First(&contract.Order, contract.OrderID).Error; err != nil {
			return nil, err
		}
	}

	return recordTransportationContract, nil
}

func (tcr *TransportationContractImpl) CreateNewTransportationContract(newTransportationContract *models.TransportationContract) (uint, error) {
	if err := tcr.db.Create(&newTransportationContract).Error; err != nil {
		return 0, err
	}

	return newTransportationContract.ID, nil
}
