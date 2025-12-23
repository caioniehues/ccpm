# CCPM TUI Wireframes & Design System

> **Aesthetic**: Neo-Brutalist Terminal
> **Toolkit**: Charm (Bubbletea, Bubbles, Lipgloss, Huh)
> **Design Date**: 2025-12-23

---

## Design Philosophy

```
╔══════════════════════════════════════════════════════════════════════════════╗
║                                                                              ║
║    "The terminal is not a limitation — it's a canvas with constraints       ║
║     that force creativity. Every character is intentional."                  ║
║                                                                              ║
╚══════════════════════════════════════════════════════════════════════════════╝
```

### Core Principles

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                             │
│   1. BOLD GEOMETRY          Heavy blocks create visual anchors              │
│      █████████████          Not timid lines, but confident shapes           │
│                                                                             │
│   2. ASYMMETRIC TENSION     Break the grid. Create visual flow.             │
│      ┌──────────┐           Offset elements. Surprise the eye.              │
│      │          └────────┐                                                  │
│      └───────────────────┘                                                  │
│                                                                             │
│   3. CHROMATIC RESTRAINT    One dominant. One accent. Black canvas.         │
│      ░░░░░░░░░░░░░░░░░░░    Let color mean something.                       │
│      ░░░░░░████░░░░░░░░░                                                    │
│      ░░░░░░░░░░░░░░░░░░░                                                    │
│                                                                             │
│   4. GENEROUS VOID          Whitespace is not waste.                        │
│                             It's breathing room.                            │
│                             Let elements float.                             │
│                                                                             │
│   5. FUNCTIONAL BEAUTY      Every pixel earns its place.                    │
│      ✓ Status at a glance   No decoration without purpose.                  │
│      → Clear hierarchy                                                      │
│      ◆ Intentional details                                                  │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## Design System

### Color Palette (ANSI True Color)

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  CCPM COLOR SYSTEM                                                          │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  ████████  VOID           #0D0D0D    Background. The canvas.                │
│  ████████  CHARCOAL       #1A1A2E    Elevated surfaces.                     │
│  ████████  GRAPHITE       #2D2D44    Borders, dividers.                     │
│  ████████  SLATE          #4A4A6A    Muted text, inactive.                  │
│  ████████  SILVER         #8888AA    Secondary text.                        │
│  ████████  PEARL          #E8E8F0    Primary text.                          │
│                                                                             │
│  ████████  ELECTRIC       #00D4FF    Primary accent. Progress. Active.      │
│  ████████  PLASMA         #FF006E    Urgent. Errors. Attention.             │
│  ████████  VOLT           #ADFF02    Success. Complete. Go.                 │
│  ████████  AMBER          #FFB800    Warning. In-progress. Caution.         │
│  ████████  LAVENDER       #B388FF    Info. Links. Interactive.              │
│                                                                             │
│  USAGE RULES:                                                               │
│  ─────────────                                                              │
│  • ELECTRIC for primary actions, progress, focus states                     │
│  • VOLT only for success/completion (earned, not given)                     │
│  • PLASMA sparingly — it demands attention                                  │
│  • AMBER for "working on it" states                                         │
│  • LAVENDER for secondary interactive elements                              │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Typography (Monospace Hierarchy)

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  TYPOGRAPHY SYSTEM                                                          │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  ██████╗  ██████╗ ██████╗ ███╗   ███╗                                       │
│  ██╔════╝██╔════╝ ██╔══██╗████╗ ████║     DISPLAY (Figlet-style)           │
│  ██║     ██║      ██████╔╝██╔████╔██║     Hero moments only.               │
│  ██║     ██║      ██╔═══╝ ██║╚██╔╝██║     App title. Major states.         │
│  ╚██████╗╚██████╗ ██║     ██║ ╚═╝ ██║                                       │
│   ╚═════╝ ╚═════╝ ╚═╝     ╚═╝     ╚═╝                                       │
│                                                                             │
│  ═══════════════════════════════════                                        │
│  SECTION HEADERS                         BOLD + UNDERLINE                   │
│  ═══════════════════════════════════     Major divisions.                   │
│                                                                             │
│  ┌─ Component Title ─────────────────┐   BOLD + BOX                         │
│  │                                   │   Card/panel headers.                │
│  └───────────────────────────────────┘                                      │
│                                                                             │
│  LABEL TEXT                              BOLD + CAPS                        │
│  For field labels and categories.        Always uppercase.                  │
│                                                                             │
│  Body text for descriptions and          REGULAR                            │
│  longer content. Comfortable reading.    Default state.                     │
│                                                                             │
│  Secondary information, timestamps       DIM                                │
│  and metadata. Recedes visually.         De-emphasized.                     │
│                                                                             │
│  interactive element                     UNDERLINE + COLOR                  │
│  Clickable/navigable items.              Shows affordance.                  │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Icon System

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  ICON VOCABULARY                                                            │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  STATUS ICONS                                                               │
│  ─────────────                                                              │
│  ◉  Active / Selected / Current         (filled circle)                    │
│  ○  Inactive / Available                 (empty circle)                     │
│  ◐  Partial / In-progress               (half circle)                       │
│  ●  Complete (small)                    (bullet)                            │
│                                                                             │
│  TASK STATUS                                                                │
│  ───────────                                                                │
│  ▣  Completed                           VOLT color                          │
│  ▶  In Progress                         AMBER color, animated               │
│  ▢  Pending                             SLATE color                         │
│  ⊘  Blocked                             PLASMA color                        │
│  ⊗  Failed                              PLASMA color                        │
│                                                                             │
│  NAVIGATION                                                                 │
│  ──────────                                                                 │
│  ❯  Cursor / Selected item              ELECTRIC color                      │
│  ›  Breadcrumb separator                GRAPHITE color                      │
│  ←  Back                                                                    │
│  ↑↓ Scroll indicators                                                       │
│  ⏎  Enter / Confirm                                                         │
│                                                                             │
│  DOCUMENT TYPES                                                             │
│  ──────────────                                                             │
│  ◈  PRD (Product Requirements)          LAVENDER color                      │
│  ◆  Epic                                ELECTRIC color                      │
│  ◇  Task                                SILVER color                        │
│  ◁  Issue (GitHub synced)               AMBER color                         │
│                                                                             │
│  WORKFLOW                                                                   │
│  ────────                                                                   │
│  ✓  Approved                            VOLT color                          │
│  ⟳  Syncing                             ELECTRIC, animated                  │
│  ⚑  Flagged / Important                 PLASMA color                        │
│  ⌂  Home / Dashboard                                                        │
│                                                                             │
│  DECORATIVE                                                                 │
│  ──────────                                                                 │
│  ░▒▓█  Density gradient (progress fills)                                    │
│  ╭╮╰╯  Rounded corners (friendly elements)                                  │
│  ┌┐└┘  Sharp corners (data/serious elements)                                │
│  ═══   Heavy dividers (major sections)                                      │
│  ───   Light dividers (minor sections)                                      │
│  ╌╌╌   Dashed (pending/incomplete)                                          │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## Component Library

