<objective>
Execute LOW priority fixes from CCPM remediation plan.

Purpose: Cleanup and consolidation (remove duplicates, merge reference files, standardize patterns)
Input: @.prompts/005-ccpm-remediation-plan/remediation-plan.md
Output: Modified files + SUMMARY.md
</objective>

<context>
Remediation plan: @.prompts/005-ccpm-remediation-plan/remediation-plan.md
Phase: 4 - low-priority
Total issues: ~55
</context>

<fixes>
## Duplicate Workflow Files (7 pairs)

Delete short-name files and update SKILL.md routing to point to -issue.md variants.

Fix 1: Delete ccpm/skills/ccpm-issue/workflows/analyze.md
- Keep: analyze-issue.md (116 lines, more complete)
- Update: SKILL.md routing table

Fix 2: Delete ccpm/skills/ccpm-issue/workflows/start.md
- Keep: start-issue.md
- Update: SKILL.md routing table

Fix 3: Delete ccpm/skills/ccpm-issue/workflows/close.md
- Keep: close-issue.md
- Update: SKILL.md routing table

Fix 4: Delete ccpm/skills/ccpm-issue/workflows/edit.md
- Keep: edit-issue.md
- Update: SKILL.md routing table

Fix 5: Delete ccpm/skills/ccpm-issue/workflows/reopen.md
- Keep: reopen-issue.md
- Update: SKILL.md routing table

Fix 6: Delete ccpm/skills/ccpm-issue/workflows/sync.md
- Keep: sync-issue.md
- Update: SKILL.md routing table

Fix 7: Delete ccpm/skills/ccpm-issue/workflows/status.md
- Keep: status-issue.md
- Update: SKILL.md routing table

## Duplicate Reference Files (5 files)

Consolidate to shared-references/ (per user decision).

Fix 8: Merge datetime.md
- Source: ccpm/rules/datetime.md (has 73% more content)
- Target: ccpm/skills/shared-references/datetime.md
- Action: Merge extra content from rules/ into shared-references/, delete rules/datetime.md

Fix 9: Move agent-coordination.md
- Source: ccpm/rules/agent-coordination.md
- Target: ccpm/skills/shared-references/agent-coordination.md
- Action: Create shared-references version, update all @-references, delete rules version

Fix 10: Merge frontmatter-operations.md
- Source: ccpm/rules/frontmatter-operations.md
- Target: ccpm/skills/shared-references/frontmatter-operations.md (has validation_rules)
- Action: Merge content, keep shared-references version, delete rules version

Fix 11: Merge github-operations.md
- Source: ccpm/rules/github-operations.md
- Target: ccpm/skills/shared-references/github-operations.md
- Action: Merge both versions (different fix instructions), delete rules version

Fix 12: Delete duplicate worktree-operations.md
- Source: ccpm/rules/worktree-operations.md
- Target: ccpm/skills/shared-references/worktree-operations.md
- Action: Keep shared-references version (near-identical), delete rules version

## Standardize Status Values

Fix 13: Standardize status vocabulary across files
- Files: parallel-orchestrator.md, parallel-worker.md, agent-coordination.md, epic-lifecycle.md
- Current: running/success/failed, waiting/running/completed/failed, open/in_progress/closed
- Change to: pending | running | completed | failed

## Workflow File Structure (Context)

Fix 14-16: Fix markdown headings in context workflows
- ccpm/skills/ccpm-context/workflows/create-context.md (lines 4-110)
- ccpm/skills/ccpm-context/workflows/prime-context.md (lines 4-129)
- ccpm/skills/ccpm-context/workflows/update-context.md (lines 4-189)
- Change: Replace `## Step N` with XML tags: `<step_preflight>`, `<step_analysis>`, etc.

## Missing Workflow Files

Fix 17: Address ccpm-prd empty workflows
- File: ccpm/skills/ccpm-prd/workflows/.gitkeep
- Issue: Routing table references 5 commands but no workflow files
- Action: Either create workflow files OR update routing to use commands directly

