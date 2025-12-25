# CCPM Architecture

This document explains the layered architecture of Claude Code PM (CCPM), including how commands, skills, workflows, and agents work together.

## Table of Contents

- [Overview](#overview)
- [The Three-Layer Architecture](#the-three-layer-architecture)
- [Layer 1: Commands](#layer-1-commands)
- [Layer 2: Skills](#layer-2-skills)
- [Layer 3: Workflows](#layer-3-workflows)
- [Agents](#agents)
- [Data Flow](#data-flow)
- [File Structure](#file-structure)
- [Design Principles](#design-principles)
- [Common Patterns](#common-patterns)

## Overview

CCPM uses a three-layer architecture to organize functionality:

```
┌─────────────────────────────────────────────────────────────────┐
│                         USER INPUT                               │
│                    /pm:prd-new my-feature                        │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│  LAYER 1: COMMANDS                                               │
│  ┌─────────────────────────────────────────────────────────────┐│
│  │ ccpm/commands/pm/prd-new.md                                 ││
│  │ - Entry point for user invocation                           ││
│  │ - Minimal routing logic                                     ││
│  │ - References workflow file                                  ││
│  └─────────────────────────────────────────────────────────────┘│
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│  LAYER 2: SKILLS                                                 │
│  ┌─────────────────────────────────────────────────────────────┐│
│  │ ccpm/skills/ccpm-prd/SKILL.md                               ││
│  │ - Domain knowledge and context                              ││
│  │ - Intake prompts for user interaction                       ││
│  │ - Routes to appropriate workflow                            ││
│  └─────────────────────────────────────────────────────────────┘│
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│  LAYER 3: WORKFLOWS                                              │
│  ┌─────────────────────────────────────────────────────────────┐│
│  │ ccpm/skills/ccpm-prd/workflows/new-prd.md                   ││
│  │ - Detailed step-by-step implementation                      ││
│  │ - Preflight checks and validation                           ││
│  │ - Error handling                                            ││
│  │ - Success criteria                                          ││
│  └─────────────────────────────────────────────────────────────┘│
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│  AGENTS (when needed)                                            │
│  ┌─────────────────────────────────────────────────────────────┐│
│  │ ccpm/agents/prd-architect.md                                ││
│  │ - Heavy lifting in isolated context                         ││
│  │ - Returns concise summary                                   ││
│  └─────────────────────────────────────────────────────────────┘│
└─────────────────────────────────────────────────────────────────┘
```

## The Three-Layer Architecture

### Why Three Layers?

| Layer | Purpose | Analogy |
|-------|---------|---------|
| **Commands** | Entry point, routing | Restaurant menu |
| **Skills** | Domain knowledge, context | Chef's expertise |
| **Workflows** | Step-by-step execution | Recipe |

This separation provides:
- **Maintainability**: Change workflows without touching commands
- **Reusability**: Multiple commands can share workflows
- **Clarity**: Each layer has a single responsibility
- **Testability**: Workflows can be validated independently

## Layer 1: Commands

Commands are the entry points users invoke with `/command-name`.

### Location
```
ccpm/commands/
├── pm/              # Project management commands
│   ├── prd-new.md
│   ├── epic-sync.md
│   └── issue-start.md
├── context/         # Context management commands
│   ├── create.md
│   └── prime.md
└── testing/         # Testing commands
    ├── prime.md
    └── run.md
```

### Command Structure

```yaml
---
description: Brief description shown in help
argument-hint: <required_arg> [optional_arg]
allowed-tools: Bash, Read, Write, Glob, Task
---

<objective>
What this command accomplishes.
</objective>

<context>
Load workflow: @ccpm/skills/{skill}/workflows/{workflow}.md
Load reference: @ccpm/skills/shared-references/{reference}.md
</context>

<process>
1. Load the workflow
2. Execute steps
3. Return results
</process>

<success_criteria>
- Criterion 1
- Criterion 2
</success_criteria>
```

### Key Principles

1. **Minimal Logic**: Commands should be thin wrappers
2. **Reference Workflows**: Use `Load workflow:` to specify implementation
3. **Clear Tools**: Declare all required tools in frontmatter
4. **User-Facing**: Written for the user to understand

## Layer 2: Skills

Skills provide domain knowledge and context for a functional area.

### Location
```
ccpm/skills/
├── ccpm-prd/        # PRD management
├── ccpm-epic/       # Epic lifecycle
├── ccpm-issue/      # Issue tracking
├── ccpm-context/    # Context management
├── ccpm-testing/    # Test execution
├── ccpm-worktree/   # Git worktree management
└── shared-references/  # Common reference docs
```

### Skill Structure

```
ccpm/skills/{skill-name}/
├── SKILL.md         # Main skill file
├── workflows/       # Step-by-step implementations
│   ├── action1.md
│   ├── action2.md
│   └── action3.md
├── references/      # Domain-specific docs
│   └── patterns.md
└── templates/       # File templates
    └── template.md
```

### SKILL.md Anatomy

```markdown
---
name: skill-name
description: When to use this skill
---

<objective>
What this skill manages.
</objective>

<essential_principles>
Domain knowledge Claude needs to understand.
</essential_principles>

<intake>
Interactive prompt for user intent discovery.
</intake>

<routing>
| Response | Workflow | Purpose |
|----------|----------|---------|
| "new"    | workflows/new.md | Create new item |
| "edit"   | workflows/edit.md | Edit existing |
</routing>

<workflows_index>
List of all available workflows.
</workflows_index>

<success_criteria>
How to know operations succeeded.
</success_criteria>
```

### The Six Core Skills

| Skill | Purpose | Workflows |
|-------|---------|-----------|
| **ccpm-prd** | Product Requirements Documents | new, edit, parse |
| **ccpm-epic** | Epic lifecycle management | decompose, sync, start, close, merge, edit, refresh, oneshot |
| **ccpm-issue** | GitHub issue operations | analyze, start, status, sync, edit, close, reopen, show |
| **ccpm-context** | Project context | create, prime, update |
| **ccpm-testing** | Test execution | prime, run |
| **ccpm-worktree** | Git worktree management | create, list, merge, remove (inline) |

## Layer 3: Workflows

Workflows contain the actual implementation steps.

### Workflow Structure

```markdown
# Workflow: Action Name

Brief description.

## Input
- `$ARGUMENTS`: What the user provides

## Preflight Checks

1. **Validation Step**
   ```bash
   # Command to validate
   ```
   If fails: "Error message"

2. **Another Check**
   ...

## Execution Steps

### 1. First Step
Detailed instructions...

### 2. Second Step
More instructions...

## Output

```
Expected output format
```

## Error Handling

- **Error case 1**: Recovery action
- **Error case 2**: Recovery action

## Success Criteria

- [ ] Criterion 1
- [ ] Criterion 2
```

### Workflow Best Practices

1. **Be Explicit**: Don't assume knowledge
2. **Include Bash**: Show exact commands when applicable
3. **Handle Errors**: Every step should have failure handling
4. **Define Success**: Clear, checkable criteria

## Agents

Agents are specialized workers for heavy or context-intensive tasks.

### Location
```
ccpm/agents/
├── code-analyzer.md      # Bug hunting
├── file-analyzer.md      # Log/output summarization
├── test-runner.md        # Test execution
├── parallel-worker.md    # Parallel coordination
├── prd-architect.md      # PRD structure design
├── epic-planner.md       # Epic decomposition
├── task-decomposer.md    # Task breakdown
├── github-syncer.md      # GitHub operations
├── worktree-manager.md   # Worktree lifecycle
└── parallel-orchestrator.md  # Multi-agent coordination
```

### When to Use Agents

| Scenario | Use Agent? | Why |
|----------|------------|-----|
| Read 1-2 files | No | Simple, direct tools work |
| Read 10+ files | Yes | Protects main context |
| Run tests | Yes | Output is verbose |
| Complex analysis | Yes | Needs focused context |
| Simple edit | No | Overkill |

### Agent Philosophy

> "Agents are context firewalls, not knowledge specialists."

Agents don't have different knowledge than Claude—they have **isolated context**. Use them to:
- Process large amounts of data
- Return concise summaries
- Work in parallel without context collision

## Data Flow

### Example: Creating a New PRD

```
User: /pm:prd-new user-auth

1. COMMAND LOOKUP
   └─ ccpm/commands/pm/prd-new.md found

2. COMMAND PARSING
   └─ $ARGUMENTS = "user-auth"
   └─ allowed-tools: Bash, Read, Write, Glob

3. WORKFLOW LOADING
   └─ Load: ccpm/skills/ccpm-prd/workflows/new-prd.md
   └─ Load: ccpm/skills/shared-references/datetime.md

4. WORKFLOW EXECUTION
   ├─ Preflight: Validate "user-auth" is kebab-case ✓
   ├─ Preflight: Check no existing PRD ✓
   ├─ Preflight: Create .claude/prds/ directory ✓
   ├─ Preflight: Get current datetime ✓
   ├─ Brainstorm: Phase 1 - Problem Definition
   ├─ Brainstorm: Phase 2 - Solution Exploration
   ├─ Brainstorm: Phase 3 - Success Criteria
   ├─ Brainstorm: Phase 4 - Dependencies
   ├─ Create: .claude/prds/user-auth.md
   └─ Validate: File exists, has content

5. COMPLETION
   └─ Summary displayed to user
   └─ Next steps suggested
```

### Example: Epic with Agent Delegation

```
User: /pm:epic-decompose payment-system

1. WORKFLOW EXECUTION
   ├─ Read epic: .claude/epics/payment-system/epic.md
   ├─ Analyze scope: 15+ tasks identified
   └─ Decision: Delegate to epic-planner agent

2. AGENT DELEGATION
   ├─ Spawn: epic-planner agent
   ├─ Task: "Create task files for payment-system epic"
   ├─ Agent works in isolated context
   ├─ Agent creates 15 task files
   └─ Agent returns: "Created 15 tasks: [summary]"

3. MAIN THREAD CONTINUES
   ├─ Receives concise summary
   ├─ Updates epic.md with task list
   └─ Suggests: /pm:epic-sync payment-system
```

## File Structure

### Complete CCPM Layout

```
ccpm/
├── agents/                    # Agent definitions
│   ├── code-analyzer.md
│   ├── epic-planner.md
│   ├── file-analyzer.md
│   ├── github-syncer.md
│   ├── parallel-orchestrator.md
│   ├── parallel-worker.md
│   ├── prd-architect.md
│   ├── task-decomposer.md
│   ├── test-runner.md
│   └── worktree-manager.md
│
├── commands/                  # User-invokable commands
│   ├── pm/                    # Project management
│   │   ├── prd-new.md
│   │   ├── prd-edit.md
│   │   ├── prd-parse.md
│   │   ├── epic-decompose.md
│   │   ├── epic-sync.md
│   │   ├── epic-start.md
│   │   ├── issue-analyze.md
│   │   ├── issue-start.md
│   │   └── ... (46 total)
│   ├── context/
│   │   ├── create.md
│   │   ├── prime.md
│   │   └── update.md
│   └── testing/
│       ├── prime.md
│       └── run.md
│
├── skills/                    # Domain skills
│   ├── ccpm-prd/
│   │   ├── SKILL.md
│   │   ├── workflows/
│   │   │   ├── new-prd.md
│   │   │   ├── edit-prd.md
│   │   │   └── parse-prd.md
│   │   ├── references/
│   │   └── templates/
│   │
│   ├── ccpm-epic/
│   │   ├── SKILL.md
│   │   ├── workflows/
│   │   │   ├── decompose-epic.md
│   │   │   ├── sync-epic.md
│   │   │   ├── start-epic.md
│   │   │   ├── close-epic.md
│   │   │   ├── merge-epic.md
│   │   │   ├── edit-epic.md
│   │   │   ├── refresh-epic.md
│   │   │   └── oneshot-epic.md
│   │   └── references/
│   │
│   ├── ccpm-issue/
│   │   ├── SKILL.md
│   │   ├── workflows/
│   │   │   ├── analyze-issue.md
│   │   │   ├── start-issue.md
│   │   │   ├── status-issue.md
│   │   │   ├── sync-issue.md
│   │   │   ├── edit-issue.md
│   │   │   ├── close-issue.md
│   │   │   ├── reopen-issue.md
│   │   │   └── show.md
│   │   └── references/
│   │
│   ├── ccpm-context/
│   │   ├── SKILL.md
│   │   └── workflows/
│   │       ├── create-context.md
│   │       ├── prime-context.md
│   │       └── update-context.md
│   │
│   ├── ccpm-testing/
│   │   ├── SKILL.md
│   │   └── workflows/
│   │       ├── prime-testing.md
│   │       └── run-tests.md
│   │
│   ├── ccpm-worktree/
│   │   └── SKILL.md          # Inline workflows
│   │
│   └── shared-references/     # Cross-skill references
│       ├── agent-coordination.md
│       ├── datetime.md
│       ├── frontmatter-operations.md
│       ├── github-operations.md
│       └── worktree-operations.md
│
├── rules/                     # Claude behavior rules
├── scripts/                   # Shell scripts
├── hooks/                     # Git hooks
└── context/                   # Project context
```

## Design Principles

### 1. Single Responsibility
Each layer does one thing:
- Commands: Route user input
- Skills: Provide domain context
- Workflows: Execute steps

### 2. No Circular References
```
WRONG:
Command → Skill → Command (loop!)

RIGHT:
Command → Workflow (implementation)
Skill → Workflow (same implementation)
```

### 3. Progressive Disclosure
- Commands show minimal info
- Skills add context when needed
- Workflows provide full detail

### 4. Context Preservation
- Use agents for heavy processing
- Return summaries, not raw data
- Keep main thread focused

### 5. Fail Fast
- Validate inputs immediately
- Clear error messages
- Never leave partial state

## Common Patterns

### Pattern: Command → Workflow

```yaml
# Command
<context>
Load workflow: @ccpm/skills/ccpm-prd/workflows/new-prd.md
</context>
```

### Pattern: Skill Routing

```markdown
<routing>
| Response | Workflow | Purpose |
|----------|----------|---------|
| "new" | workflows/new-prd.md | Create new PRD |
| "edit" | workflows/edit-prd.md | Edit existing |
</routing>
```

### Pattern: Agent Delegation

```markdown
For complex tasks, delegate to agent:

Task:
  description: "Analyze issue for parallel work"
  subagent_type: "task-decomposer"
  prompt: |
    Analyze issue #1234 and identify parallel streams...
```

### Pattern: Shared References

```yaml
<context>
Load reference: @ccpm/skills/shared-references/datetime.md
Load reference: @ccpm/skills/shared-references/github-operations.md
</context>
```

## Extending CCPM

### Adding a New Command

1. Create `ccpm/commands/pm/my-command.md`
2. Reference existing or new workflow
3. Document in COMMANDS.md

### Adding a New Workflow

1. Create `ccpm/skills/{skill}/workflows/my-workflow.md`
2. Add to skill's routing table
3. Add to workflows_index

### Adding a New Skill

1. Create directory `ccpm/skills/my-skill/`
2. Create `SKILL.md` with standard structure
3. Add workflows directory
4. Document in this file

### Adding a New Agent

1. Create `ccpm/agents/my-agent.md`
2. Define clear purpose and pattern
3. Document in AGENTS.md

---

## Summary

The CCPM architecture separates concerns into three layers:

| Layer | Location | Purpose |
|-------|----------|---------|
| Commands | `ccpm/commands/` | User entry points |
| Skills | `ccpm/skills/*/SKILL.md` | Domain context |
| Workflows | `ccpm/skills/*/workflows/` | Implementation |

Plus agents for context-isolated heavy work.

This design enables:
- **Maintainable** code through separation of concerns
- **Reusable** workflows across commands
- **Scalable** architecture for new features
- **Testable** components in isolation
