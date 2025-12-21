---
description: Configure CCPM settings and preferences
allowed-tools: Bash(mkdir:*), Bash(touch:*), Bash(cat:*), Read, Write, Glob, AskUserQuestion
---

<objective>
Configure CCPM settings for the current project, including GitHub integration and workflow preferences.
</objective>

<process>
**1. Check Prerequisites**

```
!`test -d .claude && echo "✅ CCPM initialized" || echo "❌ Run /pm:init first"`
```

**2. Gather Configuration**

Use AskUserQuestion to collect preferences:

- **GitHub Integration**: Default labels to create, issue template preferences, auto-sync settings
- **Workflow Settings**: Epic naming convention, task numbering format, context update frequency
- **Team Settings** (if applicable): Assignee defaults, review requirements

**3. Save Configuration**

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

**4. Validate Setup**

Run `/ccpm:doctor` to verify configuration.
</process>

<success_criteria>
- Configuration file created at .claude/config.yaml
- Settings validated and applied
- Doctor check passes
</success_criteria>
