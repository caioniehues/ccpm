---
description: Create initial project context documentation
allowed-tools: Bash, Read, Write, Glob
---

<objective>
Create comprehensive project context documentation in `.claude/context/` by analyzing the current codebase and establishing baseline documentation for future sessions.
</objective>

<context>
Invoke the ccpm-context skill with action: create
Load skill: @ccpm/skills/ccpm-context/SKILL.md
</context>

<process>
1. Load the ccpm-context skill
2. Execute the "create" action following skill guidance
3. Generate all 9 context files with proper frontmatter
</process>

<success_criteria>
- All context files created in .claude/context/
- Each file has valid frontmatter with real datetime
- Summary provided to user
</success_criteria>
