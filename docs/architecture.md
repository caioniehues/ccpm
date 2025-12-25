# CCPM System Architecture

This document describes the architecture of Claude Code PM (CCPM), a system designed for high-context, parallel development with AI agents.

## Overview

CCPM uses a three-layer architecture to separate user intent from implementation details, with specialized agents handling heavy computational or context-intensive tasks.

```
┌─────────────────────────────────────────────────────────────────┐
│                         USER INPUT                              │
│                    /pm:prd-new my-feature                       │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│  LAYER 1: COMMANDS (Entry Points)                               │
│  ┌─────────────────────────────────────────────────────────────┐│
│  │ ccpm/commands/pm/prd-new.md                                 ││
│  │ - Thin wrapper for user invocation                          ││
│  │ - Minimal routing logic                                     ││
│  │ - References specific workflows                             ││
│  └─────────────────────────────────────────────────────────────┘│
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│  LAYER 2: SKILLS (Domain Knowledge)                             │
│  ┌─────────────────────────────────────────────────────────────┐│
│  │ ccpm/skills/ccpm-prd/SKILL.md                               ││
│  │ - Domain expertise (PRDs, Epics, Issues)                    ││
│  │ - Intake prompts and intent discovery                       ││
│  │ - Routes to appropriate workflow based on intent            ││
│  └─────────────────────────────────────────────────────────────┘│
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│  LAYER 3: WORKFLOWS (Implementation)                            │
│  ┌─────────────────────────────────────────────────────────────┐│
│  │ ccpm/skills/ccpm-prd/workflows/new-prd.md                   ││
│  │ - Detailed step-by-step execution                           ││
│  │ - Preflight checks, validation, and error handling          ││
│  │ - Defines success criteria                                  ││
│  └─────────────────────────────────────────────────────────────┘│
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│  AGENTS (Context-Isolated Workers)                              │
│  ┌─────────────────────────────────────────────────────────────┐│
│  │ ccpm/agents/prd-architect.md                                ││
│  │ - Performs heavy lifting in isolated context                ││
│  │ - Reads many files, returns concise summaries               ││
│  │ - Prevents main context pollution                           ││
│  └─────────────────────────────────────────────────────────────┘│
└─────────────────────────────────────────────────────────────────┘
```

## Layer Explanations

### Layer 1: Commands (Entry Points)
Commands are the user-facing interface, located in `ccpm/commands/`. They correspond directly to slash commands (e.g., `/pm:prd-new`).
- **Purpose**: Route user input to the correct skill and workflow.
- **Characteristics**: Minimal logic, declarative configuration, clear argument hints.
- **Key Directive**: Uses `Load workflow:` to delegate execution to Layer 3.

### Layer 2: Skills (Domain Knowledge)
Skills represent functional domains, located in `ccpm/skills/`.
- **Purpose**: Provide the "brain" for a specific domain (e.g., Managing PRDs, Syncing Epics).
- **Structure**: Each skill has a `SKILL.md` defining its objectives, principles, and routing logic.
- **Role**: Disambiguates user intent (e.g., "create new" vs "edit existing") and selects the right workflow.

### Layer 3: Workflows (Implementation)
Workflows are the executable procedures, located in `ccpm/skills/{skill}/workflows/`.
- **Purpose**: Define the exact steps to complete a task.
- **Components**:
    - **Preflight Checks**: Validation before starting (e.g., "Does file exist?").
    - **Execution Steps**: Detailed instructions, often with specific Bash commands.
    - **Error Handling**: Recovery paths for common failures.
    - **Success Criteria**: verifiable conditions for completion.

### Agents (Context-Isolated Workers)
Agents are specialized sub-instructions located in `ccpm/agents/`.
- **Purpose**: Handle high-volume or complex tasks without overwhelming the main conversation context.
- **Mechanism**: They run in a separate context window, process data (reading files, running tests), and return only a concise summary.
- **Philosophy**: "Agents are context firewalls."

## Data Flow

### Example: Creating a New PRD

