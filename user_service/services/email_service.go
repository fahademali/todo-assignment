package services

import (
	"fmt"
	"user_service/config"

	"gopkg.in/gomail.v2"
)

type IEmailService interface {
	SendEmail(token string, receiverEmail string) error
}

type EmailService struct {
}

func NewEmailService() IEmailService {
	return &EmailService{}
}

func (es *EmailService) SendEmail(token string, receiverEmail string) error {
	verificationLink := fmt.Sprintf("http://localhost:8080/verify-user/%s", token)
	emailBody := `
	<html>
		<body>
			<p>Hello!</p>
			<p>Click the following link to verify your email address:</p>
			<p><a href="` + verificationLink + `">Verify Email Address</a></p>
		</body>
	</html>
	`

	fmt.Println(config.AppConfig.SENDER_EMAIL,
		config.AppConfig.SMTP_SERVER,
		config.AppConfig.SMTP_PORT,
		config.AppConfig.SENDER_EMAIL,
		config.AppConfig.SENDER_APP_PASS)

	m := gomail.NewMessage()
	m.SetHeader("From", config.AppConfig.SENDER_EMAIL)
	m.SetHeader("To", receiverEmail)
	m.SetHeader("Subject", "Verify your email address")
	m.SetBody("text/html", emailBody)

	d := gomail.NewDialer(config.AppConfig.SMTP_SERVER, config.AppConfig.SMTP_PORT, config.AppConfig.SENDER_EMAIL, config.AppConfig.SENDER_APP_PASS)

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
