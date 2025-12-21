# Merge Epic Workflow

Merge completed epic from worktree back to main branch.

## Input
- `$ARGUMENTS`: Epic name

## Preflight Checks

1. **Verify worktree exists:**
   ```bash
   git worktree list | grep "epic-$ARGUMENTS"
   ```

2. **Check for active agents:**
   Read `.claude/epics/$ARGUMENTS/execution-status.md`
   If active agents: "Stop agents first with /pm:epic-stop"

## Execution Steps

### 1. Pre-Merge Validation

Navigate to worktree:
```bash
cd ../epic-$ARGUMENTS

# Check uncommitted changes
if [[ $(git status --porcelain) ]]; then
  echo "Uncommitted changes:"
  git status --short
  echo "Commit or stash before merging"
  exit 1
fi

# Fetch latest
git fetch origin
git status -sb
```

### 2. Run Tests (Recommended)

Detect project type and run tests:
```bash
# Node.js
[ -f package.json ] && npm test

# Python
[ -f requirements.txt ] && pytest

# Go
[ -f go.mod ] && go test ./...

# Rust
[ -f Cargo.toml ] && cargo test
```

If tests fail: Ask "Continue anyway? (yes/no)"

### 3. Update Epic Documentation

Update `.claude/epics/$ARGUMENTS/epic.md`:
```yaml
status: completed
updated: {datetime}
completed: {datetime}
```

### 4. Merge to Main

```bash
cd {main-repo-path}

# Ensure main is current
git checkout main
git pull origin main

# Merge with --no-ff to preserve history
git merge epic/$ARGUMENTS --no-ff -m "Merge epic: $ARGUMENTS

Completed tasks:
- Task 1
- Task 2

Closes epic #{epic_issue}"
```

### 5. Handle Merge Conflicts

If conflicts detected:
```
Merge conflicts detected!

Conflicts in:
{list of files}

Options:
1. Resolve manually, then:
   git add {files}
   git commit

2. Abort:
   git merge --abort

Worktree preserved at: ../epic-$ARGUMENTS
```

### 6. Post-Merge Cleanup

On successful merge:
```bash
# Push to remote
git push origin main

# Clean up worktree
git worktree remove ../epic-$ARGUMENTS

# Delete branch
git branch -d epic/$ARGUMENTS
git push origin --delete epic/$ARGUMENTS 2>/dev/null || true

# Archive epic
mkdir -p .claude/epics/archived/
mv .claude/epics/$ARGUMENTS .claude/epics/archived/
```

### 7. Update GitHub Issues

```bash
# Extract and close epic issue
epic_issue=$(grep 'github:' .claude/epics/archived/$ARGUMENTS/epic.md | grep -oE '[0-9]+$')
gh issue close $epic_issue -c "Epic completed and merged to main"

# Close task issues
for task_file in .claude/epics/archived/$ARGUMENTS/[0-9]*.md; do
  issue_num=$(grep 'github:' "$task_file" | grep -oE '[0-9]+$')
  [ -n "$issue_num" ] && gh issue close $issue_num -c "Completed in epic merge"
done
```

## Output

```
Epic Merged: $ARGUMENTS

Summary:
  Branch: epic/$ARGUMENTS -> main
  Commits merged: {count}
  Files changed: {count}
  Issues closed: {count}

Cleanup:
  Worktree removed
  Branch deleted
  Epic archived

Next:
  Deploy if needed
  Start new epic: /pm:prd-new {feature}
```

## Error Handling

**Uncommitted changes:**
- List changes
- Suggest commit or stash

**Tests fail:**
- Show failures
- Ask for confirmation

**Merge conflicts:**
- List conflicted files
- Preserve worktree
- Provide resolution instructions
