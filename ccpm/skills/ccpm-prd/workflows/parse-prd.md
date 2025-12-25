# Workflow: Parse PRD to Epic

Convert a Product Requirements Document into a technical implementation epic.

## Input
- `$ARGUMENTS`: Feature name (kebab-case)

## Preflight Checks

1. **Verify PRD Exists:**
   ```bash
   test -f .claude/prds/$ARGUMENTS.md
   ```
   If not found: "PRD not found: .claude/prds/$ARGUMENTS.md. Create with /pm:prd-new first."

2. **Read PRD:**
   ```bash
   cat .claude/prds/$ARGUMENTS.md
   ```
   Parse all sections for technical analysis.

3. **Check Existing Epic:**
   ```bash
   test -f .claude/epics/$ARGUMENTS/epic.md
   ```
   If exists: 
   - Ask: "Epic already exists. Overwrite? (yes/no)"
   - Only proceed with explicit 'yes'

4. **Create Epic Directory:**
   ```bash
   mkdir -p .claude/epics/$ARGUMENTS
   ```

5. **Get Current DateTime:**
   ```bash
   date -u +"%Y-%m-%dT%H:%M:%SZ"
   ```

## Technical Analysis

Analyze the PRD and determine:

### 1. Architecture Approach
Based on requirements, identify:
- System components needed
- Data models required
- API endpoints/interfaces
- Integration points
- Technology choices

### 2. Task Categories
Group work into logical categories (max 10):
- Infrastructure/Setup
- Data Layer
- Business Logic
- API/Interface
- UI/Frontend
- Integration
- Testing
- Documentation
- Deployment
- Migration

### 3. Dependency Mapping
Identify:
- What must be built first (prerequisites)
- What can be built in parallel
- External dependencies to resolve

### 4. Effort Estimation
For each category:
- XS: < 1 hour
- S: 1-2 hours
- M: 2-4 hours
- L: 4-8 hours
- XL: > 8 hours (flag for splitting)

## Create Epic File

Create `.claude/epics/$ARGUMENTS/epic.md`:

```markdown
---
name: $ARGUMENTS
status: planning
created: {datetime}
updated: {datetime}
prd: .claude/prds/$ARGUMENTS.md
github: null
worktrees: null
---

# Epic: {Feature Title from PRD}

## Overview

### Source PRD
{Link to PRD and brief summary of what we're building}

### Objective
{Technical objective derived from PRD problem statement}

### Success Criteria
{Technical success criteria derived from PRD}

## Technical Approach

### Architecture
{System architecture decisions}

### Components
1. **{Component 1}**: {purpose and approach}
2. **{Component 2}**: {purpose and approach}
3. **{Component 3}**: {purpose and approach}

### Data Models
{Key data structures and their relationships}

### APIs/Interfaces
{External interfaces this feature exposes or consumes}

### Technology Choices
| Area | Choice | Rationale |
|------|--------|-----------|
| {area} | {technology} | {why} |

## Task Breakdown Preview

### Phase 1: Foundation
- [ ] {Task category 1} [{effort}]
- [ ] {Task category 2} [{effort}]

### Phase 2: Core Implementation
- [ ] {Task category 3} [{effort}]
- [ ] {Task category 4} [{effort}]

### Phase 3: Integration & Polish
- [ ] {Task category 5} [{effort}]
- [ ] {Task category 6} [{effort}]

**Total Estimated Effort**: {sum of efforts}

## Dependencies

### Prerequisites
- {What must exist before starting}

### External Dependencies
- {Third-party services, APIs to integrate}

### Internal Dependencies
- {Other epics or features this depends on}

## Risks and Mitigations

| Risk | Impact | Mitigation |
|------|--------|------------|
| {risk} | {H/M/L} | {mitigation strategy} |

## Open Questions

- [ ] {Question that needs resolution before implementation}
- [ ] {Another question}

## Definition of Done

- [ ] All tasks completed and merged
- [ ] Tests passing with adequate coverage
- [ ] Documentation updated
- [ ] PRD acceptance criteria verified
- [ ] Deployed to production (if applicable)
```

## Update PRD Status

Update the PRD to reflect that epic was created:

```yaml
---
name: $ARGUMENTS
description: {unchanged}
status: in-progress
created: {PRESERVE}
updated: {new datetime}
epic: .claude/epics/$ARGUMENTS/epic.md
---
```

## Post-Creation

1. **Validate Epic:**
   ```bash
   test -s .claude/epics/$ARGUMENTS/epic.md
   wc -l .claude/epics/$ARGUMENTS/epic.md  # Should be 50+ lines
   ```

2. **Display Summary:**
   ```
   ✅ Epic Created: .claude/epics/$ARGUMENTS/epic.md
   
   📋 Technical Summary:
      Feature: {name}
      Components: {count}
      Task Categories: {count}
      Estimated Effort: {total}
   
   📊 Task Breakdown:
      Phase 1: {count} tasks (foundation)
      Phase 2: {count} tasks (core)
      Phase 3: {count} tasks (polish)
   
   ⚠️ Open Questions: {count} items need resolution
   
   ⏭️ Next Steps:
      1. Review epic and resolve open questions
      2. When ready: /pm:epic-decompose $ARGUMENTS
   ```

## Error Handling

- **PRD not found**: Suggest `/pm:prd-new` or `/pm:prd-list`
- **Epic directory creation fails**: Check permissions
- **Insufficient PRD content**: Request more information from user
- **File write fails**: Report error, clean up partial files

## Using prd-architect Agent

For complex PRDs, delegate to prd-architect agent:

```yaml
Task:
  description: "Analyze PRD and create epic structure"
  subagent_type: "prd-architect"
  prompt: |
    Analyze PRD at .claude/prds/$ARGUMENTS.md
    Create comprehensive epic at .claude/epics/$ARGUMENTS/epic.md
    
    Include:
    - Technical architecture decisions
    - Component breakdown
    - Task categories with effort estimates
    - Dependency mapping
    - Risk assessment
```

## Success Criteria

PRD parsing is complete when:
- [ ] PRD exists and was fully read
- [ ] User confirmed overwrite (if epic existed)
- [ ] Epic directory created
- [ ] epic.md created with valid frontmatter
- [ ] Technical approach documented
- [ ] Task breakdown preview included (≤10 categories)
- [ ] Dependencies mapped
- [ ] Risks identified
- [ ] PRD status updated to in-progress
- [ ] Real datetime used everywhere
- [ ] Summary provided with next steps
