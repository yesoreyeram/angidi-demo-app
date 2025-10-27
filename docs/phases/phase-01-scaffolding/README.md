# Phase 1: Repository Scaffolding

**Status**: ✅ Completed  
**Duration**: 1 week  
**Start Date**: 2025-10-26  
**Completion Date**: 2025-10-26

---

## Overview

Phase 1 establishes the foundational project structure, development tooling, and CI/CD infrastructure for both backend and frontend applications. This phase focuses on creating a well-organized, maintainable codebase with automated quality checks and testing frameworks.

---

## Goals & Objectives

### Primary Goals
1. Initialize Go backend project with idiomatic structure
2. Set up React/Next.js frontend application
3. Implement comprehensive linting and code formatting
4. Establish CI/CD pipeline with GitHub Actions
5. Define and enforce coding standards
6. Set up E2E testing infrastructure
7. Create developer documentation and setup guides

### Success Criteria
- ✅ Go backend compiles and passes all linters
- ✅ Next.js frontend builds and runs in development mode
- ✅ All CI checks pass (build, lint, test)
- ✅ E2E test infrastructure is operational
- ✅ Comprehensive setup documentation exists
- ✅ Code formatting is automated
- ✅ Pre-commit hooks are configured

---

## System Design Concepts Introduced

### Project Organization
- **Clean Architecture**: Separation of concerns with clear boundaries
- **Hexagonal Architecture**: Domain-centric design with ports and adapters
- **Package Structure**: Idiomatic Go project layout
- **Module Organization**: Frontend component and feature organization

### Development Practices
- **Continuous Integration**: Automated build and test on every commit
- **Code Quality Automation**: Linting, formatting, and static analysis
- **Test-Driven Development**: Testing infrastructure from day one
- **Developer Experience**: Fast feedback loops and easy onboarding

### Infrastructure as Code
- **Configuration as Code**: YAML-based CI/CD pipelines
- **Reproducible Builds**: Consistent build environments
- **Automated Quality Gates**: Enforce standards automatically

---

## Tools & Technologies

### Backend Stack
- **Go 1.21+**: Backend programming language
- **Go Modules**: Dependency management
- **golangci-lint**: Comprehensive Go linting suite
- **gofmt/goimports**: Code formatting
- **Make**: Build automation

### Frontend Stack
- **Node.js 20 LTS**: JavaScript runtime
- **React 18+**: UI library
- **TypeScript 5+**: Type-safe JavaScript
- **Next.js 14+**: React framework with SSR/SSG
- **Tailwind CSS 3+**: Utility-first CSS framework
- **ESLint**: JavaScript/TypeScript linting
- **Prettier**: Code formatting
- **pnpm**: Fast, disk-efficient package manager

### Development Tools
- **GitHub Actions**: CI/CD automation
- **Playwright**: E2E testing framework
- **Husky**: Git hooks management
- **EditorConfig**: Consistent editor settings

### Testing Frameworks
- **Go Testing Package**: Built-in unit testing
- **testify**: Enhanced assertions for Go
- **Playwright**: Browser automation for E2E tests
- **Jest** (optional): JavaScript unit testing

---

## Prerequisites

### Knowledge Requirements
- Basic understanding of Go programming
- Familiarity with React and TypeScript
- Experience with Git version control
- Understanding of CI/CD concepts

### Environment Setup
- Git 2.40+
- Go 1.21+
- Node.js 20 LTS
- pnpm 8+ (or npm 10+)
- Code editor (VS Code recommended)

### Recommended VS Code Extensions
- Go (golang.go)
- Prettier (esbenp.prettier-vscode)
- ESLint (dbaeumer.vscode-eslint)
- Tailwind CSS IntelliSense (bradlc.vscode-tailwindcss)
- EditorConfig (editorconfig.editorconfig)

---

## Project Structure

### Backend Structure

```
backend/
├── cmd/
│   ├── api/              # API server entry point
│   │   └── main.go
│   ├── worker/           # Background worker entry point
│   │   └── main.go
│   └── migration/        # Database migration tool
│       └── main.go
├── internal/             # Private application code
│   ├── user/             # User domain
│   │   ├── handler.go    # HTTP handlers
│   │   ├── service.go    # Business logic
│   │   ├── repository.go # Data access
│   │   └── model.go      # Domain models
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
│   ├── local.yaml
│   ├── dev.yaml
│   └── prod.yaml
├── .golangci.yml         # Linter configuration
├── Makefile              # Build automation
├── go.mod                # Go module definition
├── go.sum                # Dependency checksums
└── README.md             # Backend documentation
```

