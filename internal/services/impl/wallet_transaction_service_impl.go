package impl

import (
	"errors"

	"github.com/asfung/elara/internal/entities"
	"github.com/asfung/elara/internal/models"
	"github.com/asfung/elara/internal/repositories"
	"github.com/asfung/elara/internal/services"
)

type walletTransactionServiceImpl struct {
	walletTransactionRepo repositories.WalletTransactionRepository
	walletService         services.WalletService
}

func NewWalletTransactionServiceImpl(
	walletTransactionRepo repositories.WalletTransactionRepository,
	walletService services.WalletService,
) services.WalletTransactionService {
	return &walletTransactionServiceImpl{
		walletTransactionRepo: walletTransactionRepo,
		walletService:         walletService,
	}
}

func (w *walletTransactionServiceImpl) CreateWalletTransaction(req models.AddWalletTransactionRequest) (entities.WalletTransaction, entities.Wallet, error) {
	wallet, err := w.walletService.GetWalletById(req.WalletID)
	if err != nil {
		return entities.WalletTransaction{}, entities.Wallet{}, err
	}

	if req.Type == "debit" && wallet.Balance < req.Amount {
		return entities.WalletTransaction{}, entities.Wallet{}, errors.New("insufficient balance")
	}

	walletTransaction, err := entities.NewWalletTransaction(
		req.WalletID,
		req.Type,
		wallet.Currency,
		req.ReferenceID,
		req.Amount,
	)
	if err != nil {
		return entities.WalletTransaction{}, entities.Wallet{}, err
	}

	if req.Type == "credit" {
		wallet.Balance += req.Amount
	} else if req.Type == "debit" {
		wallet.Balance -= req.Amount
	}

	createdWalletTransaction, err := w.walletTransactionRepo.Create(*walletTransaction)
	if err != nil {
		return entities.WalletTransaction{}, entities.Wallet{}, err
	}

	if _, err := w.walletService.UpdateWalletBalance(wallet.ID, req.Amount, walletTransaction.Type == "credit"); err != nil {
		return entities.WalletTransaction{}, entities.Wallet{}, err
	}

	return createdWalletTransaction, wallet, nil

}

func (w *walletTransactionServiceImpl) UpdateWalletTransaction(req models.UpdateWalletTransactionRequest) (entities.WalletTransaction, entities.Wallet, error) {
	existingTxn, err := w.walletTransactionRepo.FindById(req.ID)
	if err != nil {
		return entities.WalletTransaction{}, entities.Wallet{}, err
	}

	wallet, err := w.walletService.GetWalletById(existingTxn.WalletID)
	if err != nil {
		return entities.WalletTransaction{}, entities.Wallet{}, err
	}

	switch existingTxn.Type {
	case "credit":
		wallet.Balance -= existingTxn.Amount
	case "debit":
		wallet.Balance += existingTxn.Amount
	}

	if req.Type == "debit" && wallet.Balance < req.Amount {
		return entities.WalletTransaction{}, entities.Wallet{}, errors.New("insufficient balance")
	}

	switch req.Type {
	case "credit":
		wallet.Balance += req.Amount
	case "debit":
		wallet.Balance -= req.Amount
	}

	existingTxn.Type = req.Type
	existingTxn.Amount = req.Amount
	existingTxn.ReferenceID = req.ReferenceID
	existingTxn.Currency = wallet.Currency
	existingTxn.Status = req.Status

	updatedTxn, err := w.walletTransactionRepo.Update(*existingTxn)
	if err != nil {
		return entities.WalletTransaction{}, entities.Wallet{}, err
	}

	if _, err := w.walletService.UpdateWalletBalance(wallet.ID, req.Amount, existingTxn.Type == "credit"); err != nil {
		return entities.WalletTransaction{}, entities.Wallet{}, err
	}

	return updatedTxn, wallet, nil
}

func (w *walletTransactionServiceImpl) GetWalletTransactionById(id string) (entities.WalletTransaction, entities.Wallet, error) {
	walletTransaction, err := w.walletTransactionRepo.FindById(id)
	if err != nil {
		return entities.WalletTransaction{}, entities.Wallet{}, err
	}
	wallet, err := w.walletService.GetWalletById(walletTransaction.WalletID)
	if err != nil {
		return entities.WalletTransaction{}, entities.Wallet{}, err
	}
	return *walletTransaction, wallet, nil
}

func (w *walletTransactionServiceImpl) DeleteWalletTransaction(id string) error {
	return w.walletTransactionRepo.Delete(id)
}

func (w *walletTransactionServiceImpl) GetWalletTransactionByUserIdPaginated(req models.RequestParams, userID string) (models.PaginaterResolver, error) {
	return w.walletTransactionRepo.PaginateFindByUserId(req, userID)
}

func (w *walletTransactionServiceImpl) GetWalletTransactionByWalletIdPaginated(req models.RequestParams, walletID string) (models.PaginaterResolver, error) {
	return w.walletTransactionRepo.PaginateFindByWalletId(req, walletID)
}
