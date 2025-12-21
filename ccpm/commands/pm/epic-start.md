---
description: Launch parallel agents to work on epic tasks
argument-hint: <epic_name>
allowed-tools: Bash, Read, Write, Glob, Task
---

<objective>
Launch parallel agents to work on epic "$ARGUMENTS" tasks in shared branch.
</objective>

<context>
Invoke the ccpm-epic skill with action: start
Load skill: @ccpm/skills/ccpm-epic/SKILL.md
Load reference: @ccpm/skills/shared-references/agent-coordination.md
Load reference: @ccpm/skills/shared-references/worktree-operations.md
</context>

<process>
1. Load the ccpm-epic skill
2. Execute the "start" action following skill guidance
3. Create/enter branch, identify ready issues, launch agents
</process>

<success_criteria>
- Branch epic/$ARGUMENTS created or entered
- Ready issues identified with analysis files
- Parallel agents launched with proper coordination
- Execution status tracking established
</success_criteria>
