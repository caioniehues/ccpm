---
description: Full bidirectional sync between local and GitHub
argument-hint: [epic_name]
allowed-tools: Bash, Read, Write, LS
---

<objective>
Perform bidirectional synchronization between local task files and GitHub issues.
</objective>

<process>
**Usage**: `/pm:sync [epic_name]`

If epic_name provided, sync only that epic. Otherwise sync all.

**1. Pull from GitHub**

Get current state of all issues:
```bash
gh issue list --label "epic" --limit 1000 --json number,title,state,body,labels,updatedAt
gh issue list --label "task" --limit 1000 --json number,title,state,body,labels,updatedAt
```

**2. Update Local from GitHub**

For each GitHub issue:
- Find corresponding local file by issue number
- Compare states:
  - If GitHub state newer (updatedAt > local updated), update local
  - If GitHub closed but local open, close local
  - If GitHub reopened but local closed, reopen local
- Update frontmatter to match GitHub state

**3. Push Local to GitHub**

For each local task/epic:
- If has GitHub URL but GitHub issue not found, it was deleted - mark local as archived
- If no GitHub URL, create new issue (like epic-sync)
- If local updated > GitHub updatedAt, push changes:
  ```bash
  gh issue edit {number} --body-file {local_file}
  ```

**4. Handle Conflicts**

If both changed (local and GitHub updated since last sync):
- Show both versions
- Ask user: "Local and GitHub both changed. Keep: (local/github/merge)?"
- Apply user's choice

**5. Update Sync Timestamps**

Update all synced files with last_sync timestamp.

**6. Output**

```
🔄 Sync Complete

Pulled from GitHub:
  Updated: {count} files
  Closed: {count} issues

Pushed to GitHub:
  Updated: {count} issues
  Created: {count} new issues

Conflicts resolved: {count}

Status:
  ✅ All files synced
  {or list any sync failures}
```
</process>

<success_criteria>
- GitHub issues pulled and local files updated
- Local changes pushed to GitHub
- Conflicts detected and resolved with user input
- Sync timestamps updated on all files
</success_criteria>
