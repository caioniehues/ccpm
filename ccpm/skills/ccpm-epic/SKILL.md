---
name: ccpm-epic
description: Manage epics in CCPM workflow. Supports decomposing to tasks, syncing to GitHub, starting parallel work, merging, editing, closing, oneshot operations, and refreshing progress.
---

<objective>
Manage the full lifecycle of epics in the CCPM project management system. Epics are the primary work containers that hold related tasks and provide structure for parallel development workflows.

This skill handles eight core actions:
- **decompose**: Break epic into concrete, actionable tasks
- **sync**: Push epic and tasks to GitHub as issues
- **start**: Launch parallel agents to work on epic tasks
- **merge**: Merge completed epic back to main branch
- **edit**: Modify epic details after creation
- **close**: Mark epic as complete when all tasks done
- **oneshot**: Decompose and sync in one operation
- **refresh**: Update epic progress based on task states
</objective>

<shared_references>
Load before any operation:
- @ccpm/skills/shared-references/datetime.md
- @ccpm/skills/shared-references/frontmatter-operations.md
- @ccpm/skills/shared-references/github-operations.md
- @ccpm/skills/shared-references/worktree-operations.md
- @ccpm/skills/shared-references/agent-coordination.md
</shared_references>

<action name="decompose">
<description>
Break an epic into concrete, actionable tasks with proper dependency tracking.
</description>

<preflight>
1. Verify epic exists: `.claude/epics/$ARGUMENTS/epic.md`
2. Check for existing tasks - ask before overwriting
3. Validate epic frontmatter (name, status, created, prd)
4. Warn if epic status is "completed"
</preflight>

<process>
1. **Read Epic**
   - Load epic from `.claude/epics/$ARGUMENTS/epic.md`
   - Understand technical approach and requirements

2. **Determine Parallel Strategy**
   - Small epic (<5 tasks): Create sequentially
   - Medium epic (5-10): Batch into 2-3 groups
   - Large epic (>10): Launch parallel agents (max 5 concurrent)

3. **Create Task Files**
   Each task file (001.md, 002.md, etc.):
   ```markdown
   ---
   name: [Task Title]
   status: open
   created: {datetime}
   updated: {datetime}
   github: [placeholder]
   depends_on: []
   parallel: true
   conflicts_with: []
   ---

   # Task: [Title]

   ## Description
   [Clear description]

   ## Acceptance Criteria
   - [ ] Criterion 1
   - [ ] Criterion 2

   ## Technical Details
   [Implementation approach, files affected]

   ## Effort Estimate
   - Size: XS/S/M/L/XL
   - Hours: {estimate}
   ```

4. **Validate Dependencies**
   - Ensure referenced dependencies exist
   - Check for circular dependencies
   - Warn if issues found

5. **Update Epic**
   Add "Tasks Created" section with summary
</process>

<output>
```
✅ Created {count} tasks for epic: $ARGUMENTS
  Parallel: {count}, Sequential: {count}
  Total effort: {hours}h

Next: /pm:epic-sync $ARGUMENTS
```
</output>
</action>

<action name="sync">
<description>
Push epic and tasks to GitHub as issues, rename files to issue numbers.
</description>

<preflight>
1. Verify epic exists
2. Check remote isn't CCPM template repository
3. Count task files - warn if none found
</preflight>

<process>
1. **Create Epic Issue**
   - Strip frontmatter from epic.md
   - Create GitHub issue with labels: epic, epic:{name}, feature/bug
   - Store issue number

2. **Create Task Sub-Issues**
   - Small batch (<5): Sequential creation
   - Large batch: Parallel agents
   - Use gh-sub-issue if available, fallback to gh issue
   - Apply labels: task, epic:{name}

3. **Rename Task Files**
   - Build old→new ID mapping
   - Rename 001.md → {issue_id}.md
   - Update depends_on/conflicts_with references
   - Update github field in frontmatter

4. **Update Epic with Task List** (fallback only)
   If not using gh-sub-issue, add task list to epic body

5. **Update Epic Frontmatter**
   - Add github URL
   - Update timestamp
   - Update Tasks Created section with real issue IDs

6. **Create Mapping File**
   Save to `.claude/epics/$ARGUMENTS/github-mapping.md`

7. **Create Worktree**
   `git worktree add ../epic-$ARGUMENTS -b epic/$ARGUMENTS`
</process>

<output>
```
✅ Synced to GitHub
  - Epic: #{number}
  - Tasks: {count} sub-issues
  - Files renamed to issue IDs
  - Worktree: ../epic-$ARGUMENTS

Next: /pm:epic-start $ARGUMENTS
```
</output>
</action>

<action name="start">
<description>
Launch parallel agents to work on epic tasks in shared branch.
</description>

<preflight>
1. Verify epic exists
2. Check GitHub sync (github field in frontmatter)
3. Check for uncommitted changes
</preflight>

<process>
1. **Create/Enter Branch**
   - Check uncommitted changes first
   - Create or checkout `epic/$ARGUMENTS`
   - Push with tracking

2. **Identify Ready Issues**
   - Parse task frontmatter for status, depends_on, parallel
   - Categorize: Ready, Blocked, In Progress, Complete

3. **Analyze Ready Issues**
   - Check for analysis files
   - Run analysis if missing

4. **Launch Parallel Agents**
   For each ready issue with analysis:
   - Create stream progress files
   - Launch agents with Task tool
   - Agents work in same branch
   - Commit format: "Issue #{id}: {change}"

