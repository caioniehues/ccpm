<objective>
Execute MEDIUM priority fixes from CCPM remediation plan.

Purpose: Fix over-permissioning (remove unnecessary tools from leaf-node agents, add missing constraints)
Input: @.prompts/005-ccpm-remediation-plan/remediation-plan.md
Output: Modified files + SUMMARY.md
</objective>

<context>
Remediation plan: @.prompts/005-ccpm-remediation-plan/remediation-plan.md
Phase: 3 - medium-priority
Total issues: 15
</context>

<fixes>
## Tool Permission Fixes

Fix 1: Trim code-analyzer tools
- File: ccpm/agents/code-analyzer.md
- Line: 4
- Current: Glob, Grep, LS, Read, WebFetch, TodoWrite, WebSearch, Search, Task, Agent
- Change to: Glob, Grep, Read
- Reason: Leaf node should not spawn agents; code analysis is local-only

Fix 2: Trim epic-planner tools
- File: ccpm/agents/epic-planner.md
- Line: 30
- Current: includes Task, Agent
- Change to: Remove Task, Agent (keep Glob, Grep, Read, Write, TodoWrite)
- Reason: Planning agent produces plans, does not execute them

Fix 3: Trim file-analyzer tools
- File: ccpm/agents/file-analyzer.md
- Line: 4
- Current: Glob, Grep, LS, Read, WebFetch, TodoWrite, WebSearch, Search, Task, Agent
- Change to: Glob, Grep, LS, Read
- Reason: Leaf node for file analysis only

Fix 4: Trim test-runner tools
- File: ccpm/agents/test-runner.md
- Line: 4
- Current: Glob, Grep, LS, Read, WebFetch, TodoWrite, WebSearch, Search, Task, Agent (plus Bash from critical fixes)
- Change to: Bash, Glob, Grep, Read, TodoWrite
- Reason: Test execution needs Bash but not web tools or agent spawning

Fix 5: Trim parallel-worker tools
- File: ccpm/agents/parallel-worker.md
- Line: 4
- Current: (after critical fixes) has LS, WebFetch, TodoWrite, WebSearch
- Change to: Keep Glob, Grep, Read, Bash, Task (if orchestrator) OR remove Task (if leaf)
- Decision: Remove Agent tool, keep Task for orchestration; remove WebFetch, WebSearch, LS

Fix 6: Trim worktree-manager tools
- File: ccpm/agents/worktree-manager.md
- Line: 4
- Current: Bash, Read, Write, Glob
- Change to: Bash, Read, Glob
- Reason: Worktree operations are git commands; Write not needed

## Constraints Section Fixes

Fix 7: Add constraints to parallel-worker
- File: ccpm/agents/parallel-worker.md
- Line: After role section
- Add new section:
```xml
<constraints>
- MUST check for file conflicts before writing
- MUST use atomic file operations
- NEVER modify files being edited by another agent
- ALWAYS report status to execution-status.md
- MUST wait for dependencies before starting task
</constraints>
```

Fix 8: Add constraints to test-runner
- File: ccpm/agents/test-runner.md
- Line: After role section
- Add new section:
```xml
<constraints>
- MUST run tests in isolated environment when possible
- NEVER modify source code during test execution
- ALWAYS capture and report test output
- MUST report failures with actionable information
- NEVER run destructive commands (rm -rf, drop database, etc.)
</constraints>
```

Fix 9: Convert code-analyzer Operating Principles to constraints
- File: ccpm/agents/code-analyzer.md
- Line: 119-126
- Current: Operating Principles section without modal verbs
- Change to: Proper constraints section with MUST/NEVER/ALWAYS

Fix 10: Convert file-analyzer Important Guidelines to constraints
- File: ccpm/agents/file-analyzer.md
- Line: 140-145
- Current: Important Guidelines section
- Change to: Proper constraints section with MUST/NEVER/ALWAYS

## Miscellaneous Fixes

Fix 11: Verify github-syncer sub-agent reference
- File: ccpm/agents/github-syncer.md
- Line: 140
- Current: References "Parallel creation via sub-agents"
- Verify: After adding Task tool (critical fix), verify this reference works
- Fix if needed: Update reference to use parallel-worker agent

Fix 12: Add trigger keywords to parallel-worker description
- File: ccpm/agents/parallel-worker.md
- Line: 3
- Current: Description lacks clear trigger keywords
- Change to: Add "Use when: executing parallel tasks, file operations in worktree"
</fixes>

<verification>
After applying all fixes:
1. Verify code-analyzer has exactly: Glob, Grep, Read
2. Verify epic-planner has no Task or Agent tools
3. Verify file-analyzer has exactly: Glob, Grep, LS, Read
4. Verify test-runner has exactly: Bash, Glob, Grep, Read, TodoWrite
5. Verify parallel-worker has constraints section with modal verbs
6. Verify test-runner has constraints section with modal verbs
7. Search for WebFetch in leaf-node agents - should only be in orchestrators
8. Search for Agent tool in non-orchestrator agents - should find none
</verification>

<summary_requirements>
Create .prompts/008-ccpm-remediation-medium/SUMMARY.md

One-liner: "15 over-permissioning issues fixed in 7 files"

Key Findings:
- List agents with reduced tool permissions
- List agents with new constraints sections

Next Step: Run 009-ccpm-remediation-low
</summary_requirements>

<success_criteria>
- All leaf-node agents have minimal required tools
- No leaf nodes have Task or Agent tools
- All code-execution agents have constraints sections
- Modal verbs (MUST/NEVER/ALWAYS) used in constraints
- SUMMARY.md created
</success_criteria>
