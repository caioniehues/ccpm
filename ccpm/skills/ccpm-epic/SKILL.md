---
name: ccpm-epic
description: Manages epic lifecycle for CCPM including planning (decompose), execution (start, status), tracking (show, refresh), and completion (close, merge). Coordinates epic creation from PRDs, task decomposition, GitHub synchronization, and parallel agent execution. Use when working with epics, PRDs, or coordinating multi-task development workflows.
---

<objective>
Manage epic lifecycle from planning through completion including PRD parsing, task decomposition, GitHub synchronization, parallel agent execution, and branch management.
</objective>

<essential_principles>
<epic_lifecycle>
Epics in CCPM follow a structured lifecycle:
1. **Create**: Parse PRD into epic structure
2. **Decompose**: Break epic into actionable tasks
3. **Sync**: Push epic and tasks to GitHub as issues
4. **Start**: Launch parallel agents to work on tasks
5. **Track**: Monitor progress and update status
6. **Complete**: Close epic and merge changes
</epic_lifecycle>

<key_concepts>
**Epic Directory Structure**:
```
.claude/epics/{epic-name}/
├── epic.md              # Epic overview with frontmatter
├── {issue-id}.md        # Task files (named by GitHub issue number)
├── github-mapping.md    # Mapping of tasks to GitHub issues
└── execution-status.md  # Active agent tracking
```

**Task Dependencies**:
- Tasks have `depends_on` field listing prerequisite tasks
- Tasks with `parallel: true` can run simultaneously
- Tasks with `conflicts_with` field modify same files

**GitHub Integration**:
- Epic becomes parent GitHub issue with "epic" label
- Tasks become sub-issues with "task" label
- Issue numbers replace task numbers (001.md → 1234.md)
- All issues tagged with `epic:{epic-name}` label

**Branch Strategy**:
- Each epic gets its own branch: `epic/{epic-name}`
- All agents work in the same epic branch (not separate branches)
- Branch created during sync, used during start
</key_concepts>
</essential_principles>

<intake>
What would you like to do with epics?

1. **Create** - Create a new epic from a PRD
2. **Decompose** - Break epic into actionable tasks
3. **Sync** - Push epic and tasks to GitHub as issues
4. **Oneshot** - Decompose and sync in one operation
5. **Start** - Launch parallel agents to work on epic tasks
6. **List** - List all epics with their status
7. **Show** - Show detailed information about a specific epic
8. **Status** - Show task status breakdown for an epic
9. **Edit** - Edit epic details after creation
10. **Refresh** - Update epic progress based on task states
11. **Close** - Mark epic as complete when all tasks done
12. **Merge** - Merge epic branch to main
13. **Start Worktree** - Create worktree for epic development

Provide the operation name or number, optionally with epic name.

**Wait for response before proceeding.**
</intake>

<routing>
| Response | Operation | Workflow |
|----------|-----------|----------|
| 1, "create" | Create epic from PRD | Delegate to `/pm:prd-parse` command |
| 2, "decompose", "break down", "tasks" | Decompose epic | Delegate to `/pm:epic-decompose` command |
| 3, "sync", "github", "push" | Sync to GitHub | Delegate to `/pm:epic-sync` command |
| 4, "oneshot", "one shot", "quick" | Decompose + Sync | Delegate to `/pm:epic-oneshot` command |
| 5, "start", "launch", "begin", "execute" | Start execution | Delegate to `/pm:epic-start` command |
| 6, "list", "show all", "ls" | List epics | Delegate to `/pm:epic-list` command |
| 7, "show", "view", "details", "info" | Show epic details | Delegate to `/pm:epic-show` command |
| 8, "status", "progress", "state" | Show task status | Delegate to `/pm:epic-status` command |
| 9, "edit", "modify", "update" | Edit epic | Delegate to `/pm:epic-edit` command |
| 10, "refresh", "recalculate", "update progress" | Refresh progress | Delegate to `/pm:epic-refresh` command |
| 11, "close", "complete", "finish" | Close epic | Delegate to `/pm:epic-close` command |
| 12, "merge" | Merge epic branch | Delegate to `/pm:epic-merge` command |
| 13, "worktree", "start worktree" | Create worktree | Delegate to `/pm:epic-start-worktree` command |

