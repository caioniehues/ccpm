---
name: cascade-prd
description: |
  Phase 2 of Cascade Flow: Parallel PRD section generation with consistency
  validation and tradeoff analysis. Transforms brainstorm synthesis into
  structured Product Requirements Document.
allowed-tools: Task, Read, Write, Glob, Edit, AskUserQuestion
---

# Workflow: Cascade PRD (Phase 2)

Generate a structured PRD from brainstorm synthesis using parallel section
writers, consistency validation, and background tradeoff analysis.

## Input
- Session from Phase 1 (brainstorm complete)
- `$ARGUMENTS`: Feature name or session ID

## Preflight Checks

### 1. Locate Session
```bash
# Find session by feature name or ID
if [ -d ".claude/brainstorm/$ARGUMENTS" ]; then
  session_id="$ARGUMENTS"
elif session_dir=$(find .claude/brainstorm -name "session.md" -exec grep -l "feature_name: $ARGUMENTS" {} \; | head -1); then
  session_id=$(basename "$(dirname "$session_dir")")
else
  echo "No session found for: $ARGUMENTS"
  exit 1
fi
```

### 2. Verify Phase 1 Complete
```bash
status=$(grep '^status:' ".claude/brainstorm/${session_id}/session.md" | cut -d: -f2 | tr -d ' ')
if [ "$status" != "phase_1_complete" ]; then
  echo "Phase 1 not complete. Current status: $status"
  exit 1
fi
```

### 3. Read Synthesis
Load all files from `.claude/brainstorm/${session_id}/synthesis/`:
- themes.md
- conflicts.md
- recommendations.md

### 4. Extract Feature Name
```bash
feature_name=$(grep '^feature_name:' ".claude/brainstorm/${session_id}/session.md" | cut -d: -f2 | tr -d ' ')
```

## Execution Steps

### Step 1: Initialize PRD File

Create `.claude/prds/${feature_name}.md` with frontmatter:

```markdown
---
name: {feature_name}
description: {one-line from synthesis}
status: drafting
created: {datetime}
updated: {datetime}
brainstorm_session: {session_id}
sections:
  executive_summary: pending
  requirements: pending
  success_metrics: pending
  dependencies: pending
  constraints: pending
validation:
  status: pending
  issues: 0
tradeoffs:
  status: incubating
  decisions: 0
---

# PRD: {Feature Title}

{Sections to be filled by parallel writers}
```

### Step 2: Launch Parallel Section Writers

Spawn 4 section writers simultaneously:

```yaml
# Writer 1: Executive Summary
Task:
  description: "PRD: Executive Summary"
  subagent_type: "section-writer"
  run_in_background: true
  prompt: |
    Feature: {feature_name}
    Session: {session_id}
    Section: executive_summary

    ## Source Material
    Read: .claude/brainstorm/{session_id}/synthesis/recommendations.md
    Read: .claude/brainstorm/{session_id}/synthesis/themes.md

    ## Your Task
    Write the Executive Summary section including:
    - Problem Statement (from synthesis themes)
    - Proposed Solution (from recommendations)
    - Success Criteria (measurable outcomes)
    - Target Users (from user-advocate perspective)

    ## Output Format
    Write ONLY the section content (no frontmatter) to:
    .claude/prds/{feature_name}.sections/executive_summary.md

# Writer 2: Requirements
Task:
  description: "PRD: Requirements"
  subagent_type: "section-writer"
  run_in_background: true
  prompt: |
    Feature: {feature_name}
    Session: {session_id}
    Section: requirements

    ## Source Material
    Read: .claude/brainstorm/{session_id}/synthesis/recommendations.md
    Read: .claude/brainstorm/{session_id}/perspectives/tech-skeptic.md
    Read: .claude/brainstorm/{session_id}/perspectives/simplicity-champion.md

    ## Your Task
    Write Requirements sections:
    - Functional Requirements (numbered, testable)
    - Non-Functional Requirements (performance, security, etc.)
    - Out of Scope (explicit exclusions)

    ## Output Format
    Write to: .claude/prds/{feature_name}.sections/requirements.md

# Writer 3: Success Metrics
Task:
  description: "PRD: Success Metrics"
  subagent_type: "section-writer"
  run_in_background: true
  prompt: |
    Feature: {feature_name}
    Session: {session_id}
    Section: success_metrics

    ## Source Material
    Read: .claude/brainstorm/{session_id}/synthesis/recommendations.md
    Read: .claude/brainstorm/{session_id}/perspectives/user-advocate.md

    ## Your Task
    Define measurable success criteria:
    - Acceptance Criteria (must-have for launch)
    - Key Performance Indicators (ongoing measurement)
    - Definition of Done (completion checklist)

    ## Output Format
    Write to: .claude/prds/{feature_name}.sections/success_metrics.md

# Writer 4: Dependencies & Constraints
Task:
  description: "PRD: Dependencies"
  subagent_type: "section-writer"
  run_in_background: true
  prompt: |
    Feature: {feature_name}
    Session: {session_id}
    Section: dependencies

    ## Source Material
    Read: .claude/brainstorm/{session_id}/synthesis/conflicts.md
    Read: .claude/brainstorm/{session_id}/perspectives/risk-assessor.md
    Read: .claude/brainstorm/{session_id}/research/relevant-files.md

    ## Your Task
    Document dependencies and constraints:
    - Internal Dependencies (existing code, services)
    - External Dependencies (APIs, libraries)
    - Technical Constraints
    - Business Constraints
    - Assumptions

    ## Output Format
    Write to: .claude/prds/{feature_name}.sections/dependencies.md
```

