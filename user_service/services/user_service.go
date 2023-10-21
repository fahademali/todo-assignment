package services

import (
	"fmt"
	"user_service/models"
	"user_service/repo"
)

type IUserService interface {
	Login(rb models.LoginRequest) (string, error)
	Signup(rb models.SignupRequest) (string, error)
	VerifyUser(uid string) error
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
	user, err := u.userRepo.GetByEmail(rb.Email)
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

func (u *UserService) Signup(rb models.SignupRequest) (string, error) {
	hashedPassword, err := u.cryptService.GenerateHashPassword(rb.Password)
	if err != nil {
		return "", err
	}

	rb.Password = hashedPassword

	err = u.userRepo.Insert(rb)
	if err != nil {
		return "", err
	}

	token, err := u.tokenService.GenerateAccessToken(rb.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *UserService) VerifyUser(uid string) error {
	user, err := u.userRepo.GetByEmail(uid)
	if err != nil {
		return err
	}
	fmt.Println("VERIFY USER RUNNING SERVICE...........")
	fmt.Println(user)

	user.IsVerified = true

	return u.userRepo.UpdateByEmail(uid, user)
}
