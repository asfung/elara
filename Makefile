MIGRATION_DIR=internal/migrations

start:
	air

migrate-user-create:
	@go run $(MIGRATION_DIR)/user/userMigrate.go create

migrate-user-delete:
	@go run $(MIGRATION_DIR)/user/userMigrate.go delete

migrate-user-refresh:
	@go run $(MIGRATION_DIR)/user/userMigrate.go refresh

migrate-wallet-create:
	@go run $(MIGRATION_DIR)/wallet/walletMigrate.go create

migrate-wallet-delete:
	@go run $(MIGRATION_DIR)/wallet/walletMigrate.go delete

migrate-wallet-refresh:
	@go run $(MIGRATION_DIR)/wallet/walletMigrate.go refresh
