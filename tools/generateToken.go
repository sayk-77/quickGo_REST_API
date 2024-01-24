package tools

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(id uint, secretKey []byte) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 48).Unix(),
	}
	token.Claims = claims

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("error creating token: %v", err)
	}

	return tokenString, nil
}
