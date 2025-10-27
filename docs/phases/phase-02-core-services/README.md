# Phase 2: Core Services Development

**Status**: ⏳ Not Started  
**Duration**: 2 weeks  
**Start Date**: TBD  
**Completion Date**: TBD

---

## Overview

Phase 2 marks the transition from project scaffolding to actual business logic implementation. This phase focuses on building the foundational microservices that form the core of our e-commerce platform: **User Service** for authentication and user management, and **Product Service** for catalog management. Additionally, we'll implement a basic **API Gateway** to route requests and handle cross-cutting concerns.

This phase introduces critical system design concepts including RESTful API design, JWT-based authentication, structured logging, and microservices patterns. By the end of this phase, we'll have functioning API endpoints that can handle user registration, authentication, and product CRUD operations.

---

## Goals & Objectives

### Primary Goals

1. **Implement User Service** with complete authentication flow
   - User registration with validation
   - Login with JWT token generation
   - Password hashing with bcrypt
   - User profile management
   - Token refresh mechanism

2. **Implement Product Service** with full CRUD capabilities
   - Create, read, update, delete products
   - Product listing with pagination
   - Product filtering and basic search
   - Category management
   - Product validation

3. **Create Basic API Gateway**
   - Request routing to appropriate services
   - Authentication middleware
   - CORS handling
   - Rate limiting (basic)
   - Request/response logging

4. **Establish API Standards**
   - RESTful API design principles
   - Consistent error handling
   - API versioning strategy
   - OpenAPI/Swagger documentation

5. **Implement Comprehensive Testing**
   - Unit tests for all business logic
   - Integration tests for API endpoints
   - Test coverage >80%
   - Mock external dependencies

### Success Criteria

- ✅ User can register and login successfully
- ✅ JWT tokens are generated and validated correctly
- ✅ Product CRUD operations work end-to-end
- ✅ API Gateway routes requests correctly
- ✅ All API endpoints are documented in OpenAPI spec
- ✅ Unit test coverage >80%
- ✅ Integration tests pass for all endpoints
- ✅ Structured logging captures all requests
- ✅ Error responses follow consistent format
- ✅ API security best practices implemented

---

## System Design Concepts Introduced

### 1. Microservices Decomposition

**Concept**: Breaking down the monolithic application into smaller, independent services based on business capabilities.

**Implementation in Phase 2**:
- **User Service**: Handles all user-related operations (authentication, profile management)
- **Product Service**: Manages product catalog and related operations
- **API Gateway**: Single entry point for all client requests

**Benefits**:
- Independent deployment and scaling
- Technology flexibility per service
- Better fault isolation
- Team autonomy

**Trade-offs**:
- Increased operational complexity
- Network latency between services
- Distributed data management challenges
- Eventual consistency considerations

---

### 2. RESTful API Design

**Concept**: Designing APIs following REST architectural constraints for consistency, scalability, and ease of use.

**REST Principles Applied**:
1. **Resource-Based URLs**: `/api/v1/users`, `/api/v1/products`
2. **HTTP Methods**: GET, POST, PUT, PATCH, DELETE
3. **Stateless Communication**: Each request contains all necessary information
4. **Standard Status Codes**: 200, 201, 400, 401, 404, 500, etc.
5. **HATEOAS**: Links to related resources in responses

**API Design Standards**:
```
POST   /api/v1/users/register         # Create new user
POST   /api/v1/users/login            # Authenticate user
GET    /api/v1/users/me               # Get current user profile
PUT    /api/v1/users/me               # Update user profile
POST   /api/v1/users/refresh-token    # Refresh JWT token

GET    /api/v1/products               # List products (paginated)
POST   /api/v1/products               # Create product
GET    /api/v1/products/:id           # Get product by ID
PUT    /api/v1/products/:id           # Update product
DELETE /api/v1/products/:id           # Delete product
GET    /api/v1/products/search        # Search products
```

---

### 3. JWT Authentication

**Concept**: Stateless authentication using JSON Web Tokens for secure API access.

**JWT Structure**:
- **Header**: Token type and hashing algorithm
- **Payload**: User claims (user_id, roles, expiration)
- **Signature**: HMAC or RSA signature for verification

**Implementation Details**:
```json
{
  "header": {
    "alg": "HS256",
    "typ": "JWT"
  },
  "payload": {
    "user_id": "uuid",
    "email": "user@example.com",
    "roles": ["user"],
    "iat": 1635724800,
    "exp": 1635811200
  }
}
```

**Security Considerations**:
- Short-lived access tokens (15 minutes)
- Long-lived refresh tokens (7 days)
- Secure token storage (httpOnly cookies)
- Token rotation on refresh
- Blacklist for revoked tokens (Phase 3+)

---

### 4. Structured Logging with Zap

**Concept**: Consistent, structured log output for better observability and debugging.

**Log Levels**:
- **DEBUG**: Detailed diagnostic information
- **INFO**: General informational messages
- **WARN**: Warning messages for potentially harmful situations
- **ERROR**: Error events that might still allow the application to continue
- **FATAL**: Severe errors that lead to application termination

**Log Structure**:
```json
{
  "level": "info",
  "ts": "2025-10-27T10:30:00.000Z",
  "caller": "user/handler.go:45",
  "msg": "User registration successful",
  "user_id": "uuid-123",
  "email": "user@example.com",
  "duration_ms": 125
}
```

**Best Practices**:
- Include request ID for tracing
- Log important business events
- Never log sensitive data (passwords, tokens)
- Use structured fields for queryability
- Set appropriate log levels per environment

---

### 5. Password Security

**Concept**: Secure password storage using industry-standard hashing algorithms.

**bcrypt Hashing**:
- Adaptive hashing function resistant to brute-force
- Built-in salt generation
- Configurable cost factor (work factor)
- Computationally expensive to slow down attacks

**Implementation**:
```go
// Hash password with cost factor 12
hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)

// Verify password
err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
```

**Security Requirements**:
- Minimum password length: 8 characters
- Require mix of uppercase, lowercase, numbers, special characters
- Prevent common passwords (top 10,000 list)
- Implement password history (Phase 3+)
- Support password reset flow (Phase 3+)

---

### 6. Input Validation & Sanitization

**Concept**: Validate and sanitize all user inputs to prevent injection attacks and data corruption.

