# CCPM Changelog

## [2025-12-23] - Skills Architecture & Major Refactoring Release

### 🎯 Overview
Major architectural overhaul introducing skills-based workflow management, migrating from `.claude/` to `ccpm/` directory structure, and comprehensive code quality remediation addressing 148 identified issues across skills, commands, and agents.

### 🏗️ Architecture Changes

- **Directory Structure Migration**
  - Migrated distribution files from `.claude/` to `ccpm/`
  - `.claude/` now reserved for project-specific working files (PRDs, epics)
  - Clear separation between distribution package and local workspace

- **Skills-Based Architecture (NEW)**
  - Introduced `ccpm/skills/` directory with modular skill system
  - Skills: `ccpm-context`, `ccpm-epic`, `ccpm-issue`, `ccpm-prd`, `ccpm-testing`, `ccpm-worktree`
  - Each skill contains SKILL.md, workflows, and references
  - Shared references in `ccpm/skills/shared-references/`

### ✨ Added

- **New Commands**
  - `/doctor` - Diagnose CCPM installation issues
  - `/setup` - Interactive setup wizard
  - `/self-update` - Update CCPM to latest version
  - `/uninstall` - Clean removal with backup option
  - `/version` - Display version and configuration info

- **Skill Workflow Files**
  - Epic workflows: start, sync, decompose, merge, edit, close, oneshot, refresh
  - Issue workflows: start, sync, analyze, edit, close, reopen, status, show
  - Context workflows: create, update, prime
  - Testing workflows: prime, run

- **Shared Reference Documents**
  - `frontmatter-operations.md` - YAML frontmatter handling
  - `github-operations.md` - GitHub CLI integration patterns
  - `worktree-operations.md` - Git worktree management
  - `agent-coordination.md` - Multi-agent coordination patterns
  - `datetime.md` - Date/time handling standards

### 🔄 Changed

- **Commands Transformed to Skill Routers**
  - 23 commands now route to skill-based workflows
  - Reduced duplication and improved maintainability
  - Commands act as entry points, skills contain implementation

- **Agent Enhancements**
  - Agents now integrate with relevant skills
  - Improved tool specifications and capabilities
  - Enhanced coordination patterns for parallel execution

### 🔧 Quality Remediation

- **Skills Audit**: 42 issues identified and resolved
  - Fixed XML structure compliance
  - Added missing tool specifications
  - Standardized workflow patterns

- **Commands Audit**: 61 issues identified and resolved
  - Modernized YAML frontmatter
  - Fixed skill routing patterns
  - Improved argument handling

- **Agents Audit**: 45 issues identified and resolved
  - Enhanced skill integrations
  - Fixed tool permissions
  - Improved return specifications

### 📊 Summary
- **Total Issues Addressed**: 148
- **New Skill Files**: 25+
- **Commands Updated**: 23
- **Agents Enhanced**: 10

---

## [2025-01-24] - Major Cleanup & Issue Resolution Release

### 🎯 Overview
Resolved 10 of 12 open GitHub issues, modernized command syntax, improved documentation, and enhanced system accuracy. This release focuses on stability, usability, and addressing community feedback.

