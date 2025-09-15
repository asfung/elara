package services

import (
	"github.com/asfung/elara/internal/entities"
	"github.com/asfung/elara/internal/models"
)

type OTPService interface {
	CreateOTP(req models.AddOTPRequest) (entities.OTP, error)
	UpdateOTP(req models.UpdateOTPRequest) (entities.OTP, error)
	GetOTPById(id string) (entities.OTP, error)
	DeleteOTP(id string) error
	GetOTPByCode(code string) (entities.OTP, error)
	VerifyOTP(userID string, otp string) (bool, error)
}
