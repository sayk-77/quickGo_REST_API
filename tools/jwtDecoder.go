package tools

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
)

type CustomClaims struct {
	jwt.StandardClaims
	ClientId int `json:"id"`
}

func Decoder(tokenString string) (int, error) {
	if tokenString == "" {
		return 0, fmt.Errorf("Token is missing")
	}

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ACCESS_TOKEN_SECRET_KEY")), nil
	})

	if err != nil {
		return 0, fmt.Errorf("Invalid Token")
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return 0, fmt.Errorf("Invalid token")
	}

	return claims.ClientId, nil
}
