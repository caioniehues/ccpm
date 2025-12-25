# CASCADE FLOW Usage Guide

## Overview

CASCADE FLOW is a 4-phase offline workflow that transforms how you build features:

```
/cascade:start → BRAINSTORM → PRD → DECOMPOSE → EXECUTE → /cascade:sync
                     ↑            ↑         ↑          ↑
                 5 parallel   4 parallel  DAG +    7 parallel
                  agents       writers   simulation   agents
```

---

## Quick Start

```bash
# Start a new feature
/cascade:start user-authentication

# That's it! The workflow guides you through all 4 phases.
```

---

## Commands Reference

| Command | Purpose | Example |
|---------|---------|---------|
| `/cascade:start <name>` | Begin new feature | `/cascade:start dark-mode` |
| `/cascade:resume <name>` | Continue interrupted session | `/cascade:resume dark-mode` |
| `/cascade:status [name]` | Check progress (all or specific) | `/cascade:status` |
| `/cascade:board <name>` | View live task board | `/cascade:board dark-mode` |
| `/cascade:sync <name>` | Push to GitHub (optional) | `/cascade:sync dark-mode` |

---

## Phase-by-Phase Guide

### Phase 1: BRAINSTORM

**What happens:**
- You describe your feature idea
- 5 specialist agents explore it simultaneously:
  - **User Advocate** - UX and user needs
  - **Tech Skeptic** - Challenges complexity
  - **Innovation Catalyst** - Creative alternatives
  - **Risk Assessor** - Security/performance risks
  - **Simplicity Champion** - Minimum viable approach
- Background researcher scans your codebase
- Synthesis engine combines all perspectives

**Your interaction:**
```
You: "I want to add real-time notifications to the app"

[5 agents launch in parallel]
[Codebase researcher runs in background]

Claude: "Here's what the specialists found:

THEMES (consensus):
- WebSocket preferred over polling
- Need notification preferences UI

CONFLICTS (need your decision):
- Tech Skeptic vs Innovation Catalyst on complexity

How would you like to proceed?
1. Steer - explore a direction more
2. Reject - try different approach
3. Expand - get more details
4. Approve - proceed to PRD"

You: "Approve"
```

### Phase 2: PRD GENERATION

**What happens:**
- 4 section writers work in parallel:
  - Executive Summary
  - Requirements (functional + non-functional)
  - Success Metrics
  - Dependencies & Constraints
- Consistency validator checks for conflicts
- Tradeoff analyst runs in background

**Your interaction:**
```
Claude: "PRD Generated: real-time-notifications

[Shows complete PRD]

Validation: PASSED (12/12 checks)
Tradeoffs: 3 identified

Options:
1. Approve - proceed to decomposition
2. Edit - modify specific sections
3. Regenerate - different emphasis
4. Back - return to brainstorm"

You: "Approve"
```

### Phase 3: DECOMPOSITION

**What happens:**
- Task decomposer breaks PRD into 2-8 hour tasks
- Dependency visualizer creates ASCII DAG
- Execution simulator predicts timeline
- Task board initialized

**Your interaction:**
```
Claude: "Tasks Created: 8

DEPENDENCY GRAPH:
  Wave 1: [001] [002] [003]  ← parallel
              ↓
  Wave 2: [004] [005]        ← parallel
              ↓
  Wave 3: [006] [007]        ← parallel
              ↓
  Wave 4: [008]              ← final

Forecast:
- Waves: 4
- Duration: ~12 hours
- Parallelization: 66%

Options:
1. Approve - start execution
2. Adjust tasks - modify breakdown
3. Adjust order - change waves"

You: "Approve"
```

### Phase 4: EXECUTION

**What happens:**
- Wave launcher executes tasks in dependency order
- Up to 7 parallel agents per wave
- Checkpoints created at wave boundaries
- Adaptive replanner monitors for deviations
- Live task board updates in real-time

**Your interaction:**
```
Claude: "Execution Started: real-time-notifications

Wave 1 of 4 launching with 3 parallel agents...

You can:
- Continue working - I'll update you
- /cascade:board - view live progress
- Say 'pause' - stop after current wave"

[Later...]

Claude: "Wave 1 Complete!
Tasks: 3/3 ✓
Duration: 2h 15m

Proceeding to Wave 2..."

[Eventually...]

Claude: "Execution Complete!

Summary:
- Tasks: 8/8 completed
- Duration: 11h 30m
- Waves: 4

Next: /cascade:sync real-time-notifications"
```

