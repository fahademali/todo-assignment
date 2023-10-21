package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ITokenService interface {
	GenerateAccessToken(email string) (string, error)
	GetEmailFromAccessToken(accessToken string) (string, error)
}

type TokenService struct {
	secretKey []byte
}

func NewTokenService(secretKey string) ITokenService {
	return &TokenService{secretKey: []byte(secretKey)}
}

func (ts *TokenService) GenerateAccessToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"nbf":   time.Now().Unix(),
	})
	fmt.Println("ts.secretKey")
	fmt.Println(ts.secretKey)
	tokenString, err := token.SignedString(ts.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (ts *TokenService) GetEmailFromAccessToken(accessToken string) (string, error) {
	fmt.Println("runnnning GetEmailFromAccessToken...............")

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return ts.secretKey, nil
	})
	fmt.Println(ts.secretKey)

	if err != nil {
		return "", err
	}
	fmt.Println("token")
	fmt.Println(token)

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email := claims["email"].(string)
		return email, nil
	}

	return "", fmt.Errorf("invalid token or missing email claim")
}
