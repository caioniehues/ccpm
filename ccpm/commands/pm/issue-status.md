---
description: Check issue status and current state
argument-hint: <issue_number>
allowed-tools: Bash, Read, Glob
---

<objective>
Check the current status of issue #$ARGUMENTS and provide a quick status report.
</objective>

<context>
Load workflow: @ccpm/skills/ccpm-issue/workflows/status-issue.md
Issue: $ARGUMENTS
</context>

<process>
1. Load the status-issue workflow
2. Read local task file and frontmatter
3. Check GitHub issue state
4. Show epic context if applicable
5. Suggest next actions based on status
</process>

<success_criteria>
- Issue status displayed
- Epic context shown if applicable
- Relevant next actions suggested
</success_criteria>
