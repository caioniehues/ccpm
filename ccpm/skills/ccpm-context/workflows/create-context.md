# Workflow: Create Context

<process>
<step_preflight>
**Preflight Validation**

Complete these validation steps silently (don't report progress to user):

**Context Directory Check:**
- Run: `ls -la .claude/context/ 2>/dev/null`
- If directory exists and has files:
  - Count: `ls -1 .claude/context/*.md 2>/dev/null | wc -l`
  - Ask: "⚠️ Found {count} existing context files. Overwrite all context? (yes/no)"
  - Only proceed with explicit 'yes' confirmation
  - If no, suggest: "Use update operation to refresh existing context"

**Project Type Detection:**
- Run: `find . -maxdepth 2 \( -name 'package.json' -o -name 'requirements.txt' -o -name 'pyproject.toml' -o -name 'pom.xml' -o -name 'build.gradle' -o -name 'build.gradle.kts' -o -name '*.sln' -o -name '*.csproj' -o -name 'Gemfile' -o -name 'Cargo.toml' -o -name 'go.mod' -o -name 'composer.json' -o -name 'pubspec.yaml' -o -name 'CMakeLists.txt' -o -name 'Dockerfile' -o -name 'docker-compose.yml' -o -name 'Package.swift' -o -type d -name '*.xcodeproj' -o -type d -name '*.xcworkspace' \) 2>/dev/null` to detect project type
- Run: `git status 2>/dev/null` to confirm git repository
- If not a git repo, ask: "⚠️ Not a git repository. Continue anyway? (yes/no)"

**Directory Creation:**
- Create if needed: `mkdir -p .claude/context/`
- Verify write permissions: `touch .claude/context/.test && rm .claude/context/.test`
- If permission denied: "❌ Cannot create context directory. Check permissions." and exit

**Get Current DateTime:**
- Run: `date -u +"%Y-%m-%dT%H:%M:%SZ"`
- Store for use in all context file frontmatter
</step_preflight>

<step_pre_analysis>
**Pre-Analysis**

**Validation:**
- Confirm project root directory (presence of .git, package files, etc.)
- Check for README.md or other documentation
- If README.md doesn't exist, ask user for project description

**Project Information Gathering:**
- Repository info: `git remote -v 2>/dev/null`
- Current branch: `git branch --show-current 2>/dev/null`
- Root structure: `ls -la`
- Code files: `find . -type f \( -name '*.js' -o -name '*.ts' -o -name '*.jsx' -o -name '*.tsx' -o -name '*.py' -o -name '*.rs' -o -name '*.go' -o -name '*.php' -o -name '*.swift' -o -name '*.java' -o -name '*.kt' -o -name '*.kts' -o -name '*.cs' -o -name '*.rb' -o -name '*.dart' -o -name '*.c' -o -name '*.h' -o -name '*.cpp' -o -name '*.hpp' -o -name '*.sh' \) 2>/dev/null | head -20`
- Read README.md if exists
</step_pre_analysis>

<step_create_files>
**Create Context Files**

Create each file with frontmatter template:
```yaml
---
created: [Real datetime from date command]
last_updated: [Real datetime from date command]
version: 1.0
author: Claude Code PM System
---
```

**Generate these files in .claude/context/:**

1. **progress.md** - Current project status, completed work, immediate next steps
   - Current branch, recent commits (`git log --oneline -10`), outstanding changes

2. **project-structure.md** - Directory structure and file organization
   - Key directories, file naming patterns, module organization

3. **tech-context.md** - Dependencies, technologies, development tools
   - Language version, framework versions, dev dependencies from package files

4. **system-patterns.md** - Architectural patterns and design decisions
   - Design patterns observed in code, architectural style, data flow

5. **product-context.md** - Product requirements, target users, core functionality
   - User personas (if identifiable), core features, use cases

6. **project-brief.md** - Project scope, goals, key objectives
   - What it does, why it exists, success criteria

7. **project-overview.md** - High-level summary of features and capabilities
   - Feature list, current state, integration points

8. **project-vision.md** - Long-term vision and strategic direction
   - Future goals, potential expansions, strategic priorities

9. **project-style-guide.md** - Coding standards, conventions, style preferences
   - Naming conventions, file structure patterns, comment style
</step_create_files>

<step_quality_validation>
**Quality Validation**

For each file created:
- Verify file exists: `test -f .claude/context/{filename}`
- Check not empty: `test -s .claude/context/{filename}`
- Count lines: `wc -l .claude/context/{filename}` (minimum 10 lines)
- Verify frontmatter present (starts with `---`)

If any file fails:
- Report which files succeeded
- Provide option to continue with partial context
- Never leave corrupted or incomplete files
</step_quality_validation>

<step_error_handling>
**Error Handling**

**Common Issues:**
- No write permissions: "❌ Cannot write to .claude/context/. Check permissions."
- Disk space: "❌ Insufficient disk space for context files."
- File creation failed: "❌ Failed to create {filename}. Error: {error}"

If errors occur:
- Report successfully created files
- Clean up any corrupted files
- Provide actionable guidance to user
</step_error_handling>

<step_summary>
**Post-Creation Summary**

Provide comprehensive summary:

```
📋 Context Creation Complete

📁 Created context in: .claude/context/
✅ Files created: {count}/9

📊 Context Summary:
  - Project Type: {detected_type}
  - Language: {primary_language}
  - Git Status: {clean/changes}
  - Dependencies: {count} packages

📝 File Details:
  ✅ progress.md ({lines} lines) - Current status and recent work
  ✅ project-structure.md ({lines} lines) - Directory organization
  ✅ tech-context.md ({lines} lines) - Technical stack
  ✅ system-patterns.md ({lines} lines) - Architecture patterns
  ✅ product-context.md ({lines} lines) - Product requirements
  ✅ project-brief.md ({lines} lines) - Project scope
  ✅ project-overview.md ({lines} lines) - Feature overview
  ✅ project-vision.md ({lines} lines) - Strategic vision
  ✅ project-style-guide.md ({lines} lines) - Coding standards

⏰ Created: {timestamp}
🔄 Next: Use prime operation to load context in new sessions
💡 Tip: Run update operation regularly to keep context current
```
</step_summary>
</process>

<success_criteria>
Context creation is complete when:
- [ ] All preflight checks passed
- [ ] Current datetime retrieved from system
- [ ] All 9 context files created successfully
- [ ] Each file has valid frontmatter with real datetime
- [ ] Each file has meaningful content (minimum 10 lines)
- [ ] All files validated (exist, readable, not corrupted)
- [ ] Comprehensive summary provided to user
- [ ] No corrupted or incomplete files left behind
</success_criteria>
