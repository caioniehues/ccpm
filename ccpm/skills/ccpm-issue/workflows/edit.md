# Workflow: Edit Issue Details

<required_reading>
**Read these reference files NOW:**
1. references/frontmatter-operations.md
2. references/datetime-handling.md
3. references/github-sync.md
</required_reading>

<process>
## Step 1: Get Current Issue State

Get from GitHub:
```bash
gh issue view $ARGUMENTS --json title,body,labels
```

Find local task file:
- Search for file with `github:.*issues/$ARGUMENTS` in frontmatter
- First check `.claude/epics/*/$ARGUMENTS.md` (new naming)

If not found: "❌ No local task for issue #$ARGUMENTS"

## Step 2: Interactive Edit

Ask user what to edit:
- Title
- Description/Body
- Labels
- Acceptance criteria (local only)
- Priority/Size (local only)

Get user's selections and new values.

## Step 3: Update Local File

Get current datetime following references/datetime-handling.md

Update task file with changes:
- Update frontmatter `name` if title changed
- Update body content if description changed
- Update `updated` field with current datetime

Follow references/frontmatter-operations.md for proper frontmatter updates.

## Step 4: Update GitHub

If title changed:
```bash
gh issue edit $ARGUMENTS --title "{new_title}"
```

If body changed:
```bash
gh issue edit $ARGUMENTS --body-file {updated_task_file}
```

If labels changed:
```bash
gh issue edit $ARGUMENTS --add-label "{new_labels}"
gh issue edit $ARGUMENTS --remove-label "{removed_labels}"
```

## Step 5: Output Summary

```
✅ Updated issue #$ARGUMENTS
  Changes:
    {list_of_changes_made}

Synced to GitHub: ✅
```
</process>

<success_criteria>
Edit is complete when:
- [ ] Current issue state fetched from GitHub
- [ ] Local task file located
- [ ] User selections gathered for what to edit
- [ ] Local file updated with changes
- [ ] Frontmatter updated with current datetime
- [ ] GitHub issue updated with changes
- [ ] User informed of changes made
- [ ] All changes synced successfully
</success_criteria>
