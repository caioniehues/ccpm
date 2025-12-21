---
description: Close an issue as complete
argument-hint: <issue_number> [completion_notes]
allowed-tools: Bash, Read, Write, Glob
---

<objective>
Mark issue #$ARGUMENTS as complete and close it on GitHub.
</objective>

<context>
Invoke the ccpm-issue skill with action: close
Load skill: @ccpm/skills/ccpm-issue/SKILL.md
</context>

<process>
1. Load the ccpm-issue skill
2. Execute the "close" action following skill guidance
3. Update local status and close on GitHub
</process>

<success_criteria>
- Local task marked complete
- Issue closed on GitHub
- Epic progress updated
</success_criteria>
