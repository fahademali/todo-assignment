package services

import (
	"context"
	"fmt"
	"user_service/config"
	"user_service/models"
	"user_service/repo"
)

type IUserService interface {
	Login(rb models.LoginRequest) (string, error)
	Signup(rb models.SignupRequest) (string, error)
	SignupTx(ctx context.Context, rb models.SignupRequest) error
	VerifyUser(uid string) error
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

	fmt.Println("user inserted/.....................")

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
	fmt.Println(verificationLink)

	err = u.emailService.SendEmailTx(rb.Email, "Verfify Email Address", emailBody)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *UserService) SignupTx(ctx context.Context, rb models.SignupRequest) error {
	hashedPassword, err := u.cryptService.GenerateHashPassword(rb.Password)
	if err != nil {
		return err
	}

	rb.Password = hashedPassword

	tx, err := u.userRepo.ExecTx(ctx)
	if err != nil {
		return fmt.Errorf("SignupTx: %v", err)
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	err = u.userRepo.InsertTx(ctx, tx, rb)
	if err != nil {
		return fmt.Errorf("SignupTx: %v", err)
	}

	token, err := u.tokenService.GenerateAccessToken(rb.Email)
	if err != nil {
		return err
	}

	fmt.Println("user inserted/.....................")

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
	fmt.Println(verificationLink)

	err = u.emailService.SendEmailTx(rb.Email, "Verfify Email Address", emailBody)
	if err != nil {
		return err
	}
	return tx.Commit()
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
