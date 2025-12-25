---
name: cascade-brainstorm
description: |
  Phase 1 of Cascade Flow: Multi-specialist parallel brainstorming with
  background research and idea incubation. Spawns 5 perspective agents
  simultaneously while a researcher scans the codebase in the background.
allowed-tools: Task, Read, Write, Glob, Bash, AskUserQuestion, TodoWrite
---

# Workflow: Cascade Brainstorm (Phase 1)

Launch multi-specialist parallel brainstorming session with background research
and idea incubation. This is the first phase of the Cascade Flow.

## Input
- `$ARGUMENTS`: Feature name (kebab-case) or description

## Preflight Checks

### 1. Validate or Generate Feature Name
```bash
# If argument looks like a description, extract name
if echo "$ARGUMENTS" | grep -qE '\s'; then
  # Multi-word - likely a description, will generate name
  FEATURE_NAME="pending"
  FEATURE_DESCRIPTION="$ARGUMENTS"
else
  # Single word - likely a name
  FEATURE_NAME="$ARGUMENTS"
  FEATURE_DESCRIPTION=""
fi
```

If `$ARGUMENTS` is a description, ask user:
"What should we call this feature? (kebab-case, e.g., `user-authentication`)"

### 2. Validate Feature Name Format
```bash
echo "$FEATURE_NAME" | grep -qE '^[a-z][a-z0-9]*(-[a-z0-9]+)*$'
```
If invalid: "Feature name must be kebab-case (e.g., `my-feature`)"

### 3. Check for Existing Session
```bash
existing=$(find .claude/brainstorm -name "session.md" -exec grep -l "feature_name: $FEATURE_NAME" {} \; 2>/dev/null | head -1)
if [ -n "$existing" ]; then
  session_dir=$(dirname "$existing")
  session_id=$(basename "$session_dir")
  echo "Existing session found: $session_id"
fi
```

If existing session found:
"A brainstorm session for `{feature_name}` already exists (session: `{session_id}`).
Would you like to:
1. **Resume** - Continue the existing session
2. **Start fresh** - Create a new session (old one preserved)"

### 4. Generate Session ID
```bash
session_id=$(cat /proc/sys/kernel/random/uuid 2>/dev/null || uuidgen | tr '[:upper:]' '[:lower:]')
```

### 5. Create Session Directory Structure
```bash
mkdir -p ".claude/brainstorm/${session_id}"/{perspectives,research,synthesis,interactions}
```

## Execution Steps

### Step 1: Initialize Session File

Create `.claude/brainstorm/${session_id}/session.md`:

```markdown
---
session_id: {session_id}
feature_name: {feature_name}
status: active
phase: gathering_context
created: {datetime}
updated: {datetime}
perspectives:
  user_advocate: pending
  tech_skeptic: pending
  innovation_catalyst: pending
  risk_assessor: pending
  simplicity_champion: pending
research:
  status: incubating
  started: null
  files_scanned: 0
  patterns_found: 0
synthesis:
  status: pending
  themes_identified: 0
  conflicts_detected: 0
  recommendations: 0
approval:
  status: pending
  approved_at: null
  revisions: 0
---

# Brainstorm Session: {feature_name}

## Initial Context
{To be filled after gathering context}

## Session Progress
- [ ] Context gathered
- [ ] Perspectives launched
- [ ] Research started
- [ ] Perspectives complete
- [ ] Synthesis complete
- [ ] User approved
```

### Step 2: Gather Initial Context

Ask user:

"Tell me about what you want to build. Describe:
- **The problem** you're trying to solve
- **Who** will use this feature
- **What success** looks like

Take your time - your description will guide 5 specialist agents who will
explore this idea from different angles."

**Wait for response.**

After response, update session.md:
- Add user's description under "Initial Context"
- Update phase to "perspectives_launching"
- Update timestamp

### Step 3: Launch Perspective Agents (5 in Parallel)

Inform user:
"Launching 5 specialist agents to explore your idea:
- **User Advocate** (cyan): Championing user experience
- **Tech Skeptic** (red): Questioning complexity
- **Innovation Catalyst** (magenta): Finding breakthroughs
- **Risk Assessor** (yellow): Identifying risks
- **Simplicity Champion** (green): Advocating minimalism