### ✨ Added
- **Local Mode Support** ([#201](https://github.com/automazeio/ccpm/issues/201))
  - Created `LOCAL_MODE.md` with comprehensive offline workflow guide
  - All core commands (prd-new, prd-parse, epic-decompose) work without GitHub
  - Clear distinction between local-only vs GitHub-dependent commands

- **Automatic GitHub Label Creation** ([#544](https://github.com/automazeio/ccpm/issues/544))
  - Enhanced `init.sh` to automatically create `epic` and `task` labels
  - Proper colors: `epic` (green #0E8A16), `task` (blue #1D76DB)  
  - Eliminates manual label setup during project initialization

- **Context Creation Accuracy Safeguards** ([#48](https://github.com/automazeio/ccpm/issues/48))
  - Added mandatory self-verification checkpoints in context commands
  - Implemented evidence-based analysis requirements
  - Added uncertainty flagging with `⚠️ Assumption - requires verification`
  - Enhanced both `/context:create` and `/context:update` with accuracy validation

### 🔄 Changed
- **Modernized Command Syntax** ([#531](https://github.com/automazeio/ccpm/issues/531))
  - Updated 14 PM command files to use concise `!bash` execution pattern
  - Simplified `allowed-tools` frontmatter declarations
  - Reduced token usage and improved Claude Code compatibility

- **Comprehensive README Overhaul** ([#323](https://github.com/automazeio/ccpm/issues/323))
  - Clarified PRD vs Epic terminology and definitions
  - Streamlined workflow explanations and removed redundant sections
  - Fixed installation instructions and troubleshooting guidance
  - Improved overall structure and navigation

### 📋 Research & Community Engagement
- **Multi-Tracker Support Analysis** ([#200](https://github.com/automazeio/ccpm/issues/200))
  - Researched CLI availability for Linear, Trello, Azure DevOps, Jira
  - Identified Linear as best first alternative to GitHub Issues
  - Provided detailed implementation roadmap for future development

- **GitLab Support Research** ([#588](https://github.com/automazeio/ccpm/issues/588))  
  - Confirmed strong `glab` CLI support for GitLab integration
  - Invited community contributor to submit existing GitLab implementation as PR
  - Updated project roadmap to include GitLab as priority platform

### 🐛 Clarified Platform Limitations
- **Windows Shell Compatibility** ([#609](https://github.com/automazeio/ccpm/issues/609))
  - Documented as Claude Code platform limitation (requires POSIX shell)
  - Provided workarounds and alternative solutions

- **Codex CLI Integration** ([#585](https://github.com/automazeio/ccpm/issues/585))
  - Explained future multi-AI provider support in new CLI architecture

- **Parallel Worker Agent Behavior** ([#530](https://github.com/automazeio/ccpm/issues/530))
  - Clarified agent role as coordinator, not direct coder
  - Provided implementation guidance and workarounds

### 🔒 Security
- **Privacy Documentation Fix** ([#630](https://github.com/automazeio/ccpm/issues/630))
  - Verified resolution via PR #631 (remove real repository references)

### 💡 Proposed Features
- **Bug Handling Workflow** ([#654](https://github.com/automazeio/ccpm/issues/654))
  - Designed `/pm:attach-bug` command for automated bug tracking
  - Proposed lightweight sub-issue integration with existing infrastructure
  - Community feedback requested on implementation approach

### 📊 Issues Resolved
**Closed**: 10 issues  
**Active Proposals**: 1 issue (#654)  
**Remaining Open**: 1 issue (#653)

#### Closed Issues:
- [#630](https://github.com/automazeio/ccpm/issues/630) - Privacy: Remove real repo references ✅  
- [#609](https://github.com/automazeio/ccpm/issues/609) - Windows shell error (platform limitation) ✅
- [#585](https://github.com/automazeio/ccpm/issues/585) - Codex CLI compatibility (architecture update) ✅  
- [#571](https://github.com/automazeio/ccpm/issues/571) - Figma MCP support (platform feature) ✅
- [#531](https://github.com/automazeio/ccpm/issues/531) - Use !bash in custom slash commands ✅
- [#323](https://github.com/automazeio/ccpm/issues/323) - Improve README.md ✅
- [#201](https://github.com/automazeio/ccpm/issues/201) - Local-only mode support ✅
- [#200](https://github.com/automazeio/ccpm/issues/200) - Multi-tracker support research ✅  
- [#588](https://github.com/automazeio/ccpm/issues/588) - GitLab support research ✅
- [#48](https://github.com/automazeio/ccpm/issues/48) - Context creation inaccuracies ✅
- [#530](https://github.com/automazeio/ccpm/issues/530) - Parallel worker coding operations ✅
- [#544](https://github.com/automazeio/ccpm/issues/544) - Auto-create labels during init ✅
- [#947](https://github.com/automazeio/ccpm/issues/947) - Project roadmap update ✅

### 🛠️ Technical Details
- **Files Modified**: 16 core files + documentation
- **New Files**: `LOCAL_MODE.md`, `CONTEXT_ACCURACY.md`  
- **Commands Updated**: All 14 PM slash commands modernized
- **Backward Compatibility**: Fully maintained
- **Dependencies**: No new external dependencies added

### 🏗️ Project Health
- **Issue Resolution Rate**: 83% (10/12 issues closed)
- **Documentation Coverage**: Significantly improved
- **Community Engagement**: Active contributor invitation and feedback solicitation
- **Code Quality**: Enhanced accuracy safeguards and validation

### 🚀 Next Steps
1. Community feedback on bug handling proposal (#654)
2. GitLab integration PR review and merge
3. Linear platform integration (pending demand)
4. Enhanced testing and validation workflows

---

*This release represents a major stability and usability milestone for CCPM, addressing the majority of outstanding community issues while establishing a foundation for future multi-platform support.*