# CCPM Feature Adoption Specification

> **Source**: Features adopted from `claude-code-spec-workflow`
> **Target**: CCPM (Claude Code Project Management)
> **Date**: 2025-12-23
> **Status**: Approved for Implementation

---

## Executive Summary

This document specifies exactly which features from `spec-workflow` should be adopted into CCPM, how they should be implemented, and how they integrate with existing CCPM commands.

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                          ADOPTION SUMMARY                                    │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                              │
│  ADOPT (5 features)                                                          │
│  ──────────────────                                                         │
│  ✅ Phase-based wizard with approval gates                                  │
│  ✅ ✅ APPROVED status markers in documents                                  │
│  ✅ TUI Dashboard (Charm-based)                                             │
│  ✅ Local-first issue tracking (GitHub optional)                            │
│  ✅ Steering documents pattern                                              │
│                                                                              │
│  ADAPT (2 features)                                                          │
│  ─────────────────                                                          │
│  ⚡ File-based status persistence (already partial in CCPM)                 │
│  ⚡ Context caching (integrate with /context:prime)                         │
│                                                                              │
│  SKIP (3 features)                                                           │
│  ────────────────                                                           │
│  ❌ Bug tracking workflow (CCPM already has issue tracking)                 │
│  ❌ Remote dashboard sharing (low priority for solo dev)                    │
│  ❌ Multi-project discovery (CCPM is single-project focused)                │
│                                                                              │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## Feature 1: Phase-Based Wizard with Approval Gates

### What to Adopt

The wizard pattern from spec-workflow that:
1. Pauses between phases for user review
2. Requires explicit approval before proceeding
3. Marks documents with approval status
4. Supports revision loops

### CCPM Integration

| Spec-Workflow | CCPM Equivalent | Integration |
|---------------|-----------------|-------------|
| `/spec-create` | `/pm:epic-wizard` | **NEW COMMAND** - Orchestrates existing commands |
| Phase 1: Requirements | Phase 1: PRD | Uses existing `/pm:prd-new` |
| Phase 2: Design | Phase 2: Epic | Uses existing `/pm:prd-parse` |
| Phase 3: Tasks | Phase 3: Tasks | Uses existing `/pm:epic-decompose` |
| Phase 4: Execute | Phase 4: Sync | Uses existing `/pm:epic-sync` (OPTIONAL) |

### New Command Specification

```yaml
# ccpm/commands/pm/epic-wizard.md
---
description: Guided epic creation with approval gates between phases
arguments:
  - name: name
    description: Epic name (e.g., user-authentication)
    required: true
allowed_tools:
  - Read
  - Write
  - Edit
  - Bash
  - AskUserQuestion
---
```

