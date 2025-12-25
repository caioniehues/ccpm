---
description: Merge completed epic from worktree back to main
argument-hint: <epic_name>
allowed-tools: Bash, Read, Write, Glob
---

<objective>
Merge completed epic "$ARGUMENTS" from worktree back to main branch.
</objective>

<context>
Load workflow: @ccpm/skills/ccpm-epic/workflows/merge-epic.md
Load reference: @ccpm/skills/shared-references/worktree-operations.md
Load reference: @ccpm/skills/shared-references/github-operations.md
Epic: $ARGUMENTS
</context>

<process>
1. Load the merge-epic workflow
2. Validate worktree is clean and all work committed
3. Merge epic branch to main with --no-ff
4. Remove worktree and delete branch
5. Archive epic to .claude/epics/archived/
6. Close GitHub issues
</process>

<success_criteria>
- Epic branch merged to main with --no-ff
- Worktree removed, branch deleted
- Epic archived to .claude/epics/archived/
- GitHub issues closed
</success_criteria>
