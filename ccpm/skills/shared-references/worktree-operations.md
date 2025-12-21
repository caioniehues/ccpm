<overview>
Git worktrees enable parallel development by allowing multiple working directories for the same repository. Each worktree has its own working directory but shares the git object database.
</overview>

<creating_worktrees>
Always create worktrees from a clean main branch:

```bash
# Ensure main is up to date
git checkout main
git pull origin main

# Create worktree for epic
git worktree add ../epic-{name} -b epic/{name}
```

The worktree will be created as a sibling directory to maintain clean separation.
</creating_worktrees>

<working_in_worktrees>
**Agent commits:**
- Commit directly to the worktree
- Use small, focused commits
- Commit message format: `Issue #{number}: {description}`

**File operations:**
```bash
# Working directory is the worktree
cd ../epic-{name}

# Normal git operations work
git add {files}
git commit -m "Issue #{number}: {change}"

# View worktree status
git status
```
</working_in_worktrees>

<parallel_work>
Multiple agents can work in the same worktree if they touch different files:

```bash
# Agent A works on API
git add src/api/*
git commit -m "Issue #1234: Add user endpoints"

# Agent B works on UI (no conflict!)
git add src/ui/*
git commit -m "Issue #1235: Add dashboard component"
```
</parallel_work>

<merging_worktrees>
When epic is complete, merge back to main:

```bash
# From main repository (not worktree)
cd {main-repo}
git checkout main
git pull origin main

# Merge epic branch
git merge epic/{name}

# If successful, clean up
git worktree remove ../epic-{name}
git branch -d epic/{name}
```
</merging_worktrees>

<conflict_resolution>
If merge conflicts occur:

```bash
# Conflicts will be shown
git status

# Human resolves conflicts
# Then continue merge
git add {resolved-files}
git commit
```

Never force-resolve conflicts. Always involve humans.
</conflict_resolution>

<worktree_management>
**List active worktrees:**
```bash
git worktree list
```

**Remove stale worktree:**
```bash
# If worktree directory was deleted
git worktree prune

# Force remove worktree
git worktree remove --force ../epic-{name}
```

**Check worktree status:**
```bash
cd ../epic-{name} && git status && cd -
```
</worktree_management>

<best_practices>
1. **One worktree per epic** - Not per issue
2. **Clean before create** - Always start from updated main
3. **Commit frequently** - Small commits are easier to merge
4. **Delete after merge** - Don't leave stale worktrees
5. **Use descriptive branches** - `epic/feature-name` not `feature`
</best_practices>

<troubleshooting>
**Worktree already exists:**
```bash
git worktree remove ../epic-{name}
# Then create new one
```

**Branch already exists:**
```bash
git branch -D epic/{name}
# Or use existing branch
git worktree add ../epic-{name} epic/{name}
```

**Cannot remove worktree:**
```bash
git worktree remove --force ../epic-{name}
git worktree prune
```
</troubleshooting>
