---
description: Decompose epic into tasks and sync to GitHub in one operation
argument-hint: <epic_name>
allowed-tools: Bash, Read, Write, Glob, Task
---

<objective>
Decompose epic "$ARGUMENTS" into tasks and sync to GitHub in one operation.
</objective>

<context>
Load workflow: @ccpm/skills/ccpm-epic/workflows/oneshot-epic.md
Epic: $ARGUMENTS
</context>

<process>
1. Load the oneshot-epic workflow
2. Execute decompose workflow (create task files)
3. Execute sync workflow (push to GitHub, create worktree)
4. Report combined results
</process>

<success_criteria>
- Tasks created via decompose action
- GitHub sync completed with issues created
- Worktree ready for development
- Next step: /pm:epic-start
</success_criteria>
