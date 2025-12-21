---
description: Run health checks on CCPM installation and dependencies
allowed-tools: Bash(which:*), Bash(gh:*), Bash(git:*), Bash(ls:*), Bash(test:*), Read, Glob
---

<objective>
Diagnose CCPM installation health by checking dependencies, configuration, and project structure.
</objective>

<process>
**1. Core Dependencies**

Check required tools are installed:

```
!`which git 2>/dev/null && echo "✅ git installed" || echo "❌ git missing"`
!`which gh 2>/dev/null && echo "✅ gh CLI installed" || echo "❌ gh CLI missing"`
!`gh auth status 2>&1 | grep -q "Logged in" && echo "✅ gh authenticated" || echo "⚠️ gh not authenticated"`
```

**2. Project Structure**

Verify CCPM directories exist:

```
!`test -d .claude && echo "✅ .claude directory exists" || echo "❌ .claude directory missing"`
!`test -d .claude/epics && echo "✅ epics directory exists" || echo "⚠️ epics directory missing"`
!`test -d .claude/prds && echo "✅ PRDs directory exists" || echo "⚠️ PRDs directory missing"`
!`test -d .claude/context && echo "✅ context directory exists" || echo "⚠️ context directory missing"`
```

**3. Git Repository**

Check git status:

```
!`git rev-parse --git-dir 2>/dev/null && echo "✅ git repository detected" || echo "❌ not a git repository"`
!`git remote -v 2>/dev/null | head -1 || echo "⚠️ no remote configured"`
```

**4. CCPM Installation**

Verify CCPM files:

```
!`test -f ccpm/VERSION && cat ccpm/VERSION || echo "⚠️ VERSION file missing"`
!`ls ccpm/commands/pm/*.md 2>/dev/null | wc -l | xargs -I {} echo "📁 {} PM commands found"`
!`ls ccpm/scripts/pm/*.sh 2>/dev/null | wc -l | xargs -I {} echo "📁 {} PM scripts found"`
```

**5. Summary**

Report overall health status based on checks above.
</process>

<success_criteria>
- All core dependencies installed and functional
- Project structure valid
- Git repository configured correctly
- CCPM installation complete
</success_criteria>
