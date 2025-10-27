# Phase 3: Database Integration & Persistence

**Status**: ⏳ Planning  
**Duration**: 1 week  
**Start Date**: TBD  
**Completion Date**: TBD

---

## Overview

Phase 3 marks a critical transition from in-memory data storage to persistent database storage using PostgreSQL. This phase focuses on implementing proper data persistence, database schema design, migration management, and the repository pattern. By the end of this phase, all user and product data will be stored in a relational database with proper ACID guarantees, migrations for schema evolution, and connection pooling for performance.

This phase introduces fundamental system design concepts including relational data modeling, database migrations, connection pooling, and the repository pattern. These concepts are essential for building scalable, production-ready applications that can handle millions of users and products.

---

## Goals & Objectives

### Primary Goals

1. **Integrate PostgreSQL Database**
   - Set up PostgreSQL 15+ as primary data store
   - Configure database connection with proper pooling
   - Implement health checks and monitoring
   - Support multiple environments (dev, test, prod)

2. **Design and Implement Database Schema**
   - Design normalized schema for users, products, and categories
   - Implement proper indexing strategy
   - Add foreign key constraints and relationships
   - Support data integrity with constraints

3. **Set Up Database Migration System**
   - Implement golang-migrate for schema versioning
   - Create initial migration scripts
   - Support both up and down migrations
   - Automate migration execution in CI/CD

4. **Implement Repository Pattern**
   - Replace in-memory repositories with database implementations
   - Maintain interface compatibility from Phase 2
   - Add transaction support
   - Implement proper error handling

5. **Add Database Integration Tests**
   - Use Testcontainers for isolated database testing
   - Test all CRUD operations
   - Verify transaction behavior
   - Benchmark database performance

### Success Criteria

- ✅ PostgreSQL database running locally and in CI
- ✅ All user data persisted in users table
- ✅ All product data persisted in products table
- ✅ Categories table implemented with relationships
- ✅ Migration system working (up and down)
- ✅ Repository pattern implemented for all entities
- ✅ Database connection pooling configured
- ✅ Integration tests using Testcontainers pass
- ✅ Data survives server restarts
- ✅ Performance baseline established
- ✅ Documentation complete

---

## System Design Concepts Introduced

### 1. Relational Data Modeling

**Concept**: Designing database schemas using relational principles including normalization, relationships, and constraints.

**Implementation in Phase 3**:

**Entities and Relationships**:
```
Users (1) ----< (N) Orders
Products (N) ----< (1) Categories
Products (N) ----< (N) OrderItems ----< (N) Orders
```

**Normalization**:
- **1NF**: Atomic values, no repeating groups
- **2NF**: No partial dependencies on composite keys
- **3NF**: No transitive dependencies

**Schema Design**:
```sql
-- Users table
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    name VARCHAR(100) NOT NULL,
    role VARCHAR(50) NOT NULL DEFAULT 'user',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    
    CONSTRAINT email_format CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}$'),
    CONSTRAINT role_valid CHECK (role IN ('user', 'admin', 'moderator'))
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_created_at ON users(created_at);

-- Categories table
CREATE TABLE categories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) UNIQUE NOT NULL,
    description TEXT,
    parent_id UUID REFERENCES categories(id) ON DELETE SET NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT name_not_empty CHECK (LENGTH(TRIM(name)) > 0)
);

CREATE INDEX idx_categories_name ON categories(name);
CREATE INDEX idx_categories_parent_id ON categories(parent_id);

-- Products table
CREATE TABLE products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    stock INTEGER NOT NULL DEFAULT 0,
    category_id UUID NOT NULL REFERENCES categories(id) ON DELETE RESTRICT,
    image_url VARCHAR(500),
    sku VARCHAR(100) UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    
    CONSTRAINT price_positive CHECK (price >= 0),
    CONSTRAINT stock_non_negative CHECK (stock >= 0),
    CONSTRAINT name_not_empty CHECK (LENGTH(TRIM(name)) > 0)
);

CREATE INDEX idx_products_category_id ON products(category_id);
CREATE INDEX idx_products_name ON products(name);
CREATE INDEX idx_products_price ON products(price);
CREATE INDEX idx_products_created_at ON products(created_at);
CREATE INDEX idx_products_sku ON products(sku);
```

**Benefits**:
- Data integrity through constraints
- Efficient queries with proper indexes
- Scalable schema design
- Clear relationships between entities

**Trade-offs**:
- More complex queries with joins
- Potential performance impact with many relationships
- Schema changes require migrations

---

