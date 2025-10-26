#!/bin/bash

# check-prerequisites.sh - Checks if all required tools are installed

set -e

# Source common functions
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "${SCRIPT_DIR}/common.sh"

print_header "Checking Prerequisites"

# Track if all prerequisites are met
ALL_MET=true

# Function to check if command exists
check_command() {
    local cmd=$1
    local version_cmd=$2
    local required_version=$3
    local name=$4

    if command -v "$cmd" &> /dev/null; then
        local version
        version=$($version_cmd 2>&1 || echo "unknown")
        print_success "$name is installed: $version"
        return 0
    else
        print_error "$name is NOT installed"
        ALL_MET=false
        return 1
    fi
}

# Function to check version meets requirement
check_version() {
    local current=$1
    local required=$2
    local name=$3

    if [ "$current" == "unknown" ]; then
        print_warning "$name version could not be determined"
        return 0
    fi

    # Simple version comparison (works for major.minor.patch)
    if [ "$(printf '%s\n' "$required" "$current" | sort -V | head -n1)" = "$required" ]; then
        return 0
    else
        print_warning "$name version $current is below recommended $required"
        return 1
    fi
}

echo ""
echo "=== Core Tools ==="

# Git
check_command "git" "git version 2>&1 | awk '{print \$3}'" "2.0.0" "Git"

# Docker
check_command "docker" "docker version --format '{{.Server.Version}}' 2>&1 || echo 'unknown'" "20.0.0" "Docker"

# Docker Compose
check_command "docker-compose" "docker-compose version --short 2>&1 || echo 'unknown'" "1.29.0" "Docker Compose"

echo ""
echo "=== Backend Tools ==="

# Go
if check_command "go" "go version | cut -d' ' -f3 | sed 's/go//'" "1.21.0" "Go"; then
    # Check GOPATH
    if [ -z "$GOPATH" ]; then
        print_warning "GOPATH is not set"
    else
        print_success "GOPATH: $GOPATH"
    fi
fi

# Make
check_command "make" "make --version | head -n1 | cut -d' ' -f3" "3.0" "Make"

echo ""
echo "=== Frontend Tools ==="

# Node.js
check_command "node" "node --version | sed 's/v//'" "18.0.0" "Node.js"

# npm
check_command "npm" "npm --version" "9.0.0" "npm"

# Optional: yarn
if command -v yarn &> /dev/null; then
    check_command "yarn" "yarn --version" "1.22.0" "Yarn"
fi

echo ""
echo "=== Database Tools (Optional) ==="

# PostgreSQL client
if command -v psql &> /dev/null; then
    check_command "psql" "psql --version | cut -d' ' -f3" "14.0" "PostgreSQL Client"
else
    print_info "PostgreSQL client not installed (optional for local dev)"
fi

# MongoDB client
if command -v mongosh &> /dev/null; then
    check_command "mongosh" "mongosh --version" "1.0.0" "MongoDB Shell"
elif command -v mongo &> /dev/null; then
    check_command "mongo" "mongo --version | head -n1 | cut -d' ' -f4" "5.0.0" "MongoDB Client"
else
    print_info "MongoDB client not installed (optional for local dev)"
fi

# Redis client
if command -v redis-cli &> /dev/null; then
    check_command "redis-cli" "redis-cli --version | cut -d' ' -f2" "6.0.0" "Redis CLI"
else
    print_info "Redis CLI not installed (optional for local dev)"
fi

echo ""
echo "=== Kubernetes Tools (for later phases) ==="

# kubectl
if command -v kubectl &> /dev/null; then
    check_command "kubectl" "kubectl version --client --short 2>/dev/null | cut -d' ' -f3 | sed 's/v//'" "1.28.0" "kubectl"
else
    print_info "kubectl not installed (needed for Phase 11+)"
fi

# Helm
if command -v helm &> /dev/null; then
    check_command "helm" "helm version --short | cut -d' ' -f1 | sed 's/v//'" "3.12.0" "Helm"
else
    print_info "Helm not installed (needed for Phase 11+)"
fi

echo ""
echo "=== Code Quality Tools ==="

# golangci-lint
if command -v golangci-lint &> /dev/null; then
    check_command "golangci-lint" "golangci-lint --version | grep -oP '(?<=version )[0-9.]+'" "1.54.0" "golangci-lint"
else
    print_warning "golangci-lint not installed (recommended for Go development)"
fi

# ESLint (check in node_modules if not global)
if command -v eslint &> /dev/null; then
    print_success "ESLint is installed globally"
elif [ -f "node_modules/.bin/eslint" ]; then
    print_success "ESLint is installed locally"
else
    print_info "ESLint not found (will be installed with npm install)"
fi

echo ""
echo "=== System Requirements ==="

# Check available memory
if command -v free &> /dev/null; then
    TOTAL_MEM=$(free -g | awk '/^Mem:/{print $2}')
    if [ "$TOTAL_MEM" -ge 8 ]; then
        print_success "System memory: ${TOTAL_MEM}GB (>= 8GB required)"
    else
        print_warning "System memory: ${TOTAL_MEM}GB (8GB recommended)"
    fi
fi

# Check available disk space
if command -v df &> /dev/null; then
    AVAILABLE_SPACE=$(df -BG . | awk 'NR==2 {print $4}' | sed 's/G//')
    if [ "$AVAILABLE_SPACE" -ge 20 ]; then
        print_success "Available disk space: ${AVAILABLE_SPACE}GB (>= 20GB required)"
    else
        print_warning "Available disk space: ${AVAILABLE_SPACE}GB (20GB recommended)"
    fi
fi

echo ""
echo "================================"

if [ "$ALL_MET" = true ]; then
    print_success "All required prerequisites are met! ✨"
    echo ""
    echo "You're ready to start development!"
    echo "Next steps:"
    echo "  1. Review specs/phases/phase-0-planning/"
    echo "  2. Run: ./.specify/scripts/bash/setup-plan.sh"
    echo "  3. Start Phase 1 development"
    exit 0
else
    print_error "Some prerequisites are missing."
    echo ""
    echo "Please install missing tools before continuing."
    echo "Refer to README.md for installation instructions."
    exit 1
fi
