package entities

import (
	"time"

	"github.com/lucsky/cuid"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

// TRANSACTION LIMITS
type TransactionLimit struct {
	ID            string  `json:"id"`
	UserID        string  `json:"user_id"`
	DailyLimit    float64 `json:"daily_limit"`
	MonthlyLimit  float64 `json:"monthly_limit"`
	SingleTxLimit float64 `json:"single_tx_limit"`
	Currency      string  `json:"currency"`
}

func NewTransactionLimit(userID, currency string, daily, monthly, single float64) (*TransactionLimit, error) {
	id := cuid.New()
	return &TransactionLimit{
		ID:            id,
		UserID:        userID,
		DailyLimit:    daily,
		MonthlyLimit:  monthly,
		SingleTxLimit: single,
		Currency:      currency,
	}, nil
}

// RISK FLAGS
type RiskFlag struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	TransactionID string    `json:"transaction_id"`
	FlagType      string    `json:"flag_type"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
}

func NewRiskFlag(userID, txID, flagType string) (*RiskFlag, error) {
	id, err := gonanoid.New()
	if err != nil {
		return nil, err
	}
	return &RiskFlag{
		ID:            id,
		UserID:        userID,
		TransactionID: txID,
		FlagType:      flagType,
		Status:        "pending",
		CreatedAt:     time.Now(),
	}, nil
}
