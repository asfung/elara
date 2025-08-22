package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Request DTOs

// Response DTOs
type AuthResponse struct {
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
}
type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// Entity -> Response