### Wizard Flow

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                           /pm:epic-wizard {name}                             │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                              │
│  ┌─────────────────────────────────────────────────────────────────────┐    │
│  │  PHASE 1: PRD CREATION                                              │    │
│  │  ─────────────────────                                              │    │
│  │                                                                      │    │
│  │  1. Execute: /pm:prd-new {name}                                     │    │
│  │  2. Output: .claude/prds/{name}.md                                  │    │
│  │  3. PAUSE → Ask: "Review the PRD. Approve to continue?"             │    │
│  │                                                                      │    │
│  │  User Response:                                                      │    │
│  │  ├─ "yes/approved/looks good" → Add ✅ APPROVED, proceed            │    │
│  │  └─ "no/changes needed" → Collect feedback, revise, ask again       │    │
│  └─────────────────────────────────────────────────────────────────────┘    │
│                              │                                               │
│                              ▼                                               │
│  ┌─────────────────────────────────────────────────────────────────────┐    │
│  │  PHASE 2: EPIC GENERATION                                           │    │
│  │  ────────────────────────                                           │    │
│  │                                                                      │    │
│  │  1. Execute: /pm:prd-parse {name}                                   │    │
│  │  2. Output: .claude/epics/{name}/epic.md                            │    │
│  │  3. PAUSE → Ask: "Review the Epic structure. Approve?"              │    │
│  │                                                                      │    │
│  │  User Response:                                                      │    │
│  │  ├─ Approve → Add ✅ APPROVED, proceed                              │    │
│  │  └─ Revise → Collect feedback, regenerate, ask again                │    │
│  └─────────────────────────────────────────────────────────────────────┘    │
│                              │                                               │
│                              ▼                                               │
│  ┌─────────────────────────────────────────────────────────────────────┐    │
│  │  PHASE 3: TASK DECOMPOSITION                                         │    │
│  │  ───────────────────────────                                         │    │
│  │                                                                      │    │
│  │  1. Execute: /pm:epic-decompose {name}                              │    │
│  │  2. Output: .claude/epics/{name}/001.md, 002.md, ...                │    │
│  │  3. PAUSE → Ask: "Review {N} tasks. Ready to begin work?"           │    │
│  │                                                                      │    │
│  │  User Response:                                                      │    │
│  │  ├─ Approve → Mark epic as "approved-for-work", proceed             │    │
│  │  └─ Revise → Adjust tasks, ask again                                │    │
│  └─────────────────────────────────────────────────────────────────────┘    │
│                              │                                               │
│                              ▼                                               │
│  ┌─────────────────────────────────────────────────────────────────────┐    │
│  │  PHASE 4: GITHUB SYNC (OPTIONAL)                                     │    │
│  │  ─────────────────────────────                                       │    │
│  │                                                                      │    │
│  │  Ask: "Sync to GitHub?"                                             │    │
│  │                                                                      │    │
│  │  ├─ "Yes" → Execute /pm:epic-sync {name}                            │    │
│  │  ├─ "No, local only" → Skip, work locally                           │    │
│  │  └─ "Later" → Skip, user can run /pm:epic-sync manually             │    │
│  │                                                                      │    │
│  │  Success message: "Epic ready! Run /pm:issue-start to begin."       │    │
│  └─────────────────────────────────────────────────────────────────────┘    │
│                                                                              │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Approval Phrases

Recognize these as approval:
- `yes`, `y`
- `approved`, `approve`
- `looks good`, `lgtm`
- `proceed`, `continue`
- `ok`, `okay`

Recognize these as revision request:
- `no`, `n`
- `changes`, `change`
- `revise`, `revision`
- `edit`, `modify`
- Any response containing specific feedback

---

## Feature 2: Status Markers in Documents

### What to Adopt

The `✅ APPROVED` marker pattern that:
1. Persists approval state in the document itself
2. Enables phase recovery on re-run
3. Is human-readable in markdown

### Implementation

Add to document frontmatter:

```yaml
---
status: approved  # pending | approved | revision-needed
approved_at: 2025-12-23T14:30:00Z
approved_by: user
revision_count: 0
---

# Document Title

...
```

Or as inline marker (spec-workflow style):

```markdown
# User Authentication PRD

✅ APPROVED (2025-12-23)

## Executive Summary
...
```

### CCPM Decision: Use Frontmatter

Frontmatter is preferred because:
- Already used in CCPM task files
- Machine-parseable
- Supports additional metadata (timestamps, revision count)

### Phase Recovery Logic

When `/pm:epic-wizard {name}` is run:

```
1. Check if .claude/prds/{name}.md exists
   └─ If exists AND has status: approved → Skip Phase 1

2. Check if .claude/epics/{name}/epic.md exists
   └─ If exists AND has status: approved → Skip Phase 2

3. Check if .claude/epics/{name}/001.md exists
   └─ If any task exists → Resume at Phase 3 or 4

4. Otherwise → Start from Phase 1
```

---

## Feature 3: TUI Dashboard

### What to Adopt

A terminal-native dashboard using Charm libraries, providing:
- Real-time epic/task visualization
- File watching for live updates
- Wizard mode integration
- Keyboard-driven navigation

### Implementation

See:
- [TUI-WIREFRAMES.md](./TUI-WIREFRAMES.md) for visual design (18 wireframes)
- [TUI-DASHBOARD-DESIGN.md](./TUI-DASHBOARD-DESIGN.md) for technical architecture

### New Command

