#!/bin/bash

# new-phase.sh - Creates a new phase directory with template structure

set -e

# Check arguments
if [ $# -ne 2 ]; then
    echo "Usage: $0 <phase-number> <phase-name>"
    echo ""
    echo "Examples:"
    echo "  $0 1 basic-monolith"
    echo "  $0 2 first-microservice"
    echo ""
    exit 1
fi

PHASE_NUM=$1
PHASE_NAME=$2
PHASE_DIR=".spec/phases/phase-${PHASE_NUM}-${PHASE_NAME}"

# Check if phase already exists
if [ -d "$PHASE_DIR" ]; then
    echo "❌ Error: Phase directory already exists at $PHASE_DIR"
    exit 1
fi

echo "📁 Creating new phase: Phase $PHASE_NUM - $PHASE_NAME"

# Create phase directory
mkdir -p "$PHASE_DIR"

# Create README.md
cat > "$PHASE_DIR/README.md" << EOF
# Phase $PHASE_NUM: $(echo $PHASE_NAME | sed 's/-/ /g' | sed 's/\b\(.\)/\u\1/g')

**Duration**: TBD  
**Status**: Not Started

## Overview

Brief overview of what this phase accomplishes and why it's important.

## Goals

### Primary Goals
1. Goal 1
2. Goal 2
3. Goal 3

### Learning Objectives
- Learning objective 1
- Learning objective 2
- Learning objective 3

## System Design Concepts

List the system design concepts covered in this phase:

### Concept 1: Name
Description of the concept and how it's applied

### Concept 2: Name
Description of the concept and how it's applied

## Tools Introduced

### Tool 1: Name
- **Purpose**: Why we're using this tool
- **Documentation**: Link to official docs
- **Version**: Specify version

### Tool 2: Name
- **Purpose**: Why we're using this tool
- **Documentation**: Link to official docs
- **Version**: Specify version

## Design Decisions

### 1. Decision Name

**Decision**: What was decided

**Rationale**:
- Reason 1
- Reason 2

**Alternatives Considered**:
- Alternative 1: Why not chosen
- Alternative 2: Why not chosen

**Consequences**:
- Positive consequence 1
- Potential drawback 1

## Prerequisites

Before starting this phase:
- [ ] Prerequisite 1
- [ ] Prerequisite 2
- [ ] Phase $(($PHASE_NUM - 1)) completed

## Implementation Plan

### Task 1: Name
**Duration**: X days

- [ ] Sub-task 1
- [ ] Sub-task 2
- [ ] Sub-task 3

### Task 2: Name
**Duration**: X days

- [ ] Sub-task 1
- [ ] Sub-task 2

## Testing Strategy

### Unit Tests
- Test category 1
- Test category 2

### Integration Tests
- Integration test scenario 1
- Integration test scenario 2

### E2E Tests
- E2E test scenario 1
- E2E test scenario 2

### BDD Scenarios

\`\`\`gherkin
Feature: Feature Name

  Scenario: Scenario name
    Given precondition
    When action
    Then expected result
\`\`\`

## Documentation Requirements

Documents to create or update:
- [ ] API documentation
- [ ] Setup guide
- [ ] Configuration guide
- [ ] Troubleshooting guide
- [ ] Update main README

## Success Criteria

This phase is complete when:
- [ ] All implementation tasks finished
- [ ] All tests passing (unit, integration, E2E)
- [ ] Code coverage > 80%
- [ ] Documentation complete
- [ ] Code review completed
- [ ] Security review passed (if applicable)
- [ ] Performance benchmarks met (if applicable)

## Deliverables

1. Deliverable 1
2. Deliverable 2
3. Deliverable 3

## Next Phase

**Phase $(($PHASE_NUM + 1))**: Next phase name and brief description

## Resources

- [Resource 1](https://example.com)
- [Resource 2](https://example.com)

## Notes

Additional notes, gotchas, or important information.
EOF

echo "✅ Created $PHASE_DIR/README.md"

# Create concepts.md
cat > "$PHASE_DIR/concepts.md" << EOF
# System Design Concepts - Phase $PHASE_NUM

This document provides detailed explanations of the system design concepts introduced in this phase.

## Concept 1: Name

### Definition
Clear definition of the concept.

### Why It Matters
Explain the importance and real-world relevance.

### How We Implement It
Specific implementation details for this project.

### Common Patterns
- Pattern 1
- Pattern 2

### Pitfalls to Avoid
- Pitfall 1
- Pitfall 2

### Further Reading
- [Resource 1](https://example.com)
- [Resource 2](https://example.com)

---

## Concept 2: Name

### Definition
Clear definition of the concept.

### Why It Matters
Explain the importance and real-world relevance.

### How We Implement It
Specific implementation details for this project.

### Common Patterns
- Pattern 1
- Pattern 2

### Pitfalls to Avoid
- Pitfall 1
- Pitfall 2

### Further Reading
- [Resource 1](https://example.com)
- [Resource 2](https://example.com)
EOF

echo "✅ Created $PHASE_DIR/concepts.md"

# Create design-decisions.md
cat > "$PHASE_DIR/design-decisions.md" << EOF
# Design Decisions - Phase $PHASE_NUM

Architecture Decision Records (ADRs) for Phase $PHASE_NUM.

## Format

Each decision follows this format:
- **Decision**: What was decided
- **Date**: When it was decided
- **Status**: Proposed, Accepted, Deprecated, Superseded
- **Context**: The situation requiring a decision
- **Options**: Alternatives considered
- **Decision**: The chosen option
- **Rationale**: Why this option was chosen
- **Consequences**: Expected outcomes and trade-offs

---

## ADR-${PHASE_NUM}.1: Decision Title

**Date**: YYYY-MM-DD  
**Status**: Proposed

### Context
Describe the context and problem statement.

### Options Considered

#### Option 1: Name
- Pros: Benefit 1, Benefit 2
- Cons: Drawback 1, Drawback 2

#### Option 2: Name
- Pros: Benefit 1, Benefit 2
- Cons: Drawback 1, Drawback 2

### Decision
We chose Option X because...

### Rationale
1. Reason 1
2. Reason 2
3. Reason 3

### Consequences

**Positive**:
- Consequence 1
- Consequence 2

**Negative**:
- Trade-off 1
- Trade-off 2

**Mitigations**:
- How we'll address the negative consequences

---

## ADR-${PHASE_NUM}.2: Another Decision

...
EOF

echo "✅ Created $PHASE_DIR/design-decisions.md"

# Create setup.md
cat > "$PHASE_DIR/setup.md" << EOF
# Setup Guide - Phase $PHASE_NUM

This guide walks through setting up your development environment for Phase $PHASE_NUM.

## Prerequisites

### System Requirements
- Operating System: Linux (Ubuntu 20.04+ recommended)
- RAM: Minimum 8GB
- Disk Space: Minimum 20GB free

### Required Software
- [ ] Git 2.x
- [ ] Tool 1 (version X.Y)
- [ ] Tool 2 (version X.Y)

### Required Knowledge
- Familiarity with Tool 1
- Understanding of Concept X

## Installation Steps

### Step 1: Install Tool 1

\`\`\`bash
# Installation commands
curl -fsSL https://example.com/install.sh | sh
\`\`\`

**Verify installation**:
\`\`\`bash
tool1 --version
\`\`\`

Expected output: \`v1.2.3\`

### Step 2: Install Tool 2

\`\`\`bash
# Installation commands
\`\`\`

**Verify installation**:
\`\`\`bash
tool2 --version
\`\`\`

## Configuration

### Configure Tool 1

Create configuration file:
\`\`\`bash
cat > config/tool1.conf << 'CONF'
# Configuration here
CONF
\`\`\`

### Environment Variables

Add to \`.env\`:
\`\`\`bash
VAR1=value1
VAR2=value2
\`\`\`

## Running the Project

### Start Development Environment

\`\`\`bash
# Commands to start the project
make dev
\`\`\`

### Verify Everything Works

\`\`\`bash
# Test commands
make test
\`\`\`

## Troubleshooting

### Issue 1: Description

**Symptoms**: What you see

**Solution**:
\`\`\`bash
# Fix commands
\`\`\`

### Issue 2: Description

**Symptoms**: What you see

**Solution**:
\`\`\`bash
# Fix commands
\`\`\`

## Next Steps

After setup is complete:
1. Review the implementation plan in README.md
2. Understand the concepts in concepts.md
3. Start implementing the first task
EOF

echo "✅ Created $PHASE_DIR/setup.md"

# Create testing-strategy.md
cat > "$PHASE_DIR/testing-strategy.md" << EOF
# Testing Strategy - Phase $PHASE_NUM

This document outlines the testing approach for Phase $PHASE_NUM.

## Test Coverage Goals

- Unit Test Coverage: >80%
- Integration Test Coverage: >70%
- E2E Test Coverage: Critical paths

## Unit Tests

### Component 1 Tests

**File**: \`path/to/component1_test.go\`

\`\`\`go
// Example test structure
func TestComponentName(t *testing.T) {
    // Arrange
    
    // Act
    
    // Assert
}
\`\`\`

**Test Cases**:
- [ ] Test case 1: Description
- [ ] Test case 2: Description
- [ ] Test edge case: Description
- [ ] Test error handling: Description

### Component 2 Tests

...

## Integration Tests

### Integration Scenario 1: Name

**Purpose**: Test interaction between Component A and Component B

**Setup**:
\`\`\`bash
# Setup commands
\`\`\`

**Test Cases**:
- [ ] Happy path
- [ ] Error handling
- [ ] Timeout handling

## End-to-End Tests

### E2E Scenario 1: User Journey

**Description**: User completes action X

**Steps**:
1. User does action 1
2. System responds with result 1
3. User does action 2
4. System completes flow

**Verification**:
- [ ] Checkpoint 1
- [ ] Checkpoint 2
- [ ] Final state is correct

## BDD Scenarios

### Feature: Feature Name

\`\`\`gherkin
Feature: Feature Name
  As a user type
  I want to capability
  So that benefit

  Scenario: Scenario 1
    Given initial state
    When user action
    Then expected result
    
  Scenario: Scenario 2
    Given different initial state
    When user action
    Then different expected result
\`\`\`

## Performance Tests

### Load Test 1: Endpoint Name

**Objective**: Test that endpoint handles X requests/second

**Configuration**:
- Duration: 5 minutes
- VUs (Virtual Users): 100
- Target RPS: 1000

**Acceptance Criteria**:
- p95 latency < 200ms
- Error rate < 1%

## Test Execution

### Running Tests Locally

\`\`\`bash
# Unit tests
make test-unit

# Integration tests
make test-integration

# E2E tests
make test-e2e

# All tests
make test-all
\`\`\`

### CI/CD Integration

Tests run automatically on:
- Every commit to feature branch
- Pull request creation
- Merge to main branch

## Test Data Management

### Test Database Setup

\`\`\`bash
# Create test database
make test-db-create

# Seed test data
make test-db-seed

# Clean test database
make test-db-clean
\`\`\`

### Test Data Fixtures

Location: \`test/fixtures/\`

- \`users.json\`: Test user data
- \`products.json\`: Test product data

## Mocking Strategy

### External Services

Mock the following external services:
- Service 1: Use mock implementation
- Service 2: Use test doubles

### Database

- Use test database for integration tests
- Use in-memory database for unit tests where possible

## Test Documentation

Each test should:
- Have a clear, descriptive name
- Include comments explaining the scenario
- Follow AAA pattern (Arrange, Act, Assert)
- Clean up resources after execution

## Continuous Improvement

After each test run:
- Review test failures
- Update tests for new edge cases
- Refactor slow tests
- Remove obsolete tests
EOF

echo "✅ Created $PHASE_DIR/testing-strategy.md"

# Create checklist.md
cat > "$PHASE_DIR/checklist.md" << EOF
# Phase $PHASE_NUM Completion Checklist

Track your progress through Phase $PHASE_NUM with this checklist.

## Planning & Setup

- [ ] Read phase README thoroughly
- [ ] Understand all system design concepts
- [ ] Review design decisions and rationale
- [ ] Complete environment setup
- [ ] Verify all prerequisites met

## Implementation

### Task 1: Name
- [ ] Sub-task 1
- [ ] Sub-task 2
- [ ] Sub-task 3
- [ ] Code review completed
- [ ] Tests passing

### Task 2: Name
- [ ] Sub-task 1
- [ ] Sub-task 2
- [ ] Code review completed
- [ ] Tests passing

## Testing

- [ ] Unit tests written and passing
- [ ] Unit test coverage >80%
- [ ] Integration tests written and passing
- [ ] E2E tests written and passing
- [ ] BDD scenarios implemented
- [ ] Performance tests passing (if applicable)
- [ ] All tests pass in CI/CD

## Code Quality

- [ ] Code follows style guide
- [ ] No linter warnings
- [ ] Static analysis passing
- [ ] Security scan passing
- [ ] Dependency vulnerability check passing
- [ ] Code review completed
- [ ] Refactoring completed

## Documentation

- [ ] API documentation updated
- [ ] Setup guide complete
- [ ] Configuration documented
- [ ] Troubleshooting guide updated
- [ ] README updated
- [ ] Comments added where necessary
- [ ] Architecture diagrams updated (if needed)

## Security

- [ ] Input validation implemented
- [ ] Authentication/authorization working
- [ ] Secrets not committed to repo
- [ ] Security review completed
- [ ] Vulnerabilities addressed

## Performance

- [ ] Performance benchmarks met
- [ ] Load tests passing
- [ ] No memory leaks
- [ ] Database queries optimized
- [ ] Caching implemented (if needed)

## Deployment

- [ ] Application builds successfully
- [ ] Docker containers working
- [ ] Can run locally
- [ ] Can run in development environment
- [ ] Configuration externalized
- [ ] Health checks implemented

## Final Review

- [ ] All tasks completed
- [ ] All tests passing
- [ ] Documentation complete
- [ ] Code reviewed and approved
- [ ] No known bugs
- [ ] Performance acceptable
- [ ] Security review passed
- [ ] Ready for next phase

## Phase Completion

- [ ] Update system design concepts checklist
- [ ] Update main project README
- [ ] Tag release (if applicable)
- [ ] Celebrate! 🎉

**Completion Date**: ___________

**Notes**:
EOF

echo "✅ Created $PHASE_DIR/checklist.md"

# Update phases README to include new phase
echo ""
echo "📝 Don't forget to:"
echo "  1. Update .spec/phases/README.md to include Phase $PHASE_NUM"
echo "  2. Fill in the template content in the phase directory"
echo "  3. Add specific implementation details"
echo ""
echo "✨ Phase $PHASE_NUM template created successfully at:"
echo "   $PHASE_DIR"
