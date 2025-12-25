# Cascade Flow State Machine Reference

## Overview

The Cascade Flow operates as a 4-phase state machine with approval gates
between phases and recovery checkpoints throughout.

## Master State Machine

```
                           CASCADE FLOW STATE MACHINE

┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                              │
│  ┌───────────────┐                                                          │
│  │   INITIATE    │  User runs /cascade:start <feature>                      │
│  └───────┬───────┘                                                          │
│          │                                                                   │
│          ▼                                                                   │
│  ┌───────────────────────────────────────────────────────────────────────┐  │
│  │                     PHASE 1: BRAINSTORM                                │  │
│  │  ┌─────────────────────────────────────────────────────────────────┐  │  │
│  │  │  States:                                                         │  │  │
│  │  │  - gathering_context: Initial user input                         │  │  │
│  │  │  - perspectives_launching: 5 agents spawning                     │  │  │
│  │  │  - perspectives_running: Agents exploring                        │  │  │
│  │  │  - synthesizing: Synthesis engine aggregating                    │  │  │
│  │  │  - presenting: Showing synthesis to user                         │  │  │
│  │  │  - refining: User steering/expanding                             │  │  │
│  │  └─────────────────────────────────────────────────────────────────┘  │  │
│  │                              │                                         │  │
│  │                              ▼                                         │  │
│  │                    ┌─────────────────┐                                 │  │
│  │                    │ APPROVAL_GATE_1 │                                 │  │
│  │                    │   approve/revise│                                 │  │
│  │                    └────────┬────────┘                                 │  │
│  │                             │                                          │  │
│  │            ┌────────────────┼────────────────┐                        │  │
│  │            │                │                │                        │  │
│  │         revise           approve          abandon                     │  │
│  │            │                │                │                        │  │
│  │            ▼                │                ▼                        │  │
│  │    [back to refining]       │         [session saved]                 │  │
│  │                             │                                          │  │
│  └─────────────────────────────┼──────────────────────────────────────────┘  │
│                                ▼                                             │
│  ┌───────────────────────────────────────────────────────────────────────┐  │
│  │                     PHASE 2: PRD GENERATION                           │  │
│  │  ┌─────────────────────────────────────────────────────────────────┐  │  │
│  │  │  States:                                                         │  │  │
│  │  │  - initializing: Creating PRD structure                          │  │  │
│  │  │  - sections_writing: 4 parallel writers                          │  │  │
│  │  │  - assembling: Combining sections                                │  │  │
│  │  │  - validating: Consistency check                                 │  │  │
│  │  │  - presenting: Showing PRD to user                               │  │  │
│  │  │  - editing: User making changes                                  │  │  │
│  │  └─────────────────────────────────────────────────────────────────┘  │  │
│  │                              │                                         │  │
│  │                              ▼                                         │  │
│  │                    ┌─────────────────┐                                 │  │
│  │                    │ APPROVAL_GATE_2 │                                 │  │
│  │                    └────────┬────────┘                                 │  │
│  │                             │                                          │  │
│  └─────────────────────────────┼──────────────────────────────────────────┘  │
│                                ▼                                             │
│  ┌───────────────────────────────────────────────────────────────────────┐  │
│  │                    PHASE 3: DECOMPOSITION                             │  │
│  │  ┌─────────────────────────────────────────────────────────────────┐  │  │
│  │  │  States:                                                         │  │  │
│  │  │  - decomposing: Task decomposer running                          │  │  │
│  │  │  - visualizing: Dependency graph generation                      │  │  │
│  │  │  - simulating: Execution forecast                                │  │  │
│  │  │  - risk_ordering: Reordering for early risk                      │  │  │
│  │  │  - board_init: Task board setup                                  │  │  │
│  │  │  - presenting: Showing plan to user                              │  │  │
│  │  │  - adjusting: User modifying tasks/order                         │  │  │
│  │  └─────────────────────────────────────────────────────────────────┘  │  │
│  │                              │                                         │  │
│  │                              ▼                                         │  │
│  │                    ┌─────────────────┐                                 │  │
│  │                    │ APPROVAL_GATE_3 │                                 │  │
│  │                    └────────┬────────┘                                 │  │
│  │                             │                                          │  │
│  └─────────────────────────────┼──────────────────────────────────────────┘  │
│                                ▼                                             │
│  ┌───────────────────────────────────────────────────────────────────────┐  │
│  │                    PHASE 4: EXECUTION                                 │  │
│  │  ┌─────────────────────────────────────────────────────────────────┐  │  │
│  │  │  States:                                                         │  │  │
│  │  │  - initializing: Creating checkpoints                            │  │  │
│  │  │  - wave_N_executing: Running wave N tasks                        │  │  │
│  │  │  - wave_N_checkpointing: Saving wave N state                     │  │  │
│  │  │  - replanning: Adaptive plan adjustment                          │  │  │
│  │  │  - paused: User requested pause                                  │  │  │
│  │  │  - rolling_back: Restoring from checkpoint                       │  │  │
│  │  └─────────────────────────────────────────────────────────────────┘  │  │
│  │                              │                                         │  │
│  │                              ▼                                         │  │
│  │                    ┌─────────────────┐                                 │  │
│  │                    │   COMPLETED     │                                 │  │
│  │                    └────────┬────────┘                                 │  │
│  │                             │                                          │  │
│  └─────────────────────────────┼──────────────────────────────────────────┘  │
│                                ▼                                             │
│                      ┌─────────────────┐                                     │
│                      │  SYNC (optional)│  /cascade:sync                      │
│                      └─────────────────┘                                     │
│                                                                              │
└──────────────────────────────────────────────────────────────────────────────┘
```