**Validation Strategy**:
1. **Schema Validation**: Use struct tags or validation libraries
2. **Type Checking**: Ensure correct data types
3. **Range Validation**: Check min/max values, lengths
4. **Format Validation**: Email, phone, URL patterns
5. **Business Rule Validation**: Custom logic validation

**Example Validations**:
```go
type RegisterRequest struct {
    Email    string `json:"email" validate:"required,email,max=255"`
    Password string `json:"password" validate:"required,min=8,max=128"`
    Name     string `json:"name" validate:"required,min=2,max=100"`
}

type Product struct {
    Name        string  `json:"name" validate:"required,min=3,max=255"`
    Description string  `json:"description" validate:"max=2000"`
    Price       float64 `json:"price" validate:"required,gt=0"`
    Stock       int     `json:"stock" validate:"required,gte=0"`
    CategoryID  string  `json:"category_id" validate:"required,uuid"`
}
```

---

### 7. Error Handling & Response Standardization

**Concept**: Consistent error responses across all API endpoints for better client integration.

**Standard Error Response**:
```json
{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Invalid request parameters",
    "details": [
      {
        "field": "email",
        "message": "Email format is invalid"
      },
      {
        "field": "password",
        "message": "Password must be at least 8 characters"
      }
    ],
    "request_id": "req-uuid-123"
  }
}
```

**Error Categories**:
- `VALIDATION_ERROR` (400): Invalid input data
- `AUTHENTICATION_ERROR` (401): Invalid or missing credentials
- `AUTHORIZATION_ERROR` (403): Insufficient permissions
- `NOT_FOUND` (404): Resource not found
- `CONFLICT` (409): Resource already exists
- `INTERNAL_ERROR` (500): Server error

---

### 8. API Gateway Pattern

**Concept**: Single entry point for all client requests, handling cross-cutting concerns.

**Responsibilities**:
1. **Request Routing**: Forward requests to appropriate services
2. **Authentication**: Validate JWT tokens
3. **Authorization**: Check user permissions
4. **Rate Limiting**: Prevent abuse
5. **CORS**: Handle cross-origin requests
6. **Request/Response Transformation**: Modify as needed
7. **Logging**: Centralized request logging
8. **Monitoring**: Collect metrics

**Benefits**:
- Simplified client integration
- Centralized security
- Reduced service coupling
- Easy API versioning

---

## Code Quality Improvements

### 1. Clean Architecture Implementation

**Layered Architecture**:
```
┌─────────────────────────────────────┐
│        Handler Layer (HTTP)         │  ← HTTP handlers, request/response
├─────────────────────────────────────┤
│       Service Layer (Business)      │  ← Business logic, validation
├─────────────────────────────────────┤
│    Repository Layer (Data Access)   │  ← Data persistence (mock for now)
├─────────────────────────────────────┤
│       Domain Layer (Entities)       │  ← Core business entities
└─────────────────────────────────────┘
```

**Benefits**:
- Clear separation of concerns
- Easier testing through dependency injection
- Business logic independent of infrastructure
- Framework-agnostic domain layer

---

### 2. Dependency Injection

**Pattern**: Constructor injection for loose coupling and testability.

**Example**:
```go
// Service interface
type UserService interface {
    Register(ctx context.Context, req RegisterRequest) (*User, error)
    Login(ctx context.Context, email, password string) (*AuthResponse, error)
}

// Handler depends on service interface
type UserHandler struct {
    service UserService
    logger  *zap.Logger
}

func NewUserHandler(service UserService, logger *zap.Logger) *UserHandler {
    return &UserHandler{
        service: service,
        logger:  logger,
    }
}
```

---

### 3. Interface-Driven Design

**Approach**: Define interfaces for all major components to enable mocking and testing.

**Key Interfaces**:
```go
type UserRepository interface {
    Create(ctx context.Context, user *User) error
    FindByEmail(ctx context.Context, email string) (*User, error)
    FindByID(ctx context.Context, id string) (*User, error)
    Update(ctx context.Context, user *User) error
}

type ProductRepository interface {
    Create(ctx context.Context, product *Product) error
    FindByID(ctx context.Context, id string) (*Product, error)
    List(ctx context.Context, filters ProductFilters) ([]*Product, error)
    Update(ctx context.Context, product *Product) error
    Delete(ctx context.Context, id string) error
}

type TokenService interface {
    GenerateToken(user *User) (string, error)
    ValidateToken(token string) (*TokenClaims, error)
    RefreshToken(refreshToken string) (string, string, error)
}
```

---

### 4. Error Wrapping with Context

**Pattern**: Add context to errors for better debugging.

```go
import "fmt"

func (s *userService) Register(ctx context.Context, req RegisterRequest) (*User, error) {
    if err := s.validator.Validate(req); err != nil {
        return nil, fmt.Errorf("validation failed: %w", err)
    }
    
    existingUser, err := s.repo.FindByEmail(ctx, req.Email)
    if err != nil && !errors.Is(err, ErrNotFound) {
        return nil, fmt.Errorf("failed to check existing user: %w", err)
    }
    
    if existingUser != nil {
        return nil, ErrUserAlreadyExists
    }
    
    // ... rest of implementation
}
```

---

### 5. Configuration Management

**Approach**: Environment-based configuration with validation.

**Configuration Structure**:
```yaml
server:
  host: "0.0.0.0"
  port: 8080
  read_timeout: 15s
  write_timeout: 15s
  shutdown_timeout: 30s

auth:
  jwt_secret: "${JWT_SECRET}"
  access_token_duration: "15m"
  refresh_token_duration: "168h"  # 7 days
  bcrypt_cost: 12

api:
  version: "v1"
  rate_limit: 100  # requests per minute
  
logging:
  level: "info"
  format: "json"
```

---

## Security Enhancements

### 1. Authentication & Authorization

**Security Measures**:
- ✅ Password hashing with bcrypt (cost factor 12)
- ✅ JWT tokens with short expiration (15 minutes)
- ✅ Refresh token rotation
- ✅ Secure token storage recommendations
- ✅ Protection against timing attacks
- ✅ HTTPS enforcement (production)

---

### 2. Input Validation & Sanitization

**Protection Against**:
- SQL Injection (prepared statements - Phase 3)
- XSS (input sanitization, output encoding)
- Command Injection (input validation)
- Path Traversal (path validation)
- LDAP Injection (input validation)

**Validation Library**: `github.com/go-playground/validator/v10`

---

### 3. Rate Limiting

**Implementation**:
- Per-IP rate limiting
- Per-user rate limiting (authenticated requests)
- Sliding window algorithm
- Configurable limits per endpoint

