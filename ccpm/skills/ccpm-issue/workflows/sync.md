# Workflow: Sync Local Updates to GitHub

<required_reading>
**Read these reference files NOW:**
1. references/frontmatter-operations.md
2. references/datetime-handling.md
3. references/github-sync.md
4. references/progress-tracking.md
5. references/epic-integration.md
</required_reading>

<process>
## Step 1: Preflight Validation

**Repository Protection Check:**
Follow references/github-sync.md - check remote origin:
```bash
remote_url=$(git remote get-url origin 2>/dev/null || echo "")
if [[ "$remote_url" == *"automazeio/ccpm"* ]]; then
  echo "❌ ERROR: Cannot sync to CCPM template repository!"
  echo "Update your remote: git remote set-url origin https://github.com/YOUR_USERNAME/YOUR_REPO.git"
  exit 1
fi
```

**GitHub Authentication:**
- Run: `gh auth status`
- If not authenticated: "❌ GitHub CLI not authenticated. Run: gh auth login"

**Issue Validation:**
- Run: `gh issue view $ARGUMENTS --json state`
- If issue doesn't exist: "❌ Issue #$ARGUMENTS not found"
- If issue is closed and completion < 100%: warn "⚠️ Issue is closed but work incomplete"

**Local Updates Check:**
- Check if `.claude/epics/*/updates/$ARGUMENTS/` directory exists
- If not found: "❌ No local updates found for issue #$ARGUMENTS. Run: /pm:issue-start $ARGUMENTS"
- Check if progress.md exists
- If not: "❌ No progress tracking found. Initialize with: /pm:issue-start $ARGUMENTS"

**Check Last Sync:**
- Read `last_sync` from progress.md frontmatter
- If synced recently (< 5 minutes): ask "⚠️ Recently synced. Force sync anyway? (yes/no)"
- Calculate what's new since last sync

**Verify Changes:**
- Check if there are actual updates to sync
- If no changes: "ℹ️ No new updates to sync since {last_sync}"
- Exit gracefully if nothing to sync

## Step 2: Gather Local Updates

Collect all local updates for the issue:
- Read from `.claude/epics/{epic_name}/updates/$ARGUMENTS/`
- Check for new content in:
  - `progress.md` - Development progress
  - `notes.md` - Technical notes and decisions
  - `commits.md` - Recent commits and changes
  - Any other update files

## Step 3: Update Progress Tracking Frontmatter

Get current datetime following references/datetime-handling.md

Update the progress.md file frontmatter:
```yaml
---
issue: $ARGUMENTS
started: [preserve existing date]
last_sync: [Use REAL datetime from command]
completion: [calculated percentage 0-100%]
---
```

## Step 4: Determine Incremental Changes

Compare against previous sync to identify new content:
- Look for sync timestamp markers
- Identify new sections or updates
- Gather only incremental changes since last sync

Add sync markers to local files after each sync:
```markdown
<!-- SYNCED: 2024-01-15T10:30:00Z -->
```
Only sync content added after the last marker.

## Step 5: Format Update Comment

Create comprehensive update comment:

```markdown
## 🔄 Progress Update - {current_date}

### ✅ Completed Work
{list_completed_items}

### 🔄 In Progress
{current_work_items}

### 📝 Technical Notes
{key_technical_decisions}

### 📊 Acceptance Criteria Status
- ✅ {completed_criterion}
- 🔄 {in_progress_criterion}
- ⏸️ {blocked_criterion}
- □ {pending_criterion}

### 🚀 Next Steps
{planned_next_actions}

### ⚠️ Blockers
{any_current_blockers}

### 💻 Recent Commits
{commit_summaries}

---
*Progress: {completion}% | Synced from local updates at {timestamp}*
```

**Handle GitHub's Comment Limits:**
- Max comment size: 65,536 characters
- If update exceeds limit:
  1. Split into multiple comments
  2. Or summarize with link to full details
  3. Warn user: "⚠️ Update truncated due to size. Full details in local files."

## Step 6: Post to GitHub

Use GitHub CLI to add comment:
```bash
gh issue comment #$ARGUMENTS --body-file {temp_comment_file}
```

**Error Handling:**

1. **Network Error:**
   - Message: "❌ Failed to post comment: network error"
   - Solution: "Check internet connection and retry"
   - Keep local updates intact for retry

2. **Rate Limit:**
   - Message: "❌ GitHub rate limit exceeded"
   - Solution: "Wait {minutes} minutes or use different token"
   - Save comment locally for later sync

3. **Permission Denied:**
   - Message: "❌ Cannot comment on issue (permission denied)"
   - Solution: "Check repository access permissions"

4. **Issue Locked:**
   - Message: "⚠️ Issue is locked for comments"
   - Solution: "Contact repository admin to unlock"

## Step 7: Update Local Task File

Get current datetime following references/datetime-handling.md

Update the task file frontmatter with sync information:
```yaml
---
name: [Task Title]
status: open
created: [preserve existing date]
updated: [Use REAL datetime from command]
github: https://github.com/{org}/{repo}/issues/$ARGUMENTS
---
```

## Step 8: Handle Completion (If Applicable)

If task is complete, update all relevant frontmatter:

**Task file frontmatter**:
```yaml
---
name: [Task Title]
status: closed
created: [existing date]
updated: [current date/time]
github: https://github.com/{org}/{repo}/issues/$ARGUMENTS
---
```

**Progress file frontmatter**:
```yaml
---
issue: $ARGUMENTS
started: [existing date]
last_sync: [current date/time]
completion: 100%
---
```

**Completion comment format:**
```markdown
## ✅ Task Completed - {current_date}

### 🎯 All Acceptance Criteria Met
- ✅ {criterion_1}
- ✅ {criterion_2}
- ✅ {criterion_3}

### 📦 Deliverables
- {deliverable_1}
- {deliverable_2}

### 🧪 Testing
- Unit tests: ✅ Passing
- Integration tests: ✅ Passing
- Manual testing: ✅ Complete

### 📚 Documentation
- Code documentation: ✅ Updated
- README updates: ✅ Complete

This task is ready for review and can be closed.

---
*Task completed: 100% | Synced at {timestamp}*
```

**Epic progress update:**
Recalculate epic progress based on completed tasks following references/epic-integration.md

## Step 9: Post-Sync Validation

After successful sync:
- [ ] Verify comment posted on GitHub
- [ ] Confirm frontmatter updated with sync timestamp
- [ ] Check epic progress updated if task completed
- [ ] Validate no data corruption in local files

## Step 10: Output Summary

```
☁️ Synced updates to GitHub Issue #$ARGUMENTS

📝 Update summary:
   Progress items: {progress_count}
   Technical notes: {notes_count}
   Commits referenced: {commit_count}

📊 Current status:
   Task completion: {task_completion}%
   Epic progress: {epic_progress}%
   Completed criteria: {completed}/{total}

🔗 View update: gh issue view #$ARGUMENTS --comments
```
</process>

<success_criteria>
Sync is complete when:
- [ ] All preflight checks passed
- [ ] Local updates gathered and new content identified
- [ ] Progress tracking frontmatter updated with real datetime
- [ ] Update comment formatted and posted to GitHub
- [ ] Task file frontmatter updated with sync timestamp
- [ ] Epic progress recalculated if task completed
- [ ] Post-sync validation confirms success
- [ ] User informed with comprehensive summary
- [ ] No data corruption in local files
</success_criteria>
