package service

import (
	"example.com/go/models"
	"example.com/go/pkg/database"
)

type TransportationContractService struct {
	transportationContractRepository database.TransportationContractRepository
}

func NewTransportationContractService(transportationContractService database.TransportationContractRepository) *TransportationContractService {
	return &TransportationContractService{
		transportationContractRepository: transportationContractService,
	}
}

func (tcs *TransportationContractService) GetTransportationContractById(transportationContractId int) (*models.TransportationContract, error) {
	return tcs.transportationContractRepository.GetTransportationContractById(transportationContractId)
}

func (tcs *TransportationContractService) GetAllTransportationContract() ([]*models.TransportationContract, error) {
	return tcs.transportationContractRepository.GetAllTransportationContract()
}

func (tcs *TransportationContractService) CreateNewTransportationContract(newTransportationContract *models.TransportationContract) (uint, error) {
	return tcs.transportationContractRepository.CreateNewTransportationContract(newTransportationContract)
}
