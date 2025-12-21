---
description: Reopen a closed issue
argument-hint: <issue_number> [reason]
allowed-tools: Bash, Read, Write, Glob
---

<objective>
Reopen closed issue #$ARGUMENTS with optional reason.
</objective>

<context>
Invoke the ccpm-issue skill with action: reopen
Load skill: @ccpm/skills/ccpm-issue/SKILL.md
</context>

<process>
1. Load the ccpm-issue skill
2. Execute the "reopen" action following skill guidance
3. Reopen on GitHub and update epic progress
</process>

<success_criteria>
- Issue reopened locally and on GitHub
- History preserved
- Epic progress recalculated
</success_criteria>
