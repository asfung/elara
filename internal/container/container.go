package container

import (
	"github.com/asfung/elara/database"
	"github.com/asfung/elara/internal/handlers"
	"github.com/asfung/elara/internal/repositories"
	"github.com/asfung/elara/internal/services"
	"github.com/asfung/elara/internal/services/impl"
)

type Container struct {
	AuthService        services.AuthService
	AuthHandler        *handlers.AuthHandler
	BankHandler        *handlers.BankHandler
	BankAccountHandler *handlers.BankAccountHandler
}

func NewContainer(db database.Database) *Container {
	// REPOSITORIES
	userRepo := repositories.NewUserPostgresRepository(db)
	authRepo := repositories.NewAuthPostgresRepository(db)
	bankRepo := repositories.NewBankPostgresRepository(db)
	bankAccountRepo := repositories.NewBankAccountPostgresRepository(db)

	// SERVICES
	userService := impl.NewUserServiceImpl(userRepo)
	authService := impl.NewAuthServiceImpl(authRepo, userRepo, userService)
	bankService := impl.NewBankServiceImpl(bankRepo)
	bankAccountService := impl.NewBankAccountServiceImpl(bankAccountRepo, bankService)

	// HANDLERS
	authHandler := handlers.NewAuthHandler(authService)
	bankHandler := handlers.NewBankHandler(bankService)
	bankAccountHandler := handlers.NewBankAccountHandler(bankAccountService)

	return &Container{
		AuthService:        authService,
		AuthHandler:        authHandler,
		BankHandler:        bankHandler,
		BankAccountHandler: bankAccountHandler,
	}
}
