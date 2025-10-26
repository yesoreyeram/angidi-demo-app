.PHONY: help build test lint fmt clean run install-tools deps test-coverage

help: ## Show this help
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

install-tools: ## Install development tools
	@echo "Installing development tools..."
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
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
	@golangci-lint run ./...
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

.DEFAULT_GOAL := help
