# Workflow: Close Issue

<required_reading>
**Read these reference files NOW:**
1. references/frontmatter-operations.md
2. references/datetime-handling.md
3. references/github-sync.md
4. references/epic-integration.md
5. references/progress-tracking.md
</required_reading>

<process>
## Step 1: Find Local Task File

First check if `.claude/epics/*/$ARGUMENTS.md` exists (new naming).
If not found, search for task file with `github:.*issues/$ARGUMENTS` in frontmatter (old naming).
If not found: "❌ No local task for issue #$ARGUMENTS"

## Step 2: Update Local Status

Get current datetime following references/datetime-handling.md

Update task file frontmatter following references/frontmatter-operations.md:
```yaml
status: closed
updated: {current_datetime}
```

## Step 3: Update Progress File

If progress file exists at `.claude/epics/{epic}/updates/$ARGUMENTS/progress.md`:
- Set completion: 100%
- Add completion note with timestamp
- Update last_sync with current datetime

## Step 4: Close on GitHub

Add completion comment and close:
```bash
# Add final comment
echo "✅ Task completed

$ARGUMENTS

---
Closed at: {timestamp}" | gh issue comment $ARGUMENTS --body-file -

# Close the issue
gh issue close $ARGUMENTS
```

## Step 5: Update Epic Task List on GitHub

Check the task checkbox in the epic issue:

```bash
# Get epic name from local task file path
epic_name={extract_from_path}

# Get epic issue number from epic.md
epic_issue=$(grep 'github:' .claude/epics/$epic_name/epic.md | grep -oE '[0-9]+$')

if [ ! -z "$epic_issue" ]; then
  # Get current epic body
  gh issue view $epic_issue --json body -q .body > /tmp/epic-body.md

  # Check off this task
  sed -i "s/- \[ \] #$ARGUMENTS/- [x] #$ARGUMENTS/" /tmp/epic-body.md

  # Update epic issue
  gh issue edit $epic_issue --body-file /tmp/epic-body.md

  echo "✓ Updated epic progress on GitHub"
fi
```

## Step 6: Update Epic Progress

Following references/epic-integration.md:
- Count total tasks in epic
- Count closed tasks
- Calculate new progress percentage
- Update epic.md frontmatter progress field

## Step 7: Output Summary

```
✅ Closed issue #$ARGUMENTS
  Local: Task marked complete
  GitHub: Issue closed & epic updated
  Epic progress: {new_progress}% ({closed}/{total} tasks complete)

Next: Run /pm:next for next priority task
```
</process>

<success_criteria>
Issue close is complete when:
- [ ] Local task file found
- [ ] Task file frontmatter updated with status: closed
- [ ] Progress file updated with 100% completion
- [ ] Completion comment posted to GitHub
- [ ] Issue closed on GitHub
- [ ] Epic task list updated on GitHub
- [ ] Epic progress recalculated and updated
- [ ] User informed with summary
- [ ] All timestamps are real datetimes (not placeholders)
</success_criteria>
