#!/bin/bash

# update-agent-context.sh - Updates agent context with current project state

set -e

# Source common functions
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "${SCRIPT_DIR}/common.sh"

print_header "Update Agent Context"

# Get project root
PROJECT_ROOT=$(get_project_root)
CONTEXT_FILE="$PROJECT_ROOT/.specify/memory/project-context.md"

print_step "Gathering project information..."

# Backup existing context
if [ -f "$CONTEXT_FILE" ]; then
    backup_file "$CONTEXT_FILE"
fi

# Get current phase
CURRENT_PHASE=""
PHASE_STATUS=""
PHASES_DIR="$PROJECT_ROOT/specs/phases"

for phase_dir in "$PHASES_DIR"/phase-*; do
    if [ -d "$phase_dir" ]; then
        phase_name=$(basename "$phase_dir")
        status=$(grep "Status:" "$phase_dir/README.md" 2>/dev/null | head -1 | sed 's/.*Status.*: //' | sed 's/\*\*//')
        
        if [[ "$status" == *"In Progress"* ]]; then
            CURRENT_PHASE=$phase_name
            PHASE_STATUS="In Progress"
            break
        elif [[ "$status" == *"Complete"* ]]; then
            LAST_COMPLETED_PHASE=$phase_name
        fi
    fi
done

if [ -z "$CURRENT_PHASE" ]; then
    CURRENT_PHASE=${LAST_COMPLETED_PHASE:-"phase-0-planning"}
    PHASE_STATUS="Complete"
fi

# Count files and stats
print_step "Analyzing codebase..."

TOTAL_GO_FILES=$(find "$PROJECT_ROOT" -name "*.go" ! -path "*/vendor/*" ! -path "*/node_modules/*" 2>/dev/null | wc -l)
TOTAL_TS_FILES=$(find "$PROJECT_ROOT" -name "*.ts" -o -name "*.tsx" ! -path "*/node_modules/*" 2>/dev/null | wc -l)
TOTAL_MD_FILES=$(find "$PROJECT_ROOT" -name "*.md" ! -path "*/node_modules/*" ! -path "*/.git/*" 2>/dev/null | wc -l)

# Count services
SERVICE_COUNT=0
if [ -d "$PROJECT_ROOT/services" ]; then
    SERVICE_COUNT=$(find "$PROJECT_ROOT/services" -maxdepth 1 -type d | tail -n +2 | wc -l)
fi

# Get git info
GIT_BRANCH=$(get_current_branch)
TOTAL_COMMITS=$(git rev-list --count HEAD 2>/dev/null || echo "0")
LAST_COMMIT=$(git log -1 --format="%h - %s" 2>/dev/null || echo "No commits")

# Get recent ADRs
RECENT_ADRS=""
ADR_DIR="$PROJECT_ROOT/specs/architecture/decisions"
if [ -d "$ADR_DIR" ]; then
    RECENT_ADRS=$(find "$ADR_DIR" -name "*.md" -type f | sort | tail -3 | while read adr; do
        basename "$adr" .md
    done)
fi

# Create updated context
print_step "Updating context file..."

cat > "$CONTEXT_FILE" << EOF
# Project Context and Memory

This document maintains the current state and context of the Angidi e-commerce platform project.

**Last Updated**: $(get_timestamp)
**Auto-generated**: Yes (run ./.specify/scripts/bash/update-agent-context.sh to update)

## Current State

### Active Phase
**Phase**: $CURRENT_PHASE
**Status**: $PHASE_STATUS

### Project Statistics

- **Total Commits**: $TOTAL_COMMITS
- **Current Branch**: $GIT_BRANCH
- **Last Commit**: $LAST_COMMIT

### Codebase Metrics

- **Go Files**: $TOTAL_GO_FILES
- **TypeScript Files**: $TOTAL_TS_FILES
- **Documentation Files**: $TOTAL_MD_FILES
- **Services**: $SERVICE_COUNT

## Recent Activity

