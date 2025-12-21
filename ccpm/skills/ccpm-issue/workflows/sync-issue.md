---
allowed-tools: Bash, Read, Write, LS
---

# Sync Issue Progress

Push local updates as GitHub issue comments for transparent audit trail.

## Usage
```
/issue:sync <issue_number>
```

## Preflight Checklist

Before proceeding, complete these validation steps silently.

0. **Repository Protection Check:**
   ```bash
   remote_url=$(git remote get-url origin 2>/dev/null || echo "")
   if [[ "$remote_url" == *"automazeio/ccpm"* ]]; then
     echo "❌ ERROR: Cannot sync to CCPM template repository!"
     echo "Update your remote: git remote set-url origin https://github.com/YOUR_USERNAME/YOUR_REPO.git"
     exit 1
   fi
   ```

1. **GitHub Authentication:**
   - Run: `gh auth status`
   - If not authenticated: "❌ GitHub CLI not authenticated. Run: gh auth login"

2. **Issue Validation:**
   - Run: `gh issue view $ARGUMENTS --json state`
   - If issue doesn't exist: "❌ Issue #$ARGUMENTS not found"

3. **Local Updates Check:**
   - Check if `.claude/epics/*/updates/$ARGUMENTS/` directory exists
   - If not found: "❌ No local updates found for issue #$ARGUMENTS. Run: /issue:start $ARGUMENTS"

4. **Check Last Sync:**
   - Read `last_sync` from progress.md frontmatter
   - If synced recently (< 5 minutes): "⚠️ Recently synced. Force sync anyway? (yes/no)"

## Instructions

### 1. Gather Local Updates

Collect all local updates for the issue:
- Read from `.claude/epics/{epic_name}/updates/$ARGUMENTS/`
- Check for new content in:
  - `progress.md` - Development progress
  - `notes.md` - Technical notes and decisions
  - `commits.md` - Recent commits and changes
  - Any stream files

### 2. Update Progress Tracking

Get current datetime: `date -u +"%Y-%m-%dT%H:%M:%SZ"`

Update the progress.md file frontmatter:
```yaml
---
issue: $ARGUMENTS
started: [preserve existing date]
last_sync: [current datetime]
completion: [calculated percentage 0-100%]
---
```

### 3. Format Update Comment

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

### 4. Post to GitHub

```bash
gh issue comment $ARGUMENTS --body-file {temp_comment_file}
```

### 5. Handle Completion

If task is complete (100%), add completion comment:

```markdown
## ✅ Task Completed - {current_date}

### 🎯 All Acceptance Criteria Met
- ✅ {criterion_1}
- ✅ {criterion_2}

### 📦 Deliverables
- {deliverable_1}
- {deliverable_2}

### 🧪 Testing
- Unit tests: ✅ Passing
- Integration tests: ✅ Passing

This task is ready for review and can be closed.

---
*Task completed: 100% | Synced at {timestamp}*
```

### 6. Output Summary

```
☁️ Synced updates to GitHub Issue #$ARGUMENTS

📝 Update summary:
   Progress items: {progress_count}
   Technical notes: {notes_count}
   Commits referenced: {commit_count}

📊 Current status:
   Task completion: {task_completion}%
   Epic progress: {epic_progress}%

🔗 View update: gh issue view $ARGUMENTS --comments
```

## Error Handling

- **Network Error**: Keep local updates, retry later
- **Rate Limit**: "❌ GitHub rate limit exceeded. Wait {minutes} minutes."
- **Permission Denied**: "❌ Cannot comment on issue (permission denied)"
- **Issue Locked**: "⚠️ Issue is locked for comments"

## Important Notes

- Incremental sync only - don't duplicate content
- Add sync markers to prevent re-posting
- Max comment size: 65,536 characters (split if needed)

$ARGUMENTS
