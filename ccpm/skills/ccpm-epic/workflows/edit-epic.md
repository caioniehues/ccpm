# Edit Epic Workflow

Edit epic details after creation.

## Input
- `$ARGUMENTS`: Epic name

## Preflight Checks

1. **Verify epic exists:**
   ```bash
   test -f .claude/epics/$ARGUMENTS/epic.md
   ```

## Execution Steps

### 1. Read Current Epic

Load `.claude/epics/$ARGUMENTS/epic.md`:
- Parse frontmatter
- Read all content sections

### 2. Interactive Edit

Ask user what to edit:
- Name/Title
- Description/Overview
- Architecture decisions
- Technical approach
- Dependencies
- Success criteria
- Task list

Present current values and accept changes.

### 3. Update Epic File

Get current datetime:
```bash
date -u +"%Y-%m-%dT%H:%M:%SZ"
```

Update epic.md:
- Preserve all frontmatter except `updated`
- Apply user's edits to content
- Set `updated` to current datetime

### 4. Sync to GitHub (Optional)

If epic has `github:` URL in frontmatter:

Ask: "Update GitHub issue? (yes/no)"

If yes:
```bash
# Extract issue number
issue_num=$(grep 'github:' .claude/epics/$ARGUMENTS/epic.md | grep -oE '[0-9]+$')

# Prepare body (strip frontmatter)
sed '1,/^---$/d; 1,/^---$/d' .claude/epics/$ARGUMENTS/epic.md > /tmp/epic-body.md

# Update issue
gh issue edit $issue_num --body-file /tmp/epic-body.md
```

## Output

```
Updated epic: $ARGUMENTS
  Changes: {sections edited}

{If GitHub updated}: GitHub issue synced

View: /pm:epic-show $ARGUMENTS
```

## Guidelines

- Preserve frontmatter history (created, github URL)
- Don't change task files when editing epic
- Update `updated` timestamp on any change
- Keep title concise but descriptive
