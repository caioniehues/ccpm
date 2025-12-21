# Oneshot Epic Workflow

Decompose epic into tasks and sync to GitHub in one operation.

## Input
- `$ARGUMENTS`: Epic name

## Preflight Checks

1. **Verify epic exists:**
   ```bash
   test -f .claude/epics/$ARGUMENTS/epic.md
   ```

2. **Check for existing tasks:**
   ```bash
   if ls .claude/epics/$ARGUMENTS/[0-9]*.md 2>/dev/null | grep -q .; then
     echo "Tasks already exist. This will create duplicates."
     exit 1
   fi
   ```

3. **Check if already synced:**
   ```bash
   if grep -q "github:" .claude/epics/$ARGUMENTS/epic.md; then
     echo "Already synced. Use /pm:epic-sync to update."
     exit 1
   fi
   ```

## Execution Steps

### 1. Execute Decompose

Run decompose workflow:
```
Running: decompose-epic.md $ARGUMENTS
```

This will:
- Read the epic
- Create task files (parallel agents if appropriate)
- Update epic with task summary

### 2. Execute Sync

Immediately follow with sync:
```
Running: sync-epic.md $ARGUMENTS
```

This will:
- Create epic issue on GitHub
- Create sub-issues for tasks
- Rename task files to issue IDs
- Create worktree

## Output

```
Epic Oneshot Complete: $ARGUMENTS

Step 1: Decomposition
  Tasks created: {count}

Step 2: GitHub Sync
  Epic: #{number}
  Sub-issues: {count}
  Worktree: ../epic-$ARGUMENTS

Ready for development!
  Start: /pm:epic-start $ARGUMENTS
  Single task: /pm:issue-start {number}
```

## Notes

This is a convenience wrapper that runs:
1. `decompose-epic.md`
2. `sync-epic.md`

Both handle their own validation and parallel execution.

Use when epic is ready and you want to go from epic to GitHub in one step.
