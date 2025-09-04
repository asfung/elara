package repositories

import (
	"github.com/asfung/elara/internal/entities"
	"github.com/asfung/elara/internal/models"
)

type WalletRepository interface {
	Repository[entities.Wallet]
	FindByUserId(userID string) (entities.Wallet, error)
}

type WalletTransactionRepository interface {
	Repository[entities.WalletTransaction]
	PaginateFindByUserId(req models.RequestParams, userID string) (models.PaginaterResolver, error)
	PaginateFindByWalletId(req models.RequestParams, walletID string) (models.PaginaterResolver, error)
}
