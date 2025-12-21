# Workflow: Prime Context

<process>
## Step 1: Preflight Validation

Complete these validation steps silently:

**Context Availability Check:**
- Run: `ls -la .claude/context/ 2>/dev/null`
- If directory doesn't exist or is empty:
  - Tell user: "❌ No context found. Please run create operation first."
  - Exit gracefully
- Count files: `ls -1 .claude/context/*.md 2>/dev/null | wc -l`
- Report: "📁 Found {count} context files to load"

**File Integrity Check:**
For each context file found:
- Verify readable: `test -r ".claude/context/{file}"`
- Check has content: `test -s ".claude/context/{file}"`
- Check for frontmatter (starts with `---`)
- Report issues:
  - Empty: "⚠️ {filename} is empty (skipping)"
  - Unreadable: "⚠️ Cannot read {filename} (permission issue)"
  - Missing frontmatter: "⚠️ {filename} missing frontmatter (may be corrupted)"

**Project State Check:**
- Current state: `git status --short 2>/dev/null`
- Current branch: `git branch --show-current 2>/dev/null`
- Note if not in git repository

## Step 2: Load Context in Priority Order

Load files in this sequence for optimal understanding:

**Priority 1 - Essential Context (load first):**
1. `project-overview.md` - High-level understanding
2. `project-brief.md` - Core purpose and goals
3. `tech-context.md` - Technical stack and dependencies

**Priority 2 - Current State (load second):**
4. `progress.md` - Current status and recent work
5. `project-structure.md` - Directory and file organization

**Priority 3 - Deep Context (load third):**
6. `system-patterns.md` - Architecture and design patterns
7. `product-context.md` - User needs and requirements
8. `project-style-guide.md` - Coding conventions
9. `project-vision.md` - Long-term direction

## Step 3: Validate During Loading

For each file:
- Parse frontmatter and check:
  - `created` date is valid
  - `last_updated` ≥ created date
  - `version` is present
- If frontmatter invalid, note but continue loading content
- Track: files loaded successfully vs failed

## Step 4: Gather Supplementary Information

After loading context:
- Untracked files: `git ls-files --others --exclude-standard | head -20`
- Read README.md if exists
- Check for .env.example or config templates

## Step 5: Error Recovery

**If critical files missing:**
- project-overview.md → Try README.md instead
- tech-context.md → Analyze package files directly
- progress.md → Check recent git commits

**If context incomplete:**
- Inform user which files are missing
- Suggest running update operation
- Continue with partial context but note limitations

## Step 6: Loading Summary

Provide comprehensive summary:

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
  - Status: {current_status from progress.md}
  - Branch: {git_branch}

📊 Key Metrics:
  - Last Updated: {most_recent_update}
  - Context Version: {version}
  - Files Loaded: {success_count}/{total_count}

⚠️ Warnings:
  {list any missing files or issues}

🎯 Ready State:
  ✅ Project context loaded
  ✅ Current status understood
  ✅ Ready for development work

💡 Project Summary:
  {2-3 sentence summary of project and current state}
```

## Step 7: Handle Partial Context

If some files fail to load:
- Continue with available context
- Clearly note what's missing
- Suggest remediation:
  - "Missing technical context - run create operation"
  - "Progress file corrupted - run update operation"

## Step 8: Performance Optimization

For large contexts:
- Load files in parallel when possible
- Show progress: "Loading context files... {current}/{total}"
- Skip extremely large files (>10000 lines) with warning
</process>

<success_criteria>
Context priming is complete when:
- [ ] Context availability verified
- [ ] All available files loaded in priority order
- [ ] Frontmatter validated for each file
- [ ] Project understanding established (type, language, status, branch)
- [ ] Missing or corrupted files reported
- [ ] Supplementary information gathered (untracked files, README)
- [ ] Comprehensive summary provided
- [ ] Ready state confirmed for development work
- [ ] 2-3 sentence project summary provided
</success_criteria>
