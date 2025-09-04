package models

import "github.com/asfung/elara/internal/entities"

// Request DTOs
type AddP2PTransferRequest struct {
	SenderWalletID   string  `json:"sender_wallet_id" validate:"required"`
	ReceiverWalletID string  `json:"receiver_wallet_id" validate:"required"`
	Currency         string  `json:"currency" validate:"required"`
	Method           string  `json:"method" validate:"required"`
	Message          string  `json:"message" validate:"required"`
	Amount           float64 `json:"amount" validate:"required"`
}

type UpdateP2PTransferRequest struct {
	ID      string `json:"id" validate:"required"`
	Status  string `json:"status" validate:"required"`
	Message string `json:"message" validate:"required"`
	Method  string `json:"method" validate:"required"`
}

// Response DTOs
type P2PTransferResponse struct {
	ID                  string  `json:"id"`
	SenderWalletID      string  `json:"sender_wallet_id"`
	ReceiverWalletID    string  `json:"receiver_wallet_id"`
	DebitTransactionID  string  `json:"debit_transaction_id"`
	CreditTransactionID string  `json:"credit_transaction_id"`
	Amount              float64 `json:"amount"`
	Currency            string  `json:"currency"`
	Status              string  `json:"status"`
	Method              string  `json:"method"`
	Message             string  `json:"message"`
	CreatedAt           string  `json:"created_at"`
	UpdatedAt           string  `json:"updated_at"`
	CompletedAt         *string `json:"completed_at,omitempty"`
}

// Entity -> Response
func ToP2PTransferResponse(p2p entities.P2pTransfer) P2PTransferResponse {
	return P2PTransferResponse{
		ID:                  p2p.ID,
		SenderWalletID:      p2p.SenderWalletID,
		ReceiverWalletID:    p2p.ReceiverWalletID,
		DebitTransactionID:  p2p.DebitTransactionID,
		CreditTransactionID: p2p.CreditTransactionID,
		Amount:              p2p.Amount,
		Currency:            p2p.Currency,
		Status:              p2p.Status,
		Method:              p2p.Method,
		Message:             p2p.Message,
		CreatedAt:           p2p.CreatedAt.String(),
		UpdatedAt:           p2p.UpdatedAt.String(),
		CompletedAt:         nil,
	}
}
