---
name: ccpm-prd
description: Create, parse, and edit Product Requirements Documents (PRDs). Supports brainstorming new features, converting PRDs to epics, and editing existing PRDs.
---

<objective>
Manage Product Requirements Documents that define features, user needs, and success criteria. PRDs serve as the foundation for technical implementation planning and are the starting point for the epic workflow.

This skill handles three core actions:
- **new**: Create a new PRD through structured brainstorming
- **parse**: Convert a PRD into a technical implementation epic
- **edit**: Modify an existing PRD
</objective>

<shared_references>
Load before any operation:
- @ccpm/skills/shared-references/datetime.md
- @ccpm/skills/shared-references/frontmatter-operations.md
</shared_references>

<action name="new">
<description>
Launch brainstorming session to create a comprehensive Product Requirements Document for a new feature.
</description>

<preflight>
1. **Validate feature name format:**
   - Must be kebab-case (lowercase letters, numbers, hyphens only)
   - Must start with a letter
   - If invalid: "❌ Feature name must be kebab-case. Examples: user-auth, payment-v2"

2. **Check for existing PRD:**
   - Check if `.claude/prds/$ARGUMENTS.md` already exists
   - If exists, ask: "⚠️ PRD '$ARGUMENTS' already exists. Overwrite? (yes/no)"
   - Only proceed with explicit 'yes'
   - If no, suggest: "/pm:prd-parse $ARGUMENTS to create epic from existing PRD"

3. **Verify directory structure:**
   - Create `.claude/prds/` if needed
   - Verify write permissions
</preflight>

<process>
1. **Discovery & Context**
   - Ask clarifying questions about the feature
   - Understand the problem being solved
   - Identify target users and use cases
   - Gather constraints and requirements

2. **PRD Structure**
   Create comprehensive PRD with sections:

   **Executive Summary** - Brief overview and value proposition

   **Problem Statement** - What problem are we solving? Why now?

   **User Stories** - Primary personas, user journeys, pain points

   **Requirements**
   - Functional Requirements: Core features and capabilities
   - Non-Functional Requirements: Performance, security, scalability

   **Success Criteria** - Measurable outcomes, key metrics

   **Constraints & Assumptions** - Technical, timeline, resource limitations

   **Out of Scope** - What we're explicitly NOT building

   **Dependencies** - External and internal dependencies

3. **File Format**
   Save to `.claude/prds/$ARGUMENTS.md`:
   ```markdown
   ---
   name: $ARGUMENTS
   description: [Brief one-line description]
   status: backlog
   created: [REAL datetime from date command]
   ---

   # PRD: $ARGUMENTS

   ## Executive Summary
   [Content...]

   ## Problem Statement
   [Content...]

   [Continue with all sections...]
   ```

4. **Quality Checks**
   - All sections complete (no placeholders)
   - User stories include acceptance criteria
   - Success criteria are measurable
   - Dependencies clearly identified
   - Out of scope items listed

5. **Post-Creation**
   ```
   ✅ PRD created: .claude/prds/$ARGUMENTS.md

   Summary:
     - [Brief summary of what was captured]

   Next: Ready to create implementation epic? Run: /pm:prd-parse $ARGUMENTS
   ```
</process>

<error_handling>
- If step fails, explain clearly what went wrong
- Never leave partial or corrupted files
- Provide specific guidance for resolution
</error_handling>
</action>

<action name="parse">
<description>
Convert a Product Requirements Document into a technical implementation epic.
</description>

<preflight>
1. **Verify feature_name provided:**
   - If not: "❌ <feature_name> not provided. Run: /pm:prd-parse <feature_name>"

2. **Verify PRD exists:**
   - Check `.claude/prds/$ARGUMENTS.md`
   - If not found: "❌ PRD not found: $ARGUMENTS. Create it with: /pm:prd-new $ARGUMENTS"

3. **Validate PRD frontmatter:**
   - Verify: name, description, status, created
   - If invalid: "❌ Invalid PRD frontmatter. Check: .claude/prds/$ARGUMENTS.md"

4. **Check for existing epic:**
   - Check if `.claude/epics/$ARGUMENTS/epic.md` exists
   - If exists, ask: "⚠️ Epic already exists. Overwrite? (yes/no)"
   - Only proceed with 'yes'

5. **Verify directory permissions:**
   - Ensure `.claude/epics/` can be created