### Progress Bars

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  PROGRESS BAR VARIANTS                                                      │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  STANDARD (Default)                                                         │
│  ████████████████████░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░  50%       │
│                                                                             │
│  GRADIENT (Premium feel)                                                    │
│  ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▒▒▒▒▒▒▒▒▒▒░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░  50%       │
│                                                                             │
│  BLOCKS (Bold statement)                                                    │
│  ██████████░░░░░░░░░░  50%                                                  │
│                                                                             │
│  MINIMAL (Subtle)                                                           │
│  ━━━━━━━━━━──────────  50%                                                  │
│                                                                             │
│  SEGMENTED (Task count)                                                     │
│  [■][■][■][□][□][□]  3/6                                                    │
│                                                                             │
│  EPIC HERO (Dashboard feature)                                              │
│  ╔══════════════════════════════════════════════════════════════════════╗  │
│  ║ ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ ║  │
│  ╚══════════════════════════════════════════════════════════════════════╝  │
│                                                    50% Complete (3/6 tasks) │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### List Items

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  LIST ITEM STATES                                                           │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  UNSELECTED                                                                 │
│     ▢  001: Set up authentication middleware                               │
│                                                                             │
│  SELECTED (Cursor here)                                                     │
│  ╔══════════════════════════════════════════════════════════════════════╗  │
│  ║ ❯  ▢  001: Set up authentication middleware                          ║  │
│  ╚══════════════════════════════════════════════════════════════════════╝  │
│                                                                             │
│  IN PROGRESS                                                                │
│     ▶  002: Create user model                            ◐ 40%              │
│                                                                             │
│  SELECTED + IN PROGRESS                                                     │
│  ╔══════════════════════════════════════════════════════════════════════╗  │
│  ║ ❯  ▶  002: Create user model                          ◐ 40%          ║  │
│  ║      └─ src/models/user.ts                                            ║  │
│  ╚══════════════════════════════════════════════════════════════════════╝  │
│                                                                             │
│  COMPLETED                                                                  │
│     ▣  003: Implement login endpoint                     ✓ Done             │
│                                                                             │
│  BLOCKED                                                                    │
│     ⊘  004: Add OAuth integration                        ⚑ Blocked          │
│        └─ Waiting on: API credentials                                       │
│                                                                             │
│  FAILED                                                                     │
│     ⊗  005: Deploy to staging                            ✗ Failed           │
│        └─ Error: Connection timeout                                         │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Cards & Panels

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  PANEL VARIANTS                                                             │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  STANDARD CARD                                                              │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │  Card Title                                                         │   │
│  │  ─────────────────────────────────────────────────────────────────  │   │
│  │  Content goes here. Body text with information.                     │   │
│  │                                                                     │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
│  ELEVATED CARD (Focus state)                                                │
│  ╔═════════════════════════════════════════════════════════════════════╗   │
│  ║  ◆ Active Epic                                                      ║   │
│  ║  ═════════════════════════════════════════════════════════════════  ║   │
│  ║  user-authentication                                                ║   │
│  ║  Progress: ████████████████░░░░░░░░░░░░░░░░░░░░  50%               ║   │
│  ╚═════════════════════════════════════════════════════════════════════╝   │
│                                                                             │
│  SUBTLE CARD (Secondary info)                                               │
│  ╭─────────────────────────────────────────────────────────────────────╮   │
│  │  Activity Log                                                   ↑↓  │   │
│  │  ╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌  │   │
│  │  14:32  Task 003 marked complete                                    │   │
│  │  14:28  Task 002 started                                            │   │
│  ╰─────────────────────────────────────────────────────────────────────╯   │
│                                                                             │
│  ALERT CARD (Attention)                                                     │
│  ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓   │
│  ▓  ⚑ APPROVAL REQUIRED                                               ▓   │
│  ▓  ─────────────────────────────────────────────────────────────────  ▓   │
│  ▓  PRD ready for review. Press [a] to approve or [r] to revise.      ▓   │
│  ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓   │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## View Wireframes

### 1. Main Dashboard

```
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃                                                                              ┃
┃  ██████╗  ██████╗ ██████╗ ███╗   ███╗                                        ┃
┃  ██╔════╝██╔════╝ ██╔══██╗████╗ ████║        Dashboard                       ┃
┃  ██║     ██║      ██████╔╝██╔████╔██║        ═══════════                     ┃
┃  ██║     ██║      ██╔═══╝ ██║╚██╔╝██║        Branch: epic/auth               ┃
┃  ╚██████╗╚██████╗ ██║     ██║ ╚═╝ ██║        Commit: f3a821c                 ┃
┃   ╚═════╝ ╚═════╝ ╚═╝     ╚═╝     ╚═╝                                        ┃
┃                                                                              ┃
┃━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┃
┃                                                                              ┃
┃  ╔═══════════════════════════════════════════════════════════════════════╗  ┃
┃  ║  ◆ ACTIVE EPIC                                                        ║  ┃
┃  ║  ═════════════════════════════════════════════════════════════════    ║  ┃
┃  ║                                                                       ║  ┃
┃  ║  user-authentication                                                  ║  ┃
┃  ║  Implement secure user authentication with JWT and sessions           ║  ┃
┃  ║                                                                       ║  ┃
┃  ║  ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░  ║  ┃
┃  ║                                                                       ║  ┃
┃  ║  PRD ✓        Epic ✓        Tasks ◐        Synced ○                  ║  ┃
┃  ║                                      ▲                                ║  ┃
┃  ║                                      └─ 50% Complete (3/6 tasks)      ║  ┃
┃  ╚═══════════════════════════════════════════════════════════════════════╝  ┃
┃                                                                              ┃
┃                                                                              ┃
┃  ┌─ TASKS ─────────────────────────────────────────────────────────────────┐ ┃
┃  │                                                                         │ ┃
┃  │     ▣  001: Set up auth middleware                        ✓ Done       │ ┃
┃  │     ▣  002: Create user model                             ✓ Done       │ ┃
┃  │     ▣  003: Implement login endpoint                      ✓ Done       │ ┃
┃  │  ╔═══════════════════════════════════════════════════════════════════╗ │ ┃
┃  │  ║ ❯  ▶  004: Add session management                      ◐ 40%     ║ │ ┃
┃  │  ║      └─ Working: src/middleware/session.ts                        ║ │ ┃
┃  │  ╚═══════════════════════════════════════════════════════════════════╝ │ ┃
┃  │     ▢  005: Create registration flow                      ○ Pending    │ ┃
┃  │     ▢  006: Add password reset                            ○ Pending    │ ┃
┃  │                                                                         │ ┃
┃  └─────────────────────────────────────────────────────────────────────────┘ ┃
┃                                                                              ┃
┃                                                                              ┃
┃  ╭─ ACTIVITY ──────────────────────────────────────────────────────────────╮ ┃
┃  │                                                                     ↑↓  │ ┃
┃  │  14:32  ▣ Task 003 completed                                            │ ┃
┃  │  14:28  ▶ Task 004 started                                              │ ┃
┃  │  14:15  ✓ Epic approved for work                                        │ ┃
┃  │  13:45  ◈ PRD created and approved                                      │ ┃
┃  │                                                                          │ ┃
┃  ╰──────────────────────────────────────────────────────────────────────────╯ ┃
┃                                                                              ┃
┃━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┃
┃  [e]pic  [t]ask  [p]rd  [w]izard  [s]ync  [?]help  [q]uit                   ┃
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
```

### 2. Epic Detail View

```
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃                                                                              ┃
┃  ← Dashboard › Epic Details                                                  ┃
┃                                                                              ┃
┃━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┃
┃                                                                              ┃
┃                                                                              ┃
┃     ◆ user-authentication                                                    ┃
┃     ═══════════════════════════════════════════════════════════════          ┃
┃                                                                              ┃
┃     STATUS        In Progress                                                ┃
┃     CREATED       2025-12-23 09:15                                           ┃
┃     BRANCH        epic/auth                                                  ┃
┃     SOURCE PRD    ◈ user-auth-prd.md                                         ┃
┃                                                                              ┃
┃                                                                              ┃
┃  ┌─ EPIC DOCUMENT ─────────────────────────────────────────────────────────┐ ┃
┃  │                                                                     ↑↓  │ ┃
┃  │  # User Authentication Epic                                             │ ┃
┃  │                                                                          │ ┃
┃  │  ## Overview                                                             │ ┃
┃  │  Implement secure user authentication with JWT tokens and               │ ┃
┃  │  session management for the application.                                │ ┃
┃  │                                                                          │ ┃
┃  │  ## Goals                                                                │ ┃
┃  │  - Secure login/logout flow                                             │ ┃
┃  │  - Password hashing with bcrypt                                         │ ┃
┃  │  - JWT token generation and validation                                  │ ┃
┃  │  - Session persistence with Redis                                       │ ┃
┃  │                                                                          │ ┃
┃  │  ## Success Criteria                                                     │ ┃
┃  │  1. Users can register with email/password                              │ ┃
┃  │  2. Users can login and receive JWT                                     │ ┃
┃  │  3. Protected routes require valid token                                │ ┃
┃  │  4. Sessions persist across browser restarts                            │ ┃
┃  │                                                                          │ ┃
┃  │  ─────────────────────────────────────────────────────── [more below]   │ ┃
┃  └──────────────────────────────────────────────────────────────────────────┘ ┃
┃                                                                              ┃
┃                                                                              ┃
┃  ┌─ PHASE STATUS ──────────────────────────────────────────────────────────┐ ┃
┃  │                                                                          │ ┃
┃  │    [■] PRD Created           ──▶  [■] Epic Approved  ──▶  [◐] Tasks     │ ┃
┃  │                                                                          │ ┃
┃  │         ✓ Approved                    ✓ Approved           3/6 done     │ ┃
┃  │                                                                          │ ┃
┃  └──────────────────────────────────────────────────────────────────────────┘ ┃
┃                                                                              ┃
┃                                                                              ┃
┃━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┃
┃  [←]back  [t]asks  [p]rd  [s]ync to github  [m]ark complete                  ┃
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
```