## State Transitions

### Phase 1 Transitions

| From State | Event | To State |
|------------|-------|----------|
| INITIATE | /cascade:start | gathering_context |
| gathering_context | user_response | perspectives_launching |
| perspectives_launching | agents_spawned | perspectives_running |
| perspectives_running | all_complete | synthesizing |
| synthesizing | synthesis_done | presenting |
| presenting | user_steer | refining |
| presenting | user_reject | gathering_context |
| presenting | user_expand | refining |
| presenting | user_approve | PHASE_2_START |
| refining | refinement_done | presenting |

### Phase 2 Transitions

| From State | Event | To State |
|------------|-------|----------|
| PHASE_2_START | init | initializing |
| initializing | structure_created | sections_writing |
| sections_writing | all_complete | assembling |
| assembling | assembly_done | validating |
| validating | validation_pass | presenting |
| validating | validation_fail | correcting |
| correcting | corrections_done | validating |
| presenting | user_approve | PHASE_3_START |
| presenting | user_edit | editing |
| editing | edit_done | validating |

### Phase 3 Transitions

| From State | Event | To State |
|------------|-------|----------|
| PHASE_3_START | init | decomposing |
| decomposing | tasks_created | visualizing |
| visualizing | graph_done | simulating |
| simulating | forecast_done | risk_ordering |
| risk_ordering | ordering_done | board_init |
| board_init | board_ready | presenting |
| presenting | user_approve | PHASE_4_START |
| presenting | user_adjust_tasks | decomposing |
| presenting | user_adjust_order | risk_ordering |

### Phase 4 Transitions

| From State | Event | To State |
|------------|-------|----------|
| PHASE_4_START | init | initializing |
| initializing | checkpoint_created | wave_1_executing |
| wave_N_executing | wave_complete | wave_N_checkpointing |
| wave_N_checkpointing | checkpoint_done | wave_N+1_executing |
| wave_N_checkpointing | no_more_waves | COMPLETED |
| wave_N_executing | replan_triggered | replanning |
| replanning | replan_accepted | wave_N_executing |
| wave_N_executing | user_pause | paused |
| paused | user_resume | wave_N_executing |
| wave_N_executing | rollback_requested | rolling_back |
| rolling_back | rollback_done | wave_M_executing |

## Checkpoint Strategy

| Checkpoint Type | When Created | Contents |
|-----------------|--------------|----------|
| pre-execution | Before Phase 4 starts | Epic + all task files |
| wave-boundary | After each wave completes | Epic + task states + git commit |
| manual | On user request | Full state snapshot |

## Recovery Protocol

```
To resume from any state:

1. Read session.md or epic.md for current state
2. Determine phase from status field
3. Check for incomplete operations
4. Resume from last stable state

/cascade:resume <feature-name>
```

## Error States

| State | Cause | Recovery |
|-------|-------|----------|
| agent_timeout | Agent didn't respond | Retry or skip |
| validation_failed | PRD inconsistencies | Fix and revalidate |
| task_failed | Execution error | Retry, skip, or rollback |
| deadlock | Circular dependencies | Manual resolution |

## Parallel Agent Limits

| Phase | Agent Type | Max Parallel |
|-------|------------|--------------|
| 1 | Perspective agents | 5 |
| 1 | Codebase researcher | 1 (background) |
| 2 | Section writers | 4 |
| 2 | Tradeoff analyst | 1 (background) |
| 3 | Decomposer/Visualizer | 2 |
| 4 | Wave execution | 7 |
| 4 | Monitors | 2 (background) |
