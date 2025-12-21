# CCPM Commands Audit Summary

**48 critical, 11 warning, 2 info issues across 52 CCPM commands**

## Results Overview

| Metric | Count |
|--------|-------|
| Total Commands | 52 |
| Pass | 23 (44%) |
| Fail | 29 (56%) |
| Critical Issues | 48 |
| Warning Issues | 11 |
| Info Issues | 2 |

## Batch Results

| Batch | Commands | Pass | Fail | Critical |
|-------|----------|------|------|----------|
| 1 - Root | 9 | 5 | 4 | 9 |
| 2 - Context+Testing | 5 | 5 | 0 | 0 |
| 3 - Epic | 12 | 7 | 5 | 8 |
| 4 - Issue | 8 | 7 | 1 | 2 |
| 5 - PRD | 5 | 3 | 2 | 4 |
| 6 - General PM | 13 | 0 | 13 | 25 |

## Top Issues by Category

### 1. Missing XML Structure (38 critical)
Commands missing required `<objective>`, `<process>`, `<success_criteria>` tags.

**Two patterns:**
- **Bash Delegators (14 commands)**: Only invoke a bash script with no XML
  - epic-list, epic-show, epic-status, prd-list, prd-status
  - blocked, in-progress, init, next, search, standup, status, validate

- **Pure Markdown (8 commands)**: Have content but use # headings instead of XML
  - code-rabbit, context, prompt, re-init
  - epic-start-worktree, issue-show, clean, import, sync, test-reference-update

### 2. Missing Description Field (7 critical)
YAML frontmatter missing required `description` field:
- code-rabbit, prompt, re-init
- epic-start-worktree, issue-show
- clean, sync, test-reference-update

### 3. Markdown Headings in XML (5 warning)
Commands with proper XML structure but markdown ## headings inside tags:
- doctor, self-update, setup, uninstall

### 4. Missing argument-hint (5 warning)
Commands using `$ARGUMENTS` without declaring `argument-hint`:
- epic-start-worktree, issue-show, search, sync, testing/prime

## Pattern Issues Across Multiple Commands

| Pattern | Count | Affected Commands |
|---------|-------|-------------------|
| Bash delegator, no XML | 14 | epic-list, epic-show, epic-status, prd-list, prd-status, blocked, in-progress, init, next, search, standup, status, validate |
| Markdown headings in XML | 5 | doctor, self-update, setup, uninstall + help (partial) |
| Pure markdown structure | 8 | code-rabbit, epic-start-worktree, issue-show, clean, import, sync, test-reference-update |
| Missing argument-hint | 5 | epic-start-worktree, issue-show, search, sync, testing/prime |
| Missing description | 7 | code-rabbit, prompt, re-init, epic-start-worktree, issue-show, clean, sync, test-reference-update |

## Commands by Compliance Level

### Fully Compliant (5 commands)
- version
- context/create, context/prime, testing/run
- (empty XML structure but complete)

### Structurally Compliant with Warnings (18 commands)
- doctor, self-update, setup, uninstall (markdown in XML)
- context/update, testing/prime (minor argument issues)
- All skill-delegating commands: epic-close, epic-decompose, epic-edit, epic-merge, epic-oneshot, epic-refresh, epic-start, epic-sync
- issue-analyze, issue-close, issue-edit, issue-reopen, issue-start, issue-status, issue-sync
- prd-edit, prd-new, prd-parse

### Non-Compliant (29 commands)
- **Root**: code-rabbit, context, prompt, re-init
- **Epic**: epic-list, epic-show, epic-start-worktree, epic-status
- **Issue**: issue-show
- **PRD**: prd-list, prd-status
- **General PM**: blocked, clean, help, import, in-progress, init, next, search, standup, status, sync, test-reference-update, validate

## Remediation Priority

### Priority 1: Bash Delegators (14 commands)
Add minimal XML wrapper to bash-only commands:
```markdown
<objective>
Brief description of what the script shows/does.
</objective>

!bash ccpm/scripts/pm/script-name.sh $ARGUMENTS

<success_criteria>
What success looks like for the user.
</success_criteria>
```

### Priority 2: Missing Description + XML (7 commands)
Add frontmatter description AND full XML structure:
- code-rabbit, prompt, re-init
- epic-start-worktree, issue-show
- clean, sync, test-reference-update

### Priority 3: Markdown in XML (5 commands)
Remove ## headings from inside XML tags:
- doctor, self-update, setup, uninstall, help

### Priority 4: Missing argument-hint (5 commands)
Add argument-hint to frontmatter for commands using $ARGUMENTS:
- epic-start-worktree, issue-show, search, sync, testing/prime

## Comparison with Skills Audit

| Metric | Skills (001) | Commands (002) |
|--------|--------------|----------------|
| Components | 6 | 52 |
| Critical Issues | 26 | 48 |
| Warning Issues | 12 | 11 |
| Pass Rate | 0% | 44% |

Commands show better compliance than skills overall (44% pass vs 0%), but have more total critical issues due to the larger component count.

## Next Step

Run **003-ccpm-agents-audit** to audit the 10 CCPM subagent configurations.
