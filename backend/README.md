# Angidi Backend

Go-based backend API for the Angidi e-commerce platform.

## Overview

This is the backend service for Angidi, built with Go 1.21+. It provides RESTful APIs for the e-commerce platform with features like user management, product catalog, shopping cart, and order processing.

## Prerequisites

- Go 1.21 or higher
- PostgreSQL 15+ (for data persistence)
- Make (for build automation)
- golangci-lint (for linting)
- golang-migrate (for database migrations)

## Quick Start

### Database Setup

```bash
# Install PostgreSQL (macOS)
brew install postgresql@15
brew services start postgresql@15

# Create databases
createdb angidi_dev
createdb angidi_test

# Or use the Makefile
make db-setup
```

### Installation

```bash
# Clone the repository
git clone https://github.com/yesoreyeram/angidi-demo-app.git
cd angidi-demo-app/backend

# Install development tools
make install-tools

# Download dependencies
make deps

# Run database migrations
make migrate-up
```

### Environment Configuration

Copy the example environment file and update values:

```bash
cp .env.example .env
```

Edit `.env` with your database credentials and other configuration:

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=angidi_dev
DB_SSLMODE=disable

# JWT Configuration
JWT_SECRET=your-secret-key-change-in-production-MUST-BE-STRONG

# Admin Bootstrap
ADMIN_EMAIL=admin@example.com
ADMIN_PASSWORD=ChangeThisSecureAdminPassword123!
ADMIN_NAME=System Administrator
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

### Database Management

```bash
# Run all pending migrations
make migrate-up

# Rollback last migration
make migrate-down

# Create new migration
make migrate-create NAME=add_new_table

# Setup database from scratch
make db-setup
```

## Project Structure

```
backend/
├── cmd/
│   └── api/              # API server entry point
│       └── main.go
├── internal/             # Private application code
│   ├── database/         # Database connection pool
│   ├── user/             # User domain
│   ├── product/          # Product domain
│   ├── category/         # Category domain
│   ├── cart/             # Cart domain (future)
│   ├── order/            # Order domain (future)
│   ├── gateway/          # API Gateway and routing
│   └── common/           # Shared internal code
│       └── middleware/   # HTTP middleware
├── pkg/                  # Public, reusable packages
│   ├── logger/           # Structured logging
│   ├── config/           # Configuration management
│   ├── jwt/              # JWT token service
│   └── response/         # HTTP response utilities
├── migrations/           # Database migration files
│   ├── 000001_create_users_table.up.sql
│   ├── 000001_create_users_table.down.sql
│   ├── 000002_create_categories_table.up.sql
│   └── ...
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
- `DB_HOST` - Database host (default: localhost)
- `DB_PORT` - Database port (default: 5432)
- `DB_USER` - Database user (default: postgres)
- `DB_PASSWORD` - Database password (default: postgres)
- `DB_NAME` - Database name (default: angidi_dev)
- `DB_SSLMODE` - SSL mode: disable, require, verify-ca, verify-full (default: disable)
- `JWT_SECRET` - Secret key for JWT token signing (required in production)
- `ADMIN_EMAIL` - Initial admin email (required for first-time setup)
- `ADMIN_PASSWORD` - Initial admin password (required for first-time setup, min 12 characters)
- `ADMIN_NAME` - Initial admin name (optional, defaults to "System Administrator")

### Example

```bash
# Run with custom port
SERVER_PORT=9090 make run

# Run with debug logging
LOG_LEVEL=debug make run

# Create initial admin user on first run
ADMIN_EMAIL=admin@example.com ADMIN_PASSWORD=SecureAdminPass123! make run
```

## Admin Bootstrap

### First-Time Setup

On first startup, the system can automatically create an initial admin user if none exists. This is controlled via environment variables:

```bash
export ADMIN_EMAIL="admin@yourcompany.com"
export ADMIN_PASSWORD="YourSecurePassword123!"
export ADMIN_NAME="System Administrator"

