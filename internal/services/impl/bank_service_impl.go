package impl

import (
	"errors"

	"github.com/asfung/elara/internal/entities"
	"github.com/asfung/elara/internal/models"
	"github.com/asfung/elara/internal/repositories"
	"github.com/asfung/elara/internal/services"
)

type bankServiceImpl struct {
	repo repositories.BankRepository
}

func NewBankServiceImpl(repo repositories.BankRepository) services.BankService {
	return &bankServiceImpl{
		repo: repo,
	}
}

func (b *bankServiceImpl) CreateBank(req models.AddBankRequest) (entities.Bank, error) {
	bankExist, _ := b.repo.FindBySwiftCode(req.SwiftCode)
	if bankExist.ID != "" {
		return entities.Bank{}, errors.New("bank with already exists")
	}

	bank, err := entities.NewBank(req.Name, req.SwiftCode, req.Country, req.Status)
	if err != nil {
		return entities.Bank{}, err
	}
	createdBank, err := b.repo.Create(*bank)
	if err != nil {
		return entities.Bank{}, err
	}
	return createdBank, nil
}

func (b *bankServiceImpl) UpdateBank(req models.UpdateBankRequest) (entities.Bank, error) {
	bank, err := b.repo.FindById(req.ID)
	if err != nil {
		return entities.Bank{}, err
	}
	if req.Name != "" {
		bank.Name = req.Name
	}
	if req.SwiftCode != "" {
		bank.SwiftCode = req.SwiftCode
	}
	if req.Country != "" {
		bank.Country = req.Country
	}
	if req.Status != "" {
		bank.Status = req.Status
	}

	updatedBank, err := b.repo.Update(*bank)
	if err != nil {
		return entities.Bank{}, err
	}

	return updatedBank, nil
}

func (b *bankServiceImpl) GetBankById(id string) (entities.Bank, error) {
	bank, err := b.repo.FindById(id)
	if err != nil {
		return entities.Bank{}, err
	}
	return *bank, nil
}

func (b *bankServiceImpl) DeleteBank(id string) error {
	return b.repo.Delete(id)
}

func (b *bankServiceImpl) GetBankBySwiftCode(swiftCode string) (entities.Bank, error) {
	bank, err := b.repo.FindBySwiftCode(swiftCode)
	if err != nil {
		return entities.Bank{}, err
	}
	return bank, nil
}

func (b *bankServiceImpl) GetBanksPaginated(req models.RequestParams) (models.PaginaterResolver, error) {
	return b.repo.PaginateBanks(req)
}