### Frontend Structure

```
frontend/
├── src/
│   ├── app/              # Next.js app directory
│   │   ├── layout.tsx    # Root layout
│   │   ├── page.tsx      # Home page
│   │   ├── globals.css   # Global styles
│   │   ├── (auth)/       # Auth route group
│   │   │   ├── login/
│   │   │   └── register/
│   │   ├── products/     # Products pages
│   │   ├── cart/         # Cart page
│   │   └── checkout/     # Checkout pages
│   ├── components/       # Reusable components
│   │   ├── ui/           # Base UI components
│   │   ├── layout/       # Layout components
│   │   └── features/     # Feature-specific components
│   ├── lib/              # Utility functions
│   │   ├── api/          # API client
│   │   ├── hooks/        # Custom React hooks
│   │   ├── utils/        # Helper functions
│   │   └── types/        # TypeScript types
│   ├── config/           # Configuration
│   └── styles/           # Additional styles
├── public/               # Static assets
│   ├── images/
│   └── icons/
├── tests/                # E2E tests
│   ├── e2e/
│   └── playwright.config.ts
├── .eslintrc.json        # ESLint configuration
├── .prettierrc           # Prettier configuration
├── tailwind.config.ts    # Tailwind configuration
├── next.config.js        # Next.js configuration
├── tsconfig.json         # TypeScript configuration
├── package.json          # Dependencies
├── pnpm-lock.yaml        # Dependency lock file
└── README.md             # Frontend documentation
```

### Root Structure

```
angidi-demo-app/
├── backend/              # Go backend
├── frontend/             # Next.js frontend
├── docs/                 # Documentation
├── .github/              # GitHub configuration
│   ├── workflows/        # GitHub Actions workflows
│   │   ├── backend-ci.yml
│   │   ├── frontend-ci.yml
│   │   └── e2e-tests.yml
│   └── PULL_REQUEST_TEMPLATE.md
├── .editorconfig         # Editor configuration
├── .gitignore            # Git ignore rules
├── docker-compose.yml    # Local development (Phase 4)
├── LICENSE               # MIT License
└── README.md             # Project overview
```

---

## Implementation Plan

### Task 1: Backend Scaffolding

#### 1.1 Initialize Go Module
**Duration**: 30 minutes

**Steps**:
1. Create `backend/` directory
2. Initialize Go module: `go mod init github.com/yesoreyeram/angidi-demo-app/backend`
3. Create directory structure (cmd, internal, pkg)
4. Create `.gitignore` for Go artifacts

**Deliverables**:
- `backend/go.mod` file
- Directory structure in place
- `.gitignore` configured

**Acceptance Criteria**:
- `go mod tidy` runs successfully
- Directory structure matches design
- Go module path is correct

---

#### 1.2 Setup Backend Build System
**Duration**: 1 hour

**Steps**:
1. Create `Makefile` with common tasks:
   - `make build`: Build all binaries
   - `make test`: Run all tests
   - `make lint`: Run linters
   - `make fmt`: Format code
   - `make run`: Run development server
   - `make clean`: Clean build artifacts
2. Add build scripts in `scripts/` directory
3. Configure build output directories

**Deliverables**:
- Comprehensive Makefile
- Build scripts
- Documented make targets

**Acceptance Criteria**:
- All make targets work correctly
- Help text available (`make help`)
- Cross-platform compatibility

---

#### 1.3 Configure Go Linting
**Duration**: 1 hour

**Steps**:
1. Install golangci-lint
2. Create `.golangci.yml` configuration
3. Enable recommended linters:
   - gofmt, goimports (formatting)
   - govet (correctness)
   - staticcheck (bugs)
   - gosec (security)
   - errcheck (error handling)
   - unused (dead code)
   - ineffassign (inefficiencies)
4. Configure linter rules and exclusions
5. Add `make lint` target

**Deliverables**:
- `.golangci.yml` configuration
- Linting integrated into Makefile
- Documentation of linting rules

**Acceptance Criteria**:
- Linters run without errors
- Configuration is well-documented
- Fast execution (<30 seconds on empty project)

---

#### 1.4 Create Basic Backend Structure
**Duration**: 2 hours

