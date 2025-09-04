package services

import (
	"github.com/asfung/elara/internal/entities"
	"github.com/asfung/elara/internal/models"
)

type BankAccountService interface {
	CreateBankAccount(req models.AddBankAccountRequest) (entities.BankAccount, *entities.Bank, error)
	UpdateBankAccount(req models.UpdateBankAccountRequest) (entities.BankAccount, *entities.Bank, error)
	GetBankAccountById(id string) (entities.BankAccount, *entities.Bank, error)
	DeleteBankAccount(id string) error
	GetBankAccountByUserId(userID string) (entities.BankAccount, *entities.Bank, error)
}
