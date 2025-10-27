# Angidi - E-Commerce Platform

> A comprehensive learning project for advanced system design concepts through building a highly scalable, production-grade e-commerce platform.

[![Status](https://img.shields.io/badge/status-in--planning-blue)]()
[![Phase](https://img.shields.io/badge/phase-0--planning-green)]()
[![License](https://img.shields.io/badge/license-MIT-blue)]()

---

## 📋 Project Overview

**Angidi** is an educational project designed to teach advanced system design concepts by implementing a full-featured, Amazon-like e-commerce platform. The project follows a phased approach, introducing new technologies and concepts incrementally while maintaining enterprise-grade quality standards.

### 🎯 Goals

- **Learn by Doing**: Implement 40+ system design concepts hands-on
- **Production-Ready**: Build a scalable platform handling millions of users
- **Best Practices**: Follow TDD, BDD, clean architecture, and security standards
- **Comprehensive Documentation**: Every phase thoroughly documented
- **Deployable**: Kubernetes-ready application with full observability

### ✨ Key Features

- **Microservices Architecture**: 11+ independent services
- **Multi-Database**: PostgreSQL, MongoDB, Redis, Elasticsearch
- **Event-Driven**: Kafka-based asynchronous communication
- **Cloud-Native**: Containerized with Kubernetes orchestration
- **Observable**: Prometheus, Grafana, Loki, Jaeger integration
- **Secure**: PCI-DSS compliant, enterprise security standards
- **High Performance**: <200ms P95 latency, 100K req/s throughput
- **Highly Available**: 99.99% uptime target, multi-region deployment

---

## 🏗️ Architecture

### High-Level Components

```
┌─────────────────────────────────────────────────────────┐
│                     Web & Mobile                        │
│                  (Next.js, React)                       │
└─────────────────────┬───────────────────────────────────┘
                      │
                      ▼
┌─────────────────────────────────────────────────────────┐
│                    CDN / WAF                            │
└─────────────────────┬───────────────────────────────────┘
                      │
                      ▼
┌─────────────────────────────────────────────────────────┐
│                  API Gateway                            │
│          (Auth, Rate Limiting, Routing)                 │
└──────┬──────────┬──────────┬──────────┬─────────────────┘
       │          │          │          │
       ▼          ▼          ▼          ▼
┌──────────┐ ┌────────┐ ┌─────────┐ ┌──────────┐
│  User    │ │Product │ │  Cart   │ │  Order   │ ...
│ Service  │ │Service │ │ Service │ │ Service  │
└──────────┘ └────────┘ └─────────┘ └──────────┘
       │          │          │          │
       └──────────┴──────────┴──────────┘
                      │
                      ▼
           ┌──────────────────┐
           │   Event Bus      │
           │    (Kafka)       │
           └──────────────────┘
```

**Services**: User, Product, Search, Cart, Order, Inventory, Payment, Notification, Recommendation, Review, Admin

**Data Stores**: PostgreSQL (transactions), MongoDB (catalog), Redis (cache), Elasticsearch (search)

For detailed architecture, see [Architecture Documentation](./docs/architecture/README.md).

---

## 📚 Documentation

### Core Specifications

- **[Main Specification](./docs/specs/SPEC.md)**: Complete project overview
- **[Functional Requirements](./docs/specs/FUNCTIONAL_REQUIREMENTS.md)**: 130+ feature requirements
- **[Non-Functional Requirements](./docs/specs/NON_FUNCTIONAL_REQUIREMENTS.md)**: Performance, security, scalability
- **[System Design Concepts](./docs/specs/SYSTEM_DESIGN_CONCEPTS.md)**: 40+ concepts to learn
- **[Technology Stack](./docs/specs/TECH_STACK.md)**: Tools and frameworks by phase

### Implementation Guide

- **[Architecture Guide](./docs/architecture/README.md)**: System design and components
- **[Phases Overview](./docs/phases/README.md)**: 16-phase roadmap

---

## 🚀 Learning Phases

The project is divided into 16 phases, each introducing 1-2 new tools/concepts:

| Phase | Name | Focus | Key Technologies |
|-------|------|-------|------------------|
| **0** | ✅ Planning & Specification | Requirements & Design | Markdown, Diagrams |
| **1** | ✅ Repository Scaffolding | Project Setup | Go, React, GitHub Actions |
| **2** | ✅ Core Services | User & Product Services | Chi/Gin, JWT |
| **3** | 📝 Database Integration | Data Persistence | PostgreSQL |
| **4** | ⏳ Containerization | Docker Packaging | Docker, Compose |
| **5** | ⏳ Shopping Cart | Session Management | Redis |
| **6** | ⏳ Order Processing | Transactions | ACID, Idempotency |
| **7** | ⏳ Caching Layer | Performance | Redis Cluster |
| **8** | ⏳ Search & Discovery | Product Search | Elasticsearch, MongoDB |
| **9** | ⏳ Event-Driven | Async Communication | Kafka |
| **10** | ⏳ Payment Integration | Secure Payments | Stripe, PCI-DSS |
| **11** | ⏳ Kubernetes Deployment | Orchestration | K8s, Helm |
| **12** | ⏳ Monitoring | Observability | Prometheus, Grafana, Loki |
| **13** | ⏳ Distributed Tracing | Request Tracking | Jaeger, OpenTelemetry |
| **14** | ⏳ Advanced Security | Enterprise Security | Vault, TLS |
| **15** | ⏳ Performance Tuning | Optimization | CDN, Sharding |
| **16** | ⏳ Chaos Engineering | Resilience Testing | Chaos Mesh |

**Current Status**: Phase 2 Complete ✅  
**Next Phase**: Phase 3 - Database Integration 📝 (Planning)

See [detailed phase documentation](./docs/phases/README.md) for more information.

---

## 🛠️ Technology Stack

### Backend
- **Language**: Go 1.21+
- **Framework**: Chi / Gin
- **Databases**: PostgreSQL, MongoDB, Redis, Elasticsearch
- **Messaging**: Apache Kafka
- **Logging**: Zap (structured logging)

### Frontend
- **Framework**: Next.js 14+
- **Language**: TypeScript 5+
- **UI Library**: React 18+
- **Styling**: Tailwind CSS 3+
- **State**: React Query, Zustand

### Infrastructure
- **Containers**: Docker 24+
- **Orchestration**: Kubernetes 1.28+
- **Package Manager**: Helm 3+
- **CI/CD**: GitHub Actions

### Observability
- **Metrics**: Prometheus
- **Visualization**: Grafana
- **Logging**: Loki
- **Tracing**: Jaeger
- **Instrumentation**: OpenTelemetry

See [complete tech stack](./docs/specs/TECH_STACK.md) for details.

---

## 🎓 System Design Concepts

This project covers 40+ system design concepts across 10 categories:

### Scalability
- Horizontal/Vertical Scaling
- Database Sharding
- Caching Strategies
- Load Balancing

### Reliability & High Availability
- Multi-Region Architecture
- Circuit Breaker Pattern
- Graceful Degradation
- Rate Limiting

### Performance Optimization
- CDN Integration
- Query Optimization
- Asynchronous Processing
- API Optimization

### Data Management
- CAP Theorem Trade-offs
- Event-Driven Architecture
- SAGA Pattern
- Data Replication

### Microservices
- Service Decomposition
- API Gateway Pattern
- Service Discovery
- Inter-Service Communication

### Distributed Systems
- Distributed Locking
- Idempotency
- Eventual Consistency
- Distributed Tracing

### Security
- Authentication & Authorization
- PCI-DSS Compliance
- Data Encryption
- DDoS Protection

### Observability
- Metrics Collection
- Centralized Logging
- Alerting & Monitoring
- Dashboard Creation

### Testing
- Unit Testing (TDD)
- Integration Testing
- End-to-End Testing
- Load & Performance Testing
- Chaos Engineering

### DevOps & Infrastructure
- CI/CD Pipelines
- Infrastructure as Code
- Container Orchestration
- Service Mesh

See [System Design Concepts](./docs/specs/SYSTEM_DESIGN_CONCEPTS.md) for detailed descriptions.

---

## 🏁 Quick Start

### Prerequisites
- Go 1.21+
- Node.js 20 LTS
- Docker 24+ (for later phases)
- Kubernetes 1.28+ (for later phases)

### Phase 1: Repository Setup ✅

```bash
# Clone the repository
git clone https://github.com/yesoreyeram/angidi-demo-app.git
cd angidi-demo-app

# Backend setup
cd backend
make install-tools
make deps
make build
make test
make run  # Starts on http://localhost:8080

# In a new terminal - Frontend setup
cd ../frontend
npm install
npm run build
npm run dev  # Starts on http://localhost:3000
```

For detailed setup instructions, see:
- [Backend README](./backend/README.md)
- [Frontend README](./frontend/README.md)

---

## 📊 Non-Functional Requirements

### Performance Targets
- **Latency**: <200ms P95 for critical operations
- **Throughput**: 100,000 requests/second peak
- **Availability**: 99.99% uptime (~5 min/year downtime)

### Scale Targets
- **Users**: 100M registered, 50M daily active
- **Products**: 10M+ in catalog
- **Orders**: 5M per day
- **Storage**: 50TB/year growth

### Quality Standards
- **Code Coverage**: >80% for critical paths
- **Security**: OWASP Top 10 compliant, PCI-DSS for payments
- **Documentation**: Comprehensive for all components
- **Testing**: TDD/BDD throughout

---

## 🤝 Development Principles

1. **Test-Driven Development (TDD)**: Write tests before code
2. **Behavior-Driven Development (BDD)**: User story-driven features
3. **Clean Architecture**: Separation of concerns, SOLID principles
4. **Security First**: Security considerations in every phase
5. **Incremental Development**: Small, tested, documented changes
6. **Observability**: Comprehensive logging, metrics, tracing
7. **Documentation**: Document as you build

---

## 📖 Learning Resources

This project is based on system design principles from:
- **Transcript**: Deep dive on e-commerce platform design (included in issue)
- **Best Practices**: Industry standards and patterns
- **Real-World Systems**: Amazon, eBay, Alibaba architectures

---

## 🗺️ Roadmap

### Phase 0: ✅ Planning (Complete)
- Requirements specification
- Architecture design
- Technology selection
- Phase planning

### Phase 1: ✅ Scaffolding (Complete)
- Go backend structure
- Next.js frontend setup
- CI/CD pipeline
- Coding standards

### Phase 2: ✅ Core Services (Complete)
- User & Product Services
- JWT Authentication
- API Gateway
- RESTful APIs

### Phase 3: 📝 Database Integration (Planning)
- PostgreSQL integration
- Database schema design
- Migration system
- Repository pattern implementation

### Future Phases (4-16)
See [Phases Documentation](./docs/phases/README.md) for complete roadmap.

---

## 📝 License

MIT License - see [LICENSE](./LICENSE) file for details.

---

## 🙏 Acknowledgments

This project is an educational initiative to learn and practice advanced system design concepts. It draws inspiration from real-world e-commerce platforms and system design best practices.

---

**Current Phase**: 3 - Database Integration 📝 (Planning)  
**Progress**: 19% (3/17 phases complete, Phase 3 planning in progress)  
**Last Updated**: 2025-10-27
