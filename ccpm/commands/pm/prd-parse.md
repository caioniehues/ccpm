---
description: Convert PRD to technical implementation epic
argument-hint: <feature_name>
allowed-tools: Bash, Read, Write, Glob
---

<objective>
Convert the PRD for $ARGUMENTS into a technical implementation epic with architecture decisions and task breakdown.
</objective>

<context>
Invoke the ccpm-prd skill with action: parse
Load skill: @ccpm/skills/ccpm-prd/SKILL.md
PRD source: .claude/prds/$ARGUMENTS.md
</context>

<process>
1. Load the ccpm-prd skill
2. Execute the "parse" action following skill guidance
3. Create epic at .claude/epics/$ARGUMENTS/epic.md
</process>

<success_criteria>
- Epic created with technical approach
- Task breakdown preview (10 or fewer categories)
- Architecture decisions documented
</success_criteria>
