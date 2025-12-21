---
description: Mark epic as complete when all tasks done
argument-hint: <epic_name>
allowed-tools: Bash, Read, Write, Glob
---

<objective>
Mark epic "$ARGUMENTS" as complete after verifying all tasks are done.
</objective>

<context>
Invoke the ccpm-epic skill with action: close
Load skill: @ccpm/skills/ccpm-epic/SKILL.md
Load reference: @ccpm/skills/shared-references/github-operations.md
</context>

<process>
1. Load the ccpm-epic skill
2. Execute the "close" action following skill guidance
3. Verify tasks complete, update status, close on GitHub
</process>

<success_criteria>
- All tasks verified as closed
- Epic status set to completed with 100% progress
- GitHub epic issue closed
- Archive option offered
</success_criteria>
