# CCPM Skills

Skills are Claude-executable workflows organized by domain. They provide structured guidance for specific CCPM operations.

## Architecture

CCPM uses a three-layer architecture:

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
       │ implementation)│
       └───────────────┘
```

- **Commands** provide the user interface (`/pm:prd-new`)
- **Skills** organize related workflows and provide domain context
- **Workflows** contain the actual implementation steps

Both commands and skills route to the same workflow files, avoiding circular references.

## Available Skills

### 📋 `ccpm-prd`

**Purpose**: Product Requirements Document lifecycle management

**Operations**:
| Operation | Workflow | Description |
|-----------|----------|-------------|
| Create | `new-prd.md` | Brainstorm and create new PRD through structured discovery |
| Edit | `edit-prd.md` | Modify existing PRD sections |
| Parse | `parse-prd.md` | Convert PRD to technical implementation epic |

**Quick Start**:
```bash
/pm:prd-new user-authentication    # Create new PRD
/pm:prd-parse user-authentication  # Convert to epic
/pm:prd-edit user-authentication   # Edit PRD
```

---

### 📐 `ccpm-epic`

**Purpose**: Epic lifecycle management from planning through completion

**Operations**:
| Operation | Workflow | Description |
|-----------|----------|-------------|
| Decompose | `decompose-epic.md` | Break epic into actionable tasks |
| Sync | `sync-epic.md` | Push epic and tasks to GitHub |
| Oneshot | `oneshot-epic.md` | Decompose and sync in one operation |
| Start | `start-epic.md` | Launch parallel agents |
| Edit | `edit-epic.md` | Modify epic details |
| Refresh | `refresh-epic.md` | Update progress from task states |
| Close | `close-epic.md` | Mark epic as complete |
| Merge | `merge-epic.md` | Merge epic branch to main |

**Quick Start**:
```bash
/pm:epic-decompose my-feature   # Break into tasks
/pm:epic-sync my-feature        # Push to GitHub
/pm:epic-oneshot my-feature     # Decompose + sync together
/pm:epic-start my-feature       # Launch agents
```

---

### 🎯 `ccpm-issue`

**Purpose**: Issue lifecycle with local-first GitHub sync

**Operations**:
| Operation | Workflow | Description |
|-----------|----------|-------------|
| Analyze | `analyze-issue.md` | Identify parallel work streams |
| Start | `start-issue.md` | Begin work with parallel agents |
| Status | `status-issue.md` | Check issue state and progress |
| Sync | `sync-issue.md` | Push local updates to GitHub |
| Show | `show.md` | Display full issue details |
| Edit | `edit-issue.md` | Update issue details |
| Close | `close-issue.md` | Mark issue as complete |
| Reopen | `reopen-issue.md` | Reopen closed issue |

**Quick Start**:
```bash
/pm:issue-start 123    # Start work on issue
/pm:issue-status 123   # Check progress
/pm:issue-sync 123     # Push updates to GitHub
```

---

### 🗂️ `ccpm-context`

**Purpose**: Project context management for Claude sessions

**Operations**:
| Operation | Workflow | Description |
|-----------|----------|-------------|
| Create | `create-context.md` | Establish initial context documentation |
| Prime | `prime-context.md` | Load context for new session |
| Update | `update-context.md` | Refresh context with recent changes |

**Context Files Created**:
- `progress.md` - Current status, recent work
- `project-structure.md` - Directory organization
- `tech-context.md` - Dependencies and tools
- `system-patterns.md` - Architecture patterns
- `product-context.md` - Requirements and users

**Quick Start**:
```bash
/context:create   # Create initial context
/context:prime    # Load context for session
/context:update   # Refresh after changes
```

---

### 🧪 `ccpm-testing`

**Purpose**: Test environment preparation and execution

**Operations**:
| Operation | Workflow | Description |
|-----------|----------|-------------|
| Prime | `prime-testing.md` | Detect frameworks, configure test-runner |
| Run | `run-tests.md` | Execute tests with verbose output |

**Supported Frameworks** (12+ languages):
- JavaScript: Jest, Mocha, Jasmine
- Python: Pytest, unittest
- Go: go test
- Rust: cargo test
- Java/Kotlin: JUnit, Maven, Gradle
- C#/.NET: MSTest, NUnit, xUnit
- PHP: PHPUnit, Pest
- Ruby: RSpec, Minitest
- And more...

**Quick Start**:
```bash
/testing:prime              # Setup test environment
/testing:run                # Run all tests
/testing:run path/to/test   # Run specific test
```

---

### 🌳 `ccpm-worktree`

**Purpose**: Git worktree management for parallel development

**Operations**:
| Operation | Description |
|-----------|-------------|
| Create | Create worktree for epic development |
| List | Show all active worktrees |
| Status | Check specific worktree state |
| Merge | Merge worktree branch to main |
| Remove | Clean up worktree after merge |
| Prune | Remove stale references |

**Worktree Structure**:
```
project/
├── main-repo/          (primary)
├── epic-feature-a/     (worktree)
├── epic-feature-b/     (worktree)
└── epic-refactor/      (worktree)
```

**Quick Start**:
```bash
/pm:epic-start-worktree my-feature  # Create worktree
git worktree list                    # List all worktrees
/pm:epic-merge my-feature           # Merge and cleanup
```

---

## Skill Structure

Each skill follows a consistent structure:

```
ccpm/skills/{skill-name}/
├── SKILL.md           # Main skill definition
├── references/        # Domain knowledge documents
├── templates/         # Template files (if applicable)
└── workflows/         # Executable workflow files
    ├── workflow-1.md
    ├── workflow-2.md
    └── ...
```

### SKILL.md Components

1. **Frontmatter**: Name and description for routing
2. **Objective**: What the skill accomplishes
3. **Essential Principles**: Domain knowledge and key concepts
4. **Intake**: User interaction prompt
5. **Routing**: Decision table for workflow selection
6. **Workflows Index**: Available workflow files
7. **Quick Start**: Common usage patterns
8. **Success Criteria**: How to verify correct execution

## Shared References

Common patterns used across skills are in `shared-references/`:

| Reference | Purpose |
|-----------|---------|
| `agent-coordination.md` | Multi-agent coordination patterns |
| `datetime.md` | Real datetime handling (never placeholders) |
| `frontmatter-operations.md` | YAML frontmatter structure |
| `github-operations.md` | GitHub CLI patterns |
| `worktree-operations.md` | Git worktree commands |

## Creating New Skills

1. Create skill directory: `ccpm/skills/{skill-name}/`
2. Create `SKILL.md` with required sections
3. Create `workflows/` directory with workflow files
4. Add `references/` for domain knowledge (optional)
5. Update this README with skill documentation

### Key Principles

- **Single Responsibility**: Each skill handles one domain
- **Workflow Delegation**: Skills route to workflow files, not commands
- **No Circular References**: Never route skill → command → skill
- **Real Timestamps**: Always use system clock for datetimes
- **Clear Routing**: Intake → Routing table → Workflow file
