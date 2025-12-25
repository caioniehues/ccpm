---
description: Push epic and tasks to GitHub as issues
argument-hint: <epic_name>
allowed-tools: Bash, Read, Write, Glob, Task
---

<objective>
Push epic "$ARGUMENTS" and tasks to GitHub as issues, rename files to issue numbers.
</objective>

<context>
Load workflow: @ccpm/skills/ccpm-epic/workflows/sync-epic.md
Load reference: @ccpm/skills/shared-references/github-operations.md
Load reference: @ccpm/skills/shared-references/worktree-operations.md
Epic: $ARGUMENTS
</context>

<process>
1. Load the sync-epic workflow
2. Verify epic and tasks exist locally
3. Create GitHub issues (epic as parent, tasks as sub-issues)
4. Rename local files from 001.md to {issue_id}.md
5. Create worktree for development
</process>

<success_criteria>
- Epic issue created on GitHub with proper labels
- Task sub-issues created and linked
- Files renamed from 001.md to {issue_id}.md
- Worktree created at ../epic-$ARGUMENTS
</success_criteria>
