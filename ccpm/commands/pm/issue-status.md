---
description: Check issue status and current state
argument-hint: <issue_number>
allowed-tools: Bash, Read, Glob
---

<objective>
Check the current status of issue #$ARGUMENTS and provide a quick status report.
</objective>

<context>
Invoke the ccpm-issue skill with action: status
Load skill: @ccpm/skills/ccpm-issue/SKILL.md
</context>

<process>
1. Load the ccpm-issue skill
2. Execute the "status" action following skill guidance
3. Display status with actionable suggestions
</process>

<success_criteria>
- Issue status displayed
- Epic context shown if applicable
- Relevant next actions suggested
</success_criteria>
