<objective>
Execute comprehensive audits of all 52 CCPM slash commands in parallel batches.

Purpose: Gather structured findings about command structure, YAML compliance, and integration issues
Output: commands-audit.md with XML-structured findings per command
</objective>

<context>
Project: CCPM (Claude Code Project Manager)
Location: ccpm/commands/

Previous audit: @.prompts/001-ccpm-skills-audit/skills-audit.md

Command categories:
- Root (9): code-rabbit, context, doctor, prompt, re-init, self-update, setup, uninstall, version
- Context+Testing (5): context/create, context/prime, context/update, testing/prime, testing/run
- Epic (12): pm/epic-* commands
- Issue (8): pm/issue-* commands
- PRD (5): pm/prd-* commands
- General PM (13): pm/blocked, clean, help, import, in-progress, init, next, search, standup, status, sync, test-reference-update, validate
</context>

<requirements>
Execute audits in 6 PARALLEL BATCHES (each batch runs commands in parallel):

BATCH 1 - Root (9 commands in parallel):
- ccpm/commands/code-rabbit.md
- ccpm/commands/context.md
- ccpm/commands/doctor.md
- ccpm/commands/prompt.md
- ccpm/commands/re-init.md
- ccpm/commands/self-update.md
- ccpm/commands/setup.md
- ccpm/commands/uninstall.md
- ccpm/commands/version.md

BATCH 2 - Context+Testing (5 commands in parallel):
- ccpm/commands/context/create.md
- ccpm/commands/context/prime.md
- ccpm/commands/context/update.md
- ccpm/commands/testing/prime.md
- ccpm/commands/testing/run.md

BATCH 3 - Epic (12 commands in parallel):
- ccpm/commands/pm/epic-close.md
- ccpm/commands/pm/epic-decompose.md
- ccpm/commands/pm/epic-edit.md
- ccpm/commands/pm/epic-list.md
- ccpm/commands/pm/epic-merge.md
- ccpm/commands/pm/epic-oneshot.md
- ccpm/commands/pm/epic-refresh.md
- ccpm/commands/pm/epic-show.md
- ccpm/commands/pm/epic-start.md
- ccpm/commands/pm/epic-start-worktree.md
- ccpm/commands/pm/epic-status.md
- ccpm/commands/pm/epic-sync.md

BATCH 4 - Issue (8 commands in parallel):
- ccpm/commands/pm/issue-analyze.md
- ccpm/commands/pm/issue-close.md
- ccpm/commands/pm/issue-edit.md
- ccpm/commands/pm/issue-reopen.md
- ccpm/commands/pm/issue-show.md
- ccpm/commands/pm/issue-start.md
- ccpm/commands/pm/issue-status.md
- ccpm/commands/pm/issue-sync.md

BATCH 5 - PRD (5 commands in parallel):
- ccpm/commands/pm/prd-edit.md
- ccpm/commands/pm/prd-list.md
- ccpm/commands/pm/prd-new.md
- ccpm/commands/pm/prd-parse.md
- ccpm/commands/pm/prd-status.md

BATCH 6 - General PM (13 commands in parallel):
- ccpm/commands/pm/blocked.md
- ccpm/commands/pm/clean.md
- ccpm/commands/pm/help.md
- ccpm/commands/pm/import.md
- ccpm/commands/pm/in-progress.md
- ccpm/commands/pm/init.md
- ccpm/commands/pm/next.md
- ccpm/commands/pm/search.md
- ccpm/commands/pm/standup.md
- ccpm/commands/pm/status.md
- ccpm/commands/pm/sync.md
- ccpm/commands/pm/test-reference-update.md
- ccpm/commands/pm/validate.md

For each command, use: /taches:audit-slash-command {path}

Execute each batch with ALL commands in parallel (single message, multiple Task tool calls).
Wait for batch to complete before starting next batch.
</requirements>

<output_structure>
Save to: .prompts/002-ccpm-commands-audit/commands-audit.md

Write findings INCREMENTALLY:
1. Create file with XML skeleton
2. Append each batch's findings as it completes
3. Update running totals after each batch
4. Finalize summary at end

XML structure:

```xml
<audit_results type="commands">
  <summary>
    <total_components>52</total_components>
    <critical_issues>N</critical_issues>
    <warning_issues>N</warning_issues>
    <info_issues>N</info_issues>
  </summary>

  <batch number="1" name="root">
    <component name="{command-name}">
      <path>ccpm/commands/{path}</path>
      <status>pass|fail</status>
      <issues>
        <issue severity="critical|warning|info">
          <location>file:line</location>
          <description>What's wrong</description>
          <fix>How to fix it</fix>
        </issue>
      </issues>
    </component>
  </batch>

  <batch number="2" name="context-testing">...</batch>
  <batch number="3" name="epic">...</batch>
  <batch number="4" name="issue">...</batch>
  <batch number="5" name="prd">...</batch>
  <batch number="6" name="general-pm">...</batch>

  <metadata>
    <confidence level="high">Based on automated audit tools</confidence>
  </metadata>
</audit_results>
```
</output_structure>

<summary_requirements>
Create .prompts/002-ccpm-commands-audit/SUMMARY.md

One-liner: "{N} critical, {N} warning, {N} info issues across 52 CCPM commands"

Key Findings:
- Top issues by category
- Any commands that failed audit entirely
- Pattern issues (appearing in multiple commands)

Next Step: Run 003-ccpm-agents-audit
</summary_requirements>

<success_criteria>
- All 52 commands audited
- Structured XML output with batch grouping
- SUMMARY.md with category breakdown
- Ready for downstream consumption
</success_criteria>