```yaml
# ccpm/commands/pm/dashboard.md
---
description: Launch the TUI dashboard for visual epic/task tracking
allowed_tools:
  - Bash
---
```

Behavior:
```bash
# Launch the dashboard binary
ccpm-tui

# Or with specific epic focus
ccpm-tui --epic user-authentication
```

### Technology Stack

| Component | Library |
|-----------|---------|
| Framework | Bubbletea |
| Components | Bubbles |
| Styling | Lipgloss |
| Forms | Huh |
| Markdown | Glamour |
| File Watch | fsnotify |

---

## Feature 4: Local-First Issue Tracking

### What to Adopt

The ability to track task status locally without requiring GitHub sync.

### Current CCPM Gap

Current workflow requires:
1. `/pm:epic-sync` to push issues to GitHub
2. `/pm:issue-start` to start work (requires GitHub issue)

### New Commands

#### /pm:task-start-local

```yaml
---
description: Start working on a task locally (no GitHub required)
arguments:
  - name: epic
    description: Epic name
    required: true
  - name: task
    description: Task number (e.g., 001)
    required: true
---
```

Behavior:
1. Read `.claude/epics/{epic}/{task}.md`
2. Update frontmatter:
   ```yaml
   status: in-progress
   started_at: 2025-12-23T14:30:00Z
   ```
3. Create working branch if not exists
4. Output context for starting work

#### /pm:task-complete-local

```yaml
---
description: Mark a task complete locally
arguments:
  - name: epic
    description: Epic name
    required: true
  - name: task
    description: Task number
    required: true
---
```

Behavior:
1. Update task frontmatter:
   ```yaml
   status: completed
   completed_at: 2025-12-23T16:45:00Z
   ```
2. Update epic progress

#### /pm:sync-when-ready

```yaml
---
description: Batch sync all local changes to GitHub
arguments:
  - name: epic
    description: Epic name to sync
    required: true
---
```

Behavior:
1. Find all tasks with local-only changes
2. Create/update GitHub issues for each
3. Update local files with issue numbers
4. Report sync results

### Task Frontmatter Schema

```yaml
---
id: "001"
title: "Set up auth middleware"
epic: user-authentication
status: pending  # pending | in-progress | completed | blocked | verified
started_at: null
completed_at: null
github_issue: null  # Populated after sync
dependencies: []
blocked_by: null
---
```

---

## Feature 5: Steering Documents

### What to Adopt

Three curated context documents that provide consistent project context:
- `product.md` - What we're building
- `tech.md` - How we build it
- `structure.md` - Where things go

### CCPM Integration

Create `.claude/steering/` directory:

```
.claude/
├── steering/
│   ├── product.md    # Vision, goals, target users
│   ├── tech.md       # Tech stack, standards, conventions
│   └── structure.md  # Code organization, file patterns
├── prds/
├── epics/
└── ...
```

### New Command

#### /pm:steering-setup

```yaml
---
description: Initialize steering documents for the project
---
```

Behavior:
1. Create `.claude/steering/` directory
2. Generate initial templates for each document
3. Guide user through filling them out

### Steering Document Templates

#### product.md
```markdown
# Product Context

## Vision
[What is the ultimate goal of this product?]

## Target Users
[Who are we building this for?]

## Core Features
[What are the must-have features?]

## Non-Goals
[What are we explicitly NOT building?]
```

#### tech.md
```markdown
# Technical Context

## Stack
- Language:
- Framework:
- Database:
- Infrastructure:

## Standards
[Coding standards, naming conventions]

## Dependencies
[Key libraries and why they were chosen]
```

#### structure.md
```markdown
# Project Structure

## Directory Layout
```
project/
├── src/
├── tests/
└── ...
```

## File Naming
[Conventions for file names]

## Module Organization
[How code is organized]
```

### Auto-Loading in Commands

All `/pm:` commands should:
1. Check for `.claude/steering/` existence
2. Load steering documents as context
3. Reference them in prompts

---

## Feature 6: Context Caching (Adapt)

### What to Adapt

File-based caching with TTL to reduce repeated parsing.

### Integration with /context:prime

