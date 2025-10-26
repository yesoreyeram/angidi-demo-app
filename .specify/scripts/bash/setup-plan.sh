#!/bin/bash

# setup-plan.sh - Sets up a development plan for the current phase

set -e

# Source common functions
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "${SCRIPT_DIR}/common.sh"

print_header "Setup Development Plan"

# Get project root
PROJECT_ROOT=$(get_project_root)

# Determine current phase
echo ""
print_step "Determining current phase..."

PHASES_DIR="$PROJECT_ROOT/specs/phases"
CURRENT_PHASE=""

# Look for active phase (has status: In Progress)
for phase_dir in "$PHASES_DIR"/phase-*; do
    if [ -d "$phase_dir" ]; then
        if grep -q "Status.*In Progress" "$phase_dir/README.md" 2>/dev/null; then
            CURRENT_PHASE=$(basename "$phase_dir")
            break
        fi
    fi
done

# If no in-progress phase, ask user
if [ -z "$CURRENT_PHASE" ]; then
    echo "No phase currently in progress."
    echo ""
    echo "Available phases:"
    for phase_dir in "$PHASES_DIR"/phase-*; do
        if [ -d "$phase_dir" ]; then
            phase_name=$(basename "$phase_dir")
            phase_title=$(grep "^# " "$phase_dir/README.md" 2>/dev/null | head -1 | sed 's/^# //')
            echo "  - $phase_name: $phase_title"
        fi
    done
    echo ""
    
    read -p "Enter phase to plan (e.g., phase-1-basic-monolith): " CURRENT_PHASE
    
    if [ ! -d "$PHASES_DIR/$CURRENT_PHASE" ]; then
        print_error "Phase not found: $CURRENT_PHASE"
        exit 1
    fi
fi

print_success "Planning for: $CURRENT_PHASE"

# Create plan directory
PLAN_DIR="$PROJECT_ROOT/.specify/plans/$CURRENT_PHASE"
ensure_dir "$PLAN_DIR"

# Create plan file
PLAN_FILE="$PLAN_DIR/plan.md"

if [ -f "$PLAN_FILE" ]; then
    print_warning "Plan already exists: $PLAN_FILE"
    if ! confirm "Overwrite existing plan?"; then
        print_info "Keeping existing plan"
        exit 0
    fi
    backup_file "$PLAN_FILE"
fi

# Extract phase information
PHASE_README="$PHASES_DIR/$CURRENT_PHASE/README.md"
PHASE_TITLE=$(grep "^# " "$PHASE_README" | head -1 | sed 's/^# //')
PHASE_DURATION=$(grep "Duration:" "$PHASE_README" | head -1 | sed 's/.*Duration: //' | sed 's/\*\*//')

print_step "Creating development plan..."

# Create the plan
cat > "$PLAN_FILE" << EOF
# Development Plan: $PHASE_TITLE

**Phase**: $CURRENT_PHASE
**Duration**: $PHASE_DURATION
**Created**: $(get_date)
**Status**: Draft

## Overview

This plan outlines the development approach for $PHASE_TITLE.

## Goals

Based on the phase specifications, the primary goals are:

1. [ ] Goal 1 (from phase README)
2. [ ] Goal 2 (from phase README)
3. [ ] Goal 3 (from phase README)

## Timeline

**Start Date**: $(get_date)
**Target End Date**: TBD (based on $PHASE_DURATION)

### Week 1
- [ ] Setup development environment
- [ ] Review phase specifications
- [ ] Create initial project structure
- [ ] Setup CI/CD pipeline

### Week 2
- [ ] Implement core functionality
- [ ] Write unit tests
- [ ] Write integration tests
- [ ] Documentation

### Additional Weeks
Plan based on phase duration...

## Technical Tasks

### Infrastructure
- [ ] Task 1: Description
- [ ] Task 2: Description
- [ ] Task 3: Description

### Backend Development
- [ ] Task 1: Description
- [ ] Task 2: Description
- [ ] Task 3: Description

### Frontend Development
- [ ] Task 1: Description
- [ ] Task 2: Description
- [ ] Task 3: Description

### Testing
- [ ] Unit tests (>80% coverage)
- [ ] Integration tests
- [ ] E2E tests
- [ ] Performance tests
- [ ] Security tests

### Documentation
- [ ] API documentation
- [ ] Architecture documentation
- [ ] Setup guides
- [ ] Troubleshooting guides

## Dependencies

- External dependency 1
- External dependency 2
- Blockers from previous phases

## Risks

### Risk 1: Description
**Impact**: High/Medium/Low
**Mitigation**: How to address

### Risk 2: Description
**Impact**: High/Medium/Low
**Mitigation**: How to address

## Daily Standups

Track daily progress here:

### $(get_date)
- **Done**: 
- **Today**: 
- **Blockers**: 

## Weekly Reviews

### Week 1 Review (Date: _______)
- **Completed**: 
- **Challenges**: 
- **Adjustments**: 

## Resources

- [Phase Specification]($PHASES_DIR/$CURRENT_PHASE/README.md)
- [System Design Concepts](../../specs/system-design/concepts.md)
- [Architecture Overview](../../specs/architecture/overview.md)

## Notes

Add any additional notes, decisions, or observations here.

---

**Last Updated**: $(get_date)
EOF

print_success "Created plan: $PLAN_FILE"

# Create tasks breakdown
TASKS_FILE="$PLAN_DIR/tasks.md"
cat > "$TASKS_FILE" << EOF
# Tasks Breakdown: $PHASE_TITLE

## Backlog

### High Priority
- [ ] Task 1
- [ ] Task 2
- [ ] Task 3

### Medium Priority
- [ ] Task 4
- [ ] Task 5

### Low Priority
- [ ] Task 6
- [ ] Task 7

## Sprint 1 (Dates: _____ to _____)

### In Progress
- Nothing yet

### Completed
- Nothing yet

## Sprint 2 (Dates: _____ to _____)

### In Progress
- Nothing yet

### Completed
- Nothing yet

---

**Last Updated**: $(get_date)
EOF

print_success "Created tasks: $TASKS_FILE"

# Create daily log
LOG_FILE="$PLAN_DIR/log.md"
cat > "$LOG_FILE" << EOF
# Development Log: $PHASE_TITLE

## $(get_date)

**Hours**: 0
**Focus**: Setup

### Accomplished
- Created development plan
- Reviewed phase specifications

### Challenges
- None yet

### Tomorrow
- Begin implementation

---

## Template for Future Entries

## YYYY-MM-DD

**Hours**: X
**Focus**: Area of focus

### Accomplished
- What was completed

### Challenges
- Any issues or blockers

### Tomorrow
- What's planned next

---
EOF

print_success "Created log: $LOG_FILE"

# Summary
echo ""
print_header "Plan Created Successfully"
echo ""
echo "Phase: $CURRENT_PHASE"
echo "Plan Location: $PLAN_DIR"
echo ""
echo "Files created:"
echo "  - plan.md   : Development plan and timeline"
echo "  - tasks.md  : Task breakdown and tracking"
echo "  - log.md    : Daily development log"
echo ""
echo "Next steps:"
echo "  1. Review and customize: $PLAN_FILE"
echo "  2. Break down tasks: $TASKS_FILE"
echo "  3. Start logging progress: $LOG_FILE"
echo "  4. Run: ./.specify/scripts/bash/check-prerequisites.sh"
echo ""
print_success "Ready to start development!"
