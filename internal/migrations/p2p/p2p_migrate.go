package main

import (
	"fmt"
	"os"

	"github.com/asfung/elara/config"
	"github.com/asfung/elara/database"
	"github.com/asfung/elara/internal/entities"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run internal/migrations/p2p/p2p_migrate.go [create|delete|refresh]")
		os.Exit(1)
	}

	action := os.Args[1]
	conf := config.GetConfig()
	db := database.NewPostgresDatabase(conf)

	switch action {
	case "create":
		createP2PTable(db)
	case "delete":
		deleteP2PTable(db)
	case "refresh":
		refreshP2PTable(db)
	default:
		fmt.Println("Unknown action:", action)
	}

}

func createP2PTable(db database.Database) {
	db.GetDb().Migrator().CreateTable(&entities.P2pTransfer{})
	fmt.Println("p2p tables created")
}

func deleteP2PTable(db database.Database) {
	db.GetDb().Migrator().DropTable(&entities.P2pTransfer{})
	fmt.Println("p2p tables deleted")
}

func refreshP2PTable(db database.Database) {
	createP2PTable(db)
	deleteP2PTable(db)
	fmt.Println("p2p tables refreshed")
}
