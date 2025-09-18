package service

import (
	"net/smtp"
)

type SMTPService interface {
	SendMail(to string, subject, message string) error
}

type SMTPClient struct {
	Auth    smtp.Auth
	Address string
	Email   string
	Name    string
}

type smtpService struct {
	client SMTPClient
}

func NewSMTPService(
	client SMTPClient,
) SMTPService {
	return &smtpService{
		client: client,
	}
}

func (s *smtpService) SendMail(to string, subject, message string) error {
	body := "From: " + s.client.Name + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n" +
		message

	err := smtp.SendMail(s.client.Address, s.client.Auth, s.client.Email, []string{to}, []byte(body))
	if err != nil {
		return err
	}

	return nil
}
