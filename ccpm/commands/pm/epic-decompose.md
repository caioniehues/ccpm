---
description: Break epic into concrete, actionable tasks
argument-hint: <epic_name>
allowed-tools: Bash, Read, Write, Glob, Task
---

<objective>
Break epic "$ARGUMENTS" into concrete, actionable tasks with proper dependency tracking.
</objective>

<context>
Invoke the ccpm-epic skill with action: decompose
Load skill: @ccpm/skills/ccpm-epic/SKILL.md
Load reference: @ccpm/skills/shared-references/datetime.md
</context>

<process>
1. Load the ccpm-epic skill
2. Execute the "decompose" action following skill guidance
3. Create task files with frontmatter, acceptance criteria, and effort estimates
</process>

<success_criteria>
- All tasks created in .claude/epics/$ARGUMENTS/
- Each task has proper frontmatter with depends_on/parallel fields
- Epic updated with Tasks Created section
- Next step suggested: /pm:epic-sync
</success_criteria>
