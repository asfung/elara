package services

import (
	"github.com/asfung/elara/internal/entities"
	"github.com/asfung/elara/internal/models"
)

type WalletTransactionService interface {
	CreateWalletTransaction(req models.AddWalletTransactionRequest) (entities.WalletTransaction, entities.Wallet, error)
	UpdateWalletTransaction(req models.UpdateWalletTransactionRequest) (entities.WalletTransaction, entities.Wallet, error)
	GetWalletTransactionById(id string) (entities.WalletTransaction, entities.Wallet, error)
	DeleteWalletTransaction(id string) error
	GetWalletTransactionByUserIdPaginated(req models.RequestParams, userID string) (models.PaginaterResolver, error)
	GetWalletTransactionByWalletIdPaginated(req models.RequestParams, walletID string) (models.PaginaterResolver, error)
}
