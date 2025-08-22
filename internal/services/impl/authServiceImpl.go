package impl

import (
	"errors"

	"github.com/asfung/elara/internal/entities"
	"github.com/asfung/elara/internal/models"
	"github.com/asfung/elara/internal/repositories"
	"github.com/asfung/elara/internal/services"
	"github.com/asfung/elara/utils"
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

func (a *authServiceImpl) Login(req models.LoginrRequest) (string, error) {
	email := req.Email
	password := req.Password

	user, err := a.userRepo.FindByEmail(email)
	if err != nil {
		return "error", err
	}

	if user.Password == nil || !utils.VerifyPassword(password, *user.Password) {
		return "", errors.New("invalid email or password")
	}

	token, err := utils.CreateToken(&user)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *authServiceImpl) Register(req models.AddUserRequest) (entities.User, error) {
	user, err := a.userService.CreateUser(req)
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (a *authServiceImpl) Logout() error {
	panic("unimplemented")
}

func (a *authServiceImpl) RefreshToken(req models.RefreshTokenRequest) (models.AuthResponse, error) {
	panic("unimplemented")
}
