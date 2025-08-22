package main

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/asfung/elara/config"
	"github.com/asfung/elara/database"
	"github.com/asfung/elara/di"
	"github.com/asfung/elara/server"
)

func main() {
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	conf := config.GetConfig()
	db := database.NewPostgresDatabase(conf)
	container := di.NewContainer(conf, db)
	server.NewEchoServer(conf, db, container).Start()
}
