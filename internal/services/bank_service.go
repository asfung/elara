package services

import (
	"github.com/asfung/elara/internal/entities"
	"github.com/asfung/elara/internal/models"
)

type BankService interface {
	CreateBank(req models.AddBankRequest) (entities.Bank, error)
	UpdateBank(req models.UpdateBankRequest) (entities.Bank, error)
	GetBankById(id string) (entities.Bank, error)
	DeleteBank(id string) error
	GetBankBySwiftCode(swiftCode string) (entities.Bank, error)
}
