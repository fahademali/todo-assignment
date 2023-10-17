package services

import "golang.org/x/crypto/bcrypt"

type ICryptService interface {
	ComparePasswords(hashedPassword string, password string) error
	GenerateHashPassword(password string) (string, error)
}

type CryptService struct {
}

func NewCryptService() ICryptService {
	return &CryptService{}
}

func (cs *CryptService) ComparePasswords(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (cs *CryptService) GenerateHashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
