---
description: Launch parallel agents to work on epic tasks
argument-hint: <epic_name>
allowed-tools: Bash, Read, Write, Glob, Task
---

<objective>
Launch parallel agents to work on epic "$ARGUMENTS" tasks in shared branch.
</objective>

<context>
Load workflow: @ccpm/skills/ccpm-epic/workflows/start-epic.md
Load reference: @ccpm/skills/shared-references/agent-coordination.md
Load reference: @ccpm/skills/shared-references/worktree-operations.md
Epic: $ARGUMENTS
</context>

<process>
1. Load the start-epic workflow
2. Verify epic worktree exists (create if needed)
3. Identify ready tasks (status: open, dependencies met)
4. Analyze each task for parallel work streams
5. Launch parallel agents with proper coordination
</process>

<success_criteria>
- Branch epic/$ARGUMENTS created or entered
- Ready issues identified with analysis files
- Parallel agents launched with proper coordination
- Execution status tracking established
</success_criteria>
