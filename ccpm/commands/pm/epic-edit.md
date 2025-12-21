---
description: Edit epic details after creation
argument-hint: <epic_name>
allowed-tools: Bash, Read, Write, Glob
---

<objective>
Edit epic "$ARGUMENTS" details with interactive selection of fields to modify.
</objective>

<context>
Invoke the ccpm-epic skill with action: edit
Load skill: @ccpm/skills/ccpm-epic/SKILL.md
Load reference: @ccpm/skills/shared-references/frontmatter-operations.md
</context>

<process>
1. Load the ccpm-epic skill
2. Execute the "edit" action following skill guidance
3. Apply changes locally and optionally sync to GitHub
</process>

<success_criteria>
- Changes applied to epic file
- Frontmatter updated timestamp preserved
- GitHub issue updated if requested
</success_criteria>
