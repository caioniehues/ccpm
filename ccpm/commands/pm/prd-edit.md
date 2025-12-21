---
description: Edit an existing Product Requirements Document
argument-hint: <feature_name>
allowed-tools: Read, Write, Glob
---

<objective>
Edit the existing PRD for $ARGUMENTS with user-specified changes.
</objective>

<context>
Invoke the ccpm-prd skill with action: edit
Load skill: @ccpm/skills/ccpm-prd/SKILL.md
PRD location: .claude/prds/$ARGUMENTS.md
</context>

<process>
1. Load the ccpm-prd skill
2. Execute the "edit" action following skill guidance
3. Apply user changes and check epic impact
</process>

<success_criteria>
- PRD updated with real datetime
- Epic impact notification if applicable
- Original creation date preserved
</success_criteria>