**Steps**:
1. Create entry point in `cmd/api/main.go`:
   - Parse command-line flags
   - Load configuration
   - Initialize logger
   - Setup HTTP server
   - Graceful shutdown
2. Implement configuration package in `pkg/config/`:
   - YAML configuration support
   - Environment variable overrides
   - Validation
3. Implement logging package in `pkg/logger/`:
   - Structured logging with levels
   - Context-aware logging
   - JSON output for production
4. Create health check endpoint
5. Add basic HTTP server setup

**Deliverables**:
- Working API server skeleton
- Configuration management
- Structured logging
- Health check endpoint

**Acceptance Criteria**:
- Server starts and responds to requests
- Configuration loads correctly
- Logs are structured and readable
- Health check returns 200 OK
- Graceful shutdown works

---

#### 1.5 Add Backend Unit Tests
**Duration**: 2 hours

**Steps**:
1. Create test structure matching package structure
2. Write unit tests for:
   - Configuration loading
   - Logger initialization
   - HTTP server setup
   - Health check endpoint
3. Setup test helpers and fixtures
4. Configure test coverage reporting
5. Add `make test-coverage` target

**Deliverables**:
- Unit tests for all packages
- Test coverage reporting
- Test documentation

**Acceptance Criteria**:
- All tests pass
- Code coverage >80%
- Tests run in <5 seconds
- Coverage report is generated

---

### Task 2: Frontend Scaffolding

#### 2.1 Initialize Next.js Application
**Duration**: 1 hour

**Steps**:
1. Create Next.js app with TypeScript:
   ```bash
   pnpm create next-app@latest frontend --typescript --tailwind --app --src-dir
   ```
2. Configure TypeScript strict mode
3. Setup path aliases in `tsconfig.json`
4. Remove default boilerplate
5. Create basic folder structure

**Deliverables**:
- Next.js application initialized
- TypeScript configured
- Path aliases setup
- Clean slate structure

**Acceptance Criteria**:
- `pnpm dev` starts development server
- TypeScript compilation works
- No TypeScript errors
- App runs on http://localhost:3000

---

#### 2.2 Configure Frontend Tooling
**Duration**: 2 hours

**Steps**:
1. Install and configure ESLint:
   - Extend recommended configs
   - Add React/Next.js rules
   - Add TypeScript rules
   - Add accessibility rules
2. Install and configure Prettier:
   - Set formatting rules
   - Configure integration with ESLint
3. Configure Tailwind CSS:
   - Custom theme colors
   - Custom spacing
   - Plugin configuration
4. Setup path aliases
5. Add `package.json` scripts:
   - `lint`: Run ESLint
   - `format`: Run Prettier
   - `type-check`: Run TypeScript compiler

**Deliverables**:
- ESLint configuration
- Prettier configuration
- Tailwind configuration
- npm scripts

**Acceptance Criteria**:
- Linting runs without errors
- Formatting is consistent
- Type checking passes
- Tailwind classes work

---

#### 2.3 Create Base Layout and Components
**Duration**: 3 hours

**Steps**:
1. Create root layout (`app/layout.tsx`):
   - HTML structure
   - Metadata configuration
   - Font setup
2. Create basic UI components:
   - Button
   - Input
   - Card
   - Container
   - Header
   - Footer
3. Implement layout components:
   - Navigation
   - Sidebar
   - Page wrapper
4. Add loading states
5. Create 404 and error pages

**Deliverables**:
- Root layout
- Base UI components
- Layout components
- Error pages

**Acceptance Criteria**:
- Components are typed correctly
- Components are accessible
- Layout is responsive
- Error pages work

---

#### 2.4 Setup Frontend API Client
**Duration**: 2 hours

**Steps**:
1. Create API client in `lib/api/`:
   - Base HTTP client
   - Request/response interceptors
   - Error handling
   - Type definitions
2. Create API endpoint configurations
3. Add environment variable support
4. Implement basic CRUD operations
5. Add request/response logging (dev only)

**Deliverables**:
- API client implementation
- Type definitions
- Configuration
- Documentation

**Acceptance Criteria**:
- API client is typed
- Error handling works
- Environment variables load correctly
- Logging is conditional

---

### Task 3: CI/CD Pipeline

#### 3.1 Setup Backend CI Pipeline
**Duration**: 2 hours

**Steps**:
1. Create `.github/workflows/backend-ci.yml`
2. Configure workflow triggers:
   - Push to main/develop
   - Pull requests
   - Manual dispatch
