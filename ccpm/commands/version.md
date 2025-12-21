---
description: Display CCPM version information
allowed-tools: Bash(cat:*), Bash(git:*), Read
---

<objective>
Show current CCPM version and related information.
</objective>

<process>
Display version info:

```
!`echo "CCPM - Claude Code Project Management"`
!`echo ""`
!`test -f ccpm/VERSION && echo "Version: $(cat ccpm/VERSION)" || echo "Version: unknown"`
!`git log -1 --format="Commit: %h (%ci)" 2>/dev/null || echo ""`
!`git describe --tags 2>/dev/null && echo "" || echo ""`
```
</process>

<success_criteria>
Version information displayed clearly.
</success_criteria>
