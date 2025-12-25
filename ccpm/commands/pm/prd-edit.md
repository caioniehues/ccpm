---
description: Edit an existing Product Requirements Document
argument-hint: <feature_name>
allowed-tools: Read, Write, Glob
---

<objective>
Edit the existing PRD for $ARGUMENTS with user-specified changes.
</objective>

<context>
Load workflow: @ccpm/skills/ccpm-prd/workflows/edit-prd.md
Load reference: @ccpm/skills/shared-references/datetime.md
PRD location: .claude/prds/$ARGUMENTS.md
</context>

<process>
1. Load the edit-prd workflow
2. Verify PRD exists and read current content
3. Present edit options to user (section or full edit)
4. Apply changes and update timestamp
5. Check for associated epic impact
</process>

<success_criteria>
- PRD updated with real datetime
- Epic impact notification if applicable
- Original creation date preserved
- Changes validated and summarized
</success_criteria>
