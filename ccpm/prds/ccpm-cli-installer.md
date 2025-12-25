---
name: ccpm-cli-installer
description: One-command CCPM installation with automatic symlink setup for Claude Code and OpenCode compatibility
status: implemented
created: 2025-12-25T21:35:54Z
updated: 2025-12-25T21:35:54Z
---

# PRD: CCPM CLI Installer

## Executive Summary

### Problem Statement
CCPM commands were not discoverable by Claude Code or OpenCode because the commands exist in `ccpm/commands/` but both tools look for commands in platform-specific directories (`.claude/commands/` and `.opencode/command/`). Users had to manually create directories and symlinks, a process that was error-prone and undocumented.

### Proposed Solution
A single-command installer (`curl | bash`) that:
1. Downloads CCPM files from GitHub
2. Creates symlinks to make commands discoverable
3. Sets up required directories
4. Optionally configures GitHub integration
5. Verifies the installation

### Success Criteria
- `/pm:*` commands work immediately after running the installer
- Both Claude Code (`.claude/commands/`) and OpenCode (`.opencode/command/`) are supported
- Zero manual steps required for basic functionality
- GitHub integration is optional but guided

## Context and Background

### Current State
Before this feature:
- Users downloaded CCPM manually
- Commands in `ccpm/commands/` were invisible to Claude Code
- No OpenCode support existed
- Installation instructions were incomplete
- Many users reported "command not found" errors

### Target Users
- Developers adopting CCPM for the first time
- Teams setting up CCPM in new projects
- OpenCode users (previously unsupported)

### Business Objectives
- Reduce installation friction to near-zero
- Eliminate "command not found" support issues
- Enable OpenCode users to use CCPM
- Establish a single, reliable installation path

## User Stories

### Primary User Story
**As a** developer new to CCPM
**I want to** run a single command to install CCPM
**So that** I can immediately use `/pm:*` commands without manual configuration

**Acceptance Criteria:**
- [x] Single curl command downloads and runs installer
- [x] All `/pm:*` commands available after installation
- [x] No manual symlink creation required
- [x] Works on Linux, macOS (Unix), and Windows

### Additional User Stories

**US-2: OpenCode Support**
**As an** OpenCode user
**I want to** have CCPM commands in `.opencode/command/`
**So that** I can use CCPM with my preferred tool

**Acceptance Criteria:**
- [x] Installer creates `.opencode/command/` directory
- [x] Symlinks point to `ccpm/commands/`
- [x] Commands work in OpenCode sessions

**US-3: Optional GitHub Setup**
**As a** developer using CCPM with GitHub
**I want** guided GitHub configuration during install
**So that** I can sync issues without additional setup

**Acceptance Criteria:**
- [x] Installer prompts for GitHub setup (optional)
- [x] gh CLI authentication is handled
- [x] gh-sub-issue extension installed if chosen
- [x] Required labels created

## Requirements

### Functional Requirements
FR-1: Single-command installation via `curl -sSL URL | bash`
FR-2: Download CCPM files from GitHub (latest release or main branch)
FR-3: Create `.claude/commands/` with file-level symlinks to `ccpm/commands/` files
FR-4: Create `.opencode/command/` with file-level symlinks to `ccpm/commands/` files
FR-5: Create `.claude/prds/` and `.claude/epics/` directories
FR-6: Update `.gitignore` to exclude local working files
FR-7: Interactive GitHub setup (optional): auth, labels, gh-sub-issue extension
FR-8: Verification step confirming successful installation
FR-9: Idempotent execution (safe to run multiple times)

### Non-Functional Requirements
NFR-1: Installation completes in under 60 seconds on typical connection
NFR-2: Works offline after initial download (except GitHub setup)
NFR-3: No root/sudo required for basic installation
NFR-4: Clear error messages for common failure modes

### Constraints
- Must use file-level symlinks (directory symlinks don't work with command discovery)
- Must support both bash and PowerShell
- Cannot require Node.js, Python, or other runtimes
- GitHub setup must be fully optional

### Out of Scope
- Package manager distribution (npm, brew, apt)
- Auto-update mechanism
- GUI installer
- Docker installation method

## Dependencies

### Internal Dependencies
- `ccpm/commands/` structure must remain stable
- Command files must be valid markdown with frontmatter

### External Dependencies
- curl or wget for download
- git for version control
- gh CLI (optional, for GitHub integration)

### Prerequisites
- Unix-like environment (bash) or Windows (PowerShell)
- Write access to project directory
- Internet connection for initial download

## Success Metrics

### Acceptance Criteria
- [x] `curl -sSL https://automaze.io/ccpm/install | bash` works
- [x] After install, `/pm:help` shows all commands
- [x] Symlinks correctly resolve (file content accessible)
- [x] .gitignore updated with correct entries
- [x] Installer is idempotent (running twice doesn't break anything)

### KPIs
- Installation success rate > 95%
- Average installation time < 30 seconds
- Zero "command not found" issues from correctly installed users

## Implementation Notes

### Key Design Decision: File-Level Symlinks
Directory-level symlinks (e.g., `.claude/commands/pm -> ccpm/commands/pm`) are NOT followed by Claude Code's command discovery system. We discovered this through debugging and must use file-level symlinks instead:

```bash
# WRONG - directory symlink (not followed by command discovery)
ln -sf "../../ccpm/commands/pm" ".claude/commands/pm"

# CORRECT - file-level symlinks (works with command discovery)
mkdir -p .claude/commands/pm
for cmd in ccpm/commands/pm/*.md; do
    ln -sf "../../../$cmd" ".claude/commands/pm/$(basename "$cmd")"
done
```

### Files Created/Modified
- `install/ccpm.sh` - Main installer script (382 lines)
- `install/ccpm.bat` - Windows PowerShell wrapper
- `.claude/commands/{pm,context,testing}/` - Directories with file symlinks
- `.opencode/command/{pm,context,testing}/` - Directories with file symlinks
- `.gitignore` - Updated with CCPM entries
