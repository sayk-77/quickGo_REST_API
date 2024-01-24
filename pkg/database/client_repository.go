package database

import (
	"fmt"
	"os"

	"example.com/go/models"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type ClientRepository interface {
	GetClientById(clientID int) (*models.Client, error)
	GetAllClient() ([]*models.Client, error)
	CreateNewClient(newClient *models.Client) (*models.Client, error)
	ClientLogin(email string, password string) (*models.Tokens, error)
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
		return nil, fmt.Errorf("user with email %s already exists", newClient.Email)
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

func (cr *ClientRepositoryImpl) ClientLogin(email string, password string) (*models.Tokens, error) {
	var client models.Client

	if err := cr.db.Where("email = ?", email).First(&client).Error; err != nil {
		return nil, fmt.Errorf("email not found: %v", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(client.Password), []byte(password)); err != nil {
		return nil, fmt.Errorf("invalid password: %v", err)
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": client.ID,
	})

	secretKey := []byte(os.Getenv("SECRET_KEY_TOKEN"))
	accessTokenString, err := accessToken.SignedString(secretKey)
	if err != nil {
		return nil, fmt.Errorf("error creating access token: %v", err)
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": client.ID,
	})

	refreshTokenString, err := refreshToken.SignedString(secretKey)
	if err != nil {
		return nil, fmt.Errorf("error creating refresh token: %v", err)
	}

	tokens := &models.Tokens{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}

	return tokens, nil
}
