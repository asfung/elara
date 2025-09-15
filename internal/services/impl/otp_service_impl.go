package impl

import (
	"time"

	"github.com/asfung/elara/internal/entities"
	"github.com/asfung/elara/internal/models"
	"github.com/asfung/elara/internal/repositories"
	"github.com/asfung/elara/internal/services"
	"github.com/asfung/elara/utils"
)

type otpServiceImpl struct {
	otpRepo  repositories.OTPRepository
	userRepo repositories.UserRepository
}

func NewOTPServiceImpl(otpRepo repositories.OTPRepository, userRepo repositories.UserRepository) services.OTPService {
	return &otpServiceImpl{
		otpRepo:  otpRepo,
		userRepo: userRepo,
	}
}

func (o *otpServiceImpl) CreateOTP(req models.AddOTPRequest) (entities.OTP, error) {
	otpExist, err := o.otpRepo.FindByUserId(req.UserID)
	if err != nil {
		return entities.OTP{}, err
	}

	if otpExist.Code != "" {
		// if the otp code exist and not expired it just add the extras expires
		otpExist.ExpiresAt = time.Now().Add(15 * time.Minute)
		if time.Now().After(otpExist.ExpiresAt) {
			// if otp code exist and expired it will create the new code and new expires
			otpExist.Code, _ = utils.GenerateOTP(6)
			otpExist.ExpiresAt = time.Now().Add(15 * time.Minute)
		}
		updated, err := o.otpRepo.Update(otpExist)
		if err != nil {
			return entities.OTP{}, err
		}
		return updated, nil
	}

	otp, err := entities.NewOtp(req.UserID, req.Code, req.ExpiresAt)
	if err != nil {
		return entities.OTP{}, err
	}

	created, err := o.otpRepo.Create(*otp)
	if err != nil {
		return entities.OTP{}, err
	}
	return created, nil
}

func (o *otpServiceImpl) UpdateOTP(req models.UpdateOTPRequest) (entities.OTP, error) {

	otp, err := o.otpRepo.FindById(req.ID)
	if err != nil {
		return entities.OTP{}, err
	}

	if req.UserID != "" {
		otp.UserID = req.UserID
	}
	if req.Code != "" {
		otp.Code = req.Code
	}
	if req.ExpiresAt != 0 {
		otp.ExpiresAt = time.Now().Add(req.ExpiresAt)
	}

	updated, err := o.otpRepo.Update(*otp)
	if err != nil {
		return entities.OTP{}, err
	}

	return updated, nil
}

func (o *otpServiceImpl) GetOTPById(id string) (entities.OTP, error) {
	otp, err := o.otpRepo.FindById(id)
	if err != nil {
		return entities.OTP{}, err
	}
	return *otp, nil
}

func (o *otpServiceImpl) DeleteOTP(id string) error {
	return o.otpRepo.Delete(id)
}

func (o *otpServiceImpl) GetOTPByCode(code string) (entities.OTP, error) {
	otp, err := o.otpRepo.FindByCode(code)
	if err != nil {
		return entities.OTP{}, err
	}
	return otp, nil
}

func (o *otpServiceImpl) VerifyOTP(userID string, otp string) (bool, error) {
	isValid, err := o.otpRepo.VerifyOTP(userID, otp)
	if err != nil {
		return false, err
	}

	otpEntity, err := o.otpRepo.FindByUserId(userID)
	if err != nil {
		return false, err
	}

	user, err := o.userRepo.FindByUserId(otpEntity.UserID)
	if err != nil {
		return false, err
	}

	if err := o.otpRepo.Delete(otpEntity.ID); err != nil {
		return false, err
	}
	trueVal := true
	if user.EmailVerified == nil || !*user.EmailVerified {
		user.EmailVerified = &trueVal
	}
	if _, err := o.userRepo.Update(user); err != nil {
		return false, err
	}

	return isValid, nil
}
