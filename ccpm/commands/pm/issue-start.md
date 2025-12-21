---
description: Start work on issue with parallel agents
argument-hint: <issue_number>
allowed-tools: Bash, Read, Write, Task, Glob
---

<objective>
Begin work on issue #$ARGUMENTS with parallel agents based on work stream analysis.
</objective>

<context>
Invoke the ccpm-issue skill with action: start
Load skill: @ccpm/skills/ccpm-issue/SKILL.md
</context>

<process>
1. Load the ccpm-issue skill
2. Execute the "start" action following skill guidance
3. Launch parallel agents in epic worktree
</process>

<success_criteria>
- Parallel agents launched
- Progress tracking initialized
- Issue assigned on GitHub
</success_criteria>
