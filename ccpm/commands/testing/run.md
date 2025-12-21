---
description: Execute tests with the configured test-runner
argument-hint: [test_target]
allowed-tools: Bash, Read, Task
---

<objective>
Execute tests using the configured test-runner agent. Optionally specify a test file, pattern, or suite name.
</objective>

<context>
Invoke the ccpm-testing skill with action: run
Load skill: @ccpm/skills/ccpm-testing/SKILL.md
Test target: $ARGUMENTS
</context>

<process>
1. Load the ccpm-testing skill
2. Execute the "run" action with target: $ARGUMENTS
3. Report results with pass/fail analysis
</process>

<success_criteria>
- Tests executed successfully
- Clear reporting of pass/fail status
- Actionable feedback on failures
</success_criteria>
