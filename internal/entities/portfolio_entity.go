package entities

import "github.com/lucsky/cuid"

const (
	// Asset.Type
	STOCK  = "STOCK"
	ETF    = "ETF"
	CRYPTO = "CRYPTO"

	// Portfolio.Type
	PERSONAL   = "PERSONAL"
	RETIREMENT = "RETIREMENT"
)

type Asset struct {
	ID       string `json:"id" gorm:"primaryKey"`
	Symbol   string `json:"symbol" gorm:"uniqueIndex:idx_symbol"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Exchange string `json:"exchange" gorm:"uniqueIndex:idx_exchange"`
	Currency string `json:"currency"`
	TimeStamp
}

func NewAsset(symbol, type_, name, exchange, currency string) (*Asset, error) {
	id := cuid.New()
	return &Asset{
		ID:       id,
		Symbol:   symbol,
		Type:     type_,
		Exchange: exchange,
		Currency: currency,
	}, nil
}

type Portfolio struct {
	ID     string `json:"id" gorm:"primaryKey"`
	UserID string `json:"user_id" gorm:"uniqueIndex:idx_user_id"`
	Name   string `json:"name" gorm:"uniqueIndex:idx_name"`
	Type   string `json:"type"`
	TimeStamp
}

func NewPortfolio(userID, name, type_ string) (*Portfolio, error) {
	id := cuid.New()
	return &Portfolio{
		ID:     id,
		UserID: userID,
		Name:   name,
		Type:   type_,
	}, nil
}

type PortfolioAsset struct {
	ID              string `json:"id" gorm:"primaryKey"`
	PortfolioID     string `json:"portfolio_id" gorm:"uniqueIndex:idx_portfolio_id"`
	AssetID         string `json:"asset_id" gorm:"uniqueIndex:idx_asset_id"`
	Quantity        string `json:"quantity"`
	AverageBuyPrice string `json:"average_buy_price"`
	CurrentValue    string `json:"current_value"`
	TimeStamp
}

func NewPortfolioAsset(portfolioID, assetID, quantity, averageBuyPrice, currentValue string) (*PortfolioAsset, error) {
	id := cuid.New()
	return &PortfolioAsset{
		ID:              id,
		PortfolioID:     portfolioID,
		AssetID:         assetID,
		Quantity:        quantity,
		AverageBuyPrice: averageBuyPrice,
		CurrentValue:    currentValue,
	}, nil
}
