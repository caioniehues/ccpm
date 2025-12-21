---
name: ccpm-issue
description: Manage GitHub issues within CCPM workflow. Supports analyzing for parallel work, starting work, closing, editing, reopening, syncing progress, and checking status.
---

<objective>
Manage the full lifecycle of GitHub issues within the CCPM project management system. Issues are the atomic work units within epics, with support for parallel work streams and transparent progress tracking.

This skill handles seven core actions:
- **analyze**: Identify parallel work streams for maximum efficiency
- **start**: Begin work with parallel agents based on analysis
- **close**: Mark complete and close on GitHub
- **edit**: Modify issue details locally and on GitHub
- **reopen**: Reopen a closed issue
- **sync**: Push local updates as GitHub comments
- **status**: Check issue status and current state
</objective>

<shared_references>
Load before any operation:
- @ccpm/skills/shared-references/datetime.md
- @ccpm/skills/shared-references/frontmatter-operations.md
- @ccpm/skills/shared-references/github-operations.md
- @ccpm/skills/shared-references/agent-coordination.md
</shared_references>

<action name="analyze">
<description>
Analyze an issue to identify parallel work streams for maximum efficiency.
</description>

<preflight>
1. **Find local task file:**
   - Check if `.claude/epics/*/$ARGUMENTS.md` exists
   - If not, search for file with `github:.*issues/$ARGUMENTS` in frontmatter
   - If not found: "❌ No local task for issue #$ARGUMENTS. Run: /pm:import first"

2. **Check for existing analysis:**
   - If `.claude/epics/*/$ARGUMENTS-analysis.md` exists, ask: "⚠️ Analysis exists. Overwrite?"
</preflight>

<process>
1. **Read Issue Context**
   - Get from GitHub: `gh issue view $ARGUMENTS --json title,body,labels`
   - Read local task file for requirements, acceptance criteria, dependencies

2. **Identify Parallel Work Streams**
   Common patterns:
   - Database Layer: Schema, migrations, models
   - Service Layer: Business logic, data access
   - API Layer: Endpoints, validation, middleware
   - UI Layer: Components, pages, styles
   - Test Layer: Unit tests, integration tests
   - Documentation: API docs, README updates

   Key questions:
   - What files will be created/modified?
   - Which changes can happen independently?
   - What are the dependencies between changes?

3. **Create Analysis File**
   Save to `.claude/epics/{epic_name}/$ARGUMENTS-analysis.md`:
   ```markdown
   ---
   issue: $ARGUMENTS
   title: {issue_title}
   analyzed: {current_datetime}
   estimated_hours: {total_hours}
   parallelization_factor: {1.0-5.0}
   ---

   # Parallel Work Analysis: Issue #$ARGUMENTS

   ## Overview
   {description}

   ## Parallel Streams

   ### Stream A: {Name}
   **Scope**: {What this stream handles}
   **Files**: {file_patterns}
   **Agent Type**: {backend|frontend|fullstack|database}-specialist
   **Can Start**: immediately
   **Estimated Hours**: {hours}
   **Dependencies**: none

   [Additional streams...]

   ## Coordination Points
   - Shared files requiring coordination
   - Sequential requirements

   ## Conflict Risk Assessment
   - Low/Medium/High Risk assessment

   ## Parallelization Strategy
   Recommended approach: {sequential|parallel|hybrid}
   ```

4. **Output**
   ```
   ✅ Analysis complete for issue #$ARGUMENTS

   Identified {count} parallel work streams
   Parallelization potential: {factor}x speedup

   Next: Start work with /pm:issue-start $ARGUMENTS
   ```
</process>
</action>

<action name="start">
<description>
Begin work on a GitHub issue with parallel agents based on work stream analysis.
</description>

<preflight>
1. **Get issue details:**
   - Run: `gh issue view $ARGUMENTS --json state,title,labels,body`
   - If fails: "❌ Cannot access issue #$ARGUMENTS"

