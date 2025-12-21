---
description: Create a new Product Requirements Document through brainstorming
argument-hint: <feature_name>
allowed-tools: Bash, Read, Write, Glob
---

<objective>
Launch brainstorming session to create a comprehensive PRD for feature: $ARGUMENTS
</objective>

<context>
Invoke the ccpm-prd skill with action: new
Load skill: @ccpm/skills/ccpm-prd/SKILL.md
Feature name: $ARGUMENTS
</context>

<process>
1. Load the ccpm-prd skill
2. Execute the "new" action following skill guidance
3. Conduct brainstorming and create PRD at .claude/prds/$ARGUMENTS.md
</process>

<success_criteria>
- PRD created with all required sections
- Valid frontmatter with real datetime
- Next steps provided to user
</success_criteria>
