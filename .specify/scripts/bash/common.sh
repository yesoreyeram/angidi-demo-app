#!/bin/bash

# common.sh - Common functions and utilities for spec kit scripts

# Color codes for output
export RED='\033[0;31m'
export YELLOW='\033[1;33m'
export GREEN='\033[0;32m'
export BLUE='\033[0;34m'
export CYAN='\033[0;36m'
export NC='\033[0m' # No Color

# Error tracking
export ERRORS=0
export WARNINGS=0

# Function to print colored output
print_error() {
    echo -e "${RED}❌ ERROR: $1${NC}"
    ((ERRORS++))
}

print_warning() {
    echo -e "${YELLOW}⚠️  WARNING: $1${NC}"
    ((WARNINGS++))
}

print_success() {
    echo -e "${GREEN}✅ $1${NC}"
}

print_info() {
    echo -e "${BLUE}ℹ️  $1${NC}"
}

print_header() {
    echo -e "${CYAN}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
    echo -e "${CYAN}  $1${NC}"
    echo -e "${CYAN}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
}

print_step() {
    echo -e "${BLUE}➜ $1${NC}"
}

# Function to check if a file exists
check_file_exists() {
    local file=$1
    if [ ! -f "$file" ]; then
        print_error "File '$file' not found"
        return 1
    fi
    return 0
}

# Function to check if a directory exists
check_dir_exists() {
    local dir=$1
    if [ ! -d "$dir" ]; then
        print_error "Directory '$dir' not found"
        return 1
    fi
    return 0
}

# Function to create directory if it doesn't exist
ensure_dir() {
    local dir=$1
    if [ ! -d "$dir" ]; then
        mkdir -p "$dir"
        print_success "Created directory: $dir"
    fi
}

# Function to backup a file
backup_file() {
    local file=$1
    if [ -f "$file" ]; then
        local backup="${file}.backup.$(date +%Y%m%d_%H%M%S)"
        cp "$file" "$backup"
        print_info "Backed up $file to $backup"
    fi
}

# Function to get project root
get_project_root() {
    git rev-parse --show-toplevel 2>/dev/null || echo "."
}

# Function to validate phase number
validate_phase_number() {
    local phase=$1
    if ! [[ "$phase" =~ ^[0-9]+$ ]]; then
        print_error "Phase number must be a number: $phase"
        return 1
    fi
    if [ "$phase" -lt 0 ] || [ "$phase" -gt 13 ]; then
        print_error "Phase number must be between 0 and 13: $phase"
        return 1
    fi
    return 0
}

# Function to sanitize name (convert to lowercase with hyphens)
sanitize_name() {
    local name=$1
    echo "$name" | tr '[:upper:]' '[:lower:]' | tr ' _' '--' | sed 's/[^a-z0-9-]//g'
}

# Function to get current git branch
get_current_branch() {
    git branch --show-current 2>/dev/null || echo "main"
}

# Function to check if git working directory is clean
is_git_clean() {
    if [ -n "$(git status --porcelain)" ]; then
        return 1
    fi
    return 0
}

# Function to confirm action
confirm() {
    local prompt=$1
    local default=${2:-n}
    
    if [ "$default" = "y" ]; then
        prompt="$prompt [Y/n]: "
    else
        prompt="$prompt [y/N]: "
    fi
    
    read -p "$prompt" -n 1 -r
    echo
    
    if [ "$default" = "y" ]; then
        [[ ! $REPLY =~ ^[Nn]$ ]]
    else
        [[ $REPLY =~ ^[Yy]$ ]]
    fi
}

# Function to read user input with default
read_with_default() {
    local prompt=$1
    local default=$2
    local var_name=$3
    
    read -p "$prompt [$default]: " input
    eval "$var_name=\"${input:-$default}\""
}

# Function to validate file is markdown
is_markdown() {
    local file=$1
    [[ "$file" =~ \.md$ ]]
}

# Function to count lines in file
count_lines() {
    local file=$1
    if [ -f "$file" ]; then
        wc -l < "$file"
    else
        echo "0"
    fi
}

# Function to get file size in human readable format
get_file_size() {
    local file=$1
    if [ -f "$file" ]; then
        du -h "$file" | cut -f1
    else
        echo "0"
    fi
}

# Function to check if running in CI/CD
is_ci() {
    [ -n "$CI" ] || [ -n "$GITHUB_ACTIONS" ] || [ -n "$GITLAB_CI" ]
}

# Function to get timestamp
get_timestamp() {
    date +"%Y-%m-%d %H:%M:%S"
}

# Function to get date only
get_date() {
    date +"%Y-%m-%d"
}

# Function to log message to file
log_to_file() {
    local message=$1
    local logfile=${2:-.specify/logs/script.log}
    
    ensure_dir "$(dirname "$logfile")"
    echo "[$(get_timestamp)] $message" >> "$logfile"
}

# Function to check if command exists
command_exists() {
    command -v "$1" &> /dev/null
}

# Function to get OS type
get_os() {
    case "$(uname -s)" in
        Linux*)     echo "linux";;
        Darwin*)    echo "macos";;
        CYGWIN*)    echo "windows";;
        MINGW*)     echo "windows";;
        *)          echo "unknown";;
    esac
}

# Function to initialize error tracking
init_error_tracking() {
    export ERRORS=0
    export WARNINGS=0
}

# Function to get error summary
get_error_summary() {
    if [ $ERRORS -eq 0 ] && [ $WARNINGS -eq 0 ]; then
        print_success "All checks passed! ✨"
        return 0
    elif [ $ERRORS -eq 0 ]; then
        print_warning "Completed with $WARNINGS warning(s)"
        return 0
    else
        print_error "Failed with $ERRORS error(s) and $WARNINGS warning(s)"
        return 1
    fi
}

# Function to exit with error summary
exit_with_summary() {
    echo ""
    echo "================================"
    echo "📊 Summary"
    echo "================================"
    
    if get_error_summary; then
        exit 0
    else
        exit 1
    fi
}

# Function to create directory structure
create_dir_structure() {
    local dirs=("$@")
    for dir in "${dirs[@]}"; do
        ensure_dir "$dir"
    done
}

# Export functions for use in other scripts
export -f print_error
export -f print_warning
export -f print_success
export -f print_info
export -f print_header
export -f print_step
export -f check_file_exists
export -f check_dir_exists
export -f ensure_dir
export -f backup_file
export -f get_project_root
export -f validate_phase_number
export -f sanitize_name
export -f get_current_branch
export -f is_git_clean
export -f confirm
export -f read_with_default
export -f is_markdown
export -f count_lines
export -f get_file_size
export -f is_ci
export -f get_timestamp
export -f get_date
export -f log_to_file
export -f command_exists
export -f get_os
export -f init_error_tracking
export -f get_error_summary
export -f exit_with_summary
export -f create_dir_structure
