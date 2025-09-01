package models

import "github.com/asfung/elara/internal/entities"

// ========= BANK =========
// Request DTOs
type AddBankRequest struct {
	Name      string `json:"name" validate:"required"`
	SwiftCode string `json:"swift_code" validate:"required"`
	Country   string `json:"country" validate:"required"`
	Status    string `json:"status" validate:"required"`
}
type UpdateBankRequest struct {
	ID        string `json:"id" validate:"required"`
	Name      string `json:"name" validate:"omitempty"`
	SwiftCode string `json:"swift_code" validate:"omitempty"`
	Country   string `json:"country" validate:"omitempty"`
	Status    string `json:"status" validate:"omitempty"`
}

// Response DTOs
type BankResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	SwiftCode string `json:"swift_code"`
	Country   string `json:"country"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

// Entity -> Response
func ToBankResponse(bank entities.Bank) BankResponse {
	return BankResponse{
		ID:        bank.ID,
		Name:      bank.Name,
		SwiftCode: bank.SwiftCode,
		Country:   bank.Country,
		Status:    bank.Status,
		CreatedAt: bank.CreatedAt.String(),
	}
}

// ========= BANK ACCOUNT =========

// Request DTOs
type AddBankAccountRequest struct {
	UserID        string `json:"user_id" validate:"required"`
	BankID        string `json:"bank_id" validate:"required"`
	AccountNumber string `json:"account_number" validate:"required"`
	AccountType   string `json:"account_type" validate:"required"`
}

type UpdateBankAccountRequest struct {
	ID            string `json:"id" validate:"required"`
	AccountNumber string `json:"account_number" validate:"omitempty"`
	AccountType   string `json:"account_type" validate:"omitempty"`
	Verified      *bool  `json:"verified" validate:"omitempty"`
}

// Response DTOs
type BankAccountResponse struct {
	ID            string `json:"id"`
	UserId        string `json:"user_id"`
	BankId        string `json:"bank_id"`
	AccountNumber string `json:"account_number"`
	AccountType   string `json:"account_type"`
	Verified      bool   `json:"verified"`
	CreatedAt     string `json:"created_at"`
}

// Entity -> Response
func ToBankAccountResponse(bankAccount entities.BankAccount) BankAccountResponse {
	return BankAccountResponse{
		ID:            bankAccount.ID,
		UserId:        bankAccount.UserID,
		BankId:        bankAccount.BankID,
		AccountNumber: bankAccount.AccountNumber,
		AccountType:   bankAccount.AccountType,
		Verified:      bankAccount.Verified,
		CreatedAt:     bankAccount.CreatedAt.String(),
	}
}
