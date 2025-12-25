---
description: Sync completed Cascade Flow to GitHub (optional)
argument-hint: <feature-name>
allowed-tools: Bash, Read, Write, Glob
---

<objective>
Sync a completed Cascade Flow session to GitHub. Creates issues for the epic
and tasks, links them appropriately, and updates local files with GitHub references.
</objective>

<process>
## Preflight Checks

```bash
feature="$ARGUMENTS"

# Verify epic exists and is complete
if [ ! -f ".claude/epics/${feature}/epic.md" ]; then
  echo "No epic found for: $feature"
  exit 1
fi

status=$(grep '^status:' ".claude/epics/${feature}/epic.md" | cut -d: -f2 | tr -d ' ')
if [ "$status" != "completed" ]; then
  echo "Epic not complete. Status: $status"
  echo "Complete execution first with /cascade:start $feature"
  exit 1
fi

# Verify GitHub CLI
if ! command -v gh &> /dev/null; then
  echo "GitHub CLI (gh) not found. Install with: brew install gh"
  exit 1
fi

# Verify authentication
if ! gh auth status &> /dev/null; then
  echo "Not authenticated to GitHub. Run: gh auth login"
  exit 1
fi
```

## Sync Process

1. **Create Epic Issue**
```bash
epic_title=$(grep '^title:' ".claude/epics/${feature}/epic.md" | cut -d: -f2- | xargs)
epic_body=$(cat ".claude/epics/${feature}/epic.md")

epic_issue=$(gh issue create \
  --title "Epic: $epic_title" \
  --body "$epic_body" \
  --label "epic" \
  2>&1)

epic_number=$(echo "$epic_issue" | grep -oE '[0-9]+$')
```

2. **Create Task Issues**
```bash
for task_file in ".claude/epics/${feature}"/[0-9]*.md; do
  task_id=$(basename "$task_file" .md)
  task_name=$(grep '^name:' "$task_file" | cut -d: -f2- | xargs)
  task_body=$(cat "$task_file")

  task_issue=$(gh issue create \
    --title "Task: $task_name" \
    --body "$task_body" \
    --label "task" \
    --label "epic:${feature}" \
    2>&1)

  task_number=$(echo "$task_issue" | grep -oE '[0-9]+$')

  # Update local file with GitHub reference
  sed -i "s/^github:.*/github: $task_number/" "$task_file"
done
```

3. **Update Epic with Task Links**
```bash
# Add task list to epic issue
task_list=""
for task_file in ".claude/epics/${feature}"/[0-9]*.md; do
  gh_num=$(grep '^github:' "$task_file" | cut -d: -f2 | tr -d ' ')
  task_name=$(grep '^name:' "$task_file" | cut -d: -f2- | xargs)
  status=$(grep '^status:' "$task_file" | cut -d: -f2 | tr -d ' ')
  if [ "$status" = "closed" ]; then
    task_list="$task_list\n- [x] #$gh_num $task_name"
  else
    task_list="$task_list\n- [ ] #$gh_num $task_name"
  fi
done

gh issue edit "$epic_number" --body "$(cat ".claude/epics/${feature}/epic.md")\n\n## Tasks\n$task_list"
```

4. **Create GitHub Mapping**
```bash
cat > ".claude/epics/${feature}/github-mapping.md" << EOF
---
epic_issue: $epic_number
created: $(date -u +"%Y-%m-%dT%H:%M:%SZ")
---

# GitHub Mapping: $feature

## Epic
- Issue: #$epic_number

## Tasks
$(for task_file in ".claude/epics/${feature}"/[0-9]*.md; do
  task_id=$(basename "$task_file" .md)
  gh_num=$(grep '^github:' "$task_file" | cut -d: -f2 | tr -d ' ')
  echo "- $task_id → #$gh_num"
done)
EOF
```
</process>

<output_format>
## GitHub Sync Complete: {feature_name}

### Epic Issue
- #{epic_number}: {title}
- URL: {github_url}

### Task Issues Created
| Local ID | GitHub Issue | Title |
|----------|--------------|-------|
| 001 | #{N} | {name} |
| 002 | #{N} | {name} |

### Labels Applied
- `epic` on epic issue
- `task` on all task issues
- `epic:{feature_name}` on all issues

### Files Updated
- .claude/epics/{feature_name}/github-mapping.md (created)
- All task files (github field updated)

### Next Steps
- View epic: gh issue view {epic_number}
- View project board: {url if applicable}
</output_format>

<success_criteria>
- Epic issue created with correct labels
- All task issues created and linked
- Local files updated with GitHub references
- Mapping file created for reference
</success_criteria>