---

## Directory Structure

After running CASCADE FLOW, you'll have:

```
.claude/
├── brainstorm/{session-id}/
│   ├── session.md              # Master state
│   ├── perspectives/           # 5 agent outputs
│   │   ├── user-advocate.md
│   │   ├── tech-skeptic.md
│   │   ├── innovation-catalyst.md
│   │   ├── risk-assessor.md
│   │   └── simplicity-champion.md
│   ├── research/               # Codebase findings
│   │   ├── patterns.md
│   │   ├── relevant-files.md
│   │   └── prior-art.md
│   └── synthesis/              # Combined insights
│       ├── themes.md
│       ├── conflicts.md
│       └── recommendations.md
│
├── prds/{feature}.md           # Generated PRD
│
└── epics/{feature}/
    ├── epic.md                 # Epic overview
    ├── 001.md ... 00N.md       # Task files
    ├── dependency-graph.md     # ASCII DAG
    ├── execution-forecast.md   # Timeline prediction
    ├── task-board.md           # Live status
    ├── execution-status.md     # Runtime tracking
    ├── checkpoints/            # Rollback points
    └── waves/                  # Wave tracking
```

---

## Recovery & Rollback

### Resume Interrupted Session
```bash
/cascade:resume my-feature
# Automatically detects current phase and continues
```

### Rollback to Checkpoint
During execution, if something goes wrong:
```
You: "rollback"

Claude: "Available checkpoints:
1. wave-001-20251225231500
2. wave-002-20251225233000

Which checkpoint?"

You: "1"

Claude: "Rolled back to Wave 1. Ready to restart."
```

---

## Tips & Best Practices

### 1. Give Rich Context in Phase 1
The more detail you provide, the better the perspectives:
```
# Good
"I want to add user authentication with OAuth,
supporting Google and GitHub, with role-based
access control for admin vs regular users"

# Less Good
"add login"
```

### 2. Use Steering in Brainstorm
Don't just approve - steer the exploration:
```
You: "Steer toward the simplicity-champion's approach"
# Re-runs with emphasis on minimal solution
```

### 3. Check the Board During Execution
```bash
/cascade:board my-feature
# See real-time progress, which agents are working
```

### 4. Review Perspectives Individually
The perspective files are valuable references:
```bash
# Read what the tech-skeptic found
cat .claude/brainstorm/*/perspectives/tech-skeptic.md
```

### 5. Use Checkpoints Liberally
Execution creates automatic checkpoints, but you can pause anytime:
```
You: "pause"
# Stops after current wave, checkpoint created
```

---

## Comparison with Standard CCPM

| Aspect | Standard CCPM | CASCADE FLOW |
|--------|---------------|--------------|
| PRD Creation | Sequential Q&A | 5 parallel specialists |
| Research | Manual | Background incubation |
| Task Planning | Single pass | DAG + simulation |
| Execution | Start immediately | Wave-based with checkpoints |
| Recovery | Manual | Automatic checkpoints |
| Visibility | Status commands | Live task board |

---

## Example: Full Feature Flow

```bash
# 1. Start
/cascade:start api-rate-limiting

# 2. Describe your idea
"I need to add rate limiting to our API endpoints.
Should support per-user and per-endpoint limits,
with configurable thresholds and graceful degradation."

# 3. Review synthesis, approve when ready
"Approve"

# 4. Review PRD, edit if needed
"Edit the success metrics section"
[make changes]
"Approve"

# 5. Review task breakdown
"Approve"

# 6. Watch execution (or continue other work)
/cascade:board api-rate-limiting

# 7. When complete, optionally sync to GitHub
/cascade:sync api-rate-limiting
```

---

## Troubleshooting

| Issue | Solution |
|-------|----------|
| Agent timeout | Workflow continues with available perspectives |
| Validation fails | Fix issues shown, re-validate automatically |
| Task fails | Choose: retry, skip, or rollback |
| Session lost | `/cascade:resume <name>` recovers from last state |
| Want to restart | Start new session, old one preserved |

---

## Summary

```
/cascade:start <feature>     # Begin the journey
                             # → 5 specialists brainstorm
                             # → PRD generated in parallel
                             # → Tasks decomposed with DAG
                             # → Executed in waves (max 7 parallel)
/cascade:sync <feature>      # Push to GitHub when done
```
