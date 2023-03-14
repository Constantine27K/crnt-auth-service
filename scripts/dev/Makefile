DEV_PROJECT := auth-dev
DEV_COMPOSE_FILE := docker-compose.yml
MIGRATION_DIR := ../migrations/

.PHONY: up-local-db
up-local-db: down-local-db
	#${INFO} "Starting shard database"
	@docker-compose -p $(DEV_PROJECT) -f $(DEV_COMPOSE_FILE) up -d postgres
	@docker-compose -p $(DEV_PROJECT) -f $(DEV_COMPOSE_FILE) run --rm pg-up-checker
	goose -dir ${MIGRATION_DIR} postgres "user=auth-user dbname=auth-db password=auth-password host=localhost port=5433 sslmode=disable" up

.PHONY: down-local-db
down-local-db:
	@docker-compose -p $(DEV_PROJECT) -f $(DEV_COMPOSE_FILE) down

.PHONY: migration-apply
migration-apply:
	goose -dir ${MIGRATION_DIR} postgres "user=auth-user dbname=auth-db password=auth-password host=localhost port=5433 sslmode=disable" up