package server

import (
	"fmt"

	"github.com/asfung/elara/config"
	"github.com/asfung/elara/database"
	"github.com/asfung/elara/internal/container"
	"github.com/asfung/elara/internal/handlers"
	"github.com/asfung/elara/internal/oauth"
	"github.com/asfung/elara/internal/services"
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

	oauth.InitProviders()

	api := s.app.Group("/api/v1")
	// dependencies
	c := container.NewContainer(s.db)

	s.initializeHelloHttpHandler(api)
	s.registerAuthRoutes(api, c.AuthHandler, &c.AuthService, AuthMiddleware(c.AuthService, c.RoleService))
	s.registerBankRoutes(api, c.BankHandler, AuthMiddleware(c.AuthService, c.RoleService))
	s.registerBankAccountRoutes(api, c.BankAccountHandler, AuthMiddleware(c.AuthService, c.RoleService))
	s.registerCardRoutes(api, c.CardHandler, AuthMiddleware(c.AuthService, c.RoleService))
	s.registerWalletRoutes(api, c.WalletHandler, AuthMiddleware(c.AuthService, c.RoleService))
	s.registerWalletTransactionRoutes(api, c.WalletTransactionHandler, AuthMiddleware(c.AuthService, c.RoleService))
	s.registerP2PTransferRoutes(api, c.P2PTransferHandler, AuthMiddleware(c.AuthService, c.RoleService))
	s.registerAssetRoutes(api, c.Assetandler, AuthMiddleware(c.AuthService, c.RoleService))
	s.registerPortfolioRoutes(api, c.PortfolioHandler, c.PortfolioAssetHandler, AuthMiddleware(c.AuthService, c.RoleService))

	api.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{"message": "Ok!"})
	})

	serverUrl := fmt.Sprintf(":%d", s.conf.Server.Port)
	s.app.Logger.Fatal(s.app.Start(serverUrl))
}

// ===============
// ROUTES REGISTER
// ===============
func (s *echoServer) registerAuthRoutes(api *echo.Group, authHandler *handlers.AuthHandler, authService *services.AuthService, authMiddleware echo.MiddlewareFunc) {
	// register auth routes
	authGroup := api.Group("/auth")
	authGroup.POST("/login", authHandler.Login)
	authGroup.POST("/register", authHandler.Register)
	authGroup.GET("/refresh", authHandler.RefreshToken, authMiddleware).Name = "auth.refresh.token"
	authGroup.POST("/logout", authHandler.Logout, authMiddleware, authMiddleware)

	// oauth
	oauthHandler := handlers.NewOAuthHandler(*authService, "")
	oauthGroup := authGroup.Group("/oauth")
	oauthGroup.GET("/:provider", oauthHandler.BeginAuth)
	oauthGroup.GET("/:provider/callback", oauthHandler.Callback)

	// new auth mechanism
	authGroup.POST("/authorize/continue", authHandler.CheckEmail)
	authGroup.POST("/user/register", authHandler.CreateAccount)
	authGroup.POST("/email-otp/validate", authHandler.VerifyOTP)
	authGroup.POST("/password/verify", authHandler.VerifyPassword)
}

func (s *echoServer) initializeHelloHttpHandler(e *echo.Group) {
	container := container.NewContainer(s.db)
	authHandler := handlers.NewAuthHandler(container.AuthService, container.UserService, container.OtpService)

	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{"message": "Hello!"})
	})
	e.GET("/authenticated", authHandler.Authenticated, AuthMiddleware(container.AuthService, container.RoleService))
}

func (s *echoServer) registerBankRoutes(e *echo.Group, bankHandler *handlers.BankHandler, authMiddleware echo.MiddlewareFunc) {
	bankGroup := e.Group("/bank", authMiddleware)
	bankGroup.POST("", bankHandler.CreateBank)
	bankGroup.GET("", bankHandler.GetBanks)
	bankGroup.PUT("/:id", bankHandler.UpdateBank)
	bankGroup.GET("/:id", bankHandler.GetById)
	bankGroup.DELETE("/:id", bankHandler.DeleteBank)
}

func (s *echoServer) registerBankAccountRoutes(e *echo.Group, bankAccountHandler *handlers.BankAccountHandler, authMiddleware echo.MiddlewareFunc) {
	bankGroup := e.Group("/bank-account", authMiddleware)
	bankGroup.POST("", bankAccountHandler.CreateBankAccount)
	bankGroup.PUT("/:id", bankAccountHandler.UpdateBankAccount)
	bankGroup.GET("/:id", bankAccountHandler.GetById)
	bankGroup.DELETE("/:id", bankAccountHandler.DeleteBankAccount)
}

func (s *echoServer) registerCardRoutes(e *echo.Group, cardHandler *handlers.CardHandler, authMiddleware echo.MiddlewareFunc) {
	cardGroup := e.Group("/card", authMiddleware)
	cardGroup.POST("", cardHandler.CreateCard)
	cardGroup.PUT("/:id", cardHandler.UpdateCard)
	cardGroup.GET("/:id", cardHandler.GetCardById)
	cardGroup.DELETE("/:id", cardHandler.DeleteCard)
}

