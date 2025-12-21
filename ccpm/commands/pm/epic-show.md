---
description: Show detailed information about a specific epic
argument-hint: <epic-name>
allowed-tools: Bash(ccpm/scripts/pm/epic-show.sh:*)
---

<objective>
Display detailed information about a specific epic including tasks, progress, and metadata.
</objective>

<process>
!bash ccpm/scripts/pm/epic-show.sh $ARGUMENTS
</process>

<success_criteria>
Epic details displayed including task list and progress summary.
</success_criteria>
