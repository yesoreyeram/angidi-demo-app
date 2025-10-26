# Angidi Specification Kit

This directory contains the complete specification, design documentation, and planning materials for the Angidi e-commerce platform project.

## Purpose

The Angidi project is a comprehensive learning platform designed to teach advanced system design concepts through the implementation of a highly scalable, performant, and reliable e-commerce platform similar to Amazon.

## Structure

- **`system-design/`** - System design concepts, patterns, and principles
- **`requirements/`** - Functional and non-functional requirements
- **`architecture/`** - Architecture diagrams and design decisions
- **`phases/`** - Phase-based learning modules with detailed plans
- **`scripts/`** - Shell scripts for spec kit management

## Navigation

1. Start with [System Design Concepts](./system-design/concepts.md) to understand what we'll learn
2. Review [Requirements](./requirements/functional.md) to understand what we're building
3. Check [Architecture Overview](./architecture/overview.md) for the high-level design
4. Follow the [Phase Plan](./phases/README.md) for the implementation roadmap

## How to Use

Each phase contains:
- **Overview** - Goals and objectives
- **Concepts** - System design concepts covered
- **Tools** - Technologies introduced
- **Design Decisions** - Rationale behind choices
- **Setup Instructions** - How to set up the environment
- **Testing Strategy** - Test plans and BDD scenarios
- **Documentation** - Detailed documentation for the phase

## Scripts

Use the provided scripts to manage the spec kit:

```bash
# Validate spec kit structure
./.spec/scripts/validate.sh

# Generate phase checklist
./.spec/scripts/generate-checklist.sh

# Create new phase template
./.spec/scripts/new-phase.sh <phase-number> <phase-name>
```

## Principles

- **Test-Driven Development (TDD)** - Write tests before implementation
- **Behavior-Driven Development (BDD)** - Define behavior through scenarios
- **Enterprise Security** - Follow security best practices
- **Code Quality** - Maintain high code quality standards
- **Documentation** - Comprehensive documentation for all components
- **Incremental Learning** - One or two tools per phase
