---
description: Analyze issue for parallel work streams
argument-hint: <issue_number>
allowed-tools: Bash, Read, Write, Glob
---

<objective>
Analyze issue #$ARGUMENTS to identify parallel work streams for maximum efficiency.
</objective>

<context>
Invoke the ccpm-issue skill with action: analyze
Load skill: @ccpm/skills/ccpm-issue/SKILL.md
</context>

<process>
1. Load the ccpm-issue skill
2. Execute the "analyze" action following skill guidance
3. Create analysis file with parallel streams identified
</process>

<success_criteria>
- Parallel work streams identified
- Analysis file created
- Coordination points documented
</success_criteria>
