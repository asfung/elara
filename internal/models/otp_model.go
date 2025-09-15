package models

import (
	"time"

	"github.com/asfung/elara/internal/entities"
)

// Request DTOs
type AddOTPRequest struct {
	UserID    string        `json:"user_id" validate:"required"`
	Code      string        `json:"code" validate:"required"`
	ExpiresAt time.Duration `json:"expires_at" validate:"required"`
}
type UpdateOTPRequest struct {
	ID        uint          `json:"id"`
	UserID    string        `json:"user_id" validate:"required"`
	Code      string        `json:"code" validate:"required"`
	ExpiresAt time.Duration `json:"expires_at" validate:"required"`
}

// Response DTOs
type OTPResponse struct {
	ID        uint          `json:"id"`
	UserID    string        `json:"user_id"`
	Code      string        `json:"code" validate:"required"`
	ExpiresAt time.Duration `json:"expires_at" validate:"required"`
}

// Entity -> Response
func ToOTPdResponse(otp entities.OTP) OTPResponse {
	return OTPResponse{
		ID:        otp.ID,
		UserID:    otp.UserID,
		Code:      otp.Code,
		ExpiresAt: time.Until(otp.ExpiresAt),
	}
}
