# Decompose Epic Workflow

Break epic into concrete, actionable tasks.

## Input
- `$ARGUMENTS`: Epic name

## Preflight Checks

1. **Verify epic exists:**
   ```bash
   test -f .claude/epics/$ARGUMENTS/epic.md
   ```
   If not: "Epic not found. Run /pm:prd-parse first"

2. **Check existing tasks:**
   ```bash
   ls .claude/epics/$ARGUMENTS/[0-9]*.md 2>/dev/null
   ```
   If found: Ask "Delete and recreate? (yes/no)"

3. **Validate frontmatter:**
   - Required: name, status, created, prd
   - Status should not be "completed" (warn if so)

## Execution Steps

### 1. Read Epic

Load `.claude/epics/$ARGUMENTS/epic.md`:
- Understand technical approach
- Review any task breakdown preview
- Identify scope and requirements

### 2. Analyze for Parallel Creation

Determine strategy:
- **< 5 tasks**: Create sequentially
- **5-10 tasks**: Batch into 2-3 groups, spawn agents
- **> 10 tasks**: Use epic-planner agent for decomposition

### 3. Create Task Files

For each task, create `.claude/epics/$ARGUMENTS/{number}.md`:

```markdown
---
name: {Task Title}
status: open
created: {datetime}
updated: {datetime}
github: {placeholder - updated on sync}
depends_on: []
parallel: true
conflicts_with: []
effort: M
---

# Task: {Title}

## Description
{What needs to be done}

## Acceptance Criteria
- [ ] Criterion 1
- [ ] Criterion 2
- [ ] Criterion 3

## Technical Details
- Implementation approach
- Files affected
- Key considerations

## Dependencies
- Task dependencies
- External dependencies

## Definition of Done
- [ ] Code implemented
- [ ] Tests passing
- [ ] Documentation updated
```

### 4. Task File Guidelines

**Naming:**
- Sequential: 001.md, 002.md, etc.
- Will be renamed to issue IDs on sync

**Frontmatter:**
- `name`: Descriptive title
- `status`: Always "open" for new tasks
- `created/updated`: Use `date -u +"%Y-%m-%dT%H:%M:%SZ"`
- `depends_on`: List of task numbers [001, 002]
- `parallel`: true if can run alongside others
- `conflicts_with`: Tasks modifying same files

**Effort Sizes:**
- XS: < 1 hour
- S: 1-2 hours
- M: 2-4 hours
- L: 4-8 hours
- XL: > 8 hours (consider splitting)

### 5. Parallel Task Creation

For larger epics, use epic-planner agent:

```yaml
Task:
  description: "Create task files batch {X}"
  subagent_type: "epic-planner"
  prompt: |
    Create task files for epic: $ARGUMENTS
    Tasks: {list of tasks for this batch}

    For each task:
    1. Create .claude/epics/$ARGUMENTS/{number}.md
    2. Use proper frontmatter
    3. Set dependencies correctly
```

### 6. Update Epic

Add task summary to epic.md:

```markdown
## Tasks Created
- [ ] 001.md - {Title} (parallel: true)
- [ ] 002.md - {Title} (parallel: false)

Total tasks: {count}
Parallel tasks: {parallel_count}
Sequential tasks: {sequential_count}
```

### 7. Validate

Before finalizing:
- All tasks have clear acceptance criteria
- Task sizes reasonable (1-3 days each)
- Dependencies logical
- Parallel tasks don't conflict
- Combined tasks cover all requirements

## Output

```
Created {count} tasks for epic: $ARGUMENTS

Summary:
  Parallel: {parallel_count}
  Sequential: {sequential_count}

Phase 1 (Start immediately):
  001 - {title} [M]
  002 - {title} [S]

Phase 2 (After Phase 1):
  003 - {title} [L] depends on 001
  004 - {title} [M] depends on 002

Ready to sync: /pm:epic-sync $ARGUMENTS
```

## Error Recovery

- If partial completion: list created tasks
- Provide cleanup option
- Never leave inconsistent state
