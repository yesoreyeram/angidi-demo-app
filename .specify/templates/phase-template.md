# Phase Template

Use this template when creating a new phase.

## Phase Structure

Each phase should have:

```
specs/phases/phase-X-name/
├── README.md                  # Phase overview
├── concepts.md               # System design concepts covered
├── design-decisions.md       # ADRs for this phase
├── setup.md                  # Environment setup
├── testing-strategy.md       # Test plans
└── checklist.md             # Completion criteria
```

## README.md Template

```markdown
# Phase X: [Name]

**Duration**: X weeks
**Status**: [Not Started | In Progress | Complete]

## Overview
Brief description of what this phase accomplishes.

## Goals
### Primary Goals
1. Goal 1
2. Goal 2

### Learning Objectives
- Objective 1
- Objective 2

## System Design Concepts
List concepts covered in this phase.

## Tools Introduced
- Tool 1: Purpose
- Tool 2: Purpose

## Prerequisites
- [ ] Previous phase complete
- [ ] Prerequisite 1

## Implementation Plan
Detailed task breakdown.

## Success Criteria
- [ ] All tests passing
- [ ] Documentation complete
```

## Creating a New Phase

Use the script:

```bash
./specs/scripts/new-phase.sh <number> <name>
```

Or manually:

1. Create directory: `specs/phases/phase-X-name/`
2. Copy template files
3. Fill in phase-specific content
4. Update `specs/phases/README.md`
5. Add to project timeline

## Phase Naming Convention

- Use lowercase with hyphens
- Format: `phase-X-descriptive-name`
- Examples:
  - `phase-1-basic-monolith`
  - `phase-2-microservice-decomposition`
  - `phase-3-caching-layer`

## Phase Documentation Standards

### Completeness
- All sections filled
- No TODO markers
- Examples provided

### Clarity
- Clear objectives
- Step-by-step instructions
- Explanations of decisions

### Consistency
- Follow project terminology
- Consistent formatting
- Cross-references accurate
