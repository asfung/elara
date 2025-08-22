MIGRATION_DIR=internal/migrations

start:
	air

migrate-user-create:
	@go run $(MIGRATION_DIR)/userMigrate.go create

migrate-user-delete:
	@go run $(MIGRATION_DIR)/userMigrate.go delete

migrate-user-refresh:
	@go run $(MIGRATION_DIR)/userMigrate.go refresh
