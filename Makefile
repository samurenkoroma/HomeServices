.PHONY: help migrate-up migrate-down migrate-create migrate-force migrate-status

DB_DSN ?= postgres://tiun:tiun@lab.note:5432/tiun?sslmode=disable

help:
	@echo "Commands:"
	@echo "  make migrate-up MODULE=farm           - Run up migrations"
	@echo "  make migrate-down MODULE=farm STEPS=1 - Rollback"
	@echo "  make migrate-create NAME=xxx MODULE=farm - Create migration"
	@echo "  make migrate-force VERSION=1 MODULE=farm - Force version"
	@echo "  make migrate-status MODULE=farm       - Show status"

migrate-up:
	@if [ -z "$(MODULE)" ]; then \
		echo "Usage: make migrate-up MODULE=farm"; \
		exit 1; \
	fi
	 go run cmd/migrate/main.go -module=$(MODULE) -direction=up -dsn="$(DB_DSN)" -verbose=true

migrate-down:
	@if [ -z "$(MODULE)" ]; then \
		echo "Usage: make migrate-down MODULE=farm STEPS=1"; \
		exit 1; \
	fi
	go run cmd/migrate/main.go \
		-module=$(MODULE) \
		-direction=down \
		-steps=$(STEPS) \
		-dsn="$(DB_DSN)" \
		-verbose=true

migrate-create:
	@if [ -z "$(NAME)" ] || [ -z "$(MODULE)" ]; then \
		echo "Usage: make migrate-create NAME=migration_name MODULE=farm"; \
		exit 1; \
	fi
	@mkdir -p ./internal/modules/$(MODULE)/infrastructure/persistence/postgres/migrations
	@timestamp=$$(date +%Y%m%d%H%M%S); \
	up_file="./internal/modules/$(MODULE)/infrastructure/persistence/postgres/migrations/$${timestamp}_$(NAME).up.sql"; \
	down_file="./internal/modules/$(MODULE)/infrastructure/persistence/postgres/migrations/$${timestamp}_$(NAME).down.sql"; \
	echo "-- Up migration: $(NAME)" > $$up_file; \
	echo "-- Created: $$(date)" >> $$up_file; \
	echo "" >> $$up_file; \
	echo "CREATE TABLE IF NOT EXISTS $(NAME) (" >> $$up_file; \
	echo "    id UUID PRIMARY KEY DEFAULT gen_random_uuid()" >> $$up_file; \
	echo ");" >> $$up_file; \
	echo "" >> $$up_file; \
	echo "-- Down migration: $(NAME)" > $$down_file; \
	echo "DROP TABLE IF EXISTS $(NAME);" >> $$down_file; \
	echo "Created: $$up_file"; \
	echo "Created: $$down_file"

migrate-force:
	@if [ -z "$(VERSION)" ] || [ -z "$(MODULE)" ]; then \
		echo "Usage: make migrate-force VERSION=1 MODULE=farm"; \
		exit 1; \
	fi
	cd migrations && go run cmd/migrate/main.go \
		-module=$(MODULE) \
		-force=$(VERSION) \
		-dsn="$(DB_DSN)"

migrate-status:
	@if [ -z "$(MODULE)" ]; then \
		echo "Usage: make migrate-status MODULE=farm"; \
		exit 1; \
	fi
	@psql "$(DB_DSN)" -c "SELECT version, applied_at FROM schema_migrations WHERE module = '$(MODULE)' ORDER BY version;"