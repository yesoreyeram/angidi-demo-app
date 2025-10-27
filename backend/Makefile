.PHONY: help build test lint fmt clean run install-tools deps test-coverage migrate-up migrate-down migrate-create db-setup

help: ## Show this help
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

install-tools: ## Install development tools
	@echo "Installing development tools..."
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	@echo "Tools installed successfully"

deps: ## Download dependencies
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy
	@echo "Dependencies downloaded"

build: ## Build all binaries
	@echo "Building API server..."
	@go build -o bin/api ./cmd/api
	@echo "Build complete"

test: ## Run all tests
	@echo "Running tests..."
	@go test -v ./...
	@echo "Tests complete"

test-coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

test-integration: ## Run integration tests
	@echo "Running integration tests..."
	@go test -v -tags=integration ./tests/integration/...
	@echo "Integration tests complete"

lint: ## Run linters
	@echo "Running linters..."
	@$(shell go env GOPATH)/bin/golangci-lint run ./...
	@echo "Linting complete"

fmt: ## Format code
	@echo "Formatting code..."
	@gofmt -s -w .
	@goimports -w .
	@echo "Formatting complete"

run: ## Run development server
	@echo "Starting development server..."
	@go run ./cmd/api

clean: ## Clean build artifacts
	@echo "Cleaning build artifacts..."
	@rm -rf bin/
	@rm -f coverage.out coverage.html
	@echo "Clean complete"

# Database commands
DB_URL ?= postgres://postgres:postgres@localhost:5432/angidi_dev?sslmode=disable

migrate-up: ## Run all pending migrations
	@echo "Running migrations..."
	@migrate -path ./migrations -database "$(DB_URL)" up
	@echo "Migrations complete"

migrate-down: ## Rollback last migration
	@echo "Rolling back migration..."
	@migrate -path ./migrations -database "$(DB_URL)" down 1
	@echo "Rollback complete"

migrate-create: ## Create new migration (usage: make migrate-create NAME=migration_name)
	@if [ -z "$(NAME)" ]; then \
		echo "Error: NAME is required. Usage: make migrate-create NAME=migration_name"; \
		exit 1; \
	fi
	@echo "Creating migration $(NAME)..."
	@migrate create -ext sql -dir ./migrations -seq $(NAME)
	@echo "Migration files created"

db-setup: ## Setup database (create database and run migrations)
	@echo "Setting up database..."
	@createdb angidi_dev 2>/dev/null || echo "Database angidi_dev already exists"
	@createdb angidi_test 2>/dev/null || echo "Database angidi_test already exists"
	@$(MAKE) migrate-up
	@echo "Database setup complete"

.DEFAULT_GOAL := help