### 3. Task Detail View

```
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃                                                                              ┃
┃  ← Dashboard › Epic › Task Details                                           ┃
┃                                                                              ┃
┃━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┃
┃                                                                              ┃
┃                                                                              ┃
┃     ▶ 004: Add session management                                            ┃
┃     ═══════════════════════════════════════════════════════════════          ┃
┃                                                                              ┃
┃     STATUS        ▓▓▓▓▓▓▓▓░░░░░░░░░░░░  40% In Progress                     ┃
┃     STARTED       2025-12-23 14:28                                           ┃
┃     DEPENDS ON    003: Implement login endpoint                              ┃
┃     BLOCKED BY    None                                                       ┃
┃                                                                              ┃
┃                                                                              ┃
┃  ┌─ TASK DESCRIPTION ──────────────────────────────────────────────────────┐ ┃
┃  │                                                                     ↑↓  │ ┃
┃  │  ## Objective                                                            │ ┃
┃  │  Implement session management with Redis for persistent user             │ ┃
┃  │  sessions across browser restarts.                                       │ ┃
┃  │                                                                          │ ┃
┃  │  ## Acceptance Criteria                                                  │ ┃
┃  │  ▢  Sessions persist in Redis with 7-day TTL                            │ ┃
┃  │  ▢  Session ID stored in HTTP-only cookie                               │ ┃
┃  │  ▣  Session middleware validates on each request                        │ ┃
┃  │  ▣  Logout invalidates session                                          │ ┃
┃  │                                                                          │ ┃
┃  │  ## Files to Modify                                                      │ ┃
┃  │  • src/middleware/session.ts (create)                                   │ ┃
┃  │  • src/routes/auth.ts (update)                                          │ ┃
┃  │  • src/config/redis.ts (update)                                         │ ┃
┃  │                                                                          │ ┃
┃  └──────────────────────────────────────────────────────────────────────────┘ ┃
┃                                                                              ┃
┃                                                                              ┃
┃  ┌─ REQUIREMENTS TRACE ────────────────────────────────────────────────────┐ ┃
┃  │                                                                          │ ┃
┃  │  Linked to:                                                              │ ┃
┃  │    ◈ PRD → Section 3.2: Session Management                              │ ┃
┃  │    ◆ Epic → Goal 4: Session persistence with Redis                      │ ┃
┃  │                                                                          │ ┃
┃  │  Leverages:                                                              │ ┃
┃  │    • src/lib/redis-client.ts (existing)                                 │ ┃
┃  │    • src/types/session.d.ts (existing)                                  │ ┃
┃  │                                                                          │ ┃
┃  └──────────────────────────────────────────────────────────────────────────┘ ┃
┃                                                                              ┃
┃━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┃
┃  [←]back  [c]omplete  [b]lock  [n]ext task  [g]ithub issue                   ┃
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
```

### 4. Wizard Mode (Approval Flow)

```
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃                                                                              ┃
┃                                                                              ┃
┃     ╔═══════════════════════════════════════════════════════════════════╗   ┃
┃     ║                                                                   ║   ┃
┃     ║    EPIC WIZARD                                        Step 2/4   ║   ┃
┃     ║    ════════════════════════════════════════════════════════════   ║   ┃
┃     ║                                                                   ║   ┃
┃     ║    [■]────────[■]────────[◐]────────[○]                          ║   ┃
┃     ║    PRD        Epic       Tasks      Execute                       ║   ┃
┃     ║    ✓ Done     ✓ Done     Review     Pending                       ║   ┃
┃     ║                                                                   ║   ┃
┃     ╚═══════════════════════════════════════════════════════════════════╝   ┃
┃                                                                              ┃
┃                                                                              ┃
┃  ┌─ REVIEW: EPIC DOCUMENT ─────────────────────────────────────────────────┐ ┃
┃  │                                                                     ↑↓  │ ┃
┃  │  # User Authentication Epic                                             │ ┃
┃  │                                                                          │ ┃
┃  │  ## Overview                                                             │ ┃
┃  │  Implement secure user authentication with JWT tokens and               │ ┃
┃  │  session management for the application.                                │ ┃
┃  │                                                                          │ ┃
┃  │  ## Goals                                                                │ ┃
┃  │  1. Secure login/logout flow                                            │ ┃
┃  │  2. Password hashing with bcrypt                                        │ ┃
┃  │  3. JWT token generation and validation                                 │ ┃
┃  │  4. Session persistence with Redis                                      │ ┃
┃  │                                                                          │ ┃
┃  │  ─────────────────────────────────────────────────────── [more below]   │ ┃
┃  └──────────────────────────────────────────────────────────────────────────┘ ┃
┃                                                                              ┃
┃                                                                              ┃
┃  ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓  ┃
┃  ▓                                                                       ▓  ┃
┃  ▓   ? Do you approve this Epic?                                         ▓  ┃
┃  ▓                                                                       ▓  ┃
┃  ▓     › ✓ Yes, proceed to task decomposition                            ▓  ┃
┃  ▓       ✗ No, I need to make changes                                    ▓  ┃
┃  ▓                                                                       ▓  ┃
┃  ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓  ┃
┃                                                                              ┃
┃                                                                              ┃
┃━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┃
┃  [↑/↓]navigate  [⏎]select  [esc]cancel wizard                                ┃
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
```

### 5. Wizard - Revision State

```
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃                                                                              ┃
┃                                                                              ┃
┃     ╔═══════════════════════════════════════════════════════════════════╗   ┃
┃     ║                                                                   ║   ┃
┃     ║    EPIC WIZARD                                        Step 2/4   ║   ┃
┃     ║    ════════════════════════════════════════════════════════════   ║   ┃
┃     ║                                                                   ║   ┃
┃     ║    [■]────────[⟳]────────[○]────────[○]                          ║   ┃
┃     ║    PRD        Epic       Tasks      Execute                       ║   ┃
┃     ║    ✓ Done     Revising   Pending    Pending                       ║   ┃
┃     ║                                                                   ║   ┃
┃     ╚═══════════════════════════════════════════════════════════════════╝   ┃
┃                                                                              ┃
┃                                                                              ┃
┃  ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓  ┃
┃  ▓                                                                       ▓  ┃
┃  ▓   What changes are needed?                                            ▓  ┃
┃  ▓                                                                       ▓  ┃
┃  ▓   ┌───────────────────────────────────────────────────────────────┐   ▓  ┃
┃  ▓   │                                                               │   ▓  ┃
┃  ▓   │  Add OAuth2 integration with Google as a goal. Also need      │   ▓  ┃
┃  ▓   │  to consider rate limiting for login attempts.                │   ▓  ┃
┃  ▓   │                                                               │   ▓  ┃
┃  ▓   │                                                               │   ▓  ┃
┃  ▓   │                                                               │   ▓  ┃
┃  ▓   └───────────────────────────────────────────────────────────────┘   ▓  ┃
┃  ▓                                                                       ▓  ┃
┃  ▓   Characters: 98/500                                                  ▓  ┃
┃  ▓                                                                       ▓  ┃
┃  ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓  ┃
┃                                                                              ┃
┃                                                                              ┃
┃  Tip: Be specific about what needs to change. Claude will revise            ┃
┃       the document and ask for your approval again.                          ┃
┃                                                                              ┃
┃                                                                              ┃
┃━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┃
┃  [⏎]submit feedback  [esc]cancel                                             ┃
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
```

### 6. Multi-Epic Selector

