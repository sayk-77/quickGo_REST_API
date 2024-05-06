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

func (cs *ClientService) GetAllClient() ([]*models.ClientResponse, error) {
	clientList, err := cs.clientRepository.GetAllClient()
	if err != nil {
		return nil, err
	}

	var clientResponse []*models.ClientResponse
	for _, client := range clientList {
		clientResponse = append(clientResponse, &models.ClientResponse{
			ID:          client.ID,
			FirstName:   client.FirstName,
			LastName:    client.LastName,
			Email:       client.Email,
			PhoneNumber: client.PhoneNumber,
			Address:     client.Address,
		})
	}
	return clientResponse, nil
}

func (cs *ClientService) CreateNewClient(newClient *models.Client) (*models.Client, error) {
	return cs.clientRepository.CreateNewClient(newClient)
}

func (cs *ClientService) ClientLogin(email string, password string) (string, error) {
	return cs.clientRepository.ClientLogin(email, password)
}

func (cs *ClientService) ClientUpdateData(updateClientData models.ClientResponse) error {
	return cs.clientRepository.ClientUpdateData(updateClientData)
}

func (cs *ClientService) ClientChangePassword(currentPassword string, newPassword string, id int) error {
	return cs.clientRepository.ClientChangePassword(currentPassword, newPassword, id)
}
