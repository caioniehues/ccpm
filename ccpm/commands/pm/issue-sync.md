---
description: Sync local progress to GitHub as comments
argument-hint: <issue_number>
allowed-tools: Bash, Read, Write, Glob
---

<objective>
Push local development progress for issue #$ARGUMENTS to GitHub as comments for transparent audit trail.
</objective>

<context>
Invoke the ccpm-issue skill with action: sync
Load skill: @ccpm/skills/ccpm-issue/SKILL.md
Load reference: @ccpm/skills/shared-references/github-operations.md
</context>

<process>
1. Load the ccpm-issue skill
2. Execute the "sync" action following skill guidance
3. Post update comment to GitHub
</process>

<success_criteria>
- Updates posted to GitHub issue
- Last sync timestamp updated
- Epic progress updated if complete
</success_criteria>
