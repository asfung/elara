package entities

import (
	"time"

	"github.com/lucsky/cuid"
	gonanoid "github.com/matoous/go-nanoid/v2"
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

// CARDS
type Card struct {
	ID                 string    `json:"id"`
	UserID             string    `json:"user_id"`
	CardNumberHash     string    `json:"card_number_hash"`
	CardType           string    `json:"card_type"`
	ExpiryDate         string    `json:"expiry_date"`
	Status             string    `json:"status"`
	TokenizedReference string    `json:"tokenized_reference"`
	CreatedAt          time.Time `json:"created_at"`
}

func NewCard(userID, cardHash, cardType, expiry string) (*Card, error) {
	id, err := gonanoid.New()
	if err != nil {
		return nil, err
	}
	return &Card{
		ID:                 id,
		UserID:             userID,
		CardNumberHash:     cardHash,
		CardType:           cardType,
		ExpiryDate:         expiry,
		Status:             "active",
		TokenizedReference: "",
		CreatedAt:          time.Now(),
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
