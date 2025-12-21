---
description: Edit issue details locally and on GitHub
argument-hint: <issue_number>
allowed-tools: Bash, Read, Write, Glob
---

<objective>
Edit issue #$ARGUMENTS details with interactive selection of fields to modify.
</objective>

<context>
Invoke the ccpm-issue skill with action: edit
Load skill: @ccpm/skills/ccpm-issue/SKILL.md
</context>

<process>
1. Load the ccpm-issue skill
2. Execute the "edit" action following skill guidance
3. Apply changes locally and sync to GitHub
</process>

<success_criteria>
- Changes applied to local file
- Changes synced to GitHub
- Frontmatter properly updated
</success_criteria>
