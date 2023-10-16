package services

import (
	"fmt"

	"user_service/models"
	"user_service/repo"
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
	fmt.Println(user)
	return "Login Service not implemented"
}

func (u *UserService) Signup(rb models.SignupRequestBody) string {
	fmt.Println(rb)
	return "SignUp not implemented yet"
}
