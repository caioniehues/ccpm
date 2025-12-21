<objective>
Execute HIGH priority fixes from CCPM remediation plan.

Purpose: Fix structure violations (markdown headings in XML, missing required tags, missing XML structure)
Input: @.prompts/005-ccpm-remediation-plan/remediation-plan.md
Output: Modified files + SUMMARY.md
</objective>

<context>
Remediation plan: @.prompts/005-ccpm-remediation-plan/remediation-plan.md
Phase: 2 - high-priority
Total issues: ~80 (grouped by pattern)
</context>

<fixes>
## Pattern A: Skills Missing Required Tags (6 files)

Each skill needs `<objective>` and `<quick_start>` tags added at the beginning of the file body.

Fix A1: ccpm/skills/ccpm-context/SKILL.md
- Add objective: "Manage conversation context files for CCPM workflows"
- Add quick_start: "/pm:context prime - Initialize context for current project"

Fix A2: ccpm/skills/ccpm-prd/SKILL.md
- Add objective: "Manage Product Requirements Documents (PRDs) lifecycle"
- Add quick_start: "/pm:prd-new {name} - Create a new PRD"

Fix A3: ccpm/skills/ccpm-testing/SKILL.md
- Add objective: "Manage CCPM test environment and test execution"
- Add quick_start: "/pm:testing prime - Set up test environment"

Fix A4: ccpm/skills/ccpm-worktree/SKILL.md
- Add objective: "Manage git worktrees for parallel epic development"
- Add quick_start: "/pm:epic-start-worktree {epic} - Launch parallel work on epic"

Fix A5: ccpm/skills/ccpm-epic/SKILL.md
- Add objective: "Manage epic lifecycle from planning through completion"
- Add quick_start: "/pm:prd-parse {name} - Parse PRD into epic, then /pm:epic-oneshot {name}"

Fix A6: ccpm/skills/ccpm-issue/SKILL.md
- Add objective: "Manage issue lifecycle with local-first sync to GitHub"
- Add quick_start: "/pm:issue-start {number} - Start working on an issue"

## Pattern B: Markdown Headings in XML (Skills)

Remove ## and ### headings from inside XML tags. Replace with semantic XML subtags or plain prose.

Fix B1: ccpm/skills/ccpm-context/SKILL.md (lines 6-49)
- Remove: `## How Context Management Works`, `### Context Structure`, etc.
- Replace with: XML subtags `<overview>`, `<context_structure>`, `<frontmatter_requirements>`

Fix B2: ccpm/skills/ccpm-prd/SKILL.md (lines 11-43)
- Remove: `## How PRD Management Works`, `### PRD Lifecycle`, etc.
- Replace with: XML subtags `<prd_lifecycle>`, `<file_structure>`, `<prd_sections>`

Fix B3: ccpm/skills/ccpm-testing/SKILL.md (lines 7,11,15,19,51,79)
- Remove all ## and ### headings from XML tags
- Use nested XML elements or plain text with line breaks

Fix B4: ccpm/skills/ccpm-worktree/SKILL.md (lines 7,11,24,33)
- Remove markdown headings
- Replace with XML subtags or **bold text**

Fix B5: ccpm/skills/ccpm-epic/SKILL.md (lines 6-45, 91-116, 118-156)
- Remove all markdown headings from essential_principles, command_reference, workflow_patterns
- Replace with nested semantic XML tags

Fix B6: ccpm/skills/ccpm-issue/SKILL.md (lines 6-23, 55-69, 71-84)
- Remove markdown headings from essential_principles, reference_index, workflows_index
- XML tag names are self-documenting

## Pattern C: Markdown Headings in XML (Agents)

Fix C1: ccpm/agents/code-analyzer.md (line 11)
- Convert 11 markdown headings to semantic XML tags
- Wrap body in `<role>`, `<constraints>`, `<workflow>` tags

Fix C2: ccpm/agents/file-analyzer.md (lines 11-53)
- Replace 6 markdown headings with semantic XML tags
- Wrap role in `<role>` tags

Fix C3: ccpm/agents/parallel-worker.md (lines 11-247)
- Convert 16+ markdown headings to pure XML
- Add `<role>`, `<constraints>`, `<workflow>` structure

Fix C4: ccpm/agents/test-runner.md (lines 11-206)
- Convert 13 markdown headings to semantic XML tags
- Wrap in proper XML structure

## Pattern D: Commands Missing XML Structure (25+ files)

For each command file, add the three required XML tags:
- `<objective>` - What the command does
- `<process>` - How it works (wrap existing content)
- `<success_criteria>` - What success looks like

Files needing full restructure:
- ccpm/commands/code-rabbit.md (also add description to YAML)
- ccpm/commands/prompt.md (also add description)
- ccpm/commands/re-init.md (also add description)
- ccpm/commands/context.md

Bash delegator files (wrap bash invocation in XML):
- ccpm/commands/pm/epic-list.md
- ccpm/commands/pm/epic-show.md
- ccpm/commands/pm/epic-status.md
- ccpm/commands/pm/prd-list.md
- ccpm/commands/pm/prd-status.md
- ccpm/commands/pm/blocked.md
- ccpm/commands/pm/in-progress.md
- ccpm/commands/pm/init.md
- ccpm/commands/pm/next.md
- ccpm/commands/pm/search.md (also add argument-hint)
- ccpm/commands/pm/standup.md
- ccpm/commands/pm/status.md
- ccpm/commands/pm/validate.md

Large files needing restructure:
- ccpm/commands/pm/epic-start-worktree.md (also add description, argument-hint)
- ccpm/commands/pm/issue-show.md (also add description, argument-hint)
- ccpm/commands/pm/clean.md (also add description)
- ccpm/commands/pm/import.md
- ccpm/commands/pm/sync.md (also add description, argument-hint)
- ccpm/commands/pm/test-reference-update.md (also add description)

Missing process tag only:
- ccpm/commands/pm/help.md

## Pattern E: Markdown Headings in Command XML (4 files)

Remove ## and ### headings from inside process tags:
- ccpm/commands/doctor.md (lines 10-51)
- ccpm/commands/self-update.md (lines 10-42)
- ccpm/commands/setup.md (lines 10-56)
- ccpm/commands/uninstall.md (lines 10-57)

Replace with numbered lists or plain text with **bold** emphasis.
</fixes>

<verification>
After applying all fixes:
1. Grep for `## ` inside XML tags in skills/ - should find none
2. Grep for `### ` inside XML tags in skills/ - should find none
3. Verify all 6 skills have objective and quick_start tags
4. Verify all commands have objective, process, success_criteria tags
5. Check commands for missing YAML fields (description, argument-hint)
</verification>

<summary_requirements>
Create .prompts/007-ccpm-remediation-high/SUMMARY.md

One-liner: "~80 structure violations fixed in ~45 files"

Key Findings:
- List file categories modified
- Note common patterns applied

Next Step: Run 008-ccpm-remediation-medium
</summary_requirements>

<success_criteria>
- All markdown headings removed from XML tags
- All skills have objective and quick_start tags
- All commands have proper XML structure
- All commands have required YAML fields
- SUMMARY.md created
</success_criteria>
