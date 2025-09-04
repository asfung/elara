package entities

import (
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

// P2P TRANSFERS
type P2pTransfer struct {
	ID                  string     `json:"id"`
	SenderWalletID      string     `json:"sender_wallet_id"`
	ReceiverWalletID    string     `json:"receiver_wallet_id"`
	DebitTransactionID  string     `json:"debit_transaction_id"`
	CreditTransactionID string     `json:"credit_transaction_id"`
	Amount              float64    `json:"amount"`
	Currency            string     `json:"currency"`
	Status              string     `json:"status"`
	Method              string     `json:"method"`
	Message             string     `json:"message"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
	CompletedAt         *time.Time `json:"completed_at,omitempty"`
}

const (
	P2PStatusPending = "pending"
	P2PStatusSuccess = "success"
	P2PStatusFailed  = "failed"
)

func NewP2PTransfer(senderID, receiverID, debitTxnID, creditTxnID, currency, method, message string, amount float64) (*P2pTransfer, error) {
	id, err := gonanoid.New()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	return &P2pTransfer{
		ID:                  id,
		SenderWalletID:      senderID,
		ReceiverWalletID:    receiverID,
		DebitTransactionID:  debitTxnID,
		CreditTransactionID: creditTxnID,
		Amount:              amount,
		Currency:            currency,
		Status:              P2PStatusPending,
		Method:              method,
		Message:             message,
		CreatedAt:           now,
		UpdatedAt:           now,
	}, nil
}
