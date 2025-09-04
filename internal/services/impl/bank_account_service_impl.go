package impl

import (
	"errors"

	"github.com/asfung/elara/internal/entities"
	"github.com/asfung/elara/internal/models"
	"github.com/asfung/elara/internal/repositories"
	"github.com/asfung/elara/internal/services"
)

type bankAccountServiceImpl struct {
	repo        repositories.BankAccountRepository
	bankService services.BankService
}

func NewBankAccountServiceImpl(repo repositories.BankAccountRepository, bankService services.BankService) services.BankAccountService {
	return &bankAccountServiceImpl{
		repo:        repo,
		bankService: bankService,
	}
}

func (b *bankAccountServiceImpl) CreateBankAccount(req models.AddBankAccountRequest) (entities.BankAccount, *entities.Bank, error) {
	bankAccountExist, err := b.bankService.GetBankById(req.BankID)
	if err != nil {
		return entities.BankAccount{}, &entities.Bank{}, err
	}

	if bankAccountExist.ID == "" {
		return entities.BankAccount{}, &entities.Bank{}, errors.New("bank_id not exist in bank")
	}

	bankAccount, err := entities.NewBankAccount(req.UserID, req.BankID, req.AccountNumber, req.AccountType)
	if err != nil {
		return entities.BankAccount{}, &entities.Bank{}, err
	}

	createdBankAccount, err := b.repo.Create(*bankAccount)
	if err != nil {
		return entities.BankAccount{}, &entities.Bank{}, err
	}

	bank, err := b.bankService.GetBankById(req.BankID)
	if err != nil {
		return entities.BankAccount{}, &entities.Bank{}, err
	}

	return createdBankAccount, &bank, nil
}

func (b *bankAccountServiceImpl) UpdateBankAccount(req models.UpdateBankAccountRequest) (entities.BankAccount, *entities.Bank, error) {
	bankAccount, err := b.repo.FindById(req.ID)
	if err != nil {
		return entities.BankAccount{}, &entities.Bank{}, err
	}

	if req.AccountNumber != "" {
		bankAccount.AccountNumber = req.AccountNumber
	}
	if req.AccountType != "" {
		bankAccount.AccountType = req.AccountType
	}
	if req.Verified != nil {
		bankAccount.Verified = *req.Verified
	}

	updatedBankAccount, err := b.repo.Update(*bankAccount)
	if err != nil {
		return entities.BankAccount{}, &entities.Bank{}, err
	}

	bank, err := b.bankService.GetBankById(updatedBankAccount.BankID)
	if err != nil {
		return entities.BankAccount{}, &entities.Bank{}, err
	}

	return updatedBankAccount, &bank, nil
}

func (b *bankAccountServiceImpl) GetBankAccountById(id string) (entities.BankAccount, *entities.Bank, error) {
	bankAccount, err := b.repo.FindById(id)
	if err != nil {
		return entities.BankAccount{}, &entities.Bank{}, err
	}

	bank, err := b.bankService.GetBankById(bankAccount.BankID)
	if err != nil {
		return entities.BankAccount{}, &entities.Bank{}, err
	}

	return *bankAccount, &bank, nil
}

func (b *bankAccountServiceImpl) DeleteBankAccount(id string) error {
	return b.repo.Delete(id)
}

func (b *bankAccountServiceImpl) GetBankAccountByUserId(userID string) (entities.BankAccount, *entities.Bank, error) {
	bankAccount, err := b.repo.FindByUserId(userID)
	if err != nil {
		return entities.BankAccount{}, &entities.Bank{}, err
	}

	bank, err := b.bankService.GetBankById(bankAccount.BankID)
	if err != nil {
		return entities.BankAccount{}, &entities.Bank{}, err
	}

	return bankAccount, &bank, nil
}
