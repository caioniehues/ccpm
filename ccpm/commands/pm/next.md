---
description: List next available tasks ready to be worked on
allowed-tools: Bash
---

<objective>
Display tasks that are ready to be worked on (no unmet dependencies, not started).
</objective>

<process>
!bash ccpm/scripts/pm/next.sh
</process>

<success_criteria>
Ready tasks listed in priority order with epic context.
</success_criteria>
