package repositories

import (
	"github.com/asfung/elara/database"
	"github.com/asfung/elara/internal/entities"
)

type BankPostgresRepository struct {
	*BaseRepository[entities.Bank]
}

func NewBankPostgresRepository(db database.Database) BankRepository {
	return &BankPostgresRepository{
		BaseRepository: NewBaseRepository[entities.Bank](db),
	}
}

func (b *BankPostgresRepository) FindBySwiftCode(swiftCode string) (entities.Bank, error) {
	var bank entities.Bank
	if err := b.db.GetDb().Where("swift_code = ?", swiftCode).First(&bank).Error; err != nil {
		return entities.Bank{}, err
	}
	return bank, nil
}
