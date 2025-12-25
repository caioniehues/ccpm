# Command Reference

**[中文文档 (Chinese Documentation)](../zh-docs/COMMANDS_ZH.md)**

This is the comprehensive reference for all commands available in Claude Code PM (CCPM).

## Command Structure

CCPM commands use a slash syntax followed by a category and command name:

```
/category:command [arguments]
```

- **Category**: Groups related commands (e.g., `pm`, `context`, `testing`)
- **Command**: The specific action to perform
- **Arguments**: Optional parameters (e.g., feature name, issue number)

Some system-wide commands (like `/doctor`) do not have a category prefix.

## Project Management Commands (`/pm:*`)

These commands form the core of the CCPM workflow, handling everything from requirements to issue tracking.

### PRD Management

Commands for creating and managing Product Requirements Documents.

| Command | Description | Usage |
|---------|-------------|-------|
| `/pm:prd-new` | Launch brainstorming session to create a comprehensive PRD | `/pm:prd-new <feature-name>` |
| `/pm:prd-edit` | Edit an existing PRD with user-specified changes | `/pm:prd-edit <feature-name>` |
| `/pm:prd-parse` | Convert a PRD into a technical implementation epic | `/pm:prd-parse <feature-name>` |
| `/pm:prd-list` | List all PRDs with their current status | `/pm:prd-list` |
| `/pm:prd-status` | Show a summary report of PRDs grouped by status | `/pm:prd-status` |

### Epic Management

Commands for managing technical epics and their tasks.

| Command | Description | Usage |
|---------|-------------|-------|
| `/pm:epic-wizard` | Guided workflow for creating an epic from start to finish | `/pm:epic-wizard <feature-name>` |
| `/pm:epic-decompose` | Break an epic into concrete, actionable tasks | `/pm:epic-decompose <feature-name>` |
| `/pm:epic-sync` | Push epic and tasks to GitHub as issues | `/pm:epic-sync <feature-name>` |
| `/pm:epic-oneshot` | Decompose and sync in a single operation | `/pm:epic-oneshot <feature-name>` |
| `/pm:epic-start` | Launch parallel agents to work on epic tasks | `/pm:epic-start <feature-name>` |
| `/pm:epic-start-worktree` | Launch agents in a dedicated git worktree | `/pm:epic-start-worktree <feature-name>` |
| `/pm:epic-show` | Display detailed information about an epic | `/pm:epic-show <feature-name>` |
| `/pm:epic-list` | List all epics with their progress status | `/pm:epic-list` |
| `/pm:epic-edit` | Edit epic details interactively | `/pm:epic-edit <feature-name>` |
| `/pm:epic-close` | Mark an epic as complete (verifies all tasks done) | `/pm:epic-close <feature-name>` |
| `/pm:epic-merge` | Merge a completed epic worktree back to main | `/pm:epic-merge <feature-name>` |
| `/pm:epic-refresh` | Update epic progress based on current task states | `/pm:epic-refresh <feature-name>` |

### Issue Management

Commands for working with individual GitHub issues and tasks.

| Command | Description | Usage |
|---------|-------------|-------|
| `/pm:issue-start` | Begin work on an issue with parallel agents | `/pm:issue-start <issue-number>` |
| `/pm:issue-analyze` | Analyze an issue to identify parallel work streams | `/pm:issue-analyze <issue-number>` |
| `/pm:issue-sync` | Push local progress to GitHub as comments | `/pm:issue-sync <issue-number>` |
| `/pm:issue-show` | Display issue details and sub-issues | `/pm:issue-show <issue-number>` |
| `/pm:issue-status` | Check current status of an issue | `/pm:issue-status <issue-number>` |
| `/pm:issue-edit` | Edit issue details locally and on GitHub | `/pm:issue-edit <issue-number>` |
| `/pm:issue-close` | Mark an issue as complete and close on GitHub | `/pm:issue-close <issue-number>` |
| `/pm:issue-reopen` | Reopen a closed issue | `/pm:issue-reopen <issue-number>` |

### Workflow & Tracking