```
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃                                                                              ┃
┃  SELECT EPIC                                                                 ┃
┃  ═══════════════════════════════════════════════════════════════════════════ ┃
┃                                                                              ┃
┃  Filter: █                                                      3 epics     ┃
┃                                                                              ┃
┃                                                                              ┃
┃  ╔══════════════════════════════════════════════════════════════════════════╗┃
┃  ║ ❯  ◆ user-authentication                                                ║┃
┃  ║      │                                                                   ║┃
┃  ║      │  Implement secure user authentication with JWT                    ║┃
┃  ║      │                                                                   ║┃
┃  ║      │  ████████████████████░░░░░░░░░░░░░░░░░░░░  50%                    ║┃
┃  ║      │                                                                   ║┃
┃  ║      │  IN PROGRESS  •  3/6 tasks  •  Updated 2h ago                     ║┃
┃  ║      │                                                                   ║┃
┃  ╚══════════════════════════════════════════════════════════════════════════╝┃
┃                                                                              ┃
┃     ◆ api-rate-limiting                                                      ┃
┃        │                                                                     ┃
┃        │  Add rate limiting to prevent API abuse                             ┃
┃        │                                                                     ┃
┃        │  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░  0%                       ┃
┃        │                                                                     ┃
┃        │  NOT STARTED  •  0/4 tasks  •  Created 1d ago                       ┃
┃        │                                                                     ┃
┃                                                                              ┃
┃     ◆ dashboard-redesign                                                     ┃
┃        │                                                                     ┃
┃        │  Complete overhaul of the admin dashboard                           ┃
┃        │                                                                     ┃
┃        │  ████████████████████████████████████████  100%                     ┃
┃        │                                                                     ┃
┃        │  COMPLETED  •  8/8 tasks  •  Completed 3d ago                       ┃
┃        │                                                                     ┃
┃                                                                              ┃
┃━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┃
┃  [↑/↓]navigate  [⏎]select  [/]filter  [n]ew epic  [esc]cancel                ┃
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
```

### 7. Help Overlay

```
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃                                                                              ┃
┃  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ ┃
┃  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ ┃
┃  ░░░░╔════════════════════════════════════════════════════════════════╗░░░░ ┃
┃  ░░░░║                                                                ║░░░░ ┃
┃  ░░░░║   KEYBOARD SHORTCUTS                                           ║░░░░ ┃
┃  ░░░░║   ═══════════════════════════════════════════════════════════  ║░░░░ ┃
┃  ░░░░║                                                                ║░░░░ ┃
┃  ░░░░║   NAVIGATION                       ACTIONS                     ║░░░░ ┃
┃  ░░░░║   ──────────                       ───────                     ║░░░░ ┃
┃  ░░░░║   ↑ k      Move up                 ⏎       Select/Confirm      ║░░░░ ┃
┃  ░░░░║   ↓ j      Move down               esc     Back/Cancel         ║░░░░ ┃
┃  ░░░░║   ← h      Previous section        /       Filter/Search       ║░░░░ ┃
┃  ░░░░║   → l      Next section            ?       Toggle help         ║░░░░ ┃
┃  ░░░░║   g        Go to top                                           ║░░░░ ┃
┃  ░░░░║   G        Go to bottom            VIEWS                       ║░░░░ ┃
┃  ░░░░║   tab      Next panel              ─────                       ║░░░░ ┃
┃  ░░░░║                                    e       Epic details        ║░░░░ ┃
┃  ░░░░║   WORKFLOW                         t       Task details        ║░░░░ ┃
┃  ░░░░║   ────────                         p       View PRD            ║░░░░ ┃
┃  ░░░░║   w        Launch wizard           a       Activity log        ║░░░░ ┃
┃  ░░░░║   s        Sync to GitHub                                      ║░░░░ ┃
┃  ░░░░║   c        Mark complete           SYSTEM                      ║░░░░ ┃
┃  ░░░░║   b        Mark blocked            ──────                      ║░░░░ ┃
┃  ░░░░║   n        Next task               r       Refresh             ║░░░░ ┃
┃  ░░░░║                                    q       Quit                ║░░░░ ┃
┃  ░░░░║                                                                ║░░░░ ┃
┃  ░░░░╚════════════════════════════════════════════════════════════════╝░░░░ ┃
┃  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ ┃
┃  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ ┃
┃                                                                              ┃
┃  Press any key to close                                                      ┃
┃                                                                              ┃
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
```

### 8. Loading State

```
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃                                                                              ┃
┃                                                                              ┃
┃                                                                              ┃
┃                                                                              ┃
┃                                                                              ┃
┃                                                                              ┃
┃                                                                              ┃
┃                                                                              ┃
┃                                                                              ┃
┃                           ⣾⣽⣻⢿⡿⣟⣯⣷                                          ┃
┃                                                                              ┃
┃                      Syncing with GitHub...                                  ┃
┃                                                                              ┃
┃                      Pushing 6 issues to repository                          ┃
┃                                                                              ┃
┃                      ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░                          ┃
┃                                                                              ┃
┃                                                                              ┃
┃                                                                              ┃
┃                                                                              ┃
┃                                                                              ┃
┃                                                                              ┃
┃                                                                              ┃
┃                                                                              ┃
┃                                                                              ┃
┃━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┃
┃  [esc] Cancel                                                                ┃
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
```

### 9. Error State

```
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃                                                                              ┃
┃  ← Dashboard                                                                 ┃
┃                                                                              ┃
┃━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┃
┃                                                                              ┃
┃                                                                              ┃
┃                                                                              ┃
┃                                                                              ┃
┃  ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓  ┃
┃  ▓                                                                       ▓  ┃
┃  ▓   ⊗ SYNC FAILED                                                       ▓  ┃
┃  ▓   ═══════════════════════════════════════════════════════════════     ▓  ┃
┃  ▓                                                                       ▓  ┃
┃  ▓   Unable to sync with GitHub repository.                              ▓  ┃
┃  ▓                                                                       ▓  ┃
┃  ▓   Error: Authentication failed. Your GitHub token may have expired.  ▓  ┃
┃  ▓                                                                       ▓  ┃
┃  ▓   ─────────────────────────────────────────────────────────────────   ▓  ┃
┃  ▓                                                                       ▓  ┃
┃  ▓   Suggestions:                                                        ▓  ┃
┃  ▓   • Run 'gh auth login' to re-authenticate                           ▓  ┃
┃  ▓   • Check your network connection                                    ▓  ┃
┃  ▓   • Verify repository access permissions                             ▓  ┃
┃  ▓                                                                       ▓  ┃
┃  ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓  ┃
┃                                                                              ┃
┃                                                                              ┃
┃                                                                              ┃
┃                                                                              ┃
┃━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┃
┃  [r]etry  [←]back to dashboard  [c]opy error  [?]help                        ┃
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
```

### 10. Empty State

```
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃                                                                              ┃
┃  ██████╗  ██████╗ ██████╗ ███╗   ███╗                                        ┃
┃  ██╔════╝██╔════╝ ██╔══██╗████╗ ████║        Dashboard                       ┃
┃  ██║     ██║      ██████╔╝██╔████╔██║        ═══════════                     ┃
┃  ██║     ██║      ██╔═══╝ ██║╚██╔╝██║                                        ┃
┃  ╚██████╗╚██████╗ ██║     ██║ ╚═╝ ██║                                        ┃
┃   ╚═════╝ ╚═════╝ ╚═╝     ╚═╝     ╚═╝                                        ┃
┃                                                                              ┃
┃━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┃
┃                                                                              ┃
┃                                                                              ┃
┃                                                                              ┃
┃                                                                              ┃
┃                                                                              ┃
┃                                                                              ┃
┃                         ╭───────────────────────╮                            ┃
┃                         │                       │                            ┃
┃                         │    ◇ No epics yet     │                            ┃
┃                         │                       │                            ┃
┃                         │    Start by creating  │                            ┃
┃                         │    your first PRD     │                            ┃
┃                         │                       │                            ┃
┃                         │    Press [w] to       │                            ┃
┃                         │    launch the wizard  │                            ┃
┃                         │                       │                            ┃
┃                         ╰───────────────────────╯                            ┃
┃                                                                              ┃
┃                                                                              ┃
┃                                                                              ┃
┃                                                                              ┃
┃                                                                              ┃
┃━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┃
┃  [w]izard  [i]mport from github  [?]help  [q]uit                             ┃
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
```

### 11. PRD Detail View

