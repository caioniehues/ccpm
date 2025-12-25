---
description: Analyze issue for parallel work streams
argument-hint: <issue_number>
allowed-tools: Bash, Read, Write, Glob
---

<objective>
Analyze issue #$ARGUMENTS to identify parallel work streams for maximum efficiency.
</objective>

<context>
Load workflow: @ccpm/skills/ccpm-issue/workflows/analyze-issue.md
Load reference: @ccpm/skills/shared-references/agent-coordination.md
Issue: $ARGUMENTS
</context>

<process>
1. Load the analyze-issue workflow
2. Read issue details and requirements
3. Identify independent work streams
4. Document file scopes to avoid conflicts
5. Create analysis file with parallel streams
</process>

<success_criteria>
- Parallel work streams identified
- Analysis file created
- Coordination points documented
</success_criteria>
