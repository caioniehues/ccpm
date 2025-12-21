# Workflow: Start Work on Issue

<required_reading>
**Read these reference files NOW:**
1. references/frontmatter-operations.md
2. references/datetime-handling.md
3. references/parallel-work.md
4. references/progress-tracking.md
</required_reading>

<process>
## Step 1: Preflight Checks

**Get issue details:**
```bash
gh issue view $ARGUMENTS --json state,title,labels,body
```
If it fails: "❌ Cannot access issue #$ARGUMENTS. Check number or run: gh auth login"

**Find local task file:**
- First check if `.claude/epics/*/$ARGUMENTS.md` exists (new naming)
- If not found, search for file containing `github:.*issues/$ARGUMENTS` in frontmatter (old naming)
- If not found: "❌ No local task for issue #$ARGUMENTS. This issue may have been created outside the PM system."

**Check for analysis:**
```bash
test -f .claude/epics/*/$ARGUMENTS-analysis.md || echo "❌ No analysis found for issue #$ARGUMENTS

Run: /pm:issue-analyze $ARGUMENTS first
Or: /pm:issue-start $ARGUMENTS --analyze to do both"
```
If no analysis exists and no --analyze flag, stop execution.

## Step 2: Ensure Worktree Exists

Check if epic worktree exists:
```bash
# Find epic name from task file path
epic_name={extracted_from_path}

# Check worktree
if ! git worktree list | grep -q "epic-$epic_name"; then
  echo "❌ No worktree for epic. Run: /pm:epic-start $epic_name"
  exit 1
fi
```

## Step 3: Read Analysis

Read `.claude/epics/{epic_name}/$ARGUMENTS-analysis.md`:
- Parse parallel streams
- Identify which can start immediately
- Note dependencies between streams

## Step 4: Setup Progress Tracking

Get current datetime following references/datetime-handling.md

Create workspace structure:
```bash
mkdir -p .claude/epics/{epic_name}/updates/$ARGUMENTS
```

Update task file frontmatter `updated` field with current datetime following references/frontmatter-operations.md

## Step 5: Launch Parallel Agents

For each stream that can start immediately:

Create `.claude/epics/{epic_name}/updates/$ARGUMENTS/stream-{X}.md`:
```markdown
---
issue: $ARGUMENTS
stream: {stream_name}
agent: {agent_type}
started: {current_datetime}
status: in_progress
---

# Stream {X}: {stream_name}

## Scope
{stream_description}

## Files
{file_patterns}

## Progress
- Starting implementation
```

Launch agent using Task tool:
```yaml
Task:
  description: "Issue #$ARGUMENTS Stream {X}"
  subagent_type: "{agent_type}"
  prompt: |
    You are working on Issue #$ARGUMENTS in the epic worktree.

    Worktree location: ../epic-{epic_name}/
    Your stream: {stream_name}

    Your scope:
    - Files to modify: {file_patterns}
    - Work to complete: {stream_description}

    Requirements:
    1. Read full task from: .claude/epics/{epic_name}/{task_file}
    2. Work ONLY in your assigned files
    3. Commit frequently with format: "Issue #$ARGUMENTS: {specific change}"
    4. Update progress in: .claude/epics/{epic_name}/updates/$ARGUMENTS/stream-{X}.md
    5. Follow coordination rules in /rules/agent-coordination.md

    If you need to modify files outside your scope:
    - Check if another stream owns them
    - Wait if necessary
    - Update your progress file with coordination notes

    Complete your stream's work and mark as completed when done.
```

## Step 6: GitHub Assignment

```bash
# Assign to self and mark in-progress
gh issue edit $ARGUMENTS --add-assignee @me --add-label "in-progress"
```

## Step 7: Output Summary

```
✅ Started parallel work on issue #$ARGUMENTS

Epic: {epic_name}
Worktree: ../epic-{epic_name}/

Launching {count} parallel agents:
  Stream A: {name} (Agent-1) ✓ Started
  Stream B: {name} (Agent-2) ✓ Started
  Stream C: {name} - Waiting (depends on A)

Progress tracking:
  .claude/epics/{epic_name}/updates/$ARGUMENTS/

Monitor with: /pm:epic-status {epic_name}
Sync updates: /pm:issue-sync $ARGUMENTS
```
</process>

<success_criteria>
Work is started when:
- [ ] All preflight checks passed
- [ ] Worktree exists and is accessible
- [ ] Analysis file read and parsed
- [ ] Progress tracking workspace created
- [ ] Task file frontmatter updated with current datetime
- [ ] Stream files created for parallel agents
- [ ] Agents launched via Task tool
- [ ] Issue assigned on GitHub with "in-progress" label
- [ ] User informed of launched streams and next steps
</success_criteria>
