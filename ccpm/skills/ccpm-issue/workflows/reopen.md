# Workflow: Reopen Issue

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

Search for task file with `github:.*issues/$ARGUMENTS` in frontmatter.
If not found: "❌ No local task for issue #$ARGUMENTS"

## Step 2: Update Local Status

Get current datetime following references/datetime-handling.md

Update task file frontmatter following references/frontmatter-operations.md:
```yaml
status: open
updated: {current_datetime}
```

## Step 3: Reset Progress

If progress file exists:
- Keep original started date
- Reset completion to previous value or 0%
- Add note about reopening with reason

Do not delete previous progress, just reset status.

## Step 4: Reopen on GitHub

```bash
# Reopen with comment
echo "🔄 Reopening issue

Reason: $ARGUMENTS

---
Reopened at: {timestamp}" | gh issue comment $ARGUMENTS --body-file -

# Reopen the issue
gh issue reopen $ARGUMENTS
```

## Step 5: Update Epic Progress

Following references/epic-integration.md, recalculate epic progress with this task now open again.

## Step 6: Output Summary

```
🔄 Reopened issue #$ARGUMENTS
  Reason: {reason_if_provided}
  Epic progress: {updated_progress}%

Start work with: /pm:issue-start $ARGUMENTS
```
</process>

<success_criteria>
Issue reopen is complete when:
- [ ] Local task file found
- [ ] Task file frontmatter updated with status: open
- [ ] Progress file reset (keeping history)
- [ ] Reopen comment posted to GitHub with reason
- [ ] Issue reopened on GitHub
- [ ] Epic progress recalculated
- [ ] User informed with summary and next steps
- [ ] All timestamps are real datetimes (not placeholders)
</success_criteria>
