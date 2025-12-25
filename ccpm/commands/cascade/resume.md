---
description: Resume a Cascade Flow session from checkpoint
argument-hint: <session-id-or-feature-name>
allowed-tools: Skill, Task, Read, Write, Glob, Bash
---

<objective>
Resume an existing Cascade Flow session from its last checkpoint.
Automatically detects current phase and continues from there.
</objective>

<process>
1. Find session by ID or feature name
2. Read session state to determine current phase
3. Resume from last incomplete step

## Phase Detection

```bash
# Find session
if [ -d ".claude/brainstorm/$ARGUMENTS" ]; then
  session_id="$ARGUMENTS"
  session_file=".claude/brainstorm/${session_id}/session.md"
elif session_dir=$(find .claude/brainstorm -name "session.md" -exec grep -l "feature_name: $ARGUMENTS" {} \; | head -1); then
  session_id=$(basename "$(dirname "$session_dir")")
  session_file="$session_dir"
fi

# Read current phase
if [ -f "$session_file" ]; then
  phase=$(grep '^phase:' "$session_file" | cut -d: -f2 | tr -d ' ')
  status=$(grep '^status:' "$session_file" | cut -d: -f2 | tr -d ' ')
  feature_name=$(grep '^feature_name:' "$session_file" | cut -d: -f2 | tr -d ' ')
fi
```

## Resume Logic

| Status | Action |
|--------|--------|
| active (Phase 1) | Load cascade-brainstorm.md |
| phase_1_complete | Load cascade-prd.md |
| phase_2_complete | Load cascade-decompose.md |
| ready_for_execution | Load cascade-execute.md |
| completed | Report completion, offer sync |

</process>

Resume session for: $ARGUMENTS

<success_criteria>
- Session located and loaded
- Current phase identified
- Workflow resumed from correct point
- User informed of session state
</success_criteria>
