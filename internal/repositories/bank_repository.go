package repositories

import "github.com/asfung/elara/internal/entities"

type BankRepository interface {
	Repository[entities.Bank]
	FindBySwiftCode(swiftCode string) (entities.Bank, error)
}

type BankAccountRepository interface {
	Repository[entities.BankAccount]
	FindByUserId(userID string) (entities.BankAccount, error)
}
