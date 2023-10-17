package services

import (
	"fmt"
	"time"

	"user_service/models"
	"user_service/repo"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	Login(rb models.LoginRequestBody) string
	Signup(rb models.SignupRequestBody) string
}

type UserService struct {
	userRepo repo.IUserRepo
}

func NewUserService(userRepo repo.IUserRepo) IUserService {
	return &UserService{userRepo: userRepo}
}

func (u *UserService) Login(rb models.LoginRequestBody) string {
	user, err := u.userRepo.GetUserByEmail(rb.Email)
	if err != nil {
		return err.Error()
	}

	err = comparePasswords(user.Password, rb.Password)
	if err != nil {
		return err.Error()
	}

	token, err := generateAccessToken(rb.Email)
	if err != nil {
		return err.Error()
	}

	return token
}

func (u *UserService) Signup(rb models.SignupRequestBody) string {
	_, err := u.userRepo.GetUserByEmail(rb.Email)
	if err == nil {
		message := fmt.Sprintf("Account is already linked with %v", rb.Email)
		return message
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rb.Password), bcrypt.DefaultCost)
	if err != nil {
		return err.Error()
	}

	rb.Password = string(hashedPassword)
	_, err = u.userRepo.InsertUser(rb)
	if err != nil {
		return err.Error()
	}

	return "User has been created"
}

func comparePasswords(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func generateAccessToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"nbf":   time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte("mysecretkey"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
