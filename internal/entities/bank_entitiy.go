package entities

import (
	"time"

	"github.com/lucsky/cuid"
)

// BANKS
type Bank struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	SwiftCode string    `json:"swift_code"`
	Country   string    `json:"country"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

func NewBank(name, swiftCode, country, status string) (*Bank, error) {
	id := cuid.New()
	return &Bank{
		ID:        id,
		Name:      name,
		SwiftCode: swiftCode,
		Country:   country,
		Status:    status,
		CreatedAt: time.Now(),
	}, nil
}

// BANK ACCOUNTS
type BankAccount struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	BankID        string    `json:"bank_id"`
	AccountNumber string    `json:"account_number"`
	AccountType   string    `json:"account_type"`
	Verified      bool      `json:"verified"`
	CreatedAt     time.Time `json:"created_at"`
}

func NewBankAccount(userID, bankID, accountNumber, accountType string) (*BankAccount, error) {
	id := cuid.New()
	return &BankAccount{
		ID:            id,
		UserID:        userID,
		BankID:        bankID,
		AccountNumber: accountNumber,
		AccountType:   accountType,
		Verified:      false,
		CreatedAt:     time.Now(),
	}, nil
}
