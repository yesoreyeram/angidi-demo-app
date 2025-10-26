# Copilot Instructions for Angidi E-Commerce Platform

## Project Overview

Angidi is an educational project for learning advanced system design concepts through building a production-grade, scalable e-commerce platform. This is a learning-focused project implementing enterprise-grade patterns incrementally across 13 phases.

## Project Goals

- **Learning First**: Teach advanced system design concepts through hands-on implementation
- **Production-Grade**: Build using enterprise patterns and best practices
- **Incremental Complexity**: Introduce 1-2 new tools/concepts per phase
- **Well-Documented**: Comprehensive documentation for every decision and pattern

## Architecture Principles

1. **Microservices Architecture**: Independent, scalable services
2. **Event-Driven Communication**: Asynchronous messaging via Kafka
3. **Database per Service**: Each service owns its data
4. **API-First Design**: Well-defined contracts between services
5. **Fault Tolerance**: Circuit breakers, retries, graceful degradation
6. **Observability**: Comprehensive logging, metrics, and tracing
7. **Security by Design**: Defense in depth, encryption, least privilege

## Technology Stack

### Backend
- **Primary Language**: Go
- **Databases**: PostgreSQL (orders, payments, users), MongoDB (products, reviews), Redis (cache, sessions), Elasticsearch (search)
- **Message Broker**: Apache Kafka
- **API Gateway**: Kong / NGINX

### Frontend
- **Framework**: React + Next.js
- **Language**: TypeScript
- **Styling**: Tailwind CSS

### Infrastructure
- **Containers**: Docker
- **Orchestration**: Kubernetes
- **CI/CD**: GitHub Actions
- **Observability**: Prometheus, Grafana, Loki, Jaeger

## Phase-Based Development

The project follows 13 phases over ~28 weeks:

1. **Phase 0**: Planning & Repository Setup (Current)
2. **Phase 1**: Basic Monolith
3. **Phase 2**: Microservice Decomposition
4. **Phase 3**: Caching Layer
5. **Phase 4**: Message Queue & Events
6. **Phase 5**: Search & Discovery
7. **Phase 6**: Advanced Database Patterns
8. **Phase 7**: Observability - Metrics & Logs
9. **Phase 8**: Distributed Tracing
10. **Phase 9**: Resilience Patterns
11. **Phase 10**: Security Hardening
12. **Phase 11**: Kubernetes Deployment
13. **Phase 12**: Performance Optimization
14. **Phase 13**: Chaos Engineering

## Coding Standards

### Go Code
- Follow standard Go conventions and formatting (`gofmt`, `golint`)
- Use meaningful variable names
- Write comprehensive tests (>80% coverage)
- Document public APIs and complex logic
- Use dependency injection for testability

### TypeScript/React Code
- Use TypeScript strict mode
- Follow React best practices and hooks patterns
- Use functional components
- Implement proper error boundaries
- Write unit tests with Jest and React Testing Library

### Testing Strategy
- **Test-Driven Development (TDD)**: Write tests before implementation
- **Behavior-Driven Development (BDD)**: Define behavior through scenarios
- **Unit Tests**: >80% coverage
- **Integration Tests**: Test service interactions
- **E2E Tests**: Critical user journeys

## Documentation Requirements

- Every design decision should be documented with rationale
- Use Architecture Decision Records (ADRs) for significant choices
- Keep README files up-to-date in each service
- Document API contracts using OpenAPI/Swagger
- Provide setup and troubleshooting guides

## Security Requirements

- Never commit secrets or credentials
- Use environment variables for configuration
- Implement input validation on all endpoints
- Follow OWASP security best practices
- Use secure dependencies (scan regularly)
- Encrypt data in transit (TLS) and at rest (AES-256)
- Implement proper authentication and authorization

## Code Review Guidelines

- Keep changes focused and small
- Write clear commit messages
- Include tests with every change
- Update documentation as needed
- No breaking changes without discussion
- Follow semantic versioning

## Specification Files

The project uses the GitHub Spec Kit structure:

- `.specify/memory/` - Project memory and context
- `.specify/scripts/` - Automation scripts
- `.specify/templates/` - Reusable templates
- `specs/` - Technical specifications
- `.github/prompts/` - AI agent prompts

## Special Instructions for Copilot

1. **Prefer `.specify/` structure**: Use the `.specify/` directory for all specification-related files
2. **Consult specs first**: Always check `specs/` directory for requirements before implementing
3. **Follow phase plan**: Respect the phase-based approach and don't introduce tools early
4. **Maintain consistency**: Follow existing patterns in the codebase
5. **Document as you go**: Update specs and docs with every change
6. **Test thoroughly**: Run all tests before committing
7. **Security conscious**: Always consider security implications

## Current Phase Context

**Active Phase**: Phase 0 - Planning & Repository Setup

**Objectives**:
- ✅ Create specification kit structure
- ✅ Define system design concepts
- ✅ Document functional and non-functional requirements
- ✅ Design high-level architecture
- ✅ Plan all phases

**Next Steps**:
- Begin Phase 1: Basic Monolith Application
- Set up Go project structure
- Implement basic REST API
- Set up PostgreSQL database
- Create simple React frontend

## Key Files to Reference

- `specs/requirements/functional.md` - What features to build
- `specs/requirements/non-functional.md` - Quality attributes and metrics
- `specs/architecture/overview.md` - System architecture design
- `specs/phases/README.md` - Phase roadmap
- `.specify/memory/project-context.md` - Current project state

## Common Patterns

### Service Structure
```
service-name/
├── cmd/           # Application entry points
├── internal/      # Private application code
│   ├── api/       # API handlers
│   ├── models/    # Data models
│   ├── service/   # Business logic
│   └── repository/ # Data access
├── pkg/           # Public libraries
├── tests/         # Tests
└── README.md      # Service documentation
```

### Error Handling
- Use structured errors with context
- Log errors with appropriate severity
- Return user-friendly error messages
- Never expose internal details to clients

### API Design
- RESTful conventions
- Use proper HTTP methods and status codes
- Version APIs (e.g., `/api/v1/`)
- Include pagination for list endpoints
- Implement rate limiting

## Quick Reference

- **Main README**: Project overview and getting started
- **Specs Directory**: All technical specifications
- **Phase Plans**: `.spec/phases/` (legacy) and `specs/phases/`
- **Scripts**: `.specify/scripts/` for automation
- **Validation**: Run `.specify/scripts/validate.sh` before commits
