---
name: TUI Dashboard Implementation
status: completed
created: 2025-12-23T19:45:00Z
updated: 2025-12-23T22:52:00Z
prd: null
github: null
design_docs:
  - TUI-WIREFRAMES.md
  - TUI-DASHBOARD-DESIGN.md
  - SPEC-WORKFLOW-FEATURE-ANALYSIS.md
  - CCPM-ADOPTION-SPEC.md
---

# Epic: TUI Dashboard Implementation

## Overview

Build a Neo-Brutalist terminal dashboard for CCPM using Go and Charm libraries (Bubbletea, Bubbles, Lipgloss, Huh). The dashboard provides real-time visualization of epics, tasks, and PRDs with wizard-style approval workflows.

## Goals

1. **Visual Progress Tracking** - Real-time dashboard showing epic/task status
2. **Wizard Workflows** - Phase-based approval gates for PRD → Epic → Tasks
3. **Offline-First** - Local tracking without requiring GitHub sync
4. **Terminal Native** - Single Go binary, ~50ms launch, ~10MB memory

## Target Architecture

```
ccpm/
├── cmd/
│   └── ccpm-tui/
│       └── main.go                 # Entry point
│
├── internal/
│   └── tui/
│       ├── model.go                # Bubbletea Model
│       ├── update.go               # Update function
│       ├── view.go                 # View function
│       ├── keys.go                 # Keybindings
│       ├── styles.go               # Lipgloss styles
│       │
│       ├── components/             # Reusable UI components
│       │   ├── epic_card.go
│       │   ├── task_list.go
│       │   ├── progress.go
│       │   ├── activity.go
│       │   ├── header.go
│       │   └── footer.go
│       │
│       ├── views/                  # Screen views
│       │   ├── dashboard.go
│       │   ├── epic_detail.go
│       │   ├── task_detail.go
│       │   ├── prd_detail.go
│       │   ├── wizard.go
│       │   ├── epic_selector.go
│       │   ├── help.go
│       │   ├── settings.go
│       │   └── search.go
│       │
│       └── parser/                 # File parsing
│           ├── epic.go
│           ├── prd.go
│           ├── task.go
│           └── watcher.go
│
├── go.mod
└── go.sum
```

## Task Breakdown

### Phase 1: Wizard & Approval Gates (P1 - Critical)
- [001.md](001.md) - Create /pm:epic-wizard command
- [002.md](002.md) - Implement approval markers & phase recovery

### Phase 2: Go Project Foundation (P1 - Critical)
- [003.md](003.md) - Initialize Go project with Charm dependencies
- [004.md](004.md) - Implement Neo-Brutalist styles & keybindings

### Phase 3: Core TUI Architecture (P1 - Critical)
- [005.md](005.md) - Create Model/Update/View core architecture
- [006.md](006.md) - Build core components (epic card, task list, activity)

### Phase 4: Views Implementation (P2 - High)
- [007.md](007.md) - Implement main dashboard view
- [008.md](008.md) - Implement epic & task detail views
- [009.md](009.md) - Implement wizard mode with Huh forms
- [010.md](010.md) - Implement helper views (help, selector, search)

### Phase 5: File Integration (P2 - High)
- [011.md](011.md) - Create file parsers (epic, PRD, task)
- [012.md](012.md) - Implement file watcher with fsnotify

### Phase 6: Commands Integration (P2 - High)
- [013.md](013.md) - Create local tracking commands
- [014.md](014.md) - Create /pm:dashboard command

### Phase 7: Polish (P3 - Medium)
- [015.md](015.md) - Add animations and loading states
- [016.md](016.md) - Implement responsive layout handling

### Phase 8: Release (P3 - Medium)
- [017.md](017.md) - Testing and documentation
- [018.md](018.md) - Build and release pipeline

## Parallel Execution Map