3. Define jobs:
   - **Build**: Compile Go code
   - **Lint**: Run golangci-lint
   - **Test**: Run unit tests with coverage
   - **Security**: Run gosec security scanner
4. Setup Go caching for faster builds
5. Configure test result reporting
6. Add coverage reporting

**Deliverables**:
- Backend CI workflow
- Test reporting
- Coverage reporting
- Security scanning

**Acceptance Criteria**:
- All jobs pass on clean code
- Build time <2 minutes
- Test coverage is reported
- Security issues are detected

---

#### 3.2 Setup Frontend CI Pipeline
**Duration**: 2 hours

**Steps**:
1. Create `.github/workflows/frontend-ci.yml`
2. Configure workflow triggers:
   - Push to main/develop
   - Pull requests
   - Manual dispatch
3. Define jobs:
   - **Install**: Install dependencies with caching
   - **Lint**: Run ESLint and type checking
   - **Build**: Build Next.js application
   - **Test**: Run unit tests (if any)
4. Setup Node.js caching for faster builds
5. Configure build artifact upload
6. Add bundle size reporting

**Deliverables**:
- Frontend CI workflow
- Dependency caching
- Build artifacts
- Bundle size tracking

**Acceptance Criteria**:
- All jobs pass on clean code
- Build time <3 minutes
- Build artifacts are uploaded
- Bundle size is tracked

---

#### 3.3 Setup E2E Testing Pipeline
**Duration**: 3 hours

**Steps**:
1. Install Playwright:
   ```bash
   pnpm create playwright
   ```
2. Configure Playwright for:
   - Multiple browsers (Chromium, Firefox, WebKit)
   - Parallel execution
   - Video recording on failure
   - Screenshot on failure
3. Create `.github/workflows/e2e-tests.yml`:
   - Start backend server
   - Start frontend server
   - Wait for servers to be ready
   - Run Playwright tests
   - Upload test results and artifacts
4. Create sample E2E tests:
   - Homepage loads
   - Navigation works
   - Basic user flows
5. Configure test reports

**Deliverables**:
- Playwright configuration
- E2E CI workflow
- Sample E2E tests
- Test reporting

**Acceptance Criteria**:
- E2E tests run in CI
- Tests pass on clean code
- Artifacts (screenshots, videos) are uploaded
- Test results are reported
- Multiple browsers are tested

---

#### 3.4 Add Pre-commit Hooks
**Duration**: 1 hour

**Steps**:
1. Install Husky:
   ```bash
   pnpm add -D husky lint-staged
   ```
2. Configure pre-commit hooks:
   - Run linters on staged files
   - Run formatters on staged files
   - Run type checking
   - Prevent commits to main branch
3. Setup commit message linting (optional)
4. Document hook behavior

**Deliverables**:
- Husky configuration
- Pre-commit hooks
- Documentation

**Acceptance Criteria**:
- Hooks run on commit
- Hooks can be bypassed if needed
- Hooks are fast (<10 seconds)
- Documentation is clear

---

### Task 4: Documentation

#### 4.1 Create Backend README
**Duration**: 1 hour

**Steps**:
1. Create `backend/README.md` with:
   - Project overview
   - Prerequisites
   - Installation instructions
   - Development workflow
   - Available make targets
   - Testing instructions
   - Project structure explanation
   - Contributing guidelines
   - Troubleshooting section

**Deliverables**:
- Comprehensive backend README

**Acceptance Criteria**:
- New developers can get started
- All commands are documented
- Troubleshooting covers common issues

---

#### 4.2 Create Frontend README
**Duration**: 1 hour

**Steps**:
1. Create `frontend/README.md` with:
   - Project overview
   - Prerequisites
   - Installation instructions
   - Development workflow
   - Available npm/pnpm scripts
   - Environment variables
   - Component development guide
   - Testing instructions
   - Build and deployment
   - Troubleshooting section

**Deliverables**:
- Comprehensive frontend README

**Acceptance Criteria**:
- New developers can get started
- All scripts are documented
- Environment setup is clear

---

#### 4.3 Update Root README
**Duration**: 30 minutes

**Steps**:
1. Update main `README.md`:
   - Quick start guide for Phase 1
   - Link to backend and frontend READMEs
   - Update badges (build status)
   - Update progress indicators
   - Add Phase 1 completion status

**Deliverables**:
- Updated root README

**Acceptance Criteria**:
- Quick start works
- Links are correct
- Badges are accurate

---

