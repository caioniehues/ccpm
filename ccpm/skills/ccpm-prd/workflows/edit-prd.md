# Workflow: Edit PRD

Edit an existing Product Requirements Document with user-specified changes.

## Input
- `$ARGUMENTS`: Feature name (kebab-case)

## Preflight Checks

1. **Verify PRD Exists:**
   ```bash
   test -f .claude/prds/$ARGUMENTS.md
   ```
   If not found: "PRD not found: .claude/prds/$ARGUMENTS.md. Use /pm:prd-new to create."

2. **Read Current PRD:**
   ```bash
   cat .claude/prds/$ARGUMENTS.md
   ```
   Parse frontmatter and content sections.

3. **Check for Associated Epic:**
   ```bash
   test -d .claude/epics/$ARGUMENTS
   ```
   If exists: Note that changes may need to be reflected in epic.

4. **Get Current DateTime:**
   ```bash
   date -u +"%Y-%m-%dT%H:%M:%SZ"
   ```
   Store for updated timestamp.

## Edit Options

Present options to user:

```
📝 Editing PRD: $ARGUMENTS

Which section would you like to edit?
1. Executive Summary (problem, solution, success criteria)
2. User Stories (add, modify, remove stories)
3. Requirements (functional, non-functional, constraints)
4. Out of Scope (add or remove exclusions)
5. Dependencies (internal, external, prerequisites)
6. Success Metrics (acceptance criteria, KPIs)
7. Full PRD (review and edit any section)
8. Status only (change status: backlog/in-progress/completed/on-hold)

Enter your choice:
```

**Wait for user response before proceeding.**

## Edit Process

### For Section-Specific Edits (Options 1-6)

1. **Display Current Section:**
   Show the current content of the selected section.

2. **Collect Changes:**
   Ask: "What changes would you like to make to this section?"
   
3. **Apply Changes:**
   - Update the section content
   - Preserve other sections unchanged
   - Update `updated` field in frontmatter

4. **Confirm Changes:**
   Show diff of changes, ask for confirmation.

### For Full PRD Edit (Option 7)

1. **Display Full PRD:**
   Show complete PRD content.

2. **Identify Changes:**
   Ask: "What would you like to change? You can specify multiple sections."

3. **Apply Changes:**
   Make all requested changes.

4. **Validate:**
   Ensure no sections were accidentally deleted.

### For Status Change (Option 8)

1. **Show Current Status:**
   ```
   Current status: {status}
   
   Available statuses:
   - backlog: Not started
   - in-progress: Active development
   - completed: Feature shipped
   - on-hold: Paused
   
   Enter new status:
   ```

2. **Update Status:**
   Change only the `status` field in frontmatter.

## Update Frontmatter

After any edit, update the frontmatter:

```yaml
---
name: $ARGUMENTS
description: {unchanged or updated}
status: {unchanged or updated}
created: {PRESERVE ORIGINAL - never change}
updated: {new datetime from preflight}
---
```

**CRITICAL**: Never modify the `created` field.

## Epic Impact Check

If epic exists (`.claude/epics/$ARGUMENTS/`):

1. **Assess Impact:**
   - Requirements changes → May need task updates
   - Scope changes → May need epic re-decomposition
   - Status changes → Update epic status

2. **Notify User:**
   ```
   ⚠️ This PRD has an associated epic.
   
   Changes to requirements or scope may require:
   - Updating epic.md
   - Modifying task files
   - Re-syncing with GitHub
   
   Consider running: /pm:epic-refresh $ARGUMENTS
   ```

## Save Changes

1. **Write Updated PRD:**
   Write the complete PRD back to `.claude/prds/$ARGUMENTS.md`

2. **Validate:**
   ```bash
   test -s .claude/prds/$ARGUMENTS.md
   head -5 .claude/prds/$ARGUMENTS.md | grep -q "^---"  # Check frontmatter
   ```

3. **Display Summary:**
   ```
   ✅ PRD Updated: .claude/prds/$ARGUMENTS.md
   
   📋 Changes Made:
      - {Section 1}: {brief description}
      - {Section 2}: {brief description}
   
   ⏰ Updated: {datetime}
   📝 Original Created: {preserved created datetime}
   
   {Epic impact warning if applicable}
   ```

## Error Handling

- **PRD not found**: Suggest `/pm:prd-new` or `/pm:prd-list`
- **Invalid section choice**: Re-display options
- **File write fails**: Report error, preserve original
- **Frontmatter corruption**: Reconstruct from known fields

## Success Criteria

PRD edit is complete when:
- [ ] PRD file exists and was readable
- [ ] User specified which section(s) to edit
- [ ] Changes applied correctly
- [ ] `updated` field set to current datetime
- [ ] `created` field preserved unchanged
- [ ] File validated after write
- [ ] Epic impact assessed (if applicable)
- [ ] Summary provided to user
