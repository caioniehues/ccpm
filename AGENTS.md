# Agents

**[中文文档 (Chinese Documentation)](doc/AGENTS_ZH.md)**

Specialized agents that do heavy work and return concise summaries to preserve context.

## Core Philosophy

> “Don't anthropomorphize subagents. Use them to organize your prompts and elide context. Subagents are best when they can do lots of work but then provide small amounts of information back to the main conversation thread.”
>
> – Adam Wolff, Anthropic

## Available Agents

### 🔍 `code-analyzer`
- **Purpose**: Hunt bugs across multiple files without polluting main context
- **Pattern**: Search many files → Analyze code → Return bug report
- **Usage**: When you need to trace logic flows, find bugs, or validate changes
- **Returns**: Concise bug report with critical findings only

### 📄 `file-analyzer`
- **Purpose**: Read and summarize verbose files (logs, outputs, configs)
- **Pattern**: Read files → Extract insights → Return summary
- **Usage**: When you need to understand log files or analyze verbose output
- **Returns**: Key findings and actionable insights (80-90% size reduction)

### 🧪 `test-runner`
- **Purpose**: Execute tests without dumping output to main thread
- **Pattern**: Run tests → Capture to log → Analyze results → Return summary
- **Usage**: When you need to run tests and understand failures
- **Returns**: Test results summary with failure analysis

### 🔀 `parallel-worker`
- **Purpose**: Coordinate multiple parallel work streams for an issue
- **Pattern**: Read analysis → Spawn sub-agents → Consolidate results → Return summary
- **Usage**: When executing parallel work streams in a worktree
- **Returns**: Consolidated status of all parallel work

### 📋 `prd-architect`
- **Purpose**: Design and structure Product Requirements Documents following CCPM conventions
- **Pattern**: Analyze requirements → Apply PRD template → Create hierarchical structure → Return structured PRD
- **Usage**: When creating new PRDs, reviewing existing PRD structure, or ensuring PRDs follow best practices
- **Returns**: Well-structured PRD with numbered requirements, acceptance criteria, and success metrics

### 📐 `epic-planner`
- **Purpose**: Plan epic execution strategy with dependency mapping and sequencing
- **Pattern**: Analyze epic → Build dependency graph → Identify critical path → Return execution plan
- **Usage**: When you need to plan epic execution, determine task ordering, or identify parallel opportunities
- **Returns**: Comprehensive execution plan with phases, dependencies, risks, and parallel streams

### 🧩 `task-decomposer`
- **Purpose**: Break down epics into ordered, dependent tasks
- **Pattern**: Analyze requirements → Enumerate tasks → Map dependencies → Return task breakdown
- **Usage**: When decomposing features into actionable work items with dependency ordering
- **Returns**: Structured task breakdown with phases, dependency graph, and critical path

### 🔄 `github-syncer`
- **Purpose**: Bidirectional synchronization between CCPM files and GitHub issues
- **Pattern**: Read local files → Create/update GitHub issues → Update local frontmatter → Return sync report
- **Usage**: When syncing epics/tasks to GitHub, posting progress updates, or importing GitHub changes
- **Returns**: Sync summary with created issues, updated files, and next steps

### 🌳 `worktree-manager`
- **Purpose**: Manage git worktrees for parallel development
- **Pattern**: Create worktree → Monitor status → Merge changes → Cleanup → Return status
- **Usage**: When setting up epic branches, coordinating parallel work, or merging completed epics
- **Returns**: Worktree status with paths, branches, commit counts, and recommended actions

### 🎯 `parallel-orchestrator`
- **Purpose**: Coordinate parallel task execution across worktrees
- **Pattern**: Parse epic → Build dependency graph → Spawn parallel agents → Aggregate results → Return summary
- **Usage**: When starting an epic with multiple parallel tasks or orchestrating multi-agent workflows
- **Returns**: Execution summary with task status, duration, changes made, and next steps

## Why Agents?

Agents are **context firewalls** that protect the main conversation from information overload:

```
Without Agent:
Main thread reads 10 files → Context explodes → Loses coherence

With Agent:
Agent reads 10 files → Main thread gets 1 summary → Context preserved
```

## How Agents Preserve Context

1. **Heavy Lifting** - Agents do the messy work (reading files, running tests, implementing features)
2. **Context Isolation** - Implementation details stay in the agent, not the main thread
3. **Concise Returns** - Only essential information returns to main conversation
4. **Parallel Execution** - Multiple agents can work simultaneously without context collision

## Example Usage

```bash
# Analyzing code for bugs
Task: "Search for memory leaks in the codebase"
Agent: code-analyzer
Returns: "Found 3 potential leaks: [concise list]"
Main thread never sees: The hundreds of files examined

# Running tests
Task: "Run authentication tests"
Agent: test-runner
Returns: "2/10 tests failed: [failure summary]"
Main thread never sees: Verbose test output and logs

# Parallel implementation
Task: "Implement issue #1234 with parallel streams"
Agent: parallel-worker
Returns: "Completed 4/4 streams, 15 files modified"
Main thread never sees: Individual implementation details

# Creating a PRD
Task: "Create PRD for user authentication feature"
Agent: prd-architect
Returns: "PRD created with 12 requirements, 5 NFRs, 8 acceptance criteria"
Main thread never sees: Template analysis and structure decisions

# Planning epic execution
Task: "Plan execution for authentication epic"
Agent: epic-planner
Returns: "3 phases, 8 tasks, critical path: auth-core → session → oauth"
Main thread never sees: Dependency graph construction details

# Syncing to GitHub
Task: "Push authentication epic to GitHub"
Agent: github-syncer
Returns: "Created epic #42, 8 task issues (#43-#50), all linked"
Main thread never sees: Individual issue creation and frontmatter updates
```

## Creating New Agents

New agents should follow these principles:

1. **Single Purpose** - Each agent has one clear job
2. **Context Reduction** - Return 10-20% of what you process
3. **No Roleplay** - Agents aren't "experts", they're task executors
4. **Clear Pattern** - Define input → processing → output pattern
5. **Error Handling** - Gracefully handle failures and report clearly

## Anti-Patterns to Avoid

❌ **Creating "specialist" agents** (database-expert, api-expert)
   Agents don't have different knowledge - they're all the same model

❌ **Returning verbose output**
   Defeats the purpose of context preservation

❌ **Making agents communicate with each other**
   Use a coordinator agent instead (like parallel-worker)

❌ **Using agents for simple tasks**
   Only use agents when context reduction is valuable

## Integration with PM System

Agents integrate seamlessly with the PM command system:

- `/pm:issue-analyze` → Identifies work streams
- `/pm:issue-start` → Spawns parallel-worker agent
- parallel-worker → Spawns multiple sub-agents
- Sub-agents → Work in parallel in the worktree
- Results → Consolidated back to main thread

This creates a hierarchy that maximizes parallelism while preserving context at every level.
