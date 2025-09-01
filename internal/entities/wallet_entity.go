package entities

import (
	"time"

	"github.com/lucsky/cuid"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

// WALLETS
type Wallet struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Balance   float64   `json:"balance"`
	Currency  string    `json:"currency"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

func NewWallet(userID, currency string) (*Wallet, error) {
	id, err := gonanoid.New()
	if err != nil {
		return nil, err
	}
	return &Wallet{
		ID:        id,
		UserID:    userID,
		Balance:   0,
		Currency:  currency,
		Status:    "active",
		CreatedAt: time.Now(),
	}, nil
}

// WALLET TRANSACTIONS
type WalletTransaction struct {
	ID          string    `json:"id"`
	WalletID    string    `json:"wallet_id"`
	Type        string    `json:"type"`
	Amount      float64   `json:"amount"`
	Currency    string    `json:"currency"`
	Status      string    `json:"status"`
	ReferenceID string    `json:"reference_id"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewWalletTransaction(walletID, txType, currency, reference string, amount float64) (*WalletTransaction, error) {
	id := cuid.New()
	return &WalletTransaction{
		ID:          id,
		WalletID:    walletID,
		Type:        txType,
		Amount:      amount,
		Currency:    currency,
		Status:      "pending",
		ReferenceID: reference,
		CreatedAt:   time.Now(),
	}, nil
}