1.  **User Invocation**: `/pm:prd-new user-auth`
2.  **Command Layer**: `ccpm/commands/pm/prd-new.md` identifies the task and loads the `new-prd` workflow.
3.  **Workflow Layer**: `ccpm/skills/ccpm-prd/workflows/new-prd.md` executes:
    *   **Preflight**: Checks if `user-auth` is valid and doesn't exist.
    *   **Execution**:
        *   Loads `datetime.md` reference.
        *   Creates `.claude/prds/user-auth.md`.
        *   (Optional) May spawn `prd-architect` agent if complex analysis is needed.
    *   **Validation**: Verifies file creation.
4.  **Completion**: Returns success message to user.

### Example: Epic Decomposition with Agents

1.  **User Invocation**: `/pm:epic-decompose payment-system`
2.  **Workflow Layer**: Reads the epic file.
3.  **Agent Delegation**: Spawns `task-decomposer` agent.
    *   *Agent Context*: Reads epic, analyzes requirements, generates 15 task files.
    *   *Agent Return*: "Created 15 task files." (Concise)
4.  **Main Context**: Receives summary, updates `epic.md` list, and finishes. Main context is not polluted with the content of 15 new files.

## File Structure

The project follows a strict directory layout to support this architecture:

```
ccpm/
├── agents/                    # Agent definitions (e.g., code-analyzer.md)
│   ├── code-analyzer.md
│   ├── epic-planner.md
│   └── ...
├── commands/                  # User entry points
│   ├── pm/                    # Main PM commands
│   │   ├── prd-new.md
│   │   ├── epic-sync.md
│   │   └── ...
│   ├── context/               # Context management
│   └── testing/               # Test execution
├── skills/                    # Domain knowledge bundles
│   ├── ccpm-prd/              # PRD Skill
│   │   ├── SKILL.md
│   │   ├── workflows/         # Executable workflows
│   │   └── templates/
│   ├── ccpm-epic/             # Epic Skill
│   ├── ccpm-issue/            # Issue Skill
│   ├── ccpm-worktree/         # Worktree Skill
│   └── shared-references/     # Cross-skill docs (datetime, git)
├── rules/                     # Global behavior rules
├── scripts/                   # Helper shell scripts
├── hooks/                     # Git hooks
└── cmd/                       # Binary tools (TUI)
    └── ccpm-tui/
```

## Design Principles

1.  **Single Responsibility**:
    *   Commands route.
    *   Skills understand.
    *   Workflows execute.
    *   Agents process.

2.  **No Circular References**:
    *   Dependencies flow downwards: Command → Workflow, Skill → Workflow.
    *   Never Command → Skill → Command.

3.  **Progressive Disclosure**:
    *   Users see simple commands.
    *   The system loads complex context only when required by the specific workflow.

4.  **Context Preservation**:
    *   Main conversation thread remains high-level.
    *   "Noisy" operations (reading logs, searching code, generating boilerplate) are offloaded to agents.

5.  **Fail Fast**:
    *   Preflight checks must run before any state change.
    *   Validation is mandatory after every major step.

## Common Patterns

### Command → Workflow
Commands should not contain logic. They simply point to the implementation:
```yaml
<context>
Load workflow: @ccpm/skills/ccpm-prd/workflows/new-prd.md
</context>
```

### Skill Routing
Skills use a routing table to map user intent to specific workflows:
```markdown
<routing>
| Response | Workflow | Purpose |
|----------|----------|---------|
| "new"    | workflows/new-prd.md | Create new PRD |
| "edit"   | workflows/edit-prd.md | Edit existing PRD |
</routing>
```

### Agent Delegation
Use agents when the input or output text volume is high:
```markdown
Task:
  description: "Analyze issue for parallel work"
  subagent_type: "task-decomposer"
```

## Extending CCPM

### Adding a New Command
1.  Create `ccpm/commands/pm/my-command.md`.
2.  Define the objective and reference a workflow.
3.  Register in `COMMANDS.md`.

### Adding a New Workflow
1.  Identify the parent Skill (or create a new one).
2.  Create `ccpm/skills/{skill}/workflows/my-workflow.md`.
3.  Implement Preflight, Execution, and Success Criteria.
4.  Add to the Skill's `workflows_index`.

### Adding a New Agent
1.  Create `ccpm/agents/my-agent.md`.
2.  Define the Input → Process → Output pattern.
3.  Ensure the return value is a summary, not raw data.
