package server

import (
	"fmt"

	"github.com/asfung/elara/config"
	"github.com/asfung/elara/database"
	"github.com/asfung/elara/internal/handlers"
	"github.com/asfung/elara/internal/repositories"
	"github.com/asfung/elara/internal/services/impl"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type echoServer struct {
	app  *echo.Echo
	db   database.Database
	conf *config.Config
}

func NewEchoServer(conf *config.Config, db database.Database) Server {
	echoApp := echo.New()
	echoApp.Logger.SetLevel(log.DEBUG)
	RegisterValidator(echoApp)

	return &echoServer{
		app:  echoApp,
		db:   db,
		conf: conf,
	}
}

func (s *echoServer) Start() {
	BaseMiddleware(s.app)

	api := s.app.Group("/api/v1")

	s.initializeHelloHttpHandler(api)
	s.registerAuthRoutes(api)

	api.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{"message": "Ok!"})
	})

	serverUrl := fmt.Sprintf(":%d", s.conf.Server.Port)
	s.app.Logger.Fatal(s.app.Start(serverUrl))
}

func (s *echoServer) registerAuthRoutes(api *echo.Group) {
	// authHandler := s.container.AuthHandler

	userRepo := repositories.NewUserPostgresRepository(s.db)
	authRepo := repositories.NewAuthPostgresRepository(s.db)
	userService := impl.NewUserServiceImpl(userRepo)
	authService := impl.NewAuthServiceImpl(authRepo, userRepo, userService)
	authHandler := handlers.NewAuthHandler(authService)

	// register auth routes
	authGroup := api.Group("/auth")
	authGroup.POST("/login", authHandler.Login)
	authGroup.POST("/register", authHandler.Register)
	authGroup.POST("/refresh", authHandler.RefreshToken).Name = "auth.refresh.token"
	authGroup.POST("/logout", authHandler.Logout)
}

func (s *echoServer) initializeHelloHttpHandler(e *echo.Group) {
	userRepo := repositories.NewUserPostgresRepository(s.db)
	authRepo := repositories.NewAuthPostgresRepository(s.db)
	userService := impl.NewUserServiceImpl(userRepo)
	authService := impl.NewAuthServiceImpl(authRepo, userRepo, userService)
	authHandler := handlers.NewAuthHandler(authService)

	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{"message": "Hello!"})
	})
	e.GET("/authenticated", authHandler.Authenticated, AuthMiddleware(authService))
}