## Reference File Structure

Fix 18: Fix frontmatter-operations.md
- File: ccpm/skills/ccpm-issue/references/frontmatter-operations.md
- Lines: 8, 35, 53, 76, 113, 129, 154
- Remove: Markdown headings from inside XML tags

Fix 19: Fix datetime-handling.md
- File: ccpm/skills/ccpm-issue/references/datetime-handling.md
- Lines: 13, 37, 62, 103, 129, 156
- Remove: Markdown headings from inside XML tags

Fix 20: Restructure parallel-streams.md
- File: ccpm/skills/ccpm-issue/references/parallel-streams.md
- Lines: 1-190
- Change: Wrap in semantic XML tags: `<overview>`, `<stream_patterns>`, `<parallelization_strategies>`

## Epic File Structure

Fix 21: Restructure epic-lifecycle.md
- File: ccpm/skills/ccpm-epic/references/epic-lifecycle.md
- Lines: 1-275
- Change: Migrate from markdown to pure XML structure

Fix 22: Restructure epic workflows
- Files: ccpm/skills/ccpm-epic/workflows/*.md (8 files)
- Change: Migrate from markdown headings to XML structure

## Issue Workflow Structure

Fix 23: Fix analyze-issue.md
- File: ccpm/skills/ccpm-issue/workflows/analyze-issue.md
- Lines: 11-12
- Change: Convert `## Step 1` pattern to XML step tags

## Info-Level Enhancements

Fix 24-27: Consider explicit model selection
- ccpm/agents/code-analyzer.md (line 5)
- ccpm/agents/epic-planner.md (line 31)
- ccpm/agents/parallel-orchestrator.md (line 4)
- ccpm/agents/worktree-manager.md (line 5)
- Current: model: inherit
- Consider: model: sonnet (for consistency)

Fix 28-29: Add error handling sections
- ccpm/agents/prd-architect.md (lines 196-205)
- ccpm/agents/task-decomposer.md (lines 103-110)
- Add: `<error_handling>` section for edge cases

Fix 30: Add success criteria to prd-architect
- File: ccpm/agents/prd-architect.md
- Lines: 35-37
- Add: `<success_criteria>` section

Fix 31: Fix task-decomposer description
- File: ccpm/agents/task-decomposer.md
- Line: 3
- Change: Use YAML multiline syntax for long description

Fix 32-33: Fix $ARGUMENTS placement
- ccpm/commands/context/update.md (line 28)
- ccpm/commands/testing/prime.md (line 27)
- Change: Move $ARGUMENTS into context or process tag

Fix 34: Verify worktree-manager reference
- File: ccpm/skills/ccpm-worktree/SKILL.md
- Lines: 224-235
- Verify: worktree-manager.md exists at ccpm/agents/

Fix 35: Verify show.md workflow
- File: ccpm/skills/ccpm-issue/workflows/show.md
- Verify: File is complete and properly structured
</fixes>

<verification>
After applying all fixes:
1. Verify no short-name workflow duplicates exist in ccpm-issue/workflows/
2. Verify SKILL.md routing points to -issue.md variants
3. Verify ccpm/rules/ only contains non-duplicate files
4. Verify all @-references to moved files are updated
5. Search for inconsistent status values - should find only: pending, running, completed, failed
6. Verify no markdown headings inside XML in reference files
</verification>

<summary_requirements>
Create .prompts/009-ccpm-remediation-low/SUMMARY.md

One-liner: "~55 cleanup issues fixed: 7 duplicates deleted, 5 references consolidated"

Key Findings:
- Duplicate files removed
- Reference files consolidated to shared-references/
- Status vocabulary standardized

Next Step: Remediation complete - run validation audits
</summary_requirements>

<success_criteria>
- All duplicate workflow files deleted
- All duplicate reference files consolidated to shared-references/
- SKILL.md routing tables updated
- Status vocabulary consistent across files
- All markdown headings removed from XML in reference/workflow files
- SUMMARY.md created
</success_criteria>
