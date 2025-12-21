---
description: Show task status breakdown for a specific epic
argument-hint: <epic-name>
allowed-tools: Bash(ccpm/scripts/pm/epic-status.sh:*)
---

<objective>
Display task status breakdown for a specific epic showing completed, in-progress, and pending tasks.
</objective>

<process>
!bash ccpm/scripts/pm/epic-status.sh $ARGUMENTS
</process>

<success_criteria>
Status breakdown displayed with task counts per status category.
</success_criteria>
