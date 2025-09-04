package repositories

import (
	"github.com/asfung/elara/database"
	"github.com/asfung/elara/internal/entities"
	"github.com/asfung/elara/internal/models"
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

func (b *BankPostgresRepository) PaginateBanks(req models.RequestParams) (models.PaginaterResolver, error) {
	stmt := b.db.GetDb().Model(&entities.Bank{})

	paginator := new(models.PaginaterResolver).
		Stmt(stmt).
		Model(&[]entities.Bank{}). // important: pass a slice, not single entity
		RequestParams(req)

	return paginator.Paginate()
}
