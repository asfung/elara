package entities

import (
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

// CARDS
type Card struct {
	ID                 string    `json:"id"`
	UserID             string    `json:"user_id"`
	CardNumberHash     string    `json:"card_number_hash"`
	CardType           string    `json:"card_type"`
	ExpiryDate         string    `json:"expiry_date"`
	Status             string    `json:"status"`
	TokenizedReference string    `json:"tokenized_reference"`
	CreatedAt          time.Time `json:"created_at"`
}

func NewCard(userID, cardHash, cardType, expiry string) (*Card, error) {
	id, err := gonanoid.New()
	if err != nil {
		return nil, err
	}
	return &Card{
		ID:                 id,
		UserID:             userID,
		CardNumberHash:     cardHash,
		CardType:           cardType,
		ExpiryDate:         expiry,
		Status:             "active",
		TokenizedReference: "",
		CreatedAt:          time.Now(),
	}, nil
}
