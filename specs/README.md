# Angidi Specifications

This directory contains the complete technical specifications for the Angidi e-commerce platform.

## Directory Structure

```
specs/
├── README.md                   # This file
├── system-design/             # System design concepts
│   ├── concepts.md           # All concepts with descriptions
│   └── checklist.md          # Progress tracking
├── requirements/              # Requirements documentation
│   ├── functional.md         # Functional requirements
│   └── non-functional.md     # Non-functional requirements
├── architecture/              # Architecture documentation
│   ├── overview.md           # High-level design
│   └── decisions/            # Architecture Decision Records
├── phases/                    # Phase-based plans
│   ├── README.md             # Phases overview
│   └── phase-X-name/         # Individual phase docs
└── scripts/                   # Automation scripts
    ├── validate.sh           # Spec validation
    ├── generate-checklist.sh # Progress tracker
    └── new-phase.sh          # Phase template creator
```

## Purpose

The specifications serve multiple purposes:

1. **Learning Resource**: Comprehensive documentation of system design concepts
2. **Reference Guide**: Requirements and architecture for implementation
3. **Progress Tracker**: Monitor learning and implementation progress
4. **Decision Log**: Record of architectural decisions and trade-offs

## How to Use

### For Learning

1. Start with [System Design Concepts](./system-design/concepts.md) to understand what concepts will be covered
2. Review [Requirements](./requirements/) to understand what will be built
3. Study [Architecture Overview](./architecture/overview.md) for the system design
4. Follow [Phase Plan](./phases/README.md) for the learning roadmap

### For Implementation

1. Check current phase in `phases/`
2. Review requirements in `requirements/`
3. Reference architecture decisions in `architecture/`
4. Use scripts for validation and tracking

### For Contributing

1. Follow existing documentation patterns
2. Update ADRs for significant decisions
3. Keep requirements up-to-date
4. Run validation before committing

## Relationship with Other Directories

- **`.spec/`**: Legacy specification directory (being phased out)
- **`.specify/`**: GitHub Spec Kit structure (project memory, templates)
- **`.github/prompts/`**: AI agent prompts for specialized assistance
- **`specs/`**: Main specifications directory (current, preferred)

## Automation

### Validate Specifications

```bash
# Validate GitHub Spec Kit structure
./.specify/scripts/validate-specs.sh

# Validate legacy spec structure
./.spec/scripts/validate.sh
```

### Generate Progress Checklist

```bash
./specs/scripts/generate-checklist.sh
```

### Create New Phase

```bash
./specs/scripts/new-phase.sh <number> <name>
```

## Principles

- **Test-Driven Development (TDD)** - Write tests before implementation
- **Behavior-Driven Development (BDD)** - Define behavior through scenarios
- **Enterprise Security** - Follow security best practices
- **Code Quality** - Maintain high code quality standards
- **Documentation** - Comprehensive documentation for all components
- **Incremental Learning** - One or two tools per phase

## Navigation

- **Next**: Start with [System Design Concepts](./system-design/concepts.md)
- **Requirements**: See [Functional](./requirements/functional.md) and [Non-Functional](./requirements/non-functional.md)
- **Architecture**: Read [Overview](./architecture/overview.md)
- **Phases**: Follow [Phase Plan](./phases/README.md)

---

**Current Phase**: Phase 0 - Planning & Repository Setup
**Last Updated**: 2025-10-26