```
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃                                                                              ┃
┃  ← Dashboard › PRD Details                                                   ┃
┃                                                                              ┃
┃━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┃
┃                                                                              ┃
┃                                                                              ┃
┃     ◈ user-authentication PRD                                                ┃
┃     ═══════════════════════════════════════════════════════════════          ┃
┃                                                                              ┃
┃     STATUS        ✅ APPROVED                                                ┃
┃     CREATED       2025-12-23 09:00                                           ┃
┃     APPROVED      2025-12-23 09:15                                           ┃
┃     LINKED EPIC   ◆ user-authentication                                      ┃
┃                                                                              ┃
┃                                                                              ┃
┃  ┌─ PRD DOCUMENT ─────────────────────────────────────────────────────────┐  ┃
┃  │                                                                     ↑↓  │  ┃
┃  │  # User Authentication PRD                                              │  ┃
┃  │                                                                          │  ┃
┃  │  ## Executive Summary                                                    │  ┃
┃  │  Implement secure user authentication with JWT tokens and                │  ┃
┃  │  session management for the application.                                 │  ┃
┃  │                                                                          │  ┃
┃  │  ## Problem Statement                                                    │  ┃
┃  │  Users cannot currently authenticate to access protected                 │  ┃
┃  │  resources in the application.                                           │  ┃
┃  │                                                                          │  ┃
┃  │  ## Goals                                                                │  ┃
┃  │  1. Secure login/logout flow                                             │  ┃
┃  │  2. Password hashing with bcrypt                                         │  ┃
┃  │  3. JWT token generation and validation                                  │  ┃
┃  │  4. Session persistence with Redis                                       │  ┃
┃  │                                                                          │  ┃
┃  │  ## Non-Goals                                                            │  ┃
┃  │  - Social login (OAuth) - future phase                                   │  ┃
┃  │  - Two-factor authentication - future phase                              │  ┃
┃  │                                                                          │  ┃
┃  │  ─────────────────────────────────────────────────────── [more below]    │  ┃
┃  └──────────────────────────────────────────────────────────────────────────┘  ┃
┃                                                                              ┃
┃                                                                              ┃
┃  ┌─ LINKED ARTIFACTS ─────────────────────────────────────────────────────┐  ┃
┃  │                                                                          │  ┃
┃  │    ◆ Epic: user-authentication          50% (3/6 tasks)                  │  ┃
┃  │    ◇ Tasks: 6 total                     3 complete, 1 in progress        │  ┃
┃  │                                                                          │  ┃
┃  └──────────────────────────────────────────────────────────────────────────┘  ┃
┃                                                                              ┃
┃━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┃
┃  [←]back  [e]pic  [t]asks  [r]evise prd  [d]elete                            ┃
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
```

### 12. Sync Confirmation Dialog

```
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃                                                                              ┃
┃  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ ┃
┃  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ ┃
┃  ░░░░╔════════════════════════════════════════════════════════════════╗░░░░ ┃
┃  ░░░░║                                                                ║░░░░ ┃
┃  ░░░░║   ⟳ SYNC TO GITHUB                                             ║░░░░ ┃
┃  ░░░░║   ═══════════════════════════════════════════════════════════  ║░░░░ ┃
┃  ░░░░║                                                                ║░░░░ ┃
┃  ░░░░║   This will push your epic and tasks to GitHub:                ║░░░░ ┃
┃  ░░░░║                                                                ║░░░░ ┃
┃  ░░░░║   Repository:  user/my-project                                 ║░░░░ ┃
┃  ░░░░║   Epic:        user-authentication                             ║░░░░ ┃
┃  ░░░░║   Tasks:       6 issues will be created                        ║░░░░ ┃
┃  ░░░░║                                                                ║░░░░ ┃
┃  ░░░░║   ─────────────────────────────────────────────────────────    ║░░░░ ┃
┃  ░░░░║                                                                ║░░░░ ┃
┃  ░░░░║   Changes to sync:                                             ║░░░░ ┃
┃  ░░░░║                                                                ║░░░░ ┃
┃  ░░░░║   ┌──────────────────────────────────────────────────────┐     ║░░░░ ┃
┃  ░░░░║   │  + Create issue: 001 - Set up auth middleware        │     ║░░░░ ┃
┃  ░░░░║   │  + Create issue: 002 - Create user model             │     ║░░░░ ┃
┃  ░░░░║   │  + Create issue: 003 - Implement login endpoint      │     ║░░░░ ┃
┃  ░░░░║   │  + Create issue: 004 - Add session management        │     ║░░░░ ┃
┃  ░░░░║   │  + Create issue: 005 - Create registration flow      │     ║░░░░ ┃
┃  ░░░░║   │  + Create issue: 006 - Add password reset            │     ║░░░░ ┃
┃  ░░░░║   └──────────────────────────────────────────────────────┘     ║░░░░ ┃
┃  ░░░░║                                                                ║░░░░ ┃
┃  ░░░░║       ┌─────────────────┐     ┌─────────────────┐              ║░░░░ ┃
┃  ░░░░║       │   ✓ Confirm     │     │   ✗ Cancel      │              ║░░░░ ┃
┃  ░░░░║       └─────────────────┘     └─────────────────┘              ║░░░░ ┃
┃  ░░░░║                                                                ║░░░░ ┃
┃  ░░░░╚════════════════════════════════════════════════════════════════╝░░░░ ┃
┃  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ ┃
┃  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ ┃
┃                                                                              ┃
┃  [⏎] Confirm sync  [esc] Cancel                                              ┃
┃                                                                              ┃
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
```

### 13. Notification Toasts

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  NOTIFICATION TOAST VARIANTS                                                 │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                              │
│  SUCCESS TOAST (VOLT color, auto-dismiss 3s)                                 │
│  ┌──────────────────────────────────────────────────────────────────────┐   │
│  │  ✓ Task 003 marked complete                                      ✗  │   │
│  └──────────────────────────────────────────────────────────────────────┘   │
│                                                                              │
│  ERROR TOAST (PLASMA color, requires dismiss)                                │
│  ┌──────────────────────────────────────────────────────────────────────┐   │
│  │  ⊗ Sync failed: Authentication expired. Run 'gh auth login'     ✗  │   │
│  └──────────────────────────────────────────────────────────────────────┘   │
│                                                                              │
│  WARNING TOAST (AMBER color, auto-dismiss 5s)                                │
│  ┌──────────────────────────────────────────────────────────────────────┐   │
│  │  ⚠ Epic has unsaved changes. Press [s] to sync                   ✗  │   │
│  └──────────────────────────────────────────────────────────────────────┘   │
│                                                                              │
│  INFO TOAST (LAVENDER color, auto-dismiss 3s)                                │
│  ┌──────────────────────────────────────────────────────────────────────┐   │
│  │  ℹ File changed: epic.md reloaded                                ✗  │   │
│  └──────────────────────────────────────────────────────────────────────┘   │
│                                                                              │
│  TOAST POSITIONING (Top-right corner)                                        │
│  ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓  │
│  ┃  Dashboard                         ┌───────────────────────────────┐ ┃  │
│  ┃                                    │ ✓ Synced to GitHub        ✗  │ ┃  │
│  ┃                                    └───────────────────────────────┘ ┃  │
│  ┃  ┌──────────────────────────────┐                                    ┃  │
│  ┃  │  Epic content...             │                                    ┃  │
│  ┃  │                              │                                    ┃  │
│  ┃  └──────────────────────────────┘                                    ┃  │
│  ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛  │
│                                                                              │
│  STACKED TOASTS (Multiple notifications)                                     │
│                                          ┌───────────────────────────────┐  │
│                                          │ ✓ Task 003 complete       ✗  │  │
│                                          ├───────────────────────────────┤  │
│                                          │ ⚠ 2 tasks blocked         ✗  │  │
│                                          └───────────────────────────────┘  │
│                                                                              │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 14. Wizard - Step 1: PRD Review

