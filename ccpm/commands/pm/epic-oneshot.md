---
description: Decompose epic into tasks and sync to GitHub in one operation
argument-hint: <epic_name>
allowed-tools: Bash, Read, Write, Glob, Task
---

<objective>
Decompose epic "$ARGUMENTS" into tasks and sync to GitHub in one operation.
</objective>

<context>
Invoke the ccpm-epic skill with action: oneshot
Load skill: @ccpm/skills/ccpm-epic/SKILL.md
</context>

<process>
1. Load the ccpm-epic skill
2. Execute the "oneshot" action following skill guidance
3. Run decompose then sync as orchestrated sequence
</process>

<success_criteria>
- Tasks created via decompose action
- GitHub sync completed with issues created
- Worktree ready for development
- Next step: /pm:epic-start
</success_criteria>
