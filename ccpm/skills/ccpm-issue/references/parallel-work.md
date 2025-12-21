# Parallel Work Coordination

<overview>
CCPM enables multiple agents to work on the same issue in parallel by dividing work into independent streams with clear file ownership.
</overview>

<parallel_work_principles>
## Core Principles

1. **Clear scope**: Each stream owns specific files/directories
2. **No overlap**: Streams don't modify each other's files
3. **Explicit dependencies**: Document what depends on what
4. **Progress tracking**: Each stream tracks progress separately
5. **Coordination points**: Identify and manage shared resources
</parallel_work_principles>

<work_stream_structure>
## Work Stream Structure

**Stream definition:**
```markdown
### Stream A: Database Layer
**Scope**: Database schema and migrations
**Files**:
- `db/migrations/*.sql`
- `src/models/*.ts`
**Agent Type**: backend-specialist
**Can Start**: immediately
**Estimated Hours**: 4
**Dependencies**: none
```

**Key attributes:**
- **Scope**: What this stream handles (clear, bounded)
- **Files**: Specific files or patterns this stream owns
- **Agent Type**: Which type of agent should handle this
- **Can Start**: When this stream can begin (immediately vs waiting)
- **Dependencies**: Which other streams must complete first
</work_stream_structure>

<agent_types>
## Agent Types

**backend-specialist:**
- Database, API, services
- Files: `src/`, `db/`, `api/`, `services/`

**frontend-specialist:**
- UI components, pages, styles
- Files: `components/`, `pages/`, `styles/`, `public/`

**fullstack:**
- End-to-end features spanning backend and frontend
- Files: Any combination

**database-specialist:**
- Schema design, migrations, optimization
- Files: `db/`, `migrations/`, `schema/`

**test-specialist:**
- Testing infrastructure and test suites
- Files: `tests/`, `__tests__/`, `*.test.*`, `*.spec.*`

**devops-specialist:**
- CI/CD, deployment, infrastructure
- Files: `.github/`, `docker/`, `k8s/`, `.circleci/`
</agent_types>

<file_ownership>
## File Ownership Patterns

**Non-overlapping (ideal):**
```
Stream A: src/api/*
Stream B: src/ui/*
Stream C: tests/*
```
No conflicts - streams work independently.

**Shared files (requires coordination):**
```
Stream A: src/types/models.ts (add User model)
Stream B: src/types/models.ts (add Product model)
```
Both need same file - coordinate updates or make sequential.

**Configuration files (high conflict risk):**
```
Stream A: package.json (add express)
Stream B: package.json (add react)
Stream C: package.json (add jest)
```
Make one stream own config updates, others request changes.
</file_ownership>

<coordination_strategies>
## Coordination Strategies

**Strategy 1: Complete independence**
```
Stream A: backend/ (no shared files)
Stream B: frontend/ (no shared files)
Stream C: docs/ (no shared files)
```
Launch all in parallel, no coordination needed.

**Strategy 2: Sequential dependencies**
```
Stream A: Database schema (runs first)
  ↓
Stream B: API endpoints (waits for A)
  ↓
Stream C: UI components (waits for B)
```
Launch A → wait → launch B → wait → launch C.

**Strategy 3: Hybrid (parallel + sequential)**
```
Launch together:
- Stream A: Database layer
- Stream B: API layer

Wait for both to complete, then:
- Stream C: Integration tests (depends on A + B)
```

**Strategy 4: Coordinator stream**
```
Parallel streams:
- Stream A: Feature X backend
- Stream B: Feature Y backend
- Stream C: Feature Z backend

Coordinator stream:
- Stream D: Merge configs, run tests, deploy
```
Dedicated stream handles shared resources.
</coordination_strategies>

<stream_progress_tracking>
## Stream Progress Tracking

**Stream file structure:**
```
.claude/epics/{epic}/updates/{issue}/
├── stream-A.md
├── stream-B.md
└── stream-C.md
```

**Stream file format:**
```markdown
---
issue: 123
stream: Database Layer
agent: backend-specialist
started: 2024-01-15T10:30:00Z
status: in_progress|completed|blocked
---

# Stream A: Database Layer

## Scope
Database schema and migrations for user authentication

## Files
- `db/migrations/001_users.sql`
- `src/models/User.ts`

## Progress
- [x] Created migration file
- [x] Implemented User model
- [ ] Added indexes
- [ ] Wrote tests

## Coordination Notes
- Waiting for Stream B to define API types
- Shared config updated in coordination with Stream C

## Blockers
None currently
```
</stream_progress_tracking>

<conflict_detection>
## Conflict Detection

**Low risk (different directories):**
```
Stream A: src/backend/
Stream B: src/frontend/
Stream C: docs/
```
Unlikely conflicts - proceed with parallel execution.

**Medium risk (same directory, different files):**
```
Stream A: src/models/User.ts
Stream B: src/models/Product.ts
Stream C: src/models/Order.ts
```
Some coordination needed - communicate via progress files.

**High risk (same files):**
```
Stream A: src/types/index.ts (lines 1-50)
Stream B: src/types/index.ts (lines 51-100)
Stream C: src/types/index.ts (lines 101-150)
```
High conflict risk - consider making sequential or using coordinator.
</conflict_detection>

<coordination_rules>
## Coordination Rules

**For agents working in parallel:**

1. **Read your scope**: Know exactly which files you own
2. **Stay in scope**: Don't modify files outside your ownership
3. **Update progress**: Keep stream progress file current
4. **Coordinate conflicts**: If you need a file owned by another stream:
   - Check if stream is complete
   - Add coordination note to your progress file
   - Wait or request coordination via progress updates
5. **Commit frequently**: Small commits reduce merge conflicts
6. **Clear commit messages**: `Issue #123 Stream A: Add User model`

**For shared resources:**

1. **Document in analysis**: List all shared files upfront
2. **Assign ownership**: One stream owns shared files
3. **Request changes**: Other streams document needed changes
4. **Coordinate timing**: Owner makes changes when dependencies ready
</coordination_rules>

<launching_parallel_agents>
## Launching Parallel Agents

**Use Task tool for each stream:**
```yaml
Task:
  description: "Issue #123 Stream A: Database Layer"
  subagent_type: "backend-specialist"
  prompt: |
    You are working on Issue #123 in the epic worktree.

    Worktree location: ../epic-{epic-name}/
    Your stream: Database Layer

    Your scope:
    - Files to modify: db/migrations/*.sql, src/models/*.ts
    - Work to complete: Create user authentication schema

    Requirements:
    1. Read full task from: .claude/epics/{epic}/123.md
    2. Work ONLY in your assigned files
    3. Commit frequently with format: "Issue #123 Stream A: {change}"
    4. Update progress in: .claude/epics/{epic}/updates/123/stream-A.md
    5. Follow coordination rules

    If you need to modify files outside your scope:
    - Check if another stream owns them
    - Wait if necessary
    - Update your progress file with coordination notes

    Complete your stream's work and mark as completed when done.
```

**Launch immediately vs waiting:**
```bash
# Launch streams with no dependencies immediately
Task (Stream A) &
Task (Stream B) &

# Wait for dependencies
wait

# Launch dependent stream
Task (Stream C)
```
</launching_parallel_agents>

<monitoring_parallel_work>
## Monitoring Parallel Work

**Check all stream progress:**
```bash
issue=123
epic="epic-name"

for stream in .claude/epics/$epic/updates/$issue/stream-*.md; do
  stream_name=$(grep '^stream:' "$stream" | cut -d' ' -f2-)
  status=$(grep '^status:' "$stream" | cut -d' ' -f2)

  echo "Stream: $stream_name - Status: $status"
done
```

**Identify blockers:**
```bash
# Find blocked streams
grep -l "^status: blocked" .claude/epics/$epic/updates/$issue/stream-*.md
```

**Check completion:**
```bash
# Count completed streams
completed=$(grep -l "^status: completed" .claude/epics/$epic/updates/$issue/stream-*.md | wc -l)

# Count total streams
total=$(ls .claude/epics/$epic/updates/$issue/stream-*.md | wc -l)

echo "Progress: $completed/$total streams complete"
```
</monitoring_parallel_work>

<best_practices>
## Best Practices

1. **Clear boundaries**: Define file ownership explicitly
2. **Minimize overlap**: Design streams to be independent
3. **Document dependencies**: Make dependencies explicit in analysis
4. **Coordinate early**: Identify coordination points before starting
5. **Track progress**: Each stream maintains current status
6. **Frequent commits**: Small, focused commits reduce conflicts
7. **Clear naming**: Use consistent commit message format
8. **Handle blockers**: Document blockers in progress files
9. **Test integration**: Plan for integration testing after parallel work
10. **Learn from conflicts**: Update analysis patterns when conflicts occur
</best_practices>

<success_criteria>
Parallel work coordination is successful when:
- Each stream has clear, non-overlapping file ownership
- Dependencies documented and respected
- Stream progress tracked separately
- Coordination points identified upfront
- Blockers documented in progress files
- Agents stay within assigned scope
- Shared resources managed explicitly
- Work completes faster than sequential execution
- Integration succeeds without major conflicts
</success_criteria>
