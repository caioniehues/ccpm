---
description: Re-initialize or update CLAUDE.md with rules from .claude/CLAUDE.md
allowed-tools: Bash, Read, Write, LS
---

<objective>
Update CLAUDE.md file with the latest rules from .claude/CLAUDE.md.
</objective>

<process>
1. Check if CLAUDE.md exists in the root directory
2. If not, run /init to create it
3. Read rules from .claude/CLAUDE.md
4. Update CLAUDE.md to include those rules
</process>

<success_criteria>
CLAUDE.md updated with current rules from .claude/CLAUDE.md.
</success_criteria>
