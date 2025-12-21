# Workflow: Show Issue Details

<required_reading>
**Read these reference files NOW:**
1. references/github-sync.md
2. references/epic-integration.md
</required_reading>

<process>
## Step 1: Fetch Issue Data

Use `gh issue view #$ARGUMENTS` to get GitHub issue details:
```bash
gh issue view #$ARGUMENTS --json state,title,body,labels,assignees,createdAt,updatedAt,comments
```

Look for local task file:
- First check `.claude/epics/*/$ARGUMENTS.md` (new naming)
- If not found, search for file with `github:.*issues/$ARGUMENTS` in frontmatter (old naming)

Check for related issues and sub-tasks.

## Step 2: Display Issue Overview

Show issue header:
```
🎫 Issue #$ARGUMENTS: {Issue Title}
   Status: {open/closed}
   Labels: {labels}
   Assignee: {assignee}
   Created: {creation_date}
   Updated: {last_update}

📝 Description:
{issue_description}
```

## Step 3: Show Local File Mapping

If local task file exists:
```
📁 Local Files:
   Task file: .claude/epics/{epic_name}/{task_file}
   Updates: .claude/epics/{epic_name}/updates/$ARGUMENTS/
   Last local update: {timestamp}
```

## Step 4: Display Related Issues

Show relationship context:
```
🔗 Related Issues:
   Parent Epic: #{epic_issue_number}
   Dependencies: #{dep1}, #{dep2}
   Blocking: #{blocked1}, #{blocked2}
   Sub-tasks: #{sub1}, #{sub2}
```

## Step 5: Show Recent Activity

Display recent comments and updates:
```
💬 Recent Activity:
   {timestamp} - {author}: {comment_preview}
   {timestamp} - {author}: {comment_preview}

   View full thread: gh issue view #$ARGUMENTS --comments
```

## Step 6: Display Progress Tracking

If task file exists, show progress:
```
✅ Acceptance Criteria:
   ✅ Criterion 1 (completed)
   🔄 Criterion 2 (in progress)
   ⏸️ Criterion 3 (blocked)
   □ Criterion 4 (not started)
```

## Step 7: Show Quick Actions

```
🚀 Quick Actions:
   Start work: /pm:issue-start $ARGUMENTS
   Sync updates: /pm:issue-sync $ARGUMENTS
   Add comment: gh issue comment #$ARGUMENTS --body "your comment"
   View in browser: gh issue view #$ARGUMENTS --web
```

## Step 8: Error Handling

- Handle invalid issue numbers gracefully
- Check for network/authentication issues
- Provide helpful error messages and alternatives
</process>

<success_criteria>
Issue display is complete when:
- [ ] GitHub issue data fetched successfully
- [ ] Issue overview displayed with all key information
- [ ] Local file mapping shown if exists
- [ ] Related issues and dependencies listed
- [ ] Recent activity displayed
- [ ] Progress tracking shown if available
- [ ] Quick actions provided for next steps
- [ ] Errors handled gracefully with helpful messages
</success_criteria>
