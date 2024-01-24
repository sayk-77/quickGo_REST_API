package service

import (
	"example.com/go/models"
	"example.com/go/pkg/database"
)

type ClientService struct {
	clientRepository database.ClientRepository
}

func NewClientService(clientRepository database.ClientRepository) *ClientService {
	return &ClientService{
		clientRepository: clientRepository,
	}
}

func (cs *ClientService) GetClientById(clientID int) (*models.Client, error) {
	return cs.clientRepository.GetClientById(clientID)
}

func (cs *ClientService) GetAllClient() ([]*models.Client, error) {
	return cs.clientRepository.GetAllClient()
}

func (cs *ClientService) CreateNewClient(newClient *models.Client) (*models.Client, error) {
	return cs.clientRepository.CreateNewClient(newClient)
}

func (cs *ClientService) ClientLogin(email string, password string) (*models.Tokens, error) {
	return cs.clientRepository.ClientLogin(email, password)
}
