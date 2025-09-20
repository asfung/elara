package utils

import (
	"crypto/sha256"
	"errors"
	"time"

	"github.com/asfung/elara/config"
	"github.com/asfung/elara/internal/entities"
	"github.com/asfung/elara/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwe"
)

// type User = entities.User
// type Claims = entities.Claims

// var jwtSecret = []byte(config.GetConfig().Jwt.Secret)
var jwtSecret = getKey(config.GetConfig().Jwt.Secret)

func getKey(secret string) []byte {
	h := sha256.Sum256([]byte(secret))
	return h[:] // 32-byte key
}

func CreateToken(user *entities.User, duration time.Duration) (string, error) {
	claims := models.Claims{
		ID:           user.Id,
		UserID:       user.UserID.String(),
		Email:        user.Email,
		TokenVersion: user.TokenVersion,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	encryptSignToken, err := EncryptedJWToken(signedToken)
	if err != nil {
		return "", err
	}
	return encryptSignToken, nil
}

func EncryptedJWToken(signedToken string) (string, error) {
	encrypted, err := jwe.Encrypt([]byte(signedToken), jwe.WithKey(jwa.DIRECT, jwtSecret))
	if err != nil {
		return "", err
	}
	return string(encrypted), nil
}

func VerifyToken(token string) (*models.Claims, error) {
	tokenObj, err := jwt.ParseWithClaims(token, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err == nil && tokenObj.Valid {
		if claims, ok := tokenObj.Claims.(*models.Claims); ok {
			return claims, nil
		}
	}

	// an old claims
	// claims, ok := tokenObj.Claims.(*models.Claims)
	// if !ok || !tokenObj.Valid {
	// 	return nil, errors.New("invalid token")
	// }
	// return claims, nil

	decrypted, err := jwe.Decrypt([]byte(token), jwe.WithKey(jwa.DIRECT, jwtSecret))
	if err != nil {
		return nil, errors.New("failed to decrypt the token ")
	}

	tokenObj, err = jwt.ParseWithClaims(string(decrypted), &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !tokenObj.Valid {
		return nil, errors.New("invalid token after decryption")
	}

	claims, ok := tokenObj.Claims.(*models.Claims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
