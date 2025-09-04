package models

import (
	"time"

	"github.com/asfung/elara/internal/entities"
)

// ========= CARD =========
// Request DTOs
type AddCardRequest struct {
	UserID         string `json:"user_id" validate:"required"`
	CardNumberHash string `json:"card_number_hash" validate:"required"`
	CardType       string `json:"card_type" validate:"required"`
	ExpiryDate     string `json:"expiry_date" validate:"required"`
}
type UpdateCardRequest struct {
	ID                 string `json:"id" validate:"omitempty"`
	CardNumberHash     string `json:"card_number_hash"`
	CardType           string `json:"card_type"`
	ExpiryDate         string `json:"expiry_date"`
	Status             string `json:"status"`
	TokenizedReference string `json:"tokenized_reference"`
}

// Response DTOs
type CardResponse struct {
	ID                 string    `json:"id"`
	UserID             string    `json:"user_id"`
	CardNumberHash     string    `json:"card_number_hash"`
	CardType           string    `json:"card_type"`
	ExpiryDate         string    `json:"expiry_date"`
	Status             string    `json:"status"`
	TokenizedReference string    `json:"tokenized_reference"`
	CreatedAt          time.Time `json:"created_at"`
}

// Entity -> Response
func ToCardResponse(bank entities.Card) CardResponse {
	return CardResponse{
		ID:                 bank.ID,
		UserID:             bank.UserID,
		CardNumberHash:     bank.CardNumberHash,
		CardType:           bank.CardType,
		ExpiryDate:         bank.ExpiryDate,
		Status:             bank.Status,
		TokenizedReference: bank.TokenizedReference,
		CreatedAt:          bank.CreatedAt,
	}
}
