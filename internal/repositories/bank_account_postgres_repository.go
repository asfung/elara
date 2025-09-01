package repositories

import (
	"github.com/asfung/elara/database"
	"github.com/asfung/elara/internal/entities"
)

type BankAccountPostgresRepository struct {
	*BaseRepository[entities.BankAccount]
}

func NewBankAccountPostgresRepository(db database.Database) BankAccountRepository {
	return &BankAccountPostgresRepository{
		BaseRepository: NewBaseRepository[entities.BankAccount](db),
	}
}
func (b *BankAccountPostgresRepository) FindByUserId(userID string) (entities.BankAccount, error) {
	var bankAccount entities.BankAccount
	if err := b.db.GetDb().Where("user_id = ?", userID).First(&bankAccount).Error; err != nil {
		return entities.BankAccount{}, err
	}
	return bankAccount, nil
}
