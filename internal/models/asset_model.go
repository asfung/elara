package models

import "github.com/asfung/elara/internal/entities"

// Request DTOs
type AddAssetRequest struct {
	Symbol   string `json:"symbol"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Exchange string `json:"exchange"`
	Currency string `json:"currency"`
}

type UpdateAssetRequest struct {
	ID       string `json:"id"`
	Symbol   string `json:"symbol"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Exchange string `json:"exchange"`
	Currency string `json:"currency"`
	entities.TimeStamp
}

// Response DTOs

// Entity -> Response
