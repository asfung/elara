package services

import (
	"github.com/asfung/elara/internal/entities"
	"github.com/asfung/elara/internal/models"
)

type WalletService interface {
	CreateWallet(req models.AddWalletRequest) (entities.Wallet, error)
	UpdateWallet(req models.UpdateWalletRequest) (entities.Wallet, error)
	GetWalletById(id string) (entities.Wallet, error)
	DeleteWallet(id string) error
	GetWalletByUserId(userID string) (entities.Wallet, error)
	UpdateWalletBalance(walletID string, amount float64, isCredit bool) (entities.Wallet, error)
}
