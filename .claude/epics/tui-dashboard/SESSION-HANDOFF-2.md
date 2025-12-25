# TUI Dashboard Session Handoff #2

**Last Updated**: 2025-12-23 20:45 UTC

## What We Did

### 1. Investigated Frontend Agent Issues

Confirmed that the `frontend-ui-ux-engineer` agent is broken:
- Returns empty/no output (0s response time)
- Root cause: Uses `google/gemini-3-pro-high` model which fails silently
- Config location: `~/.config/opencode/oh-my-opencode.json`
- `general` agent works fine (uses different model)

### 2. Created Go Charmbracelet/Bubble Tea TUI Expertise Skill

Built a comprehensive domain expertise skill for Go TUI development:

**Location:** `~/.claude/skills/expertise/go-tui/`

**Structure:**
```
go-tui/
├── SKILL.md                          # Router + essential MVU principles
├── workflows/
│   ├── build-new-tui.md             # Create TUI from scratch
│   ├── add-view.md                  # Add new screens
│   ├── add-component.md             # Create reusable components
│   ├── implement-animations.md      # Spinners, transitions, effects
│   ├── debug-tui.md                 # Debug rendering/state issues
│   ├── write-tests.md               # Test TUI applications
│   ├── optimize-performance.md      # Profile and optimize
│   └── build-release.md             # Cross-compile and release
└── references/
    ├── architecture.md              # MVU pattern, state management
    ├── lipgloss-styling.md          # Colors, borders, layout
    ├── bubbles-components.md        # Spinner, viewport, list, table, etc.
    ├── huh-forms.md                 # Forms and prompts
    └── anti-patterns.md             # Common mistakes to avoid
```

**Coverage:**
- Bubble Tea MVU architecture, commands, messages
- Lip Gloss styling, colors, borders, tables, lists
- Bubbles components (spinner, viewport, list, table, textinput, help, etc.)
- Huh forms, validation, Bubble Tea integration
- Full lifecycle: Build → Debug → Test → Optimize → Ship

---

## Current State

### TUI Dashboard Epic
- **Location:** `.claude/epics/tui-dashboard/`
- **Tasks 001-012:** COMPLETED
- **Tasks 013-018:** PENDING
- **Build:** SUCCESS (4.9MB binary at `ccpm/ccpm-tui`)

### Remaining Tasks

| Task | Description | Type |
|------|-------------|------|
| 013 | Create local tracking commands | Commands (logic) |
| 014 | Create /pm:dashboard command | Commands (logic) |
| 015 | Add animations and loading states | Go TUI code |
| 016 | Implement responsive layout handling | Go TUI code |
| 017 | Testing and documentation | Tests + Docs |
| 018 | Build and release pipeline | DevOps |

### Key Files
- Entry: `ccpm/cmd/ccpm-tui/main.go`
- Core: `ccpm/internal/tui/{model,update,view,keys,styles}.go`
- Views: `ccpm/internal/tui/views/*.go`
- Parsers: `ccpm/internal/tui/parser/*.go`
- Components: `ccpm/internal/tui/components/*.go`

---

## Known Issues

1. **`frontend-ui-ux-engineer` agent broken** - Uses failing Google model
2. **`background_task` tool broken** - All agents fail in background mode
3. **Google model agents fail** - document-writer, multimodal-looker also affected

### Agent Configuration
```json
// ~/.config/opencode/oh-my-opencode.json
{
  "google_auth": true,
  "agents": {
    "oracle": { "model": "zai-coding-plan/glm-4.7" },           // Works
    "librarian": { "model": "zai-coding-plan/glm-4.7" },        // Works
    "frontend-ui-ux-engineer": { "model": "google/gemini-3-pro-high" },  // BROKEN
    "document-writer": { "model": "google/gemini-3-flash" },    // BROKEN
    "multimodal-looker": { "model": "google/gemini-3-flash" }   // BROKEN
  }
}
```

**Fix Option:** Change broken agents to use `zai-coding-plan/glm-4.7` model.

---

## Prompt for Continuing Session

```
Continue implementing the CCPM TUI Dashboard.

## Current State
- Epic: `.claude/epics/tui-dashboard/` with 18 task files
- Tasks 001-012: COMPLETED
- Tasks 013-018: PENDING
- Build: SUCCESS (4.9MB binary at `ccpm/ccpm-tui`)

## New Resource
A Go TUI expertise skill was created at:
`~/.claude/skills/expertise/go-tui/`

This skill contains comprehensive guidance for:
- Bubble Tea MVU architecture
- Lip Gloss styling
- Bubbles components (spinner, viewport, list, etc.)
- Animations and transitions
- Testing TUI apps
- Building and releasing

Use this skill's references when implementing tasks 015-016.

## Remaining Tasks

| Task | File | Description |
|------|------|-------------|
| 013 | 013.md | Create local tracking commands |
| 014 | 014.md | Create /pm:dashboard command |
| 015 | 015.md | Add animations and loading states |
| 016 | 016.md | Implement responsive layout handling |
| 017 | 017.md | Testing and documentation |
| 018 | 018.md | Build and release pipeline |

## Key Context
- This is a Go TUI (not web frontend) - use Bubble Tea/Lip Gloss patterns
- The `frontend-ui-ux-engineer` agent is broken (Google API issues)
- Handle TUI work directly using the go-tui skill

## Build Command
cd /home/caio/dev/ccpm/ccpm && go build ./cmd/ccpm-tui

## Next Steps
1. Read remaining task files (013-018)
2. Start with task 013 or 014 (command logic)
3. For tasks 015-016 (UI work), reference `~/.claude/skills/expertise/go-tui/`
4. Run build after each task to verify
```

---

## Design References

- `TUI-WIREFRAMES.md` - Visual wireframes and Neo-Brutalist color palette
- `TUI-DASHBOARD-DESIGN.md` - Architecture and Charm library patterns
- `TUI-DESIGN-INDEX.md` - Index of all design documents

---

## Session Artifacts

### Created This Session
1. `~/.claude/skills/expertise/go-tui/SKILL.md`
2. `~/.claude/skills/expertise/go-tui/workflows/*.md` (8 files)
3. `~/.claude/skills/expertise/go-tui/references/*.md` (5 files)

### Verified Working
- `explore` agent via `call_omo_agent`
- `general` agent via sync `task` tool
- `oracle` agent (uses working model)
- `librarian` agent (uses working model)