Current `/context:prime` behavior:
- Reads context files on each invocation
- No caching

Enhanced behavior:
1. Cache parsed context with 30-minute TTL
2. Validate cache via file mtime comparison
3. Auto-invalidate on file changes

### Implementation

Add to context module:

```
Cache Key: file path
Cache Value: { content, parsed_data, mtime, cached_at }
TTL: 30 minutes
Validation: Compare current mtime with cached mtime
```

---

## Features NOT Adopted

### Bug Tracking Workflow

**Reason**: CCPM already has GitHub issue integration which handles bugs.

spec-workflow has:
- `.claude/bugs/` with 4-phase workflow
- report.md → analysis.md → fix.md → verification.md

CCPM alternative:
- Use GitHub issues with labels
- Track via existing `/pm:issue-*` commands

### Remote Dashboard Sharing

**Reason**: Low priority for solo developer focus.

spec-workflow has:
- Cloudflare tunnel integration
- Ngrok tunnel integration

CCPM alternative:
- Use tmux/screen for terminal sharing
- SSH access for remote work

### Multi-Project Discovery

**Reason**: CCPM is single-project focused.

spec-workflow has:
- Scans ~/Projects, ~/Code, etc.
- Discovers all projects with .claude/

CCPM alternative:
- Work within current project
- Switch projects via shell/terminal

---

## Implementation Priority

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                         IMPLEMENTATION SCHEDULE                              │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                              │
│  WEEK 1: Wizard & Approval Gates                                             │
│  ─────────────────────────────────                                          │
│  □ Create /pm:epic-wizard command                                           │
│  □ Implement approval phrase detection                                      │
│  □ Add frontmatter status fields to documents                               │
│  □ Implement phase recovery logic                                           │
│  □ Add revision loop support                                                │
│                                                                              │
│  WEEK 2: TUI Dashboard Core                                                  │
│  ──────────────────────────                                                 │
│  □ Set up Go project with Charm dependencies                                │
│  □ Implement Model/Update/View structure                                    │
│  □ Create epic/task file parsers                                            │
│  □ Add file watcher (fsnotify)                                              │
│  □ Build main dashboard view                                                │
│                                                                              │
│  WEEK 3: Local Tracking & Dashboard Views                                    │
│  ─────────────────────────────────────────                                  │
│  □ Create /pm:task-start-local command                                      │
│  □ Create /pm:task-complete-local command                                   │
│  □ Create /pm:sync-when-ready command                                       │
│  □ Build epic/task/PRD detail views                                         │
│  □ Add wizard mode to dashboard                                             │
│                                                                              │
│  WEEK 4: Steering & Polish                                                   │
│  ─────────────────────────                                                  │
│  □ Create /pm:steering-setup command                                        │
│  □ Integrate steering docs into all /pm: commands                           │
│  □ Implement context caching                                                │
│  □ Add dashboard animations and loading states                              │
│  □ Handle terminal resize                                                   │
│                                                                              │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## Success Criteria

After implementation, verify:

| Feature | Success Criteria |
|---------|------------------|
| Wizard | Can create epic from idea to tasks with review at each phase |
| Approval Gates | Documents show ✅ APPROVED status after user approval |
| Phase Recovery | Re-running wizard skips already-approved phases |
| TUI Dashboard | Visual progress tracking works offline |
| Local Tracking | Can complete full workflow without GitHub |
| Steering Docs | All /pm: commands reference steering context |

---

## Related Documents

| Document | Purpose |
|----------|---------|
| [TUI-WIREFRAMES.md](./TUI-WIREFRAMES.md) | Visual design for dashboard |
| [TUI-DASHBOARD-DESIGN.md](./TUI-DASHBOARD-DESIGN.md) | Technical architecture |
| [SPEC-WORKFLOW-FEATURE-ANALYSIS.md](./SPEC-WORKFLOW-FEATURE-ANALYSIS.md) | Feature analysis source |
| [TUI-DESIGN-INDEX.md](./TUI-DESIGN-INDEX.md) | Navigation index |

---

*This specification is approved for implementation. Implementation should follow the weekly schedule above.*
