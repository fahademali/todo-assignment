package services

import (
	"user_service/log"
	"user_service/models"

	"gopkg.in/gomail.v2"
)

type IEmailService interface {
	SendEmail(receiverEmail string, subject string, body string) error
	SendEmailToAll(recepientDetails []models.RecepietDetails) error
}

type EmailService struct {
	senderEmail      string
	senderAppPass    string
	senderSmtpServer string
	senderSmtpPort   int
}

func NewEmailService(senderEmail string, senderAppPass string, senderSmtpServer string, senderSmtpPort int) IEmailService {
	return &EmailService{senderEmail: senderEmail, senderAppPass: senderAppPass, senderSmtpServer: senderSmtpServer, senderSmtpPort: senderSmtpPort}
}

func (es *EmailService) SendEmail(receiverEmail string, subject string, body string) error {
	log := log.GetLog()
	log.Info("SENDNIG ENMAIL......")

	emailMsg := gomail.NewMessage()
	emailMsg.SetHeader("From", es.senderEmail)
	emailMsg.SetHeader("To", receiverEmail)
	emailMsg.SetHeader("Subject", subject)
	emailMsg.SetBody("text/html", body)

	dialer := gomail.NewDialer(es.senderSmtpServer, es.senderSmtpPort, es.senderEmail, es.senderAppPass)

	return dialer.DialAndSend(emailMsg)
}

func (es *EmailService) SendEmailToAll(recepientDetails []models.RecepietDetails) error {
	emailMsg := gomail.NewMessage()
	dialer := gomail.NewDialer(es.senderSmtpServer, es.senderSmtpPort, es.senderEmail, es.senderAppPass)
	emailMsg.SetHeader("From", es.senderEmail)

	for _, recepient := range recepientDetails {
		emailMsg.SetHeader("To", recepient.UserEmail)
		emailMsg.SetHeader("Subject", recepient.Subject)
		emailMsg.SetBody("text/html", recepient.Body)
		if err := dialer.DialAndSend(emailMsg); err != nil {
			log.GetLog().Error("err")
			log.GetLog().Error(err)
			return err
		}
	}

	return nil
}
