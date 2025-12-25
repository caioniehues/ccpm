---
description: Sync local progress to GitHub as comments
argument-hint: <issue_number>
allowed-tools: Bash, Read, Write, Glob
---

<objective>
Push local development progress for issue #$ARGUMENTS to GitHub as comments for transparent audit trail.
</objective>

<context>
Load workflow: @ccpm/skills/ccpm-issue/workflows/sync-issue.md
Load reference: @ccpm/skills/shared-references/github-operations.md
Load reference: @ccpm/skills/shared-references/datetime.md
Issue: $ARGUMENTS
</context>

<process>
1. Load the sync-issue workflow
2. Read local progress and changes
3. Format update as GitHub comment
4. Post comment to issue
5. Update last_synced timestamp
</process>

<success_criteria>
- Updates posted to GitHub issue
- Last sync timestamp updated
- Epic progress updated if complete
</success_criteria>
