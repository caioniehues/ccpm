# Progress Tracking

<overview>
CCPM tracks work progress through structured files in the updates directory. Progress files maintain state, sync history, and work details.
</overview>

<directory_structure>
## Progress Directory Structure

```
.claude/epics/{epic}/updates/{issue}/
├── progress.md          # Overall progress tracking
├── notes.md             # Technical notes and decisions
├── commits.md           # Commit summaries
├── stream-A.md          # Stream A progress
├── stream-B.md          # Stream B progress
└── stream-C.md          # Stream C progress
```

**Created when:**
- Issue work starts (`/pm:issue-start`)

**Updated during:**
- Work progress
- GitHub sync
- Issue completion
</directory_structure>

<progress_file_format>
## Progress File Format

**progress.md structure:**
```markdown
---
issue: 123
started: 2024-01-15T10:30:00Z
last_sync: 2024-01-15T14:45:00Z
completion: 65
---

# Progress Tracking: Issue #123

## Completed Work
- Implemented user authentication
- Created database schema
- Added API endpoints

## In Progress
- Writing unit tests
- Updating documentation

## Next Steps
- Integration testing
- Code review
- Deploy to staging

## Blockers
None currently

## Technical Decisions
- Using JWT for authentication
- PostgreSQL for user data
- Express middleware for auth checks
```

**Frontmatter fields:**
- `issue`: Issue number (integer)
- `started`: When work began (ISO 8601 UTC)
- `last_sync`: Last GitHub sync (ISO 8601 UTC)
- `completion`: Progress percentage (0-100 integer)
</progress_file_format>

<notes_file_format>
## Notes File Format

**notes.md structure:**
```markdown
---
issue: 123
---

# Technical Notes: Issue #123

## 2024-01-15 10:30 UTC

### Decision: Authentication Strategy
Using JWT tokens with 24-hour expiration. Refresh tokens stored in secure HTTP-only cookies.

**Rationale:**
- Better security than session storage
- Supports mobile apps
- Allows stateless scaling

**Alternatives considered:**
- Session-based auth (too stateful)
- OAuth only (too complex for v1)

## 2024-01-15 14:00 UTC

### Implementation Note: Database Schema
Added unique constraint on email field. Using bcrypt for password hashing with cost factor 12.

**References:**
- OWASP password storage guidelines
- Postgres performance docs
```

**Format:**
- Reverse chronological order (newest first)
- Timestamp each entry
- Document decisions with rationale
- Link to relevant resources
</notes_file_format>

<commits_file_format>
## Commits File Format

**commits.md structure:**
```markdown
---
issue: 123
---

# Commit History: Issue #123

## 2024-01-15 14:45 UTC
- `a1b2c3d` Issue #123: Add user authentication endpoints
- `e4f5g6h` Issue #123: Implement JWT token generation
- `i7j8k9l` Issue #123: Add password hashing with bcrypt

## 2024-01-15 12:30 UTC
- `m0n1o2p` Issue #123: Create user database schema
- `q3r4s5t` Issue #123: Add migration for users table
```

**Format:**
- Group by date/time
- Include commit hash and message
- Link to full diff if needed
</commits_file_format>

<stream_file_format>
## Stream File Format

See parallel-work.md for detailed stream file structure.

**Quick reference:**
```markdown
---
issue: 123
stream: Stream Name
agent: agent-type
started: 2024-01-15T10:30:00Z
status: in_progress|completed|blocked
---

# Stream A: {Stream Name}

## Scope
{What this stream handles}

## Files
{Files this stream owns}

## Progress
- [x] Completed item
- [ ] Pending item

## Coordination Notes
{Notes about coordination with other streams}

## Blockers
{Current blockers if any}
```
</stream_file_format>

<updating_progress>
## Updating Progress

**Update completion percentage:**
```bash
issue=123
epic="epic-name"
progress_file=".claude/epics/$epic/updates/$issue/progress.md"

# Calculate new completion (example: 3 of 5 criteria met)
new_completion=$((3 * 100 / 5))  # 60%

# Get current datetime
datetime=$(date -u +"%Y-%m-%dT%H:%M:%SZ")

# Update frontmatter
sed -i "s/^completion: .*/completion: $new_completion/" "$progress_file"
sed -i "s/^last_sync: .*/last_sync: $datetime/" "$progress_file"
```