**IMPORTANT**: All epic operations are implemented as slash commands in `ccpm/commands/pm/epic-*.md`. This skill routes to the appropriate command based on user intent.

**After determining the operation, delegate to the corresponding slash command.**
</routing>

<command_reference>
All commands are in `ccpm/commands/pm/`:

<command_group name="planning">
**Planning & Setup**:
- `/pm:prd-parse {epic_name}` - Create epic from PRD
- `/pm:epic-decompose {epic_name}` - Break into tasks
- `/pm:epic-sync {epic_name}` - Push to GitHub
- `/pm:epic-oneshot {epic_name}` - Decompose + sync together
</command_group>

<command_group name="execution">
**Execution**:
- `/pm:epic-start {epic_name}` - Launch parallel agents
- `/pm:epic-start-worktree {epic_name}` - Create worktree
</command_group>

<command_group name="tracking">
**Tracking & Updates**:
- `/pm:epic-list` - List all epics
- `/pm:epic-show {epic_name}` - Show epic details
- `/pm:epic-status {epic_name}` - Show task breakdown
- `/pm:epic-refresh {epic_name}` - Update progress
</command_group>

<command_group name="maintenance">
**Maintenance**:
- `/pm:epic-edit {epic_name}` - Edit epic details
- `/pm:epic-close {epic_name}` - Mark as complete
- `/pm:epic-merge {epic_name}` - Merge to main
</command_group>
</command_reference>

<workflow_patterns>
<workflow_pattern name="full_lifecycle">
**New Epic (Full Lifecycle)**:
```
1. /pm:prd-parse {epic_name}      # Create epic from PRD
2. /pm:epic-decompose {epic_name}  # Break into tasks
3. /pm:epic-sync {epic_name}       # Push to GitHub
4. /pm:epic-start {epic_name}      # Launch agents
5. /pm:epic-status {epic_name}     # Monitor progress
6. /pm:epic-refresh {epic_name}    # Update progress
7. /pm:epic-close {epic_name}      # Mark complete
8. /pm:epic-merge {epic_name}      # Merge to main
```
</workflow_pattern>

<workflow_pattern name="quick_start">
**Quick Start (Oneshot)**:
```
1. /pm:prd-parse {epic_name}       # Create epic from PRD
2. /pm:epic-oneshot {epic_name}    # Decompose + sync together
3. /pm:epic-start {epic_name}      # Launch agents
```
</workflow_pattern>

<workflow_pattern name="monitoring">
**Monitoring Existing Epic**:
```
/pm:epic-list                      # See all epics
/pm:epic-show {epic_name}          # View specific epic
/pm:epic-status {epic_name}        # Check task breakdown
/pm:epic-refresh {epic_name}       # Update from GitHub
```
</workflow_pattern>

<workflow_pattern name="manual_work">
**Manual Work (No Agents)**:
```
1. /pm:epic-oneshot {epic_name}    # Decompose + sync
2. /pm:epic-start-worktree {epic_name}  # Create worktree
3. # Work manually in worktree
4. /pm:epic-refresh {epic_name}    # Update progress
5. /pm:epic-close {epic_name}      # When done
```
</workflow_pattern>
</workflow_patterns>

<quick_start>
**Quick start:**

```bash
# Create epic from PRD
/pm:prd-parse my-feature

# Decompose + sync to GitHub in one step
/pm:epic-oneshot my-feature

# Launch parallel agents to work on epic
/pm:epic-start my-feature

# Check epic progress
/pm:epic-status my-feature
```
</quick_start>

<success_criteria>
This skill successfully routes requests when:
- User intent is correctly identified from their request
- Appropriate epic command is invoked with correct arguments
- Epic name is extracted if provided in user request
- User understands which operation will be performed

The underlying epic commands handle actual execution and validation.
</success_criteria>