**Default Limits**:
- Anonymous: 100 requests/minute
- Authenticated: 1000 requests/minute
- Login endpoint: 5 attempts/minute

---

### 4. CORS Configuration

**Secure CORS Setup**:
```go
corsConfig := cors.Config{
    AllowedOrigins:   []string{"https://yourdomain.com"},
    AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
    AllowedHeaders:   []string{"Content-Type", "Authorization"},
    AllowCredentials: true,
    MaxAge:          3600,
}
```

---

### 5. Security Headers

**HTTP Security Headers**:
```
X-Content-Type-Options: nosniff
X-Frame-Options: DENY
X-XSS-Protection: 1; mode=block
Strict-Transport-Security: max-age=31536000; includeSubDomains
Content-Security-Policy: default-src 'self'
```

---

### 6. Secrets Management

**Best Practices**:
- Never commit secrets to version control
- Use environment variables for secrets
- Rotate secrets regularly
- Use different secrets per environment
- Prepare for Vault integration (Phase 14)

---

## Testability Improvements

### 1. Test-Driven Development (TDD)

**Approach**: Write tests before implementation.

**TDD Cycle**:
1. **Red**: Write failing test
2. **Green**: Write minimal code to pass test
3. **Refactor**: Improve code while keeping tests green

---

### 2. Test Structure

**Organization**:
```
internal/user/
├── handler.go
├── handler_test.go      # Handler tests
├── service.go
├── service_test.go      # Service tests
├── repository.go
├── repository_mock.go   # Mock implementation
└── model_test.go        # Model tests
```

---

### 3. Unit Testing Strategy

**Coverage Goals**:
- Handlers: >80%
- Services: >90%
- Utilities: >95%

**Test Categories**:
1. **Happy Path**: Test expected successful scenarios
2. **Error Cases**: Test all error conditions
3. **Edge Cases**: Boundary values, empty inputs
4. **Concurrent Access**: Test thread safety

**Example Unit Test**:
```go
func TestUserService_Register(t *testing.T) {
    tests := []struct {
        name    string
        request RegisterRequest
        setup   func(*mock.MockUserRepository)
        want    *User
        wantErr bool
    }{
        {
            name: "successful registration",
            request: RegisterRequest{
                Email:    "test@example.com",
                Password: "SecurePass123!",
                Name:     "Test User",
            },
            setup: func(repo *mock.MockUserRepository) {
                repo.EXPECT().
                    FindByEmail(gomock.Any(), "test@example.com").
                    Return(nil, ErrNotFound)
                repo.EXPECT().
                    Create(gomock.Any(), gomock.Any()).
                    Return(nil)
            },
            wantErr: false,
        },
        {
            name: "duplicate email",
            request: RegisterRequest{
                Email:    "existing@example.com",
                Password: "SecurePass123!",
                Name:     "Test User",
            },
            setup: func(repo *mock.MockUserRepository) {
                repo.EXPECT().
                    FindByEmail(gomock.Any(), "existing@example.com").
                    Return(&User{}, nil)
            },
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            ctrl := gomock.NewController(t)
            defer ctrl.Finish()

            repo := mock.NewMockUserRepository(ctrl)
            tt.setup(repo)

            service := NewUserService(repo, validator, logger)
            user, err := service.Register(context.Background(), tt.request)

            if tt.wantErr {
                assert.Error(t, err)
            } else {
                assert.NoError(t, err)
                assert.NotNil(t, user)
            }
        })
    }
}
```

---

### 4. Integration Testing

**Scope**: Test HTTP endpoints end-to-end.

**Test Setup**:
```go
func setupTestServer(t *testing.T) *httptest.Server {
    // Initialize services with test dependencies
    userService := setupUserService()
    productService := setupProductService()
    
    // Create router
    router := setupRouter(userService, productService)
    
    return httptest.NewServer(router)
}

func TestUserRegistration_Integration(t *testing.T) {
    server := setupTestServer(t)
    defer server.Close()

    reqBody := `{"email":"test@example.com","password":"SecurePass123!","name":"Test User"}`
    resp, err := http.Post(server.URL+"/api/v1/users/register", "application/json", strings.NewReader(reqBody))
    
    require.NoError(t, err)
    defer resp.Body.Close()
    
    assert.Equal(t, http.StatusCreated, resp.StatusCode)
    
    var user User
    err = json.NewDecoder(resp.Body).Decode(&user)
    require.NoError(t, err)
    assert.NotEmpty(t, user.ID)
    assert.Equal(t, "test@example.com", user.Email)
}
```

---

### 5. Mocking Strategy

**Tools**:
- `github.com/golang/mock/gomock`: Generate mocks from interfaces
- `github.com/stretchr/testify/mock`: Manual mock implementations

**Mock Generation**:
```bash
mockgen -source=internal/user/repository.go -destination=internal/user/repository_mock.go -package=user
```

---

### 6. Test Helpers & Fixtures

**Reusable Test Utilities**:
```go
// Test data builders
func NewTestUser(email string) *User {
    return &User{
        ID:        uuid.New().String(),
        Email:     email,
        Name:      "Test User",
        CreatedAt: time.Now(),
    }
}

func NewTestProduct(name string, price float64) *Product {
    return &Product{
        ID:          uuid.New().String(),
        Name:        name,
        Price:       price,
        Stock:       100,
        CategoryID:  "test-category",
        CreatedAt:   time.Now(),
    }
}
```

---

## Architectural Improvements

### 1. Hexagonal Architecture (Ports & Adapters)

**Structure**:
```
Core Domain (Business Logic)
    ↑                    ↑
    │                    │
Primary Ports      Secondary Ports
(Inbound)          (Outbound)
    ↑                    ↑
    │                    │
Primary Adapters   Secondary Adapters
(HTTP, gRPC)       (Database, Cache, Message Queue)
```

**Benefits**:
- Domain logic independent of infrastructure
- Easy to swap implementations
- Better testability
- Clear boundaries

---

### 2. Domain-Driven Design (DDD) Principles

**Concepts Applied**:
- **Entities**: User, Product (with identity)
- **Value Objects**: Email, Money, Address (immutable)
- **Aggregates**: User aggregate, Product aggregate
- **Repositories**: Data access abstraction
- **Services**: Business operations

---

### 3. SOLID Principles

**Implementation**:

