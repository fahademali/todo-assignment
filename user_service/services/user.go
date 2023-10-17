package services

import (
	"fmt"

	"user_service/models"
	"user_service/repo"
)

type IUserService interface {
	Login(rb models.LoginRequest) (string, error)
	Signup(rb models.SignupRequest) error
	VerifyUser() error
}

type UserService struct {
	userRepo     repo.IUserRepo
	cryptService ICryptService
	tokenService ITokenService
}

func NewUserService(userRepo repo.IUserRepo, cryptService ICryptService, tokenService ITokenService) IUserService {
	return &UserService{userRepo: userRepo, cryptService: cryptService, tokenService: tokenService}
}

func (u *UserService) Login(rb models.LoginRequest) (string, error) {
	user, err := u.userRepo.GetUserByEmail(rb.Email)
	if err != nil {
		return "", err
	}

	err = u.cryptService.ComparePasswords(user.Password, rb.Password)
	if err != nil {
		return "", err
	}

	token, err := u.tokenService.GenerateAccessToken(rb.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *UserService) Signup(rb models.SignupRequest) error {
	_, err := u.userRepo.GetUserByEmail(rb.Email)
	if err == nil {
		errMessage := fmt.Errorf("account is already linked with %s", rb.Email)
		return errMessage
	}

	hashedPassword, err := u.cryptService.GenerateHashPassword(rb.Password)
	if err != nil {
		return err
	}

	rb.Password = hashedPassword

	err = u.userRepo.InsertUser(rb)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserService) VerifyUser() error {
	//TODO: remove hard code values
	err := u.userRepo.VerifyUser("fahad@gmail.com")
	if err != nil {
		return err
	}
	return nil
}
