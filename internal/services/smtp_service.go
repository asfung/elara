package services

type SmtpService interface {
	SendEmail(to string, subject string, data interface{}) error
}
