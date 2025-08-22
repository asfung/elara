package services

import (
	"github.com/asfung/elara/internal/entities"
	"github.com/asfung/elara/internal/models"
)

type AuthService interface {
	Login(req models.LoginrRequest) (string, error)
	Register(req models.AddUserRequest) (entities.User, error)
	RefreshToken(req models.RefreshTokenRequest) (models.AuthResponse, error)
	Logout() error
}
