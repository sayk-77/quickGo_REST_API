package service

import (
	"example.com/go/models"
	"example.com/go/pkg/database"
	"example.com/go/tools"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"os"
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
	client, err := cs.clientRepository.ClientLogin(email)
	if err != nil {
		return "", fmt.Errorf("Пользователь не найден")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(client.Password), []byte(password)); err != nil {
		return "", fmt.Errorf("Данные не верны")
	}

	accessToken, err := cs.generateAccessToken(int(client.ID))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (cs *ClientService) generateAccessToken(clientID int) (string, error) {
	accessTokenSecretKey := []byte(os.Getenv("ACCESS_TOKEN_SECRET_KEY"))

	accessToken, err := tools.GenerateToken(uint(clientID), accessTokenSecretKey)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (cs *ClientService) ClientUpdateData(updateClientData models.ClientResponse) error {
	return cs.clientRepository.ClientUpdateData(updateClientData)
}

func (cs *ClientService) ClientChangePassword(currentPassword string, newPassword string, id int) error {
	return cs.clientRepository.ClientChangePassword(currentPassword, newPassword, id)
}

func (cs *ClientService) DeleteClient(id int) error {
	return cs.clientRepository.DeleteClient(id)
}
