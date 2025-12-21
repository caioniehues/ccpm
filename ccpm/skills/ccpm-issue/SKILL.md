---
name: ccpm-issue
description: Manages GitHub issue lifecycle for CCPM including creation, tracking, parallel work analysis, status updates, and bidirectional sync. Use when working with CCPM issues, task management, or GitHub issue operations.
---

<essential_principles>
## How CCPM Issue Management Works

### Principle 1: Local-First with GitHub Sync
CCPM maintains local task files as the source of truth, syncing bidirectionally with GitHub issues. All operations update local files first, then sync to GitHub.

### Principle 2: Frontmatter-Driven State
Task state (status, dates, progress) lives in YAML frontmatter. All operations must preserve frontmatter structure and follow datetime conventions.

### Principle 3: Parallel Work Coordination
Issue analysis identifies independent work streams. Multiple agents can work in parallel when properly coordinated through file scope and progress tracking.

### Principle 4: Epic Context
Issues exist within epics. Operations must maintain epic progress calculations and update parent epic state when tasks change.

### Principle 5: Transparent Audit Trail
All work syncs to GitHub as comments, creating a transparent history of progress, decisions, and completion for stakeholders.
</essential_principles>

<intake>
**What would you like to do with issue management?**

1. **Analyze** - Identify parallel work streams for maximum efficiency
2. **Start** - Begin work on an issue with parallel agents
3. **Status** - Check current state and progress
4. **Sync** - Push local updates to GitHub as comments
5. **Show** - Display full issue details and context
6. **Edit** - Update issue details locally and on GitHub
7. **Close** - Mark issue as complete and close
8. **Reopen** - Reopen a closed issue

**Wait for response before proceeding.**
</intake>

<routing>
| Response | Workflow | Purpose |
|----------|----------|---------|
| 1, "analyze", "analysis", "parallel", "streams" | `workflows/analyze.md` | Analyze issue for parallel work streams |
| 2, "start", "begin", "launch" | `workflows/start.md` | Start work with parallel agents |
| 3, "status", "check", "state" | `workflows/status.md` | Check issue status and progress |
| 4, "sync", "push", "update github" | `workflows/sync.md` | Sync local updates to GitHub |
| 5, "show", "display", "view", "details" | `workflows/show.md` | Display full issue information |
| 6, "edit", "modify", "change" | `workflows/edit.md` | Edit issue details |
| 7, "close", "complete", "finish" | `workflows/close.md` | Close completed issue |
| 8, "reopen", "resume" | `workflows/reopen.md` | Reopen closed issue |

**After reading the workflow, follow it exactly.**
</routing>

<reference_index>
## Shared Knowledge

All domain knowledge in `references/`:

**Core Patterns:**
- `frontmatter-operations.md` - YAML frontmatter structure and update patterns
- `datetime-handling.md` - Real datetime operations and formatting
- `github-sync.md` - GitHub CLI operations and sync patterns
- `epic-integration.md` - Epic progress calculation and updates

**Coordination:**
- `parallel-work.md` - Multi-agent coordination and file scope
- `progress-tracking.md` - Progress file structure and updates
</reference_index>

<workflows_index>
## Available Workflows

| Workflow | Purpose | Typical Use |
|----------|---------|-------------|
| analyze.md | Analyze issue for parallel work streams | Before starting complex work |
| start.md | Begin work with parallel agents | Start implementation |
| status.md | Check issue status and progress | Quick status check |
| sync.md | Sync local updates to GitHub | Share progress updates |
| show.md | Display full issue details | View complete context |
| edit.md | Edit issue details | Update requirements |
| close.md | Close completed issue | Mark work complete |
| reopen.md | Reopen closed issue | Resume work |
</workflows_index>

<success_criteria>
Issue management is working correctly when:

- All operations update local files before GitHub
- Frontmatter remains valid after every operation
- Epic progress recalculates when tasks change
- GitHub sync creates transparent audit trail
- Parallel work coordination prevents conflicts
- Real datetimes (not placeholders) in all timestamps
- File paths and references remain valid
</success_criteria>
