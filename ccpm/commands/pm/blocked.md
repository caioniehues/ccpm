---
description: List all blocked tasks across epics
allowed-tools: Bash
---

<objective>
Display all tasks that are currently blocked by unmet dependencies across all epics.
</objective>

<process>
!bash ccpm/scripts/pm/blocked.sh
</process>

<success_criteria>
Blocked tasks listed with their blocking dependencies identified.
</success_criteria>
