# Workflow: Create New PRD

Create a comprehensive Product Requirements Document through structured brainstorming.

## Input
- `$ARGUMENTS`: Feature name (kebab-case)

## Preflight Checks

1. **Validate Feature Name:**
   ```bash
   # Check kebab-case format
   echo "$ARGUMENTS" | grep -qE '^[a-z][a-z0-9]*(-[a-z0-9]+)*$'
   ```
   If invalid: "Feature name must be kebab-case (e.g., user-authentication)"

2. **Check Existing PRD:**
   ```bash
   test -f .claude/prds/$ARGUMENTS.md
   ```
   If exists: "PRD already exists. Use edit operation or choose different name."

3. **Ensure Directory:**
   ```bash
   mkdir -p .claude/prds
   ```

4. **Get Current DateTime:**
   ```bash
   date -u +"%Y-%m-%dT%H:%M:%SZ"
   ```
   Store for frontmatter.

## Brainstorming Session

Conduct a structured discovery session. Ask questions in this order:

### Phase 1: Problem Definition
Ask the user:
1. "What problem does this feature solve?"
2. "Who experiences this problem? (target users)"
3. "What happens if we don't build this?"

### Phase 2: Solution Exploration
Ask the user:
1. "How should users interact with this feature?"
2. "What are the 3-5 core capabilities?"
3. "What should this feature explicitly NOT do? (out of scope)"

### Phase 3: Success Criteria
Ask the user:
1. "How will we know this feature is successful?"
2. "What are the measurable acceptance criteria?"
3. "Are there any constraints (technical, timeline, budget)?"

### Phase 4: Dependencies
Ask the user:
1. "What existing systems does this integrate with?"
2. "Are there any external dependencies?"
3. "What needs to exist before we can build this?"

**IMPORTANT**: Wait for user response after each question before proceeding.

## Create PRD File

After brainstorming, create `.claude/prds/$ARGUMENTS.md`:

```markdown
---
name: $ARGUMENTS
description: {One-line summary from brainstorming}
status: backlog
created: {datetime from preflight}
updated: {datetime from preflight}
---

# PRD: {Feature Title}

## Executive Summary

### Problem Statement
{From Phase 1 answers}

### Proposed Solution
{From Phase 2 answers - 1-2 paragraphs}

### Success Criteria
{From Phase 3 answers - high level}

## Context and Background

### Current State
{How things work today without this feature}

### Target Users
{From Phase 1 - who experiences the problem}

### Business Objectives
{Why this matters to the business}

## User Stories

### Primary User Story
**As a** {user type}
**I want to** {action}
**So that** {benefit}

**Acceptance Criteria:**
- [ ] {Criterion 1}
- [ ] {Criterion 2}
- [ ] {Criterion 3}

### Additional User Stories
{Additional stories from brainstorming}

## Requirements

### Functional Requirements
FR-1: {Requirement with clear acceptance criteria}
FR-2: {Next requirement}
FR-3: {Next requirement}

### Non-Functional Requirements
NFR-1: {Performance, security, scalability requirement}
NFR-2: {Next requirement}

### Constraints
- {Technical constraints from Phase 3}
- {Timeline constraints}
- {Resource constraints}

### Out of Scope
- {Explicit exclusion 1 from Phase 2}
- {Explicit exclusion 2}

## Dependencies

### Internal Dependencies
- {Systems this integrates with}

### External Dependencies
- {Third-party services, APIs}

### Prerequisites
- {What must exist first}

## Success Metrics

### Acceptance Criteria
- [ ] {Measurable criterion 1}
- [ ] {Measurable criterion 2}
- [ ] {Measurable criterion 3}

### KPIs
- {Metric to track success}
```

## Post-Creation

1. **Validate File:**
   ```bash
   test -s .claude/prds/$ARGUMENTS.md
   wc -l .claude/prds/$ARGUMENTS.md  # Should be 50+ lines
   ```

2. **Display Summary:**
   ```
   ✅ PRD Created: .claude/prds/$ARGUMENTS.md
   
   📋 Summary:
      Feature: {name}
      Status: backlog
      User Stories: {count}
      Requirements: {count}
   
   ⏭️ Next Steps:
      1. Review and refine the PRD
      2. When ready: /pm:prd-parse $ARGUMENTS
   ```

## Error Handling

- **Directory creation fails**: Check permissions, suggest manual creation
- **File write fails**: Report error, don't leave partial file
- **User abandons brainstorming**: Ask if they want to save partial progress

## Success Criteria

PRD creation is complete when:
- [ ] Feature name validated as kebab-case
- [ ] No existing PRD with same name
- [ ] Brainstorming session completed (all phases)
- [ ] PRD file created with valid frontmatter
- [ ] Real datetime used (not placeholder)
- [ ] All sections populated with content (no placeholders)
- [ ] File validated (exists, non-empty, 50+ lines)
- [ ] Next steps provided to user
