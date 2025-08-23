package entities

import (
	"github.com/google/uuid"
)

type User struct {
	Id                uint32    `json:"id"`
	Provider          string    `json:"provider"`
	Email             string    `json:"email" gorm:"uniqueIndex:idx_users_email" label:"email"`
	Name              string    `json:"name"`
	FirstName         *string   `json:"firstName"`
	LastName          *string   `json:"lastName"`
	Username          string    `json:"username" gorm:"uniqueIndex:idx_users_username" label:"username"`
	Password          *string   `json:"password"`
	Description       *string   `json:"description"`
	UserID            uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()" json:"userId"`
	AvatarURL         *string   `json:"avatarUrl"`
	Location          *string   `json:"location"`
	AccessToken       *string   `json:"accessToken"`
	AccessTokenSecret *string   `json:"accessTokenSecret"`
	RefreshToken      *string   `json:"refreshToken"`
	IDToken           *string   `json:"idToken"`
	TimeStamp
}
