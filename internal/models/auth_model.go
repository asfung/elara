package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Request DTOs
type LoginRequest struct {
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
	AccessToken          string        `json:"access_token"`
	RefreshToken         string        `json:"refresh_token,omitempty"`
	AccessTokenFormatted string        `json:"access_token_formatted"`
	ExpiresAt            time.Duration `json:"expires_at"`
}
type Claims struct {
	ID           uint32 `json:"id"`
	UserID       string `json:"user_id"`
	Email        string `json:"email"`
	TokenVersion int    `json:"token_version"`
	jwt.RegisteredClaims
}

// Entity -> Response