**S - Single Responsibility**:
- Each service handles one domain concept
- Handlers only handle HTTP concerns
- Repositories only handle data persistence

**O - Open/Closed**:
- Extend behavior through interfaces
- Use middleware for cross-cutting concerns

**L - Liskov Substitution**:
- Interfaces can be replaced with any implementation
- Mock repositories work same as real ones in tests

**I - Interface Segregation**:
- Small, focused interfaces
- Clients depend only on methods they use

**D - Dependency Inversion**:
- Depend on abstractions (interfaces), not concretions
- Inject dependencies via constructors

---

### 4. Middleware Chain Pattern

**Implementation**:
```go
router.Use(
    middleware.RequestID(),
    middleware.Logger(logger),
    middleware.Recovery(logger),
    middleware.CORS(corsConfig),
    middleware.RateLimit(rateLimiter),
)

// Protected routes
protected := router.Group("/api/v1")
protected.Use(middleware.Authentication(tokenService))
```

---

### 5. Repository Pattern

**Abstraction**: Separate data access logic from business logic.

**Benefits**:
- Swap data sources without changing business logic
- Easy mocking for tests
- Centralized data access logic
- Support for multiple data sources

---

## UI/UX Improvements

### 1. Frontend API Integration

**API Client Enhancement**:
```typescript
// Enhanced API client with authentication
class APIClient {
  private baseURL: string;
  private accessToken: string | null = null;

  async register(data: RegisterData): Promise<User> {
    const response = await this.post('/users/register', data);
    return response.data;
  }

  async login(email: string, password: string): Promise<AuthResponse> {
    const response = await this.post('/users/login', { email, password });
    this.setAccessToken(response.data.access_token);
    return response.data;
  }

  async getProducts(filters?: ProductFilters): Promise<ProductList> {
    const response = await this.get('/products', { params: filters });
    return response.data;
  }

  private setAccessToken(token: string) {
    this.accessToken = token;
    localStorage.setItem('access_token', token);
  }
}
```

---

### 2. Authentication Flow

**User Journey**:
1. User visits registration page
2. Fills out form with validation
3. Submits form → API call to `/api/v1/users/register`
4. Success → Redirect to login or auto-login
5. Login → Receive JWT tokens
6. Store tokens securely (httpOnly cookie recommended)
7. Include token in subsequent requests
8. Handle token expiration and refresh

---

### 3. Form Validation

**Client-Side Validation**:
```typescript
const registerSchema = z.object({
  email: z.string().email('Invalid email format'),
  password: z
    .string()
    .min(8, 'Password must be at least 8 characters')
    .regex(/[A-Z]/, 'Password must contain uppercase letter')
    .regex(/[a-z]/, 'Password must contain lowercase letter')
    .regex(/[0-9]/, 'Password must contain number'),
  name: z
    .string()
    .min(2, 'Name must be at least 2 characters')
    .max(100, 'Name must not exceed 100 characters'),
});
```

---

### 4. Error Handling & User Feedback

**Consistent Error Display**:
```typescript
interface APIError {
  code: string;
  message: string;
  details?: Array<{
    field: string;
    message: string;
  }>;
}

function handleAPIError(error: APIError) {
  // Display user-friendly error messages
  if (error.code === 'VALIDATION_ERROR' && error.details) {
    error.details.forEach(detail => {
      showFieldError(detail.field, detail.message);
    });
  } else {
    showGlobalError(error.message);
  }
}
```

---

### 5. Loading States & Feedback

**UI States**:
- Loading: Show spinner or skeleton
- Success: Show success message, redirect
- Error: Show error message, allow retry
- Empty: Show empty state with call-to-action

---

### 6. Product Catalog UI

**Features**:
- Product grid with images
- Pagination controls
- Filters (category, price range)
- Sort options (price, name, popularity)
- Product quick view
- Add to cart button (Phase 5)

---

## Implementation Plan

### Task 1: Project Structure Setup (4 hours)

#### 1.1 Create Service Directories
**Duration**: 30 minutes

**Steps**:
1. Create directory structure:
```
backend/internal/
├── user/
│   ├── handler.go
│   ├── service.go
│   ├── repository.go
│   ├── model.go
│   ├── errors.go
│   └── validation.go
├── product/
│   ├── handler.go
│   ├── service.go
│   ├── repository.go
│   ├── model.go
│   ├── errors.go
│   └── validation.go
└── gateway/
    ├── router.go
    ├── middleware.go
    └── server.go
```

**Acceptance Criteria**:
- All directories created
- Each directory has placeholder files
- Package declarations are correct

---

#### 1.2 Setup Common Packages
**Duration**: 1 hour

**Steps**:
1. Create `pkg/jwt/` for token management
2. Create `pkg/validator/` for input validation
3. Create `pkg/errors/` for error handling utilities
4. Create `pkg/response/` for standardized API responses

**Deliverables**:
- JWT token generation and validation utilities
- Validator wrapper with custom rules
- Error response builder
- Success response builder

---

#### 1.3 Update Configuration
**Duration**: 30 minutes

**Steps**:
1. Add JWT configuration to config files
2. Add API gateway configuration
3. Add service-specific configurations
4. Update configuration loading logic

---

### Task 2: User Service Implementation (12 hours)

#### 2.1 Define Domain Models
**Duration**: 1 hour

**Deliverables**:
```go
type User struct {
    ID           string    `json:"id"`
    Email        string    `json:"email"`
    PasswordHash string    `json:"-"`
    Name         string    `json:"name"`
    Role         string    `json:"role"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
}

type RegisterRequest struct {
    Email    string `json:"email" validate:"required,email,max=255"`
    Password string `json:"password" validate:"required,min=8,max=128"`
    Name     string `json:"name" validate:"required,min=2,max=100"`
}

type LoginRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required"`
}

type AuthResponse struct {
    AccessToken  string `json:"access_token"`
    RefreshToken string `json:"refresh_token"`
    ExpiresIn    int    `json:"expires_in"`
    User         *User  `json:"user"`
}
```

---

#### 2.2 Implement Repository Interface
**Duration**: 2 hours

**Implementation**:
- In-memory repository for Phase 2
- Interface for future database implementation
- CRUD operations for users
- Email uniqueness check

**Deliverables**:
```go
type UserRepository interface {
    Create(ctx context.Context, user *User) error
    FindByID(ctx context.Context, id string) (*User, error)
    FindByEmail(ctx context.Context, email string) (*User, error)
    Update(ctx context.Context, user *User) error
    Delete(ctx context.Context, id string) error
}

