package server

import (
	"fmt"

	"github.com/asfung/elara/config"
	"github.com/asfung/elara/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	return &echoServer{
		app:  echoApp,
		db:   db,
		conf: conf,
	}
}

func (s *echoServer) Start() {
	s.app.Pre(middleware.RemoveTrailingSlash())
	s.app.Use(middleware.Recover())
	s.app.Use(middleware.Logger())

	api := s.app.Group("/api/v1")

	s.initializeHelloHttpHandler(api)

	api.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{"message": "Ok!"})
	})

	serverUrl := fmt.Sprintf(":%d", s.conf.Server.Port)
	s.app.Logger.Fatal(s.app.Start(serverUrl))
}

func (s *echoServer) initializeHelloHttpHandler(api *echo.Group) {
	api.GET("/hello", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{"message": "Hello!"})
	})
}
