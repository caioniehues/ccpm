---
description: Mark epic as complete when all tasks done
argument-hint: <epic_name>
allowed-tools: Bash, Read, Write, Glob
---

<objective>
Mark epic "$ARGUMENTS" as complete after verifying all tasks are done.
</objective>

<context>
Load workflow: @ccpm/skills/ccpm-epic/workflows/close-epic.md
Load reference: @ccpm/skills/shared-references/github-operations.md
Epic: $ARGUMENTS
</context>

<process>
1. Load the close-epic workflow
2. Verify all tasks have status: closed
3. Update epic status to completed with 100% progress
4. Close GitHub epic issue
5. Offer archive option
</process>

<success_criteria>
- All tasks verified as closed
- Epic status set to completed with 100% progress
- GitHub epic issue closed
- Archive option offered
</success_criteria>
