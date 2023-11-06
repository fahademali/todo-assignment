package services

import (
	"fmt"
	"todo_service/models"
	"user_service/log"

	"github.com/golang-jwt/jwt/v5"
)

type ITokenService interface {
	GetInfoFromToken(accessToken string) (models.UserInfo, error)
}

type TokenService struct {
	secretKey []byte
}

func NewTokenService(secretKey string) ITokenService {
	return &TokenService{secretKey: []byte(secretKey)}
}

func (ts *TokenService) GetInfoFromToken(accessToken string) (models.UserInfo, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return ts.secretKey, nil
	})

	if err != nil {
		return models.UserInfo{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return models.UserInfo{}, fmt.Errorf("invalid token or missing email claim")

	}
	//ASK: Why do i get the need to assert it to float64 even tho the generated id was of type int
	id, ok := claims["id"].(float64)
	log.GetLog().Warn(id)
	log.GetLog().Warn(int64(id))
	if !ok {
		return models.UserInfo{}, fmt.Errorf("invalid token or missing id claim")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return models.UserInfo{}, fmt.Errorf("invalid token or missing email claim")
	}
	isVerified, ok := claims["isVerified"].(bool)
	if !ok {
		return models.UserInfo{}, fmt.Errorf("invalid token or missing isVerified claim")
	}
	validFrom, ok := claims["validFrom"].(float64)
	if !ok {
		return models.UserInfo{}, fmt.Errorf("invalid token or missing validFrom claim")
	}
	role, ok := claims["role"].(string)
	if !ok {
		return models.UserInfo{}, fmt.Errorf("invalid token or missing role claim")
	}

	return models.UserInfo{
		ID:         int64(id),
		Email:      email,
		IsVerified: isVerified,
		ValidFrom:  validFrom,
		Role:       role,
	}, nil
}
