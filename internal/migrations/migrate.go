package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/asfung/elara/config"
	"github.com/asfung/elara/database"
	"github.com/asfung/elara/database/seeders"
	"github.com/asfung/elara/internal/entities"
	"gorm.io/gorm"
)

const (
	ENTITY = "entity"
)

var entityRegistry = map[string][]interface{}{
	"bank_entity": {
		&entities.Bank{},
		&entities.BankAccount{},
	},
	"bill_entity": {
		&entities.Biller{},
		&entities.BillPayment{},
	},
	"budget_entity": {
		&entities.Category{},
		&entities.Expense{},
	},
	"card_entity": {
		&entities.Card{},
	},
	"otp_entity": {
		&entities.OTP{},
	},
	"p2p_entity": {
		&entities.P2pTransfer{},
	},
	"security_entity": {
		&entities.TransactionLimit{},
		&entities.RiskFlag{},
	},
	"user_entity": {
		&entities.User{},
	},
	"role_entity": {
		&entities.Role{},
	},
	"wallet_entity": {
		&entities.Wallet{},
		&entities.WalletTransaction{},
	},
	"portfolio_entity": {
		&entities.Asset{},
		&entities.Portfolio{},
		&entities.PortfolioAsset{},
	},
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run migrate.go [entity_file_name] [create|delete|refresh]")
		os.Exit(1)
	}

	// entityName := strings.TrimSuffix(os.Args[1]+ENTITY, ".go")
	entityName := strings.TrimSuffix(os.Args[1], ".go")
	action := os.Args[2]

	models, ok := entityRegistry[entityName]
	if !ok {
		fmt.Printf("Unknown entity: %s\n", entityName)
		os.Exit(1)
	}

	conf := config.GetConfig()
	db := database.NewPostgresDatabase(conf)

	switch action {
	case "create":
		createTables(db.GetDb(), models, entityName)
		if entityName == "role_entity" {
			seeders.SeedRoleEntity(db.GetDb())
		}
	case "delete":
		deleteTables(db.GetDb(), models, entityName)
	case "refresh":
		refreshTables(db.GetDb(), models, entityName)
		if entityName == "role_entity" {
			seeders.SeedRoleEntity(db.GetDb())
		}
	default:
		fmt.Println("Unknown action:", action)
	}
}

func createTables(db *gorm.DB, models []interface{}, name string) {
	for _, m := range models {
		if err := db.Migrator().CreateTable(m); err != nil {
			fmt.Println("Error creating table:", err)
			return
		}
	}
	fmt.Printf("Tables for %s created\n", name)
}

func deleteTables(db *gorm.DB, models []interface{}, name string) {
	for _, m := range models {
		if err := db.Migrator().DropTable(m); err != nil {
			fmt.Println("Error deleting table:", err)
			return
		}
	}
	fmt.Printf("Tables for %s deleted\n", name)
}

func refreshTables(db *gorm.DB, models []interface{}, name string) {
	for _, m := range models {
		if err := db.Migrator().DropTable(m); err != nil {
			fmt.Println("Error dropping table:", err)
			return
		}
	}
	for _, m := range models {
		if err := db.Migrator().CreateTable(m); err != nil {
			fmt.Println("Error creating table:", err)
			return
		}
	}
	fmt.Printf("Tables for %s refreshed\n", name)
}