**Add progress entry:**
```bash
progress_file=".claude/epics/$epic/updates/$issue/progress.md"

# Append new completed work
cat >> "$progress_file" << EOF

- Implemented feature X
- Fixed bug Y
EOF
```

**Calculate completion from acceptance criteria:**
```bash
# Count total criteria
total=$(grep -c "^- \[ \]\\|^- \[x\]" task.md)

# Count completed criteria
completed=$(grep -c "^- \[x\]" task.md)

# Calculate percentage
if [ $total -gt 0 ]; then
  completion=$((completed * 100 / total))
else
  completion=0
fi

echo "Completion: $completion% ($completed/$total criteria)"
```
</updating_progress>

<sync_markers>
## Sync Markers

**Purpose:**
Track what's been synced to GitHub to avoid duplicate comments.

**Add sync marker:**
```bash
datetime=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
progress_file=".claude/epics/$epic/updates/$issue/progress.md"

# Add marker at bottom
echo "<!-- SYNCED: $datetime -->" >> "$progress_file"
```

**Find last sync:**
```bash
# Extract last sync timestamp
last_sync=$(grep "<!-- SYNCED:" "$progress_file" | tail -1 | grep -oE '[0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}Z')

if [ -z "$last_sync" ]; then
  echo "Never synced"
else
  echo "Last synced: $last_sync"
fi
```

**Get new content since last sync:**
```bash
# Get line number of last sync marker
last_line=$(grep -n "<!-- SYNCED:" "$progress_file" | tail -1 | cut -d: -f1)

if [ -z "$last_line" ]; then
  # Never synced, get all content
  cat "$progress_file"
else
  # Get content after last sync
  tail -n +$((last_line + 1)) "$progress_file"
fi
```
</sync_markers>

<progress_calculation>
## Progress Calculation Methods

**Method 1: Acceptance criteria**
```bash
# From task.md checkboxes
total=$(grep -c "^- \[ \]\\|^- \[x\]" task.md)
completed=$(grep -c "^- \[x\]" task.md)
completion=$((completed * 100 / total))
```

**Method 2: Stream completion**
```bash
# From parallel streams
total_streams=$(ls .claude/epics/$epic/updates/$issue/stream-*.md | wc -l)
completed_streams=$(grep -l "^status: completed" .claude/epics/$epic/updates/$issue/stream-*.md | wc -l)
completion=$((completed_streams * 100 / total_streams))
```

**Method 3: Manual estimate**
```bash
# Developer sets based on judgement
completion=75  # 75% complete
```

**Method 4: Time-based**
```bash
# Based on estimated vs actual hours
estimated_hours=8
actual_hours=6
completion=$((actual_hours * 100 / estimated_hours))
# Cap at 100%
[ $completion -gt 100 ] && completion=100
```
</progress_calculation>

<progress_states>
## Progress States

**0% - Not started:**
- Task exists but work hasn't begun
- No progress.md file yet

**1-99% - In progress:**
- Work actively happening
- Regular updates to progress files
- Sync to GitHub periodically

**100% - Complete:**
- All acceptance criteria met
- All streams completed
- Ready for close

**Blocked:**
- Progress stalled due to blocker
- Document blocker in progress file
- Update when blocker resolved
</progress_states>

<best_practices>
## Best Practices

1. **Update frequently**: Keep progress current, don't batch updates
2. **Use real timestamps**: Get actual datetime, never placeholders
3. **Document decisions**: Add notes for important technical choices
4. **Track blockers**: Document what's blocking progress
5. **Sync regularly**: Push updates to GitHub for transparency
6. **Calculate objectively**: Base completion on measurable criteria
7. **Preserve history**: Don't delete old progress, append new
8. **Use sync markers**: Avoid duplicate GitHub comments
9. **Separate concerns**: Use different files for different content types
10. **Validate percentages**: Ensure 0-100 range, no decimals

**Update workflow:**
1. Do work
2. Update progress files
3. Calculate new completion %
4. Sync to GitHub
5. Repeat
</best_practices>

<success_criteria>
Progress tracking is working correctly when:
- Progress files created when work starts
- Frontmatter valid and up-to-date
- Completion percentage accurate (0-100)
- Timestamps real (not placeholders)
- Sync markers prevent duplicate comments
- Notes document technical decisions
- Commits tracked with hashes
- Stream progress independent
- History preserved (not overwritten)
- GitHub sync reflects current state
</success_criteria>
