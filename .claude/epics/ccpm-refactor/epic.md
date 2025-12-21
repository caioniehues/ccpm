---
name: CCPM Commands Refactoring
status: in-progress
created: 2025-12-21T00:00:00Z
updated: 2025-12-21T13:36:39Z
prd: null
github: null
worktrees: worktrees.md
---

# Epic: Refactor CCPM Commands Using Taches Patterns

## Overview
Refactor 46 CCPM commands to follow taches skill patterns, creating 6 specialist skills and 6 workflow-specific subagents aligned with the PM pipeline.

## Target Workflow
```
PRD Creation → Epic Planning → Task Decomposition → GitHub Sync → Worktree Setup → Parallel Execution
```

## Architecture

### Skills (6)
| Skill | Commands Absorbed | Location |
|-------|-------------------|----------|
| `ccpm-prd` | prd-new, prd-parse, prd-edit | `ccpm/skills/ccpm-prd/` |
| `ccpm-epic` | epic-start, epic-sync, epic-decompose, epic-merge, epic-edit, epic-close, epic-oneshot, epic-refresh | `ccpm/skills/ccpm-epic/` |
| `ccpm-issue` | issue-analyze, issue-start, issue-close, issue-edit, issue-reopen, issue-sync, issue-status | `ccpm/skills/ccpm-issue/` |
| `ccpm-context` | context/create, context/prime, context/update | `ccpm/skills/ccpm-context/` |
| `ccpm-testing` | testing/prime, testing/run | `ccpm/skills/ccpm-testing/` |
| `ccpm-worktree` | epic-start-worktree, worktree management | `ccpm/skills/ccpm-worktree/` |

### Subagents (6 New + 4 Enhanced)
**New**: prd-architect, epic-planner, task-decomposer, github-syncer, worktree-manager, parallel-orchestrator
**Enhanced**: code-analyzer, file-analyzer, parallel-worker, test-runner

## Task Breakdown

### Phase 0: Foundation (Parallel Group A)
- [001.md](001.md) - Create skill directory structure
- [002.md](002.md) - Create shared-references from rules/

### Phase 1: Scripts Update (Can run after Phase 0)
- [003.md](003.md) - Update 14 script-delegating commands to taches YAML

### Phase 2-6: Skills & Subagents (Parallel Group B - after Phase 0)
- [004.md](004.md) - Create ccpm-context skill + workflows
- [005.md](005.md) - Create ccpm-testing skill + workflows
- [006.md](006.md) - Create ccpm-prd skill + workflows + prd-architect subagent
- [007.md](007.md) - Create ccpm-issue skill + workflows + task-decomposer subagent
- [008.md](008.md) - Create ccpm-worktree skill + workflows + worktree-manager subagent
- [009.md](009.md) - Create ccpm-epic skill + workflows + epic-planner & github-syncer subagents

### Phase 7: Orchestration (After Phase 2-6)
- [010.md](010.md) - Create parallel-orchestrator subagent

### Phase 8: Agent Enhancements (Parallel Group C - after Phase 7) ✅
- [011.md](011.md) - ✅ Enhance code-analyzer with issue tracking
- [012.md](012.md) - ✅ Enhance file-analyzer with context skill integration
- [013.md](013.md) - ✅ Enhance parallel-worker with dependency awareness
- [014.md](014.md) - ✅ Enhance test-runner with testing skill integration

### Phase 9: Command Migration (After skills created)
- [015.md](015.md) - Transform PRD commands to routers
- [016.md](016.md) - Transform Epic commands to routers
- [017.md](017.md) - Transform Issue commands to routers
- [018.md](018.md) - Transform Context commands to routers
- [019.md](019.md) - Transform Testing commands to routers

### Phase 10: Cleanup
- [020.md](020.md) - Update AGENTS.md documentation
- [021.md](021.md) - Validate and test all skills

## Parallel Execution Map

```
                    ┌─────────────────────────────────────────────────┐
                    │              PARALLEL GROUP A                    │
                    │  (Can run simultaneously in separate worktrees)  │
                    ├─────────────────────────────────────────────────┤
Phase 0:            │  [001] Directory Structure                       │
                    │  [002] Shared References                         │
                    └─────────────────────────────────────────────────┘
                                          │
                    ┌─────────────────────┴─────────────────────┐
                    │                                           │
                    ▼                                           ▼
    ┌───────────────────────────┐       ┌─────────────────────────────────────────┐
    │ [003] Script Updates      │       │           PARALLEL GROUP B              │
    │ (Sequential)              │       │ (6 skills can run in 6 worktrees)       │
    └───────────────────────────┘       ├─────────────────────────────────────────┤
                                        │ [004] ccpm-context skill                │
                                        │ [005] ccpm-testing skill                │
                                        │ [006] ccpm-prd skill + prd-architect    │
                                        │ [007] ccpm-issue skill + task-decomposer│
                                        │ [008] ccpm-worktree skill + worktree-mgr│
                                        │ [009] ccpm-epic skill + epic-planner    │
                                        └─────────────────────────────────────────┘
                                                          │
                                                          ▼
                                        ┌─────────────────────────────────────────┐
                                        │ [010] parallel-orchestrator subagent    │
                                        │ (Sequential - depends on all skills)    │
                                        └─────────────────────────────────────────┘
                                                          │
                    ┌─────────────────────────────────────┴─────────────────────────────────────┐
                    │                                                                           │
                    ▼                                                                           ▼
    ┌─────────────────────────────────────────┐         ┌─────────────────────────────────────────┐
    │       PARALLEL GROUP C ✅ COMPLETE      │         │           PARALLEL GROUP D              │
    │ (4 agent enhancements - 4 worktrees)    │         │ (5 command migrations - 5 worktrees)    │
    ├─────────────────────────────────────────┤         ├─────────────────────────────────────────┤
    │ [011] ✅ Enhance code-analyzer          │         │ [015] Transform PRD commands            │
    │ [012] ✅ Enhance file-analyzer          │         │ [016] Transform Epic commands           │
    │ [013] ✅ Enhance parallel-worker        │         │ [017] Transform Issue commands          │
    │ [014] ✅ Enhance test-runner            │         │ [018] Transform Context commands        │
    └─────────────────────────────────────────┘         │ [019] Transform Testing commands        │
                    │                                   └─────────────────────────────────────────┘
                    │                                                     │
                    └─────────────────────────┬───────────────────────────┘
                                              │
                                              ▼
                              ┌─────────────────────────────────────────┐
                              │           PARALLEL GROUP E              │
                              │ (2 cleanup tasks - 2 worktrees)         │
                              ├─────────────────────────────────────────┤
                              │ [020] Update AGENTS.md                  │
                              │ [021] Validate and test                 │
                              └─────────────────────────────────────────┘
```

## Success Criteria
- [ ] 6 skills created with router pattern
- [ ] 6 new workflow-aligned subagents created
- [x] 4 existing agents enhanced (PR #5)
- [ ] 27 commands transformed to minimal routers
- [ ] 14 shell scripts retained unchanged
- [ ] All skills < 500 lines
- [ ] AGENTS.md documentation updated