#### 4.4 Create Development Guide
**Duration**: 2 hours

**Steps**:
1. Create `docs/DEVELOPMENT.md` with:
   - Development environment setup
   - IDE configuration
   - Code style guide
   - Git workflow
   - Branch naming conventions
   - PR guidelines
   - Code review checklist
   - Debugging tips
   - Performance profiling

**Deliverables**:
- Comprehensive development guide

**Acceptance Criteria**:
- Covers all development aspects
- Includes examples
- References are accurate

---

### Task 5: E2E Testing Infrastructure

#### 5.1 Setup E2E Testing Framework
**Duration**: 3 hours

**Steps**:
1. Create E2E test structure:
   ```
   frontend/tests/e2e/
   ├── fixtures/        # Test data and helpers
   ├── pages/           # Page object models
   ├── specs/           # Test specifications
   └── utils/           # Utility functions
   ```
2. Implement page object models for:
   - Homepage
   - Navigation
   - Common interactions
3. Create test fixtures and helpers:
   - Test data generators
   - API mocking utilities
   - Authentication helpers
4. Write initial E2E tests:
   - **Home Page Tests**:
     - Page loads successfully
     - Navigation is visible
     - Footer is visible
   - **Navigation Tests**:
     - Can navigate to different pages
     - Active page is highlighted
     - Mobile menu works
5. Configure test parallelization
6. Setup test reporting

**Deliverables**:
- E2E test framework
- Page object models
- Initial E2E tests
- Test utilities
- Test documentation

**Acceptance Criteria**:
- Tests run successfully locally
- Tests run in CI/CD pipeline
- Page objects are reusable
- Test coverage for critical paths
- Tests are maintainable and readable

---

#### 5.2 Add Integration Testing Setup
**Duration**: 2 hours

**Steps**:
1. Create integration test structure in `backend/tests/integration/`
2. Setup test database configuration
3. Create test helpers:
   - HTTP request helpers
   - Database cleanup utilities
   - Test server setup
4. Write sample integration tests:
   - Health check endpoint
   - API server startup/shutdown
   - Configuration loading
5. Configure integration test running in CI

**Deliverables**:
- Integration test framework
- Test helpers
- Sample integration tests
- CI integration

**Acceptance Criteria**:
- Integration tests run locally
- Integration tests run in CI
- Tests are isolated
- Fast execution (<30 seconds)

---

## Testing Strategy

### Unit Testing

#### Backend Unit Tests
**Scope**:
- Configuration loading and validation
- Logger initialization and output
- Utility functions
- Business logic (when added in later phases)

**Tools**: Go testing package, testify

**Coverage Target**: >80% for all packages

**Location**: `*_test.go` files alongside implementation

**Execution**:
```bash
make test
make test-coverage
```

---

#### Frontend Unit Tests
**Scope**:
- Utility functions
- Custom hooks
- API client functions
- Component logic (when complex)

**Tools**: Jest, React Testing Library (to be added)

**Coverage Target**: >70% for utilities and hooks

**Location**: `__tests__/` directories or `*.test.ts(x)` files

**Execution**:
```bash
pnpm test
pnpm test:coverage
```

---

### Integration Testing

#### Backend Integration Tests
**Scope**:
- HTTP server functionality
- Endpoint integration
- Configuration integration
- Health check endpoints

**Tools**: Go testing package, httptest

**Location**: `backend/tests/integration/`

**Execution**:
```bash
make test-integration
```

---

### End-to-End Testing

#### E2E Test Scenarios (Phase 1)

**1. Homepage Loading**
```typescript
test('homepage loads successfully', async ({ page }) => {
  await page.goto('/');
  await expect(page).toHaveTitle(/Angidi/);
  await expect(page.locator('nav')).toBeVisible();
  await expect(page.locator('footer')).toBeVisible();
});
```

**2. Navigation**
```typescript
test('navigation links work', async ({ page }) => {
  await page.goto('/');
  await page.click('nav >> text=Products');
  await expect(page).toHaveURL(/products/);
  await expect(page.locator('nav >> text=Products')).toHaveClass(/active/);
});
```

**3. Responsive Design**
```typescript
test('mobile menu works on small screens', async ({ page }) => {
  await page.setViewportSize({ width: 375, height: 667 });
  await page.goto('/');
  await page.click('[aria-label="Menu"]');
  await expect(page.locator('nav')).toBeVisible();
});
```

