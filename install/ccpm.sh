#!/bin/bash

# CCPM - Claude Code Project Management Installer
# Works with both Claude Code and OpenCode
# https://github.com/automazeio/ccpm

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

# Configuration
REPO_URL="https://github.com/automazeio/ccpm.git"
CCPM_VERSION="1.0.0"

# Helper functions
print_header() {
    echo ""
    echo -e "${CYAN}"
    echo " ██████╗ ██████╗██████╗ ███╗   ███╗"
    echo "██╔════╝██╔════╝██╔══██╗████╗ ████║"
    echo "██║     ██║     ██████╔╝██╔████╔██║"
    echo "╚██████╗╚██████╗██║     ██║ ╚═╝ ██║"
    echo " ╚═════╝ ╚═════╝╚═╝     ╚═╝     ╚═╝"
    echo -e "${NC}"
    echo "┌─────────────────────────────────────┐"
    echo "│  Claude Code Project Management     │"
    echo "│  Works with Claude Code & OpenCode  │"
    echo "└─────────────────────────────────────┘"
    echo ""
}

print_step() {
    echo -e "${BLUE}▶${NC} $1"
}

print_success() {
    echo -e "${GREEN}✓${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}⚠${NC} $1"
}

print_error() {
    echo -e "${RED}✗${NC} $1"
}

prompt_yes_no() {
    local prompt="$1"
    local default="${2:-y}"
    
    if [[ "$default" == "y" ]]; then
        prompt="$prompt [Y/n]: "
    else
        prompt="$prompt [y/N]: "
    fi
    
    read -p "$prompt" response
    response=${response:-$default}
    
    [[ "$response" =~ ^[Yy]$ ]]
}

# Check if we're in a valid project directory
check_project_directory() {
    if [[ "$(pwd)" == "$HOME" ]]; then
        print_error "Cannot install in home directory. Please cd into your project."
        exit 1
    fi
}

# Check for existing CCPM installation
check_existing_installation() {
    if [[ -d "ccpm" ]]; then
        print_warning "CCPM directory already exists."
        if prompt_yes_no "Update existing installation?" "y"; then
            return 0
        else
            print_error "Installation cancelled."
            exit 0
        fi
    fi
    return 1
}

# Download CCPM files
download_ccpm() {
    print_step "Downloading CCPM..."
    
    if [[ -d "ccpm" ]]; then
        # Update existing
        print_step "Updating existing CCPM files..."
        local temp_dir=$(mktemp -d)
        git clone --depth 1 "$REPO_URL" "$temp_dir" 2>/dev/null
        rm -rf ccpm
        mv "$temp_dir/ccpm" ./ccpm
        rm -rf "$temp_dir"
    else
        # Fresh install
        local temp_dir=$(mktemp -d)
        git clone --depth 1 "$REPO_URL" "$temp_dir" 2>/dev/null
        mv "$temp_dir/ccpm" ./ccpm
        rm -rf "$temp_dir"
    fi
    
    print_success "CCPM files downloaded"
}

# Create directory structure
create_directories() {
    print_step "Creating directory structure..."
    
    # Create .claude directories
    mkdir -p .claude/prds
    mkdir -p .claude/epics
    mkdir -p .claude/commands
    
    # Create .opencode directories
    mkdir -p .opencode/command
    
    print_success "Directories created"
}

