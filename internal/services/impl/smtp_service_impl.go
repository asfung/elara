package impl

import (
	"bytes"
	"html/template"
	"net/smtp"
	"os"

	"github.com/asfung/elara/internal/services"
)

var (
	SMTP_FROM     = os.Getenv("SMTP_FROM")
	SMTP_PASSWORD = os.Getenv("SMTP_PASSWORD")
	SMTP_HOST     = os.Getenv("SMTP_HOST")
	SMTP_PORT     = os.Getenv("SMTP_PORT")

	// SMTP_FROM="paungcuy15@gmail.com"
	// SMTP_PASSWORD=vtrsrtkahhcxsyck
	// SMTP_HOST="smtp.gmail.com"
	// SMTP_PORT="587"
)

type smtpServiceImpl struct {
	Host         string
	Port         string
	Username     string
	Password     string
	TemplateHTML string
}

func NewSmtpServiceImpl(host, port, username, password, templateHTML string) services.SmtpService {
	return &smtpServiceImpl{
		Host:         host,
		Port:         port,
		Username:     username,
		Password:     password,
		TemplateHTML: templateHTML,
	}
}

func (s *smtpServiceImpl) SendEmail(to string, subject string, data interface{}) error {
	tmpl, err := template.ParseFiles(s.TemplateHTML)
	if err != nil {
		return err
	}

	var bodyBuffer bytes.Buffer
	if err := tmpl.Execute(&bodyBuffer, data); err != nil {
		return err
	}

	msg := []byte(
		"To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/html; charset=\"utf-8\"\r\n" +
			"\r\n" +
			bodyBuffer.String(),
	)

	auth := smtp.PlainAuth("", s.Username, s.Password, s.Host)

	addr := s.Host + ":" + s.Port
	if err := smtp.SendMail(addr, auth, SMTP_FROM, []string{to}, msg); err != nil {
		return err
	}

	return nil
}
