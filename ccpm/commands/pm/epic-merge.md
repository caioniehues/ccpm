---
description: Merge completed epic from worktree back to main
argument-hint: <epic_name>
allowed-tools: Bash, Read, Write, Glob
---

<objective>
Merge completed epic "$ARGUMENTS" from worktree back to main branch.
</objective>

<context>
Invoke the ccpm-epic skill with action: merge
Load skill: @ccpm/skills/ccpm-epic/SKILL.md
Load reference: @ccpm/skills/shared-references/worktree-operations.md
Load reference: @ccpm/skills/shared-references/github-operations.md
</context>

<process>
1. Load the ccpm-epic skill
2. Execute the "merge" action following skill guidance
3. Validate, merge, cleanup worktree, close GitHub issues
</process>

<success_criteria>
- Epic branch merged to main with --no-ff
- Worktree removed, branch deleted
- Epic archived to .claude/epics/archived/
- GitHub issues closed
</success_criteria>
