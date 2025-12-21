---
description: Configure CCPM settings and preferences
allowed-tools: Bash(mkdir:*), Bash(touch:*), Bash(cat:*), Read, Write, Glob, AskUserQuestion
---

<objective>
Configure CCPM settings for the current project, including GitHub integration and workflow preferences.
</objective>

<process>
## Setup Workflow

### 1. Check Prerequisites
```
!`test -d .claude && echo "✅ CCPM initialized" || echo "❌ Run /pm:init first"`
```

### 2. Gather Configuration

Use AskUserQuestion to collect preferences:

1. **GitHub Integration**
   - Default labels to create
   - Issue template preferences
   - Auto-sync settings

2. **Workflow Settings**
   - Epic naming convention
   - Task numbering format
   - Context update frequency

3. **Team Settings** (if applicable)
   - Assignee defaults
   - Review requirements

### 3. Save Configuration

Write settings to `.claude/config.yaml`:

```yaml
github:
  auto_sync: true
  labels:
    - epic
    - task
    - in-progress
    - blocked
workflow:
  epic_prefix: ""
  task_format: "NNN"
  context_frequency: "session"
```

### 4. Validate Setup
Run `/ccpm:doctor` to verify configuration.
</process>

<success_criteria>
- Configuration file created at .claude/config.yaml
- Settings validated and applied
- Doctor check passes
</success_criteria>
