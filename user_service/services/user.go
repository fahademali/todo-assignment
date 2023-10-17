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
	Login(rb models.LoginRequest) string
	Signup(rb models.SignupRequest) (string, error)
	VerifyUser() error
}

type UserService struct {
	userRepo repo.IUserRepo
}

func NewUserService(userRepo repo.IUserRepo) IUserService {
	return &UserService{userRepo: userRepo}
}

func (u *UserService) Login(rb models.LoginRequest) string {
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

func (u *UserService) Signup(rb models.SignupRequest) (string, error) {
	_, err := u.userRepo.GetUserByEmail(rb.Email)
	if err == nil {
		errMessage := fmt.Errorf("account is already linked with %s", rb.Email)
		return "", errMessage
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rb.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	rb.Password = string(hashedPassword)

	err = u.userRepo.InsertUser(rb)
	if err != nil {
		return "", err
	}

	return "User has been created", nil
}

func (u *UserService) VerifyUser() error {
	//TODO: remove hard code values
	err := u.userRepo.VerifyUser("fahad@gmail.com")
	if err != nil {
		return err
	}
	return nil
}

func comparePasswords(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func generateAccessToken(email string) (string, error) {
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