type InMemoryUserRepository struct {
    users sync.Map
    mutex sync.RWMutex
}
```

---

#### 2.3 Implement Service Layer
**Duration**: 4 hours

**Business Logic**:
- User registration with validation
- Password hashing with bcrypt
- Login with password verification
- JWT token generation
- Token refresh logic
- User profile updates

**Deliverables**:
```go
type UserService interface {
    Register(ctx context.Context, req RegisterRequest) (*User, error)
    Login(ctx context.Context, req LoginRequest) (*AuthResponse, error)
    GetProfile(ctx context.Context, userID string) (*User, error)
    UpdateProfile(ctx context.Context, userID string, req UpdateProfileRequest) (*User, error)
    RefreshToken(ctx context.Context, refreshToken string) (*AuthResponse, error)
}
```

---

#### 2.4 Implement HTTP Handlers
**Duration**: 3 hours

**Endpoints**:
- `POST /api/v1/users/register`: User registration
- `POST /api/v1/users/login`: User login
- `GET /api/v1/users/me`: Get current user profile
- `PUT /api/v1/users/me`: Update user profile
- `POST /api/v1/users/refresh-token`: Refresh JWT token

**Deliverables**:
- Request parsing and validation
- Response formatting
- Error handling
- HTTP status codes

---

#### 2.5 Unit & Integration Tests
**Duration**: 2 hours

**Test Coverage**:
- Service layer tests with mocked repository
- Handler tests with mocked service
- Integration tests for complete flow
- Edge cases and error scenarios

**Target**: >80% code coverage

---

### Task 3: Product Service Implementation (10 hours)

#### 3.1 Define Domain Models
**Duration**: 1 hour

**Deliverables**:
```go
type Product struct {
    ID          string    `json:"id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    Price       float64   `json:"price"`
    Stock       int       `json:"stock"`
    CategoryID  string    `json:"category_id"`
    ImageURL    string    `json:"image_url"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

type CreateProductRequest struct {
    Name        string  `json:"name" validate:"required,min=3,max=255"`
    Description string  `json:"description" validate:"max=2000"`
    Price       float64 `json:"price" validate:"required,gt=0"`
    Stock       int     `json:"stock" validate:"required,gte=0"`
    CategoryID  string  `json:"category_id" validate:"required"`
    ImageURL    string  `json:"image_url" validate:"omitempty,url"`
}

type ProductFilters struct {
    CategoryID  string  `json:"category_id,omitempty"`
    MinPrice    float64 `json:"min_price,omitempty"`
    MaxPrice    float64 `json:"max_price,omitempty"`
    Search      string  `json:"search,omitempty"`
    Page        int     `json:"page" validate:"min=1"`
    PageSize    int     `json:"page_size" validate:"min=1,max=100"`
}
```

---

#### 3.2 Implement Repository Interface
**Duration**: 2 hours

**Implementation**:
- In-memory repository with filtering
- Pagination support
- Search functionality
- CRUD operations

---

#### 3.3 Implement Service Layer
**Duration**: 3 hours

**Business Logic**:
- Product creation with validation
- Product updates
- Product deletion
- Product listing with filters
- Product search
- Pagination logic

---

#### 3.4 Implement HTTP Handlers
**Duration**: 2 hours

**Endpoints**:
- `GET /api/v1/products`: List products with pagination
- `POST /api/v1/products`: Create product (admin only)
- `GET /api/v1/products/:id`: Get product by ID
- `PUT /api/v1/products/:id`: Update product (admin only)
- `DELETE /api/v1/products/:id`: Delete product (admin only)
- `GET /api/v1/products/search`: Search products

---

#### 3.5 Unit & Integration Tests
**Duration**: 2 hours

**Test Coverage**: >80%

---

### Task 4: API Gateway Implementation (8 hours)

#### 4.1 Setup Router
**Duration**: 2 hours

**Framework**: Chi or Gin

**Router Setup**:
```go
func SetupRouter(
    userHandler *user.Handler,
    productHandler *product.Handler,
    logger *zap.Logger,
) http.Handler {
    r := chi.NewRouter()
    
    // Global middleware
    r.Use(middleware.RequestID)
    r.Use(middleware.Logger(logger))
    r.Use(middleware.Recoverer)
    r.Use(middleware.CORS(corsConfig))
    r.Use(middleware.RateLimit(100))
    
    // Public routes
    r.Post("/api/v1/users/register", userHandler.Register)
    r.Post("/api/v1/users/login", userHandler.Login)
    r.Get("/api/v1/products", productHandler.ListProducts)
    r.Get("/api/v1/products/{id}", productHandler.GetProduct)
    
    // Protected routes
    r.Group(func(r chi.Router) {
        r.Use(middleware.Authentication(tokenService))
        
        r.Get("/api/v1/users/me", userHandler.GetProfile)
        r.Put("/api/v1/users/me", userHandler.UpdateProfile)
        
        // Admin only routes
        r.Group(func(r chi.Router) {
            r.Use(middleware.RequireRole("admin"))
            
            r.Post("/api/v1/products", productHandler.CreateProduct)
            r.Put("/api/v1/products/{id}", productHandler.UpdateProduct)
            r.Delete("/api/v1/products/{id}", productHandler.DeleteProduct)
        })
    })
    
    return r
}
```

---

#### 4.2 Implement Middleware
**Duration**: 3 hours

**Middleware Components**:

**1. Request ID Middleware**:
```go
func RequestID() func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            requestID := uuid.New().String()
            ctx := context.WithValue(r.Context(), requestIDKey, requestID)
            w.Header().Set("X-Request-ID", requestID)
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}
```

**2. Authentication Middleware**:
```go
func Authentication(tokenService TokenService) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            token := extractToken(r)
            if token == "" {
                writeError(w, ErrUnauthorized)
                return
            }
            
            claims, err := tokenService.ValidateToken(token)
            if err != nil {
                writeError(w, ErrInvalidToken)
                return
            }
            
            ctx := context.WithValue(r.Context(), userClaimsKey, claims)
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}
```

**3. Rate Limiting Middleware**:
```go
func RateLimit(requestsPerMinute int) func(http.Handler) http.Handler {
    limiter := rate.NewLimiter(rate.Limit(requestsPerMinute/60), requestsPerMinute)
    
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            if !limiter.Allow() {
                writeError(w, ErrRateLimitExceeded)
                return
            }
            next.ServeHTTP(w, r)
        })
    }
}
```

**4. CORS Middleware**: Use standard library or chi/gin CORS

**5. Logging Middleware**: Log all requests with duration

**6. Recovery Middleware**: Recover from panics gracefully

---

#### 4.3 Error Handling
**Duration**: 1 hour

**Standardized Error Responses**:
```go
type ErrorResponse struct {
    Error ErrorDetail `json:"error"`
}

