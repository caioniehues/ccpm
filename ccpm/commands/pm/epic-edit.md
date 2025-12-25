---
description: Edit epic details after creation
argument-hint: <epic_name>
allowed-tools: Bash, Read, Write, Glob
---

<objective>
Edit epic "$ARGUMENTS" details with interactive selection of fields to modify.
</objective>

<context>
Load workflow: @ccpm/skills/ccpm-epic/workflows/edit-epic.md
Load reference: @ccpm/skills/shared-references/frontmatter-operations.md
Epic: $ARGUMENTS
</context>

<process>
1. Load the edit-epic workflow
2. Read current epic state
3. Present edit options to user
4. Apply changes and update timestamp
5. Optionally sync to GitHub
</process>

<success_criteria>
- Changes applied to epic file
- Frontmatter updated timestamp preserved
- GitHub issue updated if requested
</success_criteria>
