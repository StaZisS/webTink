package service

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"io/ioutil"
	"net/smtp"
	"os"
	listing "web"
	"web/pkg/repository"
)

type EmailService struct {
	repo repository.Email
}

func NewEmailService(repo repository.Email) *EmailService {
	return &EmailService{repo: repo}
}

func (s *EmailService) SendEmail(email listing.Email) error {
	validate := validator.New()
	if err := validate.Struct(email); err != nil {
		return fmt.Errorf("validate: %w", err)
	}
	auth := smtp.PlainAuth("", os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASSWORD"), os.Getenv("SMTP_HOST"))
	sample, err := ioutil.ReadFile("../configs/email_html.txt")
	if err != nil {
		return fmt.Errorf("read file: %w", err)
	}

	subject := "Subject: У вас сообщение от TEAM XD!\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := fmt.Sprintf(string(sample), email.Name, email.Message)
	msg := []byte(subject + mime + body)

	err = smtp.SendMail(os.Getenv("SMTP_HOST")+":"+os.Getenv("SMTP_PORT"), auth, os.Getenv("SMTP_USER"), []string{email.EmailValue}, msg)
	if err != nil {
		return fmt.Errorf("send mail: %w", err)
	}
	return nil
}