### Latest Commits
\`\`\`
$(git log -5 --oneline 2>/dev/null || echo "No commits")
\`\`\`

### Recent Architecture Decisions

$( [ -n "$RECENT_ADRS" ] && echo "$RECENT_ADRS" || echo "No ADRs yet" )

## Active Work

### Current Focus

Based on $CURRENT_PHASE, the current focus areas are:

1. [ ] Primary objective 1
2. [ ] Primary objective 2
3. [ ] Primary objective 3

### In Progress Features

- Feature 1 (if any)
- Feature 2 (if any)

### Blockers

- None currently identified

## Key Decisions

### Architecture Decisions

Review [Architecture Decisions](../architecture/decisions/) for complete ADR history.

### Technology Stack

**Backend**:
- Language: Go $(go version 2>/dev/null | cut -d' ' -f3 || echo "not installed")
- Databases: PostgreSQL, MongoDB, Redis, Elasticsearch
- Message Broker: Apache Kafka

**Frontend**:
- Framework: React + Next.js
- Language: TypeScript
- Styling: Tailwind CSS

**Infrastructure**:
- Containers: Docker $(docker --version 2>/dev/null | cut -d' ' -f3 | tr -d ',' || echo "not installed")
- Orchestration: Kubernetes
- CI/CD: GitHub Actions

## Open Questions

1. Question 1 (if any)
2. Question 2 (if any)
3. Question 3 (if any)

## Risks and Issues

### Active Risks

None currently identified.

### Known Issues

None currently identified.

## Dependencies

### External Dependencies

- List external services or APIs
- Third-party integrations

### Internal Dependencies

- Service dependencies
- Shared libraries

## Metrics and Goals

### Performance Targets

- API Response Time: p95 < 200ms
- Database Query Time: p95 < 100ms
- Error Rate: < 1%

### Quality Targets

- Code Coverage: > 80%
- Test Pass Rate: 100%
- Security Scan: No critical issues

### Phase Completion

- Phases Completed: $(find "$PHASES_DIR" -name "README.md" -exec grep -l "Status.*Complete" {} \; 2>/dev/null | wc -l) / 13
- Current Phase Progress: $(grep -c "\[x\]" "$PHASES_DIR/$CURRENT_PHASE/README.md" 2>/dev/null || echo "0") tasks completed

## Team Context

### Roles

- Developer: Implementation
- Architect: Design decisions
- DevOps: Infrastructure
- QA: Testing strategy

### Communication

- All design decisions documented in ADRs
- Weekly progress reviews
- Phase retrospectives

## Resources

### Documentation

- Main README: [/README.md](../../README.md)
- Specs: [/specs/](../../specs/)
- Agents: [/AGENTS.md](../../AGENTS.md)

### Tools and Scripts

- Prerequisites Check: \`./.specify/scripts/bash/check-prerequisites.sh\`
- Create Feature: \`./.specify/scripts/bash/create-new-feature.sh\`
- Setup Plan: \`./.specify/scripts/bash/setup-plan.sh\`
- Validate Specs: \`./.specify/scripts/validate-specs.sh\`

## Next Steps

### Immediate Actions

1. [ ] Complete current phase objectives
2. [ ] Update documentation
3. [ ] Run tests and validation

### Upcoming Phases

- Next Phase: $(ls "$PHASES_DIR" | grep "^phase-" | sort | grep -A1 "$CURRENT_PHASE" | tail -1)
- Timeline: See [Phase Roadmap](../../specs/phases/README.md)

## Notes

Additional context and observations:

- Project is in active development
- Following specification-driven development approach
- Emphasis on learning and documentation

---

**Maintained by**: Automated script
**Review Frequency**: Updated on demand
**Next Review**: Before phase transition
EOF

print_success "Updated context file: $CONTEXT_FILE"

# Update agent file with current context
AGENT_CONTEXT="$PROJECT_ROOT/.github/prompts/current-context.md"
cat > "$AGENT_CONTEXT" << EOF
# Current Project Context

**Last Updated**: $(get_timestamp)

## Quick Facts

- **Current Phase**: $CURRENT_PHASE
- **Status**: $PHASE_STATUS
- **Branch**: $GIT_BRANCH
- **Services**: $SERVICE_COUNT
- **Go Files**: $TOTAL_GO_FILES
- **TS Files**: $TOTAL_TS_FILES

## Focus

Based on the current phase, agents should focus on:

1. Reviewing phase objectives in \`specs/phases/$CURRENT_PHASE/\`
2. Following established patterns from previous phases
3. Maintaining consistency with project constitution
4. Ensuring comprehensive testing and documentation

## Key Resources

- Project Context: \`.specify/memory/project-context.md\`
- Constitution: \`.specify/memory/constitution.md\`
- Phase Spec: \`specs/phases/$CURRENT_PHASE/README.md\`

---

This file is auto-generated. Run \`./.specify/scripts/bash/update-agent-context.sh\` to update.
EOF

print_success "Updated agent context: $AGENT_CONTEXT"

# Summary
echo ""
print_header "Context Updated Successfully"
echo ""
echo "Current Phase: $CURRENT_PHASE ($PHASE_STATUS)"
echo "Files: $TOTAL_GO_FILES Go, $TOTAL_TS_FILES TypeScript, $TOTAL_MD_FILES Docs"
echo "Services: $SERVICE_COUNT"
echo ""
echo "Updated files:"
echo "  - $CONTEXT_FILE"
echo "  - $AGENT_CONTEXT"
echo ""
print_success "Agent context is up to date!"
