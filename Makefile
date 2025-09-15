MIGRATION_DIR=internal/migrations

start:
	air

migrate-create:
	@go run ${MIGRATION_DIR}/migrate.go create
migrate-delete:
	@go run ${MIGRATION_DIR}/migrate.go delete
migrate-refresh:
	@go run ${MIGRATION_DIR}/migrate.go refresh

migrate-user-create:
	@go run $(MIGRATION_DIR)/user/user_migrate.go create

migrate-user-delete:
	@go run $(MIGRATION_DIR)/user/user_migrate.go delete

migrate-user-refresh:
	@go run $(MIGRATION_DIR)/user/user_migrate.go refresh

migrate-wallet-create:
	@go run $(MIGRATION_DIR)/wallet/wallet_migrate.go create

migrate-wallet-delete:
	@go run $(MIGRATION_DIR)/wallet/wallet_migrate.go delete

migrate-wallet-refresh:
	@go run $(MIGRATION_DIR)/wallet/wallet_migrate.go refresh