**4. Accessibility**
```typescript
test('keyboard navigation works', async ({ page }) => {
  await page.goto('/');
  await page.keyboard.press('Tab');
  await page.keyboard.press('Enter');
  // Verify navigation occurred
});
```

**Tools**: Playwright

**Browsers**: Chromium, Firefox, WebKit

**Location**: `frontend/tests/e2e/`

**Execution**:
```bash
pnpm test:e2e
pnpm test:e2e:ui  # Run with UI mode
```

---

### Performance Testing
**Note**: Performance testing will be added in Phase 15. Phase 1 establishes baseline metrics.

**Baseline Metrics to Collect**:
- Backend server startup time
- Frontend build time
- Page load time
- Bundle size

---

## Configuration Management

### Backend Configuration

#### Configuration Files
**Location**: `backend/configs/`

**Files**:
- `local.yaml`: Local development settings
- `dev.yaml`: Development environment
- `test.yaml`: Testing environment
- `prod.yaml`: Production environment (template)

**Structure**:
```yaml
server:
  host: "localhost"
  port: 8080
  read_timeout: 10s
  write_timeout: 10s
  shutdown_timeout: 30s

logging:
  level: "info"
  format: "json"
  output: "stdout"

cors:
  allowed_origins:
    - "http://localhost:3000"
  allowed_methods:
    - "GET"
    - "POST"
    - "PUT"
    - "DELETE"
  allowed_headers:
    - "Content-Type"
    - "Authorization"
```

---

### Frontend Configuration

#### Environment Variables
**Files**:
- `.env.local`: Local development (git-ignored)
- `.env.development`: Development defaults
- `.env.production`: Production defaults
- `.env.example`: Template for required variables

**Variables**:
```bash
NEXT_PUBLIC_API_URL=http://localhost:8080
NEXT_PUBLIC_APP_NAME="Angidi"
NEXT_PUBLIC_APP_VERSION="0.1.0"
```

---

## Code Quality Standards

### Go Code Standards