type ErrorDetail struct {
    Code      string            `json:"code"`
    Message   string            `json:"message"`
    Details   []ValidationError `json:"details,omitempty"`
    RequestID string            `json:"request_id"`
}

type ValidationError struct {
    Field   string `json:"field"`
    Message string `json:"message"`
}
```

---

#### 4.4 Integration Testing
**Duration**: 2 hours

**Test Scenarios**:
- Complete user registration flow
- Login and authenticated requests
- Rate limiting behavior
- CORS headers
- Error responses
- Token validation

---

### Task 5: JWT Token Service (4 hours)

#### 5.1 Token Generation
**Duration**: 2 hours

**Implementation**:
```go
type TokenService struct {
    secretKey              []byte
    accessTokenDuration    time.Duration
    refreshTokenDuration   time.Duration
}

func (s *TokenService) GenerateTokens(user *User) (*AuthTokens, error) {
    accessToken, err := s.generateAccessToken(user)
    if err != nil {
        return nil, err
    }
    
    refreshToken, err := s.generateRefreshToken(user)
    if err != nil {
        return nil, err
    }
    
    return &AuthTokens{
        AccessToken:  accessToken,
        RefreshToken: refreshToken,
        ExpiresIn:    int(s.accessTokenDuration.Seconds()),
    }, nil
}
```

---

#### 5.2 Token Validation
**Duration**: 1 hour

**Implementation**:
- Parse JWT token
- Validate signature
- Check expiration
- Extract claims

---

#### 5.3 Token Refresh
**Duration**: 1 hour

**Implementation**:
- Validate refresh token
- Generate new access token
- Optionally rotate refresh token

---

### Task 6: API Documentation (6 hours)

#### 6.1 OpenAPI Specification
**Duration**: 4 hours

**Create OpenAPI 3.0 Spec**:
```yaml
openapi: 3.0.0
info:
  title: Angidi E-Commerce API
  version: 1.0.0
  description: RESTful API for e-commerce platform

servers:
  - url: http://localhost:8080/api/v1
    description: Development server

paths:
  /users/register:
    post:
      summary: Register new user
      tags:
        - Users
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/RegisterRequest'
      responses:
        '201':
          description: User created successfully
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/User'
        '400':
          description: Validation error
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'
```

**Full Documentation**:
- All endpoints documented
- Request/response schemas
- Authentication requirements
- Error responses
- Examples

---

#### 6.2 Setup Swagger UI
**Duration**: 1 hour

**Implementation**:
- Serve OpenAPI spec at `/api/docs/openapi.yaml`
- Setup Swagger UI at `/api/docs`
- Auto-generate from code annotations (optional)

---

#### 6.3 API Usage Guide
**Duration**: 1 hour

**Create Documentation**:
- Getting started guide
- Authentication flow examples
- Common use cases
- Code examples (curl, JavaScript, Go)

---

### Task 7: Frontend Integration (8 hours)

#### 7.1 Update API Client
**Duration**: 2 hours

**Enhancements**:
- Add authentication endpoints
- Add product endpoints
- Token management
- Error handling
- Request interceptors

---

#### 7.2 Authentication Pages
**Duration**: 3 hours

**Pages to Create**:
1. **Registration Page** (`/register`)
   - Registration form
   - Client-side validation
   - Error handling
   - Success redirect

2. **Login Page** (`/login`)
   - Login form
   - Remember me option
   - Forgot password link (UI only)
   - Error handling

---

#### 7.3 Product Catalog Page
**Duration**: 3 hours

**Features**:
- Product grid display
- Pagination controls
- Filter sidebar
- Sort dropdown
- Loading states
- Empty states
- Error states

---

### Task 8: Testing & Quality Assurance (8 hours)

#### 8.1 Unit Tests
**Duration**: 3 hours

**Coverage**:
- All services >90%
- All handlers >80%
- Utilities >95%

---

#### 8.2 Integration Tests
**Duration**: 3 hours

**Test Scenarios**:
- Complete user flows
- Error scenarios
- Authentication flows
- Product CRUD operations

---

#### 8.3 E2E Tests
**Duration**: 2 hours

**Playwright Tests**:
- User registration flow
- Login flow
- Browse products
- View product details

---

### Task 9: Documentation & Deployment Prep (4 hours)

#### 9.1 Update README Files
**Duration**: 1 hour

**Updates**:
- Backend README with new endpoints
- Frontend README with new pages
- API documentation link
- Environment variables

---

#### 9.2 Configuration Examples
**Duration**: 1 hour

**Provide**:
- .env.example with all required variables
- config.yaml.example
- Development setup guide

---

#### 9.3 Troubleshooting Guide
**Duration**: 1 hour

**Common Issues**:
- JWT token errors
- CORS issues
- Rate limiting
- Validation errors

---

#### 9.4 Migration Guide
**Duration**: 1 hour

**Document**:
- Changes from Phase 1
- Breaking changes (if any)
- New dependencies
- Configuration changes

---

## Testing Strategy

### Unit Testing

**Coverage Targets**:
- Service Layer: >90%
- Handler Layer: >80%
- Utility Functions: >95%

**Test Structure**:
```go
func TestUserService_Register(t *testing.T) {
    tests := []struct {
        name    string
        request RegisterRequest
        setup   func(*gomock.Controller) UserRepository
        want    *User
        wantErr error
    }{
        // Test cases
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            ctrl := gomock.NewController(t)
            defer ctrl.Finish()
            
            repo := tt.setup(ctrl)
            service := NewUserService(repo, validator, logger)
            
            got, err := service.Register(context.Background(), tt.request)
            
            if tt.wantErr != nil {
                assert.Error(t, err)
                assert.Equal(t, tt.wantErr, err)
            } else {
                assert.NoError(t, err)
                assert.NotNil(t, got)
            }
        })
    }
}
```

---

### Integration Testing

**HTTP Integration Tests**:
```go
func TestUserRegistration_Integration(t *testing.T) {
    server := setupTestServer(t)
    defer server.Close()
    
    tests := []struct {
        name       string
        payload    interface{}
        wantStatus int
        validate   func(*testing.T, *http.Response)
    }{
        {
            name: "successful registration",
            payload: map[string]string{
                "email":    "test@example.com",
                "password": "SecurePass123!",
                "name":     "Test User",
            },
            wantStatus: http.StatusCreated,
            validate: func(t *testing.T, resp *http.Response) {
                var user User
                err := json.NewDecoder(resp.Body).Decode(&user)
                require.NoError(t, err)
                assert.NotEmpty(t, user.ID)
                assert.Equal(t, "test@example.com", user.Email)
            },
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            body, _ := json.Marshal(tt.payload)
            resp, err := http.Post(
                server.URL+"/api/v1/users/register",
                "application/json",
                bytes.NewReader(body),
            )
            require.NoError(t, err)
            defer resp.Body.Close()
            
            assert.Equal(t, tt.wantStatus, resp.StatusCode)
            if tt.validate != nil {
                tt.validate(t, resp)
            }
        })
    }
}
```

---

### E2E Testing

**Playwright Test Examples**:
```typescript
test.describe('User Authentication', () => {
  test('user can register and login', async ({ page }) => {
    // Registration
    await page.goto('/register');
    await page.fill('[name="email"]', 'test@example.com');
    await page.fill('[name="password"]', 'SecurePass123!');
    await page.fill('[name="name"]', 'Test User');
    await page.click('button[type="submit"]');
    
    // Should redirect to login or dashboard
    await expect(page).toHaveURL(/\/(login|dashboard)/);
    
    // Login
    await page.goto('/login');
    await page.fill('[name="email"]', 'test@example.com');
    await page.fill('[name="password"]', 'SecurePass123!');
    await page.click('button[type="submit"]');
    
    // Should be logged in
    await expect(page).toHaveURL('/dashboard');
    await expect(page.locator('text=Test User')).toBeVisible();
  });
});

