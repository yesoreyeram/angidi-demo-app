# Angidi - Enterprise E-Commerce Platform

> A comprehensive learning project for advanced system design concepts through building a scalable, performant, and reliable e-commerce platform

[![Status](https://img.shields.io/badge/status-in--planning-yellow)]()
[![Phase](https://img.shields.io/badge/phase-0-blue)]()
[![License](https://img.shields.io/badge/license-MIT-green)]()

## 🎯 Project Overview

**Angidi** is an educational project designed to teach advanced system design concepts through hands-on implementation of a production-grade e-commerce platform. The project follows a phase-based approach, incrementally introducing system design patterns, architectural concepts, and modern technologies.

### What Makes This Project Unique?

- **🎓 Learning-First Approach**: Each phase introduces 1-2 new tools/concepts with detailed explanations
- **📚 Comprehensive Documentation**: Extensive specs covering every design decision and concept
- **🧪 Test-Driven Development**: TDD and BDD principles applied throughout
- **🔒 Enterprise-Grade**: Security, observability, and scalability built-in from the start
- **🏗️ Real-World Architecture**: Patterns used by companies like Amazon, Netflix, and Google
- **📊 System Design Mastery**: Covers all major distributed systems concepts

## 🌟 Goals

### Primary Goals

1. **Learn Advanced System Design Concepts** - Master scalability, reliability, consistency, observability, and more
2. **Build a Production-Ready Platform** - Create a fully functional e-commerce system
3. **Follow Best Practices** - Apply enterprise-grade security, testing, and documentation standards
4. **Understand Trade-offs** - Learn when and why to choose different architectural patterns

### Non-Goals

- This is **not** a production business venture
- This is **not** about learning basic programming (prerequisites required)
- This is **not** a copy-paste tutorial (understanding is key)

## 👥 Target Audience

This project is ideal for:

- **Software Engineers** wanting to level up their system design skills
- **Backend Developers** looking to understand distributed systems
- **Students** preparing for system design interviews
- **Technical Leads** learning architectural patterns
- **Anyone** interested in building scalable systems

### Prerequisites

**Required Knowledge**:
- Proficiency in at least one programming language (Go or JavaScript/TypeScript preferred)
- Basic understanding of REST APIs and HTTP
- Familiarity with Git and command-line tools
- Understanding of relational databases
- Basic Docker knowledge (helpful but not required)

**Required Tools**:
- Linux environment (Ubuntu 20.04+ recommended) or macOS
- Git 2.x
- Docker 20.x
- 8GB+ RAM
- 20GB+ free disk space

## 🏛️ Architecture

Angidi is designed as a **microservices-based e-commerce platform** with the following characteristics:

### High-Level Features

- **User Management**: Registration, authentication, profiles, roles
- **Product Catalog**: Flexible product listings with categories and variants
- **Search & Discovery**: Full-text search with filters and recommendations
- **Shopping Cart**: Persistent cart with real-time pricing
- **Order Management**: Complete order lifecycle tracking
- **Inventory**: Real-time stock management with overselling prevention
- **Payments**: Secure payment processing with PCI-DSS compliance
- **Notifications**: Multi-channel notifications (email, SMS, push)
- **Reviews & Ratings**: Product and seller ratings
- **Admin Panel**: Comprehensive administrative functions

### Technology Stack

#### Backend
- **Language**: Go (primary), Node.js (select services)
- **Databases**: PostgreSQL, MongoDB, Redis, Elasticsearch
- **Message Broker**: Apache Kafka
- **API Gateway**: Kong / NGINX

#### Frontend
- **Framework**: React + Next.js
- **Language**: TypeScript
- **Styling**: Tailwind CSS
- **State Management**: Redux / Zustand

#### Infrastructure
- **Containers**: Docker
- **Orchestration**: Kubernetes
- **IaC**: Terraform
- **CI/CD**: GitHub Actions
- **Cloud**: AWS / GCP (cloud-agnostic design)

#### Observability
- **Metrics**: Prometheus + Grafana
- **Logging**: Loki / ELK Stack
- **Tracing**: Jaeger / Zipkin
- **APM**: Custom dashboards

### Key Design Principles

1. **Microservices Architecture** - Independent, scalable services
2. **Event-Driven Communication** - Asynchronous messaging via Kafka
3. **Database per Service** - Each service owns its data
4. **API-First Design** - Well-defined contracts between services
5. **Fault Tolerance** - Circuit breakers, retries, graceful degradation
6. **Observability** - Comprehensive logging, metrics, and tracing
7. **Security by Design** - Defense in depth, encryption, least privilege

## 🗺️ Project Phases

The project is divided into **13 phases**, each building upon the previous:

| Phase | Name | Key Concepts | Duration |
|-------|------|--------------|----------|
| **0** | Planning & Repository Setup | Requirements, Architecture | 1 week |
| **1** | Basic Monolith | MVC, REST APIs, Database | 2 weeks |
| **2** | Microservice Decomposition | Service boundaries, Containers | 2 weeks |
| **3** | Caching Layer | Caching strategies, Redis | 1 week |
| **4** | Message Queue & Events | Event-driven architecture, Kafka | 2 weeks |
| **5** | Search & Discovery | Full-text search, Elasticsearch | 2 weeks |
| **6** | Advanced Database Patterns | Sharding, replication, CAP theorem | 2 weeks |
| **7** | Observability - Metrics & Logs | Prometheus, Grafana, Loki | 2 weeks |
| **8** | Distributed Tracing | Request tracing, Jaeger | 1 week |
| **9** | Resilience Patterns | Circuit breakers, bulkheads | 2 weeks |
| **10** | Security Hardening | OAuth, encryption, compliance | 2 weeks |
| **11** | Kubernetes Deployment | Container orchestration, scaling | 3 weeks |
| **12** | Performance Optimization | Load testing, profiling | 2 weeks |
| **13** | Chaos Engineering | Fault injection, resilience testing | 2 weeks |

**Total Duration**: ~28 weeks (7 months)

**Current Phase**: Phase 0 - Planning & Repository Setup ✅

## 📚 Documentation Structure

```
.spec/                          # Specification Kit
├── README.md                   # Spec kit overview
├── system-design/              # System design concepts
│   └── concepts.md            # All concepts with checklist
├── requirements/               # Requirements documentation
│   ├── functional.md          # Functional requirements
│   └── non-functional.md      # Non-functional requirements
├── architecture/               # Architecture documentation
│   └── overview.md            # High-level architecture
├── phases/                     # Phase-based planning
│   ├── README.md              # Phases overview
│   └── phase-X-name/          # Individual phase docs
│       ├── README.md          # Phase overview
│       ├── concepts.md        # Concepts covered
│       ├── design-decisions.md # ADRs
│       ├── setup.md           # Setup instructions
│       ├── testing-strategy.md # Test plans
│       └── checklist.md       # Completion checklist
└── scripts/                    # Automation scripts
    ├── validate.sh            # Validate spec kit
    ├── generate-checklist.sh  # Generate progress checklist
    └── new-phase.sh           # Create new phase template
```

## 🚀 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/yesoreyeram/angidi-demo-app.git
cd angidi-demo-app
```

### 2. Explore the Specification

```bash
# Read the spec kit overview
cat .spec/README.md

# Review system design concepts
cat .spec/system-design/concepts.md

# Check current phase
cat .spec/phases/phase-0-planning/README.md
```

### 3. Validate the Spec Kit

```bash
# Validate the specification structure
./.spec/scripts/validate.sh
```

### 4. Follow Phase 0

Start with Phase 0 to understand the planning and setup:

```bash
cd .spec/phases/phase-0-planning
cat README.md
```

### 5. Set Up Your Environment

As you progress through phases, each phase has a `setup.md` file with specific instructions:

```bash
# Example for future phases
cat .spec/phases/phase-1-basic-monolith/setup.md
```

## 🧪 Testing Strategy

The project follows **Test-Driven Development (TDD)** and **Behavior-Driven Development (BDD)**:

### Test Pyramid

1. **Unit Tests** (>80% coverage)
   - Test individual components
   - Fast, isolated, deterministic
   
2. **Integration Tests** (>70% coverage)
   - Test service interactions
   - Test with real databases (test instances)
   
3. **End-to-End Tests** (Critical paths)
   - Test complete user journeys
   - Test with all services running

### Running Tests

```bash
# Once implementation begins (Phase 1+)

# Run all tests
make test

# Run specific test suites
make test-unit
make test-integration
make test-e2e

# Run with coverage
make test-coverage
```

## 📈 Progress Tracking

Track your progress through the system design concepts:

```bash
# Generate progress checklist
./.spec/scripts/generate-checklist.sh

# View checklist
cat .spec/system-design/checklist.md
```

## 🔧 Development Workflow

### Branch Strategy

- `main` - Stable, completed phases
- `phase-N-name` - Work-in-progress for specific phase
- `feature/description` - Specific features within a phase

### Commit Convention

```
type(scope): subject

body

footer
```

**Types**: feat, fix, docs, style, refactor, test, chore

**Example**:
```
feat(user-service): add user registration endpoint

- Implement POST /api/users/register
- Add input validation
- Hash passwords with bcrypt
- Return JWT token

Closes #123
```

### Pull Request Process

1. Complete work in feature branch
2. Ensure all tests pass
3. Update documentation
4. Create PR with description
5. Address review comments
6. Merge after approval

## 📊 Non-Functional Requirements

The platform aims to achieve:

- **Performance**: <200ms API response time (p95)
- **Availability**: 99.99% uptime (~5 min downtime/year)
- **Scalability**: Handle 100K requests/second
- **Throughput**: Process 5M orders/day
- **Data**: Support 100M users, 10M+ products
- **Security**: PCI-DSS compliant, GDPR compliant
- **Observability**: Full logging, metrics, and tracing

See [Non-Functional Requirements](.spec/requirements/non-functional.md) for details.

## 🔒 Security

Security is a first-class concern:

- **Authentication**: JWT + OAuth 2.0
- **Authorization**: Role-based access control (RBAC)
- **Encryption**: TLS 1.3 in transit, AES-256 at rest
- **Secrets Management**: HashiCorp Vault / AWS Secrets Manager
- **Compliance**: PCI-DSS for payments, GDPR for data privacy
- **Security Scanning**: Automated vulnerability scanning in CI/CD

## 🤝 Contributing

This is primarily an educational project, but contributions are welcome:

1. **Improvements**: Better explanations, corrections, optimizations
2. **Documentation**: Enhanced docs, diagrams, examples
3. **Tests**: Additional test cases, better coverage
4. **Tools**: Alternative implementations, new tools

Please open an issue first to discuss significant changes.

## 📖 Learning Resources

### System Design
- "Designing Data-Intensive Applications" by Martin Kleppmann
- "System Design Interview" by Alex Xu
- AWS Well-Architected Framework
- Google Cloud Architecture Center

### Distributed Systems
- MIT 6.824 Distributed Systems Course
- "Distributed Systems" by Maarten van Steen & Andrew Tanenbaum

### Microservices
- "Building Microservices" by Sam Newman
- "Microservices Patterns" by Chris Richardson

### Go Programming
- "The Go Programming Language" by Donovan & Kernighan
- Go official documentation

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

This project draws inspiration from:

- Real-world e-commerce platforms (Amazon, Shopify, eBay)
- System design courses and books
- Open-source projects demonstrating best practices
- The developer community's collective knowledge

## 📞 Contact & Support

- **Issues**: [GitHub Issues](https://github.com/yesoreyeram/angidi-demo-app/issues)
- **Discussions**: [GitHub Discussions](https://github.com/yesoreyeram/angidi-demo-app/discussions)

---

**Current Status**: Phase 0 - Planning & Repository Setup 🚧

**Last Updated**: 2025-10-26

Happy learning! 🚀
