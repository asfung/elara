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
		fmt.Println("Usage: go run internal/migrations/wallet/walletMigrate.go [create|delete|refresh]")
		os.Exit(1)
	}

	action := os.Args[1]
	conf := config.GetConfig()
	db := database.NewPostgresDatabase(conf)

	switch action {
	case "create":
		createWalletTable(db)
	case "delete":
		deleteWalletTable(db)
	case "refresh":
		refreshWalletTable(db)
	default:
		fmt.Println("Unknown action:", action)
	}

}

func createWalletTable(db database.Database) {
	db.GetDb().Migrator().CreateTable(&entities.Bank{})
	db.GetDb().Migrator().CreateTable(&entities.Wallet{})
	db.GetDb().Migrator().CreateTable(&entities.BankAccount{})
	db.GetDb().Migrator().CreateTable(&entities.Card{})
	db.GetDb().Migrator().CreateTable(&entities.WalletTransaction{})
	db.GetDb().Migrator().CreateTable(&entities.P2pTransfer{})
	db.GetDb().Migrator().CreateTable(&entities.Biller{})
	db.GetDb().Migrator().CreateTable(&entities.BillPayment{})
	db.GetDb().Migrator().CreateTable(&entities.TransactionLimit{})
	db.GetDb().Migrator().CreateTable(&entities.RiskFlag{})
	db.GetDb().Migrator().CreateTable(&entities.Category{})
	db.GetDb().Migrator().CreateTable(&entities.Expense{})
	fmt.Println("wallet tables created")
}

func deleteWalletTable(db database.Database) {
	db.GetDb().Migrator().DropTable(&entities.Bank{})
	db.GetDb().Migrator().DropTable(&entities.Wallet{})
	db.GetDb().Migrator().DropTable(&entities.BankAccount{})
	db.GetDb().Migrator().DropTable(&entities.Card{})
	db.GetDb().Migrator().DropTable(&entities.WalletTransaction{})
	db.GetDb().Migrator().DropTable(&entities.P2pTransfer{})
	db.GetDb().Migrator().DropTable(&entities.Biller{})
	db.GetDb().Migrator().DropTable(&entities.BillPayment{})
	db.GetDb().Migrator().DropTable(&entities.TransactionLimit{})
	db.GetDb().Migrator().DropTable(&entities.RiskFlag{})
	db.GetDb().Migrator().DropTable(&entities.Category{})
	db.GetDb().Migrator().DropTable(&entities.Expense{})
	fmt.Println("wallet tables deleted")
}

func refreshWalletTable(db database.Database) {
	createWalletTable(db)
	deleteWalletTable(db)
	fmt.Println("wallet tables refreshed")
}
