package utils

import (
	"errors"
	"time"

	"github.com/asfung/elara/internal/entities"
	"github.com/asfung/elara/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

// type User = entities.User
// type Claims = entities.Claims

var jwtSecret = []byte("53cr3t_k3y")

func CreateToken(user *entities.User) (string, error) {
	claims := models.Claims{
		UserID: user.UserID.String(),
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenStr string) (*models.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &models.Claims{}, func(token *jwt.Token) (any, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*models.Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}

func RefreshToken(refreshToken string, user *entities.User) (string, error) {
	// check if refreshToken matches user's stored refresh token from table User

	if refreshToken != user.RefreshToken {
		return "", errors.New("invalid refresh token")
	}

	return CreateToken(user)
}
