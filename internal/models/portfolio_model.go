package models

// Request DTOs
type AddAPortfolioRequest struct {
	UserID string `json:"user_id" gorm:"uniqueIndex:idx_user_id"`
	Name   string `json:"name" gorm:"uniqueIndex:idx_name"`
	Type   string `json:"type"`
}
type AddAPortfolioAssetRequest struct {
	PortfolioID     string `json:"portfolio_id"`
	AssetID         string `json:"asset_id"`
	Quantity        string `json:"quantity"`
	AverageBuyPrice string `json:"average_buy_price"`
	CurrentValue    string `json:"current_value"`
}

type UpdatePortfolioRequest struct {
	ID     string `json:"id" gorm:"primaryKey"`
	UserID string `json:"user_id" gorm:"uniqueIndex:idx_user_id"`
	Name   string `json:"name" gorm:"uniqueIndex:idx_name"`
	Type   string `json:"type"`
}
type UpdatePortfolioAssetRequest struct {
	ID              string `json:"id"`
	PortfolioID     string `json:"portfolio_id"`
	AssetID         string `json:"asset_id"`
	Quantity        string `json:"quantity"`
	AverageBuyPrice string `json:"average_buy_price"`
	CurrentValue    string `json:"current_value"`
}

// Response DTOs

// Entity -> Response