Plus a **Codebase Researcher** scanning your project in the background..."

Spawn all 6 agents in parallel:

```yaml
# Agent 1: User Advocate
Task:
  description: "Brainstorm: User Advocate"
  subagent_type: "user-advocate"
  run_in_background: true
  prompt: |
    Session: {session_id}
    Feature: {feature_name}

    ## User's Description
    {initial_context}

    ## Your Task
    Analyze this feature idea from the user's perspective.
    Focus on UX, pain points, usability, and user value.

    Write your analysis to:
    .claude/brainstorm/{session_id}/perspectives/user-advocate.md

    Follow the output format in your agent definition.

# Agent 2: Tech Skeptic
Task:
  description: "Brainstorm: Tech Skeptic"
  subagent_type: "tech-skeptic"
  run_in_background: true
  prompt: |
    Session: {session_id}
    Feature: {feature_name}

    ## User's Description
    {initial_context}

    ## Your Task
    Critically analyze this feature idea.
    Question complexity, challenge over-engineering, suggest proven alternatives.

    Write your analysis to:
    .claude/brainstorm/{session_id}/perspectives/tech-skeptic.md

# Agent 3: Innovation Catalyst
Task:
  description: "Brainstorm: Innovation Catalyst"
  subagent_type: "innovation-catalyst"
  run_in_background: true
  prompt: |
    Session: {session_id}
    Feature: {feature_name}

    ## User's Description
    {initial_context}

    ## Your Task
    Explore creative and unconventional approaches.
    Find breakthrough opportunities and adjacent possibilities.

    Write your analysis to:
    .claude/brainstorm/{session_id}/perspectives/innovation-catalyst.md

# Agent 4: Risk Assessor
Task:
  description: "Brainstorm: Risk Assessor"
  subagent_type: "risk-assessor"
  run_in_background: true
  prompt: |
    Session: {session_id}
    Feature: {feature_name}

    ## User's Description
    {initial_context}

    ## Your Task
    Identify all risks: security, performance, scalability, reliability.
    Provide mitigations for each risk.

    Write your analysis to:
    .claude/brainstorm/{session_id}/perspectives/risk-assessor.md

# Agent 5: Simplicity Champion
Task:
  description: "Brainstorm: Simplicity Champion"
  subagent_type: "simplicity-champion"
  run_in_background: true
  prompt: |
    Session: {session_id}
    Feature: {feature_name}

    ## User's Description
    {initial_context}

    ## Your Task
    Find the minimum viable solution.
    Identify what can be eliminated, simplified, or deferred.

    Write your analysis to:
    .claude/brainstorm/{session_id}/perspectives/simplicity-champion.md

# Agent 6: Codebase Researcher (Background)
Task:
  description: "Background: Codebase Research"
  subagent_type: "codebase-researcher"
  run_in_background: true
  timeout: 600000
  prompt: |
    Session: {session_id}
    Feature: {feature_name}

    ## User's Description
    {initial_context}

    ## Your Task
    Scan the codebase for relevant patterns, prior art, and integration points.
    Run continuously in background, updating findings as you discover them.

    Write findings to:
    .claude/brainstorm/{session_id}/research/

    Files to create:
    - patterns.md
    - relevant-files.md
    - prior-art.md
```

Update session.md:
- Set all perspective statuses to "running"
- Set research.status to "incubating"
- Set research.started to current datetime

### Step 4: User Interaction While Agents Work

While agents are running, engage user:

"While the specialists explore your idea, let me ask a few clarifying questions:

1. **Priority**: Is speed of delivery or long-term maintainability more important?
2. **Scope**: Are you thinking MVP or full-featured?
3. **Constraints**: Any technical or business constraints I should know about?"

Record responses in `.claude/brainstorm/{session_id}/interactions/{timestamp}-clarifications.md`

### Step 5: Monitor Agent Completion

Check for completion every 30 seconds:

```bash
for perspective in user-advocate tech-skeptic innovation-catalyst risk-assessor simplicity-champion; do
  if [ -s ".claude/brainstorm/${session_id}/perspectives/${perspective}.md" ]; then
    echo "${perspective}: complete"
  else
    echo "${perspective}: pending"
  fi
done
```

