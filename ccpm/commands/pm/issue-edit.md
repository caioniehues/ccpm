---
description: Edit issue details locally and on GitHub
argument-hint: <issue_number>
allowed-tools: Bash, Read, Write, Glob
---

<objective>
Edit issue #$ARGUMENTS details with interactive selection of fields to modify.
</objective>

<context>
Load workflow: @ccpm/skills/ccpm-issue/workflows/edit-issue.md
Load reference: @ccpm/skills/shared-references/frontmatter-operations.md
Load reference: @ccpm/skills/shared-references/github-operations.md
Issue: $ARGUMENTS
</context>

<process>
1. Load the edit-issue workflow
2. Read current issue state
3. Present edit options to user
4. Apply changes to local file
5. Sync changes to GitHub
</process>

<success_criteria>
- Changes applied to local file
- Changes synced to GitHub
- Frontmatter properly updated
</success_criteria>
