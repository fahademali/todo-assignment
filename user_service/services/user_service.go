package services

import (
	"context"
	"fmt"
	"user_service/config"
	"user_service/log"
	"user_service/models"
	"user_service/repo"
)

type IUserService interface {
	Login(requestBody models.LoginRequest) (string, error)
	Signup(ctx context.Context, requestBody models.SignupRequest) error
	VerifyUser(email string) error
	GrantAdminRole(email string) error
}

type UserService struct {
	userRepo     repo.IUserRepo
	cryptService ICryptService
	tokenService ITokenService
	emailService IEmailService
}

func NewUserService(userRepo repo.IUserRepo, cryptService ICryptService, tokenService ITokenService, emailService IEmailService) IUserService {
	return &UserService{userRepo: userRepo, cryptService: cryptService, tokenService: tokenService, emailService: emailService}
}

func (u *UserService) Login(requestBody models.LoginRequest) (string, error) {
	user, err := u.userRepo.GetByEmail(requestBody.Email)
	if err != nil {
		return "", err
	}

	err = u.cryptService.ComparePasswords(user.Password, requestBody.Password)
	if err != nil {
		return "", err
	}

	token, err := u.tokenService.GenerateAccessToken(user.Email, user.Role, user.IsVerified)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *UserService) Signup(ctx context.Context, requestBody models.SignupRequest) error {
	hashedPassword, err := u.cryptService.GenerateHashPassword(requestBody.Password)
	if err != nil {
		return err
	}

	requestBody.Password = hashedPassword

	tx, err := u.userRepo.ExecTx(ctx)
	if err != nil {
		return fmt.Errorf("SignupTx: %v", err)
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	err = u.userRepo.Insert(ctx, tx, requestBody)
	if err != nil {
		return fmt.Errorf("SignupTx: %v", err)
	}

	token, err := u.tokenService.GenerateAccessToken(requestBody.Email, "user", false)
	if err != nil {
		return err
	}

	verificationLink := fmt.Sprintf("%s/verify-user/%s", config.AppConfig.BASE_URL, token)
	emailBody := `
	<html>
		<body>
			<p>Hello!</p>
			<p>Click the following link to verify your email address:</p>
			<p><a href="` + verificationLink + `">Verify Email Address</a></p>
		</body>
	</html>
	`

	err = u.emailService.SendEmail(requestBody.Email, "Verfify Email Address", emailBody)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (u *UserService) VerifyUser(email string) error {
	user, err := u.userRepo.GetByEmail(email)
	if err != nil {
		return err
	}

	user.IsVerified = true
	log.GetLog().Info(user)
	return u.userRepo.UpdateByEmail(email, user)
}

func (u *UserService) GrantAdminRole(email string) error {
	err := u.userRepo.SetAdminRole(email)
	if err != nil {
		return err
	}
	return nil
}
