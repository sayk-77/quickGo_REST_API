package tools

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
)

type CustomClaims struct {
	jwt.StandardClaims
}

func Decoder(tokenString string) (map[string]interface{}, error) {
	if tokenString == "" {
		return nil, fmt.Errorf("Token is missing")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ACCESS_TOKEN_SECRET_KEY")), nil
	})

	if err != nil {
		return nil, fmt.Errorf("Invalid Token")
	}

	if !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("Invalid data in the token")
	}

	return claims, nil
}
