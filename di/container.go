package di

import (
	"github.com/asfung/elara/config"
	"github.com/asfung/elara/database"
	"github.com/asfung/elara/internal/handlers"
	"github.com/asfung/elara/internal/repositories"
	"github.com/asfung/elara/internal/services/impl"
)

type Container struct {
	AuthHandler *handlers.AuthHandler
}

func NewContainer(conf *config.Config, db database.Database) *Container {
	// repositories
	userRepo := repositories.NewUserPostgresRepository(db)
	authRepo := repositories.NewAuthPostgresRepository(db)

	// services
	userService := impl.NewUserServiceImpl(userRepo)
	authService := impl.NewAuthServiceImpl(authRepo, userRepo, userService)

	// handlers
	authHandler := handlers.NewAuthHandler(authService)

	return &Container{
		AuthHandler: authHandler,
	}
}
