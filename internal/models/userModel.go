package models

import (
	"github.com/asfung/elara/internal/entities"
	"github.com/google/uuid"
)

// Request DTOs
type AddUserRequest struct {
	Username string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}
type UpdateUserRequest struct {
	Id        uint32 `validate:"required"`
	Username  string `validate:"omitempty"`
	Email     string `validate:"omitempty,email"`
	FirstName *string
	LastName  *string
}
type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

// Response DTOs
type UserResponse struct {
	UserId    uuid.UUID
	AvatarUrl *string
	Username  string
	Email     string
	FirstName *string
	LastName  *string
}

// Entity -> Response
func ToUserResponse(user entities.User) UserResponse {
	return UserResponse{
		UserId: user.UserID,
		AvatarUrl: user.AvatarURL,
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}
