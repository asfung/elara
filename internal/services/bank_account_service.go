package services

import (
	"github.com/asfung/elara/internal/entities"
	"github.com/asfung/elara/internal/models"
)

type BankAccountService interface {
	CreateBankAccount(req models.AddBankAccountRequest) (entities.BankAccount, error)
	UpdateBankAccount(req models.UpdateBankAccountRequest) (entities.BankAccount, error)
	GetBankAccountById(id string) (entities.BankAccount, error)
	DeleteBankAccount(id string) error
	GetBankAccountByUserId(userID string) (entities.BankAccount, error)
}