### 2. Database Migrations

**Concept**: Version control for database schemas, allowing safe and repeatable schema changes across environments.

**Migration Strategy**:

**Tools**: golang-migrate (https://github.com/golang-migrate/migrate)

**Migration Structure**:
```
backend/migrations/
├── 000001_create_users_table.up.sql
├── 000001_create_users_table.down.sql
├── 000002_create_categories_table.up.sql
├── 000002_create_categories_table.down.sql
├── 000003_create_products_table.up.sql
├── 000003_create_products_table.down.sql
├── 000004_add_indexes.up.sql
├── 000004_add_indexes.down.sql
```

**Best Practices**:
1. **Always provide down migrations** for rollback capability
2. **One logical change per migration** for atomic changes
3. **Test migrations in both directions** (up and down)
4. **Never modify existing migrations** once applied to production
5. **Use transactions** for safety (when supported)
6. **Add migrations to version control**

**Migration Example**:
```sql
-- 000001_create_users_table.up.sql
BEGIN;

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    name VARCHAR(100) NOT NULL,
    role VARCHAR(50) NOT NULL DEFAULT 'user',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_email ON users(email);

COMMIT;

-- 000001_create_users_table.down.sql
BEGIN;

DROP TABLE IF EXISTS users;

COMMIT;
```

**Execution**:
```bash
# Apply all pending migrations
migrate -path ./migrations -database "postgresql://user:pass@localhost:5432/dbname?sslmode=disable" up

# Rollback last migration
migrate -path ./migrations -database "postgresql://user:pass@localhost:5432/dbname?sslmode=disable" down 1

# Check current version
migrate -path ./migrations -database "postgresql://user:pass@localhost:5432/dbname?sslmode=disable" version
```

---

### 3. Connection Pooling

**Concept**: Reusing database connections to reduce overhead and improve performance.

**Implementation with pgx**:

```go
package database

import (
    "context"
    "fmt"
    "time"
    
    "github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
    Host            string
    Port            int
    User            string
    Password        string
    Database        string
    SSLMode         string
    MaxConns        int32
    MinConns        int32
    MaxConnLifetime time.Duration
    MaxConnIdleTime time.Duration
}

func NewPool(ctx context.Context, cfg Config) (*pgxpool.Pool, error) {
    connString := fmt.Sprintf(
        "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
        cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode,
    )
    
    config, err := pgxpool.ParseConfig(connString)
    if err != nil {
        return nil, fmt.Errorf("failed to parse connection string: %w", err)
    }
    
    // Configure pool
    config.MaxConns = cfg.MaxConns                      // Max connections in pool
    config.MinConns = cfg.MinConns                      // Min idle connections
    config.MaxConnLifetime = cfg.MaxConnLifetime        // Max connection lifetime
    config.MaxConnIdleTime = cfg.MaxConnIdleTime        // Max idle time
    config.HealthCheckPeriod = 1 * time.Minute          // Health check frequency
    
    pool, err := pgxpool.NewWithConfig(ctx, config)
    if err != nil {
        return nil, fmt.Errorf("failed to create connection pool: %w", err)
    }
    
    // Verify connection
    if err := pool.Ping(ctx); err != nil {
        return nil, fmt.Errorf("failed to ping database: %w", err)
    }
    
    return pool, nil
}
```

**Pool Configuration Guidelines**:
- **MaxConns**: Set based on database limits and application load
  - Small apps: 10-25
  - Medium apps: 25-50
  - Large apps: 50-100+
- **MinConns**: Keep 2-5 connections always open for quick response
- **MaxConnLifetime**: 1-2 hours to avoid stale connections
- **MaxConnIdleTime**: 5-15 minutes to release unused connections

**Benefits**:
- Reduced connection overhead
- Better resource utilization
- Improved performance
- Automatic connection management

---

### 4. Repository Pattern

**Concept**: Abstraction layer between business logic and data access, providing a collection-like interface for domain objects.

**Implementation**:

**Interface (remains from Phase 2)**:
```go
type UserRepository interface {
    Create(ctx context.Context, user *User) error
    FindByID(ctx context.Context, id string) (*User, error)
    FindByEmail(ctx context.Context, email string) (*User, error)
    Update(ctx context.Context, user *User) error
    Delete(ctx context.Context, id string) error
}
```

**PostgreSQL Implementation**:
```go
type PostgresUserRepository struct {
    pool *pgxpool.Pool
}

func NewPostgresUserRepository(pool *pgxpool.Pool) *PostgresUserRepository {
    return &PostgresUserRepository{pool: pool}
}

func (r *PostgresUserRepository) Create(ctx context.Context, user *User) error {
    query := `
        INSERT INTO users (id, email, password_hash, name, role, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
    `
    
    _, err := r.pool.Exec(ctx, query,
        user.ID,
        user.Email,
        user.PasswordHash,
        user.Name,
        user.Role,
        user.CreatedAt,
        user.UpdatedAt,
    )
    
    if err != nil {
        return fmt.Errorf("failed to create user: %w", err)
    }
    
    return nil
}

func (r *PostgresUserRepository) FindByEmail(ctx context.Context, email string) (*User, error) {
    query := `
        SELECT id, email, password_hash, name, role, created_at, updated_at
        FROM users
        WHERE email = $1 AND deleted_at IS NULL
    `
    
    var user User
    err := r.pool.QueryRow(ctx, query, email).Scan(
        &user.ID,
        &user.Email,
        &user.PasswordHash,
        &user.Name,
        &user.Role,
        &user.CreatedAt,
        &user.UpdatedAt,
    )
    
    if err != nil {
        if errors.Is(err, pgx.ErrNoRows) {
            return nil, ErrUserNotFound
        }
        return nil, fmt.Errorf("failed to find user: %w", err)
    }
    
    return &user, nil
}
```

**Benefits**:
- Separation of concerns
- Easy to test with mocks
- Swappable implementations
- Centralized data access logic

---

### 5. ACID Transactions

**Concept**: Database transactions ensuring Atomicity, Consistency, Isolation, and Durability.

**Implementation**:

```go
func (r *PostgresUserRepository) CreateWithProfile(ctx context.Context, user *User, profile *UserProfile) error {
    // Begin transaction
    tx, err := r.pool.Begin(ctx)
    if err != nil {
        return fmt.Errorf("failed to begin transaction: %w", err)
    }
    defer tx.Rollback(ctx) // Rollback if not committed
    
    // Insert user
    _, err = tx.Exec(ctx, `
        INSERT INTO users (id, email, password_hash, name, role, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
    `, user.ID, user.Email, user.PasswordHash, user.Name, user.Role, user.CreatedAt, user.UpdatedAt)
    
    if err != nil {
        return fmt.Errorf("failed to insert user: %w", err)
    }
    
    // Insert profile
    _, err = tx.Exec(ctx, `
        INSERT INTO user_profiles (user_id, phone, address, created_at)
        VALUES ($1, $2, $3, $4)
    `, user.ID, profile.Phone, profile.Address, profile.CreatedAt)
    
    if err != nil {
        return fmt.Errorf("failed to insert profile: %w", err)
    }
    
    // Commit transaction
    if err = tx.Commit(ctx); err != nil {
        return fmt.Errorf("failed to commit transaction: %w", err)
    }
    
    return nil
}
```

**ACID Properties**:
- **Atomicity**: All or nothing - either all operations succeed or all fail
- **Consistency**: Database moves from one valid state to another
- **Isolation**: Concurrent transactions don't interfere
- **Durability**: Committed data persists even after system failure

---

### 6. Soft Deletes

**Concept**: Mark records as deleted instead of physically removing them, allowing data recovery and audit trails.

**Implementation**:

```sql
-- Add deleted_at column
ALTER TABLE users ADD COLUMN deleted_at TIMESTAMP NULL;
ALTER TABLE products ADD COLUMN deleted_at TIMESTAMP NULL;

-- Update queries to exclude soft-deleted records
SELECT * FROM users WHERE deleted_at IS NULL;

-- Soft delete implementation
UPDATE users SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1;

-- Restore soft-deleted record
UPDATE users SET deleted_at = NULL WHERE id = $1;
```

**Benefits**:
- Data recovery capability
- Audit trail maintenance
- Referential integrity preserved
- Compliance requirements met

**Trade-offs**:
- Increased storage requirements
- More complex queries
- Unique constraints need special handling

---

## Implementation Plan

### Task 1: PostgreSQL Setup (4 hours)

#### 1.1 Local PostgreSQL Setup
**Duration**: 1 hour

**Steps**:
1. Install PostgreSQL 15+ locally
2. Create development database
3. Configure authentication
4. Test connection

**Deliverables**:
```bash
# PostgreSQL installation (macOS)
brew install postgresql@15

# Start PostgreSQL
brew services start postgresql@15

# Create database
createdb angidi_dev
createdb angidi_test

# Create user
createuser -P angidi_user
# Set password: angidi_pass

# Grant privileges
psql -d angidi_dev -c "GRANT ALL PRIVILEGES ON DATABASE angidi_dev TO angidi_user;"
psql -d angidi_test -c "GRANT ALL PRIVILEGES ON DATABASE angidi_test TO angidi_user;"
```

---

#### 1.2 Database Configuration
**Duration**: 1 hour

**Configuration File** (`config/database.yaml`):
```yaml
database:
  development:
    host: localhost
    port: 5432
    user: angidi_user
    password: angidi_pass
    database: angidi_dev
    sslmode: disable
    pool:
      max_conns: 25
      min_conns: 5
      max_conn_lifetime: 1h
      max_conn_idle_time: 15m
  
  test:
    host: localhost
    port: 5432
    user: angidi_user
    password: angidi_pass
    database: angidi_test
    sslmode: disable
    pool:
      max_conns: 10
      min_conns: 2
      max_conn_lifetime: 30m
      max_conn_idle_time: 5m
  
  production:
    host: ${DB_HOST}
    port: ${DB_PORT}
    user: ${DB_USER}
    password: ${DB_PASSWORD}
    database: ${DB_NAME}
    sslmode: require
    pool:
      max_conns: 100
      min_conns: 10
      max_conn_lifetime: 2h
      max_conn_idle_time: 30m
```

---

#### 1.3 Connection Pool Setup
**Duration**: 1 hour

**Package**: `backend/internal/database/pool.go`

**Implementation**: Connection pool manager with health checks

---

#### 1.4 CI/CD Database Setup
**Duration**: 1 hour

**GitHub Actions**:
```yaml
# .github/workflows/backend-tests.yml
services:
  postgres:
    image: postgres:15-alpine
    env:
      POSTGRES_USER: angidi_user
      POSTGRES_PASSWORD: angidi_pass
      POSTGRES_DB: angidi_test
    options: >-
      --health-cmd pg_isready
      --health-interval 10s
      --health-timeout 5s
      --health-retries 5
    ports:
      - 5432:5432
```

---

### Task 2: Database Schema Design (6 hours)

#### 2.1 ERD Design
**Duration**: 2 hours

**Entity Relationship Diagram**:
```
┌─────────────┐         ┌──────────────┐         ┌─────────────┐
│    Users    │         │  Categories  │         │  Products   │
├─────────────┤         ├──────────────┤         ├─────────────┤
│ id (PK)     │         │ id (PK)      │         │ id (PK)     │
│ email       │         │ name         │         │ name        │
│ password    │         │ description  │         │ description │
│ name        │         │ parent_id(FK)│───┐     │ price       │
│ role        │         │ created_at   │   │     │ stock       │
│ created_at  │         │ updated_at   │   │     │ category_id │
│ updated_at  │         └──────────────┘   │     │ image_url   │
│ deleted_at  │                            │     │ sku         │
└─────────────┘                            │     │ created_at  │
                                           │     │ updated_at  │
                                           │     │ deleted_at  │
                                           │     └─────────────┘
                                           │            │
                                           └────────────┘
                                             (belongs to)
```

---

#### 2.2 Schema SQL Files
**Duration**: 2 hours

**Create migration files** for all tables with proper constraints and indexes.

---

#### 2.3 Seed Data Scripts
**Duration**: 1 hour

**Purpose**: Populate database with test data for development.

**File**: `backend/migrations/seeds/001_seed_categories.sql`

```sql
-- Insert default categories
INSERT INTO categories (id, name, description) VALUES
    (gen_random_uuid(), 'Electronics', 'Electronic devices and accessories'),
    (gen_random_uuid(), 'Clothing', 'Apparel and fashion items'),
    (gen_random_uuid(), 'Books', 'Books and publications'),
    (gen_random_uuid(), 'Home & Garden', 'Home and garden products'),
    (gen_random_uuid(), 'Sports', 'Sports and outdoor equipment');

-- Insert sample products
-- (Similar structure)
```

---

#### 2.4 Schema Validation
**Duration**: 1 hour

**Validation Checklist**:
- All tables have primary keys
- Foreign keys properly defined
- Indexes on frequently queried columns
- Constraints for data integrity
- Proper data types selected
- Timestamps for audit trail

---

### Task 3: Migration System (6 hours)

#### 3.1 Install golang-migrate
**Duration**: 30 minutes

**Installation**:
```bash
# macOS
brew install golang-migrate

# Linux
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz
sudo mv migrate /usr/local/bin/

# Verify
migrate -version
```

---

#### 3.2 Create Migration Files
**Duration**: 3 hours

**Migration files** for:
1. Users table
2. Categories table
3. Products table
4. Indexes
5. Constraints

**Naming Convention**: `YYYYMMDDHHMMSS_description.up.sql` / `.down.sql`

---

#### 3.3 Migration CLI Tool
**Duration**: 2 hours

**Purpose**: Simplify migration commands.

**File**: `backend/cmd/migrate/main.go`

```go
package main

import (
    "flag"
    "fmt"
    "log"
    
    "github.com/golang-migrate/migrate/v4"
    _ "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
    var (
        migrationsPath = flag.String("path", "./migrations", "Path to migrations")
        dbURL          = flag.String("database", "", "Database URL")
        direction      = flag.String("direction", "up", "Migration direction: up or down")
        steps          = flag.Int("steps", 0, "Number of migrations to apply")
    )
    flag.Parse()
    
    m, err := migrate.New(
        fmt.Sprintf("file://%s", *migrationsPath),
        *dbURL,
    )
    if err != nil {
        log.Fatal(err)
    }
    defer m.Close()
    
    switch *direction {
    case "up":
        if *steps == 0 {
            err = m.Up()
        } else {
            err = m.Steps(*steps)
        }
    case "down":
        if *steps == 0 {
            err = m.Down()
        } else {
            err = m.Steps(-*steps)
        }
    default:
        log.Fatal("Invalid direction")
    }
    
    if err != nil && err != migrate.ErrNoChange {
        log.Fatal(err)
    }
    
    version, dirty, err := m.Version()
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Current version: %d, Dirty: %v\n", version, dirty)
}
```

---

#### 3.4 Makefile Integration
**Duration**: 30 minutes

**Add to Makefile**:
```makefile
# Database migrations
.PHONY: migrate-up migrate-down migrate-create migrate-version

migrate-up:
migrate -path ./migrations -database "postgresql://angidi_user:angidi_pass@localhost:5432/angidi_dev?sslmode=disable" up

migrate-down:
migrate -path ./migrations -database "postgresql://angidi_user:angidi_pass@localhost:5432/angidi_dev?sslmode=disable" down 1

migrate-create:
@read -p "Enter migration name: " name; \
migrate create -ext sql -dir ./migrations -seq $$name

migrate-version:
migrate -path ./migrations -database "postgresql://angidi_user:angidi_pass@localhost:5432/angidi_dev?sslmode=disable" version
```

---

### Task 4: Repository Implementation (12 hours)

#### 4.1 Update User Repository
**Duration**: 4 hours

**Tasks**:
- Implement PostgreSQL user repository
- Replace in-memory implementation
- Add transaction support
- Handle unique constraint violations
- Implement soft deletes

**Files**:
- `internal/user/repository_postgres.go`
- Update `internal/user/repository.go` interface if needed

---

#### 4.2 Update Product Repository
**Duration**: 4 hours

**Tasks**:
- Implement PostgreSQL product repository
- Add pagination support
- Implement filtering and search
- Add category relationship
- Handle foreign key constraints

**Files**:
- `internal/product/repository_postgres.go`

---

#### 4.3 Category Repository
**Duration**: 2 hours

**Tasks**:
- Create category domain model
- Implement category repository
- Support hierarchical categories
- Add CRUD operations

**Files**:
- `internal/category/model.go`
- `internal/category/repository.go`
- `internal/category/repository_postgres.go`

---

#### 4.4 Repository Testing
**Duration**: 2 hours

**Tasks**:
- Write unit tests for repositories
- Test error scenarios
- Test constraint violations
- Test transaction behavior

---

### Task 5: Service Layer Updates (6 hours)

#### 5.1 Update User Service
**Duration**: 2 hours

**Changes**:
- Update to use PostgreSQL repository
- Add transaction handling
- Update error messages
- Handle database-specific errors

---

#### 5.2 Update Product Service
**Duration**: 2 hours

**Changes**:
- Update to use PostgreSQL repository
- Add category validation
- Update filtering logic
- Handle database-specific errors

---

#### 5.3 Category Service
**Duration**: 2 hours

**Tasks**:
- Create category service
- Implement business logic
- Add validation
- Handle hierarchical operations

---

### Task 6: Integration Testing (8 hours)

#### 6.1 Testcontainers Setup
**Duration**: 2 hours

**Purpose**: Run PostgreSQL in Docker for tests.

**Package**: `github.com/testcontainers/testcontainers-go`

**Implementation**:
```go
package testutil

import (
    "context"
    "fmt"
    "testing"
    
    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/testcontainers/testcontainers-go"
    "github.com/testcontainers/testcontainers-go/wait"
)

func SetupTestDatabase(t *testing.T) (*pgxpool.Pool, func()) {
    ctx := context.Background()
    
    req := testcontainers.ContainerRequest{
        Image:        "postgres:15-alpine",
        ExposedPorts: []string{"5432/tcp"},
        Env: map[string]string{
            "POSTGRES_USER":     "test",
            "POSTGRES_PASSWORD": "test",
            "POSTGRES_DB":       "testdb",
        },
        WaitingFor: wait.ForLog("database system is ready to accept connections"),
    }
    
    container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
        ContainerRequest: req,
        Started:          true,
    })
    if err != nil {
        t.Fatal(err)
    }
    
    host, err := container.Host(ctx)
    if err != nil {
        t.Fatal(err)
    }
    
    port, err := container.MappedPort(ctx, "5432")
    if err != nil {
        t.Fatal(err)
    }
    
    connString := fmt.Sprintf("postgres://test:test@%s:%s/testdb?sslmode=disable", host, port.Port())
    
    pool, err := pgxpool.New(ctx, connString)
    if err != nil {
        t.Fatal(err)
    }
    
    // Run migrations
    runMigrations(connString)
    
    cleanup := func() {
        pool.Close()
        container.Terminate(ctx)
    }
    
    return pool, cleanup
}
```

---

#### 6.2 Repository Integration Tests
**Duration**: 3 hours

**Test Coverage**:
- CRUD operations
- Unique constraints
- Foreign key constraints
- Soft deletes
- Pagination
- Filtering

---

#### 6.3 Service Integration Tests
**Duration**: 2 hours

**Test Coverage**:
- End-to-end flows
- Transaction behavior
- Error handling
- Data validation

---

#### 6.4 Performance Benchmarks
**Duration**: 1 hour

**Benchmark Tests**:
```go
func BenchmarkUserRepository_Create(b *testing.B) {
    pool, cleanup := SetupTestDatabase(b)
    defer cleanup()
    
    repo := NewPostgresUserRepository(pool)
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        user := &User{
            ID:    uuid.New().String(),
            Email: fmt.Sprintf("user%d@example.com", i),
            Name:  "Test User",
        }
        repo.Create(context.Background(), user)
    }
}
```

---

### Task 7: Documentation (4 hours)

#### 7.1 Database Schema Documentation
**Duration**: 1 hour

**File**: `backend/docs/database/SCHEMA.md`

**Content**:
- ERD diagrams
- Table descriptions
- Column descriptions
- Relationship explanations
- Index strategy

---

#### 7.2 Migration Guide
**Duration**: 1 hour

**File**: `backend/docs/database/MIGRATIONS.md`

**Content**:
- How to create migrations
- How to apply migrations
- How to rollback
- Best practices
- Common issues

---

#### 7.3 Setup Instructions
**Duration**: 1 hour

**Updates to README**:
- PostgreSQL installation
- Database setup
- Migration execution
- Environment variables
- Troubleshooting

---

#### 7.4 API Documentation Updates
**Duration**: 1 hour

**Updates**:
- Add category endpoints to OpenAPI spec
- Update error responses
- Document database-related errors
- Add examples

---

### Task 8: Performance Optimization (4 hours)

#### 8.1 Index Optimization
**Duration**: 2 hours

**Tasks**:
- Analyze query patterns
- Add appropriate indexes
- Test index effectiveness
- Document index strategy

---

#### 8.2 Query Optimization
**Duration**: 2 hours

**Tasks**:
- Use EXPLAIN ANALYZE
- Optimize slow queries
- Reduce N+1 queries
- Add query caching strategy

---

## Testing Strategy

### Unit Testing

**Target Coverage**: >85%

**Components to Test**:
- Repository implementations
- Service layer
- Migration helper functions
- Connection pool management

**Example**:
```go
func TestPostgresUserRepository_Create(t *testing.T) {
    pool, cleanup := SetupTestDatabase(t)
    defer cleanup()
    
    repo := NewPostgresUserRepository(pool)
    
    user := &User{
        ID:           uuid.New().String(),
        Email:        "test@example.com",
        PasswordHash: "hashed_password",
        Name:         "Test User",
        Role:         "user",
        CreatedAt:    time.Now(),
        UpdatedAt:    time.Now(),
    }
    
    err := repo.Create(context.Background(), user)
    assert.NoError(t, err)
    
    // Verify user was created
    found, err := repo.FindByEmail(context.Background(), user.Email)
    assert.NoError(t, err)
    assert.Equal(t, user.Email, found.Email)
}
```

---

### Integration Testing

**Test with Testcontainers**:
- Isolated PostgreSQL instance per test suite
- Real database operations
- Migration testing
- Transaction testing

---

### Performance Testing

**Benchmarks**:
- Insert operations
- Query operations
- Pagination
- Filtering
- Connection pool efficiency

**Target Metrics**:
- User creation: <50ms
- User lookup by email: <10ms
- Product listing (50 items): <100ms
- Product search: <200ms

---

## Dependencies

### Backend Dependencies

**New Dependencies for Phase 3**:

```go
// go.mod additions
require (
    github.com/jackc/pgx/v5 v5.5.0                        // PostgreSQL driver
    github.com/jackc/pgx/v5/pgxpool v5.5.0                // Connection pooling
    github.com/golang-migrate/migrate/v4 v4.16.2          // Database migrations
    github.com/testcontainers/testcontainers-go v0.26.0   // Testing with Docker
)
```

**Optional Dependencies**:
```go
// If using GORM
github.com/jinzhu/gorm v1.9.16
gorm.io/driver/postgres v1.5.4
gorm.io/gorm v1.25.5
```

---

## Security Considerations

### Phase 3 Security Checklist

- [ ] Database credentials in environment variables
- [ ] SSL/TLS for production database connections
- [ ] Prepared statements prevent SQL injection
- [ ] Row-level security (RLS) evaluated
- [ ] Database user has minimal required permissions
- [ ] Sensitive data encrypted at rest (future)
- [ ] Connection strings not logged
- [ ] Database backups configured (production)
- [ ] Audit logging for sensitive operations
- [ ] Password fields never logged or exposed

**SQL Injection Prevention**:
```go
// GOOD - Parameterized query
query := "SELECT * FROM users WHERE email = $1"
row := db.QueryRow(ctx, query, email)

// BAD - String concatenation (vulnerable to SQL injection)
query := fmt.Sprintf("SELECT * FROM users WHERE email = '%s'", email)
```

---

## Performance Considerations

### Performance Targets

**Database Operations**:
- Single record insert: <50ms
- Single record lookup by PK: <5ms
- Single record lookup by indexed field: <10ms
- Batch insert (100 records): <500ms
- Paginated list (50 items): <100ms
- Complex query with joins: <200ms

### Optimization Strategies

**Indexing**:
- Primary keys automatically indexed
- Foreign keys indexed
- Frequently queried columns indexed
- Composite indexes for multi-column queries

**Connection Pooling**:
- Reuse connections
- Configure based on load
- Monitor pool statistics

**Query Optimization**:
- Select only needed columns
- Use LIMIT for pagination
- Avoid N+1 queries
- Use EXPLAIN ANALYZE

**Caching** (Phase 7):
- Cache frequently accessed data
- Invalidate on updates
- Use Redis for distributed caching

---

## Acceptance Criteria

### Phase 3 Completion Checklist

#### Database Setup
- [ ] PostgreSQL 15+ installed and running
- [ ] Development database created
- [ ] Test database created
- [ ] Connection pooling configured
- [ ] Health checks working
- [ ] CI/CD database service configured

#### Schema & Migrations
- [ ] Users table migration created
- [ ] Categories table migration created
- [ ] Products table migration created
- [ ] All migrations have up and down scripts
- [ ] Indexes created for performance
- [ ] Constraints enforced
- [ ] Seed data scripts created
- [ ] Migration CLI tool working

#### Repository Implementation
- [ ] User repository uses PostgreSQL
- [ ] Product repository uses PostgreSQL
- [ ] Category repository implemented
- [ ] All CRUD operations working
- [ ] Soft deletes implemented
- [ ] Transaction support added
- [ ] Error handling updated
- [ ] In-memory repositories removed/deprecated

#### Testing
- [ ] Testcontainers setup working
- [ ] Repository integration tests pass (>85% coverage)
- [ ] Service integration tests pass
- [ ] Transaction tests pass
- [ ] Performance benchmarks established
- [ ] All existing tests still pass

#### Documentation
- [ ] Database schema documented
- [ ] ERD diagram created
- [ ] Migration guide written
- [ ] Setup instructions updated
- [ ] API documentation updated
- [ ] Troubleshooting guide complete

#### Quality Gates
- [ ] All linters pass
- [ ] All tests pass (>85% coverage)
- [ ] No SQL injection vulnerabilities
- [ ] Performance targets met
- [ ] Data persists across restarts
- [ ] Migrations work both directions
- [ ] No breaking changes to API

---

## Troubleshooting

### Common Issues

#### Issue 1: Connection Refused

**Symptoms**:
- Cannot connect to PostgreSQL
- "connection refused" errors

**Solutions**:
1. Verify PostgreSQL is running: `pg_isready`
2. Check PostgreSQL port: `lsof -i :5432`
3. Verify host/port in connection string
4. Check firewall settings
5. Verify credentials

---

#### Issue 2: Migration Failed

**Symptoms**:
- Migration errors
- Database in dirty state

**Solutions**:
1. Check migration SQL syntax
2. Verify database permissions
3. Check for conflicting migrations
4. Force version if needed: `migrate force <version>`
5. Manually rollback if necessary

---

#### Issue 3: Slow Queries

**Symptoms**:
- High latency
- Timeout errors

**Solutions**:
1. Run EXPLAIN ANALYZE on slow queries
2. Check if indexes are being used
3. Add missing indexes
4. Optimize query structure
5. Consider pagination for large results

---

#### Issue 4: Connection Pool Exhausted

**Symptoms**:
- "too many connections" errors
- Timeouts acquiring connection

**Solutions**:
1. Increase max connections
2. Decrease connection lifetime
3. Fix connection leaks
4. Implement connection timeout
5. Monitor pool statistics

---

## Migration from Phase 2

### Breaking Changes

**Data Storage**:
- In-memory storage replaced with PostgreSQL
- Data now persists across restarts
- Requires database setup before running

### New Environment Variables

```bash
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=angidi_user
DB_PASSWORD=angidi_pass
DB_NAME=angidi_dev
DB_SSLMODE=disable

# Connection Pool
DB_MAX_CONNS=25
DB_MIN_CONNS=5
DB_MAX_CONN_LIFETIME=1h
DB_MAX_CONN_IDLE_TIME=15m
```

### Migration Steps

1. **Install PostgreSQL**:
   ```bash
   brew install postgresql@15
   ```

2. **Create databases**:
   ```bash
   createdb angidi_dev
   createdb angidi_test
   ```

3. **Run migrations**:
   ```bash
   make migrate-up
   ```

4. **Update environment variables**:
   - Copy `.env.example` to `.env`
   - Update database credentials

5. **Run application**:
   ```bash
   make run
   ```

---

## Next Steps (Phase 4)

Phase 4 will introduce:
- **Docker containerization** for all services
- **Docker Compose** for local development
- **Multi-stage builds** for optimization
- **Health checks** and monitoring
- **Container orchestration** basics

**Preparation**:
- Learn Docker basics
- Understand container networking
- Study multi-stage builds
- Review Docker Compose syntax

---

## References

### Internal Documentation
- [Main Specification](../../specs/SPEC.md)
- [Technology Stack](../../specs/TECH_STACK.md)
- [System Design Concepts](../../specs/SYSTEM_DESIGN_CONCEPTS.md)
- [Architecture Guide](../../architecture/README.md)
- [Phase 2: Core Services](../phase-02-core-services/README.md)
- [All Phases Overview](../README.md)

### External Resources

#### PostgreSQL
- [PostgreSQL Documentation](https://www.postgresql.org/docs/15/)
- [PostgreSQL Tutorial](https://www.postgresqltutorial.com/)
- [PostgreSQL Performance Tips](https://wiki.postgresql.org/wiki/Performance_Optimization)

#### pgx Driver
- [pgx Documentation](https://pkg.go.dev/github.com/jackc/pgx/v5)
- [pgx Pool Guide](https://pkg.go.dev/github.com/jackc/pgx/v5/pgxpool)

#### Migrations
- [golang-migrate](https://github.com/golang-migrate/migrate)
- [Database Migration Best Practices](https://www.dbms tools.com/migration-best-practices)

#### Testing
- [Testcontainers Go](https://golang.testcontainers.org/)
- [PostgreSQL Testing Strategies](https://www.postgresql.org/docs/current/regress.html)

#### Database Design
- [Database Normalization](https://en.wikipedia.org/wiki/Database_normalization)
- [SQL Indexing Guide](https://use-the-index-luke.com/)
- [CAP Theorem](https://en.wikipedia.org/wiki/CAP_theorem)

---

**Version**: 1.0.0  
**Last Updated**: 2025-10-27  
**Status**: ⏳ Planning  
**Next Review**: Upon implementation start
