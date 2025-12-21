---
description: Execute complex prompts from a file when direct input fails
allowed-tools: Bash, Read, Write, LS
---

<objective>
Execute a prompt written to a file when complex prompts with numerous @ references fail in direct input.
</objective>

<process>
This is an ephemeral command for handling complex prompts.

Some complex prompts (with numerous @ references) may fail if entered directly into the prompt input.

If that happens, write your prompt to this file and run `/prompt` to execute it.
</process>

<success_criteria>
Prompt content processed and executed successfully.
</success_criteria>
