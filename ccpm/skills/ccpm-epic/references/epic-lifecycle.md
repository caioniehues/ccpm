# Epic Lifecycle Reference

Complete documentation for managing epics from creation to completion.

## Epic States

```
┌─────────────────────────────────────────────────────────────────────────┐
│                          EPIC LIFECYCLE                                  │
├─────────────────────────────────────────────────────────────────────────┤
│                                                                          │
│  Created ──► Decomposed ──► Synced ──► Started ──► Merged ──► Closed    │
│     │            │            │           │           │          │       │
│     ▼            ▼            ▼           ▼           ▼          ▼       │
│   PRD        Task Files    GitHub     Worktree     Main       Archived  │
│  Parsed       Created      Issues      Active      Branch      Data     │
│                                                                          │
│                            ┌──────────────┐                              │
│                            │  Refreshed   │◄──── Progress updates        │
│                            └──────────────┘                              │
└─────────────────────────────────────────────────────────────────────────┘
```

### State Descriptions

| State | Description | Files Present |
|-------|-------------|---------------|
| Created | PRD parsed, epic.md exists | epic.md |
| Decomposed | Tasks broken down | epic.md, 001.md, 002.md... |
| Synced | Pushed to GitHub | epic.md, {issue_id}.md, github-mapping.md |
| Started | Agents working | execution-status.md, worktree |
| Merged | Code in main branch | archived/ |
| Closed | Epic complete | .archived/{epic}/ |

## File Structure

```
.claude/epics/{epic_name}/
├── epic.md                    # Epic definition
├── {issue_number}.md          # Task files (renamed from 001.md after sync)
├── {issue_number}-analysis.md # Work stream analysis (optional)
├── github-mapping.md          # Issue ID mappings
├── execution-status.md        # Active agent tracking
└── updates/
    └── {issue_number}/
        └── stream-{X}.md      # Agent progress updates
```

## Epic File Format

```markdown
---
name: Feature Name
status: backlog|in-progress|completed
created: 2024-01-15T10:30:00Z
updated: 2024-01-15T10:30:00Z
prd: ../prds/feature-name.md
github: https://github.com/owner/repo/issues/123
progress: 0%
completed: null
---

# Epic: Feature Name

## Overview
High-level description of the epic.

## Technical Approach
Architecture decisions and implementation strategy.

## Success Criteria
- Criterion 1
- Criterion 2

## Tasks Created
- [ ] #123 - Task Name (parallel: true)
- [ ] #124 - Task Name (parallel: false)

Total tasks: 5
Parallel tasks: 3
Sequential tasks: 2
```

## Task File Format

```markdown
---
name: Task Title
status: open|in_progress|closed
created: 2024-01-15T10:30:00Z
updated: 2024-01-15T10:30:00Z
github: https://github.com/owner/repo/issues/124
depends_on: []
parallel: true
conflicts_with: []
effort: M
---

# Task: Title

## Description
What needs to be done.

## Acceptance Criteria
- [ ] Criterion 1
- [ ] Criterion 2

## Technical Details
Implementation notes.

## Definition of Done
- [ ] Code implemented
- [ ] Tests passing
```

## Workflow Sequence

### 1. Epic Creation
```
User Request → PRD Parsing → epic.md Created
```
- PRD content extracted and structured
- Initial frontmatter set (status: backlog, progress: 0%)
- No tasks yet

### 2. Decomposition
```
epic.md → Task Analysis → 001.md, 002.md, ...
```
- Tasks identified from requirements
- Dependencies mapped
- Parallel flags set
- Epic updated with task summary

### 3. GitHub Sync
```
Task Files → GitHub Issues → File Rename
```
- Epic issue created
- Sub-issues for each task
- Files renamed: 001.md → 123.md
- References updated (depends_on)
- Mapping file created

### 4. Execution Start
```
Worktree Created → Ready Issues → Agents Launched
```
- git worktree for isolated development
- Dependency analysis for ready issues
- Parallel agents spawned
- Status tracking initialized

### 5. Progress Tracking
```
Agent Updates → Status Check → Epic Refresh
```
- Agents update progress files
- Status queries show current state
- Epic progress recalculated
- GitHub checkboxes updated

### 6. Merge
```
Tests Pass → Merge to Main → Cleanup
```
- Pre-merge validation
- Tests run
- Merge with --no-ff
- Worktree removed
- Branch deleted

### 7. Close
```
All Tasks Done → Archive → GitHub Closed
```
- Verify all tasks closed
- Epic archived
- GitHub issues closed
- PRD marked complete

## Dependency Management

### Dependency Types

**Hard Dependencies** (blocks start):
- Database schema before API
- Types before implementation
- Core logic before tests

**Soft Dependencies** (can start, blocks completion):
- Documentation can draft during implementation
- Tests can scaffold before code complete

### Dependency Notation

In task frontmatter:
```yaml
depends_on: [123, 124]      # Must complete before this starts
conflicts_with: [125, 126]  # Modifies same files
```

### Dependency Resolution

When starting epic:
1. Build dependency graph
2. Identify tasks with no unmet dependencies
3. Launch parallel agents for ready tasks
4. As tasks complete, check newly-ready tasks
5. Launch new agents for newly-ready work

## Parallelization Strategies

### By Architectural Layer
```
Database Layer ──────────────────────►
Service Layer  ──────────────────────►  (after DB types)
API Layer      ──────────────────────►  (after Services)
UI Layer       ──────────────────────►  (after API types)
```

### By Feature Slice
```
Feature A (Backend) ──────►
Feature A (Frontend) ─────►  (parallel)
Feature B (Full)     ─────►  (parallel)
```

### Conflict Avoidance
- Group tasks by directory/module
- Assign clear file ownership
- Use conflicts_with for coordination
- Run tests after each stream merge

## Error Recovery

### Partial Sync
- Some issues created, some failed
- Retry failed tasks only
- Check github-mapping.md for successes

### Merge Conflicts
- Preserve worktree
- Show conflicted files
- Provide resolution steps
- Allow abort option

### Agent Failures
- Log failure in execution-status.md
- Continue with other agents
- Offer retry option

## Best Practices

### Epic Size
- 5-15 tasks per epic
- 1-2 weeks of work
- Clear completion criteria

### Task Granularity
- 1-4 hours per task
- Single responsibility
- Testable completion

### Parallel Execution
- Maximize parallelization
- Minimize dependencies
- Clear file ownership
- Frequent commits

### Progress Visibility
- Regular status refreshes
- GitHub sync for stakeholders
- Clear done criteria
