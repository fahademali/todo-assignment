package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ITokenService interface {
	GenerateAccessToken(email string, role string, isVerified bool) (string, error)
	GetEmailFromAccessToken(accessToken string) (string, error)
	ParseToken(accessToken string, secretKey string) (jwt.MapClaims, error)
}

type TokenService struct {
	secretKey []byte
}

func NewTokenService(secretKey string) ITokenService {
	return &TokenService{secretKey: []byte(secretKey)}
}

func (ts *TokenService) GenerateAccessToken(email string, role string, isVerified bool) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":      email,
		"role":       role,
		"isVerified": isVerified,
		"nbf":        time.Now().Unix(),
	})
	tokenString, err := token.SignedString(ts.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (ts *TokenService) GetEmailFromAccessToken(accessToken string) (string, error) {

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return ts.secretKey, nil
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

func (ts *TokenService) ParseToken(accessToken string, secretKey string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		return claims, nil
	}

	return nil, fmt.Errorf("invalid token or missing email claim")

}