### Step 3: Launch Background Tradeoff Analysis

```yaml
Task:
  description: "Background: Tradeoff Analysis"
  subagent_type: "tradeoff-analyst"
  run_in_background: true
  timeout: 300000
  prompt: |
    Feature: {feature_name}
    Session: {session_id}

    ## Source Material
    Read all files in: .claude/brainstorm/{session_id}/

    ## Your Task
    Analyze architectural and design tradeoffs.
    Identify key decisions that need to be made.

    ## Output
    Write to: .claude/prds/{feature_name}.sections/tradeoffs.md
```

### Step 4: Monitor Section Completion

Wait for all 4 section writers to complete:

```bash
for section in executive_summary requirements success_metrics dependencies; do
  while [ ! -s ".claude/prds/${feature_name}.sections/${section}.md" ]; do
    sleep 5
  done
  echo "${section}: complete"
done
```

### Step 5: Assemble PRD

Combine all sections into the main PRD file:

```bash
cat > ".claude/prds/${feature_name}.md" << 'EOF'
{frontmatter}

# PRD: {Feature Title}

## Executive Summary
{content from executive_summary.md}

## Requirements
{content from requirements.md}

## Success Metrics
{content from success_metrics.md}

## Dependencies & Constraints
{content from dependencies.md}

## Tradeoff Analysis
{content from tradeoffs.md - if available}

---
*Generated by Cascade Flow from brainstorm session {session_id}*
EOF
```

### Step 6: Run Consistency Validation

```yaml
Task:
  description: "Validate PRD Consistency"
  subagent_type: "consistency-validator"
  prompt: |
    Feature: {feature_name}

    ## Your Task
    Read the complete PRD: .claude/prds/{feature_name}.md

    Validate:
    1. Cross-section consistency
    2. Terminology consistency
    3. Requirement conflicts
    4. Missing references
    5. Success criteria coverage

    ## Output
    Write validation report to:
    .claude/prds/{feature_name}.validation.md

    Format:
    - PASS: {item} (if valid)
    - FAIL: {item} - {issue} (if invalid)
    - WARN: {item} - {concern} (if questionable)
```

### Step 7: Handle Validation Issues

If validation finds FAIL items:
- Present issues to user
- Offer to fix automatically or manually
- Re-run validation after fixes

### Step 8: Present PRD to User

Display the complete PRD with diff-based editing option:

"## PRD Generated: {feature_name}

{Display PRD content with section headers}

---

**Validation Status**: {PASS/FAIL with count}

**Tradeoffs Identified**: {count from tradeoffs.md}

---

**Options:**
1. **Approve** - Proceed to task decomposition
2. **Edit** - Make specific changes (diff-based)
3. **Regenerate** - Re-run with different emphasis
4. **Back** - Return to brainstorm for more exploration"

### Step 9: Handle User Response

#### If Approve:
- Update PRD status to "approved"
- Update brainstorm session status
- Proceed to Phase 3

#### If Edit:
- Ask which section to modify
- Present current content
- Accept changes (diff-based using Edit tool)
- Re-run consistency validation
- Loop back to Step 8

#### If Regenerate:
- Ask for different emphasis/focus
- Re-run section writers with new guidance
- Loop back to Step 5

#### If Back:
- Return to brainstorm workflow
- Allow additional exploration

### Step 10: Transition to Phase 3

Update PRD frontmatter:
- status: "approved"
- approved_at: {datetime}

Update brainstorm session:
- status: "phase_2_complete"

"PRD approved! Ready for task decomposition.

Proceeding to Phase 3: Decomposition..."

Load workflow: cascade-decompose.md

## Output Summary

```
PRD Generated: {feature_name}

Sections:
  - Executive Summary: {status}
  - Requirements: {status}
  - Success Metrics: {status}
  - Dependencies: {status}
  - Tradeoffs: {status}

Validation:
  - Passed: {count}
  - Failed: {count}
  - Warnings: {count}

Status: Approved
Location: .claude/prds/{feature_name}.md

Proceeding to Phase 3...
```

## Error Handling

### Section Writer Timeout
If any section writer times out:
- Use synthesis content directly for that section
- Mark section as "simplified"
- Continue with available sections

### Validation Failure
If validation finds critical issues:
- Block approval until resolved
- Offer automatic fix suggestions
- Allow manual override with acknowledgment