2. **Find local task file:**
   - Check `.claude/epics/*/$ARGUMENTS.md`
   - If not found: "❌ No local task for issue #$ARGUMENTS"

3. **Check for analysis:**
   - If no `.claude/epics/*/$ARGUMENTS-analysis.md` exists:
     "❌ No analysis found. Run: /pm:issue-analyze $ARGUMENTS first"
</preflight>

<process>
1. **Ensure Worktree Exists**
   - Extract epic name from task file path
   - Check if epic worktree exists
   - If not: "❌ No worktree for epic. Run: /pm:epic-start {epic_name}"

2. **Read Analysis**
   - Parse parallel streams
   - Identify which can start immediately
   - Note dependencies between streams

3. **Setup Progress Tracking**
   - Get current datetime
   - Create workspace: `.claude/epics/{epic_name}/updates/$ARGUMENTS/`
   - Update task file `updated` field

4. **Launch Parallel Agents**
   For each stream that can start immediately:
   - Create stream progress file
   - Launch agent with Task tool
   - Agent works in worktree, commits with "Issue #$ARGUMENTS: {change}"

5. **GitHub Assignment**
   - Run: `gh issue edit $ARGUMENTS --add-assignee @me --add-label "in-progress"`

6. **Output**
   ```
   ✅ Started parallel work on issue #$ARGUMENTS

   Epic: {epic_name}
   Worktree: ../epic-{epic_name}/

   Launching {count} parallel agents:
     Stream A: {name} ✓ Started
     Stream B: {name} ✓ Started

   Monitor with: /pm:epic-status {epic_name}
   ```
</process>
</action>

<action name="close">
<description>
Mark an issue as complete and close it on GitHub.
</description>

<process>
1. **Find Local Task File**
   - Check `.claude/epics/*/$ARGUMENTS.md`
   - If not found: "❌ No local task for issue #$ARGUMENTS"

2. **Update Local Status**
   - Get current datetime
   - Update task file: `status: closed`, `updated: {datetime}`

3. **Update Progress File**
   - Set completion: 100%
   - Add completion note with timestamp

4. **Close on GitHub**
   ```bash
   # Add completion comment
   gh issue comment $ARGUMENTS --body "✅ Task completed..."

   # Close the issue
   gh issue close $ARGUMENTS
   ```

5. **Update Epic Task List**
   - Check off task in epic issue body
   - Update epic progress percentage

6. **Output**
   ```
   ✅ Closed issue #$ARGUMENTS
     Local: Task marked complete
     GitHub: Issue closed & epic updated
     Epic progress: {new_progress}%

   Next: Run /pm:next for next priority task
   ```
</process>
</action>

<action name="edit">
<description>
Edit issue details locally and on GitHub.
</description>

<process>
1. **Get Current Issue State**
   - From GitHub: `gh issue view $ARGUMENTS --json title,body,labels`
   - Find local task file

2. **Interactive Edit**
   Ask user what to edit:
   - Title
   - Description/Body
   - Labels
   - Acceptance criteria (local only)
   - Priority/Size (local only)

3. **Update Local File**
   - Get current datetime
   - Update changed fields
   - Update `updated` field

4. **Update GitHub**
   - If title changed: `gh issue edit $ARGUMENTS --title "{new}"`
   - If body changed: `gh issue edit $ARGUMENTS --body-file {file}`
   - If labels changed: add/remove labels

5. **Output**
   ```
   ✅ Updated issue #$ARGUMENTS
     Changes: {list}
   Synced to GitHub: ✅
   ```
</process>
</action>

<action name="reopen">
<description>
Reopen a closed issue.
</description>

<process>
1. **Find Local Task File**
   - If not found: "❌ No local task for issue #$ARGUMENTS"

2. **Update Local Status**
   - Get current datetime
   - Update: `status: open`, `updated: {datetime}`

3. **Reset Progress**
   - Keep original started date
   - Add note about reopening with reason

4. **Reopen on GitHub**
   ```bash
   gh issue comment $ARGUMENTS --body "🔄 Reopening: {reason}"
   gh issue reopen $ARGUMENTS
   ```

