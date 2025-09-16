package container

import (
	"os"
	"path/filepath"

	"github.com/asfung/elara/database"
	"github.com/asfung/elara/internal/handlers"
	"github.com/asfung/elara/internal/repositories"
	"github.com/asfung/elara/internal/services"
	"github.com/asfung/elara/internal/services/impl"
)

type Container struct {
	AuthService services.AuthService
	UserService services.UserService
	RoleService services.RoleService
	SmtpService services.SmtpService
	OtpService  services.OTPService

	AuthHandler              *handlers.AuthHandler
	BankHandler              *handlers.BankHandler
	BankAccountHandler       *handlers.BankAccountHandler
	CardHandler              *handlers.CardHandler
	WalletHandler            *handlers.WalletHandler
	WalletTransactionHandler *handlers.WalletTransactionHandler
	P2PTransferHandler       *handlers.P2PTransferHandler
}

func NewContainer(db database.Database) *Container {
	SMTP_FROM := os.Getenv("SMTP_FROM")
	SMTP_PASSWORD := os.Getenv("SMTP_PASSWORD")
	SMTP_HOST := os.Getenv("SMTP_HOST")
	SMTP_PORT := os.Getenv("SMTP_PORT")
	cwd, _ := os.Getwd()
	TemplateHTML := filepath.Join(cwd, "templates", "email.html")

	// REPOSITORIES
	userRepo := repositories.NewUserPostgresRepository(db)
	roleRepo := repositories.NewRolePostgresRepository(db)
	authRepo := repositories.NewAuthPostgresRepository(db)
	bankRepo := repositories.NewBankPostgresRepository(db)
	bankAccountRepo := repositories.NewBankAccountPostgresRepository(db)
	cardRepo := repositories.NewCardPostrgresRepository(db)
	walletRepo := repositories.NewWalletPostgresRepository(db)
	walletTransactionRepo := repositories.NewWalletTransactionPostgresRepository(db)
	p2pTransferRepo := repositories.NewP2PTransferPostgresRepository(db)
	otpRepo := repositories.NewOTPPotgresRepository(db)

	// SERVICES
	otpService := impl.NewOTPServiceImpl(otpRepo, userRepo)
	smtpService := impl.NewSmtpServiceImpl(SMTP_HOST, SMTP_PORT, SMTP_FROM, SMTP_PASSWORD, TemplateHTML)
	userService := impl.NewUserServiceImpl(userRepo)
	roleService := impl.NewRoleServiceImpl(roleRepo)
	authService := impl.NewAuthServiceImpl(authRepo, userRepo, userService, otpService, smtpService)
	bankService := impl.NewBankServiceImpl(bankRepo)
	bankAccountService := impl.NewBankAccountServiceImpl(bankAccountRepo, bankService)
	cardService := impl.NewCardServiceImpl(cardRepo, userService)
	walletService := impl.NewWalletServiceImpl(walletRepo, userService)
	walletTransactionService := impl.NewWalletTransactionServiceImpl(walletTransactionRepo, walletService)
	p2pTransferService := impl.NewP2PTransferServiceImpl(p2pTransferRepo, walletTransactionRepo, walletService, userService)

	// HANDLERS
	authHandler := handlers.NewAuthHandler(authService, userService, otpService)
	bankHandler := handlers.NewBankHandler(bankService)
	bankAccountHandler := handlers.NewBankAccountHandler(bankAccountService)
	cardHandler := handlers.NewCardHandler(cardService)
	walletHandler := handlers.NewWalletHandler(walletService)
	walletTransactionHandler := handlers.NewWalletTransactionHandler(walletTransactionService)
	p2pTransferHandler := handlers.NewP2PTransferHandler(p2pTransferService)

	return &Container{
		AuthService: authService,
		RoleService: roleService,
		SmtpService: smtpService,
		OtpService:  otpService,
		UserService: userService,

		AuthHandler:              authHandler,
		BankHandler:              bankHandler,
		BankAccountHandler:       bankAccountHandler,
		CardHandler:              cardHandler,
		WalletHandler:            walletHandler,
		WalletTransactionHandler: walletTransactionHandler,
		P2PTransferHandler:       p2pTransferHandler,
	}
}
