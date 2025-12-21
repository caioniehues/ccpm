# Sync Epic Workflow

Push epic and tasks to GitHub as issues.

## Input
- `$ARGUMENTS`: Epic name

## Preflight Checks

1. **Verify epic exists:**
   ```bash
   test -f .claude/epics/$ARGUMENTS/epic.md
   ```

2. **Check for tasks:**
   ```bash
   ls .claude/epics/$ARGUMENTS/[0-9]*.md 2>/dev/null | wc -l
   ```
   If none: "No tasks. Run /pm:epic-decompose first"

3. **Check remote:**
   Ensure not syncing to template repository.

## Execution Steps

### 1. Detect Repository

```bash
remote_url=$(git remote get-url origin)
REPO=$(echo "$remote_url" | sed 's|.*github.com[:/]||' | sed 's|\.git$||')
```

### 2. Create Epic Issue

```bash
# Strip frontmatter, prepare body
sed '1,/^---$/d; 1,/^---$/d' .claude/epics/$ARGUMENTS/epic.md > /tmp/epic-body.md

# Create issue
epic_number=$(gh issue create \
  --repo "$REPO" \
  --title "Epic: $ARGUMENTS" \
  --body-file /tmp/epic-body.md \
  --label "epic,epic:$ARGUMENTS" \
  --json number -q .number)
```

### 3. Create Task Sub-Issues

Check for gh-sub-issue extension:
```bash
gh extension list | grep -q "yahsan2/gh-sub-issue"
```

For each task file:
```bash
task_name=$(grep '^name:' "$task_file" | sed 's/^name: *//')
sed '1,/^---$/d; 1,/^---$/d' "$task_file" > /tmp/task-body.md

# With sub-issues:
gh sub-issue create --parent $epic_number --title "$task_name" \
  --body-file /tmp/task-body.md --label "task,epic:$ARGUMENTS"

# Without: use gh issue create
```

**Parallel Strategy:**
- < 5 tasks: Sequential creation
- >= 5 tasks: Use github-syncer agent for parallel batches

### 4. Rename Task Files

Build mapping and rename:
```bash
# old_num:new_issue_id mapping
while IFS=: read -r task_file task_number; do
  new_name="$(dirname "$task_file")/${task_number}.md"

  # Update depends_on/conflicts_with references
  # Rename file
  mv "$task_file" "$new_name"

  # Update github field in frontmatter
done < /tmp/task-mapping.txt
```

### 5. Update Epic Issue (Fallback)

If not using sub-issues, add task list to epic body:
```markdown
## Tasks
- [ ] #123 Task Name
- [ ] #124 Task Name
```

### 6. Create Mapping File

Create `.claude/epics/$ARGUMENTS/github-mapping.md`:
```markdown
# GitHub Issue Mapping

Epic: #123 - https://github.com/{repo}/issues/123

Tasks:
- #124: Task Name - https://github.com/{repo}/issues/124
- #125: Task Name - https://github.com/{repo}/issues/125

Synced: {datetime}
```

### 7. Create Worktree

```bash
git checkout main && git pull origin main
git worktree add ../epic-$ARGUMENTS -b epic/$ARGUMENTS
```

## Output

```
Synced to GitHub
  Epic: #{epic_number} - {title}
  Tasks: {count} sub-issues created
  Labels: epic, task, epic:{name}
  Files renamed: 001.md -> {issue_id}.md
  Worktree: ../epic-$ARGUMENTS

Next:
  Start: /pm:epic-start $ARGUMENTS
  View: https://github.com/{repo}/issues/{epic_number}
```

## Error Handling

**Issue creation fails:**
- Report successes
- Note failures
- Don't rollback (partial sync is fine)

**Already synced:**
- Check for github: field in frontmatter
- Warn about duplicate issues
