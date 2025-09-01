package entities

import (
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

// P2P TRANSFERS
type P2pTransfer struct {
	ID               string    `json:"id"`
	SenderWalletID   string    `json:"sender_wallet_id"`
	ReceiverWalletID string    `json:"receiver_wallet_id"`
	Amount           float64   `json:"amount"`
	Currency         string    `json:"currency"`
	Status           string    `json:"status"`
	Method           string    `json:"method"`
	Message          string    `json:"message"`
	CreatedAt        time.Time `json:"created_at"`
}

func NewP2PTransfer(senderID, receiverID, currency, method, message string, amount float64) (*P2pTransfer, error) {
	id, err := gonanoid.New()
	if err != nil {
		return nil, err
	}
	return &P2pTransfer{
		ID:               id,
		SenderWalletID:   senderID,
		ReceiverWalletID: receiverID,
		Amount:           amount,
		Currency:         currency,
		Status:           "pending",
		Method:           method,
		Message:          message,
		CreatedAt:        time.Now(),
	}, nil
}