test.describe('Product Catalog', () => {
  test('user can browse products', async ({ page }) => {
    await page.goto('/products');
    
    // Products should load
    await expect(page.locator('.product-card')).toHaveCount(10);
    
    // Pagination works
    await page.click('button:has-text("Next")');
    await expect(page).toHaveURL(/page=2/);
    
    // Filter works
    await page.selectOption('[name="category"]', 'electronics');
    await expect(page.locator('.product-card')).toHaveCount.toBeGreaterThan(0);
  });
});
```

---

## Dependencies

### Backend Dependencies

#### New Dependencies for Phase 2

```go
// go.mod additions
require (
    github.com/go-chi/chi/v5 v5.0.10              // HTTP router
    github.com/go-chi/cors v1.2.1                 // CORS middleware
    github.com/golang-jwt/jwt/v5 v5.1.0           // JWT tokens
    golang.org/x/crypto v0.15.0                   // bcrypt
    github.com/go-playground/validator/v10 v10.16.0 // Validation
    github.com/google/uuid v1.4.0                 // UUID generation
    golang.org/x/time v0.4.0                      // Rate limiting
)
```

**OR if using Gin**:
```go
require (
    github.com/gin-gonic/gin v1.9.1               // HTTP framework
    github.com/golang-jwt/jwt/v5 v5.1.0           // JWT tokens
    golang.org/x/crypto v0.15.0                   // bcrypt
    github.com/go-playground/validator/v10 v10.16.0 // Validation
    github.com/google/uuid v1.4.0                 // UUID generation
)
```

#### Testing Dependencies

```go
require (
    github.com/stretchr/testify v1.8.4            // Test assertions
    github.com/golang/mock v1.6.0                 // Mocking
    go.uber.org/mock v0.3.0                       // Alternative mocking
)
```

---

### Frontend Dependencies

#### New Dependencies for Phase 2

```json
{
  "dependencies": {
    "axios": "^1.6.0",                    // HTTP client
    "zustand": "^4.4.6",                  // State management
    "react-hook-form": "^7.48.2",         // Form handling
    "zod": "^3.22.4",                     // Schema validation
    "@hookform/resolvers": "^3.3.2"       // Form validation integration
  },
  "devDependencies": {
    "@types/node": "^20.9.0",
    "@types/react": "^18.2.37",
    "typescript": "^5.2.2"
  }
}
```

---

## Security Considerations

### Phase 2 Security Checklist

- [ ] Passwords hashed with bcrypt (cost ≥12)
- [ ] JWT tokens with short expiration
- [ ] Refresh token rotation implemented
- [ ] Input validation on all endpoints
- [ ] SQL injection prevention (Phase 3)
- [ ] XSS prevention (output encoding)
- [ ] CSRF protection (Phase 3+)
- [ ] Rate limiting on authentication endpoints
- [ ] HTTPS enforcement (production)
- [ ] Security headers configured
- [ ] CORS properly configured
- [ ] Secrets not committed to version control
- [ ] Error messages don't leak sensitive info
- [ ] Logging doesn't include sensitive data

---

## Performance Considerations

### Performance Targets

**API Response Times**:
- User registration: <500ms
- User login: <300ms
- Product list: <200ms
- Product detail: <100ms

**Optimization Strategies**:
- Use efficient data structures (maps for lookups)
- Minimize allocations in hot paths
- Use connection pooling (Phase 3)
- Implement pagination for lists
- Add indexes on frequently queried fields (Phase 3)

---

## Acceptance Criteria

### Phase 2 Completion Checklist

#### User Service
- [ ] User registration endpoint working
- [ ] Email validation prevents duplicates
- [ ] Password hashing with bcrypt
- [ ] Login endpoint returns JWT tokens
- [ ] Token validation middleware working
- [ ] User profile endpoints functional
- [ ] Token refresh mechanism implemented
- [ ] Unit tests >90% coverage
- [ ] Integration tests pass

#### Product Service
- [ ] Product CRUD endpoints working
- [ ] Product listing with pagination
- [ ] Product filtering by category/price
- [ ] Product search functionality
- [ ] Input validation on all fields
- [ ] Admin-only endpoints protected
- [ ] Unit tests >90% coverage
- [ ] Integration tests pass

#### API Gateway
- [ ] Router configured with all endpoints
- [ ] Authentication middleware working
- [ ] Authorization (role-based) working
- [ ] Rate limiting functional
- [ ] CORS configured properly
- [ ] Request logging enabled
- [ ] Error handling consistent
- [ ] Recovery from panics

#### Documentation
- [ ] OpenAPI spec complete
- [ ] Swagger UI accessible
- [ ] API usage guide created
- [ ] README updated
- [ ] Environment variables documented
- [ ] Troubleshooting guide complete

#### Frontend
- [ ] Registration page functional
- [ ] Login page functional
- [ ] Product catalog page working
- [ ] Authentication state management
- [ ] Error handling implemented
- [ ] Loading states shown
- [ ] E2E tests pass

#### Quality Gates
- [ ] All linters pass
- [ ] All unit tests pass (>80% coverage)
- [ ] All integration tests pass
- [ ] All E2E tests pass
- [ ] No security vulnerabilities
- [ ] Performance targets met
- [ ] Documentation complete

---

## Troubleshooting

### Common Issues

#### Issue 1: JWT Token Validation Fails

**Symptoms**:
- 401 Unauthorized errors on protected endpoints
- "Invalid token" errors

**Solutions**:
1. Check token is being sent in Authorization header
2. Verify header format: `Authorization: Bearer <token>`
3. Ensure JWT secret matches between generation and validation
4. Check token hasn't expired
5. Verify token claims structure

```bash
# Debug JWT token
echo "your.jwt.token" | cut -d'.' -f2 | base64 -d | jq
```

---

#### Issue 2: CORS Errors

**Symptoms**:
- Browser console shows CORS errors
- Requests fail from frontend

**Solutions**:
1. Verify CORS middleware is configured
2. Check allowed origins include frontend URL
3. Ensure allowed methods include required HTTP methods
4. Add `Content-Type` and `Authorization` to allowed headers
5. Set `AllowCredentials: true` if using cookies

```go
cors.Config{
    AllowedOrigins:   []string{"http://localhost:3000"},
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
    AllowedHeaders:   []string{"Content-Type", "Authorization"},
    AllowCredentials: true,
}
```

---

#### Issue 3: Rate Limiting Too Aggressive

**Symptoms**:
- 429 Too Many Requests errors
- Legitimate requests being blocked

**Solutions**:
1. Increase rate limit threshold
2. Implement per-user rate limiting instead of global
3. Whitelist certain IPs (admin, monitoring)
4. Use sliding window instead of fixed window
5. Provide clear error messages with retry-after header

---

#### Issue 4: Password Hashing Too Slow

**Symptoms**:
- Registration/login taking >1 second
- High CPU usage during authentication

**Solutions**:
1. Reduce bcrypt cost factor (but not below 10)
2. Implement async password hashing
3. Add caching for repeated login attempts (Phase 3+)
4. Consider using hardware acceleration

```go
// Adjust cost based on environment
cost := 12  // Production
if env == "development" {
    cost = 10  // Faster for dev
}
hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
```

---

#### Issue 5: In-Memory Data Loss

**Symptoms**:
- Users/products disappear after server restart
- Data inconsistency

**Expected Behavior** (Phase 2):
- This is expected! We're using in-memory storage
- Data is meant to be lost on restart
- Phase 3 will add persistent database

**Workarounds**:
- Seed data on startup
- Use long-running server for testing
- Document this limitation clearly

---

## Migration from Phase 1

### Breaking Changes

**None** - Phase 2 is additive.

### New Environment Variables

```bash
# JWT Configuration
JWT_SECRET=your-secret-key-here  # CHANGE IN PRODUCTION!
JWT_ACCESS_TOKEN_DURATION=15m
JWT_REFRESH_TOKEN_DURATION=168h