```
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃                                                                              ┃
┃                                                                              ┃
┃     ╔═══════════════════════════════════════════════════════════════════╗   ┃
┃     ║                                                                   ║   ┃
┃     ║    EPIC WIZARD                                        Step 1/4   ║   ┃
┃     ║    ════════════════════════════════════════════════════════════   ║   ┃
┃     ║                                                                   ║   ┃
┃     ║    [◐]────────[○]────────[○]────────[○]                          ║   ┃
┃     ║    PRD        Epic       Tasks      Sync                         ║   ┃
┃     ║    Review     Pending    Pending    Pending                       ║   ┃
┃     ║                                                                   ║   ┃
┃     ╚═══════════════════════════════════════════════════════════════════╝   ┃
┃                                                                              ┃
┃                                                                              ┃
┃  ┌─ REVIEW: PRD DOCUMENT ─────────────────────────────────────────────────┐  ┃
┃  │                                                                     ↑↓  │  ┃
┃  │  # User Authentication PRD                                              │  ┃
┃  │                                                                          │  ┃
┃  │  ## Executive Summary                                                    │  ┃
┃  │  Implement secure user authentication with JWT tokens and                │  ┃
┃  │  session management for the application.                                 │  ┃
┃  │                                                                          │  ┃
┃  │  ## Problem Statement                                                    │  ┃
┃  │  Users cannot currently authenticate to access protected                 │  ┃
┃  │  resources in the application.                                           │  ┃
┃  │                                                                          │  ┃
┃  │  ─────────────────────────────────────────────────────── [more below]    │  ┃
┃  └──────────────────────────────────────────────────────────────────────────┘  ┃
┃                                                                              ┃
┃                                                                              ┃
┃  ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓  ┃
┃  ▓                                                                       ▓  ┃
┃  ▓   ? Do you approve this PRD?                                          ▓  ┃
┃  ▓                                                                       ▓  ┃
┃  ▓     › ✓ Yes, proceed to Epic generation                               ▓  ┃
┃  ▓       ✗ No, I need to make changes                                    ▓  ┃
┃  ▓                                                                       ▓  ┃
┃  ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓  ┃
┃                                                                              ┃
┃                                                                              ┃
┃━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┃
┃  [↑/↓]navigate  [⏎]select  [esc]cancel wizard                                ┃
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
```

### 15. Wizard - Step 3: Tasks Review

```
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃                                                                              ┃
┃                                                                              ┃
┃     ╔═══════════════════════════════════════════════════════════════════╗   ┃
┃     ║                                                                   ║   ┃
┃     ║    EPIC WIZARD                                        Step 3/4   ║   ┃
┃     ║    ════════════════════════════════════════════════════════════   ║   ┃
┃     ║                                                                   ║   ┃
┃     ║    [■]────────[■]────────[◐]────────[○]                          ║   ┃
┃     ║    PRD        Epic       Tasks      Sync                         ║   ┃
┃     ║    ✓ Done     ✓ Done     Review     Pending                       ║   ┃
┃     ║                                                                   ║   ┃
┃     ╚═══════════════════════════════════════════════════════════════════╝   ┃
┃                                                                              ┃
┃                                                                              ┃
┃  ┌─ REVIEW: 6 TASKS GENERATED ────────────────────────────────────────────┐  ┃
┃  │                                                                     ↑↓  │  ┃
┃  │  ▢  001: Set up authentication middleware                               │  ┃
┃  │      └─ Create Express middleware for JWT validation                    │  ┃
┃  │                                                                          │  ┃
┃  │  ▢  002: Create user model                                              │  ┃
┃  │      └─ Define User schema with password hashing                        │  ┃
┃  │                                                                          │  ┃
┃  │  ▢  003: Implement login endpoint                                       │  ┃
┃  │      └─ POST /auth/login with JWT response                              │  ┃
┃  │                                                                          │  ┃
┃  │  ▢  004: Add session management                                         │  ┃
┃  │      └─ Redis-backed session storage                                    │  ┃
┃  │                                                                          │  ┃
┃  │  ▢  005: Create registration flow                                       │  ┃
┃  │      └─ POST /auth/register with validation                             │  ┃
┃  │                                                                          │  ┃
┃  │  ▢  006: Add password reset                                             │  ┃
┃  │      └─ Email-based password reset flow                                 │  ┃
┃  │                                                                          │  ┃
┃  └──────────────────────────────────────────────────────────────────────────┘  ┃
┃                                                                              ┃
┃                                                                              ┃
┃  ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓  ┃
┃  ▓                                                                       ▓  ┃
┃  ▓   ? Do you approve these 6 tasks?                                     ▓  ┃
┃  ▓                                                                       ▓  ┃
┃  ▓     › ✓ Yes, ready to begin work                                      ▓  ┃
┃  ▓       ✗ No, adjust the task breakdown                                 ▓  ┃
┃  ▓                                                                       ▓  ┃
┃  ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓  ┃
┃                                                                              ┃
┃                                                                              ┃
┃━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┃
┃  [↑/↓]navigate  [⏎]select  [esc]cancel wizard                                ┃
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
```

### 16. Wizard - Step 4: GitHub Sync Choice

```
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃                                                                              ┃
┃                                                                              ┃
┃     ╔═══════════════════════════════════════════════════════════════════╗   ┃
┃     ║                                                                   ║   ┃
┃     ║    EPIC WIZARD                                        Step 4/4   ║   ┃
┃     ║    ════════════════════════════════════════════════════════════   ║   ┃
┃     ║                                                                   ║   ┃
┃     ║    [■]────────[■]────────[■]────────[◐]                          ║   ┃
┃     ║    PRD        Epic       Tasks      Sync                         ║   ┃
┃     ║    ✓ Done     ✓ Done     ✓ Done     Choose                        ║   ┃
┃     ║                                                                   ║   ┃
┃     ╚═══════════════════════════════════════════════════════════════════╝   ┃
┃                                                                              ┃
┃                                                                              ┃
┃  ┌─ EPIC SUMMARY ─────────────────────────────────────────────────────────┐  ┃
┃  │                                                                          │  ┃
┃  │  ◆ user-authentication                                                   │  ┃
┃  │                                                                          │  ┃
┃  │  PRD:    .claude/prds/user-authentication.md          ✅ APPROVED        │  ┃
┃  │  Epic:   .claude/epics/user-authentication/epic.md    ✅ APPROVED        │  ┃
┃  │  Tasks:  6 tasks ready for execution                  ✅ APPROVED        │  ┃
┃  │                                                                          │  ┃
┃  │  ─────────────────────────────────────────────────────────────────────   │  ┃
┃  │                                                                          │  ┃
┃  │  Ready to begin work! You can now start executing tasks.                 │  ┃
┃  │                                                                          │  ┃
┃  └──────────────────────────────────────────────────────────────────────────┘  ┃
┃                                                                              ┃
┃                                                                              ┃
┃  ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓  ┃
┃  ▓                                                                       ▓  ┃
┃  ▓   ? Sync epic to GitHub?                                              ▓  ┃
┃  ▓                                                                       ▓  ┃
┃  ▓     › ⟳ Yes, create GitHub issues now                                 ▓  ┃
┃  ▓       ○ No, work locally only                                         ▓  ┃
┃  ▓       ⏭ Skip for now, I'll sync later                                 ▓  ┃
┃  ▓                                                                       ▓  ┃
┃  ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓  ┃
┃                                                                              ┃
┃                                                                              ┃
┃  Tip: Working locally is great for offline development. You can              ┃
┃       sync to GitHub anytime with [s] from the dashboard.                    ┃
┃                                                                              ┃
┃━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┃
┃  [↑/↓]navigate  [⏎]select  [esc]cancel wizard                                ┃
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
```

### 17. Settings/Preferences View

