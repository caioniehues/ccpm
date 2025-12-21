---
description: Update epic progress based on task states
argument-hint: <epic_name>
allowed-tools: Bash, Read, Write, Glob
---

<objective>
Update epic "$ARGUMENTS" progress based on current task states.
</objective>

<context>
Invoke the ccpm-epic skill with action: refresh
Load skill: @ccpm/skills/ccpm-epic/SKILL.md
Load reference: @ccpm/skills/shared-references/frontmatter-operations.md
</context>

<process>
1. Load the ccpm-epic skill
2. Execute the "refresh" action following skill guidance
3. Calculate progress, update status, sync GitHub checkboxes
</process>

<success_criteria>
- Progress accurately calculated from task states
- Epic status updated (backlog/in-progress/completed)
- GitHub task list checkboxes synced
- Next action suggested based on progress
</success_criteria>
