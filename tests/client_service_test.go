package tests

import (
	"testing"

	"example.com/go/models"
	"example.com/go/pkg/service"
)

type MockRepository struct {
	clients []*models.Client
}

func (m *MockRepository) GetClientById(clientID int) (*models.Client, error) {
	if len(m.clients) > 0 {
		return m.clients[0], nil
	}
	return nil, nil
}

func (m *MockRepository) GetAllClient() ([]*models.Client, error) {
	return m.clients, nil
}

func (m *MockRepository) CreateNewClient(newClient *models.Client) (*models.Client, error) {
	m.clients = append(m.clients, newClient)
	return newClient, nil
}

func (m *MockRepository) ClientUpdateData(updateClientData models.ClientResponse) error {
	return nil
}

func (m *MockRepository) ClientChangePassword(currentPassword string, newPassword string, id int) error {
	return nil
}

func (m *MockRepository) ClientFindByEmail(email string) error {
	return nil
}

func (m *MockRepository) ClientPasswordRecovery(email string, password string) error {
	return nil
}

func (m *MockRepository) DeleteClient(id int) error {
	return nil
}

func (m *MockRepository) ClientLogin(email string) (*models.Client, error) {
	for _, client := range m.clients {
		if client.Email == email {
			return client, nil
		}
	}
	return nil, nil
}

func TestClientGetById_Success(t *testing.T) {
	mockRepo := &MockRepository{
		clients: []*models.Client{{FirstName: "John", LastName: "Doe"}},
	}

	clientService := service.NewClientService(mockRepo)

	clientID := 1
	client, err := clientService.GetClientById(clientID)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if client == nil {
		t.Error("Expected a client, got nil")
	}
}

func TestClientGetAll_Success(t *testing.T) {
	mockRepo := &MockRepository{
		clients: []*models.Client{
			{FirstName: "John", LastName: "Doe"},
			{FirstName: "Jane", LastName: "Doe"},
		},
	}

	clientService := service.NewClientService(mockRepo)

	allClients, err := clientService.GetAllClient()

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if len(allClients) == 0 {
		t.Error("Expected at least one client, got none")
	}
}

func TestClientCreateNew_Success(t *testing.T) {
	mockRepo := &MockRepository{}

	clientService := service.NewClientService(mockRepo)

	newClient := &models.Client{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		Password:  "password123",
	}

	createdClient, err := clientService.CreateNewClient(newClient)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if createdClient == nil {
		t.Error("Expected a created client, got nil")
	}
}
