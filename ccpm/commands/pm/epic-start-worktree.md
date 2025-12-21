---
description: Launch parallel agents to work on epic tasks in a shared worktree
argument-hint: <epic_name>
allowed-tools: Bash, Read, Write, LS, Task
---

<objective>
Start parallel execution of epic tasks by creating a worktree and launching coordinated agents.
</objective>

<process>
**Usage**: `/pm:epic-start <epic_name>`

**1. Quick Check**

Verify epic exists:
```bash
test -f .claude/epics/$ARGUMENTS/epic.md || echo "❌ Epic not found. Run: /pm:prd-parse $ARGUMENTS"
```

Check GitHub sync (look for `github:` field in epic frontmatter).
If missing: "❌ Epic not synced. Run: /pm:epic-sync $ARGUMENTS first"

Check for worktree:
```bash
git worktree list | grep "epic-$ARGUMENTS"
```

**2. Create or Enter Worktree**

Follow `shared-references/worktree-operations.md`:

```bash
if ! git worktree list | grep -q "epic-$ARGUMENTS"; then
  git checkout main
  git pull origin main
  git worktree add ../epic-$ARGUMENTS -b epic/$ARGUMENTS
  echo "✅ Created worktree: ../epic-$ARGUMENTS"
else
  echo "✅ Using existing worktree: ../epic-$ARGUMENTS"
fi
```

**3. Identify Ready Issues**

Read all task files in `.claude/epics/$ARGUMENTS/`:
- Parse frontmatter for `status`, `depends_on`, `parallel` fields
- Check GitHub issue status if needed
- Build dependency graph

Categorize issues:
- **Ready**: No unmet dependencies, not started
- **Blocked**: Has unmet dependencies
- **In Progress**: Already being worked on
- **Complete**: Finished

**4. Analyze Ready Issues**

For each ready issue without analysis:
```bash
if ! test -f .claude/epics/$ARGUMENTS/{issue}-analysis.md; then
  echo "Analyzing issue #{issue}..."
fi
```

**5. Launch Parallel Agents**

For each ready issue with analysis, use Task tool:
```yaml
Task:
  description: "Issue #{issue} Stream {X}"
  subagent_type: "parallel-worker"
  prompt: |
    Working in worktree: ../epic-$ARGUMENTS/
    Issue: #{issue} - {title}
    Stream: {stream_name}
    Your scope:
    - Files: {file_patterns}
    - Work: {stream_description}
    Read full requirements from:
    - .claude/epics/$ARGUMENTS/{task_file}
    - .claude/epics/$ARGUMENTS/{issue}-analysis.md
    Follow coordination rules in shared-references/agent-coordination.md
    Commit frequently with message format: "Issue #{issue}: {specific change}"
```

**6. Track Active Agents**

Create/update `.claude/epics/$ARGUMENTS/execution-status.md` with:
- Active agents and their assignments
- Queued issues waiting on dependencies
- Completed work

**7. Monitor and Coordinate**

```
Agents launched successfully!

Monitor progress: /pm:epic-status $ARGUMENTS
View worktree changes: cd ../epic-$ARGUMENTS && git status
Stop all agents: /pm:epic-stop $ARGUMENTS
Merge when complete: /pm:epic-merge $ARGUMENTS
```

**8. Handle Dependencies**

As agents complete streams:
- Check if any blocked issues are now ready
- Launch new agents for newly-ready work
- Update execution-status.md

**Output Format**:
```
🚀 Epic Execution Started: $ARGUMENTS

Worktree: ../epic-$ARGUMENTS
Branch: epic/$ARGUMENTS

Launching {total} agents across {issue_count} issues:

Issue #1234: Database Schema
  ├─ Stream A: Schema creation (Agent-1) ✓ Started
  └─ Stream B: Migrations (Agent-2) ✓ Started

Blocked Issues (2):
  - #1236: UI Components (depends on #1234)

Monitor with: /pm:epic-status $ARGUMENTS
```

**Error Handling**:
- If agent launch fails, report and offer to continue with others
- If worktree creation fails, suggest `git worktree prune`
</process>

<success_criteria>
- Worktree created or verified
- Ready issues identified and analyzed
- Parallel agents launched for ready work
- Execution status tracked
- Monitoring instructions provided
</success_criteria>
