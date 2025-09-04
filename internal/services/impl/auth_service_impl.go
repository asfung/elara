package impl

import (
	"errors"
	"fmt"
	"time"

	"github.com/asfung/elara/internal/entities"
	"github.com/asfung/elara/internal/models"
	"github.com/asfung/elara/internal/repositories"
	"github.com/asfung/elara/internal/services"
	"github.com/asfung/elara/utils"
	"github.com/charmbracelet/log"
)

type authServiceImpl struct {
	authRepo    repositories.AuthRepository
	userRepo    repositories.UserRepository
	userService services.UserService
}

func NewAuthServiceImpl(
	authRepo repositories.AuthRepository,
	userRepo repositories.UserRepository,
	userSvc services.UserService,
) services.AuthService {
	return &authServiceImpl{
		authRepo:    authRepo,
		userService: userSvc,
		userRepo:    userRepo,
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
	fmt.Println(req.RefreshToken)
	claims, err := utils.VerifyToken(req.RefreshToken)
	if err != nil {
		return models.AuthResponse{}, errors.New("invalid or expired refresh token")
	}
	fmt.Println(claims)

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
