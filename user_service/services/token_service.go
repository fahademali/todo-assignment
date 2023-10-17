package services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ITokenService interface {
	GenerateAccessToken(email string) (string, error)
}

type TokenService struct {
}

func NewTokenService() ITokenService {
	return &TokenService{}
}


func (ts *TokenService) GenerateAccessToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"nbf":   time.Now().Unix(),
	})
	//TODO: get the key from .env variables
	tokenString, err := token.SignedString([]byte("mysecretkey"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
