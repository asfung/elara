package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Request DTOs
type LoginrRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}
type RegisterRequest struct {
	Username string `validate:"required"`
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

// Response DTOs
type AuthResponse struct {
	AccessToken  string        `json:"access_token"`
	RefreshToken string        `json:"refresh_token"`
	ExpiresAt    time.Duration `json:"expires_at"`
}
type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// Entity -> Response
