---
name: cascade-execute
description: |
  Phase 4 of Cascade Flow: Wave-based parallel execution with adaptive
  replanning, checkpoint recovery, and real-time progress tracking.
allowed-tools: Task, Read, Write, Glob, Bash, TaskOutput, TodoWrite
---

# Workflow: Cascade Execute (Phase 4)

Execute tasks in dependency-aware waves with adaptive replanning,
checkpoint recovery, and real-time task board updates.

## Input
- Approved decomposition from Phase 3
- `$ARGUMENTS`: Feature name (epic name)

## Preflight Checks

### 1. Verify Epic Ready
```bash
epic_status=$(grep '^status:' ".claude/epics/${ARGUMENTS}/epic.md" | cut -d: -f2 | tr -d ' ')
if [ "$epic_status" != "ready_for_execution" ]; then
  echo "Epic not ready for execution. Status: $epic_status"
  exit 1
fi
```

### 2. Read Execution Forecast
Load `.claude/epics/${ARGUMENTS}/execution-forecast.md`:
- Number of waves
- Tasks per wave
- Dependencies

### 3. Initialize Execution Status
Create `.claude/epics/${ARGUMENTS}/execution-status.md`:

```markdown
---
epic: {feature_name}
started: {datetime}
status: initializing
current_wave: 0
total_waves: {N}
agents_active: 0
agents_max: 7
checkpoints: []
---

# Execution Status: {Feature Name}

## Current State
- Wave: Initializing
- Active Agents: 0
- Completed Tasks: 0

## Timeline
| Event | Time | Details |
|-------|------|---------|
| Started | {datetime} | Execution initialized |
```

## Execution Steps

### Step 1: Create Wave Directories

```bash
waves=$(grep 'total_waves:' ".claude/epics/${ARGUMENTS}/execution-forecast.md" | cut -d: -f2 | tr -d ' ')
for i in $(seq 1 $waves); do
  mkdir -p ".claude/epics/${ARGUMENTS}/waves/wave-$(printf '%03d' $i)"
done
```

### Step 2: Create Initial Checkpoint

```bash
checkpoint_id="wave-000-$(date +%Y%m%d%H%M%S)"
mkdir -p ".claude/epics/${ARGUMENTS}/checkpoints/${checkpoint_id}"
# Save current state
tar -czf ".claude/epics/${ARGUMENTS}/checkpoints/${checkpoint_id}/state.tar.gz" \
  ".claude/epics/${ARGUMENTS}"/*.md 2>/dev/null || true
```

### Step 3: Launch Wave Launcher Agent

```yaml
Task:
  description: "Execute: {epic_name}"
  subagent_type: "wave-launcher"
  prompt: |
    Epic: {feature_name}

    ## Your Task
    Execute all tasks in waves following the execution forecast.

    ## Files to Read
    - .claude/epics/{feature_name}/execution-forecast.md
    - .claude/epics/{feature_name}/dependency-graph.md
    - All task files: .claude/epics/{feature_name}/[0-9]*.md

    ## Execution Protocol
    1. For each wave:
       a. Create checkpoint
       b. Identify ready tasks
       c. Spawn parallel-worker agents (max 7)
       d. Monitor completion
       e. Update task-board.md
       f. Transition to next wave

    2. Handle failures:
       - Log in execution-status.md
       - Decide: retry, skip, or pause
       - Continue with non-blocked tasks

    3. Track progress:
       - Update execution-status.md after each task
       - Update task-board.md in real-time

    ## Output
    Update: .claude/epics/{feature_name}/execution-status.md
    Update: .claude/epics/{feature_name}/task-board.md
```

### Step 4: Launch Background Monitors

```yaml
# Monitor 1: Progress Dashboard
Task:
  description: "Monitor: Progress"
  run_in_background: true
  timeout: 3600000  # 1 hour
  prompt: |
    Epic: {feature_name}

    ## Your Task
    Monitor execution progress and update task board.

    Every 30 seconds:
    1. Read execution-status.md
    2. Read all task files for status
    3. Calculate progress percentage
    4. Update task-board.md with current state

    ## Output
    Continuously update: .claude/epics/{feature_name}/task-board.md

# Monitor 2: Adaptive Replanner
Task:
  description: "Monitor: Replanner"
  run_in_background: true
  timeout: 3600000
  prompt: |
    Epic: {feature_name}

    ## Your Task
    Monitor for execution deviations that require replanning.

    Triggers:
    - Task takes 2x longer than estimate
    - Task completes 50% faster than estimate
    - Task failure affects critical path
    - Unexpected blocking dependency

    On trigger:
    1. Analyze impact
    2. Generate replan suggestion
    3. Write to: .claude/epics/{feature_name}/replan-suggestion.md
```

