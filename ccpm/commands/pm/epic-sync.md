---
description: Push epic and tasks to GitHub as issues
argument-hint: <epic_name>
allowed-tools: Bash, Read, Write, Glob, Task
---

<objective>
Push epic "$ARGUMENTS" and tasks to GitHub as issues, rename files to issue numbers.
</objective>

<context>
Invoke the ccpm-epic skill with action: sync
Load skill: @ccpm/skills/ccpm-epic/SKILL.md
Load reference: @ccpm/skills/shared-references/github-operations.md
Load reference: @ccpm/skills/shared-references/worktree-operations.md
</context>

<process>
1. Load the ccpm-epic skill
2. Execute the "sync" action following skill guidance
3. Create GitHub issues, rename files, create worktree
</process>

<success_criteria>
- Epic issue created on GitHub with proper labels
- Task sub-issues created and linked
- Files renamed from 001.md to {issue_id}.md
- Worktree created at ../epic-$ARGUMENTS
</success_criteria>
