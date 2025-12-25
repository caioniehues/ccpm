# CCPM Agent System

The Claude Code PM (CCPM) system relies on a network of specialized agents to handle heavy workloads while preserving the main conversation context. This architecture enables the system to manage large-scale projects without hitting context window limits or losing coherence.

## Core Philosophy

Agents act as **context firewalls**. Instead of dumping hundreds of lines of code, logs, or requirements into the main chat, specialized agents process this information and return only the high-signal summary that matters for decision-making.

> "Don't anthropomorphize subagents. Use them to organize your prompts and elide context. Subagents are best when they can do lots of work but then provide small amounts of information back to the main conversation thread."
>
> – Adam Wolff, Anthropic

## Available Agents

| Agent | Purpose | Tools |
|-------|---------|-------|
| [`code-analyzer`](#code-analyzer) | Hunt bugs and trace logic across files | Glob, Grep, Read |
| [`file-analyzer`](#file-analyzer) | Summarize verbose files (logs, configs) | Glob, Grep, LS, Read |
| [`test-runner`](#test-runner) | Execute tests and analyze failures | Bash, Glob, Grep, Read |
| [`parallel-worker`](#parallel-worker) | Execute work streams in a worktree | Bash, Task, TodoWrite |
| [`prd-architect`](#prd-architect) | Design and structure PRDs | Read, Write, Edit |
| [`epic-planner`](#epic-planner) | Plan epic execution strategy | Glob, Read, Write |
| [`task-decomposer`](#task-decomposer) | Break down epics into tasks | Read, Write, Glob |
| [`github-syncer`](#github-syncer) | Sync bi-directionally with GitHub | Bash, Task, Write |
| [`worktree-manager`](#worktree-manager) | Manage git worktrees | Bash, Read, Glob |
| [`parallel-orchestrator`](#parallel-orchestrator) | Coordinate multi-agent execution | Task, Bash, Glob |

## Detailed Agent Reference

### 🔍 code-analyzer
**Purpose**: Hunt bugs across multiple files without polluting main context.
- **Pattern**: Search many files → Analyze code → Return bug report
- **When to use**: When you need to trace logic flows, find bugs, or validate changes.
- **Capabilities**:
  - Logic tracing across files
  - Bug pattern recognition (race conditions, leaks, null refs)
  - Change impact analysis
- **Output**: Concise bug report with critical findings and actionable fixes.

### 📄 file-analyzer
**Purpose**: Read and summarize verbose files (logs, outputs, configs).
- **Pattern**: Read files → Extract insights → Return summary
- **When to use**: When you need to understand log files, debug outputs, or analyze verbose configurations.
- **Capabilities**:
  - Extracts errors, exceptions, and stack traces
  - Identifies patterns and anomalies
  - Reduces token usage by 80-90%
- **Output**: Hierarchical summary of key findings and actionable insights.

### 🧪 test-runner
**Purpose**: Execute tests without dumping output to main thread.
- **Pattern**: Run tests → Capture to log → Analyze results → Return summary
- **When to use**: After code changes or during debugging to validate functionality.
- **Capabilities**:
  - Auto-detects test frameworks
  - Captures full logs (never shown to main thread)
  - Analyzes failure patterns
- **Output**: Test results summary with specific failure analysis and fix recommendations.

### 🔀 parallel-worker
**Purpose**: Coordinate multiple parallel work streams for an issue.
- **Pattern**: Read analysis → Spawn sub-agents → Consolidate results → Return summary
- **When to use**: When implementing complex issues that can be split into independent streams.
- **Capabilities**:
  - Spawns sub-agents for specific file changes
  - Manages dependencies between streams
  - Consolidates git commits
- **Output**: Consolidated status of all parallel work streams.

### 📋 prd-architect
**Purpose**: Design and structure Product Requirements Documents.
- **Pattern**: Analyze requirements → Apply PRD template → Create hierarchical structure
- **When to use**: Creating new PRDs, reviewing structure, or ensuring best practices.
- **Capabilities**:
  - Transforms ideas into structured requirements
  - Defines success metrics and acceptance criteria
  - Ensures CCPM planning conventions
- **Output**: Well-structured PRD ready for implementation planning.

### 📐 epic-planner
**Purpose**: Plan epic execution strategy with dependency mapping.
- **Pattern**: Analyze epic → Build dependency graph → Identify critical path
- **When to use**: Before starting an epic to determine order of operations.
- **Capabilities**:
  - Identifies critical path
  - Maps technical and resource dependencies
  - Plans phases and validation gates
- **Output**: Comprehensive execution plan with phases and risk assessment.

### 🧩 task-decomposer
**Purpose**: Break down epics into ordered, dependent tasks.
- **Pattern**: Analyze requirements → Enumerate tasks → Map dependencies
- **When to use**: converting high-level epics into actionable 2-8 hour tasks.
- **Capabilities**:
  - Identifies independent vs sequential work
  - Scopes tasks appropriately
  - Generates frontmatter for tracking
- **Output**: Structured task breakdown with defined deliverables.

### 🔄 github-syncer
**Purpose**: Bidirectional synchronization between CCPM files and GitHub issues.
- **Pattern**: Read local files ↔ Call GitHub CLI ↔ Update metadata
- **When to use**: Syncing epics to GitHub, posting progress updates, or pulling comments.
- **Capabilities**:
  - Creates/Updates issues and comments
  - Enforces repository protection (prevents template pollution)
  - Maintains sync timestamps and audit trails
- **Output**: Sync report confirming data integrity between local and remote.

### 🌳 worktree-manager
**Purpose**: Manage git worktrees for parallel development.
- **Pattern**: Create worktree → Monitor status → Merge changes → Cleanup
- **When to use**: Setting up isolated environments for epic development.
- **Capabilities**:
  - Creates clean worktrees from main
  - Monitors worktree health and conflicts
  - Safely merges and cleans up resources
- **Output**: Worktree status reports and operation results.

### 🎯 parallel-orchestrator
**Purpose**: Coordinate parallel task execution across worktrees.
- **Pattern**: Parse epic → Build dependency graph → Spawn agents → Aggregate results
- **When to use**: Running an entire epic's worth of tasks in parallel where possible.
- **Capabilities**:
  - Manages task dependencies (blocks/unblocks)
  - Orchestrates up to 7 parallel agents
  - Aggregates results into epic progress
- **Output**: Execution summary with task status and next steps.

## When to Use Which Agent

| Scenario | Recommended Agent |
|----------|-------------------|
| "Found a bug, need to fix it" | `code-analyzer` |
| "Tests failed, why?" | `test-runner` |
| "Analyze these huge log files" | `file-analyzer` |
| "Plan this big feature" | `prd-architect` then `epic-planner` |
| "Break this down into tasks" | `task-decomposer` |
| "Sync my progress to GitHub" | `github-syncer` |
| "Start working on this epic" | `worktree-manager` then `parallel-orchestrator` |
| "Implement this specific task" | `parallel-worker` |

## How Agents Preserve Context

The "Context Firewall" pattern works by delegating high-token operations to sub-agents.

**Without Agents:**
1. Main thread reads 10 source files (50k tokens)
2. Main thread runs tests and reads logs (20k tokens)
3. Context fills up, model "forgets" initial instructions
4. Performance degrades, cost explodes

**With Agents:**
1. Main thread spawns `code-analyzer`
2. Agent reads 50k tokens of code in its own context
3. Agent returns 500-token summary to main thread
4. Main thread retains 99% of context window for strategic reasoning

## Creating New Agents

New agents should follow the **Specialized Executor** pattern:

1. **Single Purpose**: Do one thing extremely well (e.g., "Analyze Logs", not "Fix Bugs")
2. **Defined Input/Output**: Strict schema for what goes in and what comes out.
3. **Context Reduction**: The output must be significantly smaller than the input (aim for >90% reduction).
4. **No Roleplay**: Agents are tools, not personalities. Keep prompts functional.

### Template
```markdown
---
name: my-new-agent
description: Concise description of what this agent does
tools: Read, Write, Glob
---

<role>
You are a [Specialist Type]. Your goal is to [Specific Goal].
</role>

<core_responsibilities>
- Responsibility 1
- Responsibility 2
</core_responsibilities>

<output_format>
Structure your response as:
## Summary
## Key Findings
</output_format>
```

## Anti-Patterns to Avoid

❌ **The "God Agent"**
Creating an agent that tries to do planning, coding, and testing all at once.
*Fix: Split into `planner`, `coder`, and `tester` agents.*

❌ **Verbose Returns**
Agents that return full file contents or long logs to the main thread.
*Fix: Enforce strict summarization in the agent's system prompt.*

❌ **Chatty Agents**
Agents that ask for clarification from the main thread.
*Fix: Provide all necessary context in the initial prompt.*

❌ **Recursive Spawning**
Agents spawning agents spawning agents without limit.
*Fix: Use a flat orchestration layer (like `parallel-orchestrator`) to manage depth.*
