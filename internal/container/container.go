package container

import (
	"github.com/asfung/elara/database"
	"github.com/asfung/elara/internal/handlers"
	"github.com/asfung/elara/internal/repositories"
	"github.com/asfung/elara/internal/services"
	"github.com/asfung/elara/internal/services/impl"
)

type Container struct {
	AuthService              services.AuthService
	AuthHandler              *handlers.AuthHandler
	BankHandler              *handlers.BankHandler
	BankAccountHandler       *handlers.BankAccountHandler
	CardHandler              *handlers.CardHandler
	WalletHandler            *handlers.WalletHandler
	WalletTransactionHandler *handlers.WalletTransactionHandler
	P2PTransferHandler       *handlers.P2PTransferHandler
}

func NewContainer(db database.Database) *Container {
	// REPOSITORIES
	userRepo := repositories.NewUserPostgresRepository(db)
	authRepo := repositories.NewAuthPostgresRepository(db)
	bankRepo := repositories.NewBankPostgresRepository(db)
	bankAccountRepo := repositories.NewBankAccountPostgresRepository(db)
	cardRepo := repositories.NewCardPostrgresRepository(db)
	walletRepo := repositories.NewWalletPostgresRepository(db)
	walletTransactionRepo := repositories.NewWalletTransactionPostgresRepository(db)
	p2pTransferRepo := repositories.NewP2PTransferPostgresRepository(db)

	// SERVICES
	userService := impl.NewUserServiceImpl(userRepo)
	authService := impl.NewAuthServiceImpl(authRepo, userRepo, userService)
	bankService := impl.NewBankServiceImpl(bankRepo)
	bankAccountService := impl.NewBankAccountServiceImpl(bankAccountRepo, bankService)
	cardService := impl.NewCardServiceImpl(cardRepo, userService)
	walletService := impl.NewWalletServiceImpl(walletRepo, userService)
	walletTransactionService := impl.NewWalletTransactionServiceImpl(walletTransactionRepo, walletService)
	p2pTransferService := impl.NewP2PTransferServiceImpl(p2pTransferRepo, walletTransactionRepo, walletService, userService)

	// HANDLERS
	authHandler := handlers.NewAuthHandler(authService)
	bankHandler := handlers.NewBankHandler(bankService)
	bankAccountHandler := handlers.NewBankAccountHandler(bankAccountService)
	cardHandler := handlers.NewCardHandler(cardService)
	walletHandler := handlers.NewWalletHandler(walletService)
	walletTransactionHandler := handlers.NewWalletTransactionHandler(walletTransactionService)
	p2pTransferHandler := handlers.NewP2PTransferHandler(p2pTransferService)

	return &Container{
		AuthService:              authService,
		AuthHandler:              authHandler,
		BankHandler:              bankHandler,
		BankAccountHandler:       bankAccountHandler,
		CardHandler:              cardHandler,
		WalletHandler:            walletHandler,
		WalletTransactionHandler: walletTransactionHandler,
		P2PTransferHandler:       p2pTransferHandler,
	}
}
