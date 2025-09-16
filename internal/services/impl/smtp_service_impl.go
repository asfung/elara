package impl

import (
	"bytes"
	"html/template"
	"net/smtp"

	"github.com/asfung/elara/internal/services"
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
		"From: " + s.Username + "\r\n" +
			"To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/html; charset=\"utf-8\"\r\n" +
			"\r\n" +
			bodyBuffer.String(),
	)

	auth := smtp.PlainAuth("", s.Username, s.Password, s.Host)
	addr := s.Host + ":" + s.Port

	// log.Infof("SMTP Config -> host=%s port=%s user=%s", s.Host, s.Port, s.Username)
	// log.Infof("sending email to %s via %s", to, addr)

	if err := smtp.SendMail(addr, auth, s.Username, []string{to}, msg); err != nil {
		return err
	}
	return nil
}