### Step 5: User Interaction During Execution

Inform user:

"## Execution Started: {feature_name}

**Wave 1 of {N}** launching with **{count}** parallel agents...

You can:
- **Continue working** - I'll update you on progress
- **View board** - `/cascade:board {feature_name}`
- **Check status** - `/cascade:status {feature_name}`
- **Pause** - Say \"pause\" to stop after current wave"

### Step 6: Monitor Wave Progress

For each wave:

1. **Wait for completion** or user interrupt
2. **Create checkpoint** at wave boundary
3. **Check for replan suggestions**
4. **Transition to next wave** if approved

```bash
# Check wave status
wave_status=$(grep "^  wave_${current_wave}:" ".claude/epics/${ARGUMENTS}/execution-status.md")
```

### Step 7: Handle Replanning

If replan-suggestion.md exists:

"## Replanning Suggested

**Reason**: {from replan-suggestion.md}

**Current State**:
- Wave {N} of {M}
- Completed: {count}
- In Progress: {count}

**Suggestion**: {replan recommendation}

**Options**:
1. **Accept** - Apply replan and continue
2. **Ignore** - Continue with original plan
3. **Pause** - Stop for manual review"

### Step 8: Handle Failures

If task fails:

"## Task Failed: #{task_id}

**Error**: {from agent output}

**Impact**:
- Dependent tasks: {list}
- Critical path affected: Yes/No

**Options**:
1. **Retry** - Re-run the task
2. **Skip** - Mark as skipped, continue
3. **Pause** - Stop execution for manual fix
4. **Rollback** - Restore from checkpoint"

### Step 9: Wave Completion

After each wave:

```bash
# Create checkpoint
checkpoint_id="wave-${wave_num}-$(date +%Y%m%d%H%M%S)"
mkdir -p ".claude/epics/${ARGUMENTS}/checkpoints/${checkpoint_id}"
```

Update user:

"## Wave {N} Complete

**Tasks Completed**: {count}/{total}
**Duration**: {time}
**Status**: All successful / {N} failed

**Next Wave**: {count} tasks ready

Proceeding to Wave {N+1}..."

### Step 10: Execution Complete

When all waves complete:

Update epic.md:
- status: "completed"
- completed_at: {datetime}

Update execution-status.md:
- status: "completed"

Final report:

"## Execution Complete: {feature_name}

### Summary
| Metric | Value |
|--------|-------|
| Total Tasks | {count} |
| Completed | {count} |
| Failed | {count} |
| Skipped | {count} |
| Duration | {time} |
| Waves | {count} |

### Files Modified
{list from git status or task outputs}

### Next Steps
1. Review changes in codebase
2. Run tests: `/cascade:test {feature_name}`
3. Sync to GitHub: `/cascade:sync {feature_name}` (optional)

**Checkpoints available for rollback**: {list}"

## Rollback Protocol

If user requests rollback:

```bash
# List available checkpoints
ls -la ".claude/epics/${ARGUMENTS}/checkpoints/"

# Restore from checkpoint
checkpoint_id=$1
tar -xzf ".claude/epics/${ARGUMENTS}/checkpoints/${checkpoint_id}/state.tar.gz" -C .

# Reset task statuses
for task in ".claude/epics/${ARGUMENTS}"/*.md; do
  sed -i 's/status: .*/status: open/' "$task"
done
```

Inform user:
"Rolled back to checkpoint: {checkpoint_id}
Wave {N} ready to restart."

## Error Handling

### Agent Crash
- Detect via TaskOutput timeout
- Mark task as failed
- Log in execution-status.md
- Attempt retry or continue

### Dependency Deadlock
- Detect circular waiting
- Break deadlock by pausing one task
- Notify user for resolution

### Resource Exhaustion
- Monitor agent count
- Queue excess tasks
- Gradual release as slots open

## Output Summary

```
Execution Complete: {feature_name}

Results:
  Tasks: {completed}/{total}
  Duration: {time}
  Waves: {count}
  Failures: {count}

Artifacts:
  - Epic: .claude/epics/{feature_name}/epic.md
  - Status: .claude/epics/{feature_name}/execution-status.md
  - Board: .claude/epics/{feature_name}/task-board.md
  - Checkpoints: {count} available

Next: /cascade:sync {feature_name} (optional)
```
