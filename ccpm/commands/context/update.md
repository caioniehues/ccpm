---
description: Update project context with recent changes
allowed-tools: Bash, Read, Write, Glob
---

<objective>
Update project context documentation in `.claude/context/` to reflect recent changes. Run at end of development sessions.
</objective>

<context>
Invoke the ccpm-context skill with action: update
Load skill: @ccpm/skills/ccpm-context/SKILL.md
</context>

<process>
1. Load the ccpm-context skill
2. Execute the "update" action following skill guidance
3. Update changed files, skip unchanged ones
</process>

<success_criteria>
- Changed context files updated with real datetime
- Unchanged files preserved
- Update summary provided to user
</success_criteria>

$ARGUMENTS
