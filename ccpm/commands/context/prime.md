---
description: Load project context for a new session
allowed-tools: Bash, Read, Glob
---

<objective>
Load essential project context from `.claude/context/` to establish understanding of the codebase at session start.
</objective>

<context>
Invoke the ccpm-context skill with action: prime
Load skill: @ccpm/skills/ccpm-context/SKILL.md
</context>

<process>
1. Load the ccpm-context skill
2. Execute the "prime" action following skill guidance
3. Load context files in priority order
</process>

<success_criteria>
- Context files loaded successfully
- Project understanding established
- Ready state confirmed to user
</success_criteria>
