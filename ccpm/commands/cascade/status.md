---
description: Show current Cascade Flow session status
argument-hint: [feature-name]
allowed-tools: Read, Glob, Bash
---

<objective>
Display the current status of a Cascade Flow session including phase,
progress, and next steps. If no argument, shows all active sessions.
</objective>

<process>
## If Feature Name Provided

```bash
feature="$ARGUMENTS"

# Check brainstorm session
if session=$(find .claude/brainstorm -name "session.md" -exec grep -l "feature_name: $feature" {} \; | head -1); then
  session_id=$(basename "$(dirname "$session")")
  echo "## Brainstorm Session: $session_id"
  grep -E '^(status|phase|feature_name):' "$session"
fi

# Check PRD
if [ -f ".claude/prds/${feature}.md" ]; then
  echo "## PRD Status"
  grep -E '^(status|created|updated):' ".claude/prds/${feature}.md" | head -5
fi

# Check Epic
if [ -d ".claude/epics/${feature}" ]; then
  echo "## Epic Status"
  if [ -f ".claude/epics/${feature}/epic.md" ]; then
    grep -E '^(status|current_wave):' ".claude/epics/${feature}/epic.md"
  fi
  if [ -f ".claude/epics/${feature}/task-board.md" ]; then
    echo "## Task Board"
    grep -A5 '## Summary' ".claude/epics/${feature}/task-board.md"
  fi
fi
```

## If No Argument (Show All)

```bash
echo "## Active Cascade Sessions"
echo ""

for session in .claude/brainstorm/*/session.md; do
  if [ -f "$session" ]; then
    feature=$(grep '^feature_name:' "$session" | cut -d: -f2 | tr -d ' ')
    status=$(grep '^status:' "$session" | cut -d: -f2 | tr -d ' ')
    phase=$(grep '^phase:' "$session" | cut -d: -f2 | tr -d ' ')
    echo "- $feature: $status ($phase)"
  fi
done

echo ""
echo "## Active Epics"
for epic in .claude/epics/*/epic.md; do
  if [ -f "$epic" ]; then
    name=$(grep '^name:' "$epic" | cut -d: -f2 | tr -d ' ')
    status=$(grep '^status:' "$epic" | cut -d: -f2 | tr -d ' ')
    echo "- $name: $status"
  fi
done
```

</process>

<output_format>
## Cascade Status: {feature_name}

### Current Phase
{Phase name and status}

### Progress
| Phase | Status |
|-------|--------|
| 1. Brainstorm | {status} |
| 2. PRD | {status} |
| 3. Decompose | {status} |
| 4. Execute | {status} |

### Next Steps
{Recommended action}

### Files
- Session: .claude/brainstorm/{session_id}/
- PRD: .claude/prds/{feature_name}.md
- Epic: .claude/epics/{feature_name}/
</output_format>

<success_criteria>
- Status accurately reflects current state
- All relevant files located
- Progress clearly displayed
- Next steps recommended
</success_criteria>