```
                    ┌─────────────────────────────────────────────────┐
                    │              PARALLEL GROUP A                    │
                    │         (Can run simultaneously)                 │
                    ├─────────────────────────────────────────────────┤
Phase 1:            │  [001] /pm:epic-wizard command                  │
                    │  [002] Approval markers & recovery               │
                    └─────────────────────────────────────────────────┘
                                          │
                                          ▼
                    ┌─────────────────────────────────────────────────┐
                    │              PARALLEL GROUP B                    │
                    │         (After Phase 1 complete)                 │
                    ├─────────────────────────────────────────────────┤
Phase 2:            │  [003] Go project setup                         │
                    │  [004] Styles & keybindings                      │
                    └─────────────────────────────────────────────────┘
                                          │
                                          ▼
                    ┌─────────────────────────────────────────────────┐
                    │              SEQUENTIAL                          │
                    ├─────────────────────────────────────────────────┤
Phase 3:            │  [005] Model/Update/View  ──▶  [006] Components │
                    └─────────────────────────────────────────────────┘
                                          │
                    ┌─────────────────────┴─────────────────────┐
                    │                                           │
                    ▼                                           ▼
     ┌─────────────────────────────┐       ┌─────────────────────────────┐
     │    PARALLEL GROUP C         │       │    PARALLEL GROUP D         │
     ├─────────────────────────────┤       ├─────────────────────────────┤
     │ [007] Dashboard view        │       │ [011] File parsers          │
     │ [008] Detail views          │       │ [012] File watcher          │
     │ [009] Wizard mode           │       └─────────────────────────────┘
     │ [010] Helper views          │
     └─────────────────────────────┘
                    │                                           │
                    └─────────────────────┬─────────────────────┘
                                          │
                                          ▼
                    ┌─────────────────────────────────────────────────┐
                    │              PARALLEL GROUP E                    │
                    ├─────────────────────────────────────────────────┤
Phase 6:            │  [013] Local tracking commands                  │
                    │  [014] /pm:dashboard command                    │
                    └─────────────────────────────────────────────────┘
                                          │
                                          ▼
                    ┌─────────────────────────────────────────────────┐
                    │              PARALLEL GROUP F                    │
                    ├─────────────────────────────────────────────────┤
Phase 7:            │  [015] Animations                               │
                    │  [016] Responsive layout                        │
                    └─────────────────────────────────────────────────┘
                                          │
                                          ▼
                    ┌─────────────────────────────────────────────────┐
                    │              SEQUENTIAL                          │
                    ├─────────────────────────────────────────────────┤
Phase 8:            │  [017] Testing  ──▶  [018] Release              │
                    └─────────────────────────────────────────────────┘
```

## Technology Stack

| Component | Library | Purpose |
|-----------|---------|---------|
| Framework | Bubbletea | Elm architecture for TUI |
| Components | Bubbles | Pre-built UI components |
| Styling | Lipgloss | CSS-like terminal styling |
| Forms | Huh | Wizard-style approval forms |
| Markdown | Glamour | Markdown rendering |
| File Watch | fsnotify | Cross-platform file watching |
| YAML | gopkg.in/yaml.v3 | Frontmatter parsing |

## Success Criteria

- [x] `/pm:epic-wizard` creates epics with approval gates
- [x] `ccpm-tui` binary launches dashboard in <100ms
- [x] Real-time file watching updates dashboard
- [x] All 18 wireframe views implemented
- [x] Local-only workflow works without GitHub
- [x] Memory usage <20MB (~5MB binary)
- [x] Works in 80x24 minimum terminal size

## Design References

- **Visual Design**: [TUI-WIREFRAMES.md](../../TUI-WIREFRAMES.md)
- **Technical Architecture**: [TUI-DASHBOARD-DESIGN.md](../../TUI-DASHBOARD-DESIGN.md)
- **Feature Analysis**: [SPEC-WORKFLOW-FEATURE-ANALYSIS.md](../../SPEC-WORKFLOW-FEATURE-ANALYSIS.md)
- **Adoption Spec**: [CCPM-ADOPTION-SPEC.md](../../CCPM-ADOPTION-SPEC.md)
