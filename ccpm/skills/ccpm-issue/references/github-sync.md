# GitHub Sync Operations

<overview>
CCPM syncs bidirectionally with GitHub issues. Local files are the source of truth, GitHub is the communication layer.
</overview>

<sync_principles>
## Core Principles

1. **Local first**: Always update local files before GitHub
2. **Transparent audit trail**: All work syncs as GitHub comments
3. **Never force**: Respect GitHub's rate limits and permissions
4. **Validate repository**: Never sync to template repositories
</sync_principles>

<repository_protection>
## Repository Protection Check

**Always check before GitHub operations:**
```bash
remote_url=$(git remote get-url origin 2>/dev/null || echo "")

if [[ "$remote_url" == *"automazeio/ccpm"* ]]; then
  echo "❌ ERROR: Cannot sync to CCPM template repository!"
  echo "Update your remote: git remote set-url origin https://github.com/YOUR_USERNAME/YOUR_REPO.git"
  exit 1
fi
```

**Protected repositories:**
- `automazeio/ccpm` - Official template
- Any repository with `template` in the name
- Repositories you don't have write access to
</repository_protection>

<github_cli_operations>
## GitHub CLI Operations

**Authentication:**
```bash
# Check authentication status
gh auth status

# Login if needed
gh auth login
```

**View issue:**
```bash
# Basic view
gh issue view 123

# JSON output
gh issue view 123 --json state,title,body,labels,assignees

# With comments
gh issue view 123 --comments

# In browser
gh issue view 123 --web
```

**Edit issue:**
```bash
# Update title
gh issue edit 123 --title "New Title"

# Update body from file
gh issue edit 123 --body-file task.md

# Add labels
gh issue edit 123 --add-label "bug,priority-high"

# Remove labels
gh issue edit 123 --remove-label "needs-triage"

# Assign to self
gh issue edit 123 --add-assignee @me
```

**Comment on issue:**
```bash
# Add comment from string
gh issue comment 123 --body "Progress update"

# Add comment from file
gh issue comment 123 --body-file update.md

# Add comment from stdin
echo "Progress update" | gh issue comment 123 --body-file -
```

**Close/Reopen issue:**
```bash
# Close issue
gh issue close 123

# Reopen issue
gh issue reopen 123

# Close with comment
echo "Completed" | gh issue comment 123 --body-file -
gh issue close 123
```
</github_cli_operations>

<sync_patterns>
## Common Sync Patterns

**Initial sync (create GitHub issue from local):**
```bash
# Create issue from task file
gh issue create --title "Task Title" --body-file task.md --label "epic:name"

# Get issue number
issue_number=$(gh issue list --limit 1 --json number -q '.[0].number')

# Update local task file with GitHub URL
echo "github: https://github.com/$(gh repo view --json nameWithOwner -q .nameWithOwner)/issues/$issue_number" >> task.md
```

**Progress sync (local to GitHub):**
```bash
# Create update comment
cat > /tmp/update.md << EOF
## Progress Update

### Completed
- Item 1
- Item 2

### In Progress
- Item 3
EOF

# Post to GitHub
gh issue comment 123 --body-file /tmp/update.md
```

**Status sync (GitHub to local):**
```bash
# Get GitHub issue state
state=$(gh issue view 123 --json state -q .state)

# Update local task file
if [ "$state" = "CLOSED" ]; then
  sed -i 's/^status: .*/status: closed/' task.md
elif [ "$state" = "OPEN" ]; then
  sed -i 's/^status: .*/status: open/' task.md
fi
```
</sync_patterns>

<rate_limiting>
## Rate Limiting

**GitHub API limits:**
- Authenticated: 5,000 requests/hour
- Unauthenticated: 60 requests/hour

**Check rate limit:**
```bash
gh api rate_limit
```

**Handle rate limit errors:**
```bash
if ! gh issue view 123 2>/dev/null; then
  # Check if rate limited
  if gh api rate_limit | grep -q '"remaining": 0'; then
    echo "❌ GitHub rate limit exceeded"
    # Get reset time
    reset_time=$(gh api rate_limit --jq '.rate.reset')
    echo "Wait until: $(date -d @$reset_time)"
    exit 1
  fi
fi
```
</rate_limiting>

<error_handling>
## Error Handling

**Common errors and solutions:**

**Authentication error:**
```bash
if ! gh auth status &>/dev/null; then
  echo "❌ GitHub CLI not authenticated"
  echo "Run: gh auth login"
  exit 1
fi
```

**Network error:**
```bash
if ! gh issue view 123 &>/dev/null; then
  echo "❌ Cannot access GitHub"
  echo "Check internet connection"
  exit 1
fi
```

**Permission denied:**
```bash
if ! gh issue edit 123 --title "Test" &>/dev/null; then
  echo "❌ No write access to repository"
  echo "Check repository permissions"
  exit 1
fi
```

**Issue not found:**
```bash
if ! gh issue view 123 &>/dev/null; then
  echo "❌ Issue #123 not found"
  echo "Check issue number"
  exit 1
fi
```

**Issue locked:**
```bash
# Check if issue is locked
locked=$(gh issue view 123 --json isLocked -q .isLocked)

if [ "$locked" = "true" ]; then
  echo "⚠️ Issue is locked for comments"
  echo "Contact repository admin to unlock"
  exit 1
fi
```
</error_handling>

<comment_size_limits>
## Comment Size Limits

**GitHub limits:**
- Maximum comment size: 65,536 characters
- Maximum comment length: ~16,000 words

**Handle large updates:**
```bash
update_size=$(wc -c < update.md)

if [ $update_size -gt 65000 ]; then
  echo "⚠️ Update too large ($update_size chars)"
  echo "Splitting into multiple comments..."

  # Split into chunks
  split -C 60000 update.md /tmp/chunk-

  # Post each chunk
  for chunk in /tmp/chunk-*; do
    gh issue comment 123 --body-file "$chunk"
  done
else
  gh issue comment 123 --body-file update.md
fi
```
</comment_size_limits>

<best_practices>
## Best Practices

1. **Always validate repository before sync**
2. **Check authentication before operations**
3. **Handle errors gracefully with helpful messages**
4. **Respect rate limits**
5. **Keep local files as source of truth**
6. **Use `--body-file` for multi-line content**
7. **Add timestamps to sync comments**
8. **Validate issue exists before editing**
9. **Test with dry-run when possible**
10. **Log all GitHub operations for audit trail**
</best_practices>

<success_criteria>
GitHub sync is working correctly when:
- Repository protection check prevents template sync
- Authentication validated before operations
- All errors handled with actionable messages
- Rate limits respected
- Local files updated before GitHub
- Comments include timestamps
- Issue state stays in sync
- No data loss during sync
</success_criteria>
