---
description: Create a new Product Requirements Document through brainstorming
argument-hint: <feature_name>
allowed-tools: Bash, Read, Write, Glob
---

<objective>
Launch brainstorming session to create a comprehensive PRD for feature: $ARGUMENTS
</objective>

<context>
Load workflow: @ccpm/skills/ccpm-prd/workflows/new-prd.md
Load reference: @ccpm/skills/shared-references/datetime.md
Feature name: $ARGUMENTS
</context>

<process>
1. Load the new-prd workflow
2. Execute preflight checks (validate name, check existing, create directory)
3. Conduct structured brainstorming session (4 phases)
4. Create PRD at .claude/prds/$ARGUMENTS.md with all sections
5. Validate and provide next steps
</process>

<success_criteria>
- PRD created with all required sections
- Valid frontmatter with real datetime
- No placeholder content
- Next steps provided: /pm:prd-parse $ARGUMENTS
</success_criteria>
