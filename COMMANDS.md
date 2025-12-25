# Commands

**[中文文档 (Chinese Documentation)](doc/COMMANDS_ZH.md)**

Complete reference of all commands available in the Claude Code PM system.

> **Note**: Project Management commands (`/pm:*`) are documented in the main [README.md](README.md#command-reference).

## Table of Contents

- [Context Commands](#context-commands)
- [Testing Commands](#testing-commands)
- [Utility Commands](#utility-commands)
- [System Commands](#system-commands)
- [Review Commands](#review-commands)

## Context Commands

Commands for managing project context in `ccpm/context/`.

### `/context:create`
- **Purpose**: Create initial project context documentation
- **Usage**: `/context:create`
- **Description**: Analyzes the project structure and creates comprehensive baseline documentation in `.claude/context/`. Includes project overview, architecture, dependencies, and patterns.
- **When to use**: At project start or when context needs full rebuild
- **Output**: Multiple context files covering different aspects of the project

### `/context:update`
- **Purpose**: Update existing context with recent changes
- **Usage**: `/context:update`
- **Description**: Refreshes context documentation based on recent code changes, new features, or architectural updates. Preserves existing context while adding new information.
- **When to use**: After significant changes or before major work sessions
- **Output**: Updated context files with change tracking

### `/context:prime`
- **Purpose**: Load context into current conversation
- **Usage**: `/context:prime`
- **Description**: Reads all context files and loads them into the current conversation's memory. Essential for maintaining project awareness.
- **When to use**: At the start of any work session
- **Output**: Confirmation of loaded context

## Testing Commands

Commands for test configuration and execution.

### `/testing:prime`
- **Purpose**: Configure testing setup
- **Usage**: `/testing:prime`
- **Description**: Detects and configures the project's testing framework, creates testing configuration, and prepares the test-runner agent.
- **When to use**: Initial project setup or when testing framework changes
- **Output**: `ccpm/testing-config.md` with test commands and patterns

### `/testing:run`
- **Purpose**: Execute tests with intelligent analysis
- **Usage**: `/testing:run [test_target]`
- **Description**: Runs tests using the test-runner agent which captures output to logs and returns only essential results to preserve context.
- **Options**:
   - No arguments: Run all tests
   - File path: Run specific test file
   - Pattern: Run tests matching pattern
- **Output**: Test summary with failures analyzed, no verbose output in main thread

## Utility Commands

General utility and maintenance commands.

### `/prompt`
- **Purpose**: Handle complex prompts with multiple references
- **Usage**: Write your prompt in the file, then type `/prompt`
- **Description**: Ephemeral command for when complex prompts with numerous @ references fail in direct input. The prompt is written to the command file first, then executed.
- **When to use**: When Claude's UI rejects complex prompts
- **Output**: Executes the written prompt

### `/re-init`
- **Purpose**: Update or create CLAUDE.md with PM rules
- **Usage**: `/re-init`
- **Description**: Updates the project's CLAUDE.md file with rules from `.claude/CLAUDE.md`, ensuring Claude instances have proper instructions.
- **When to use**: After cloning PM system or updating rules
- **Output**: Updated CLAUDE.md in project root

## System Commands

Commands for system maintenance and diagnostics.

### `/doctor`
- **Purpose**: Diagnose CCPM installation issues
- **Usage**: `/doctor`
- **Description**: Runs comprehensive diagnostics on your CCPM installation, checking for missing files, incorrect configurations, and common setup issues.
- **When to use**: After installation or when experiencing issues
- **Output**: Diagnostic report with recommendations

### `/setup`
- **Purpose**: Initial CCPM setup wizard
- **Usage**: `/setup`
- **Description**: Interactive setup wizard that guides you through initial CCPM configuration, including GitHub authentication and directory structure creation.
- **When to use**: First-time setup or reconfiguration
- **Output**: Configured CCPM environment

### `/self-update`
- **Purpose**: Update CCPM to latest version
- **Usage**: `/self-update`
- **Description**: Checks for updates and applies the latest CCPM version from the repository.
- **When to use**: When updates are available or periodically
- **Output**: Updated CCPM installation

### `/uninstall`
- **Purpose**: Remove CCPM from project
- **Usage**: `/uninstall`
- **Description**: Cleanly removes CCPM configuration and files while preserving your PRDs and epics.
- **When to use**: When removing CCPM from a project
- **Output**: Clean uninstallation with backup option

### `/version`
- **Purpose**: Display CCPM version information
- **Usage**: `/version`
- **Description**: Shows the current CCPM version, installation date, and configuration status.
- **When to use**: When checking version or reporting issues
- **Output**: Version and configuration information

## Review Commands

Commands for handling external code review tools.

### `/code-rabbit`
- **Purpose**: Process CodeRabbit review comments intelligently
- **Usage**: `/code-rabbit` then paste comments
- **Description**: Evaluates CodeRabbit suggestions with context awareness, accepting valid improvements while ignoring context-unaware suggestions. Spawns parallel agents for multi-file reviews.
- **Features**:
   - Understands CodeRabbit lacks full context
   - Accepts: Real bugs, security issues, resource leaks
   - Ignores: Style preferences, irrelevant patterns
   - Parallel processing for multiple files
- **Output**: Summary of accepted/ignored suggestions with reasoning

## Command Patterns

All commands follow consistent patterns:

### Allowed Tools
Each command specifies its required tools in frontmatter:
- `Read, Write, LS` - File operations
- `Bash` - System commands
- `Task` - Sub-agent spawning
- `Grep` - Code searching

### Error Handling
Commands follow fail-fast principles:
- Check prerequisites first
- Clear error messages with solutions
- Never leave partial state

### Context Preservation
Commands that process lots of information:
- Use agents to shield main thread from verbose output
- Return summaries, not raw data
- Preserve only essential information

## Creating Custom Commands

To add new commands:

1. **Create file**: `ccpm/commands/category/command-name.md`
2. **Add frontmatter**:
   ```yaml
   ---
   allowed-tools: Read, Write, LS
   ---
   ```
3. **Structure content**:
   - Purpose and usage
   - Preflight checks
   - Step-by-step instructions
   - Error handling
   - Output format

4. **Follow patterns**:
   - Keep it simple (no over-validation)
   - Fail fast with clear messages
   - Use agents for heavy processing
   - Return concise output

## Integration with Agents

Commands often use agents for heavy lifting:

- **test-runner**: Executes tests, analyzes results
- **file-analyzer**: Summarizes verbose files
- **code-analyzer**: Hunts bugs across codebase
- **parallel-worker**: Coordinates parallel execution
- **prd-architect**: Creates structured PRD documents
- **epic-planner**: Plans epic execution with dependency mapping
- **task-decomposer**: Breaks epics into ordered tasks
- **github-syncer**: Syncs CCPM files with GitHub issues
- **worktree-manager**: Manages git worktrees for parallel work
- **parallel-orchestrator**: Coordinates parallel task execution

This keeps the main conversation context clean while doing complex work.

## Skills and Workflow Routing

CCPM uses a three-layer architecture to avoid circular references and ensure clean execution:

```
┌─────────────┐     ┌─────────────┐
│  Commands   │     │   Skills    │
│  /pm:*      │     │  ccpm-*     │
└──────┬──────┘     └──────┬──────┘
       │                   │
       └───────┬───────────┘
               │
               ▼
       ┌───────────────┐
       │   Workflows   │
       │  (the actual  │
       │  implementation)│
       └───────────────┘
```

### How It Works

1. **Commands** (e.g., `/pm:prd-new`) provide the user interface
2. **Skills** (e.g., `ccpm-prd`) organize related workflows and provide context
3. **Workflows** (e.g., `new-prd.md`) contain the actual implementation steps

### Why This Matters

- **No circular references**: Commands don't invoke skills that invoke commands
- **Single source of truth**: Implementation logic lives in one place (workflows)
- **Reusability**: Both commands and skills can reference the same workflow
- **Maintainability**: Update a workflow once, all callers get the update

### Available Skills

| Skill | Purpose | Workflows |
|-------|---------|-----------|
| `ccpm-prd` | PRD lifecycle management | new-prd, edit-prd, parse-prd |
| `ccpm-epic` | Epic lifecycle management | decompose, sync, start, close, merge, edit, oneshot, refresh |
| `ccpm-issue` | Issue lifecycle management | analyze, start, status, sync, edit, close, reopen |
| `ccpm-context` | Project context management | create-context, update-context, prime-context |
| `ccpm-testing` | Test execution | prime-testing, run-tests |
| `ccpm-worktree` | Git worktree management | (worktree operations) |

Skills are located in `ccpm/skills/` and documented in each skill's `SKILL.md` file.

## Notes

- Commands are markdown files interpreted as instructions
- The `/` prefix triggers command execution
- Commands can spawn agents for context preservation
- All PM commands (`/pm:*`) are documented in the main README
- Commands follow rules defined in `ccpm/rules/`
