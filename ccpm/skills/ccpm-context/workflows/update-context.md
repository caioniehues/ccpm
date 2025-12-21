# Workflow: Update Context

<process>
## Step 1: Preflight Validation

Complete these validation steps silently:

**Context Validation:**
- Run: `ls -la .claude/context/ 2>/dev/null`
- If directory doesn't exist or is empty:
  - Tell user: "❌ No context to update. Please run create operation first."
  - Exit gracefully
- Count files: `ls -1 .claude/context/*.md 2>/dev/null | wc -l`
- Report: "📁 Found {count} context files to check for updates"

**Change Detection:**

Git Changes:
- Uncommitted: `git status --short`
- Recent commits: `git log --oneline -10`
- Files changed: `git diff --stat HEAD~5..HEAD 2>/dev/null`

File Modifications:
- Context file ages: `find .claude/context -name "*.md" -type f -exec ls -lt {} + | head -5`

Dependency Changes:
- Check for changes in package files:
  - Node.js: `git diff HEAD~5..HEAD package.json 2>/dev/null`
  - Python: `git diff HEAD~5..HEAD requirements.txt pyproject.toml 2>/dev/null`
  - Java: `git diff HEAD~5..HEAD pom.xml build.gradle build.gradle.kts 2>/dev/null`
  - C#/.NET: `git diff HEAD~5..HEAD *.sln *.csproj 2>/dev/null`
  - Ruby: `git diff HEAD~5..HEAD Gemfile Gemfile.lock 2>/dev/null`
  - Rust: `git diff HEAD~5..HEAD Cargo.toml Cargo.lock 2>/dev/null`
  - Go: `git diff HEAD~5..HEAD go.mod go.sum 2>/dev/null`
  - PHP: `git diff HEAD~5..HEAD composer.json composer.lock 2>/dev/null`
  - Dart/Flutter: `git diff HEAD~5..HEAD pubspec.yaml pubspec.lock 2>/dev/null`
  - Swift: `git diff HEAD~5..HEAD Package.swift Package.resolved 2>/dev/null`
  - C/C++: `git diff HEAD~5..HEAD CMakeLists.txt 2>/dev/null`

**Get Current DateTime:**
- Run: `date -u +"%Y-%m-%dT%H:%M:%SZ"`
- Store for updating `last_updated` field

## Step 2: Systematic Change Analysis

Determine which files need updates:

**progress.md - ALWAYS UPDATE:**
- Check: `git log --oneline -5`
- Update: Latest completed work, current blockers, next steps
- Include completion percentages if applicable

**project-structure.md - Update if Changed:**
- Check: `git diff --name-status HEAD~10..HEAD | grep -E '^A'` for new files
- Update: Only if significant structural changes (new directories, reorganization)
- Skip if only file content changed

**tech-context.md - Update if Dependencies Changed:**
- Check: Package files for new dependencies or version changes
- Update: New libraries, upgraded versions, new dev tools
- Include security updates or breaking changes

**system-patterns.md - Update if Architecture Changed:**
- Check: New design patterns, architectural decisions
- Update: Only for significant architectural changes
- Skip for minor refactoring

**product-context.md - Update if Requirements Changed:**
- Check: New features implemented, user feedback incorporated
- Update: New user stories, changed requirements
- Include any pivot in product direction

**project-brief.md - Rarely Update:**
- Check: Only if fundamental project goals changed
- Update: Major scope changes, new objectives
- Usually remains stable

**project-overview.md - Update for Major Milestones:**
- Check: Major features completed, significant progress
- Update: Feature status, capability changes
- Update when reaching project milestones

**project-vision.md - Rarely Update:**
- Check: Strategic direction changes
- Update: Only for major vision shifts
- Usually remains stable

**project-style-guide.md - Update if Conventions Changed:**
- Check: New linting rules, style decisions
- Update: Convention changes, new patterns adopted
- Include examples of new patterns

## Step 3: Smart Update Strategy

For each file that needs updating:

1. **Read existing file** to understand current content
2. **Identify specific sections** that need updates
3. **Preserve frontmatter** but update `last_updated`:
   ```yaml
   ---
   created: [preserve original]
   last_updated: [Real datetime from date command]
   version: [increment if major update, e.g., 1.0 → 1.1]
   author: Claude Code PM System
   ---
   ```
4. **Make targeted updates** - don't rewrite entire file
5. **Add update notes** if significant:
   ```markdown
   ## Update History
   - {date}: {summary of what changed}
   ```

## Step 4: Update Validation

After updating each file:
- Verify frontmatter still valid
- Check file size reasonable (not corrupted)
- Ensure markdown formatting preserved
- Confirm updates accurately reflect changes

## Step 5: Skip Optimization

**Skip files without changes:**
- If no relevant changes detected, skip the file
- Don't update timestamp if content unchanged
- This preserves accurate "last modified" information
- Report skipped files in summary

## Step 6: Error Handling

**Common Issues:**
- File locked: "❌ Cannot update {file} - may be open in editor"
- Permission denied: "❌ Cannot write to {file} - check permissions"
- Corrupted file: "⚠️ {file} appears corrupted - skipping update"
- Disk space: "❌ Insufficient disk space for updates"

If update fails:
- Report successfully updated files
- Note failed files and why
- Preserve original files (don't corrupt)

## Step 7: Update Summary

Provide detailed summary:

```
🔄 Context Update Complete

📊 Update Statistics:
  - Files Scanned: {total_count}
  - Files Updated: {updated_count}
  - Files Skipped: {skipped_count} (no changes needed)
  - Errors: {error_count}

📝 Updated Files:
  ✅ progress.md - Updated recent commits, current status
  ✅ tech-context.md - Added {count} new dependencies
  ✅ project-structure.md - Noted new /{directory} directory

⏭️ Skipped Files (no changes):
  - project-brief.md (last updated: {time_ago})
  - project-vision.md (last updated: {time_ago})
  - system-patterns.md (last updated: {time_ago})

⚠️ Issues:
  {any warnings or errors}

⏰ Last Update: {timestamp}
🔄 Next: Run this operation regularly to keep context current
💡 Tip: Major changes? Consider running create operation for full refresh
```

## Step 8: Incremental Update Tracking

Track what was updated:
- Note which sections of each file were modified
- Keep changes focused and surgical
- Don't regenerate unchanged content
- Preserve formatting and structure

## Step 9: Performance Optimization

For large projects:
- Process files in parallel when possible
- Show progress: "Updating context files... {current}/{total}"
- Skip very large files with warning
- Use git diff to quickly identify changed areas
</process>

<success_criteria>
Context update is complete when:
- [ ] Context validation passed
- [ ] Current datetime retrieved from system
- [ ] Change detection completed (git, dependencies, files)
- [ ] Each file analyzed for update necessity
- [ ] Only files with actual changes updated
- [ ] Real datetime used for `last_updated` field
- [ ] Surgical updates made (not complete regeneration)
- [ ] All updates validated
- [ ] Skipped files reported with reasons
- [ ] Comprehensive summary provided
- [ ] No files corrupted or left in invalid state
</success_criteria>
