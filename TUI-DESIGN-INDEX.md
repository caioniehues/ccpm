# CCPM TUI Dashboard Design Index

> **Last Updated**: 2025-12-23
> **Status**: Design Complete, Ready for Implementation

---

## Quick Navigation

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                          CCPM TUI DOCUMENTATION                              │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                              │
│   ┌─────────────────────┐    ┌─────────────────────┐    ┌─────────────────┐ │
│   │  TUI-WIREFRAMES.md  │    │ TUI-DASHBOARD-      │    │ SPEC-WORKFLOW-  │ │
│   │                     │    │ DESIGN.md           │    │ FEATURE-        │ │
│   │  Visual Design      │    │ Technical Arch      │    │ ANALYSIS.md     │ │
│   │  ───────────────    │    │ ──────────────      │    │ ─────────────   │ │
│   │  • 18 wireframes    │    │ • Elm architecture  │    │ • Feature       │ │
│   │  • Color palette    │    │ • Component mapping │    │   comparison    │ │
│   │  • Icon system      │    │ • Go code patterns  │    │ • Gap analysis  │ │
│   │  • Typography       │    │ • File structure    │    │ • Roadmap       │ │
│   │  • Animations       │    │ • Keybindings       │    │                 │ │
│   └─────────────────────┘    └─────────────────────┘    └─────────────────┘ │
│              │                          │                        │          │
│              └──────────────────────────┼────────────────────────┘          │
│                                         │                                   │
│                                         ▼                                   │
│                          ┌─────────────────────────┐                        │
│                          │  CCPM-ADOPTION-SPEC.md  │                        │
│                          │                         │                        │
│                          │  What to adopt from     │                        │
│                          │  spec-workflow          │                        │
│                          └─────────────────────────┘                        │
│                                                                              │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## Document Overview

| Document | Purpose | Key Sections |
|----------|---------|--------------|
| [TUI-WIREFRAMES.md](./TUI-WIREFRAMES.md) | Visual design system and all view wireframes | Design Philosophy, Color Palette, 18 View Wireframes, Animation Specs, Lipgloss Styles |
| [TUI-DASHBOARD-DESIGN.md](./TUI-DASHBOARD-DESIGN.md) | Technical architecture using Charm libraries | Bubbletea Model, Component Mapping, File Structure, Implementation Phases |
| [SPEC-WORKFLOW-FEATURE-ANALYSIS.md](./SPEC-WORKFLOW-FEATURE-ANALYSIS.md) | Feature comparison with spec-workflow | Wizard Patterns, Approval Gates, Offline-First, Gap Analysis |
| [CCPM-ADOPTION-SPEC.md](./CCPM-ADOPTION-SPEC.md) | What CCPM should adopt | Priority Features, Implementation Order, Technical Requirements |

---

## Design Decision

### Why TUI over Web Dashboard?

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                           DECISION RATIONALE                                 │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                              │
│  ✅ TUI (Charm Libraries) Selected                                           │
│  ────────────────────────────────                                           │
│                                                                              │
│  1. TARGET USER: Solo developers who live in the terminal                   │
│     → No context switching between terminal and browser                     │
│                                                                              │
│  2. SIMPLICITY: Single Go binary                                            │
│     → No Node.js, npm packages, or web server required                      │
│                                                                              │
│  3. OFFLINE-FIRST: 100% offline operation                                   │
│     → No network dependencies whatsoever                                    │
│                                                                              │
│  4. PERFORMANCE: ~50ms launch time, ~10MB memory                            │
│     → Instant feedback, minimal resource usage                              │
│                                                                              │
│  5. SSH COMPATIBLE: Works in remote sessions                                │
│     → Pair programming via tmux/ssh just works                              │
│                                                                              │
│  6. CHARM ECOSYSTEM: Excellent developer experience                         │
│     → Bubbletea (Elm arch), Bubbles (components), Lipgloss (styling)        │
│     → Huh (forms/wizards), Glamour (markdown)                               │
│                                                                              │
│  ❌ Web Dashboard Rejected                                                   │
│  ────────────────────────                                                   │
│                                                                              │
│  • Context switching slows down workflow                                    │
│  • Heavier dependencies (Node.js, React, npm packages)                      │
│  • ~2s launch time vs 50ms for TUI                                          │
│  • ~150MB memory vs 10MB for TUI                                            │
│  • Remote access requires tunnels (Cloudflare/ngrok)                        │
│                                                                              │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## Design System Summary

### Neo-Brutalist Color Palette

