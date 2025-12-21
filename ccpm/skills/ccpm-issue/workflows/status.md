# Workflow: Check Issue Status

<required_reading>
**Read these reference files NOW:**
1. references/github-sync.md
2. references/progress-tracking.md
</required_reading>

<process>
## Step 1: Fetch Issue Status

Use GitHub CLI to get current status:
```bash
gh issue view #$ARGUMENTS --json state,title,labels,assignees,updatedAt
```

## Step 2: Display Status Overview

Show concise status information:
```
🎫 Issue #$ARGUMENTS: {Title}

📊 Status: {OPEN/CLOSED}
   Last update: {timestamp}
   Assignee: {assignee or "Unassigned"}

🏷️ Labels: {label1}, {label2}, {label3}
```

## Step 3: Check Epic Context

If issue is part of an epic:
```
📚 Epic Context:
   Epic: {epic_name}
   Epic progress: {completed_tasks}/{total_tasks} tasks complete
   This task: {task_position} of {total_tasks}
```

## Step 4: Check Local Sync Status

Compare local and GitHub state:
```
💾 Local Sync:
   Local file: {exists/missing}
   Last local update: {timestamp}
   Sync status: {in_sync/needs_sync/local_ahead/remote_ahead}
```

## Step 5: Status Indicators

Use clear visual indicators:
- 🟢 Open and ready
- 🟡 Open with blockers
- 🔴 Open and overdue
- ✅ Closed and complete
- ❌ Closed without completion

## Step 6: Suggest Next Actions

Based on status, suggest actions:
```
🚀 Suggested Actions:
   - Start work: /pm:issue-start $ARGUMENTS
   - Sync updates: /pm:issue-sync $ARGUMENTS
   - Close issue: gh issue close #$ARGUMENTS
   - Reopen issue: gh issue reopen #$ARGUMENTS
```

## Step 7: Batch Status (Optional)

If checking multiple issues (comma-separated list):
```
/pm:issue-status 123,124,125
```

Show summary table for all issues.
</process>

<success_criteria>
Status check is complete when:
- [ ] GitHub issue status fetched successfully
- [ ] Status overview displayed with clear indicators
- [ ] Epic context shown if applicable
- [ ] Local sync status compared
- [ ] Actionable next steps suggested
- [ ] Output is concise and informative
</success_criteria>
