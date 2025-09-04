package impl

import (
	"fmt"

	"github.com/asfung/elara/internal/entities"
	"github.com/asfung/elara/internal/models"
	"github.com/asfung/elara/internal/repositories"
	"github.com/asfung/elara/internal/services"
)

type walletServiceImpl struct {
	repo        repositories.WalletRepository
	userService services.UserService
}

func NewWalletServiceImpl(repo repositories.WalletRepository, userService services.UserService) services.WalletService {
	return &walletServiceImpl{
		repo:        repo,
		userService: userService,
	}
}

func (w *walletServiceImpl) CreateWallet(req models.AddWalletRequest) (entities.Wallet, error) {
	_, err := w.userService.GetUserByUserId(req.UserID)
	if err != nil {
		return entities.Wallet{}, err
	}

	card, err := entities.NewWallet(req.UserID, req.Currency)
	if err != nil {
		return entities.Wallet{}, err
	}

	createdWallet, err := w.repo.Create(*card)
	if err != nil {
		return entities.Wallet{}, err
	}

	return createdWallet, nil
}

func (w *walletServiceImpl) UpdateWallet(req models.UpdateWalletRequest) (entities.Wallet, error) {
	wallet, err := w.repo.FindById(req.ID)
	if err != nil {
		return entities.Wallet{}, err
	}

	if req.Balance != 0 {
		wallet.Balance = req.Balance
	}
	if req.Status != "" {
		wallet.Status = req.Status
	}

	updatedWallet, err := w.repo.Update(*wallet)
	if err != nil {
		return entities.Wallet{}, err
	}

	return updatedWallet, nil
}

func (w *walletServiceImpl) GetWalletById(id string) (entities.Wallet, error) {
	wallet, err := w.repo.FindById(id)
	if err != nil {
		return entities.Wallet{}, err
	}
	return *wallet, nil
}

func (w *walletServiceImpl) DeleteWallet(id string) error {
	return w.repo.Delete(id)
}

func (w *walletServiceImpl) GetWalletByUserId(userID string) (entities.Wallet, error) {
	wallet, err := w.repo.FindByUserId(userID)
	if err != nil {
		return entities.Wallet{}, err
	}
	return wallet, nil
}

func (w *walletServiceImpl) UpdateWalletBalance(walletID string, amount float64, isCredit bool) (entities.Wallet, error) {
	wallet, err := w.repo.FindById(walletID)
	if err != nil {
		return entities.Wallet{}, err
	}

	if isCredit {
		wallet.Balance += amount
	} else {
		if wallet.Balance < amount {
			return entities.Wallet{}, fmt.Errorf("insufficient balance")
		}
		wallet.Balance -= amount
	}

	// wallet.UpdatedAt = time.Now()

	updatedWallet, err := w.repo.Update(*wallet)
	if err != nil {
		return entities.Wallet{}, err
	}

	return updatedWallet, nil
}
