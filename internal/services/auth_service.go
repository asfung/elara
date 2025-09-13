package services

import (
	"github.com/asfung/elara/internal/entities"
	"github.com/asfung/elara/internal/models"
	"github.com/markbates/goth"
)

type AuthService interface {
	Login(req models.LoginRequest) (string, string, error)
	Register(req models.AddUserRequest) (entities.User, error)
	RefreshToken(req models.RefreshTokenRequest) (models.AuthResponse, error)
	Logout(token string) error
	Verify(token string) (*entities.User, error)
	OAuthLoginFromGothUser(goth.User) (accessToken string, refreshToken string, err error)
	GetUserByEmail(email string) (entities.User, error)
}
