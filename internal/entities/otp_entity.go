package entities

import "time"

type OTP struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    string `gorm:"index,uniqueIndex:idx_user_id"`
	Code      string `gorm:"uniqueIndex:idx_user_code"`
	ExpiresAt time.Time
}

func NewOtp(userID string, code string, ttl time.Duration) (*OTP, error) {
	return &OTP{
		UserID:    userID,
		Code:      code,
		ExpiresAt: time.Now().Add(ttl),
	}, nil
}
