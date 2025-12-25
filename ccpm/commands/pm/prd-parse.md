---
description: Convert PRD to technical implementation epic
argument-hint: <feature_name>
allowed-tools: Bash, Read, Write, Glob, Task
---

<objective>
Convert the PRD for $ARGUMENTS into a technical implementation epic with architecture decisions and task breakdown.
</objective>

<context>
Load workflow: @ccpm/skills/ccpm-prd/workflows/parse-prd.md
Load reference: @ccpm/skills/shared-references/datetime.md
PRD source: .claude/prds/$ARGUMENTS.md
Epic destination: .claude/epics/$ARGUMENTS/epic.md
</context>

<process>
1. Load the parse-prd workflow
2. Verify PRD exists and read content
3. Perform technical analysis (architecture, components, dependencies)
4. Create epic with task breakdown preview
5. Update PRD status to in-progress
</process>

<success_criteria>
- Epic created with technical approach
- Task breakdown preview (≤10 categories)
- Architecture decisions documented
- PRD status updated
- Next steps: /pm:epic-decompose $ARGUMENTS
</success_criteria>
