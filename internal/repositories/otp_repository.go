package repositories

import "github.com/asfung/elara/internal/entities"

type OTPRepository interface {
	Repository[entities.OTP]
	FindByCode(code string) (entities.OTP, error)
	FindByUserId(userID string) (entities.OTP, error)
	VerifyOTP(userID string, code string) (bool, error)
}
