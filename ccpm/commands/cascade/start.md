---
description: Begin new Cascade Flow session with multi-specialist brainstorming
argument-hint: <feature-name>
allowed-tools: Skill, Task, Read, Write, Glob, Bash, AskUserQuestion
---

<objective>
Start a new Cascade Flow session for the specified feature. Initiates Phase 1
(multi-specialist brainstorming) with 5 parallel perspective agents and
background codebase research.
</objective>

<process>
1. Invoke the ccpm-cascade skill with the feature name
2. The skill will route to cascade-brainstorm workflow
3. Follow the 4-phase flow: Brainstorm → PRD → Decompose → Execute
</process>

Invoke the ccpm-cascade skill for: $ARGUMENTS

<success_criteria>
- Brainstorm session created in .claude/brainstorm/
- 5 perspective agents launched successfully
- Background research initiated
- User engaged in interactive refinement
- Session ready for Phase 2 upon approval
</success_criteria>
