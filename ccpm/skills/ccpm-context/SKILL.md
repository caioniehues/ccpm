---
name: ccpm-context
description: Manages project context for CCPM by creating, updating, and priming context documentation in .claude/context/. Use when working with project context, establishing baseline documentation, or loading context for new sessions.
---

<objective>
Manage conversation context files for CCPM workflows including creation, priming, and updates of project documentation in `.claude/context/`.
</objective>

<essential_principles>
<overview>
Context management in CCPM maintains comprehensive project documentation that helps Claude understand your project's current state, architecture, and progress.
</overview>

<context_structure>
Context files live in `.claude/context/` and include:
- **progress.md** - Current status, recent work, next steps
- **project-structure.md** - Directory organization and file patterns
- **tech-context.md** - Dependencies, technologies, and tools
- **system-patterns.md** - Architectural patterns and design decisions
- **product-context.md** - Requirements, users, and functionality
- **project-brief.md** - Scope, goals, and objectives
- **project-overview.md** - Feature summary and capabilities
- **project-vision.md** - Long-term direction and strategy
- **project-style-guide.md** - Coding standards and conventions
</context_structure>

<frontmatter_requirements>
All context files MUST include frontmatter with real datetime:
```yaml
---
created: [Real datetime from system clock]
last_updated: [Real datetime from system clock]
version: 1.0
author: Claude Code PM System
---
```
</frontmatter_requirements>

<context_lifecycle>
1. **Create** - Establish initial baseline documentation by analyzing current project state
2. **Prime** - Load essential context for new agent sessions to understand the project
3. **Update** - Refresh context to reflect current state (run regularly, especially after significant changes)
</context_lifecycle>

<key_principles>
- **Always use real datetime** from system clock (`date -u +"%Y-%m-%dT%H:%M:%SZ"`)
- **Validate before proceeding** - Check for existing context, permissions, git status
- **Handle errors gracefully** - Provide specific guidance when issues occur
- **Make surgical updates** - Don't regenerate entire files unnecessarily
- **Preserve accurate timestamps** - Only update `last_updated` when content actually changes
</key_principles>
</essential_principles>

<intake>
**What would you like to do with project context?**

1. **Create** - Establish initial context documentation (analyze project and create all context files)
2. **Prime** - Load existing context for a new session (read context to understand project)
3. **Update** - Refresh context to reflect current state (update files based on recent changes)
4. Something else (please specify)

**Wait for response before proceeding.**
</intake>

<routing>
| Response | Workflow | When to Use |
|----------|----------|-------------|
| 1, "create", "establish", "initialize", "new" | `workflows/create-context.md` | First time setup, complete rebuild, or when context is missing |
| 2, "prime", "load", "read", "understand" | `workflows/prime-context.md` | Starting a new session, getting up to speed on the project |
| 3, "update", "refresh", "sync", "current" | `workflows/update-context.md` | After development work, before ending session, when changes occurred |
| 4, other | Clarify intent, then route to appropriate workflow |

**After reading the workflow, follow it exactly.**
</routing>

<workflows_index>
| Workflow | Purpose |
|----------|---------|
| create-context.md | Create initial project context by analyzing codebase and establishing baseline documentation |
| prime-context.md | Load essential context for new agent session by reading existing documentation |
| update-context.md | Update context to reflect current project state based on recent changes |
</workflows_index>

<quick_start>
**Quick start:**

```bash
# Load context for a new session
/pm:context prime

# Create initial context documentation
/pm:context create

# Update context after changes
/pm:context update
```
</quick_start>

<success_criteria>
Context management is successful when:

**For Create:**
- All 9 context files created with valid frontmatter
- Real datetime used (not placeholders)
- Each file has meaningful content (minimum 10 lines)
- Files accurately reflect current project state
- Summary provided showing what was created

**For Prime:**
- All available context files loaded in priority order
- Project understanding established (type, language, status, branch)
- Any missing or corrupted files reported
- Ready state confirmed for development work
- 2-3 sentence project summary provided

**For Update:**
- Only files with actual changes updated
- Real datetime used for `last_updated` field
- Surgical updates made (not complete regeneration)
- Skipped files reported (with reason)
- Summary shows what changed and what didn't
</success_criteria>