| Color | Hex | Usage |
|-------|-----|-------|
| VOID | `#0D0D0D` | Background canvas |
| CHARCOAL | `#1A1A2E` | Elevated surfaces |
| GRAPHITE | `#2D2D44` | Borders, dividers |
| SLATE | `#4A4A6A` | Muted text, inactive |
| SILVER | `#8888AA` | Secondary text |
| PEARL | `#E8E8F0` | Primary text |
| ELECTRIC | `#00D4FF` | Primary accent, active |
| PLASMA | `#FF006E` | Urgent, errors |
| VOLT | `#ADFF02` | Success, complete |
| AMBER | `#FFB800` | Warning, in-progress |
| LAVENDER | `#B388FF` | Info, links |

### Wireframe Inventory (18 Views)

| # | View | Description |
|---|------|-------------|
| 1 | Main Dashboard | Epic progress, task list, activity log |
| 2 | Epic Detail | Full epic content with phase status |
| 3 | Task Detail | Task description, acceptance criteria, traces |
| 4 | Wizard - Epic Review | Step 2/4 approval flow |
| 5 | Wizard - Revision | Feedback input for rejected documents |
| 6 | Multi-Epic Selector | List view with filter |
| 7 | Help Overlay | Keyboard shortcuts reference |
| 8 | Loading State | Spinner with message |
| 9 | Error State | Error display with suggestions |
| 10 | Empty State | First-run experience |
| 11 | PRD Detail | PRD content with linked artifacts |
| 12 | Sync Confirmation | Pre-sync review dialog |
| 13 | Notification Toasts | Success/error/warning/info toasts |
| 14 | Wizard - Step 1 | PRD review |
| 15 | Wizard - Step 3 | Tasks review |
| 16 | Wizard - Step 4 | GitHub sync choice |
| 17 | Settings | Preferences and GitHub config |
| 18 | Search/Filter | Modal with type and status filters |

---

## Implementation Priority

### From SPEC-WORKFLOW-FEATURE-ANALYSIS.md

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                         ADOPTION PRIORITY                                    │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                              │
│  P1 (CRITICAL) - Week 1                                                      │
│  ─────────────────────                                                      │
│  □ Phase-based wizard UX with approval gates                                │
│  □ ✅ APPROVED markers in document frontmatter                              │
│  □ Phase recovery (resume from last approved phase)                         │
│                                                                              │
│  P2 (HIGH VALUE) - Week 2                                                    │
│  ────────────────────────                                                   │
│  □ TUI Dashboard (main dashboard view)                                      │
│  □ File watcher for real-time updates                                       │
│  □ Epic/Task detail views                                                   │
│                                                                              │
│  P3 (ENHANCEMENT) - Week 3                                                   │
│  ─────────────────────────                                                  │
│  □ Local issue tracking (GitHub optional)                                   │
│  □ File-based status persistence                                            │
│  □ Steering documents integration                                           │
│                                                                              │
│  P4 (POLISH) - Week 4                                                        │
│  ────────────────────                                                       │
│  □ Animations and loading states                                            │
│  □ Responsive terminal sizing                                               │
│  □ Settings view                                                            │
│                                                                              │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## Technology Stack

| Component | Library | Purpose |
|-----------|---------|---------|
| Framework | [Bubbletea](https://github.com/charmbracelet/bubbletea) | Elm architecture for TUI |
| Components | [Bubbles](https://github.com/charmbracelet/bubbles) | Pre-built UI components |
| Styling | [Lipgloss](https://github.com/charmbracelet/lipgloss) | CSS-like styling for terminal |
| Forms | [Huh](https://github.com/charmbracelet/huh) | Wizard-style approval forms |
| Markdown | [Glamour](https://github.com/charmbracelet/glamour) | Markdown rendering |
| File Watch | [fsnotify](https://github.com/fsnotify/fsnotify) | Cross-platform file watching |

---

## Getting Started

1. **Design Understanding**: Start with [TUI-WIREFRAMES.md](./TUI-WIREFRAMES.md) to understand the visual design
2. **Technical Architecture**: Read [TUI-DASHBOARD-DESIGN.md](./TUI-DASHBOARD-DESIGN.md) for implementation patterns
3. **Feature Context**: Review [SPEC-WORKFLOW-FEATURE-ANALYSIS.md](./SPEC-WORKFLOW-FEATURE-ANALYSIS.md) for feature rationale
4. **Adoption Spec**: Check [CCPM-ADOPTION-SPEC.md](./CCPM-ADOPTION-SPEC.md) for prioritized implementation

---

*This index provides navigation and context for the CCPM TUI Dashboard design documentation.*
