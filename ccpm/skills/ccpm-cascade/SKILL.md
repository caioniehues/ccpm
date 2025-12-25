---
name: ccpm-cascade
description: |
  Revolutionary 4-phase offline workflow for AI-assisted development.
  Transforms sequential interactions into multi-dimensional, parallel-first architecture
  with background agent incubation, multi-perspective ideation, real-time synthesis,
  and adaptive execution.
---

<objective>
Orchestrate the CASCADE FLOW workflow through four phases:
1. BRAINSTORM - Multi-specialist parallel ideation with background research
2. PRD GENERATION - Parallel section writing with consistency validation
3. DECOMPOSITION - DAG visualization, execution simulation, live task board
4. IMPLEMENTATION - Wave-based execution with adaptive replanning and checkpoints
</objective>

<essential_principles>

## Offline-First Architecture
All operations work without GitHub. Sync is optional at the end via `/cascade:sync`.

## Parallel-First Design
- Phase 1: 5 perspective agents explore simultaneously
- Phase 2: 4 section writers work in parallel
- Phase 3: Visualizer and simulator run concurrently
- Phase 4: Up to 7 task agents per wave

## Background Incubation
Codebase researcher runs continuously while user interacts.
Ideas continue developing in background during review phases.

## Approval Gates
Explicit user approval required between phases.
Supports revision loops at each gate.

## Phase Recovery
Sessions can be resumed from any checkpoint.
State persisted in frontmatter for crash recovery.

## Context Firewall
Agents return only synthesized summaries.
Implementation details stay within agent context.

</essential_principles>

<directory_structure>
```
.claude/
├── brainstorm/{session-id}/           # Phase 1
│   ├── session.md                     # Master session state
│   ├── perspectives/                  # 5 specialist outputs
│   │   ├── user-advocate.md
│   │   ├── tech-skeptic.md
│   │   ├── innovation-catalyst.md
│   │   ├── risk-assessor.md
│   │   └── simplicity-champion.md
│   ├── research/                      # Background findings
│   │   ├── patterns.md
│   │   ├── relevant-files.md
│   │   └── prior-art.md
│   ├── synthesis/                     # Aggregated insights
│   │   ├── themes.md
│   │   ├── conflicts.md
│   │   └── recommendations.md
│   └── interactions/                  # User steering history
│       └── {timestamp}-{type}.md
├── prds/{name}.md                     # Phase 2 output
├── epics/{name}/                      # Phases 3-4
│   ├── epic.md
│   ├── dependency-graph.md            # ASCII DAG
│   ├── execution-forecast.md          # Predicted timeline
│   ├── task-board.md                  # Real-time status
│   ├── {NNN}.md                       # Task files
│   ├── checkpoints/                   # Rollback points
│   └── waves/                         # Wave execution
└── cascade-flow/                      # Workflow metadata
    ├── active-sessions.md
    └── recovery/
```
</directory_structure>

<intake>
What would you like to do?

1. **Start new feature** - Begin brainstorming a new feature with multi-specialist ideation
2. **Resume session** - Continue from a previous checkpoint
3. **Check status** - View current cascade session status
4. **View board** - Display live task board for active execution
5. **Sync to GitHub** - Push completed work to GitHub (optional)
</intake>

<routing>
| User Intent | Action |
|-------------|--------|
| Start new / brainstorm / ideate | Load workflow: cascade-brainstorm.md |
| Resume / continue | Load workflow: cascade-resume.md |
| Status / progress | Load workflow: cascade-status.md |
| Board / tasks / live | Load workflow: cascade-board.md |
| Sync / push / github | Load workflow: cascade-sync.md |
| PRD / requirements (after brainstorm) | Load workflow: cascade-prd.md |
| Decompose / tasks (after PRD) | Load workflow: cascade-decompose.md |
| Execute / implement (after decompose) | Load workflow: cascade-execute.md |
</routing>

<state_machine>
```
INITIATE
    │
    ▼
PHASE_1_BRAINSTORM ──────────────────────────────────────┐
    │ spawn: 5 perspective agents (parallel)             │
    │ spawn: codebase-researcher (background)            │
    │ run: synthesis-engine                              │
    │                                                    │
    ▼                                                    │
APPROVAL_GATE_1 ◄────────── revise ──────────────────────┤
    │                                                    │
    │ approve                                            │
    ▼                                                    │
PHASE_2_PRD ─────────────────────────────────────────────┤
    │ spawn: 4 section-writer agents (parallel)          │
    │ run: consistency-validator                         │
    │ spawn: tradeoff-analyst (background)               │
    │                                                    │
    ▼                                                    │
APPROVAL_GATE_2 ◄────────── revise ──────────────────────┤
    │                                                    │
    │ approve                                            │
    ▼                                                    │
PHASE_3_DECOMPOSE ───────────────────────────────────────┤
    │ run: dependency-visualizer                         │
    │ run: execution-simulator                           │
    │ run: task-board-manager                            │
    │                                                    │
    ▼                                                    │
APPROVAL_GATE_3 ◄────────── adjust ──────────────────────┤
    │                                                    │
    │ approve                                            │
    ▼                                                    │
PHASE_4_EXECUTE                                          │
    │ for each wave:                                     │
    │   create: checkpoint                               │
    │   spawn: wave agents (max 7 parallel)              │
    │   monitor: adaptive-replanner (background)         │
    │   update: task-board (real-time)                   │
    │                                                    │
    ├──────────── rollback ──────────────────────────────┤
    │                                                    │
    ▼                                                    │
COMPLETED                                                │
    │                                                    │
    ├──────────── /cascade:sync (optional) ──────────────┘
    │
    ▼
SYNCED (optional)
```
</state_machine>

<agents_used>
## Perspective Agents (Phase 1)
- **user-advocate**: Champions user needs, UX focus
- **tech-skeptic**: Challenges complexity, questions over-engineering
- **innovation-catalyst**: Creative alternatives, breakthrough ideas
- **risk-assessor**: Security, performance, maintenance risks
- **simplicity-champion**: YAGNI advocate, scope reduction

## Synthesis Agents
- **synthesis-engine**: Aggregates perspectives into coherent insights
- **codebase-researcher**: Background pattern scanning (incubating)

## Execution Agents
- **wave-launcher**: Dependency-aware parallel execution
- **parallel-worker**: Executes individual tasks (existing)
- **parallel-orchestrator**: Coordinates waves (existing, enhanced)
</agents_used>

<workflows_index>
| Workflow | Purpose |
|----------|---------|
| cascade-brainstorm.md | Phase 1: Multi-specialist ideation |
| cascade-prd.md | Phase 2: Parallel PRD generation |
| cascade-decompose.md | Phase 3: Task breakdown with visualization |
| cascade-execute.md | Phase 4: Wave-based implementation |
| cascade-resume.md | Resume from checkpoint |
| cascade-status.md | Display session status |
| cascade-board.md | Live task board view |
| cascade-sync.md | Optional GitHub sync |
</workflows_index>

<success_criteria>
1. All 4 phases work completely offline
2. Phase 1 spawns 5 agents in parallel successfully
3. Background research runs while user interacts
4. Synthesis produces coherent output from diverse inputs
5. Wave execution respects dependency graph
6. Checkpoints enable rollback recovery
7. Task board updates in real-time
</success_criteria>
