#!/bin/bash

# validate-specs.sh - Validates the specification structure and content

set -e

SPECS_DIR="specs"
SPECIFY_DIR=".specify"
LEGACY_SPEC_DIR=".spec"
ERRORS=0
WARNINGS=0

echo "🔍 Validating Angidi Specifications..."
echo ""

# Color codes
RED='\033[0;31m'
YELLOW='\033[1;33m'
GREEN='\033[0;32m'
NC='\033[0m'

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

# Check specs directory structure
echo "📁 Checking specs directory structure..."

if [ ! -d "$SPECS_DIR" ]; then
    print_warning "Specs directory '$SPECS_DIR' not found (will migrate from .spec/)"
fi

if [ -d "$SPECS_DIR" ]; then
    required_specs_dirs=(
        "requirements"
        "architecture"
        "phases"
    )
    
    for dir in "${required_specs_dirs[@]}"; do
        if [ ! -d "$SPECS_DIR/$dir" ]; then
            print_error "Required directory '$SPECS_DIR/$dir' not found"
        else
            print_success "Directory '$SPECS_DIR/$dir' exists"
        fi
    done
fi

# Check .specify directory structure
echo ""
echo "📁 Checking .specify directory structure..."

required_specify_dirs=(
    "memory"
    "scripts"
    "templates"
)

for dir in "${required_specify_dirs[@]}"; do
    if [ ! -d "$SPECIFY_DIR/$dir" ]; then
        print_error "Required directory '$SPECIFY_DIR/$dir' not found"
    else
        print_success "Directory '$SPECIFY_DIR/$dir' exists"
    fi
done

# Check .github structure
echo ""
echo "📁 Checking .github directory structure..."

if [ ! -d ".github/prompts" ]; then
    print_error ".github/prompts directory not found"
else
    print_success ".github/prompts directory exists"
    
    # Check for agent prompts
    prompt_count=$(find .github/prompts -name "speckit.*.prompt.md" | wc -l)
    if [ "$prompt_count" -eq 0 ]; then
        print_warning "No agent prompts found in .github/prompts/"
    else
        print_success "Found $prompt_count agent prompt(s)"
    fi
fi

if [ ! -f ".github/copilot-instructions.md" ]; then
    print_error ".github/copilot-instructions.md not found"
else
    print_success ".github/copilot-instructions.md exists"
fi

# Check required files
echo ""
echo "📄 Checking required files..."

required_files=(
    "AGENTS.md"
    "README.md"
    ".specify/memory/project-context.md"
)

for file in "${required_files[@]}"; do
    if [ ! -f "$file" ]; then
        print_error "Required file '$file' not found"
    else
        if [ ! -s "$file" ]; then
            print_warning "File '$file' is empty"
        else
            print_success "File '$file' exists and has content"
        fi
    fi
done

# Check markdown formatting
echo ""
echo "📝 Checking markdown formatting..."

md_files=$(find . -name "*.md" ! -path "./.git/*" ! -path "./node_modules/*")

for md_file in $md_files; do
    # Check for trailing whitespace
    if grep -q " $" "$md_file"; then
        print_warning "Trailing whitespace found in $md_file"
    fi
    
    # Check for proper heading
    first_heading=$(grep -m 1 "^#" "$md_file" 2>/dev/null || true)
    if [ -z "$first_heading" ]; then
        print_warning "No heading found in $md_file"
    fi
done

# Check VSCode settings
echo ""
echo "⚙️  Checking VSCode configuration..."

if [ ! -f ".vscode/settings.json" ]; then
    print_warning ".vscode/settings.json not found (optional)"
else
    print_success ".vscode/settings.json exists"
fi

# Summary
echo ""
echo "================================"
echo "📊 Validation Summary"
echo "================================"

if [ $ERRORS -eq 0 ] && [ $WARNINGS -eq 0 ]; then
    print_success "All checks passed! ✨"
    exit 0
elif [ $ERRORS -eq 0 ]; then
    echo -e "${YELLOW}⚠️  Validation completed with $WARNINGS warning(s)${NC}"
    exit 0
else
    echo -e "${RED}❌ Validation failed with $ERRORS error(s) and $WARNINGS warning(s)${NC}"
    exit 1
fi
