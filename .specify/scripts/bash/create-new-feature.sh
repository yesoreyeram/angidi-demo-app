#!/bin/bash

# create-new-feature.sh - Creates a new feature with specs and implementation structure

set -e

# Source common functions
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "${SCRIPT_DIR}/common.sh"

print_header "Create New Feature"

# Check if feature name is provided
if [ $# -lt 1 ]; then
    echo "Usage: $0 <feature-name> [service-name]"
    echo ""
    echo "Examples:"
    echo "  $0 user-authentication user-service"
    echo "  $0 product-search search-service"
    echo "  $0 checkout-process order-service"
    exit 1
fi

FEATURE_NAME=$1
SERVICE_NAME=${2:-""}

# Sanitize names
FEATURE_NAME_CLEAN=$(sanitize_name "$FEATURE_NAME")
SERVICE_NAME_CLEAN=$(sanitize_name "$SERVICE_NAME")

print_info "Creating feature: $FEATURE_NAME_CLEAN"
if [ -n "$SERVICE_NAME_CLEAN" ]; then
    print_info "For service: $SERVICE_NAME_CLEAN"
fi

# Confirm action
echo ""
if ! confirm "Create this feature?"; then
    print_info "Operation cancelled"
    exit 0
fi

# Get project root
PROJECT_ROOT=$(get_project_root)

# Create feature directory structure
FEATURE_DIR="$PROJECT_ROOT/specs/features/$FEATURE_NAME_CLEAN"
ensure_dir "$FEATURE_DIR"

# Create feature specification
SPEC_FILE="$FEATURE_DIR/spec.md"
if [ ! -f "$SPEC_FILE" ]; then
    cat > "$SPEC_FILE" << EOF
# Feature: $(echo "$FEATURE_NAME" | sed 's/-/ /g' | sed 's/\b\(.\)/\u\1/g')

**Status**: Draft
**Service**: ${SERVICE_NAME:-TBD}
**Created**: $(get_date)
**Owner**: TBD

## Overview

Brief description of the feature and its purpose.

## Business Requirements

### Goals
- Goal 1
- Goal 2
- Goal 3

### Success Metrics
- Metric 1: Description
- Metric 2: Description
- Metric 3: Description

## Functional Requirements

### FR-1: Requirement Title
**Description**: What the system should do

**Acceptance Criteria**:
- [ ] Criterion 1
- [ ] Criterion 2
- [ ] Criterion 3

### FR-2: Another Requirement
**Description**: What the system should do

**Acceptance Criteria**:
- [ ] Criterion 1
- [ ] Criterion 2

## Non-Functional Requirements

### Performance
- Response time: < XXXms (p95)
- Throughput: XXX requests/second
- Resource usage: CPU < XX%, Memory < XXXMiB

### Security
- Authentication required: Yes/No
- Authorization: Required roles/permissions
- Data encryption: At rest / In transit
- PII handling: Describe requirements

### Scalability
- Expected load: XXX users, XXX requests/sec
- Scaling approach: Horizontal / Vertical
- Caching strategy: Describe

## Technical Design

### Architecture

\`\`\`
[Add architecture diagram or description]
\`\`\`

### API Design

#### Endpoints

**POST /api/v1/resource**
- Description: Create a new resource
- Request body: \`{ "field": "value" }\`
- Response: \`{ "id": "xxx", "status": "created" }\`
- Status codes: 201, 400, 401, 500

**GET /api/v1/resource/:id**
- Description: Get resource by ID
- Path params: \`id\` (string, required)
- Response: \`{ "id": "xxx", "data": {...} }\`
- Status codes: 200, 404, 500

### Data Model

\`\`\`go
type Resource struct {
    ID        string    \`json:"id" db:"id"\`
    Name      string    \`json:"name" db:"name"\`
    CreatedAt time.Time \`json:"created_at" db:"created_at"\`
    UpdatedAt time.Time \`json:"updated_at" db:"updated_at"\`
}
\`\`\`

### Database Schema

\`\`\`sql
CREATE TABLE resources (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_resources_name ON resources(name);
\`\`\`

## Testing Strategy

### Unit Tests
- Test case 1: Description
- Test case 2: Description
- Test case 3: Description

### Integration Tests
- Integration scenario 1
- Integration scenario 2

### E2E Tests
- E2E scenario 1: User journey description

### BDD Scenarios

\`\`\`gherkin
Feature: $(echo "$FEATURE_NAME" | sed 's/-/ /g' | sed 's/\b\(.\)/\u\1/g')

  Scenario: Happy path
    Given precondition
    When user action
    Then expected result

  Scenario: Error case
    Given precondition
    When invalid action
    Then error message shown
\`\`\`

## Implementation Plan

### Phase 1: Setup
- [ ] Create service structure
- [ ] Set up database schema
- [ ] Configure API routes

### Phase 2: Core Implementation
- [ ] Implement business logic
- [ ] Add validation
- [ ] Error handling

### Phase 3: Testing
- [ ] Write unit tests
- [ ] Write integration tests
- [ ] Write E2E tests

### Phase 4: Documentation
- [ ] API documentation
- [ ] Setup guide
- [ ] Troubleshooting guide

## Dependencies

- Dependency 1: Description
- Dependency 2: Description

## Risks and Mitigations

### Risk 1: Description
**Impact**: High/Medium/Low
**Probability**: High/Medium/Low
**Mitigation**: How to address

### Risk 2: Description
**Impact**: High/Medium/Low
**Probability**: High/Medium/Low
**Mitigation**: How to address

## Open Questions

1. Question 1?
2. Question 2?
3. Question 3?

## References

- Link to related docs
- Link to design discussions
- Link to ADRs

---

**Last Updated**: $(get_date)
EOF
    print_success "Created specification: $SPEC_FILE"
else
    print_warning "Specification already exists: $SPEC_FILE"
fi

# Create tasks file
TASKS_FILE="$FEATURE_DIR/tasks.md"
if [ ! -f "$TASKS_FILE" ]; then
    cat > "$TASKS_FILE" << EOF
# Tasks: $(echo "$FEATURE_NAME" | sed 's/-/ /g' | sed 's/\b\(.\)/\u\1/g')

## Backlog

- [ ] Task 1: Description
- [ ] Task 2: Description
- [ ] Task 3: Description

## In Progress

Currently no tasks in progress.

## Completed

None yet.

---

**Last Updated**: $(get_date)
EOF
    print_success "Created tasks file: $TASKS_FILE"
fi

# Create checklist file
CHECKLIST_FILE="$FEATURE_DIR/checklist.md"
if [ ! -f "$CHECKLIST_FILE" ]; then
    cat > "$CHECKLIST_FILE" << EOF
# Feature Checklist: $(echo "$FEATURE_NAME" | sed 's/-/ /g' | sed 's/\b\(.\)/\u\1/g')

## Planning
- [ ] Feature spec complete
- [ ] Requirements reviewed
- [ ] Technical design approved
- [ ] Dependencies identified

## Implementation
- [ ] Code structure created
- [ ] Core logic implemented
- [ ] API endpoints implemented
- [ ] Database schema created
- [ ] Error handling added
- [ ] Logging added

## Testing
- [ ] Unit tests written (>80% coverage)
- [ ] Integration tests written
- [ ] E2E tests written
- [ ] BDD scenarios implemented
- [ ] Performance tested
- [ ] Security tested

## Documentation
- [ ] API documentation complete
- [ ] Setup guide written
- [ ] Configuration documented
- [ ] Troubleshooting guide created

## Review
- [ ] Code review passed
- [ ] Security review passed
- [ ] Performance benchmarks met
- [ ] Documentation reviewed

## Deployment
- [ ] Staging deployment successful
- [ ] Production deployment plan ready
- [ ] Rollback plan documented
- [ ] Monitoring configured

---

**Completion Date**: ___________
EOF
    print_success "Created checklist: $CHECKLIST_FILE"
fi

# Create implementation directory if service name provided
if [ -n "$SERVICE_NAME_CLEAN" ]; then
    IMPL_DIR="$PROJECT_ROOT/services/$SERVICE_NAME_CLEAN/features/$FEATURE_NAME_CLEAN"
    ensure_dir "$IMPL_DIR"
    
    # Create basic README
    README_FILE="$IMPL_DIR/README.md"
    if [ ! -f "$README_FILE" ]; then
        cat > "$README_FILE" << EOF
# Feature: $(echo "$FEATURE_NAME" | sed 's/-/ /g' | sed 's/\b\(.\)/\u\1/g')

Implementation for the $FEATURE_NAME_CLEAN feature in $SERVICE_NAME_CLEAN.

## Files

- \`handler.go\` - HTTP handlers
- \`service.go\` - Business logic
- \`repository.go\` - Data access
- \`models.go\` - Data models

## Related Specs

See [Feature Specification](../../../specs/features/$FEATURE_NAME_CLEAN/spec.md)

## Testing

Run tests:
\`\`\`bash
go test -v ./...
\`\`\`

Run with coverage:
\`\`\`bash
go test -cover -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
\`\`\`
EOF
        print_success "Created implementation README: $README_FILE"
    fi
fi

# Summary
echo ""
print_header "Feature Created Successfully"
echo ""
echo "Feature: $FEATURE_NAME_CLEAN"
echo "Location: $FEATURE_DIR"
echo ""
echo "Next steps:"
echo "  1. Edit spec: $SPEC_FILE"
echo "  2. Define tasks: $TASKS_FILE"
echo "  3. Track progress: $CHECKLIST_FILE"
if [ -n "$SERVICE_NAME_CLEAN" ]; then
    echo "  4. Implement in: $IMPL_DIR"
fi
echo ""
print_success "Ready to start development!"
