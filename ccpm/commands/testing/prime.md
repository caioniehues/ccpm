---
description: Configure testing environment with framework detection
allowed-tools: Bash, Read, Write, Glob
---

<objective>
Prepare the testing environment by detecting the test framework, validating dependencies, and configuring the test-runner agent.
</objective>

<context>
Invoke the ccpm-testing skill with action: prime
Load skill: @ccpm/skills/ccpm-testing/SKILL.md
</context>

<process>
1. Load the ccpm-testing skill
2. Execute the "prime" action following skill guidance
3. Detect framework, validate dependencies, save configuration
</process>

<success_criteria>
- Test framework detected and configured
- Configuration saved to .claude/testing-config.md
- Ready status confirmed to user
</success_criteria>

$ARGUMENTS
