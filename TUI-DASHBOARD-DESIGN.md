# CCPM TUI Dashboard Design

> **Using Charm Libraries**: Bubbletea, Bubbles, Lipgloss, Huh
> **Language**: Go
> **Design Date**: 2025-12-23

---

## Related Documents

| Document | Purpose |
|----------|---------|
| [TUI-WIREFRAMES.md](./TUI-WIREFRAMES.md) | Visual wireframes, design system, color palette, animations |
| [SPEC-WORKFLOW-FEATURE-ANALYSIS.md](./SPEC-WORKFLOW-FEATURE-ANALYSIS.md) | Feature comparison and adoption roadmap |
| [CCPM-ADOPTION-SPEC.md](./CCPM-ADOPTION-SPEC.md) | What CCPM should adopt from spec-workflow |

---

## Executive Summary

A terminal-native dashboard for CCPM using [Charm](https://github.com/charmbracelet) libraries, providing:
- Real-time epic/task visualization
- Wizard-style approval workflows
- File-watching for live updates
- Zero external dependencies (no browser, no server)

---

## Charm Library Stack

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                        CHARM LIBRARY ECOSYSTEM                               │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                              │
│  ┌─────────────────────────────────────────────────────────────────────┐    │
│  │                         BUBBLETEA                                    │    │
│  │                   (The Elm Architecture)                             │    │
│  │  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐                  │    │
│  │  │   MODEL     │  │   UPDATE    │  │    VIEW     │                  │    │
│  │  │  (State)    │──│  (Events)   │──│  (Render)   │                  │    │
│  │  └─────────────┘  └─────────────┘  └─────────────┘                  │    │
│  └─────────────────────────────────────────────────────────────────────┘    │
│                                    │                                         │
│           ┌────────────────────────┼────────────────────────┐               │
│           │                        │                        │               │
│           ▼                        ▼                        ▼               │
│  ┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐         │
│  │    BUBBLES      │    │    LIPGLOSS     │    │      HUH        │         │
│  │  (Components)   │    │   (Styling)     │    │    (Forms)      │         │
│  │                 │    │                 │    │                 │         │
│  │ • Progress      │    │ • Colors        │    │ • Input         │         │
│  │ • Table         │    │ • Borders       │    │ • Select        │         │
│  │ • List          │    │ • Padding       │    │ • MultiSelect   │         │
│  │ • Spinner       │    │ • Layout        │    │ • Confirm       │         │
│  │ • Viewport      │    │ • Alignment     │    │ • Groups        │         │
│  │ • Help          │    │ • Themes        │    │ • Wizards       │         │
│  └─────────────────┘    └─────────────────┘    └─────────────────┘         │
│                                                                              │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## TUI Dashboard Layout

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                              │
│  ╔═══════════════════════════════════════════════════════════════════════╗  │
│  ║  CCPM Dashboard                                    Branch: epic/auth  ║  │
│  ╠═══════════════════════════════════════════════════════════════════════╣  │
│  ║                                                                        ║  │
│  ║  ┌──────────────────────────────────────────────────────────────────┐ ║  │
│  ║  │  ACTIVE EPIC: user-authentication                                │ ║  │
│  ║  │  ──────────────────────────────────────────────────────────────  │ ║  │
│  ║  │  Status: IN PROGRESS                    PRD: ✅ APPROVED         │ ║  │
│  ║  │                                                                  │ ║  │
│  ║  │  Progress                                                        │ ║  │
│  ║  │  ████████████████████░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░  50% (3/6)  │ ║  │
│  ║  └──────────────────────────────────────────────────────────────────┘ ║  │
│  ║                                                                        ║  │
│  ║  ┌──────────────────────────────────────────────────────────────────┐ ║  │
│  ║  │  TASKS                                                [j/k nav]  │ ║  │
│  ║  │  ────────────────────────────────────────────────────────────    │ ║  │
│  ║  │    ✅  001: Set up auth middleware                               │ ║  │
│  ║  │    ✅  002: Create user model                                    │ ║  │
│  ║  │    ✅  003: Implement login endpoint                             │ ║  │
│  ║  │  ▶ 🔄  004: Add session management            ← IN PROGRESS     │ ║  │
│  ║  │    ⏳  005: Create registration flow                             │ ║  │
│  ║  │    ⏳  006: Add password reset                                   │ ║  │
│  ║  └──────────────────────────────────────────────────────────────────┘ ║  │
│  ║                                                                        ║  │
│  ║  ┌──────────────────────────────────────────────────────────────────┐ ║  │
│  ║  │  ACTIVITY LOG                                    [auto-refresh]  │ ║  │
│  ║  │  ────────────────────────────────────────────────────────────    │ ║  │
│  ║  │  14:32  Task 003 marked complete                                 │ ║  │
│  ║  │  14:28  Task 004 started                                         │ ║  │
│  ║  │  14:15  Epic approved for work                                   │ ║  │
│  ║  │  13:45  PRD created and approved                                 │ ║  │
│  ║  └──────────────────────────────────────────────────────────────────┘ ║  │
│  ║                                                                        ║  │
│  ║  ┌──────────────────────────────────────────────────────────────────┐ ║  │
│  ║  │ [e]Epic  [t]Tasks  [p]PRD  [s]Sync  [w]Wizard  [?]Help  [q]Quit  │ ║  │
│  ║  └──────────────────────────────────────────────────────────────────┘ ║  │
│  ╚═══════════════════════════════════════════════════════════════════════╝  │
│                                                                              │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## Component Mapping

### CCPM Feature → Charm Component

| CCPM Feature | Charm Component | Description |
|--------------|-----------------|-------------|
| Epic progress | `bubbles/progress` | Animated progress bar with percentage |
| Task list | `bubbles/list` | Navigable list with custom item renderer |
| Task details | `bubbles/viewport` | Scrollable markdown content |
| Activity log | `bubbles/viewport` | Auto-scrolling event feed |
| Navigation | `bubbles/help` | Auto-generated keybinding help |
| Epic selector | `bubbles/list` | Multi-epic navigation |
| Status spinner | `bubbles/spinner` | Loading/syncing indicator |
| Approval wizard | `huh` forms | Multi-step approval workflow |
| Markdown preview | `glamour` | Render PRD/epic/task content |
| Styled output | `lipgloss` | Colors, borders, layout |

---

## Architecture

### Elm Architecture for CCPM

```go
// Model - Application State
type Model struct {
    // Data
    epics       []Epic
    activeEpic  *Epic
    tasks       []Task
    activeTask  int
    activityLog []ActivityEntry

    // UI State
    view        ViewMode  // dashboard | wizard | details | help
    loading     bool
    lastError   error

    // Components
    taskList    list.Model
    progress    progress.Model
    viewport    viewport.Model
    spinner     spinner.Model
    help        help.Model

    // File Watcher
    watcher     *fsnotify.Watcher

    // Dimensions
    width       int
    height      int
}

// ViewMode enum
type ViewMode int
const (
    ViewDashboard ViewMode = iota
    ViewWizard
    ViewDetails
    ViewHelp
)
```

### Message Flow

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                           MESSAGE FLOW                                       │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                              │
│  ┌─────────────┐                                                            │
│  │   INPUTS    │                                                            │
│  └──────┬──────┘                                                            │
│         │                                                                    │
│         ├──────────────────────┬──────────────────────┐                     │
│         ▼                      ▼                      ▼                     │
│  ┌─────────────┐       ┌─────────────┐       ┌─────────────┐               │
│  │  Keyboard   │       │ File Watch  │       │   Timer     │               │
│  │  tea.KeyMsg │       │ FileChanged │       │  tea.Tick   │               │
│  └──────┬──────┘       └──────┬──────┘       └──────┬──────┘               │
│         │                     │                     │                       │
│         └──────────────────────┼──────────────────────┘                     │
│                               │                                              │
│                               ▼                                              │
│                      ┌─────────────────┐                                    │
│                      │     UPDATE      │                                    │
│                      │   (Reducer)     │                                    │
│                      └────────┬────────┘                                    │
│                               │                                              │
│         ┌─────────────────────┼─────────────────────┐                       │
│         ▼                     ▼                     ▼                       │
│  ┌─────────────┐       ┌─────────────┐       ┌─────────────┐               │
│  │ Update List │       │ Update Prog │       │ Parse Files │               │
│  └─────────────┘       └─────────────┘       └─────────────┘               │
│                               │                                              │
│                               ▼                                              │
│                      ┌─────────────────┐                                    │
│                      │      VIEW       │                                    │
│                      │   (Render)      │                                    │
│                      └────────┬────────┘                                    │
│                               │                                              │
│                               ▼                                              │
│                      ┌─────────────────┐                                    │
│                      │  Terminal Out   │                                    │
│                      └─────────────────┘                                    │
│                                                                              │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## Key Features Implementation

### 1. Real-Time File Watching

```go
// File watcher integration
func watchFiles(p *tea.Program) {
    watcher, _ := fsnotify.NewWatcher()
    defer watcher.Close()

    // Watch .claude/epics and .claude/prds
    watcher.Add(".claude/epics")
    watcher.Add(".claude/prds")

    for {
        select {
        case event := <-watcher.Events:
            if event.Op&fsnotify.Write == fsnotify.Write {
                // Parse the changed file
                epic, tasks := parseEpicFiles(event.Name)

                // Send update to Bubbletea program
                p.Send(FileChangedMsg{
                    Path:  event.Name,
                    Epic:  epic,
                    Tasks: tasks,
                })
            }
        }
    }
}

// Message type
type FileChangedMsg struct {
    Path  string
    Epic  Epic
    Tasks []Task
}
```

### 2. Progress Bar with Animation

```go
import "github.com/charmbracelet/bubbles/progress"

func newProgress() progress.Model {
    return progress.New(
        progress.WithDefaultGradient(),
        progress.WithWidth(50),
        progress.WithoutPercentage(),
    )
}

func (m Model) renderProgress() string {
    completed := countCompleted(m.tasks)
    total := len(m.tasks)
    percent := float64(completed) / float64(total)

    return lipgloss.JoinVertical(
        lipgloss.Left,
        "Progress",
        m.progress.ViewAs(percent) + fmt.Sprintf(" %d/%d", completed, total),
    )
}
```

### 3. Task List with Custom Rendering

```go
import "github.com/charmbracelet/bubbles/list"

// Custom item delegate for task rendering
type taskDelegate struct{}

func (d taskDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
    task := item.(Task)

    // Status icon
    icon := "⏳"
    if task.Completed {
        icon = "✅"
    } else if task.InProgress {
        icon = "🔄"
    }

    // Cursor indicator
    cursor := "  "
    if index == m.Index() {
        cursor = "▶ "
    }

    // Style based on status
    style := lipgloss.NewStyle()
    if task.InProgress {
        style = style.Bold(true).Foreground(lipgloss.Color("212"))
    }

    fmt.Fprintf(w, "%s%s %s: %s\n", cursor, icon, task.ID, style.Render(task.Title))
}
```

### 4. Approval Wizard with Huh

```go
import "github.com/charmbracelet/huh"

func runApprovalWizard(phase string, content string) (bool, error) {
    var approved bool
    var feedback string

    form := huh.NewForm(
        // Phase 1: Show content
        huh.NewGroup(
            huh.NewNote().
                Title(fmt.Sprintf("Review %s", phase)).
                Description(content),
        ),

        // Phase 2: Approval
        huh.NewGroup(
            huh.NewConfirm().
                Title(fmt.Sprintf("Approve %s?", phase)).
                Description("Do you want to proceed to the next phase?").
                Affirmative("Yes, approve").
                Negative("No, revise").
                Value(&approved),
        ),

        // Phase 3: Feedback (if not approved)
        huh.NewGroup(
            huh.NewText().
                Title("What changes are needed?").
                Value(&feedback),
        ).WithHideFunc(func() bool { return approved }),
    )

    err := form.Run()
    return approved, err
}
```

### 5. Styled Layout with Lipgloss

> **Note**: Uses the Neo-Brutalist color palette defined in [TUI-WIREFRAMES.md](./TUI-WIREFRAMES.md#color-palette-ansi-true-color)

```go
import "github.com/charmbracelet/lipgloss"

// Neo-Brutalist Color Palette (from TUI-WIREFRAMES.md)
var (
    // Base Colors
    Void      = lipgloss.Color("#0D0D0D")  // Background - the canvas
    Charcoal  = lipgloss.Color("#1A1A2E")  // Elevated surfaces
    Graphite  = lipgloss.Color("#2D2D44")  // Borders, dividers
    Slate     = lipgloss.Color("#4A4A6A")  // Muted text, inactive
    Silver    = lipgloss.Color("#8888AA")  // Secondary text
    Pearl     = lipgloss.Color("#E8E8F0")  // Primary text

    // Accent Colors
    Electric  = lipgloss.Color("#00D4FF")  // Primary accent, progress, active
    Plasma    = lipgloss.Color("#FF006E")  // Urgent, errors, attention
    Volt      = lipgloss.Color("#ADFF02")  // Success, complete, go
    Amber     = lipgloss.Color("#FFB800")  // Warning, in-progress, caution
    Lavender  = lipgloss.Color("#B388FF")  // Info, links, interactive
)

// Styles using Neo-Brutalist palette
var (
    titleStyle = lipgloss.NewStyle().
        Bold(true).
        Foreground(Electric).
        MarginBottom(1)

    boxStyle = lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder()).
        BorderForeground(Electric).
        Padding(1, 2)

    statusBarStyle = lipgloss.NewStyle().
        Background(Charcoal).
        Foreground(Silver).
        Padding(0, 1)

    helpStyle = lipgloss.NewStyle().
        Foreground(Slate).
        MarginTop(1)

    // Additional semantic styles
    successStyle = lipgloss.NewStyle().
        Foreground(Volt).
        Bold(true)

    errorStyle = lipgloss.NewStyle().
        Foreground(Plasma).
        Bold(true)

    warningStyle = lipgloss.NewStyle().
        Foreground(Amber).
        Bold(true)

    infoStyle = lipgloss.NewStyle().
        Foreground(Lavender)
)

func (m Model) renderHeader() string {
    title := titleStyle.Render("CCPM Dashboard")
    branch := statusBarStyle.Render("Branch: " + m.currentBranch)

    return lipgloss.JoinHorizontal(
        lipgloss.Top,
        title,
        lipgloss.PlaceHorizontal(m.width-lipgloss.Width(title)-2, lipgloss.Right, branch),
    )
}
```

### 6. Activity Log Viewport

```go
import "github.com/charmbracelet/bubbles/viewport"

func (m *Model) initViewport() {
    m.viewport = viewport.New(m.width-4, 6)
    m.viewport.SetContent(m.formatActivityLog())
}

func (m Model) formatActivityLog() string {
    var lines []string
    for _, entry := range m.activityLog {
        timestamp := entry.Time.Format("15:04")
        line := fmt.Sprintf("%s  %s",
            lipgloss.NewStyle().Foreground(mutedColor).Render(timestamp),
            entry.Message,
        )
        lines = append(lines, line)
    }
    return strings.Join(lines, "\n")
}
```

---

## Wizard Mode: Epic Creation Flow

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                        WIZARD MODE: /pm:epic-wizard                          │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                              │
│  ╔═══════════════════════════════════════════════════════════════════════╗  │
│  ║                                                                        ║  │
│  ║   Epic Creation Wizard                                    Step 1 of 4  ║  │
│  ║   ────────────────────────                                             ║  │
│  ║                                                                        ║  │
│  ║   ┌────────────────────────────────────────────────────────────────┐  ║  │
│  ║   │                                                                │  ║  │
│  ║   │  PHASE 1: Product Requirements Document                       │  ║  │
│  ║   │                                                                │  ║  │
│  ║   │  ┌──────────────────────────────────────────────────────────┐ │  ║  │
│  ║   │  │  # User Authentication PRD                               │ │  ║  │
│  ║   │  │                                                          │ │  ║  │
│  ║   │  │  ## Executive Summary                                    │ │  ║  │
│  ║   │  │  Implement secure user authentication with...            │ │  ║  │
│  ║   │  │                                                          │ │  ║  │
│  ║   │  │  ## Problem Statement                                    │ │  ║  │
│  ║   │  │  Users cannot currently log in to...                     │ │  ║  │
│  ║   │  │                                                   [↓ more] │  ║  │
│  ║   │  └──────────────────────────────────────────────────────────┘ │  ║  │
│  ║   │                                                                │  ║  │
│  ║   └────────────────────────────────────────────────────────────────┘  ║  │
│  ║                                                                        ║  │
│  ║   ┌────────────────────────────────────────────────────────────────┐  ║  │
│  ║   │  ? Do you approve this PRD?                                    │  ║  │
│  ║   │                                                                │  ║  │
│  ║   │    › Yes, proceed to Epic generation                          │  ║  │
│  ║   │      No, I need to make changes                                │  ║  │
│  ║   └────────────────────────────────────────────────────────────────┘  ║  │
│  ║                                                                        ║  │
│  ║   [↑/↓] Navigate   [enter] Select   [esc] Cancel                      ║  │
│  ║                                                                        ║  │
│  ╚═══════════════════════════════════════════════════════════════════════╝  │
│                                                                              │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## Views & Navigation

### View Hierarchy

```
                                   ┌─────────────────┐
                                   │   Dashboard     │
                                   │   (Default)     │
                                   └────────┬────────┘
                                            │
              ┌───────────────┬─────────────┼─────────────┬───────────────┐
              ▼               ▼             ▼             ▼               ▼
       ┌────────────┐  ┌────────────┐  ┌────────────┐  ┌────────────┐  ┌────────────┐
       │  [e] Epic  │  │ [t] Tasks  │  │  [p] PRD   │  │ [w] Wizard │  │  [?] Help  │
       │  Details   │  │  Details   │  │   View     │  │   Mode     │  │   View     │
       └────────────┘  └────────────┘  └────────────┘  └────────────┘  └────────────┘
              │               │             │             │
              │               │             │             │
              ▼               ▼             ▼             ▼
       ┌────────────┐  ┌────────────┐  ┌────────────┐  ┌────────────┐
       │  Viewport  │  │  Viewport  │  │  Glamour   │  │    Huh     │
       │  (scroll)  │  │  (scroll)  │  │  Markdown  │  │   Forms    │
       └────────────┘  └────────────┘  └────────────┘  └────────────┘
```

### Keybindings

```go
type keyMap struct {
    Up        key.Binding
    Down      key.Binding
    Enter     key.Binding
    Back      key.Binding
    Epic      key.Binding
    Tasks     key.Binding
    PRD       key.Binding
    Sync      key.Binding
    Wizard    key.Binding
    Help      key.Binding
    Quit      key.Binding
}

var keys = keyMap{
    Up:     key.NewBinding(key.WithKeys("up", "k"),    key.WithHelp("↑/k", "up")),
    Down:   key.NewBinding(key.WithKeys("down", "j"),  key.WithHelp("↓/j", "down")),
    Enter:  key.NewBinding(key.WithKeys("enter"),      key.WithHelp("enter", "select")),
    Back:   key.NewBinding(key.WithKeys("esc"),        key.WithHelp("esc", "back")),
    Epic:   key.NewBinding(key.WithKeys("e"),          key.WithHelp("e", "epic details")),
    Tasks:  key.NewBinding(key.WithKeys("t"),          key.WithHelp("t", "task details")),
    PRD:    key.NewBinding(key.WithKeys("p"),          key.WithHelp("p", "view PRD")),
    Sync:   key.NewBinding(key.WithKeys("s"),          key.WithHelp("s", "sync to GitHub")),
    Wizard: key.NewBinding(key.WithKeys("w"),          key.WithHelp("w", "wizard mode")),
    Help:   key.NewBinding(key.WithKeys("?"),          key.WithHelp("?", "help")),
    Quit:   key.NewBinding(key.WithKeys("q", "ctrl+c"),key.WithHelp("q", "quit")),
}
```

---

## File Structure

```
ccpm/
├── cmd/
│   └── dashboard/
│       └── main.go              # Entry point
│
├── internal/
│   └── tui/
│       ├── model.go             # Bubbletea Model
│       ├── update.go            # Update function (event handling)
│       ├── view.go              # View function (rendering)
│       ├── keys.go              # Keybinding definitions
│       ├── styles.go            # Lipgloss styles
│       │
│       ├── components/
│       │   ├── epic_list.go     # Epic selection list
│       │   ├── task_list.go     # Task list with status
│       │   ├── progress.go      # Progress bar wrapper
│       │   ├── activity.go      # Activity log viewport
│       │   └── header.go        # Dashboard header
│       │
│       ├── views/
│       │   ├── dashboard.go     # Main dashboard view
│       │   ├── details.go       # Epic/task detail view
│       │   ├── wizard.go        # Wizard mode (Huh forms)
│       │   └── help.go          # Help view
│       │
│       └── parser/
│           ├── epic.go          # Parse .claude/epics/
│           ├── prd.go           # Parse .claude/prds/
│           └── watcher.go       # fsnotify file watcher
│
├── go.mod
└── go.sum
```

---

## Dependencies

```go
// go.mod
module github.com/your/ccpm-tui

go 1.21

require (
    github.com/charmbracelet/bubbletea v1.0.0
    github.com/charmbracelet/bubbles v0.20.0
    github.com/charmbracelet/lipgloss v1.0.0
    github.com/charmbracelet/huh v0.6.0
    github.com/charmbracelet/glamour v0.8.0
    github.com/fsnotify/fsnotify v1.7.0
    gopkg.in/yaml.v3 v3.0.1
)
```

---

## Integration with CCPM

### Command: `/pm:dashboard`

```bash
# Launch TUI dashboard
ccpm-tui

# Or via Claude Code slash command
/pm:dashboard
```

### Slash Command Definition

```yaml
# ccpm/commands/pm/dashboard.md
---
description: Launch the TUI dashboard for visual epic/task tracking
---

Launch the CCPM TUI dashboard for real-time visualization of:
- Epic progress and status
- Task completion tracking
- Activity log
- Approval workflow wizards

## Usage

Run the dashboard binary:
\`\`\`bash
ccpm-tui
\`\`\`

Or use the wizard mode for new epics:
\`\`\`bash
ccpm-tui --wizard
\`\`\`

## Features

- **Real-time updates**: File watcher detects changes instantly
- **Keyboard navigation**: Vim-style j/k navigation
- **Wizard mode**: Step-through approval gates
- **Markdown preview**: View PRDs/epics with syntax highlighting
- **Offline first**: No network required, reads local files only
```

---

## Comparison: TUI vs Web Dashboard

> **✅ TUI (Charm) was selected** for CCPM based on solo developer focus and terminal-native workflow.

```
┌────────────────────────────────────────────────────────────────────────┐
│                    TUI vs WEB DASHBOARD COMPARISON                      │
├────────────────────┬───────────────────────┬───────────────────────────┤
│ Feature            │ TUI (Charm) ✅ CHOSEN │ Web (Fastify + React)     │
├────────────────────┼───────────────────────┼───────────────────────────┤
│ Launch time        │ ~50ms                 │ ~2s (server + browser)    │
│ Memory usage       │ ~10MB                 │ ~150MB (Node + browser)   │
│ Dependencies       │ Single binary         │ Node.js, npm packages     │
│ Context switch     │ None (stays in term)  │ Alt-tab to browser        │
│ SSH compatible     │ ✅ Yes                │ ❌ Needs port forwarding  │
│ Remote sharing     │ Pair with tmux        │ Cloudflare/ngrok tunnels  │
│ Offline            │ ✅ 100%               │ ✅ 100% (localhost)       │
│ Mobile viewing     │ ❌ No                 │ ✅ Yes (with tunnel)      │
│ Rich graphics      │ Unicode/ASCII only    │ Full HTML/CSS/SVG         │
│ Mouse support      │ ✅ Yes (Bubbletea)    │ ✅ Yes                    │
│ Distribution       │ Single Go binary      │ npm package               │
│ Updates            │ Download new binary   │ npm update                │
├────────────────────┼───────────────────────┼───────────────────────────┤
│ BEST FOR           │ Solo terminal users   │ Team dashboards           │
└────────────────────┴───────────────────────┴───────────────────────────┘
```

---

## Implementation Phases

### Phase 1: Core Dashboard (3-4 days)

- [ ] Set up Go project with Charm dependencies
- [ ] Implement Model/Update/View structure
- [ ] Create task list component
- [ ] Add progress bar
- [ ] Implement file parser for epics/tasks
- [ ] Add basic file watching

### Phase 2: Views & Navigation (2-3 days)

- [ ] Dashboard view (main)
- [ ] Detail views (epic, task, PRD)
- [ ] Help view with keybindings
- [ ] Viewport for scrollable content
- [ ] Glamour markdown rendering

### Phase 3: Wizard Mode (2-3 days)

- [ ] Huh form for PRD approval
- [ ] Huh form for epic approval
- [ ] Huh form for task approval
- [ ] Phase recovery (resume from marker)
- [ ] Revision feedback loop

### Phase 4: Polish (1-2 days)

- [ ] Responsive layout (terminal resize)
- [ ] Color themes
- [ ] Error handling
- [ ] Build/release pipeline
- [ ] Integration with CCPM commands

---

## Example Output

```
$ ccpm-tui

  ╭──────────────────────────────────────────────────────────────────────╮
  │  CCPM Dashboard                                  Branch: epic/auth   │
  ├──────────────────────────────────────────────────────────────────────┤
  │                                                                      │
  │  ┌─ Active Epic ─────────────────────────────────────────────────┐  │
  │  │  user-authentication                                          │  │
  │  │  ████████████████████░░░░░░░░░░░░░░░░░░░░░░░░░░░░░  50% (3/6) │  │
  │  └───────────────────────────────────────────────────────────────┘  │
  │                                                                      │
  │  Tasks                                                               │
  │    ✅ 001: Set up auth middleware                                   │
  │    ✅ 002: Create user model                                        │
  │    ✅ 003: Implement login endpoint                                 │
  │  ▶ 🔄 004: Add session management                                   │
  │    ⏳ 005: Create registration flow                                 │
  │    ⏳ 006: Add password reset                                       │
  │                                                                      │
  │  Activity                                                            │
  │  14:32  Task 003 completed                                          │
  │  14:28  Task 004 started                                            │
  │  14:15  Epic approved                                               │
  │                                                                      │
  │  [e]Epic [t]Task [p]PRD [w]Wizard [s]Sync [?]Help [q]Quit          │
  ╰──────────────────────────────────────────────────────────────────────╯
```

---

*Design complete. Ready for implementation.*
