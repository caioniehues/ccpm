---
description: Search across PRDs, epics, and tasks
argument-hint: <search_term>
allowed-tools: Bash
---

<objective>
Search for content across PRDs, epics, and tasks matching the provided search term.
</objective>

<process>
!bash ccpm/scripts/pm/search.sh $ARGUMENTS
</process>

<success_criteria>
Search results displayed with file paths and matching content.
</success_criteria>