5. **Update Epic Progress**
   - Recalculate with this task now open

6. **Output**
   ```
   🔄 Reopened issue #$ARGUMENTS
     Reason: {reason}
     Epic progress: {updated}%

   Start work with: /pm:issue-start $ARGUMENTS
   ```
</process>
</action>

<action name="sync">
<description>
Push local updates as GitHub issue comments for transparent audit trail.
</description>

<preflight>
1. **Repository Protection Check:**
   - Verify not syncing to template repository

2. **GitHub Authentication:**
   - Run: `gh auth status`
   - If not authenticated: "❌ Run: gh auth login"

3. **Issue Validation:**
   - Run: `gh issue view $ARGUMENTS --json state`
   - If not found: "❌ Issue #$ARGUMENTS not found"

4. **Local Updates Check:**
   - Check `.claude/epics/*/updates/$ARGUMENTS/` exists
   - If not: "❌ No local updates. Run: /pm:issue-start $ARGUMENTS"

5. **Check Last Sync:**
   - If synced < 5 minutes ago, ask: "⚠️ Recently synced. Force?"
   - If no changes: "ℹ️ No new updates to sync"
</preflight>

<process>
1. **Gather Local Updates**
   - Read from `.claude/epics/{epic}/updates/$ARGUMENTS/`
   - Check progress.md, notes.md, commits.md

2. **Update Progress Tracking**
   - Get current datetime
   - Update `last_sync` in frontmatter

3. **Format Update Comment**
   ```markdown
   ## 🔄 Progress Update - {date}

   ### ✅ Completed Work
   {items}

   ### 🔄 In Progress
   {items}

   ### 📊 Acceptance Criteria Status
   - ✅/🔄/⏸️/□ {criteria}

   ---
   *Progress: {completion}%*
   ```

4. **Post to GitHub**
   - Run: `gh issue comment $ARGUMENTS --body-file {file}`

5. **Update Local Files**
   - Update task file `updated` field
   - Update epic progress if task completed

6. **Output**
   ```
   ☁️ Synced updates to GitHub Issue #$ARGUMENTS

   Progress: {completion}%
   Epic progress: {epic_progress}%

   View: gh issue view #$ARGUMENTS --comments
   ```
</process>

<error_handling>
- Network error: Keep local updates for retry
- Rate limit: Save for later sync
- Permission denied: Check repository access
- Issue locked: Contact admin
</error_handling>
</action>

<action name="status">
<description>
Check issue status (open/closed) and current state.
</description>

<process>
1. **Fetch Issue Status**
   - Run: `gh issue view $ARGUMENTS --json state,title,labels,assignees,updatedAt`

2. **Status Display**
   ```
   🎫 Issue #$ARGUMENTS: {Title}

   📊 Status: {OPEN/CLOSED}
      Last update: {timestamp}
      Assignee: {assignee or "Unassigned"}

   🏷️ Labels: {labels}
   ```

3. **Epic Context** (if applicable)
   ```
   📚 Epic: {epic_name}
      Progress: {completed}/{total} tasks
   ```

4. **Local Sync Status**
   ```
   💾 Local: {exists/missing}
      Sync status: {in_sync/needs_sync}
   ```

5. **Status Indicators**
   - 🟢 Open and ready
   - 🟡 Open with blockers
   - 🔴 Open and overdue
   - ✅ Closed and complete
   - ❌ Closed without completion

6. **Suggested Actions**
   Based on status, suggest relevant commands
</process>
</action>

<success_criteria>
- **analyze**: Parallel work streams identified with clear separation
- **start**: Agents launched with proper coordination
- **close**: Issue closed locally and on GitHub, epic updated
- **edit**: Changes applied locally and synced to GitHub
- **reopen**: Issue reopened with history preserved
- **sync**: Updates posted as GitHub comments
- **status**: Clear status report with actionable suggestions
- All operations use real datetime
- Frontmatter properly maintained
</success_criteria>
