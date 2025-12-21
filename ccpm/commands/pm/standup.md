---
description: Generate daily standup report with today's activity and recent changes
allowed-tools: Bash
---

<objective>
Generate a standup report summarizing recent activity, in-progress work, and upcoming tasks.
</objective>

<process>
!bash ccpm/scripts/pm/standup.sh
</process>

<success_criteria>
Standup report displayed with yesterday's completions, today's work, and blockers.
</success_criteria>
