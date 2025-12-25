---
description: Display live task board for Cascade Flow execution
argument-hint: <feature-name>
allowed-tools: Read, Glob, Bash
---

<objective>
Display the real-time task board for an executing Cascade Flow session.
Shows current wave, task statuses, progress, and agent activity.
</objective>

<process>
Read and display the task board:

```bash
feature="$ARGUMENTS"
board_file=".claude/epics/${feature}/task-board.md"
status_file=".claude/epics/${feature}/execution-status.md"

if [ ! -f "$board_file" ]; then
  echo "No task board found for: $feature"
  echo "Is execution running? Try /cascade:status $feature"
  exit 1
fi

# Display board
cat "$board_file"

# Add live status if executing
if [ -f "$status_file" ]; then
  echo ""
  echo "---"
  echo "## Live Execution Status"
  grep -A20 '# Execution Status' "$status_file"
fi
```
</process>

<output_format>
# Task Board: {feature_name}

## Summary
| Status | Count | Percentage |
|--------|-------|------------|
| Completed | {N} | {%} |
| In Progress | {N} | {%} |
| Ready | {N} | {%} |
| Blocked | {N} | {%} |

## Current Wave: {N}

### In Progress
┌────────────────────────────────────────────────────────────────┐
│ {id} │ {name} │ {progress_bar} {%} │ {agent} │
└────────────────────────────────────────────────────────────────┘

### Ready (Next)
- {id}: {name} (waiting for: {deps})

### Completed
- [x] {id}: {name} ({duration})

### Blocked
- {id}: {name} (blocked by: {reason})

## Agent Activity
| Agent | Task | Status | Duration |
|-------|------|--------|----------|
| {id} | {task} | {status} | {time} |

---
Last updated: {timestamp}
</output_format>

<success_criteria>
- Board displayed with current state
- All task statuses accurate
- Progress visualization clear
- Agent activity visible
</success_criteria>
