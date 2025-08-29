package entities

import (
	"time"

	"github.com/lucsky/cuid"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

// BILLERS
type Biller struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Category     string `json:"category"`
	ProviderCode string `json:"provider_code"`
	Country      string `json:"country"`
	Status       string `json:"status"`
}

func NewBiller(name, category, providerCode, country string) (*Biller, error) {
	id := cuid.New()
	return &Biller{
		ID:           id,
		Name:         name,
		Category:     category,
		ProviderCode: providerCode,
		Country:      country,
		Status:       "active",
	}, nil
}

// BILL PAYMENTS
type BillPayment struct {
	ID            string    `json:"id"`
	WalletID      string    `json:"wallet_id"`
	BillerID      string    `json:"biller_id"`
	Amount        float64   `json:"amount"`
	Currency      string    `json:"currency"`
	Status        string    `json:"status"`
	BillReference string    `json:"bill_reference"`
	CreatedAt     time.Time `json:"created_at"`
}

func NewBillPayment(walletID, billerID, billRef, currency string, amount float64) (*BillPayment, error) {
	id, err := gonanoid.New()
	if err != nil {
		return nil, err
	}
	return &BillPayment{
		ID:            id,
		WalletID:      walletID,
		BillerID:      billerID,
		Amount:        amount,
		Currency:      currency,
		Status:        "pending",
		BillReference: billRef,
		CreatedAt:     time.Now(),
	}, nil
}
