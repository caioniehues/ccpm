---
name: cascade-decompose
description: |
  Phase 3 of Cascade Flow: Task decomposition with dependency graph visualization,
  execution simulation, and live task board initialization.
allowed-tools: Task, Read, Write, Glob, Bash, AskUserQuestion, TodoWrite
---

# Workflow: Cascade Decompose (Phase 3)

Break down the approved PRD into actionable tasks with dependency visualization,
execution simulation, and task board setup.

## Input
- Approved PRD from Phase 2
- `$ARGUMENTS`: Feature name

## Preflight Checks

### 1. Verify PRD Approved
```bash
prd_status=$(grep '^status:' ".claude/prds/${ARGUMENTS}.md" | cut -d: -f2 | tr -d ' ')
if [ "$prd_status" != "approved" ]; then
  echo "PRD not approved. Current status: $prd_status"
  exit 1
fi
```

### 2. Create Epic Directory
```bash
mkdir -p ".claude/epics/${ARGUMENTS}"/{checkpoints,waves,updates}
```

### 3. Read PRD Content
Load complete PRD from `.claude/prds/${ARGUMENTS}.md`

## Execution Steps

### Step 1: Initialize Epic File

Create `.claude/epics/${ARGUMENTS}/epic.md`:

```markdown
---
name: {feature_name}
title: {from PRD title}
status: decomposing
created: {datetime}
updated: {datetime}
prd: .claude/prds/{feature_name}.md
brainstorm_session: {from PRD}
decomposition:
  status: in_progress
  tasks_created: 0
  estimated_effort: null
visualization:
  status: pending
simulation:
  status: pending
  waves_predicted: 0
  estimated_duration: null
task_board:
  status: pending
  initialized: false
approval:
  status: pending
---

# Epic: {Title}

## Overview
{From PRD executive summary}

## Source PRD
See: .claude/prds/{feature_name}.md

## Tasks
{To be generated}

## Execution Plan
{To be generated}
```

### Step 2: Launch Task Decomposition

```yaml
Task:
  description: "Decompose into tasks"
  subagent_type: "task-decomposer"
  prompt: |
    Epic: {feature_name}

    ## Source PRD
    Read: .claude/prds/{feature_name}.md

    ## Your Task
    Break down this PRD into actionable tasks.

    For each task, create a file:
    .claude/epics/{feature_name}/{NNN}.md

    Where NNN is 001, 002, 003, etc.

    Each task file must have frontmatter:
    ---
    id: "{NNN}"
    name: "{Task Name}"
    status: open
    created: {datetime}
    effort: XS|S|M|L|XL
    depends_on: []
    conflicts_with: []
    parallel: true|false
    ---

    ## Guidelines
    - Tasks should be 2-8 hours of work
    - Include clear acceptance criteria
    - Identify dependencies accurately
    - Flag tasks that can run in parallel

    ## Output
    After creating all tasks, write summary to:
    .claude/epics/{feature_name}/decomposition-summary.md
```

### Step 3: Generate Dependency Graph

```yaml
Task:
  description: "Visualize dependencies"
  subagent_type: "dependency-visualizer"
  prompt: |
    Epic: {feature_name}

    ## Your Task
    Read all task files in: .claude/epics/{feature_name}/

    Generate ASCII dependency graph showing:
    - Task relationships
    - Parallel opportunities
    - Critical path
    - Wave groupings

    ## Output
    Write to: .claude/epics/{feature_name}/dependency-graph.md

    Include:
    - ASCII diagram
    - Critical path identification
    - Parallelization factor
    - Wave breakdown
```

### Step 4: Run Execution Simulation

```yaml
Task:
  description: "Simulate execution"
  subagent_type: "execution-simulator"
  prompt: |
    Epic: {feature_name}

    ## Your Task
    Read:
    - All task files in .claude/epics/{feature_name}/
    - .claude/epics/{feature_name}/dependency-graph.md

    Simulate wave-based execution:
    - Predict number of waves
    - Estimate duration per wave
    - Identify bottlenecks
    - Suggest optimizations

    ## Output
    Write to: .claude/epics/{feature_name}/execution-forecast.md
```

