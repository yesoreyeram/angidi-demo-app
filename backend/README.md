# Angidi Backend

Go-based backend API for the Angidi e-commerce platform.

## Overview

This is the backend service for Angidi, built with Go 1.21+. It provides RESTful APIs for the e-commerce platform with features like user management, product catalog, shopping cart, and order processing.

## Prerequisites

- Go 1.21 or higher
- Make (for build automation)
- golangci-lint (for linting)

## Quick Start

### Installation

```bash
# Clone the repository
git clone https://github.com/yesoreyeram/angidi-demo-app.git
cd angidi-demo-app/backend

# Install development tools
make install-tools

# Download dependencies
make deps
```

### Development

```bash
# Run the development server
make run

# The server will start on http://localhost:8080
# Health check endpoint: http://localhost:8080/health
```

### Building

```bash
# Build the API server binary
make build

# Binary will be created at bin/api
./bin/api
```

### Testing

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run integration tests
make test-integration
```

### Code Quality

```bash
# Run linters
make lint

# Format code
make fmt
```

## Project Structure

```
backend/
├── cmd/
│   └── api/              # API server entry point
│       └── main.go
├── internal/             # Private application code
│   ├── user/             # User domain
│   ├── product/          # Product domain
│   ├── cart/             # Cart domain
│   ├── order/            # Order domain
│   └── common/           # Shared internal code
│       ├── middleware/   # HTTP middleware
│       ├── validator/    # Input validation
│       └── errors/       # Error handling
├── pkg/                  # Public, reusable packages
│   ├── logger/           # Structured logging
│   ├── config/           # Configuration management
│   ├── database/         # Database utilities
│   └── http/             # HTTP utilities
├── api/                  # API specifications
│   └── openapi/          # OpenAPI/Swagger specs
├── tests/                # Integration and E2E tests
│   ├── integration/
│   └── e2e/
├── scripts/              # Build and deployment scripts
├── configs/              # Configuration files
├── Makefile              # Build automation
└── README.md             # This file
```

## Configuration

Configuration is managed through YAML files in the `configs/` directory and can be overridden with environment variables.

### Configuration Files

- `configs/local.yaml` - Local development settings
- `configs/dev.yaml` - Development environment (to be added)
- `configs/prod.yaml` - Production environment (to be added)

### Environment Variables

- `CONFIG_PATH` - Path to configuration file (default: configs/local.yaml)
- `SERVER_HOST` - Server host (default: localhost)
- `SERVER_PORT` - Server port (default: 8080)
- `LOG_LEVEL` - Log level: debug, info, warn, error (default: info)

### Example

```bash
# Run with custom port
SERVER_PORT=9090 make run

# Run with debug logging
LOG_LEVEL=debug make run
```

## API Endpoints

### Health Check

```bash
GET /health
```

Returns server health status.

**Response:**
```json
{
  "status": "healthy",
  "timestamp": "2025-10-26T23:00:00Z"
}
```

### Welcome

```bash
GET /
```

Returns welcome message and version.

**Response:**
```json
{
  "message": "Welcome to Angidi API",
  "version": "0.1.0"
}
```

## Development Workflow

### Adding a New Feature

1. Create feature branch
2. Write tests first (TDD)
3. Implement feature
4. Run tests and linters
5. Submit pull request

### Code Style

- Follow [Effective Go](https://golang.org/doc/effective_go.html)
- Use `gofmt` for formatting
- Write documentation comments for exported items
- Keep functions small and focused

### Testing

- Write unit tests for all packages
- Aim for >80% code coverage
- Use table-driven tests when applicable
- Mock external dependencies

## Makefile Targets

Run `make help` to see all available targets:

```
Available targets:
  help                 Show this help
  install-tools        Install development tools
  deps                 Download dependencies
  build                Build all binaries
  test                 Run all tests
  test-coverage        Run tests with coverage
  test-integration     Run integration tests
  lint                 Run linters
  fmt                  Format code
  run                  Run development server
  clean                Clean build artifacts
```

## Troubleshooting

### Go Module Download Fails

```bash
# Set Go proxy
export GOPROXY=https://proxy.golang.org,direct

# Clear module cache
go clean -modcache

# Retry
make deps
```

### Linter Fails

```bash
# Update golangci-lint
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run with verbose output
golangci-lint run -v
```

### Port Already in Use

```bash
# Change the port
SERVER_PORT=9090 make run
```

## Dependencies

Main dependencies:
- `go.uber.org/zap` - Structured logging
- `gopkg.in/yaml.v3` - YAML parsing

Development dependencies:
- `golangci-lint` - Linting

## License

MIT License - see LICENSE file for details.

## Contributing

See DEVELOPMENT.md for development guidelines and contribution process.

---

**Current Phase**: Phase 1 - Repository Scaffolding  
**Status**: ✅ Complete  
**Last Updated**: 2025-10-26
