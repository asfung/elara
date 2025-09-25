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
	RefreshToken string `json:"refresh_token" validate:"required"`
}

// Response DTOs
type UserResponse struct {
	UserId       uuid.UUID `json:"user_id"`
	AvatarUrl    *string   `json:"avatar_url"`
	Username     string    `json:"username"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	FirstName    *string   `json:"first_name"`
	LastName     *string   `json:"last_name"`
	Subscription *string   `json:"subscription"`
}

// Entity -> Response
func ToUserResponse(user entities.User) UserResponse {
	return UserResponse{
		UserId:       user.UserID,
		AvatarUrl:    user.AvatarURL,
		Username:     user.Username,
		Name:         user.Name,
		Email:        user.Email,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Subscription: user.Subscription,
	}
}
