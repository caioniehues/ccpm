# Refresh Epic Workflow

Update epic progress based on task states.

## Input
- `$ARGUMENTS`: Epic name

## Execution Steps

### 1. Count Task Status

Scan task files in `.claude/epics/$ARGUMENTS/`:

```bash
total=0
closed=0
open=0

for task_file in .claude/epics/$ARGUMENTS/[0-9]*.md; do
  [ -f "$task_file" ] || continue
  total=$((total + 1))

  status=$(grep '^status:' "$task_file" | cut -d: -f2 | tr -d ' ')
  case "$status" in
    closed) closed=$((closed + 1)) ;;
    *) open=$((open + 1)) ;;
  esac
done
```

### 2. Calculate Progress

```bash
if [ $total -gt 0 ]; then
  progress=$((closed * 100 / total))
else
  progress=0
fi
```

### 3. Update GitHub Task List

If epic has `github:` URL:

```bash
epic_issue=$(grep 'github:' .claude/epics/$ARGUMENTS/epic.md | grep -oE '[0-9]+$')

if [ -n "$epic_issue" ]; then
  # Get current body
  gh issue view $epic_issue --json body -q .body > /tmp/epic-body.md

  # Update checkboxes based on task status
  for task_file in .claude/epics/$ARGUMENTS/[0-9]*.md; do
    task_issue=$(grep 'github:' "$task_file" | grep -oE '[0-9]+$')
    task_status=$(grep '^status:' "$task_file" | cut -d: -f2 | tr -d ' ')

    if [ "$task_status" = "closed" ]; then
      sed -i "s/- \[ \] #$task_issue/- [x] #$task_issue/" /tmp/epic-body.md
    else
      sed -i "s/- \[x\] #$task_issue/- [ ] #$task_issue/" /tmp/epic-body.md
    fi
  done

  # Update issue
  gh issue edit $epic_issue --body-file /tmp/epic-body.md
fi
```

### 4. Determine Epic Status

```bash
if [ $progress -eq 0 ]; then
  new_status="backlog"
elif [ $progress -eq 100 ]; then
  new_status="completed"
else
  new_status="in-progress"
fi
```

### 5. Update Epic Frontmatter

```bash
current_date=$(date -u +"%Y-%m-%dT%H:%M:%SZ")

# Update status, progress, updated fields
sed -i "s/^status:.*/status: $new_status/" .claude/epics/$ARGUMENTS/epic.md
sed -i "s/^progress:.*/progress: ${progress}%/" .claude/epics/$ARGUMENTS/epic.md
sed -i "s/^updated:.*/updated: $current_date/" .claude/epics/$ARGUMENTS/epic.md
```

## Output

```
Epic refreshed: $ARGUMENTS

Tasks:
  Closed: {closed}
  Open: {open}
  Total: {total}

Progress: {old}% -> {new}%
Status: {old} -> {new}
GitHub: Task list updated

{If complete}: Run /pm:epic-close $ARGUMENTS
{If in progress}: Run /pm:next for priority tasks
```

## Notes

- Useful after manual task edits or GitHub sync
- Only modifies epic status, not task files
- Preserves all other frontmatter fields
- Updates GitHub checkboxes to match local state
