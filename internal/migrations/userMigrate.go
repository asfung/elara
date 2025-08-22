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
		fmt.Println("Usage: go run internal/migrations/userMigrate.go [create|delete|refresh]")
		os.Exit(1)
	}

	action := os.Args[1]
	conf := config.GetConfig()
	db := database.NewPostgresDatabase(conf)

	switch action {
	case "create":
		createUserTable(db)
	case "delete":
		deleteUserTable(db)
	case "refresh":
		refreshUserTable(db)
	default:
		fmt.Println("Unknown action:", action)
	}

}

func createUserTable(db database.Database) {
	db.GetDb().Migrator().CreateTable(&entities.User{})
	fmt.Println("User table created")
}

func deleteUserTable(db database.Database) {
	db.GetDb().Migrator().DropTable(&entities.User{})
	fmt.Println("User table deleted")
}

func refreshUserTable(db database.Database) {
	db.GetDb().Migrator().DropTable(&entities.User{})
	db.GetDb().Migrator().CreateTable(&entities.User{})
	fmt.Println("User table refreshed")
}
