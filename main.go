package main

import (
	_ "net/http/pprof"

	"github.com/asfung/elara/config"
	"github.com/asfung/elara/database"
	"github.com/asfung/elara/server"
	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Warn(".env not loaded", err)
	}
	conf := config.GetConfig()
	db := database.NewPostgresDatabase(conf)
	server.NewEchoServer(conf, db).Start()
}
