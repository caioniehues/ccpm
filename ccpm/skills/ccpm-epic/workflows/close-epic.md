# Close Epic Workflow

Mark an epic as complete when all tasks are done.

## Input
- `$ARGUMENTS`: Epic name

## Preflight Checks

1. **Verify all tasks complete:**
   ```bash
   # Check all task files
   for task_file in .claude/epics/$ARGUMENTS/[0-9]*.md; do
     status=$(grep '^status:' "$task_file" | cut -d: -f2 | tr -d ' ')
     if [ "$status" != "closed" ]; then
       echo "Open task: $task_file"
     fi
   done
   ```

   If any open: "Cannot close. Open tasks remain: {list}"

## Execution Steps

### 1. Update Epic Status

Get current datetime:
```bash
date -u +"%Y-%m-%dT%H:%M:%SZ"
```

Update epic.md frontmatter:
```yaml
status: completed
progress: 100%
updated: {datetime}
completed: {datetime}
```

### 2. Update PRD Status

If epic references a PRD:
- Find PRD file from epic frontmatter
- Update PRD status to "complete"

### 3. Close Epic on GitHub

If epic has `github:` URL:
```bash
issue_num=$(grep 'github:' .claude/epics/$ARGUMENTS/epic.md | grep -oE '[0-9]+$')
gh issue close $issue_num --comment "Epic completed - all tasks done"
```

### 4. Archive Option

Ask: "Archive completed epic? (yes/no)"

If yes:
```bash
mkdir -p .claude/epics/.archived/
mv .claude/epics/$ARGUMENTS .claude/epics/.archived/

# Create archive summary
cat > .claude/epics/.archived/$ARGUMENTS/ARCHIVE-INFO.md << EOF
# Archive Info

Epic: $ARGUMENTS
Completed: {datetime}
Tasks completed: {count}
Duration: {days from created to completed}
EOF
```

## Output

```
Epic closed: $ARGUMENTS
  Tasks completed: {count}
  Duration: {days}

{If archived}: Archived to .claude/epics/.archived/

Next: /pm:next to see priority work
```

## Guidelines

- Only close with all tasks complete
- Preserve all data when archiving
- Update related PRD status
- GitHub issue closed with completion message