Commands for daily workflow and status tracking.

| Command | Description | Usage |
|---------|-------------|-------|
| `/pm:dashboard` | Launch TUI dashboard for visual tracking | `/pm:dashboard` |
| `/pm:next` | List next prioritized tasks ready for work | `/pm:next` |
| `/pm:status` | Show overall project status summary | `/pm:status` |
| `/pm:standup` | Generate a daily standup report | `/pm:standup` |
| `/pm:blocked` | List tasks currently blocked by dependencies | `/pm:blocked` |
| `/pm:in-progress` | List all tasks currently being worked on | `/pm:in-progress` |
| `/pm:search` | Search across PRDs, epics, and tasks | `/pm:search <term>` |

### System & Maintenance

Commands for maintaining the PM system itself.

| Command | Description | Usage |
|---------|-------------|-------|
| `/pm:init` | Initialize CCPM structure and configuration | `/pm:init` |
| `/pm:import` | Import existing GitHub issues into CCPM | `/pm:import` |
| `/pm:sync` | Bidirectional sync between local files and GitHub | `/pm:sync` |
| `/pm:validate` | Check system integrity and configuration | `/pm:validate` |
| `/pm:clean` | Archive completed work and remove stale files | `/pm:clean` |
| `/pm:help` | Show summary of available PM commands | `/pm:help` |

## Context Commands (`/context:*`)

Commands for managing project documentation and Claude's context window.

| Command | Description | Usage |
|---------|-------------|-------|
| `/context:create` | Analyze project and create initial documentation | `/context:create` |
| `/context:update` | Refresh context docs with recent changes | `/context:update` |
| `/context:prime` | Load project context into current conversation | `/context:prime` |

## Testing Commands (`/testing:*`)

Commands for configuring and running tests.

| Command | Description | Usage |
|---------|-------------|-------|
| `/testing:prime` | Detect and configure project testing framework | `/testing:prime` |
| `/testing:run` | Execute tests and analyze failures | `/testing:run [target]` |

## System Utility Commands

General commands for system health and maintenance.

| Command | Description | Usage |
|---------|-------------|-------|
| `/doctor` | Diagnose installation and configuration issues | `/doctor` |
| `/setup` | Interactive setup wizard for first-time users | `/setup` |
| `/self-update` | Update CCPM to the latest version | `/self-update` |
| `/uninstall` | Remove CCPM from the project | `/uninstall` |
| `/version` | Display version information | `/version` |
| `/re-init` | Update CLAUDE.md with latest PM rules | `/re-init` |
| `/prompt` | Execute complex prompts from a file | `/prompt` |
| `/code-rabbit` | Process CodeRabbit review comments | `/code-rabbit` |

## Command Patterns

All CCPM commands follow consistent architectural patterns.

### Frontmatter Configuration
Each command file defines its capabilities in the YAML frontmatter:

```yaml
---
description: Brief description of what the command does
argument-hint: [optional args] <required args>
allowed-tools: Bash, Read, Write, Glob, Task
---
```

### Tool Permissions
Commands are explicit about the tools they need:
- **Bash**: For system operations and scripts
- **Read/Write/LS/Glob**: For file manipulation
- **Task**: For spawning sub-agents (critical for parallel work)

### Agent Integration
Many commands spawn specialized agents to preserve the main context:
- `epic-decompose` → Uses **task-decomposer** agent
- `epic-sync` → Uses **github-syncer** agent
- `issue-start` → Uses **parallel-worker** agent

## Creating Custom Commands

You can extend CCPM by adding your own commands.

1.  **Create the file**: Add a markdown file in `ccpm/commands/<category>/<command>.md`.
2.  **Define frontmatter**:
    ```yaml
    ---
    description: My custom command
    allowed-tools: Bash, Read
    ---
    ```
3.  **Write instructions**: Use XML tags for structure (optional but recommended).
    ```xml
    <objective>
    Explain what the command should achieve.
    </objective>

    <steps>
    1. Step one
    2. Step two
    </steps>
    ```

The system will automatically pick up the new command (no restart required).
