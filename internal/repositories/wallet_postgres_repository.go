package repositories

import (
	"github.com/asfung/elara/database"
	"github.com/asfung/elara/internal/entities"
)

type WalletPostgresRepository struct {
	*BaseRepository[entities.Wallet]
}

func NewWalletPostgresRepository(db database.Database) WalletRepository {
	return &WalletPostgresRepository{
		BaseRepository: NewBaseRepository[entities.Wallet](db),
	}
}

func (w *WalletPostgresRepository) FindByUserId(userID string) (entities.Wallet, error) {
	var wallet entities.Wallet
	if err := w.db.GetDb().Where("user_id = ?", userID).First(&wallet).Error; err != nil {
		return entities.Wallet{}, err
	}
	return wallet, nil
}