5. **Track Active Agents**
   Create/update `execution-status.md`

6. **Handle Dependencies**
   As agents complete, launch newly-ready work
</process>

<output>
```
🚀 Epic Execution Started: $ARGUMENTS

Branch: epic/$ARGUMENTS

Launching {count} agents:
  Issue #1234: Schema
    ├─ Stream A: ✓ Started
    └─ Stream B: ✓ Started

Monitor: /pm:epic-status $ARGUMENTS
```
</output>
</action>

<action name="merge">
<description>
Merge completed epic from worktree back to main branch.
</description>

<preflight>
1. Verify worktree exists
2. Check for active agents
</preflight>

<process>
1. **Pre-Merge Validation**
   - Check uncommitted changes in worktree
   - Fetch and check branch status

2. **Run Tests** (optional)
   - Detect project type
   - Run appropriate test command
   - Warn if tests fail

3. **Update Epic Documentation**
   - Set status to "completed"
   - Update completion date

4. **Attempt Merge**
   - Checkout main, pull
   - Merge with --no-ff
   - Include feature list in commit

5. **Handle Merge Conflicts**
   If conflicts, provide resolution guidance

6. **Post-Merge Cleanup**
   - Push to remote
   - Remove worktree
   - Delete branch (local and remote)
   - Archive epic to `.claude/epics/archived/`

7. **Close GitHub Issues**
   - Close epic issue
   - Close all task issues
</process>

<output>
```
✅ Epic Merged: $ARGUMENTS

  Commits: {count}
  Files: {count}
  Issues closed: {count}

Cleanup:
  ✓ Worktree removed
  ✓ Branch deleted
  ✓ Epic archived
```
</output>
</action>

<action name="edit">
<description>
Edit epic details after creation.
</description>

<process>
1. **Read Current Epic**
   - Parse frontmatter
   - Read content sections

2. **Interactive Edit**
   Ask user what to edit:
   - Name/Title
   - Description/Overview
   - Architecture decisions
   - Technical approach
   - Dependencies
   - Success criteria

3. **Update Epic File**
   - Preserve frontmatter except `updated`
   - Apply user's edits
   - Update `updated` field

4. **Option to Update GitHub**
   If has GitHub URL, ask to sync:
   `gh issue edit {number} --body-file epic.md`
</process>

<output>
```
✅ Updated epic: $ARGUMENTS
  Changes: {sections_edited}
  GitHub: {updated/skipped}
```
</output>
</action>

<action name="close">
<description>
Mark epic as complete when all tasks are done.
</description>

<process>
1. **Verify All Tasks Complete**
   - Check all task files have `status: closed`
   - Block if any open tasks remain

2. **Update Epic Status**
   ```yaml
   status: completed
   progress: 100%
   updated: {datetime}
   completed: {datetime}
   ```

3. **Update PRD Status**
   If epic references PRD, set to "complete"

4. **Close on GitHub**
   `gh issue close {number} --comment "Epic completed"`

5. **Archive Option**
   Ask to archive to `.claude/epics/.archived/`
</process>

<output>
```
✅ Epic closed: $ARGUMENTS
  Tasks: {count}
  Duration: {days}

Next: /pm:next
```
</output>
</action>

<action name="oneshot">
<description>
Decompose epic into tasks and sync to GitHub in one operation.
</description>

<preflight>
1. Verify epic exists
2. No existing tasks (would create duplicates)
3. Not already synced to GitHub
</preflight>

<process>
1. **Execute Decompose**
   Run decompose action with all validation

2. **Execute Sync**
   Immediately run sync action

This is a convenience wrapper that orchestrates:
- `/pm:epic-decompose $ARGUMENTS`
- `/pm:epic-sync $ARGUMENTS`
</process>

<output>
```
🚀 Epic Oneshot Complete: $ARGUMENTS

Step 1: Decomposition ✓
  Tasks: {count}

Step 2: GitHub Sync ✓
  Epic: #{number}
  Worktree: ../epic-$ARGUMENTS

Ready: /pm:epic-start $ARGUMENTS
```
</output>
</action>

<action name="refresh">
<description>
Update epic progress based on task states.
</description>

<process>
1. **Count Task Status**
   - Total tasks
   - Closed tasks
   - Open tasks

2. **Calculate Progress**
   `progress = (closed / total) * 100`

3. **Update GitHub Task List**
   Sync checkbox states to epic issue body

4. **Determine Epic Status**
   - 0%: backlog
   - 1-99%: in-progress
   - 100%: completed

5. **Update Epic Frontmatter**
   ```yaml
   status: {calculated}
   progress: {calculated}%
   updated: {datetime}
   ```
</process>

<output>
```
🔄 Epic refreshed: $ARGUMENTS

Tasks: {closed}/{total}
Progress: {old}% → {new}%
Status: {old} → {new}
GitHub: ✓ Updated

{If 100%}: Run /pm:epic-close $ARGUMENTS
```
</output>
</action>

<success_criteria>
- **decompose**: Tasks created with proper dependencies and estimates
- **sync**: Issues created on GitHub, files renamed, worktree created
- **start**: Parallel agents launched with proper coordination
- **merge**: Clean merge with cleanup and archiving
- **edit**: Changes applied locally and optionally to GitHub
- **close**: Epic marked complete with all tasks verified
- **oneshot**: Decompose + sync completed successfully
- **refresh**: Progress accurately calculated and synced
- All operations use real datetime
- Frontmatter properly maintained
- GitHub state in sync with local state
</success_criteria>
