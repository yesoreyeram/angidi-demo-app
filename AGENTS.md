# Angidi Project Agents

This document describes the specialized AI agents available for the Angidi e-commerce platform project.

## Overview

The project uses specialized AI agents (via GitHub Copilot) to assist with different aspects of system design, implementation, and documentation. Each agent has specific expertise and can be invoked through the `.github/prompts/` directory.

## Available Agents

### 1. System Design Expert (`speckit.system-design.prompt.md`)

**Expertise**: Distributed systems architecture, microservices design, database scaling

**Use When**:
- Designing system architecture
- Making technology choices
- Evaluating scalability approaches
- Understanding trade-offs between options
- Designing for high availability

**Example Questions**:
- "How should we handle inventory management to prevent overselling?"
- "What's the best approach for implementing distributed transactions?"
- "Should we use PostgreSQL or MongoDB for the product catalog?"
- "How do we achieve 99.99% uptime?"

**Key Focus Areas**:
- Microservices patterns
- Event-driven architecture
- CAP theorem and consistency models
- Caching strategies
- Load balancing and sharding

### 2. Requirements Analyst (`speckit.requirements.prompt.md`)

**Expertise**: Functional and non-functional requirements documentation

**Use When**:
- Defining new features
- Writing user stories
- Creating acceptance criteria
- Documenting edge cases
- Establishing metrics and SLOs

**Example Questions**:
- "What are the requirements for user authentication?"
- "What edge cases should we handle for order cancellation?"
- "What are the performance requirements for search?"
- "How do we measure order processing success?"

**Key Focus Areas**:
- Functional requirements with clear acceptance criteria
- Non-functional requirements (performance, scalability, security)
- User stories and use cases
- Business metrics and KPIs
- Testable requirements

### 3. Testing Strategist (`speckit.testing.prompt.md`)

**Expertise**: Test-driven development, behavior-driven development, test automation

**Use When**:
- Designing test strategies
- Writing BDD scenarios
- Planning integration tests
- Setting up performance tests
- Implementing security tests

**Example Questions**:
- "What BDD scenarios do we need for checkout?"
- "How should we test the payment service integration?"
- "What performance tests should we run for the search service?"
- "How do we test distributed transactions?"

**Key Focus Areas**:
- Unit, integration, and E2E testing
- BDD scenario writing (Gherkin)
- Test automation strategies
- Performance and load testing
- Security testing

## How to Use Agents

### In GitHub Copilot Chat

You can reference these agents in your Copilot conversations:

```
@workspace /ask Using the system design expert agent,
how should we implement the recommendation service?
```

### In Code Comments

Reference agents in your code for context-aware suggestions:

```go
// @agent system-design: Need to implement caching strategy for product catalog
// Requirements: <200ms p95 latency, handle 10M products
```

### In PRs and Issues

Mention agents in PR descriptions or issues:

```markdown
## Design Question

@system-design-agent: Should we use a separate service for recommendations
or add it to the product service?

Consider:
- 100M users
- Real-time personalization
- ML model updates
```

## Agent Interaction Patterns

### Sequential Consultation

1. **Requirements Analyst**: Define what needs to be built
2. **System Design Expert**: Design how to build it
3. **Testing Strategist**: Define how to verify it

### Collaborative Problem Solving

For complex problems, consult multiple agents:

```markdown
## Problem: Implement real-time inventory updates

@requirements-agent: What are the requirements?
- Must prevent overselling
- Should handle 10K updates/sec
- Need strong consistency

@system-design-agent: Proposed solution?
- Redis for distributed locks
- PostgreSQL for persistence
- Kafka for async updates

@testing-agent: How to test?
- Unit tests for lock acquisition
- Integration tests for race conditions
- Load tests at 10K updates/sec
```

## Agent Best Practices

### Do's
✅ Provide clear context and constraints
✅ Ask specific, focused questions
✅ Include scale and performance requirements
✅ Reference existing specs and documentation
✅ Request code examples when helpful

### Don'ts
❌ Ask overly broad questions
❌ Ignore project constraints and phase plan
❌ Request implementation without requirements
❌ Skip documenting agent recommendations
❌ Contradict established architectural decisions

## Agent Specialization Map

```
Feature Development Flow:
┌─────────────────┐
│  Requirements   │──> Define functional & non-functional requirements
│     Analyst     │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│ System Design   │──> Design architecture & choose technologies
│     Expert      │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│    Testing      │──> Create test strategy & BDD scenarios
│   Strategist    │
└─────────────────┘
```

## Integration with Specs

Agent recommendations should be documented in:

- **Design Decisions**: `specs/architecture/decisions/`
- **Requirements**: `specs/requirements/`
- **Test Plans**: `specs/phases/phase-X/testing-strategy.md`
- **ADRs**: Architecture Decision Records

## Memory and Context

Agents have access to:

- `.specify/memory/` - Project context and decisions
- `specs/` - All specifications and requirements
- `.spec/` - Legacy spec kit (transitioning)
- Phase plans and current phase context

## Continuous Improvement

As the project evolves:

1. **Update agent prompts** with new patterns learned
2. **Add examples** from actual implementations
3. **Refine specializations** based on usage
4. **Document anti-patterns** to avoid

## Agent Versioning

Agents are versioned with the spec kit:

- **v1.0**: Initial agent prompts (Phase 0)
- **v1.1**: Enhanced with Phase 1 learnings
- **v2.0**: Major updates after microservices decomposition (Phase 2)

## Getting Help

If agents provide conflicting advice:

1. Check the current phase plan for context
2. Consult existing ADRs for precedent
3. Escalate to team discussion if needed
4. Document the decision for future reference

## Contributing to Agents

To improve agent prompts:

1. Identify gaps or unclear guidance
2. Propose updates via PR
3. Test with real scenarios
4. Update this document with new capabilities

---

**Note**: These agents are assistive tools. Final decisions should always consider:
- Project constraints and phase plan
- Team expertise and preferences
- Existing architectural choices
- Cost and complexity trade-offs
