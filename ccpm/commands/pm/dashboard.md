---
description: Launch the TUI dashboard for visual epic/task tracking
argument-hint: [--wizard] [--epic <name>]
allowed-tools: Bash, Read
---

<objective>
Launch the CCPM TUI dashboard for real-time visualization of epics, tasks,
and PRDs. The dashboard provides keyboard-driven navigation and wizard-style
approval workflows.
</objective>

<process>
## 1. Locate Binary

Check for ccpm-tui binary in order:

```bash
# Check locations in priority order
BINARY=""
if [ -x "./ccpm/ccpm-tui" ]; then
    BINARY="./ccpm/ccpm-tui"
elif [ -x "./ccpm-tui" ]; then
    BINARY="./ccpm-tui"
elif [ -x "$GOPATH/bin/ccpm-tui" ]; then
    BINARY="$GOPATH/bin/ccpm-tui"
elif [ -x "$HOME/.local/bin/ccpm-tui" ]; then
    BINARY="$HOME/.local/bin/ccpm-tui"
elif command -v ccpm-tui &> /dev/null; then
    BINARY="ccpm-tui"
fi
```

## 2. Parse Arguments

- `--wizard`: Start in wizard mode for new epic creation
- `--epic {name}`: Open dashboard with specific epic selected
- No args: Open main dashboard

## 3. Launch Dashboard

Execute the binary with appropriate flags:
```bash
$BINARY [--wizard] [--epic {name}]
```

## 4. Handle Missing Binary

If binary not found, display:

```
═══════════════════════════════════════════════════════════════
ERROR: TUI Dashboard binary not found
═══════════════════════════════════════════════════════════════

The TUI dashboard requires building the Go binary first.

To build:
  cd ccpm && go build -o ccpm-tui ./cmd/ccpm-tui

To install globally:
  go install github.com/automazeio/ccpm/cmd/ccpm-tui@latest

Note: Requires Go 1.21 or later
═══════════════════════════════════════════════════════════════
```

## 5. Handle Wrong Directory

If .claude/ directory not found:

```
═══════════════════════════════════════════════════════════════
ERROR: Not in a CCPM project
═══════════════════════════════════════════════════════════════

The dashboard requires a .claude/ directory.
Run this command from your project root.

To initialize CCPM:
  /pm:init
═══════════════════════════════════════════════════════════════
```
</process>

<success_criteria>
- Dashboard launches successfully when binary exists
- Wizard mode activates with --wizard flag
- Specific epic selected with --epic flag
- Clear error message if binary missing
- Clear error message if not in CCPM project
</success_criteria>