```
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃                                                                              ┃
┃  ← Dashboard › Settings                                                      ┃
┃                                                                              ┃
┃━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┃
┃                                                                              ┃
┃                                                                              ┃
┃     ⚙ Settings                                                               ┃
┃     ═══════════════════════════════════════════════════════════════          ┃
┃                                                                              ┃
┃                                                                              ┃
┃  ┌─ APPEARANCE ───────────────────────────────────────────────────────────┐  ┃
┃  │                                                                          │  ┃
┃  │  Theme                                                                   │  ┃
┃  │  ╔═════════════════════════════════════════════════════════════════╗    │  ┃
┃  │  ║ › Neo-Brutalist (Default)                                       ║    │  ┃
┃  │  ║   Minimal Dark                                                  ║    │  ┃
┃  │  ║   High Contrast                                                 ║    │  ┃
┃  │  ╚═════════════════════════════════════════════════════════════════╝    │  ┃
┃  │                                                                          │  ┃
┃  │  Animation Speed        [■■■■■░░░░░]  Medium                             │  ┃
┃  │  Show Icons             [■] Enabled                                      │  ┃
┃  │                                                                          │  ┃
┃  └──────────────────────────────────────────────────────────────────────────┘  ┃
┃                                                                              ┃
┃  ┌─ GITHUB ───────────────────────────────────────────────────────────────┐  ┃
┃  │                                                                          │  ┃
┃  │  Repository         user/my-project                                      │  ┃
┃  │  Auth Status        ✅ Authenticated                                     │  ┃
┃  │  Auto-sync          [□] Disabled                                         │  ┃
┃  │                                                                          │  ┃
┃  │  [Re-authenticate]  [Clear credentials]                                  │  ┃
┃  │                                                                          │  ┃
┃  └──────────────────────────────────────────────────────────────────────────┘  ┃
┃                                                                              ┃
┃  ┌─ BEHAVIOR ─────────────────────────────────────────────────────────────┐  ┃
┃  │                                                                          │  ┃
┃  │  File Watching      [■] Enabled (auto-reload on changes)                 │  ┃
┃  │  Confirm on Exit    [■] Enabled                                          │  ┃
┃  │  Toast Duration     [■■■■░░░░░░]  3 seconds                              │  ┃
┃  │                                                                          │  ┃
┃  └──────────────────────────────────────────────────────────────────────────┘  ┃
┃                                                                              ┃
┃━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┃
┃  [←]back  [⏎]toggle  [r]eset defaults                                        ┃
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
```

### 18. Search/Filter Modal

```
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃                                                                              ┃
┃  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ ┃
┃  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ ┃
┃  ░░░░╔════════════════════════════════════════════════════════════════╗░░░░ ┃
┃  ░░░░║                                                                ║░░░░ ┃
┃  ░░░░║   🔍 SEARCH                                                    ║░░░░ ┃
┃  ░░░░║   ═══════════════════════════════════════════════════════════  ║░░░░ ┃
┃  ░░░░║                                                                ║░░░░ ┃
┃  ░░░░║   ┌──────────────────────────────────────────────────────────┐ ║░░░░ ┃
┃  ░░░░║   │ auth█                                                    │ ║░░░░ ┃
┃  ░░░░║   └──────────────────────────────────────────────────────────┘ ║░░░░ ┃
┃  ░░░░║                                                                ║░░░░ ┃
┃  ░░░░║   FILTERS                                                      ║░░░░ ┃
┃  ░░░░║   ────────                                                     ║░░░░ ┃
┃  ░░░░║                                                                ║░░░░ ┃
┃  ░░░░║   Type:     [■] Epics  [■] Tasks  [■] PRDs                    ║░░░░ ┃
┃  ░░░░║   Status:   [■] All    [□] Pending  [□] In Progress  [□] Done ║░░░░ ┃
┃  ░░░░║                                                                ║░░░░ ┃
┃  ░░░░║   ─────────────────────────────────────────────────────────    ║░░░░ ┃
┃  ░░░░║                                                                ║░░░░ ┃
┃  ░░░░║   RESULTS (3 matches)                                          ║░░░░ ┃
┃  ░░░░║   ────────────────────                                         ║░░░░ ┃
┃  ░░░░║                                                                ║░░░░ ┃
┃  ░░░░║   ❯ ◆ user-authentication                     Epic            ║░░░░ ┃
┃  ░░░░║       └─ Implement secure user authentication...              ║░░░░ ┃
┃  ░░░░║                                                                ║░░░░ ┃
┃  ░░░░║     ◇ 001: Set up auth middleware              Task            ║░░░░ ┃
┃  ░░░░║       └─ Create Express middleware for JWT...                 ║░░░░ ┃
┃  ░░░░║                                                                ║░░░░ ┃
┃  ░░░░║     ◈ user-auth-prd                            PRD             ║░░░░ ┃
┃  ░░░░║       └─ Product requirements for authentication...           ║░░░░ ┃
┃  ░░░░║                                                                ║░░░░ ┃
┃  ░░░░╚════════════════════════════════════════════════════════════════╝░░░░ ┃
┃  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ ┃
┃                                                                              ┃
┃  [↑/↓]navigate  [⏎]select  [tab]toggle filter  [esc]close                    ┃
┃                                                                              ┃
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
```

---

## Navigation Flow

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                           NAVIGATION STATE MACHINE                           │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                              │
│                           ┌─────────────────┐                                │
│                           │                 │                                │
│              ┌────────────│    DASHBOARD    │────────────┐                   │
│              │            │    (Default)    │            │                   │
│              │            └────────┬────────┘            │                   │
│              │                     │                     │                   │
│         [e] Epic             [t] Task              [w] Wizard                │
│              │                     │                     │                   │
│              ▼                     ▼                     ▼                   │
│  ┌─────────────────┐   ┌─────────────────┐   ┌─────────────────┐            │
│  │                 │   │                 │   │                 │            │
│  │   EPIC DETAIL   │   │   TASK DETAIL   │   │   WIZARD MODE   │            │
│  │                 │   │                 │   │   (Multi-step)  │            │
│  └────────┬────────┘   └────────┬────────┘   └────────┬────────┘            │
│           │                     │                     │                      │
│           │                     │                     │                      │
│      [←] or [esc]          [←] or [esc]          [esc] Cancel                │
│           │                     │                     │                      │
│           └──────────────────────┴──────────────────────┘                     │
│                                 │                                            │
│                                 ▼                                            │
│                    ┌─────────────────────┐                                   │
│                    │                     │                                   │
│                    │     BACK TO         │                                   │
│                    │     DASHBOARD       │                                   │
│                    │                     │                                   │
│                    └─────────────────────┘                                   │
│                                                                              │
│                                                                              │
│  OVERLAYS (Can appear on any view)                                           │
│  ──────────────────────────────────                                          │
│                                                                              │
│  ┌─────────────┐      ┌─────────────┐      ┌─────────────┐                   │
│  │   [?] Help  │      │ Epic Select │      │   Loading   │                   │
│  │   Overlay   │      │    Modal    │      │   Spinner   │                   │
│  └─────────────┘      └─────────────┘      └─────────────┘                   │
│                                                                              │
│  Dismiss: any key     Dismiss: [esc]      Auto-dismiss                       │
│                                                                              │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## Responsive Behavior

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                          TERMINAL SIZE HANDLING                              │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                              │
│  MINIMUM SIZE: 80 columns × 24 rows                                          │
│                                                                              │
│  ┌──────────────────────────────────────────────────────────────────────┐   │
│  │                                                                      │   │
│  │  SMALL (80-99 cols)           MEDIUM (100-119 cols)                 │   │
│  │  ────────────────────         ──────────────────────                │   │
│  │  • Single column layout       • Task list gains more width          │   │
│  │  • Abbreviated labels         • Full labels                         │   │
│  │  • Compact progress bars      • Standard progress bars              │   │
│  │  • Activity hidden            • Activity visible                    │   │
│  │                                                                      │   │
│  │                                                                      │   │
│  │  LARGE (120+ cols)            EXTRA TALL (40+ rows)                 │   │
│  │  ─────────────────            ──────────────────────                │   │
│  │  • Side-by-side panels        • Expanded viewports                  │   │
│  │  • Epic + Tasks visible       • More task items visible             │   │
│  │  • Full activity log          • Full activity history               │   │
│  │  • Extra metadata             • Additional metadata                 │   │
│  │                                                                      │   │
│  └──────────────────────────────────────────────────────────────────────┘   │
│                                                                              │
│  SIZE TOO SMALL MESSAGE:                                                     │
│                                                                              │
│  ┌──────────────────────────────────────────────────────────────────────┐   │
│  │                                                                      │   │
│  │     ╭─────────────────────────────────────╮                          │   │
│  │     │                                     │                          │   │
│  │     │   Terminal too small                │                          │   │
│  │     │                                     │                          │   │
│  │     │   Minimum: 80×24                    │                          │   │
│  │     │   Current: 60×20                    │                          │   │
│  │     │                                     │                          │   │
│  │     │   Please resize your terminal       │                          │   │
│  │     │                                     │                          │   │
│  │     ╰─────────────────────────────────────╯                          │   │
│  │                                                                      │   │
│  └──────────────────────────────────────────────────────────────────────┘   │
│                                                                              │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## Animation Specifications

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                            ANIMATION SPECS                                   │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                              │
│  SPINNER (Loading states)                                                    │
│  ─────────────────────────                                                  │
│  Frames: ⠋ ⠙ ⠹ ⠸ ⠼ ⠴ ⠦ ⠧ ⠇ ⠏                                               │
│  Speed: 80ms per frame                                                       │
│  Color: ELECTRIC (cyan)                                                      │
│                                                                              │
│  PROGRESS BAR (Smooth fill)                                                  │
│  ──────────────────────────                                                 │
│  Animation: Gradient sweep on update                                         │
│  Duration: 200ms                                                             │
│  Easing: ease-out                                                            │
│                                                                              │
│  CURSOR (Selection indicator)                                                │
│  ────────────────────────────                                               │
│  Frames: ❯ (static with blink)                                              │
│  Blink rate: 530ms                                                           │
│  Color: ELECTRIC                                                             │
│                                                                              │
│  IN-PROGRESS INDICATOR                                                       │
│  ─────────────────────                                                      │
│  Frames: ▶ (pulsing brightness)                                             │
│  Pulse rate: 1000ms                                                          │
│  Color: AMBER → dim → AMBER                                                  │
│                                                                              │
│  SYNC INDICATOR                                                              │
│  ──────────────                                                             │
│  Frames: ⟳ (rotating)                                                        │
│  Rotation: Character substitution ⟳ → ↻ → ⟳                                  │
│  Speed: 150ms per frame                                                      │
│  Color: ELECTRIC                                                             │
│                                                                              │
│  SUCCESS FLASH                                                               │
│  ─────────────                                                              │
│  Effect: Brief green highlight on row                                        │
│  Duration: 300ms                                                             │
│  Color: VOLT background, then fade                                           │
│                                                                              │
│  ERROR SHAKE                                                                 │
│  ───────────                                                                │
│  Effect: Horizontal jitter of error card                                     │
│  Duration: 200ms                                                             │
│  Movement: ±2 characters, 3 cycles                                           │
│                                                                              │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## Lipgloss Style Definitions

