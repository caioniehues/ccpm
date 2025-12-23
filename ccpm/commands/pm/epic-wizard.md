---
description: Guided epic creation with approval gates between phases
argument-hint: <epic-name>
allowed-tools: Bash, Read, Write, Edit, Glob, Task
---

<objective>
Guide user through complete epic creation workflow for "$ARGUMENTS" with explicit approval gates between each phase.
</objective>

<context>
This wizard orchestrates the full PRD → Epic → Tasks flow with approval pauses:

Phase 1: PRD Creation → Approval Gate
Phase 2: Epic Generation → Approval Gate  
Phase 3: Task Decomposition → Approval Gate
Phase 4: GitHub Sync (optional)

Approval phrases: "yes", "approved", "looks good", "proceed", "continue", "ok", "lgtm", "y"
Revision phrases: "no", "changes", "revise", "edit", "modify", "n"
</context>

<process>
## Phase Recovery (Run First)

Before starting any phase, check for existing progress:

```
Phase Detection Logic:
─────────────────────
1. Check .claude/prds/$ARGUMENTS.md
   ├── NOT EXISTS → Start Phase 1
   ├── EXISTS + status: pending → Start Phase 1 (review existing)
   └── EXISTS + status: approved → Skip to Phase 2

2. Check .claude/epics/$ARGUMENTS/epic.md  
   ├── NOT EXISTS → Start Phase 2
   ├── EXISTS + status: pending → Start Phase 2 (review existing)
   └── EXISTS + status: approved → Skip to Phase 3

3. Check .claude/epics/$ARGUMENTS/001.md (any task file)
   ├── NOT EXISTS → Start Phase 3
   └── EXISTS → Check epic status

4. Check epic status field
   ├── status: approved-for-work → Skip to Phase 4
   └── Otherwise → Start Phase 3 (review tasks)
```

When resuming from a later phase, show:
```
═══════════════════════════════════════════════════════════════
RESUMING EPIC WIZARD: $ARGUMENTS
═══════════════════════════════════════════════════════════════

Previously completed:
  ✓ Phase 1: PRD created and approved
  ✓ Phase 2: Epic created and approved
  
Resuming from Phase 3: Task Decomposition
═══════════════════════════════════════════════════════════════
```

## Phase 1: PRD Creation

1. Execute PRD creation:
   - Follow /pm:prd-new behavior for "$ARGUMENTS"
   - Create .claude/prds/$ARGUMENTS.md with comprehensive requirements
   
2. APPROVAL GATE 1:
   ```
   ═══════════════════════════════════════════════════════════════
   ✓ PRD CREATED: .claude/prds/$ARGUMENTS.md
   ═══════════════════════════════════════════════════════════════
   
   Please review the PRD above.
   
   → Approve: "yes", "approved", "looks good", "proceed", "lgtm"
   → Revise:  "no", "changes needed", "revise" (then describe changes)
   
   Do you approve this PRD?
   ```
   
3. If revision requested:
   - Ask: "What changes would you like?"
   - Apply changes to PRD
   - Show updated PRD
   - Return to approval gate

4. On approval:
   - Update PRD frontmatter: `approval_status: approved`
   - Proceed to Phase 2

## Phase 2: Epic Generation

1. Execute epic generation:
   - Follow /pm:prd-parse behavior for "$ARGUMENTS"
   - Create .claude/epics/$ARGUMENTS/epic.md
   
2. APPROVAL GATE 2:
   ```
   ═══════════════════════════════════════════════════════════════
   ✓ EPIC CREATED: .claude/epics/$ARGUMENTS/epic.md
   ═══════════════════════════════════════════════════════════════
   
   Please review the epic structure and technical approach.
   
   → Approve: "yes", "approved", "looks good", "proceed", "lgtm"
   → Revise:  "no", "changes needed", "revise" (then describe changes)
   
   Do you approve this epic?
   ```

3. If revision requested:
   - Ask: "What changes would you like?"
   - Apply changes to epic
   - Show updated epic
   - Return to approval gate

4. On approval:
   - Update epic frontmatter: `approval_status: approved`
   - Proceed to Phase 3

## Phase 3: Task Decomposition

1. Execute task decomposition:
   - Follow /pm:epic-decompose behavior for "$ARGUMENTS"
   - Create .claude/epics/$ARGUMENTS/001.md, 002.md, etc.
   
2. APPROVAL GATE 3:
   ```
   ═══════════════════════════════════════════════════════════════
   ✓ TASKS CREATED: {N} tasks in .claude/epics/$ARGUMENTS/
   ═══════════════════════════════════════════════════════════════
   
   Task Summary:
   - 001.md: [task name] (effort: S/M/L)
   - 002.md: [task name] (effort: S/M/L)
   - ...
   
   → Approve: "yes", "approved", "looks good", "proceed", "lgtm"
   → Revise:  "no", "changes needed", "revise" (then describe changes)
   
   Do you approve these tasks?
   ```

3. If revision requested:
   - Ask: "What changes would you like?"
   - Apply changes to tasks
   - Show updated task list
   - Return to approval gate

4. On approval:
   - Update epic frontmatter: `status: approved-for-work`
   - Proceed to Phase 4

## Phase 4: GitHub Sync (Optional)

1. Ask user:
   ```
   ═══════════════════════════════════════════════════════════════
   GITHUB SYNC OPTIONS
   ═══════════════════════════════════════════════════════════════
   
   Would you like to sync to GitHub?
   
   → "yes"  - Create GitHub issues now
   → "no"   - Keep local only (can sync later with /pm:epic-sync)
   → "skip" - Skip and proceed to work
   
   Your choice?
   ```

2. If "yes":
   - Execute /pm:epic-sync $ARGUMENTS
   - Report created issue numbers
   
3. If "no" or "skip":
   - Continue with local-only workflow

## Completion

Show final summary:
```
═══════════════════════════════════════════════════════════════
✓ EPIC WIZARD COMPLETE: $ARGUMENTS
═══════════════════════════════════════════════════════════════

Created:
  ✓ PRD: .claude/prds/$ARGUMENTS.md
  ✓ Epic: .claude/epics/$ARGUMENTS/epic.md
  ✓ Tasks: {N} task files
  {✓ GitHub: Issues #XXX-#YYY created | ○ GitHub: Local only}

Next Steps:
  → /pm:next                    - Start first priority task
  → /pm:epic-show $ARGUMENTS    - View epic details
  → /pm:epic-sync $ARGUMENTS    - Sync to GitHub (if skipped)

Ready to build!
```
</process>

<success_criteria>
- All phases completed with explicit user approval
- PRD created with approval_status marker
- Epic created with approval_status marker
- Tasks created and approved
- User informed of next steps
- Clean exit with actionable guidance
</success_criteria>
