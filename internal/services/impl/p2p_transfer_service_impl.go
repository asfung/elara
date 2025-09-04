package impl

import (
	"errors"
	"time"

	"github.com/asfung/elara/internal/entities"
	"github.com/asfung/elara/internal/models"
	"github.com/asfung/elara/internal/repositories"
	"github.com/asfung/elara/internal/services"
)

type p2pTransferServiceImpl struct {
	p2pTransferRepo       repositories.P2PTransferRepository
	walletTransactionRepo repositories.WalletTransactionRepository
	walletService         services.WalletService
	userService           services.UserService
}

func NewP2PTransferServiceImpl(
	repo repositories.P2PTransferRepository,
	walletTransactionRepo repositories.WalletTransactionRepository,
	walletService services.WalletService,
	userService services.UserService,
) services.P2PTransferService {
	return &p2pTransferServiceImpl{
		p2pTransferRepo:       repo,
		walletTransactionRepo: walletTransactionRepo,
		walletService:         walletService,
		userService:           userService,
	}
}

func (p *p2pTransferServiceImpl) CreateP2PTransfer(req models.AddP2PTransferRequest) (entities.P2pTransfer, error) {
	senderWallet, err := p.walletService.GetWalletById(req.SenderWalletID)
	if err != nil {
		return entities.P2pTransfer{}, err
	}
	receiverWallet, err := p.walletService.GetWalletById(req.ReceiverWalletID)
	if err != nil {
		return entities.P2pTransfer{}, err
	}

	if senderWallet.Balance < req.Amount {
		return entities.P2pTransfer{}, errors.New("insufficient balance")
	}

	// 1. debit sender
	debitTxn, err := entities.NewWalletTransaction(senderWallet.ID, "debit", senderWallet.Currency, req.Message, req.Amount)
	if err != nil {
		return entities.P2pTransfer{}, err
	}
	createdDebit, err := p.walletTransactionRepo.Create(*debitTxn)
	if err != nil {
		return entities.P2pTransfer{}, err
	}

	// 2. credit receiver
	creditTxn, err := entities.NewWalletTransaction(receiverWallet.ID, "credit", req.Currency, "", req.Amount)
	if err != nil {
		return entities.P2pTransfer{}, err
	}
	createdCredit, err := p.walletTransactionRepo.Create(*creditTxn)
	if err != nil {
		return entities.P2pTransfer{}, err
	}

	// 3. update balances
	senderWallet.Balance -= req.Amount
	receiverWallet.Balance += req.Amount
	if _, err := p.walletService.UpdateWallet(models.UpdateWalletRequest{
		ID:      senderWallet.ID,
		Balance: senderWallet.Balance,
	}); err != nil {
		return entities.P2pTransfer{}, err
	}
	if _, err := p.walletService.UpdateWallet(models.UpdateWalletRequest{
		ID:      receiverWallet.ID,
		Balance: receiverWallet.Balance,
	}); err != nil {
		return entities.P2pTransfer{}, err
	}

	// 4. create transfer record
	newTransfer, err := entities.NewP2PTransfer(
		senderWallet.ID,
		receiverWallet.ID,
		createdDebit.ID,
		createdCredit.ID,
		req.Currency,
		req.Method,
		req.Message,
		req.Amount,
	)
	if err != nil {
		return entities.P2pTransfer{}, err
	}
	createdTransfer, err := p.p2pTransferRepo.Create(*newTransfer)
	if err != nil {
		return entities.P2pTransfer{}, err
	}

	return createdTransfer, nil
}

func (p *p2pTransferServiceImpl) UpdateP2PTransfer(req models.UpdateP2PTransferRequest) (entities.P2pTransfer, error) {
	transfer, err := p.p2pTransferRepo.FindById(req.ID)
	if err != nil {
		return entities.P2pTransfer{}, err
	}

	transfer.Message = req.Message
	transfer.Method = req.Method
	transfer.Status = req.Status
	transfer.UpdatedAt = time.Now()

	updated, err := p.p2pTransferRepo.Update(*transfer)
	if err != nil {
		return entities.P2pTransfer{}, err
	}
	return updated, nil
}

func (p *p2pTransferServiceImpl) GetP2PTransferById(id string) (entities.P2pTransfer, error) {
	tf, err := p.p2pTransferRepo.FindById(id)
	if err != nil {
		return entities.P2pTransfer{}, err
	}
	return *tf, nil
}

func (p *p2pTransferServiceImpl) DeleteP2PTransfer(id string) error {
	return p.p2pTransferRepo.Delete(id)
}
