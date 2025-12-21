---
description: Manage project context (create, update, or prime context documentation)
argument-hint: [operation]
allowed-tools: Skill(ccpm-context)
---

<objective>
Invoke the ccpm-context skill for context management operations.
</objective>

<process>
Invoke the ccpm-context skill for: $ARGUMENTS

Available operations:
- prime: Generate initial context documentation from codebase
- create: Create a new context file
- update: Update existing context files
</process>

<success_criteria>
Context operation completed via ccpm-context skill.
</success_criteria>
