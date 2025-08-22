package models

import (
	"github.com/asfung/elara/internal/entities"
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
	Id        uint32
	Username  string
	Email     string
	FirstName *string
	LastName  *string
}

// Entity -> Response
func ToUserResponse(user entities.User) UserResponse {
	return UserResponse{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}
