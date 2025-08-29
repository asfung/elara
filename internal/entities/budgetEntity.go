package entities

import (
	"time"

	"github.com/lucsky/cuid"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

// CATEGORIES (optional - budgeting)
type Category struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ParentID string `json:"parent_id"`
}

func NewCategory(name, parentID string) (*Category, error) {
	id := cuid.New()
	return &Category{
		ID:       id,
		Name:     name,
		ParentID: parentID,
	}, nil
}

// EXPENSES (optional - budgeting)
type Expense struct {
	ID                  string    `json:"id"`
	WalletTransactionID string    `json:"wallet_transaction_id"`
	CategoryID          string    `json:"category_id"`
	Amount              float64   `json:"amount"`
	Note                string    `json:"note"`
	CreatedAt           time.Time `json:"created_at"`
}

func NewExpense(txID, catID, note string, amount float64) (*Expense, error) {
	id, err := gonanoid.New()
	if err != nil {
		return nil, err
	}
	return &Expense{
		ID:                  id,
		WalletTransactionID: txID,
		CategoryID:          catID,
		Amount:              amount,
		Note:                note,
		CreatedAt:           time.Now(),
	}, nil
}