Update session.md as each perspective completes.

When all 5 perspectives are complete, proceed to synthesis.

### Step 6: Launch Synthesis Engine

```yaml
Task:
  description: "Synthesize perspectives"
  subagent_type: "synthesis-engine"
  prompt: |
    Session: {session_id}
    Feature: {feature_name}

    ## Your Task
    Read all perspective files and research findings.
    Produce synthesis documents.

    Input directories:
    - .claude/brainstorm/{session_id}/perspectives/
    - .claude/brainstorm/{session_id}/research/

    Output directory:
    - .claude/brainstorm/{session_id}/synthesis/

    Create:
    - themes.md (common themes)
    - conflicts.md (disagreements)
    - recommendations.md (actionable synthesis)
```

Update session.md:
- Set synthesis.status to "running"

### Step 7: Present Synthesis to User

Read synthesis files and present to user:

"## Brainstorm Synthesis Complete

### Key Themes (Consensus)
{List themes from themes.md}

### Conflicts (Need Your Decision)
{List conflicts from conflicts.md with options}

### Recommendations
{Prioritized list from recommendations.md}

### Research Findings
{Summary from research/ directory}

---

**How would you like to proceed?**
1. **Steer** - \"I want to explore {direction} more\"
2. **Reject** - \"This isn't quite right, let's try {alternative}\"
3. **Expand** - \"Tell me more about {specific point}\"
4. **Approve** - \"This looks good, proceed to PRD\""

### Step 8: Handle User Response

#### If Steer:
```bash
echo "User steered toward: {direction}" > ".claude/brainstorm/${session_id}/interactions/$(date +%s)-steer.md"
```
- Re-spawn relevant perspective agent(s) with new direction
- Re-run synthesis
- Loop back to Step 7

#### If Reject:
```bash
echo "User rejected: {reason}" > ".claude/brainstorm/${session_id}/interactions/$(date +%s)-reject.md"
```
- Record rejection reason
- Spawn idea-incubator to explore alternatives
- Re-gather context if needed
- Loop back to Step 3

#### If Expand:
```bash
echo "User requested expansion: {topic}" > ".claude/brainstorm/${session_id}/interactions/$(date +%s)-expand.md"
```
- Spawn targeted deep-dive agent
- Present expanded analysis
- Loop back to Step 7

#### If Approve:
```bash
# Update session status
# Create approval record
```
- Update session.md: approval.status = "approved", approval.approved_at = datetime
- Create checkpoint
- Proceed to Phase 2 (PRD Generation)

### Step 9: Transition to Phase 2

Update session.md:
- status: "phase_1_complete"
- phase: "ready_for_prd"

Inform user:
"Brainstorm approved! Your ideas have been captured.

**Session**: `{session_id}`
**Feature**: `{feature_name}`

Ready for Phase 2: PRD Generation. Run `/cascade:start {feature_name}` to continue,
or the workflow will proceed automatically."

Load workflow: cascade-prd.md with session context.

## Output Summary

```
Brainstorm Complete: {feature_name}

Session: .claude/brainstorm/{session_id}/

Perspectives Gathered:
  - User Advocate: {status}
  - Tech Skeptic: {status}
  - Innovation Catalyst: {status}
  - Risk Assessor: {status}
  - Simplicity Champion: {status}

Synthesis:
  Themes: {count}
  Conflicts: {count} (resolved: {count})
  Recommendations: {count}

Background Research:
  Patterns found: {count}
  Files analyzed: {count}
  Prior art: {count}

Status: Approved for PRD Generation
Revisions: {count}

Proceeding to Phase 2...
```

## Error Handling

### Agent Timeout
If perspective agent doesn't complete in 5 minutes:
- Report which agent timed out
- Continue with completed perspectives
- Note gap in synthesis

### Research Stall
If codebase researcher produces no output:
- Note in synthesis as "research unavailable"
- Proceed with perspective-only synthesis

### User Abandons
If user stops responding:
- Save session state
- Session can be resumed with `/cascade:resume {session_id}`

## Recovery Protocol

Session can be resumed at any point by reading session.md and:
1. Checking which perspectives are complete
2. Checking synthesis status
3. Resuming from last incomplete step
