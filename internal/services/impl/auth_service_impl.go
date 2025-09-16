package impl

import (
	"errors"
	"time"

	"github.com/asfung/elara/internal/entities"
	"github.com/asfung/elara/internal/models"
	"github.com/asfung/elara/internal/repositories"
	"github.com/asfung/elara/internal/services"
	"github.com/asfung/elara/utils"
	"github.com/charmbracelet/log"
	"github.com/markbates/goth"
	"gorm.io/gorm"
)

type authServiceImpl struct {
	authRepo    repositories.AuthRepository
	userRepo    repositories.UserRepository
	userService services.UserService
	otpService  services.OTPService
	smtpService services.SmtpService
}

func NewAuthServiceImpl(
	authRepo repositories.AuthRepository,
	userRepo repositories.UserRepository,
	userSvc services.UserService,
	otpService services.OTPService,
	smtpService services.SmtpService,
) services.AuthService {
	return &authServiceImpl{
		authRepo:    authRepo,
		userService: userSvc,
		userRepo:    userRepo,
		otpService:  otpService,
		smtpService: smtpService,
	}
}

func (a *authServiceImpl) Login(req models.LoginRequest) (string, string, error) {
	email := req.Email
	password := req.Password

	user, err := a.userRepo.FindByEmail(email)
	if err != nil {
		return "error", "error", err
	}

	if user.Password == nil || !utils.VerifyPassword(password, *user.Password) {
		return "", "", errors.New("invalid email or password")
	}

	accessToken, err := utils.CreateToken(&user, 24*time.Hour*7)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := utils.CreateToken(&user, 24*time.Hour*30)
	if err != nil {
		return "", "", err
	}

	user.AccessToken = &accessToken
	user.RefreshToken = &refreshToken
	if _, err := a.userRepo.Update(user); err != nil {
		log.Error(err)
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (a *authServiceImpl) Register(req models.AddUserRequest) (entities.User, error) {
	user, err := a.userService.CreateUser(req)
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (a *authServiceImpl) Logout(token string) error {
	claims, err := utils.VerifyToken(token)
	if err != nil {
		return errors.New("invalid or expired token")
	}

	user, err := a.userRepo.FindByUserId(claims.UserID)
	if err != nil {
		return err
	}

	user.TokenVersion++
	user.AccessToken = nil
	user.RefreshToken = nil
	_, err = a.userRepo.Update(user)
	return err

}

func (a *authServiceImpl) RefreshToken(req models.RefreshTokenRequest) (models.AuthResponse, error) {
	claims, err := utils.VerifyToken(req.RefreshToken)
	if err != nil {
		return models.AuthResponse{}, errors.New("invalid or expired refresh token")
	}

	user, err := a.userRepo.FindByUserId(claims.UserID)
	if err != nil {
		return models.AuthResponse{}, err
	}

	if user.RefreshToken == nil || *user.RefreshToken != req.RefreshToken {
		return models.AuthResponse{}, errors.New("refresh token mismatch")
	}
	if claims.TokenVersion != user.TokenVersion {
		return models.AuthResponse{}, errors.New("refresh token revoked")
	}

	user.TokenVersion++

	accessToken, err := utils.CreateToken(&user, 24*time.Hour*7)
	if err != nil {
		return models.AuthResponse{}, err
	}

	newRefreshToken, err := utils.CreateToken(&user, 24*time.Hour*30)
	if err != nil {
		return models.AuthResponse{}, err
	}

	user.AccessToken = &accessToken
	user.RefreshToken = &newRefreshToken
	if _, err := a.userRepo.Update(user); err != nil {
		return models.AuthResponse{}, err
	}

	return models.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}, nil
}

func (a *authServiceImpl) Verify(token string) (*entities.User, error) {
	claims, err := utils.VerifyToken(token)
	if err != nil {
		return nil, err
	}

	user, err := a.userRepo.FindByUserId(claims.UserID)
	if err != nil {
		return nil, err
	}

	if claims.TokenVersion != user.TokenVersion {
		return nil, errors.New("token revoked")
	}

	return &user, nil
}

func (a *authServiceImpl) OAuthLoginFromGothUser(gUser goth.User) (string, string, error) {

	u, err := a.userRepo.FindByProvider(gUser.Provider, gUser.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newUser := entities.User{
				Provider:       gUser.Provider,
				ProviderUserID: gUser.UserID, // its actually UserID from provider
				Email:          gUser.Email,
				Name:           gUser.Name,
				FirstName:      &gUser.FirstName,
				LastName:       &gUser.LastName,
				Username:       utils.GenerateUsername(gUser.NickName),
				AvatarURL:      &gUser.AvatarURL,
				Location:       &gUser.Location,
			}
			created, err := a.userRepo.Create(newUser)
			if err != nil {
				return "", "", err
			}
			u = created
		} else {
			return "", "", err
		}
	}

	accessToken, refreshToken, err := a.CreateTokensForUser(u)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (a *authServiceImpl) CreateTokensForUser(u entities.User) (string, string, error) {
	user, err := a.userRepo.FindByEmail(u.Email)
	if err != nil {
		return "", "", err
	}

	accessToken, err := utils.CreateToken(&user, 15*time.Minute)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := utils.CreateToken(&user, 24*time.Hour*30)
	if err != nil {
		return "", "", err
	}

	user.AccessToken = &accessToken
	user.RefreshToken = &refreshToken
	if _, err := a.userRepo.Update(user); err != nil {
		log.Error(err)
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (a *authServiceImpl) GetUserByEmail(email string) (entities.User, error) {
	user, err := a.userRepo.FindByEmail(email)
	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}

func (a *authServiceImpl) CreateAccountWithPassword(email, password string) (entities.User, error) {
	hashed, err := utils.HashPassword(password)
	if err != nil {
		return entities.User{}, err
	}

	emailPrefix, err := utils.GetEmailPrefix(email)
	if err != nil {
		return entities.User{}, err
	}

	nameAndUsernameGen := utils.GenerateUsername(emailPrefix)
	user := entities.User{
		Email:    email,
		Password: &hashed,
		Name:     nameAndUsernameGen,
		Username: nameAndUsernameGen,
		// Description: utils.StringPtr("pending_verification"),
	}

	accountCreated, err := a.userRepo.Create(user)
	if err != nil {
		return entities.User{}, err
	}

	otp, err := utils.GenerateOTP(6)
	if err != nil {
		return entities.User{}, err
	}

	otpCreated, err := a.otpService.CreateOTP(models.AddOTPRequest{UserID: accountCreated.UserID.String(), Code: otp, ExpiresAt: 15 * time.Minute})
	if err != nil {
		return entities.User{}, err
	}

	toEmail := accountCreated.Email
	data := map[string]interface{}{
		"Name": accountCreated.Name,
		"OTP":  otpCreated.Code,
	}
	err = a.smtpService.SendEmail(toEmail, "Elara Code Verification", data)
	if err != nil {
		log.Error("AuthServiceImpl.CreateAccountWithPassword Error", err)
	} else {
		log.Info("send email successfully")
	}

	return accountCreated, nil
}

func (a *authServiceImpl) VerifyPassword(email, password string) (bool, error) {
	user, err := a.userRepo.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.Password == nil || !utils.VerifyPassword(password, *user.Password) {
		return false, errors.New("invalid email or password")
	}

	otp, err := utils.GenerateOTP(6)
	if err != nil {
		return false, err
	}

	otpCreated, err := a.otpService.CreateOTP(models.AddOTPRequest{UserID: user.UserID.String(), Code: otp, ExpiresAt: 15 * time.Minute})
	if err != nil {
		return false, err
	}

	toEmail := user.Email
	data := map[string]interface{}{
		"Name": user.Name,
		"OTP":  otpCreated.Code,
	}
	err = a.smtpService.SendEmail(toEmail, "Elara Code Verification", data)
	if err != nil {
		log.Error("AuthServiceImpl.CreateAccountWithPassword Error", err)
	} else {
		log.Info("send email successfully")
	}

	return true, nil
}