#### Style Guide
- Follow [Effective Go](https://golang.org/doc/effective_go.html)
- Use `gofmt` for formatting (via `make fmt`)
- Use descriptive variable and function names
- Write documentation comments for exported items
- Keep functions small and focused
- Prefer composition over inheritance

#### Naming Conventions
- **Packages**: lowercase, single word
- **Files**: lowercase with underscores (e.g., `user_handler.go`)
- **Functions/Methods**: CamelCase (exported), camelCase (unexported)
- **Constants**: CamelCase or SCREAMING_SNAKE_CASE
- **Variables**: camelCase

#### Error Handling
- Always handle errors explicitly
- Use custom error types when appropriate
- Wrap errors with context: `fmt.Errorf("context: %w", err)`
- Log errors with appropriate levels

#### Testing
- Test file naming: `*_test.go`
- Test function naming: `TestFunctionName`
- Use table-driven tests when applicable
- Mock external dependencies
- Aim for >80% coverage

---

### TypeScript/React Code Standards

#### Style Guide
- Use TypeScript strict mode
- Use functional components with hooks
- Use ESLint and Prettier for consistency
- Write self-documenting code with good naming
- Keep components small and focused
- Separate concerns (UI, logic, data)

#### Naming Conventions
- **Files**: kebab-case (e.g., `user-profile.tsx`)
- **Components**: PascalCase (e.g., `UserProfile`)
- **Functions/Variables**: camelCase
- **Constants**: SCREAMING_SNAKE_CASE
- **Types/Interfaces**: PascalCase with descriptive names

#### Component Structure
```typescript
// 1. Imports
import React from 'react';
import { SomeType } from '@/lib/types';

// 2. Types/Interfaces
interface UserProfileProps {
  userId: string;
  onUpdate?: () => void;
}

// 3. Component
export function UserProfile({ userId, onUpdate }: UserProfileProps) {
  // Hooks
  const [user, setUser] = React.useState<User | null>(null);
  
  // Effects
  React.useEffect(() => {
    // ...
  }, [userId]);
  
  // Handlers
  const handleUpdate = () => {
    // ...
    onUpdate?.();
  };
  
  // Render
  return (
    <div>
      {/* JSX */}
    </div>
  );
}
```

#### Testing
- Test file naming: `*.test.ts(x)`
- Use React Testing Library for component tests
- Test user interactions, not implementation
- Aim for >70% coverage on critical paths

---

### Git Workflow

#### Branch Naming
- **Feature**: `feature/short-description`
- **Bugfix**: `bugfix/issue-number-description`
- **Hotfix**: `hotfix/issue-number-description`
- **Release**: `release/version-number`

#### Commit Messages
Follow Conventional Commits:
```
type(scope): subject

body (optional)

footer (optional)
```

**Types**: feat, fix, docs, style, refactor, test, chore

**Examples**:
```
feat(backend): add health check endpoint
fix(frontend): resolve navigation menu overflow
docs(readme): update installation instructions
```

#### Pull Request Process
1. Create feature branch from `main`
2. Make small, focused commits
3. Write descriptive commit messages
4. Push branch and create PR
5. Fill out PR template
6. Request review from team
7. Address review comments
8. Squash merge to main (optional)

---

## Acceptance Criteria

### Phase 1 Completion Checklist

#### Backend Infrastructure
- [ ] Go module initialized and configured
- [ ] Project structure follows best practices
- [ ] Makefile with all essential targets
- [ ] golangci-lint configured and passing
- [ ] Basic API server with health check endpoint
- [ ] Configuration management implemented
- [ ] Structured logging implemented
- [ ] Unit tests with >80% coverage
- [ ] Integration test framework setup
- [ ] Backend README complete

#### Frontend Infrastructure
- [ ] Next.js application initialized
- [ ] TypeScript strict mode enabled
- [ ] ESLint and Prettier configured
- [ ] Tailwind CSS configured with custom theme
- [ ] Base layout and components created
- [ ] API client implemented
- [ ] Path aliases configured
- [ ] Frontend README complete

#### CI/CD Pipeline
- [ ] Backend CI workflow operational
- [ ] Frontend CI workflow operational
- [ ] E2E testing workflow operational
- [ ] Pre-commit hooks configured
- [ ] All CI checks passing
- [ ] Build artifacts generated
- [ ] Test coverage reported

#### E2E Testing
- [ ] Playwright installed and configured
- [ ] Page object models created
- [ ] Homepage E2E tests passing
- [ ] Navigation E2E tests passing
- [ ] Responsive design tests passing
- [ ] Accessibility tests passing
- [ ] E2E tests run in CI/CD
- [ ] Test reports generated

#### Documentation
- [ ] Backend README comprehensive
- [ ] Frontend README comprehensive
- [ ] Root README updated
- [ ] Development guide created
- [ ] Code standards documented
- [ ] Git workflow documented
- [ ] Troubleshooting guides complete

#### Quality Gates
- [ ] All linters pass with no errors
- [ ] All tests pass
- [ ] Code coverage meets targets (>80% backend, >70% frontend)
- [ ] No security vulnerabilities detected
- [ ] Documentation is complete and accurate
- [ ] Quick start guide works for new developers

---

## Troubleshooting

### Common Issues

#### Issue 1: Go Module Download Fails
**Symptom**: `go mod download` fails with network errors

**Solution**:
```bash
# Set Go proxy
export GOPROXY=https://proxy.golang.org,direct

# Clear module cache
go clean -modcache

# Retry
go mod download
```

---

#### Issue 2: Node Modules Installation Fails
**Symptom**: `pnpm install` fails or hangs

**Solution**:
```bash
# Clear pnpm cache
pnpm store prune

# Remove node_modules and lock file
rm -rf node_modules pnpm-lock.yaml

# Retry
pnpm install
```

---

#### Issue 3: Linter Fails with Unfamiliar Errors
**Symptom**: golangci-lint reports unexpected errors

**Solution**:
```bash
# Update golangci-lint
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run with verbose output
golangci-lint run -v

# Check configuration
golangci-lint linters
```

---

#### Issue 4: E2E Tests Fail in CI but Pass Locally
**Symptom**: Playwright tests fail in GitHub Actions

**Solution**:
- Check CI logs for timing issues
- Increase timeout values
- Add explicit waits for elements
- Verify environment variables
- Check browser versions

---

#### Issue 5: Pre-commit Hooks Don't Run
**Symptom**: Husky hooks not executing

**Solution**:
```bash
# Reinstall Husky
rm -rf .husky
pnpm prepare

# Verify installation
git commit --dry-run
```

---

## Performance Metrics

### Build Performance

#### Backend
- **Initial Build Time**: Target <30 seconds
- **Incremental Build Time**: Target <5 seconds
- **Lint Time**: Target <10 seconds
- **Test Time**: Target <5 seconds

#### Frontend
- **Install Time**: Target <60 seconds (with cache)
- **Build Time**: Target <60 seconds
- **Lint Time**: Target <10 seconds
- **Type Check Time**: Target <15 seconds

#### E2E Tests
- **Total Test Time**: Target <3 minutes (all browsers)
- **Single Test Time**: Target <10 seconds
- **Parallel Execution**: 4 workers

---

### Runtime Performance

#### Backend
- **Startup Time**: Target <1 second
- **Memory Usage (Idle)**: Target <50 MB
- **Health Check Response**: Target <10ms

#### Frontend
- **Page Load Time (Dev)**: Target <2 seconds
- **Page Load Time (Prod)**: Target <1 second
- **Bundle Size**: Target <200 KB (gzipped)
- **Lighthouse Score**: Target >90

---

## Dependencies

### Backend Dependencies

#### Direct Dependencies
```go
// go.mod (minimal dependencies for Phase 1)
module github.com/yesoreyeram/angidi-demo-app/backend

go 1.21

require (
    github.com/spf13/viper v1.18.0     // Configuration
    go.uber.org/zap v1.26.0            // Logging
    gopkg.in/yaml.v3 v3.0.1            // YAML parsing
)
```

#### Development Dependencies
- golangci-lint (v1.55+)
- testify (for enhanced test assertions)

---

### Frontend Dependencies

#### Direct Dependencies
```json
{
  "dependencies": {
    "next": "^14.0.0",
    "react": "^18.2.0",
    "react-dom": "^18.2.0",
    "tailwindcss": "^3.3.0"
  }
}
```

#### Development Dependencies
```json
{
  "devDependencies": {
    "@types/node": "^20.0.0",
    "@types/react": "^18.2.0",
    "@types/react-dom": "^18.2.0",
    "typescript": "^5.3.0",
    "eslint": "^8.54.0",
    "eslint-config-next": "^14.0.0",
    "prettier": "^3.1.0",
    "playwright": "^1.40.0",
    "husky": "^8.0.0",
    "lint-staged": "^15.1.0"
  }
}
```

---

## Security Considerations

### Phase 1 Security Measures

#### Code Security
- **Static Analysis**: gosec for Go, ESLint security plugins
- **Dependency Scanning**: Dependabot enabled
- **Secret Scanning**: GitHub secret scanning enabled
- **Code Review**: Required for all PRs

#### CI/CD Security
- **Secrets Management**: GitHub Actions secrets
- **Token Permissions**: Minimal required permissions
- **Branch Protection**: Require CI to pass before merge
- **Signed Commits**: Recommended but optional in Phase 1

#### Development Security
- **Git Ignore**: Exclude sensitive files
- **Environment Files**: Template-based with `.env.example`
- **Logging**: No sensitive data in logs
- **Error Messages**: Generic errors in production

---

## Lessons Learned

**Note**: This section will be populated upon completion of Phase 1.

### What Worked Well
- TBD

### Challenges Faced
- TBD

### Improvements for Next Phase
- TBD

---

## References

### Internal Documentation
- [Main Specification](../../specs/SPEC.md)
- [Technology Stack](../../specs/TECH_STACK.md)
- [System Design Concepts](../../specs/SYSTEM_DESIGN_CONCEPTS.md)
- [Functional Requirements](../../specs/FUNCTIONAL_REQUIREMENTS.md)
- [Non-Functional Requirements](../../specs/NON_FUNCTIONAL_REQUIREMENTS.md)
- [Architecture Guide](../../architecture/README.md)
- [All Phases Overview](../README.md)

### External Resources

#### Go Resources
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)

#### React/Next.js Resources
- [Next.js Documentation](https://nextjs.org/docs)
- [React Documentation](https://react.dev)
- [TypeScript Handbook](https://www.typescriptlang.org/docs/handbook/intro.html)
- [Tailwind CSS Documentation](https://tailwindcss.com/docs)

#### Testing Resources
- [Playwright Documentation](https://playwright.dev)
- [Go Testing Package](https://pkg.go.dev/testing)
- [React Testing Library](https://testing-library.com/react)

#### CI/CD Resources
- [GitHub Actions Documentation](https://docs.github.com/en/actions)
- [golangci-lint Configuration](https://golangci-lint.run/usage/configuration/)

---

**Version**: 1.0.0  
**Last Updated**: 2025-10-27  
**Status**: ✅ Completed  
**Next Review**: N/A
