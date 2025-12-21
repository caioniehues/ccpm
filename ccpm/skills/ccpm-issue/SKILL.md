---
name: ccpm-issue
description: Manages GitHub issue lifecycle for CCPM including creation, tracking, parallel work analysis, status updates, and bidirectional sync. Use when working with CCPM issues, task management, or GitHub issue operations.
---

<objective>
Manage issue lifecycle with local-first sync to GitHub including parallel work analysis, status tracking, and bidirectional synchronization.
</objective>

<essential_principles>
<principle name="local_first">
**Local-First with GitHub Sync**
CCPM maintains local task files as the source of truth, syncing bidirectionally with GitHub issues. All operations update local files first, then sync to GitHub.
</principle>

<principle name="frontmatter_driven">
**Frontmatter-Driven State**
Task state (status, dates, progress) lives in YAML frontmatter. All operations must preserve frontmatter structure and follow datetime conventions.
</principle>

<principle name="parallel_coordination">
**Parallel Work Coordination**
Issue analysis identifies independent work streams. Multiple agents can work in parallel when properly coordinated through file scope and progress tracking.
</principle>

<principle name="epic_context">
**Epic Context**
Issues exist within epics. Operations must maintain epic progress calculations and update parent epic state when tasks change.
</principle>

<principle name="audit_trail">
**Transparent Audit Trail**
All work syncs to GitHub as comments, creating a transparent history of progress, decisions, and completion for stakeholders.
</principle>
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
| 1, "analyze", "analysis", "parallel", "streams" | `workflows/analyze-issue.md` | Analyze issue for parallel work streams |
| 2, "start", "begin", "launch" | `workflows/start-issue.md` | Start work with parallel agents |
| 3, "status", "check", "state" | `workflows/status-issue.md` | Check issue status and progress |
| 4, "sync", "push", "update github" | `workflows/sync-issue.md` | Sync local updates to GitHub |
| 5, "show", "display", "view", "details" | `workflows/show.md` | Display full issue information |
| 6, "edit", "modify", "change" | `workflows/edit-issue.md` | Edit issue details |
| 7, "close", "complete", "finish" | `workflows/close-issue.md` | Close completed issue |
| 8, "reopen", "resume" | `workflows/reopen-issue.md` | Reopen closed issue |

**After reading the workflow, follow it exactly.**
</routing>

<reference_index>
All domain knowledge in `references/`:

<reference_group name="core_patterns">
**Core Patterns:**
- `frontmatter-operations.md` - YAML frontmatter structure and update patterns
- `datetime-handling.md` - Real datetime operations and formatting
- `github-sync.md` - GitHub CLI operations and sync patterns
- `epic-integration.md` - Epic progress calculation and updates
</reference_group>

<reference_group name="coordination">
**Coordination:**
- `parallel-work.md` - Multi-agent coordination and file scope
- `progress-tracking.md` - Progress file structure and updates
</reference_group>
</reference_index>

<workflows_index>
| Workflow | Purpose | Typical Use |
|----------|---------|-------------|
| analyze-issue.md | Analyze issue for parallel work streams | Before starting complex work |
| start-issue.md | Begin work with parallel agents | Start implementation |
| status-issue.md | Check issue status and progress | Quick status check |
| sync-issue.md | Sync local updates to GitHub | Share progress updates |
| show.md | Display full issue details | View complete context |
| edit-issue.md | Edit issue details | Update requirements |
| close-issue.md | Close completed issue | Mark work complete |
| reopen-issue.md | Reopen closed issue | Resume work |
</workflows_index>

<quick_start>
**Quick start:**

```bash
# Start working on an issue
/pm:issue-start 123

# Analyze issue for parallel work streams
/pm:issue-analyze 123

# Check issue status
/pm:issue-status 123

# Sync local updates to GitHub
/pm:issue-sync 123
```
</quick_start>

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
