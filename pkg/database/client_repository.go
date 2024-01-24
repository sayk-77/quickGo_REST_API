package database

import (
	"example.com/go/models"
	"gorm.io/gorm"
)

type ClientRepository interface {
	GetClientById(clientID int) (*models.Client, error)
	GetAllClient() ([]*models.Client, error)
	CreateNewClient(newClient *models.Client) (*models.Client, error)
}

type ClientRepositoryImpl struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) *ClientRepositoryImpl {
	return &ClientRepositoryImpl{
		db: db,
	}
}

func (cr *ClientRepositoryImpl) GetClientById(clientID int) (*models.Client, error) {
	var client models.Client
	if err := cr.db.First(&client, clientID).Error; err != nil {
		return nil, err
	}

	return &client, nil
}

func (cr *ClientRepositoryImpl) GetAllClient() ([]*models.Client, error) {
	var clientRecord []*models.Client
	if err := cr.db.Find(&clientRecord).Error; err != nil {
		return nil, err
	}

	return clientRecord, nil
}

func (cr *ClientRepositoryImpl) CreateNewClient(newClient *models.Client) (*models.Client, error) {
	if err := cr.db.Create(newClient).Error; err != nil {
		return nil, err
	}

	return newClient, nil
}