make run
```

**Security Requirements**:
- Password must be at least 12 characters
- Admin bootstrap only runs if no admin user exists
- Password is automatically cleared from environment after bootstrap
- For production, use secrets management (Vault, AWS Secrets Manager, etc.)

**Production Deployment**:
```bash
# Using Docker
docker run -e ADMIN_EMAIL=admin@company.com \
           -e ADMIN_PASSWORD=$(vault kv get -field=password secret/admin) \
           angidi-api:latest

# Using Kubernetes
kubectl create secret generic admin-credentials \
  --from-literal=email=admin@company.com \
  --from-literal=password=SecurePassword123!
```

See `docs/phases/phase-02-core-services/ADMIN_BOOTSTRAP.md` for detailed security considerations and alternative approaches.

## API Endpoints

### Health Check

```bash
GET /health
```

Returns server health status.

**Response:**
```json
{
  "data": {
    "status": "healthy",
    "timestamp": "2025-10-27T03:00:00Z"
  }
}
```

### User Authentication

#### Register User

```bash
POST /api/v1/users/register
```

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "SecurePass123!",
  "name": "John Doe"
}
```

**Response (201 Created):**
```json
{
  "data": {
    "id": "uuid",
    "email": "user@example.com",
    "name": "John Doe",
    "role": "user",
    "created_at": "2025-10-27T03:00:00Z",
    "updated_at": "2025-10-27T03:00:00Z"
  }
}
```

#### Login

