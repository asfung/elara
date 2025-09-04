package models

import "github.com/asfung/elara/internal/entities"

// ========= WALLET =========

// Request DTOs
type AddWalletRequest struct {
	UserID   string `json:"user_id" validate:"required"`
	Currency string `json:"currency" validate:"required"`
	Status   string `json:"status" validate:"required"`
}

type UpdateWalletRequest struct {
	ID      string  `json:"id" validate:"omitempty"`
	Balance float64 `json:"balance" validate:"omitempty"`
	Status  string  `json:"status" validate:"omitempty"`
}

// Response DTOs
type WalletResponse struct {
	ID        string  `json:"id"`
	UserID    string  `json:"user_id"`
	Balance   float64 `json:"balance"`
	Currency  string  `json:"currency"`
	Status    string  `json:"status"`
	CreatedAt string  `json:"created_at"`
}

// Entity -> Response
func ToWalletResponse(wallet entities.Wallet) WalletResponse {
	return WalletResponse{
		ID:        wallet.ID,
		UserID:    wallet.UserID,
		Balance:   wallet.Balance,
		Currency:  wallet.Currency,
		Status:    wallet.Status,
		CreatedAt: wallet.CreatedAt.String(),
	}
}

// ========= WALLET TRANSACTION =========

// Request DTOs
type AddWalletTransactionRequest struct {
	WalletID    string  `json:"wallet_id" validate:"required"`
	Type        string  `json:"type" validate:"required"`
	Amount      float64 `json:"amount" validate:"required"`
	Currency    string  `json:"currency" validate:"required"`
	Status      string  `json:"status" validate:"required"`
	ReferenceID string  `json:"reference_id" validate:"required"`
}

type UpdateWalletTransactionRequest struct {
	ID          string  `json:"id" validate:"omitempty"`
	WalletID    string  `json:"wallet_id" validate:"required"`
	Type        string  `json:"type" validate:"required"`
	Amount      float64 `json:"amount" validate:"required"`
	Currency    string  `json:"currency" validate:"required"`
	Status      string  `json:"status" validate:"required"`
	ReferenceID string  `json:"reference_id" validate:"required"`
}

// Response DTOs
type WalletTransactionResponse struct {
	ID       string          `json:"id"`
	WalletID string          `json:"wallet_id"`
	Type     string          `json:"type"`
	Amount   float64         `json:"amount"`
	Wallet   *WalletResponse `json:"wallet,omitempty"`
}

// Entity -> Response
func ToWalletTransactionResponse(walletTransaction entities.WalletTransaction) WalletTransactionResponse {
	return WalletTransactionResponse{
		ID:       walletTransaction.ID,
		WalletID: walletTransaction.WalletID,
		Type:     walletTransaction.Type,
		Amount:   walletTransaction.Amount,
	}
}

func ToWalletTransactionWithWalletResponse(walletTransaction entities.WalletTransaction, wallet *entities.Wallet) WalletTransactionResponse {
	var walletResponse *WalletResponse
	if wallet != nil {
		walletResponse = &WalletResponse{
			ID:        wallet.ID,
			UserID:    wallet.UserID,
			Balance:   wallet.Balance,
			Currency:  wallet.Currency,
			Status:    wallet.Status,
			CreatedAt: wallet.CreatedAt.String(),
		}
	}

	return WalletTransactionResponse{
		ID:       walletTransaction.ID,
		WalletID: walletTransaction.WalletID,
		Type:     walletTransaction.Type,
		Amount:   walletTransaction.Amount,
		Wallet:   walletResponse,
	}
}
