package database

import (
	"errors"
	"example.com/go/models"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type ClientRepository interface {
	GetClientById(clientID int) (*models.Client, error)
	GetAllClient() ([]*models.Client, error)
	CreateNewClient(newClient *models.Client) (*models.Client, error)
	ClientLogin(email string) (*models.Client, error)
	ClientUpdateData(updateClientData models.ClientResponse) error
	ClientChangePassword(currentPassword string, newPassword string, id int) error
	ClientFindByEmail(email string) error
	ClientPasswordRecovery(email string, password string) error
	DeleteClient(id int) error
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
	var client models.Client
	if err := cr.db.Where("email = ?", newClient.Email).First(&client).Error; err == nil {
		return nil, errors.New("Пользователь с указанными данными уже существует")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newClient.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %v", err)
	}

	newClient.Password = string(hashedPassword)

	if err := cr.db.Create(newClient).Error; err != nil {
		return nil, fmt.Errorf("error creating user: %v", err)
	}

	return newClient, nil
}

func (cr *ClientRepositoryImpl) ClientLogin(email string) (*models.Client, error) {
	var client models.Client

	if err := cr.db.Where("email = ?", email).First(&client).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Пользователь с такой электронной почтой не найден")
		}
		return nil, fmt.Errorf("Ошибка при поиске пользователя: %v", err)
	}

	return &client, nil
}

func (cr *ClientRepositoryImpl) ClientUpdateData(updateClientData models.ClientResponse) error {
	var client models.Client

	if err := cr.db.First(&client, updateClientData.ID).Error; err != nil {
		return err
	}

	client.FirstName = updateClientData.FirstName
	client.LastName = updateClientData.LastName
	client.Address = updateClientData.Address
	client.PhoneNumber = updateClientData.PhoneNumber
	client.Email = updateClientData.Email

	if err := cr.db.Save(&client).Error; err != nil {
		return err
	}

	return nil
}

func (cr *ClientRepositoryImpl) ClientChangePassword(currentPassword string, newPassword string, id int) error {
	var client models.Client
	if err := cr.db.First(&client, id).Error; err != nil {
		return fmt.Errorf("Пользователь не найден")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(client.Password), []byte(currentPassword)); err != nil {
		return fmt.Errorf("Текущий пароль не верный")
	}

	hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("Ошибка хеширования пароля")
	}

	client.Password = string(hashedNewPassword)
	if err := cr.db.Save(&client).Error; err != nil {
		return fmt.Errorf("Ошибка при сохранении пароля")
	}

	return nil
}

func (cr *ClientRepositoryImpl) ClientFindByEmail(email string) error {
	var client models.Client

	if err := cr.db.Where("email = ?", email).First(&client).Error; err != nil {
		return fmt.Errorf("Пользователь не найден")
	}

	return nil
}

func (cr *ClientRepositoryImpl) ClientPasswordRecovery(email string, password string) error {
	var client models.Client

	if err := cr.db.Where("email = ?", email).First(&client).Error; err != nil {
		return fmt.Errorf("Не найден")
	}

	client.Password = password

	if err := cr.db.Save(&client).Error; err != nil {
		return fmt.Errorf("Ошибка при изменении пароля")
	}

	return nil
}

func (cr *ClientRepositoryImpl) DeleteClient(id int) error {
	if err := cr.db.Delete(&models.Client{}, id).Error; err != nil {
		return err
	}
	return nil
}
