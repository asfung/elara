package repositories

import (
	"github.com/asfung/elara/internal/entities"
	"github.com/asfung/elara/internal/models"
)

type BankRepository interface {
	Repository[entities.Bank]
	FindBySwiftCode(swiftCode string) (entities.Bank, error)
	PaginateBanks(req models.RequestParams) (models.PaginaterResolver, error)
}

type BankAccountRepository interface {
	Repository[entities.BankAccount]
	FindByUserId(userID string) (entities.BankAccount, error)
}
