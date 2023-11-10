package services

import (
	"fmt"
	"time"
	"user_service/log"
	"user_service/models"

	"github.com/golang-jwt/jwt/v5"
)

type ITokenService interface {
	GenerateAccessToken(user models.User) (string, error)
	GetEmailFromAccessToken(accessToken string) (string, error)
	GetInfoFromToken(accessToken string) (models.UserInfo, error)
}

type TokenService struct {
	secretKey []byte
}

func NewTokenService(secretKey string) ITokenService {
	return &TokenService{secretKey: []byte(secretKey)}
}

func (ts *TokenService) GenerateAccessToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         user.ID,
		"email":      user.Email,
		"role":       user.Role,
		"isVerified": user.IsVerified,
		"validFrom":  time.Now().Unix(),
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

	log.GetLog().Info("claims")
	log.GetLog().Info(claims)
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
		Email:      email,
		IsVerified: isVerified,
		ValidFrom:  validFrom,
		Role:       role,
	}, nil

}