func (s *echoServer) registerWalletRoutes(e *echo.Group, walletHandler *handlers.WalletHandler, authMiddleware echo.MiddlewareFunc) {
	walletGroup := e.Group("/wallet", authMiddleware)
	walletGroup.POST("", walletHandler.CreateWallet)
	walletGroup.PUT("/:id", walletHandler.UpdateWallet)
	walletGroup.PUT("/balance/:id", walletHandler.UpdateWalletBalance)
	walletGroup.GET("/:id", walletHandler.GetWalletById)
	walletGroup.GET("/user/:userId", walletHandler.GetWalletByUserId)
	walletGroup.DELETE("/:id", walletHandler.DeleteWallet)
}

func (s *echoServer) registerWalletTransactionRoutes(e *echo.Group, walletTransactionHandler *handlers.WalletTransactionHandler, authMiddleware echo.MiddlewareFunc) {
	walletTransactionGroup := e.Group("/wallet-transaction", authMiddleware)
	walletTransactionGroup.POST("", walletTransactionHandler.CreateWalletTransaction)
	walletTransactionGroup.PUT("/:id", walletTransactionHandler.UpdateWalletTransaction)
	walletTransactionGroup.GET("/:id", walletTransactionHandler.GetWalletTransactionById)
	walletTransactionGroup.GET("/user/:userId", walletTransactionHandler.GetWalletTransactionByUserIdPaginated) //deprecated
	walletTransactionGroup.GET("/wallet/:walletId", walletTransactionHandler.GetWalletTransactionByWalletIdPaginated)
	walletTransactionGroup.DELETE("/:id", walletTransactionHandler.DeleteWalletTransaction)
}

func (s *echoServer) registerP2PTransferRoutes(e *echo.Group, p2pTransferHandler *handlers.P2PTransferHandler, authMiddleware echo.MiddlewareFunc) {
	p2pTransferGroup := e.Group("/p2p-transfer", authMiddleware)
	p2pTransferGroup.POST("", p2pTransferHandler.CreateP2PTransfer)
	p2pTransferGroup.PUT("/:id", p2pTransferHandler.Update2PTransfer)
	p2pTransferGroup.GET("/:id", p2pTransferHandler.GetP2PTransferById)
	p2pTransferGroup.DELETE("/:id", p2pTransferHandler.DeleteP2PTransfer)
}

func (s *echoServer) registerAssetRoutes(e *echo.Group, assetHandler *handlers.AssetHandler, authMiddleware echo.MiddlewareFunc) {
	assetGroup := e.Group("/asset", authMiddleware)
	assetGroup.GET("", assetHandler.CreateAsset)
	assetGroup.PUT("/:id", assetHandler.UpdateAsset)
	assetGroup.GET("/:id", assetHandler.GetAssetById)
	assetGroup.DELETE("", assetHandler.DeleteAsset)
}

func (s *echoServer) registerPortfolioRoutes(e *echo.Group, portfolioHandler *handlers.PortfolioHandler, portfolioAssetHandler *handlers.PortfolioAssetHandler, authMiddleware echo.MiddlewareFunc) {
	portfolioGroup := e.Group("/portfolio", authMiddleware)
	portfolioGroup.GET("", portfolioHandler.CreatePortfolio)
	portfolioGroup.PUT("/:id", portfolioHandler.UpdatePortfolio)
	portfolioGroup.GET("/:id", portfolioHandler.GetPortfolioById)
	portfolioGroup.DELETE("", portfolioHandler.DeletePortfolio)

	// portfolio asset
	portfolioAssetGroup := portfolioGroup.Group("/asset")
	portfolioAssetGroup.GET("", portfolioAssetHandler.CreatePortfolioAsset)
	portfolioAssetGroup.PUT("/:id", portfolioAssetHandler.UpdatePortfolioAsset)
	portfolioAssetGroup.GET("/:id", portfolioAssetHandler.GetPortfolioAssetById)
	portfolioAssetGroup.DELETE("", portfolioAssetHandler.DeletePortfolioAsset)

}

// func (s *echoServer) registerPortfolioAssetRoutes(e *echo.Group, portfolioAssetHandler *handlers.PortfolioAssetHandler, authMiddleware echo.MiddlewareFunc) {
// 	portfolioAssetGroup := e.Group("/portfolio-asset", authMiddleware)
// 	portfolioGroup.GET("", portfolioHandler.CreatePortfolio)
// 	portfolioGroup.PUT("/:id", portfolioHandler.UpdatePortfolio)
// 	portfolioGroup.GET("/:id", portfolioHandler.GetPortfolioById)
// 	portfolioGroup.DELETE("", portfolioHandler.DeletePortfolio)
// }
