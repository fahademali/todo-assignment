package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ITokenService interface {
	GenerateAccessToken(email string, role string, isVerified bool, secretkey string) (string, error)
	GetEmailFromAccessToken(accessToken string, secretKey string) (string, error)
}

type TokenService struct {
}

func NewTokenService() ITokenService {
	return &TokenService{}
}

func (ts *TokenService) GenerateAccessToken(email string, role string, isVerified bool, secretkey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":      email,
		"role":       role,
		"isVerified": isVerified,
		"nbf":        time.Now().Unix(),
	})
	tokenString, err := token.SignedString([]byte(secretkey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (ts *TokenService) GetEmailFromAccessToken(accessToken string, secretKey string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email := claims["email"].(string)
		return email, nil
	}

	return "", fmt.Errorf("invalid token or missing email claim")
}