### Step 5: Apply Risk-First Ordering

Review task order and adjust:
- Move high-risk tasks earlier (fail fast)
- Ensure critical path is optimized
- Balance parallel workload

Update task dependencies if needed.

### Step 6: Initialize Task Board

Create `.claude/epics/${ARGUMENTS}/task-board.md`:

```markdown
---
type: task-board
epic: {feature_name}
created: {datetime}
last_updated: {datetime}
refresh_interval: 30s
current_wave: 0
---

# Task Board: {Feature Name}

## Summary
| Status | Count | Percentage |
|--------|-------|------------|
| Open | {N} | {%} |
| In Progress | 0 | 0% |
| Completed | 0 | 0% |
| Blocked | 0 | 0% |

## Wave Overview
{From execution-forecast.md}

## All Tasks
{Aggregated from task files}

## Dependencies
{Summary from dependency-graph.md}
```

### Step 7: Present Decomposition to User

Display:

"## Task Decomposition Complete: {feature_name}

### Tasks Created: {count}

| ID | Name | Effort | Dependencies | Parallel |
|----|------|--------|--------------|----------|
{Task table}

### Dependency Graph
```
{ASCII diagram from dependency-graph.md}
```

### Execution Forecast
- **Waves**: {count}
- **Estimated Duration**: {time}
- **Parallelization Factor**: {%}
- **Critical Path**: {task IDs}

### Bottlenecks Identified
{From simulation}

---

**Options:**
1. **Approve** - Proceed to implementation
2. **Adjust Tasks** - Modify task breakdown
3. **Adjust Order** - Change execution order
4. **Back** - Return to PRD for changes"

### Step 8: Handle User Response

#### If Approve:
- Update epic status to "approved_for_execution"
- Create initial checkpoint
- Proceed to Phase 4

#### If Adjust Tasks:
- Ask which tasks to modify
- Support: add, remove, split, merge
- Re-run dependency visualization
- Re-run simulation
- Loop back to Step 7

#### If Adjust Order:
- Allow reordering of waves
- Allow dependency modifications
- Re-run simulation
- Loop back to Step 7

#### If Back:
- Return to PRD workflow
- Allow PRD modifications

### Step 9: Create Pre-Execution Checkpoint

```bash
checkpoint_id="pre-exec-$(date +%Y%m%d%H%M%S)"
mkdir -p ".claude/epics/${ARGUMENTS}/checkpoints/${checkpoint_id}"

# Save current state
cp ".claude/epics/${ARGUMENTS}/epic.md" ".claude/epics/${ARGUMENTS}/checkpoints/${checkpoint_id}/"
cp ".claude/epics/${ARGUMENTS}"/*.md ".claude/epics/${ARGUMENTS}/checkpoints/${checkpoint_id}/" 2>/dev/null || true
```

### Step 10: Transition to Phase 4

Update epic.md:
- status: "ready_for_execution"
- approval.status: "approved"
- approval.approved_at: {datetime}

"Decomposition approved! Ready for implementation.

**Epic**: {feature_name}
**Tasks**: {count}
**Waves**: {count}
**Estimated Duration**: {time}

Proceeding to Phase 4: Execution..."

Load workflow: cascade-execute.md

## Output Summary

```
Decomposition Complete: {feature_name}

Tasks:
  Created: {count}
  Total Effort: {sum of estimates}

Execution Plan:
  Waves: {count}
  Critical Path: {duration}
  Max Parallelism: {agents}

Files Created:
  - epic.md
  - 001.md through {NNN}.md
  - dependency-graph.md
  - execution-forecast.md
  - task-board.md

Status: Approved for Execution
Checkpoint: {checkpoint_id}

Proceeding to Phase 4...
```
