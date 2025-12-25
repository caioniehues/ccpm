---
description: Update epic progress based on task states
argument-hint: <epic_name>
allowed-tools: Bash, Read, Write, Glob
---

<objective>
Update epic "$ARGUMENTS" progress based on current task states.
</objective>

<context>
Load workflow: @ccpm/skills/ccpm-epic/workflows/refresh-epic.md
Load reference: @ccpm/skills/shared-references/frontmatter-operations.md
Epic: $ARGUMENTS
</context>

<process>
1. Load the refresh-epic workflow
2. Read all task files and their statuses
3. Calculate progress percentage
4. Update epic status (backlog/in-progress/completed)
5. Sync GitHub task list checkboxes
</process>

<success_criteria>
- Progress accurately calculated from task states
- Epic status updated (backlog/in-progress/completed)
- GitHub task list checkboxes synced
- Next action suggested based on progress
</success_criteria>
