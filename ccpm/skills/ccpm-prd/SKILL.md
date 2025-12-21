---
name: ccpm-prd
description: Manages Product Requirements Documents (PRDs) for CCPM including creation, editing, parsing to epics, and status tracking. Use when working with PRDs or when user mentions product requirements, feature documentation, or epic generation.
---

<objective>
Coordinate Product Requirements Document (PRD) operations including creation through structured brainstorming, editing, parsing to technical epics, and portfolio tracking. Provides a unified interface for managing the product requirements lifecycle from initial documentation through technical implementation planning.
</objective>

<essential_principles>
## How PRD Management Works

PRDs (Product Requirements Documents) are the foundation of feature development in CCPM. They capture product requirements, user stories, and success criteria before technical implementation.

### PRD Lifecycle

1. **Creation**: Brainstorm and document product requirements through structured discovery
2. **Editing**: Refine and update PRD sections as requirements evolve
3. **Parsing**: Convert PRD to technical implementation epic
4. **Tracking**: Monitor PRD status and progress across the portfolio

### File Structure

PRDs are stored in `.claude/prds/{feature-name}.md` with frontmatter:
- **name**: Feature name (kebab-case)
- **description**: Brief one-line summary
- **status**: backlog, in-progress, completed, or on-hold
- **created**: ISO 8601 datetime
- **updated**: ISO 8601 datetime (when edited)

### PRD Sections

Complete PRDs include:
- Executive Summary
- Problem Statement
- User Stories with acceptance criteria
- Functional and Non-Functional Requirements
- Success Criteria (measurable)
- Constraints & Assumptions
- Out of Scope (explicit exclusions)
- Dependencies

### Relationship to Epics

PRDs describe WHAT to build and WHY. Epics describe HOW to build it technically. Use `/pm:prd-parse` to generate an epic from a PRD.
</essential_principles>

<intake>
What would you like to do with PRDs?

1. Create a new PRD (brainstorming session)
2. Edit an existing PRD
3. Parse PRD to epic (convert product requirements to technical implementation)
4. List all PRDs
5. Show PRD status report

**Wait for response before proceeding.**
</intake>

<routing>
| Response | Command | Description |
|----------|---------|-------------|
| 1, "new", "create" | `/pm:prd-new` | Launch brainstorming for new PRD |
| 2, "edit", "modify", "update" | `/pm:prd-edit` | Edit existing PRD sections |
| 3, "parse", "epic", "convert" | `/pm:prd-parse` | Convert PRD to implementation epic |
| 4, "list", "show all", "view" | `/pm:prd-list` | List all PRDs with status |
| 5, "status", "report", "stats" | `/pm:prd-status` | Show PRD status counts |

**After determining the operation, delegate to the appropriate command.**
</routing>

<operation_details>
## Available Operations

**Create New PRD** (`/pm:prd-new <feature_name>`)
- Validates feature name format (kebab-case)
- Checks for existing PRD
- Conducts structured brainstorming session
- Creates comprehensive PRD with all required sections
- Saves to `.claude/prds/{feature-name}.md`

**Edit PRD** (`/pm:prd-edit <feature_name>`)
- Reads current PRD
- Allows selective section editing
- Updates timestamp
- Checks for associated epic impact

**Parse to Epic** (`/pm:prd-parse <feature_name>`)
- Reads PRD requirements
- Performs technical analysis
- Creates implementation epic at `.claude/epics/{feature-name}/epic.md`
- Maps product requirements to technical approach
- Suggests task breakdown

**List PRDs** (`/pm:prd-list`)
- Shows all PRDs with status
- Displays created/updated dates
- Organized by status

**Status Report** (`/pm:prd-status`)
- Shows count by status (backlog, in-progress, completed, on-hold)
- Portfolio overview
</operation_details>

<quick_start>
**Common workflow:**

```bash
# 1. Create a new PRD
/pm:prd-new user-authentication

# 2. Parse to epic when ready
/pm:prd-parse user-authentication

# 3. View all PRDs
/pm:prd-list

# 4. Edit if needed
/pm:prd-edit user-authentication
```
</quick_start>

<success_criteria>
PRD operations are successful when:

- **New PRD**: File created with valid frontmatter, all sections complete, no placeholders
- **Edit PRD**: Specified sections updated, timestamp refreshed, frontmatter preserved
- **Parse to Epic**: Epic created with technical approach, task breakdown, and PRD reference
- **List/Status**: Current PRD state accurately displayed

All operations should:
- Use real timestamps (never placeholders)
- Validate inputs before processing
- Provide clear error messages
- Suggest next steps after completion
</success_criteria>
