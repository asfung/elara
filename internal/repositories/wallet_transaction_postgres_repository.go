package repositories

import (
	"fmt"

	"github.com/asfung/elara/database"
	"github.com/asfung/elara/internal/entities"
	"github.com/asfung/elara/internal/models"
)

type WalletTransactionPostgresRepository struct {
	*BaseRepository[entities.WalletTransaction]
}

func NewWalletTransactionPostgresRepository(db database.Database) WalletTransactionRepository {
	return &WalletTransactionPostgresRepository{
		BaseRepository: NewBaseRepository[entities.WalletTransaction](db),
	}
}

// DEPRECATED
func (w *WalletTransactionPostgresRepository) PaginateFindByUserId(req models.RequestParams, userID string) (models.PaginaterResolver, error) {
	stmt := w.db.GetDb().
		Where("user_id = ?", userID).
		Model(&entities.WalletTransaction{})

	paginator := new(models.PaginaterResolver).
		Stmt(stmt).
		Model(&[]entities.WalletTransaction{}).
		RequestParams(req)

	result, err := paginator.Paginate()
	if err != nil {
		return result, err
	}

	// type assert
	txsPtr, ok := result.Data.(*[]entities.WalletTransaction)
	if !ok {
		return result, fmt.Errorf("unexpected data type in pagination")
	}
	txs := *txsPtr

	//  collect walletIDs
	walletIDs := make([]string, 0, len(txs))
	for _, tx := range txs {
		walletIDs = append(walletIDs, tx.WalletID)
	}

	// query wallets
	var wallets []entities.Wallet
	if err := w.db.GetDb().Where("id IN ?", walletIDs).Find(&wallets).Error; err != nil {
		return result, err
	}

	// build map
	walletMap := make(map[string]entities.Wallet)
	for _, wallet := range wallets {
		walletMap[wallet.ID] = wallet
	}

	// map to response
	responses := make([]models.WalletTransactionResponse, 0, len(txs))
	for _, tx := range txs {
		wallet := walletMap[tx.WalletID]
		responses = append(responses, models.ToWalletTransactionWithWalletResponse(tx, &wallet))
	}

	result.Data = responses
	return result, nil
}

func (w *WalletTransactionPostgresRepository) PaginateFindByWalletId(req models.RequestParams, walletID string) (models.PaginaterResolver, error) {
	stmt := w.db.GetDb().
		Where("wallet_id = ?", walletID).
		Model(&entities.WalletTransaction{})

	paginator := new(models.PaginaterResolver).
		Stmt(stmt).
		Model(&[]entities.WalletTransaction{}).
		RequestParams(req)

	result, err := paginator.Paginate()
	if err != nil {
		return result, err
	}

	// ✅ FIX: type assert to *[]entities.WalletTransaction
	txsPtr, ok := result.Data.(*[]entities.WalletTransaction)
	if !ok {
		return result, fmt.Errorf("unexpected data type in pagination: %T", result.Data)
	}
	txs := *txsPtr

	// collect walletIDs
	walletIDs := make([]string, 0, len(txs))
	for _, tx := range txs {
		walletIDs = append(walletIDs, tx.WalletID)
	}

	// query wallets
	var wallets []entities.Wallet
	if err := w.db.GetDb().Where("id IN ?", walletIDs).Find(&wallets).Error; err != nil {
		return result, err
	}

	// map wallets by ID
	walletMap := make(map[string]entities.Wallet)
	for _, wallet := range wallets {
		walletMap[wallet.ID] = wallet
	}

	// build []WalletTransactionResponse
	responses := make([]models.WalletTransactionResponse, 0, len(txs))
	for _, tx := range txs {
		wallet := walletMap[tx.WalletID]
		responses = append(responses, models.ToWalletTransactionWithWalletResponse(tx, &wallet))
	}

	// override Data with responses
	result.Data = responses
	return result, nil

}

// func (w *WalletTransactionPostgresRepository) FindById(id any) (*entities.WalletTransaction, error) {
// 	panic("unimplemented")
// }