# API Configuration
API_VERSION=v1
API_RATE_LIMIT=100  # requests per minute

# CORS Configuration
CORS_ALLOWED_ORIGINS=http://localhost:3000
```

### Dependency Updates

```bash
# Backend
cd backend
go mod tidy

# Frontend
cd frontend
npm install
```

---

## Lessons Learned

**Note**: This section will be populated upon completion of Phase 2.

### What Worked Well
- TBD

### Challenges Faced
- TBD

### Improvements for Next Phase
- TBD

---

## Next Steps (Phase 3)

Phase 3 will introduce:
- **PostgreSQL integration** for persistent storage
- **Database migrations** for schema management
- **Connection pooling** for performance
- **Repository pattern** with real database implementation
- **Transactional operations** for data consistency

**Preparation**:
- Review PostgreSQL documentation
- Learn about database migrations
- Study the Repository pattern
- Understand ACID properties

---

## References

### Internal Documentation
- [Main Specification](../../specs/SPEC.md)
- [Technology Stack](../../specs/TECH_STACK.md)
- [System Design Concepts](../../specs/SYSTEM_DESIGN_CONCEPTS.md)
- [Functional Requirements](../../specs/FUNCTIONAL_REQUIREMENTS.md)
- [Non-Functional Requirements](../../specs/NON_FUNCTIONAL_REQUIREMENTS.md)
- [Architecture Guide](../../architecture/README.md)
- [Phase 1: Repository Scaffolding](../phase-01-scaffolding/README.md)
- [All Phases Overview](../README.md)

### External Resources

#### Go Resources
- [Chi Router Documentation](https://go-chi.io/)
- [Gin Framework Documentation](https://gin-gonic.com/docs/)
- [JWT-Go Library](https://github.com/golang-jwt/jwt)
- [bcrypt Package](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
- [Validator v10](https://github.com/go-playground/validator)

#### Authentication & Security
- [JWT Introduction](https://jwt.io/introduction)
- [OWASP Authentication Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Authentication_Cheat_Sheet.html)
- [bcrypt: How It Works](https://en.wikipedia.org/wiki/Bcrypt)
- [CORS Explained](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS)

#### API Design
- [REST API Tutorial](https://restfulapi.net/)
- [OpenAPI Specification](https://swagger.io/specification/)
- [API Design Best Practices](https://docs.microsoft.com/en-us/azure/architecture/best-practices/api-design)

#### Testing
- [Go Testing Package](https://pkg.go.dev/testing)
- [Testify Documentation](https://github.com/stretchr/testify)
- [GoMock Documentation](https://github.com/golang/mock)

#### Frontend Integration
- [Axios Documentation](https://axios-http.com/)
- [React Hook Form](https://react-hook-form.com/)
- [Zod Validation](https://zod.dev/)

---

**Version**: 1.0.0  
**Last Updated**: 2025-10-27  
**Status**: ⏳ Not Started  
**Next Review**: Upon Phase 2 completion
