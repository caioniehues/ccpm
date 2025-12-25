---
description: Break epic into concrete, actionable tasks
argument-hint: <epic_name>
allowed-tools: Bash, Read, Write, Glob, Task
---

<objective>
Break epic "$ARGUMENTS" into concrete, actionable tasks with proper dependency tracking.
</objective>

<context>
Load workflow: @ccpm/skills/ccpm-epic/workflows/decompose-epic.md
Load reference: @ccpm/skills/shared-references/datetime.md
Epic: $ARGUMENTS
</context>

<process>
1. Load the decompose-epic workflow
2. Verify epic exists and read current state
3. Analyze for parallel task creation strategy
4. Create task files with frontmatter, acceptance criteria, and effort estimates
5. Update epic with Tasks Created section
</process>

<success_criteria>
- All tasks created in .claude/epics/$ARGUMENTS/
- Each task has proper frontmatter with depends_on/parallel fields
- Epic updated with Tasks Created section
- Next step suggested: /pm:epic-sync
</success_criteria>
