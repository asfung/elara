package repositories

import (
	"time"

	"github.com/asfung/elara/database"
	"github.com/asfung/elara/internal/entities"
)

type OTPPostgresRepository struct {
	*BaseRepository[entities.OTP]
}

func NewOTPPotgresRepository(db database.Database) OTPRepository {
	return &OTPPostgresRepository{
		BaseRepository: NewBaseRepository[entities.OTP](db),
	}
}

func (o *OTPPostgresRepository) FindByCode(code string) (entities.OTP, error) {
	var otp entities.OTP
	if err := o.db.GetDb().Where("code = ?", code).First(&otp).Error; err != nil {
		return entities.OTP{}, err
	}
	return otp, nil
}

func (o *OTPPostgresRepository) VerifyOTP(userID string, code string) (bool, error) {
	var otp entities.OTP
	if err := o.db.GetDb().Where("user_id = ? AND code = ?", userID, code).First(&otp).Error; err != nil {
		return false, err
	}
	if time.Now().After(otp.ExpiresAt) {
		return false, nil
	}
	return true, nil
}

func (o *OTPPostgresRepository) FindByUserId(userID string) (entities.OTP, error) {
	var otp entities.OTP
	if err := o.db.GetDb().Where("user_id = ?", userID).First(&otp).Error; err != nil {
		return entities.OTP{}, err
	}
	return otp, nil
}
