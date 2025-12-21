---
name: ccpm-worktree
description: Manages git worktrees for parallel epic development. Use when creating, monitoring, merging, or cleaning up worktrees for CCPM epics. Enables parallel agent work by providing isolated working directories that share the git object database.
---

<objective>
Manage git worktrees for parallel epic development, enabling isolated working directories that share the git object database for concurrent agent work.
</objective>

<essential_principles>
<overview>
Git worktrees enable parallel development by allowing multiple working directories for the same repository. Each worktree has its own working directory but shares the git object database.
</overview>

<key_concepts>
**Worktree Structure**:
```
project/
├── main-repo/          (primary working directory)
├── epic-feature-a/     (worktree for epic/feature-a)
├── epic-feature-b/     (worktree for epic/feature-b)
└── epic-refactor/      (worktree for epic/refactor)
```

**Naming Conventions**:
- Worktree directories: `../epic-{name}` (sibling to main repo)
- Branch names: `epic/{name}`
- One worktree per epic (not per task)

**Parallel Agent Work**:
- Multiple agents work in the SAME worktree
- Agents touch different files to avoid conflicts
- Commits are made directly in the worktree
- Commit message format: `Issue #{number}: {description}`
</key_concepts>

<best_practices>
1. **Clean before create** - Always start from updated main
2. **Commit frequently** - Small commits are easier to merge
3. **Delete after merge** - Don't leave stale worktrees
4. **Never force-resolve** - Conflicts require human intervention
</best_practices>
</essential_principles>

<intake>
What would you like to do with worktrees?

1. **Create** - Create a new worktree for an epic
2. **List** - Show all active worktrees and their status
3. **Status** - Check status of a specific worktree
4. **Merge** - Merge worktree branch back to main
5. **Remove** - Clean up a worktree after merge
6. **Prune** - Remove stale worktree references

Provide the operation name, optionally with epic name.

**Wait for response before proceeding.**
</intake>

<routing>
| Response | Operation | Workflow |
|----------|-----------|----------|
| 1, "create", "new", "start" | Create worktree | Use `worktree-manager` subagent or delegate to `/pm:epic-start-worktree` |
| 2, "list", "ls", "show all" | List worktrees | Execute `git worktree list` with status check |
| 3, "status", "check" | Check worktree | Check specific worktree with git status |
| 4, "merge" | Merge worktree | Delegate to `/pm:epic-merge` command |
| 5, "remove", "delete", "cleanup" | Remove worktree | Execute worktree removal workflow |
| 6, "prune" | Prune stale refs | Execute `git worktree prune` |

**After determining the operation, execute the appropriate workflow or delegate to the command/subagent.**
</routing>

<workflow name="create">
**Create Worktree Workflow**

1. **Verify Prerequisites**
   ```bash
   # Ensure we're in main repo
   git rev-parse --git-dir

   # Check for existing worktree
   git worktree list | grep "epic-{name}"
   ```

2. **Update Main Branch**
   ```bash
   git checkout main
   git pull origin main
   ```

3. **Create Worktree**
   ```bash
   git worktree add ../epic-{name} -b epic/{name}
   ```

4. **Verify Creation**
   ```bash
   git worktree list
   cd ../epic-{name} && git status && cd -
   ```

5. **Report**
   - Worktree path
   - Branch name
   - Ready for agent work
</workflow>

<workflow name="list">
**List Worktrees Workflow**

```bash
# List all worktrees
git worktree list

# For each worktree, check status
for wt in $(git worktree list --porcelain | grep worktree | cut -d' ' -f2); do
  echo "=== $wt ==="
  cd "$wt" && git status --short && git log --oneline -3 && cd -
done
```

Report summary:
- Total active worktrees
- Status of each (clean, dirty, commits ahead)
- Recommendations
</workflow>

<workflow name="merge">
**Merge Worktree Workflow**

1. **Pre-Merge Validation**
   ```bash
   cd ../epic-{name} && git status
   git log main..HEAD --oneline
   ```

2. **Return to Main**
   ```bash
   cd {main-repo}
   git checkout main
   git pull origin main
   ```

3. **Merge**
   ```bash
   git merge epic/{name}
   ```

4. **Handle Results**
   - If success: proceed to cleanup
   - If conflicts: STOP, report to user, do NOT force-resolve

5. **Cleanup** (if merge succeeded)
   ```bash
   git worktree remove ../epic-{name}
   git branch -d epic/{name}
   ```
</workflow>

<workflow name="remove">
**Remove Worktree Workflow**

1. **Verify Worktree Exists**
   ```bash
   git worktree list | grep "epic-{name}"
   ```

2. **Check for Uncommitted Changes**
   ```bash
   cd ../epic-{name} && git status
   ```

   If dirty: warn user and confirm before proceeding

3. **Remove Worktree**
   ```bash
   git worktree remove ../epic-{name}
   ```

   If fails, try force:
   ```bash
   git worktree remove --force ../epic-{name}
   ```

4. **Delete Branch** (optional, only if merged)
   ```bash
   git branch -d epic/{name}
   ```

5. **Prune Stale References**
   ```bash
   git worktree prune
   ```
</workflow>

<error_handling>
**Worktree already exists:**
```bash
git worktree remove ../epic-{name}
# Then create new one
```

**Branch already exists:**
```bash
# Use existing branch
git worktree add ../epic-{name} epic/{name}

# Or delete and start fresh (only if safe)
git branch -D epic/{name}
git worktree add ../epic-{name} -b epic/{name}
```

**Cannot remove worktree:**
```bash
git worktree remove --force ../epic-{name}
git worktree prune
```

**Merge conflicts:**
- STOP immediately
- Run `git status` to identify conflicts
- Report conflict files to user
- Provide manual resolution instructions
- DO NOT proceed with merge
</error_handling>

<subagent_reference>
For complex worktree operations, use the `worktree-manager` subagent:

```
Task tool with subagent_type: worktree-manager
```

The subagent handles:
- Full lifecycle management
- Error recovery
- Conflict detection
- Status monitoring
</subagent_reference>

<command_reference>
Related commands in `ccpm/commands/pm/`:

- `/pm:epic-start-worktree {epic_name}` - Create worktree for epic
- `/pm:epic-start {epic_name}` - Start agents (includes worktree creation)
- `/pm:epic-merge {epic_name}` - Merge epic branch to main
- `/pm:epic-status {epic_name}` - Check epic and worktree status
</command_reference>

<quick_start>
**Quick start:**

```bash
# Create worktree for an epic
/pm:epic-start-worktree my-feature

# List all worktrees
git worktree list

# Remove worktree after merge
git worktree remove ../epic-my-feature
```
</quick_start>

<success_criteria>
This skill successfully handles worktree operations when:
- User intent is correctly identified
- Appropriate workflow or command is executed
- Worktree operations complete without errors
- Conflicts are properly escalated (never auto-resolved)
- Clean state is maintained (no orphaned worktrees)
</success_criteria>
