#!/bin/bash

# validate.sh - Validates the spec kit structure and content

set -e

SPEC_DIR=".spec"
ERRORS=0
WARNINGS=0

echo "🔍 Validating Angidi Spec Kit..."
echo ""

# Color codes for output
RED='\033[0;31m'
YELLOW='\033[1;33m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color

# Function to print error
print_error() {
    echo -e "${RED}❌ ERROR: $1${NC}"
    ((ERRORS++))
}

# Function to print warning
print_warning() {
    echo -e "${YELLOW}⚠️  WARNING: $1${NC}"
    ((WARNINGS++))
}

# Function to print success
print_success() {
    echo -e "${GREEN}✅ $1${NC}"
}

# Check if .spec directory exists
if [ ! -d "$SPEC_DIR" ]; then
    print_error "Spec directory '$SPEC_DIR' does not exist"
    exit 1
fi

print_success "Spec directory exists"

# Check required top-level directories
echo ""
echo "📁 Checking directory structure..."

required_dirs=(
    "system-design"
    "requirements"
    "architecture"
    "phases"
    "scripts"
)

for dir in "${required_dirs[@]}"; do
    if [ ! -d "$SPEC_DIR/$dir" ]; then
        print_error "Required directory '$dir' not found"
    else
        print_success "Directory '$dir' exists"
    fi
done

# Check required documents
echo ""
echo "📄 Checking required documents..."

required_docs=(
    "README.md"
    "system-design/concepts.md"
    "requirements/functional.md"
    "requirements/non-functional.md"
    "architecture/overview.md"
    "phases/README.md"
)

for doc in "${required_docs[@]}"; do
    if [ ! -f "$SPEC_DIR/$doc" ]; then
        print_error "Required document '$doc' not found"
    else
        # Check if file is empty
        if [ ! -s "$SPEC_DIR/$doc" ]; then
            print_warning "Document '$doc' is empty"
        else
            print_success "Document '$doc' exists and has content"
        fi
    fi
done

# Check for broken links in markdown files
echo ""
echo "🔗 Checking for broken internal links..."

# Find all markdown files
md_files=$(find "$SPEC_DIR" -name "*.md")

for md_file in $md_files; do
    # Extract markdown links [text](path)
    links=$(grep -oP '\[.*?\]\(\K[^)]+' "$md_file" 2>/dev/null || true)
    
    for link in $links; do
        # Skip external links (http, https, mailto)
        if [[ $link =~ ^https?:// ]] || [[ $link =~ ^mailto: ]]; then
            continue
        fi
        
        # Skip anchors
        if [[ $link =~ ^# ]]; then
            continue
        fi
        
        # Resolve relative path
        dir=$(dirname "$md_file")
        target="$dir/$link"
        
        # Check if target exists
        if [ ! -f "$target" ] && [ ! -d "$target" ]; then
            print_warning "Broken link in $md_file: $link"
        fi
    done
done

# Check markdown formatting
echo ""
echo "📝 Checking markdown formatting..."

for md_file in $md_files; do
    # Check for proper heading hierarchy (should start with #)
    first_heading=$(grep -m 1 "^#" "$md_file" 2>/dev/null || true)
    if [ -z "$first_heading" ]; then
        print_warning "No heading found in $md_file"
    fi
    
    # Check for trailing whitespace
    if grep -q " $" "$md_file"; then
        print_warning "Trailing whitespace found in $md_file"
    fi
done

# Check phase structure
echo ""
echo "🎯 Checking phase structure..."

phase_dirs=$(find "$SPEC_DIR/phases" -maxdepth 1 -type d -name "phase-*" | sort)

if [ -z "$phase_dirs" ]; then
    print_warning "No phase directories found"
else
    for phase_dir in $phase_dirs; do
        phase_name=$(basename "$phase_dir")
        
        # Check for README.md in each phase
        if [ ! -f "$phase_dir/README.md" ]; then
            print_warning "Phase $phase_name is missing README.md"
        else
            print_success "Phase $phase_name has README.md"
        fi
    done
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