```go
// styles.go - Complete Lipgloss style definitions

package tui

import "github.com/charmbracelet/lipgloss"

// Color Palette
var (
    Void      = lipgloss.Color("#0D0D0D")
    Charcoal  = lipgloss.Color("#1A1A2E")
    Graphite  = lipgloss.Color("#2D2D44")
    Slate     = lipgloss.Color("#4A4A6A")
    Silver    = lipgloss.Color("#8888AA")
    Pearl     = lipgloss.Color("#E8E8F0")

    Electric  = lipgloss.Color("#00D4FF")
    Plasma    = lipgloss.Color("#FF006E")
    Volt      = lipgloss.Color("#ADFF02")
    Amber     = lipgloss.Color("#FFB800")
    Lavender  = lipgloss.Color("#B388FF")
)

// Base Styles
var (
    BaseStyle = lipgloss.NewStyle().
        Background(Void).
        Foreground(Pearl)

    MutedStyle = lipgloss.NewStyle().
        Foreground(Slate)

    AccentStyle = lipgloss.NewStyle().
        Foreground(Electric).
        Bold(true)
)

// Container Styles
var (
    AppFrame = lipgloss.NewStyle().
        Border(lipgloss.ThickBorder()).
        BorderForeground(Graphite).
        Padding(1, 2)

    ElevatedCard = lipgloss.NewStyle().
        Border(lipgloss.DoubleBorder()).
        BorderForeground(Electric).
        Padding(1, 2).
        MarginBottom(1)

    SubtleCard = lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder()).
        BorderForeground(Graphite).
        Padding(1, 2)

    AlertCard = lipgloss.NewStyle().
        Background(Charcoal).
        Border(lipgloss.ThickBorder()).
        BorderForeground(Plasma).
        Padding(1, 2)
)

// Typography Styles
var (
    TitleStyle = lipgloss.NewStyle().
        Foreground(Pearl).
        Bold(true).
        MarginBottom(1)

    SectionHeader = lipgloss.NewStyle().
        Foreground(Electric).
        Bold(true).
        Underline(true)

    LabelStyle = lipgloss.NewStyle().
        Foreground(Silver).
        Bold(true)

    BodyStyle = lipgloss.NewStyle().
        Foreground(Pearl)

    DimStyle = lipgloss.NewStyle().
        Foreground(Slate).
        Italic(true)
)

// Status Styles
var (
    SuccessStyle = lipgloss.NewStyle().
        Foreground(Volt).
        Bold(true)

    WarningStyle = lipgloss.NewStyle().
        Foreground(Amber).
        Bold(true)

    ErrorStyle = lipgloss.NewStyle().
        Foreground(Plasma).
        Bold(true)

    InfoStyle = lipgloss.NewStyle().
        Foreground(Lavender)
)

// Interactive Styles
var (
    SelectedItem = lipgloss.NewStyle().
        Border(lipgloss.DoubleBorder()).
        BorderForeground(Electric).
        Padding(0, 1)

    FocusedInput = lipgloss.NewStyle().
        Border(lipgloss.NormalBorder()).
        BorderForeground(Electric).
        Padding(0, 1)

    HelpKeyStyle = lipgloss.NewStyle().
        Foreground(Electric).
        Bold(true)

    HelpDescStyle = lipgloss.NewStyle().
        Foreground(Silver)
)

// Progress Styles
var (
    ProgressFilled = lipgloss.NewStyle().
        Foreground(Electric)

    ProgressEmpty = lipgloss.NewStyle().
        Foreground(Graphite)

    ProgressComplete = lipgloss.NewStyle().
        Foreground(Volt)
)

// Statusbar Style
var (
    StatusBar = lipgloss.NewStyle().
        Background(Charcoal).
        Foreground(Silver).
        Padding(0, 1).
        Width(100) // Full width
)
```

---

## Implementation Checklist

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                        IMPLEMENTATION CHECKLIST                              │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                              │
│  PHASE 1: Core Structure                                                     │
│  ─────────────────────────                                                  │
│  ▢  Set up Go project with Charm dependencies                               │
│  ▢  Implement Model struct with all state                                   │
│  ▢  Create Update function with message routing                             │
│  ▢  Build View function with layout composition                             │
│  ▢  Define keymap and help bindings                                         │
│  ▢  Implement styles.go with all Lipgloss styles                            │
│                                                                              │
│  PHASE 2: Components                                                         │
│  ───────────────────                                                        │
│  ▢  Epic progress bar component                                             │
│  ▢  Task list with custom delegate                                          │
│  ▢  Activity log viewport                                                   │
│  ▢  Status indicators (icons + colors)                                      │
│  ▢  Header with branch/commit info                                          │
│  ▢  Footer with keybindings                                                 │
│                                                                              │
│  PHASE 3: Views                                                              │
│  ─────────────                                                              │
│  ▢  Dashboard view (main layout)                                            │
│  ▢  Epic detail view                                                        │
│  ▢  Task detail view                                                        │
│  ▢  Multi-epic selector                                                     │
│  ▢  Help overlay                                                            │
│                                                                              │
│  PHASE 4: Wizard Mode                                                        │
│  ────────────────────                                                       │
│  ▢  Wizard progress stepper                                                 │
│  ▢  Document preview viewport                                               │
│  ▢  Approval confirmation (Huh)                                             │
│  ▢  Revision text input (Huh)                                               │
│  ▢  Phase transition animations                                             │
│                                                                              │
│  PHASE 5: File Integration                                                   │
│  ─────────────────────────                                                  │
│  ▢  Epic parser (.claude/epics/)                                            │
│  ▢  PRD parser (.claude/prds/)                                              │
│  ▢  Task parser (frontmatter + checkboxes)                                  │
│  ▢  File watcher (fsnotify)                                                 │
│  ▢  Git status integration                                                  │
│                                                                              │
│  PHASE 6: Polish                                                             │
│  ──────────────                                                             │
│  ▢  Loading spinners                                                        │
│  ▢  Error states                                                            │
│  ▢  Empty states                                                            │
│  ▢  Success animations                                                      │
│  ▢  Responsive layout handling                                              │
│  ▢  Terminal resize handling                                                │
│                                                                              │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

*Design complete. Ready for implementation.*

*Aesthetic: Neo-Brutalist Terminal*
*Framework: Charm (Bubbletea, Bubbles, Lipgloss, Huh)*
*Document Version: 1.0*
