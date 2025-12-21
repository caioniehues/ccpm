<overview>
Rules for multiple agents working in parallel within the same epic worktree. Ensures agents can work on different files simultaneously without conflicts.
</overview>

<core_principles>
1. **File-level parallelism** - Agents working on different files never conflict
2. **Explicit coordination** - When same file needed, coordinate explicitly
3. **Fail fast** - Surface conflicts immediately, don't try to be clever
4. **Human resolution** - Conflicts are resolved by humans, not agents
</core_principles>

<work_stream_assignment>
Each agent is assigned a work stream with specific file patterns:

```yaml
# From {issue}-analysis.md
Stream A: Database Layer
  Files: src/db/*, migrations/*
  Agent: backend-specialist

Stream B: API Layer
  Files: src/api/*
  Agent: api-specialist
```

Agents should **only modify files in their assigned patterns**.
</work_stream_assignment>

<file_access_coordination>
**Check before modifying shared files:**
```bash
git status {file}

# If modified by another agent, wait
if [[ $(git status --porcelain {file}) ]]; then
  echo "Waiting for {file} to be available..."
  sleep 30
  # Retry
fi
```

**Make commits atomic and focused:**
```bash
# Good - Single purpose
git add src/api/users.ts src/api/users.test.ts
git commit -m "Issue #1234: Add user CRUD endpoints"

# Bad - Mixed concerns
git add src/api/* src/db/* src/ui/*
git commit -m "Issue #1234: Multiple changes"
```
</file_access_coordination>

<agent_communication>
**Through commits:**
```bash
# Check what others have done
git log --oneline -10

# Pull latest changes
git pull origin epic/{name}
```

**Through progress files:**
```markdown
# .claude/epics/{epic}/updates/{issue}/stream-A.md
---
stream: Database Layer
agent: backend-specialist
status: in_progress
---

## Completed
- Created user table schema

## Working On
- Adding indexes

## Blocked
- None
```
</agent_communication>

<conflict_handling>
**Detection:**
```bash
# If commit fails due to conflict
git commit -m "Issue #1234: Update"
# Error: conflicts exist

echo "Conflict detected in {files}"
echo "Human intervention needed"
```

**Resolution rules:**
1. Agent detects conflict
2. Agent reports issue
3. Agent pauses work
4. Human resolves
5. Agent continues

**Never attempt automatic merge resolution.**
</conflict_handling>

<synchronization>
**Natural sync points:**
- After each commit
- Before starting new file
- When switching work streams
- Every 30 minutes of work

**Explicit sync:**
```bash
git pull --rebase origin epic/{name}

if [[ $? -ne 0 ]]; then
  echo "Sync failed - human help needed"
  exit 1
fi
```
</synchronization>

<best_practices>
1. **Commit early and often** - Smaller commits = fewer conflicts
2. **Stay in your lane** - Only modify assigned files
3. **Communicate changes** - Update progress files
4. **Pull frequently** - Stay synchronized with other agents
5. **Fail loudly** - Report issues immediately
6. **Never force** - No `--force` flags ever
</best_practices>
