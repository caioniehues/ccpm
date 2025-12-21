---
description: Remove CCPM from the current project
allowed-tools: Bash(rm:*), Bash(ls:*), Read, Glob, AskUserQuestion
---

<objective>
Safely remove CCPM installation while preserving project data if requested.
</objective>

<process>
## Uninstall Workflow

### 1. Survey Installation
```
!`ls -la .claude/ 2>/dev/null | head -10 || echo "No .claude directory found"`
!`ls -la ccpm/ 2>/dev/null | head -5 || echo "No ccpm directory found"`
```

### 2. Confirm with User
Use AskUserQuestion to confirm:

1. **Data Preservation**
   - Keep .claude/epics/ (your work)?
   - Keep .claude/prds/ (your documents)?
   - Keep .claude/context/ (your context files)?

2. **Confirmation**
   - "This will remove CCPM. Proceed?"

### 3. Remove CCPM Files (if confirmed)

**Remove CCPM system files:**
- `ccpm/` directory (commands, scripts, agents, skills)
- CCPM-specific configuration

**Optionally preserve user data:**
- `.claude/epics/` - User's epic work
- `.claude/prds/` - User's PRD documents
- `.claude/context/` - User's context files

### 4. Clean Up
- Remove GitHub labels (optional)
- Remove any CCPM-specific git hooks

### 5. Report
```
Uninstall complete.

Removed:
  - ccpm/ directory
  - CCPM configuration

Preserved (per your request):
  - .claude/epics/
  - .claude/prds/
```
</process>

<success_criteria>
- CCPM system files removed
- User data preserved as requested
- No orphaned configuration
- Clean project state
</success_criteria>
