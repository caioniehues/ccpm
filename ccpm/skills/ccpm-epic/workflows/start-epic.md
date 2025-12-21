# Start Epic Workflow

Launch parallel agents to work on epic tasks in a shared worktree.

## Input
- `$ARGUMENTS`: Epic name (e.g., "user-auth")

## Preflight Checks

1. **Verify epic exists:**
   ```bash
   test -f .claude/epics/$ARGUMENTS/epic.md || echo "Epic not found"
   ```

2. **Check GitHub sync:**
   - Look for `github:` in epic frontmatter
   - If missing: "Run /pm:epic-sync first"

3. **Check worktree:**
   ```bash
   git worktree list | grep "epic-$ARGUMENTS"
   ```

## Execution Steps

### 1. Create/Enter Worktree

```bash
if ! git worktree list | grep -q "epic-$ARGUMENTS"; then
  git checkout main && git pull origin main
  git worktree add ../epic-$ARGUMENTS -b epic/$ARGUMENTS
  echo "Created worktree: ../epic-$ARGUMENTS"
else
  echo "Using existing worktree: ../epic-$ARGUMENTS"
fi
```

### 2. Identify Ready Issues

Read all task files in `.claude/epics/$ARGUMENTS/`:
- Parse `status`, `depends_on`, `parallel` from frontmatter
- Build dependency graph

Categorize:
- **Ready**: No unmet dependencies, not started
- **Blocked**: Has unmet dependencies
- **In Progress**: Already being worked
- **Complete**: Finished

### 3. Analyze Ready Issues

For each ready issue without analysis file:
```bash
test -f .claude/epics/$ARGUMENTS/{issue}-analysis.md
```

If missing, analyze for parallel work streams (use task-decomposer agent if complex).

### 4. Launch Parallel Agents

For each ready issue with analysis:

```yaml
Task:
  description: "Issue #{issue} Stream {X}"
  subagent_type: "parallel-worker"
  prompt: |
    Working in worktree: ../epic-$ARGUMENTS/
    Issue: #{issue} - {title}
    Stream: {stream_name}

    Scope:
    - Files: {file_patterns}
    - Work: {stream_description}

    Read requirements from:
    - .claude/epics/$ARGUMENTS/{issue}.md
    - .claude/epics/$ARGUMENTS/{issue}-analysis.md

    Commit format: "Issue #{issue}: {change}"
```

### 5. Track Active Agents

Create `.claude/epics/$ARGUMENTS/execution-status.md`:

```markdown
---
started: {datetime}
worktree: ../epic-$ARGUMENTS
branch: epic/$ARGUMENTS
---

# Execution Status

## Active Agents
- Agent-1: Issue #123 Stream A - Started {time}
- Agent-2: Issue #123 Stream B - Started {time}

## Queued Issues
- Issue #456 - Waiting for #123

## Completed
- {None yet}
```

### 6. Handle Dependencies

As agents complete:
- Check if blocked issues are now ready
- Launch new agents for ready work
- Update execution-status.md

## Output

```
Epic Execution Started: $ARGUMENTS

Worktree: ../epic-$ARGUMENTS
Branch: epic/$ARGUMENTS

Launching {total} agents across {count} issues:

Issue #123: Database Schema
  ├─ Stream A: Schema (Agent-1) Started
  └─ Stream B: Migrations (Agent-2) Started

Blocked Issues:
  - #456: UI Components (depends on #123)

Monitor: /pm:epic-status $ARGUMENTS
```

## Error Handling

**Agent launch fails:**
- Report which failed
- Ask to continue with others

**Worktree creation fails:**
- Suggest: `git worktree prune`
- Check: `git worktree list`
