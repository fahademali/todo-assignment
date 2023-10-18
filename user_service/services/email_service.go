package services

import (
	"fmt"

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
			<p>Hello <b>Bob</b> and <i>Cora</i>!</p>
			<p>Click the following link to verify your email address:</p>
			<p><a href="` + verificationLink + `">Verify Email Address</a></p>
		</body>
	</html>
	`

	m := gomail.NewMessage()
	m.SetHeader("From", "valeedtest@gmail.com")
	m.SetHeader("To", receiverEmail)
	m.SetHeader("Subject", "Verify your email address")
	m.SetBody("text/html", emailBody)

	d := gomail.NewDialer("smtp.gmail.com", 587, "valeedtest@gmail.com", "anhf fraz llzc karg")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
