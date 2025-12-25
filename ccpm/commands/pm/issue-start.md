---
description: Start work on issue with parallel agents
argument-hint: <issue_number>
allowed-tools: Bash, Read, Write, Task, Glob
---

<objective>
Begin work on issue #$ARGUMENTS with parallel agents based on work stream analysis.
</objective>

<context>
Load workflow: @ccpm/skills/ccpm-issue/workflows/start-issue.md
Load reference: @ccpm/skills/shared-references/agent-coordination.md
Load reference: @ccpm/skills/shared-references/worktree-operations.md
Issue: $ARGUMENTS
</context>

<process>
1. Load the start-issue workflow
2. Verify analysis exists (run analyze if not)
3. Enter epic worktree
4. Launch parallel agents per work stream
5. Initialize progress tracking
</process>

<success_criteria>
- Parallel agents launched
- Progress tracking initialized
- Issue assigned on GitHub
</success_criteria>