```bash
POST /api/v1/users/login
```

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "SecurePass123!"
}
```

**Response (200 OK):**
```json
{
  "data": {
    "access_token": "eyJhbG...",
    "refresh_token": "eyJhbG...",
    "expires_in": 900,
    "user": {
      "id": "uuid",
      "email": "user@example.com",
      "name": "John Doe",
      "role": "user",
      "created_at": "2025-10-27T03:00:00Z",
      "updated_at": "2025-10-27T03:00:00Z"
    }
  }
}
```

#### Refresh Token

```bash
POST /api/v1/users/refresh-token
```

**Request Body:**
```json
{
  "refresh_token": "eyJhbG..."
}
```

**Response (200 OK):**
```json
{
  "data": {
    "access_token": "eyJhbG...",
    "refresh_token": "eyJhbG...",
    "expires_in": 900,
    "user": { ... }
  }
}
```

#### Get Profile (Protected)

```bash
GET /api/v1/users/me
Authorization: Bearer <access_token>
```

**Response (200 OK):**
```json
{
  "data": {
    "id": "uuid",
    "email": "user@example.com",
    "name": "John Doe",
    "role": "user",
    "created_at": "2025-10-27T03:00:00Z",
    "updated_at": "2025-10-27T03:00:00Z"
  }
}
```

#### Update Profile (Protected)

```bash
PUT /api/v1/users/me
Authorization: Bearer <access_token>
```

**Request Body:**
```json
{
  "name": "John Updated"
}
```

**Response (200 OK):**
```json
{
  "data": {
    "id": "uuid",
    "email": "user@example.com",
    "name": "John Updated",
    "role": "user",
    "created_at": "2025-10-27T03:00:00Z",
    "updated_at": "2025-10-27T03:30:00Z"
  }
}
```

### Product Management

#### List Products

```bash
GET /api/v1/products?page=1&page_size=10&category_id=cat1&search=phone&min_price=100&max_price=1000
```

**Query Parameters:**
- `page` (optional): Page number (default: 1)
- `page_size` (optional): Items per page (default: 10, max: 100)
- `category_id` (optional): Filter by category
- `search` (optional): Search in name and description
- `min_price` (optional): Minimum price filter
- `max_price` (optional): Maximum price filter

**Response (200 OK):**
```json
{
  "data": {
    "products": [
      {
        "id": "uuid",
        "name": "Product Name",
        "description": "Product description",
        "price": 99.99,
        "stock": 100,
        "category_id": "cat1",
        "image_url": "https://example.com/image.jpg",
        "created_at": "2025-10-27T03:00:00Z",
        "updated_at": "2025-10-27T03:00:00Z"
      }
    ],
    "total_count": 50,
    "page": 1,
    "page_size": 10,
    "total_pages": 5
  }
}
```

#### Get Product

```bash
GET /api/v1/products/:id
```

**Response (200 OK):**
```json
{
  "data": {
    "id": "uuid",
    "name": "Product Name",
    "description": "Product description",
    "price": 99.99,
    "stock": 100,
    "category_id": "cat1",
    "image_url": "https://example.com/image.jpg",
    "created_at": "2025-10-27T03:00:00Z",
    "updated_at": "2025-10-27T03:00:00Z"
  }
}
```

#### Create Product (Admin Only)

```bash
POST /api/v1/products
Authorization: Bearer <access_token>
```

**Request Body:**
```json
{
  "name": "New Product",
  "description": "Product description",
  "price": 99.99,
  "stock": 100,
  "category_id": "cat1",
  "image_url": "https://example.com/image.jpg"
}
```

**Response (201 Created):**
```json
{
  "data": {
    "id": "uuid",
    "name": "New Product",
    ...
  }
}
```

#### Update Product (Admin Only)

```bash
PUT /api/v1/products/:id
Authorization: Bearer <access_token>
```

**Request Body:**
```json
{
  "name": "Updated Product",
  "price": 149.99
}
```

**Response (200 OK):**
```json
{
  "data": {
    "id": "uuid",
    "name": "Updated Product",
    "price": 149.99,
    ...
  }
}
```

#### Delete Product (Admin Only)

```bash
DELETE /api/v1/products/:id
Authorization: Bearer <access_token>
```

**Response (204 No Content)**

### Error Responses

All error responses follow this format:

```json
{
  "error": {
    "code": "ERROR_CODE",
    "message": "Human-readable error message",
    "details": [
      {
        "field": "email",
        "message": "Email format is invalid"
      }
    ],
    "request_id": "uuid"
  }
}
```

**Common Error Codes:**
- `VALIDATION_ERROR` (400): Invalid request parameters
- `AUTHENTICATION_ERROR` (401): Invalid or missing credentials
- `AUTHORIZATION_ERROR` (403): Insufficient permissions
- `NOT_FOUND` (404): Resource not found
- `CONFLICT` (409): Resource already exists
- `INTERNAL_ERROR` (500): Server error
- `RATE_LIMIT_EXCEEDED` (429): Too many requests

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
- `github.com/go-chi/chi/v5` - HTTP router
- `github.com/go-chi/cors` - CORS middleware
- `github.com/golang-jwt/jwt/v5` - JWT tokens
- `golang.org/x/crypto` - bcrypt password hashing
- `github.com/go-playground/validator/v10` - Request validation
- `github.com/google/uuid` - UUID generation
- `golang.org/x/time` - Rate limiting
- `go.uber.org/zap` - Structured logging
- `gopkg.in/yaml.v3` - YAML parsing

Development dependencies:
- `github.com/stretchr/testify` - Test assertions
- `go.uber.org/mock` - Mocking
- `golangci-lint` - Linting

## Environment Variables

### JWT Configuration

```bash
JWT_SECRET=your-secret-key-change-in-production
JWT_ACCESS_TOKEN_DURATION=15m
JWT_REFRESH_TOKEN_DURATION=168h
```

### CORS Configuration

```bash
CORS_ALLOWED_ORIGINS=http://localhost:3000,http://localhost:8080
```

## License

MIT License - see LICENSE file for details.

## Contributing

See DEVELOPMENT.md for development guidelines and contribution process.

---

**Current Phase**: Phase 2 - Core Services Development  
**Status**: ✅ Complete  
**Last Updated**: 2025-10-27

**Features Implemented:**
- User authentication (registration, login, JWT tokens)
- Product CRUD operations with pagination and filters
- API Gateway with middleware (auth, rate limiting, CORS, logging)
- Comprehensive unit tests (>80% coverage for critical paths)
