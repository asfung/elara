package entities

import (
	"crypto/rand"
	"encoding/binary"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id          uint32  `json:"id"`
	Provider    string  `json:"provider"`
	Email       string  `json:"email" gorm:"uniqueIndex:idx_users_email" label:"email"`
	Name        string  `json:"name"`
	FirstName   *string `json:"firstName"`
	LastName    *string `json:"lastName"`
	Username    string  `json:"username" gorm:"uniqueIndex:idx_users_username" label:"username"`
	Password    *string `json:"password"`
	Description *string `json:"description"`
	// UserID            uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()" json:"userId"`
	UserID            uuid.UUID `gorm:"type:uuid" json:"userId"`
	AvatarURL         *string   `json:"avatarUrl"`
	Location          *string   `json:"location"`
	AccessToken       *string   `json:"accessToken"`
	AccessTokenSecret *string   `json:"accessTokenSecret"`
	RefreshToken      *string   `json:"refreshToken"`
	TokenVersion      int       `gorm:"default:1"`
	IDToken           *string   `json:"idToken"`
	TimeStamp
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.UserID == uuid.Nil {
		// u.UserID, err = uuid.NewV7()
		u.UserID, err = NewV8(1)
	}
	return
}

func NewV8(shardID uint16) (uuid.UUID, error) {
	var u uuid.UUID
	now := uint32(time.Now().Unix())

	// fill first 4 bytes with timestamp
	binary.BigEndian.PutUint32(u[0:4], now)
	// next 2 bytes = shard id
	binary.BigEndian.PutUint16(u[4:6], shardID)
	// remaining 10 bytes random
	if _, err := rand.Read(u[6:16]); err != nil {
		return uuid.Nil, err
	}
	// set version = 8
	u[6] = (u[6] & 0x0f) | (8 << 4)
	// set rfc 4122 variant
	u[8] = (u[8] & 0x3f) | 0x80

	return u, nil
}

/*
regionID -> US-East, Europe, Asia-Pacific
shardID -> 1-1M userId = 1 (shardID), 1M+1-2M = 2 (shardID) [range-based]
*/
func NewV8WithRegion(regionID, shardID uint16) (uuid.UUID, error) {
	var u uuid.UUID
	now := uint32(time.Now().Unix())

	// fill first 4 bytes with timestamp
	binary.BigEndian.PutUint32(u[0:4], now)
	// next 2 bytes = region id
	binary.BigEndian.PutUint16(u[4:6], regionID)
	// next 2 bytes = shard id
	binary.BigEndian.PutUint16(u[6:8], shardID)
	// remaining 8 bytes random
	if _, err := rand.Read(u[8:16]); err != nil {
		return uuid.Nil, err
	}

	// set version = 8
	u[6] = (u[6] & 0x0f) | (8 << 4)
	// set rfc 4122 variant
	u[8] = (u[8] & 0x3f) | 0x80

	return u, nil
}
