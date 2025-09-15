package entities

type Role struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	Name        string  `gorm:"uniqueIndex;size:50" json:"name"` // "user", "admin", "super"
	Description *string `json:"description"`
}
