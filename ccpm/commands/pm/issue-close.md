---
description: Close an issue as complete
argument-hint: <issue_number> [completion_notes]
allowed-tools: Bash, Read, Write, Glob
---

<objective>
Mark issue #$ARGUMENTS as complete and close it on GitHub.
</objective>

<context>
Load workflow: @ccpm/skills/ccpm-issue/workflows/close-issue.md
Load reference: @ccpm/skills/shared-references/github-operations.md
Load reference: @ccpm/skills/shared-references/datetime.md
Issue: $ARGUMENTS
</context>

<process>
1. Load the close-issue workflow
2. Update local task status to closed
3. Close issue on GitHub with comment
4. Update epic progress calculation
5. Suggest next steps
</process>

<success_criteria>
- Local task marked complete
- Issue closed on GitHub
- Epic progress updated
</success_criteria>
