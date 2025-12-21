# Epic Integration

<overview>
Issues exist within epics. Epic progress is calculated from task completion. All issue operations must maintain epic state.
</overview>

<epic_structure>
## Epic Directory Structure

```
.claude/epics/{epic-name}/
├── epic.md               # Epic metadata and description
├── PRD.md                # Product requirements document
├── task-001.md           # Individual task files
├── task-002.md
├── 123.md                # Tasks named by issue number (new convention)
├── 124.md
└── updates/              # Progress tracking
    ├── 123/
    │   ├── progress.md
    │   ├── notes.md
    │   └── stream-*.md
    └── 124/
        └── progress.md
```
</epic_structure>

<epic_frontmatter>
## Epic Frontmatter

```yaml
---
name: Epic Name
status: planning|in-progress|completed
created: 2024-01-15T10:30:00Z
progress: 45
prd: .claude/epics/epic-name/PRD.md
github: https://github.com/org/repo/issues/100
---
```

**Key fields:**
- `progress`: Completion percentage (0-100) calculated from tasks
- `status`: Manually set based on epic phase
- `github`: Link to epic tracking issue
</epic_frontmatter>

<progress_calculation>
## Progress Calculation

**Calculate epic progress:**
```bash
epic_dir=".claude/epics/epic-name"

# Count total tasks (exclude epic.md and PRD.md)
total_tasks=$(find "$epic_dir" -maxdepth 1 -name "*.md" ! -name "epic.md" ! -name "PRD.md" | wc -l)

# Count closed tasks
closed_tasks=$(grep -l "^status: closed" "$epic_dir"/*.md 2>/dev/null | grep -v epic.md | grep -v PRD.md | wc -l)

# Calculate progress percentage
if [ $total_tasks -gt 0 ]; then
  progress=$((closed_tasks * 100 / total_tasks))
else
  progress=0
fi

echo "Epic progress: $progress% ($closed_tasks/$total_tasks tasks)"
```

**Update epic frontmatter:**
```bash
datetime=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
sed -i "s/^progress: .*/progress: $progress/" "$epic_dir/epic.md"
sed -i "s/^updated: .*/updated: $datetime/" "$epic_dir/epic.md"
```
</progress_calculation>

<task_to_epic_link>
## Finding Epic from Task

**Extract epic name from task file path:**
```bash
task_file=".claude/epics/epic-name/123.md"

# Get epic directory
epic_dir=$(dirname "$task_file")

# Get epic name
epic_name=$(basename "$epic_dir")

echo "Epic: $epic_name"
```

**Find epic.md from task:**
```bash
task_file=".claude/epics/epic-name/123.md"
epic_file="$(dirname "$task_file")/epic.md"

if [ -f "$epic_file" ]; then
  echo "Epic file found: $epic_file"
else
  echo "❌ Epic file not found"
fi
```
</task_to_epic_link>

<epic_github_sync>
## Epic GitHub Sync

**Get epic issue number:**
```bash
epic_file=".claude/epics/epic-name/epic.md"

# Extract issue number from GitHub URL
epic_issue=$(grep '^github:' "$epic_file" | grep -oE '[0-9]+$')

if [ -z "$epic_issue" ]; then
  echo "⚠️ No GitHub issue linked for epic"
else
  echo "Epic issue: #$epic_issue"
fi
```

**Update epic task list on GitHub:**
```bash
epic_issue=$(grep '^github:' .claude/epics/epic-name/epic.md | grep -oE '[0-9]+$')
task_issue=123

if [ ! -z "$epic_issue" ]; then
  # Get current epic body
  gh issue view $epic_issue --json body -q .body > /tmp/epic-body.md

  # Check off completed task
  sed -i "s/- \[ \] #$task_issue/- [x] #$task_issue/" /tmp/epic-body.md

  # Update epic issue
  gh issue edit $epic_issue --body-file /tmp/epic-body.md

  echo "✓ Updated epic #$epic_issue task list"
fi
```

**Uncheck task (when reopened):**
```bash
# Uncheck task
sed -i "s/- \[x\] #$task_issue/- [ ] #$task_issue/" /tmp/epic-body.md

# Update epic issue
gh issue edit $epic_issue --body-file /tmp/epic-body.md
```
</epic_github_sync>

<epic_status_transitions>
## Epic Status Transitions

**Automatic transitions based on progress:**
```bash
progress=45  # From calculation

# Update epic status based on progress
if [ $progress -eq 0 ]; then
  new_status="planning"
elif [ $progress -eq 100 ]; then
  new_status="completed"
else
  new_status="in-progress"
fi

# Update epic frontmatter
sed -i "s/^status: .*/status: $new_status/" epic.md
```

**Manual status override:**
Epic status can be manually set to:
- `planning`: Epic in planning phase (even if some tasks exist)
- `in-progress`: Active development
- `completed`: All tasks done and epic delivered
- `on-hold`: Epic paused
- `cancelled`: Epic cancelled
</epic_status_transitions>

<operations_trigger_epic_update>
## Operations That Trigger Epic Update

**Issue close:**
1. Mark task as closed
2. Recalculate epic progress
3. Update epic frontmatter
4. Check off task in epic GitHub issue

**Issue reopen:**
1. Mark task as open
2. Recalculate epic progress
3. Update epic frontmatter
4. Uncheck task in epic GitHub issue

**Issue created:**
1. Add task to epic directory
2. Recalculate epic progress
3. Update epic frontmatter
4. Add task to epic GitHub issue body

**Issue deleted:**
1. Remove task from epic directory
2. Recalculate epic progress
3. Update epic frontmatter
4. Remove task from epic GitHub issue body
</operations_trigger_epic_update>

<epic_worktree>
## Epic Worktree Integration

**Check if epic has worktree:**
```bash
epic_name="epic-name"

if git worktree list | grep -q "epic-$epic_name"; then
  echo "✓ Worktree exists for epic"
  worktree_path=$(git worktree list | grep "epic-$epic_name" | awk '{print $1}')
  echo "Path: $worktree_path"
else
  echo "❌ No worktree for epic"
  echo "Run: /pm:epic-start $epic_name"
fi
```

**Worktree location:**
```
../epic-{epic-name}/  # Parallel to main working directory
```
</epic_worktree>

<best_practices>
## Best Practices

1. **Always recalculate progress** - Don't manually set progress percentages
2. **Update epic after task changes** - Keep epic state in sync
3. **Sync to GitHub epic issue** - Update task list checkboxes
4. **Preserve epic metadata** - Don't overwrite created, github fields
5. **Use real datetimes** - Update `updated` field with current time
6. **Handle missing epic gracefully** - Some tasks may not have epic
7. **Validate epic file exists** - Check before updating
8. **Use atomic operations** - Update local first, then GitHub
</best_practices>

<success_criteria>
Epic integration is working correctly when:
- Epic progress recalculates after task status changes
- Epic frontmatter stays valid after updates
- Epic GitHub issue task list stays in sync
- Task-to-epic relationship maintained
- Worktree integration works correctly
- All timestamps are real (not placeholders)
- Epic status reflects actual progress
- No orphaned tasks (tasks without epic)
</success_criteria>