</preflight>

<process>
1. **Read the PRD**
   - Load `.claude/prds/$ARGUMENTS.md`
   - Analyze all requirements and constraints
   - Understand user stories and success criteria
   - Extract description from frontmatter

2. **Technical Analysis**
   - Identify architectural decisions needed
   - Determine technology stack and approaches
   - Map functional requirements to technical components
   - Identify integration points and dependencies

3. **Create Epic File**
   Create `.claude/epics/$ARGUMENTS/epic.md`:
   ```markdown
   ---
   name: $ARGUMENTS
   status: backlog
   created: [REAL datetime]
   progress: 0%
   prd: .claude/prds/$ARGUMENTS.md
   github: [Will be updated when synced]
   ---

   # Epic: $ARGUMENTS

   ## Overview
   Brief technical summary of implementation approach

   ## Architecture Decisions
   - Key technical decisions and rationale
   - Technology choices
   - Design patterns to use

   ## Technical Approach
   ### Frontend Components
   - UI components needed
   - State management approach

   ### Backend Services
   - API endpoints required
   - Data models and schema

   ### Infrastructure
   - Deployment considerations
   - Scaling requirements

   ## Implementation Strategy
   - Development phases
   - Risk mitigation
   - Testing approach

   ## Task Breakdown Preview
   High-level task categories (limit to 10 or less):
   - [ ] Category 1: Description
   - [ ] Category 2: Description

   ## Dependencies
   - External service dependencies
   - Internal team dependencies

   ## Success Criteria (Technical)
   - Performance benchmarks
   - Quality gates

   ## Estimated Effort
   - Overall timeline estimate
   - Resource requirements
   ```

4. **Quality Validation**
   - All PRD requirements addressed
   - Task breakdown covers all areas
   - Dependencies technically accurate
   - Estimates realistic

5. **Post-Creation**
   ```
   ✅ Epic created: .claude/epics/$ARGUMENTS/epic.md

   Summary:
     - Task categories: {count}
     - Key decisions: {list}
     - Estimated effort: {estimate}

   Next: Ready to break down into tasks? Run: /pm:epic-decompose $ARGUMENTS
   ```
</process>

<important_notes>
- Aim for as few tasks as possible (10 or less)
- Look for ways to simplify and leverage existing functionality
- Never create epic with incomplete information
</important_notes>
</action>

<action name="edit">
<description>
Edit an existing Product Requirements Document.
</description>

<preflight>
1. **Verify PRD exists:**
   - Check `.claude/prds/$ARGUMENTS.md`
   - If not found: "❌ PRD not found: $ARGUMENTS"

2. **Read current PRD:**
   - Parse frontmatter
   - Read all sections
</preflight>

<process>
1. **Read Current PRD**
   - Load `.claude/prds/$ARGUMENTS.md`
   - Parse frontmatter and all sections

2. **Interactive Edit**
   Ask user what sections to edit:
   - Executive Summary
   - Problem Statement
   - User Stories
   - Requirements (Functional/Non-Functional)
   - Success Criteria
   - Constraints & Assumptions
   - Out of Scope
   - Dependencies

3. **Update PRD**
   Get current datetime: `date -u +"%Y-%m-%dT%H:%M:%SZ"`

   Update PRD file:
   - Preserve frontmatter except `updated` field
   - Apply user's edits to selected sections
   - Update `updated` field with current datetime

4. **Check Epic Impact**
   If PRD has associated epic:
   - Notify: "This PRD has epic: {epic_name}"
   - Ask: "Epic may need updating. Review epic? (yes/no)"
   - If yes: "Review with: /pm:epic-edit {epic_name}"

5. **Output**
   ```
   ✅ Updated PRD: $ARGUMENTS
     Sections edited: {list}

   {If has epic}: ⚠️ Epic may need review: {epic_name}

   Next: /pm:prd-parse $ARGUMENTS to update epic
   ```
</process>

<important_notes>
- Preserve original creation date
- Keep version history in frontmatter if needed
- Follow frontmatter-operations reference
</important_notes>
</action>

<success_criteria>
- **new**: Comprehensive PRD created through structured brainstorming
- **parse**: Technical epic created covering all PRD requirements
- **edit**: PRD updated with proper datetime and epic impact notification
- All operations use real datetime from system clock
- Clear next steps provided to user
</success_criteria>
