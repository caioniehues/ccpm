---
name: ccpm-context
description: Manage project context documentation for Claude Code sessions. Supports creating, loading (priming), and updating context files in .claude/context/.
---

<objective>
Manage comprehensive project context documentation that enables Claude Code to understand and work effectively with any codebase. Context files capture project structure, technical stack, patterns, and current progress - ensuring continuity across sessions.

This skill handles three core actions:
- **create**: Initialize full context documentation from scratch
- **prime**: Load existing context at session start
- **update**: Refresh context with recent changes
</objective>

<shared_references>
Load these before any operation:
- @ccpm/skills/shared-references/datetime.md
- @ccpm/skills/shared-references/frontmatter-operations.md
</shared_references>

<action name="create">
<description>
Create initial project context documentation in `.claude/context/` by analyzing the current project state and establishing comprehensive baseline documentation.
</description>

<preflight>
1. **Context Directory Check**
   - Run: `ls -la .claude/context/ 2>/dev/null`
   - If exists with files, count and ask user: "⚠️ Found {count} existing context files. Overwrite all? (yes/no)"
   - Only proceed with explicit 'yes' confirmation
   - If no, suggest: "Use /context:update to refresh existing context"

2. **Project Type Detection**
   - Check for: package.json (Node.js), requirements.txt/pyproject.toml (Python), Cargo.toml (Rust), go.mod (Go), pom.xml/build.gradle (Java), *.csproj (C#), composer.json (PHP), Gemfile (Ruby), pubspec.yaml (Dart), Package.swift (Swift), CMakeLists.txt (C/C++)
   - Run: `git status 2>/dev/null` to confirm git repository
   - If not git repo, ask: "⚠️ Not a git repository. Continue anyway? (yes/no)"

3. **Directory Creation**
   - Create `.claude/context/` if missing: `mkdir -p .claude/context/`
   - Verify write permissions: `touch .claude/context/.test && rm .claude/context/.test`

4. **Get Current DateTime**
   - Run: `date -u +"%Y-%m-%dT%H:%M:%SZ"` and store for frontmatter
</preflight>

<process>
1. **Pre-Analysis**
   - Confirm project root (presence of .git, package.json, etc.)
   - Check for existing documentation (README.md, docs/)
   - If no README.md, ask user for project description

2. **Systematic Project Analysis**
   - Project detection: find build/config files up to 2 levels deep
   - Git info: `git remote -v`, `git branch --show-current`
   - Codebase scan: find source files by extension (limit to 20)
   - Directory structure: `ls -la`

3. **Create Context Files with Frontmatter**
   Each file MUST include:
   ```yaml
   ---
   created: {REAL datetime from date command}
   last_updated: {REAL datetime from date command}
   version: 1.0
   author: Claude Code PM System
   ---
   ```

   Generate these files:
   - `progress.md` - Current status, completed work, next steps
   - `project-structure.md` - Directory structure, file organization
   - `tech-context.md` - Dependencies, technologies, dev tools
   - `system-patterns.md` - Architectural patterns, design decisions
   - `product-context.md` - Requirements, target users, core functionality
   - `project-brief.md` - Scope, goals, objectives
   - `project-overview.md` - Features, capabilities summary
   - `project-vision.md` - Long-term vision, strategic direction
   - `project-style-guide.md` - Coding standards, conventions

4. **Quality Validation**
   - Verify each file created successfully
   - Check minimum 10 lines of content
   - Ensure valid frontmatter and markdown

5. **Post-Creation Summary**
   ```
   📋 Context Creation Complete

   📁 Created context in: .claude/context/
   ✅ Files created: {count}/9

   📊 Context Summary:
     - Project Type: {detected_type}
     - Language: {primary_language}
     - Git Status: {clean/changes}
     - Dependencies: {count} packages

   ⏰ Created: {timestamp}
   🔄 Next: Use /context:prime to load context in new sessions
   💡 Tip: Run /context:update regularly to keep context current
   ```
</process>

<error_handling>
- **No write permissions**: "❌ Cannot write to .claude/context/. Check permissions."
- **Disk space**: "❌ Insufficient disk space for context files."
- **File creation failed**: "❌ Failed to create {filename}. Error: {error}"
- Never leave corrupted or incomplete files
</error_handling>
</action>

<action name="prime">
<description>
Load essential context for a new agent session by reading project context documentation and understanding the codebase structure.
</description>

<preflight>
1. **Context Availability**
   - Run: `ls -la .claude/context/ 2>/dev/null`
   - If missing or empty: "❌ No context found. Please run /context:create first."
   - Count files: `ls -1 .claude/context/*.md 2>/dev/null | wc -l`
   - Report: "📁 Found {count} context files to load"

2. **File Integrity Check**
   - For each file: verify readable, has content, has frontmatter
   - Report issues: empty files, unreadable files, missing frontmatter

3. **Project State**
   - Run: `git status --short 2>/dev/null`
   - Run: `git branch --show-current 2>/dev/null`
</preflight>

<process>
1. **Load in Priority Order**

   **Priority 1 - Essential (load first):**
   - `project-overview.md` - High-level understanding
   - `project-brief.md` - Core purpose and goals
   - `tech-context.md` - Technical stack

   **Priority 2 - Current State (load second):**
   - `progress.md` - Status and recent work
   - `project-structure.md` - Directory organization

   **Priority 3 - Deep Context (load third):**
   - `system-patterns.md` - Architecture patterns
   - `product-context.md` - User needs and requirements
   - `project-style-guide.md` - Coding conventions
   - `project-vision.md` - Long-term direction

2. **Validation During Loading**
   - Parse frontmatter: check created, last_updated, version
   - Track successful vs failed loads

3. **Supplementary Information**
   - Check untracked files: `git ls-files --others --exclude-standard | head -20`
   - Read README.md if exists
   - Check for .env.example

4. **Error Recovery**
   - Missing `project-overview.md`: Try README.md
   - Missing `tech-context.md`: Analyze config files directly
   - Missing `progress.md`: Check recent git commits
   - Continue with partial context but note limitations

5. **Loading Summary**
   ```
   🧠 Context Primed Successfully

   📖 Loaded Context Files:
     ✅ Essential: {count}/3 files
     ✅ Current State: {count}/2 files
     ✅ Deep Context: {count}/4 files

   🔍 Project Understanding:
     - Name: {project_name}
     - Type: {project_type}
     - Language: {primary_language}
     - Status: {current_status}
     - Branch: {git_branch}

   📊 Key Metrics:
     - Last Updated: {most_recent_update}
     - Files Loaded: {success_count}/{total_count}

   🎯 Ready State:
     ✅ Project context loaded
     ✅ Ready for development work

   💡 Project Summary:
     {2-3 sentence summary}
   ```
</process>

<error_handling>
- Handle missing files gracefully - don't fail completely
- Load what's available, note what's missing
- Suggest remediation: "/context:create to rebuild" or "/context:update to refresh"
</error_handling>
</action>

<action name="update">
<description>
Update project context documentation to reflect current state. Run at end of development sessions to keep context accurate.
</description>

<preflight>
1. **Context Validation**
   - Run: `ls -la .claude/context/ 2>/dev/null`
   - If missing or empty: "❌ No context to update. Run /context:create first."
   - Count files: `ls -1 .claude/context/*.md 2>/dev/null | wc -l`

2. **Change Detection**
   - Git changes: `git status --short`, `git log --oneline -10`
   - Recent file changes: `git diff --stat HEAD~5..HEAD 2>/dev/null`
   - Dependency changes: check package files for modifications

3. **Get Current DateTime**
   - Run: `date -u +"%Y-%m-%dT%H:%M:%SZ"` for `last_updated` field
</preflight>

<process>
1. **Systematic Change Analysis**
   Determine which files need updates:

   - **progress.md** - ALWAYS UPDATE: Recent commits, current branch, next steps
   - **project-structure.md** - If new files/directories added
   - **tech-context.md** - If dependencies changed
   - **system-patterns.md** - If architecture changed
   - **product-context.md** - If requirements changed
   - **project-brief.md** - Rarely (only if goals changed)
   - **project-overview.md** - For major milestones
   - **project-vision.md** - Rarely (only strategic shifts)
   - **project-style-guide.md** - If conventions changed

2. **Smart Update Strategy**
   For each file needing update:
   - Read existing content
   - Identify specific sections to update
   - Preserve frontmatter `created` field
   - Update `last_updated` with real datetime
   - Increment version if major update (1.0 → 1.1)
   - Make targeted updates, don't rewrite entire file

3. **Update Validation**
   - Verify valid frontmatter after update
   - Check file size is reasonable
   - Ensure markdown formatting preserved

4. **Skip Optimization**
   - Don't update files with no changes
   - Preserve accurate timestamps

5. **Update Summary**
   ```
   🔄 Context Update Complete

   📊 Update Statistics:
     - Files Scanned: {total_count}
     - Files Updated: {updated_count}
     - Files Skipped: {skipped_count}
     - Errors: {error_count}

   📝 Updated Files:
     ✅ progress.md - Updated recent commits, current status
     ✅ tech-context.md - Added new dependencies

   ⏭️ Skipped Files (no changes):
     - project-brief.md (last updated: 5 days ago)
     - project-vision.md (last updated: 2 weeks ago)

   ⏰ Last Update: {timestamp}
   🔄 Next: Run this command regularly to keep context current
   ```
</process>

<error_handling>
- **File locked**: "❌ Cannot update {file} - may be open in editor"
- **Permission denied**: "❌ Cannot write to {file} - check permissions"
- **Corrupted file**: "⚠️ {file} appears corrupted - skipping update"
- Never leave corrupted state
</error_handling>
</action>

<context_gathering_commands>
Common commands for gathering project information:
- Target directory: `.claude/context/`
- Git status: `git status --short`
- Recent commits: `git log --oneline -10`
- Changed files: `git diff --name-only HEAD~5..HEAD 2>/dev/null`
- Branch info: `git branch --show-current`
- Documentation: `find . -type f -name '*.md' -path '*/docs/*' 2>/dev/null | head -10`
- Tests: `find . -type d \( -name 'test' -o -name 'tests' -o -name '__tests__' \) 2>/dev/null | head -10`
</context_gathering_commands>

<success_criteria>
- **create**: All 9 context files created with valid frontmatter and content
- **prime**: Context loaded successfully with clear project understanding
- **update**: Changed files updated, unchanged files skipped, valid state maintained
- All operations use real datetime from system clock
- Clear summaries provided to user
- Errors handled gracefully with actionable guidance
</success_criteria>