setup_claude_commands() {
    print_step "Setting up Claude Code commands..."

    # Clean up existing command symlinks
    find .claude/commands -type l -delete 2>/dev/null || true
    find .claude/commands -type d -empty -delete 2>/dev/null || true

    # Generic subdirectory handling - auto-detects all command categories
    for subdir in ccpm/commands/*/; do
        if [[ -d "$subdir" ]]; then
            subdir_name=$(basename "$subdir")
            mkdir -p ".claude/commands/$subdir_name"
            for cmd in "$subdir"*.md; do
                if [[ -f "$cmd" ]]; then
                    ln -sf "../../../$cmd" ".claude/commands/$subdir_name/$(basename "$cmd")"
                fi
            done
        fi
    done

    # Top-level commands
    for cmd in ccpm/commands/*.md; do
        if [[ -f "$cmd" ]]; then
            ln -sf "../../$cmd" ".claude/commands/$(basename "$cmd")"
        fi
    done

    print_success "Claude Code commands linked"
}

setup_opencode_commands() {
    print_step "Setting up OpenCode commands..."

    # Clean up existing command symlinks
    find .opencode/command -type l -delete 2>/dev/null || true
    find .opencode/command -type d -empty -delete 2>/dev/null || true

    # Generic subdirectory handling
    for subdir in ccpm/commands/*/; do
        if [[ -d "$subdir" ]]; then
            subdir_name=$(basename "$subdir")
            mkdir -p ".opencode/command/$subdir_name"
            for cmd in "$subdir"*.md; do
                if [[ -f "$cmd" ]]; then
                    ln -sf "../../../$cmd" ".opencode/command/$subdir_name/$(basename "$cmd")"
                fi
            done
        fi
    done

    # Top-level commands
    for cmd in ccpm/commands/*.md; do
        if [[ -f "$cmd" ]]; then
            ln -sf "../../$cmd" ".opencode/command/$(basename "$cmd")"
        fi
    done

    print_success "OpenCode commands linked"
}

# Update .gitignore
update_gitignore() {
    print_step "Updating .gitignore..."
    
    local gitignore=".gitignore"
    local entries=(
        ""
        "# CCPM - Claude Code Project Management"
        ".claude/prds/"
        ".claude/epics/"
        ".claude/brainstorm/"
        ".claude/settings.local.json"
    )
    
    # Create .gitignore if it doesn't exist
    touch "$gitignore"
    
    # Add entries if they don't exist
    for entry in "${entries[@]}"; do
        if [[ -n "$entry" ]] && ! grep -qF "$entry" "$gitignore" 2>/dev/null; then
            echo "$entry" >> "$gitignore"
        fi
    done
    
    print_success ".gitignore updated"
}

# Check and setup GitHub CLI
setup_github() {
    print_step "Checking GitHub CLI..."
    
    if ! command -v gh &> /dev/null; then
        print_warning "GitHub CLI (gh) not found"
        if prompt_yes_no "Install GitHub CLI?" "y"; then
            if command -v brew &> /dev/null; then
                brew install gh
            elif command -v apt-get &> /dev/null; then
                sudo apt-get update && sudo apt-get install -y gh
            elif command -v dnf &> /dev/null; then
                sudo dnf install -y gh
            elif command -v pacman &> /dev/null; then
                sudo pacman -S github-cli
            else
                print_error "Could not install gh. Please install manually: https://cli.github.com/"
                return 1
            fi
        else
            print_warning "Skipping GitHub CLI setup"
            return 0
        fi
    fi
    
    print_success "GitHub CLI found: $(gh --version | head -1)"
    
    # Check authentication
    if ! gh auth status &> /dev/null; then
        print_warning "GitHub CLI not authenticated"
        if prompt_yes_no "Login to GitHub?" "y"; then
            gh auth login
        else
            print_warning "Skipping GitHub authentication"
        fi
    else
        print_success "GitHub authenticated"
    fi
    
    # Check for gh-sub-issue extension
    if ! gh extension list 2>/dev/null | grep -q "yahsan2/gh-sub-issue"; then
        if prompt_yes_no "Install gh-sub-issue extension (recommended for epic/task linking)?" "y"; then
            gh extension install yahsan2/gh-sub-issue
            print_success "gh-sub-issue extension installed"
        fi
    else
        print_success "gh-sub-issue extension found"
    fi
    
    # Create GitHub labels
    if gh repo view &> /dev/null; then
        if prompt_yes_no "Create GitHub labels (epic, task)?" "y"; then
            gh label create "epic" --color "0E8A16" --description "Epic issue containing multiple related tasks" --force 2>/dev/null && \
                print_success "Created 'epic' label" || print_warning "'epic' label may already exist"
            gh label create "task" --color "1D76DB" --description "Individual task within an epic" --force 2>/dev/null && \
                print_success "Created 'task' label" || print_warning "'task' label may already exist"
        fi
    else
        print_warning "Not a GitHub repository - skipping label creation"
    fi
}

# Verify installation
verify_installation() {
    print_step "Verifying installation..."
    
    local errors=0
    
    # Check directories
    [[ -d "ccpm" ]] || { print_error "Missing: ccpm/"; ((errors++)); }
    [[ -d ".claude/commands" ]] || { print_error "Missing: .claude/commands/"; ((errors++)); }
    [[ -d ".claude/prds" ]] || { print_error "Missing: .claude/prds/"; ((errors++)); }
    [[ -d ".claude/epics" ]] || { print_error "Missing: .claude/epics/"; ((errors++)); }
    [[ -d ".opencode/command" ]] || { print_error "Missing: .opencode/command/"; ((errors++)); }
    
    [[ -d ".claude/commands/pm" ]] || { print_error "Missing: .claude/commands/pm/"; ((errors++)); }
    [[ -d ".opencode/command/pm" ]] || { print_error "Missing: .opencode/command/pm/"; ((errors++)); }
    
    # Count commands
    local cmd_count=$(find ccpm/commands -name "*.md" | wc -l)
    
    if [[ $errors -eq 0 ]]; then
        print_success "Installation verified ($cmd_count commands available)"
        return 0
    else
        print_error "Installation has $errors error(s)"
        return 1
    fi
}

# Print final instructions
print_instructions() {
    echo ""
    echo "═══════════════════════════════════════════════════════════════"
    echo -e "${GREEN}✓ CCPM Installation Complete!${NC}"
    echo "═══════════════════════════════════════════════════════════════"
    echo ""
    echo -e "${YELLOW}⚠ IMPORTANT: Restart your Claude Code / OpenCode session${NC}"
    echo "  Commands are discovered at session start."
    echo ""
    echo "Available commands after restart:"
    echo "  /pm:help          - Show all PM commands"
    echo "  /pm:prd-new       - Create a new PRD"
    echo "  /pm:status        - Show project status"
    echo "  /cascade:start    - CASCADE FLOW (parallel brainstorming)"
    echo ""
    echo "Quick start:"
    echo "  1. Restart your session"
    echo "  2. Run: /pm:prd-new my-feature"
    echo ""
    echo "Documentation: https://github.com/automazeio/ccpm"
    echo ""
}

# Main installation flow
main() {
    print_header
    
    check_project_directory
    
    local is_update=false
    check_existing_installation && is_update=true
    
    echo -e "${BLUE}Installing CCPM...${NC}"
    echo ""
    
    download_ccpm
    create_directories
    setup_claude_commands
    setup_opencode_commands
    update_gitignore
    
    echo ""
    if prompt_yes_no "Setup GitHub integration?" "y"; then
        setup_github
    fi
    
    echo ""
    verify_installation
    
    print_instructions
}

# Run main
main "$@"
