---
description: Reopen a closed issue
argument-hint: <issue_number> [reason]
allowed-tools: Bash, Read, Write, Glob
---

<objective>
Reopen closed issue #$ARGUMENTS with optional reason.
</objective>

<context>
Load workflow: @ccpm/skills/ccpm-issue/workflows/reopen-issue.md
Load reference: @ccpm/skills/shared-references/github-operations.md
Issue: $ARGUMENTS
</context>

<process>
1. Load the reopen-issue workflow
2. Update local task status to open
3. Reopen issue on GitHub with reason
4. Recalculate epic progress
5. Document reopening in history
</process>

<success_criteria>
- Issue reopened locally and on GitHub
- History preserved
- Epic progress recalculated
</success_criteria>
