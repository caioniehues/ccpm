---
description: Clean up completed work and archive old epics
argument-hint: [--dry-run]
allowed-tools: Bash, Read, Write, LS
---

<objective>
Clean up completed work by archiving old epics and removing stale files.
</objective>

<process>
**Usage**: `/pm:clean [--dry-run]`

Options:
- `--dry-run` - Show what would be cleaned without doing it

**1. Identify Completed Epics**

Find epics with:
- `status: completed` in frontmatter
- All tasks closed
- Last update > 30 days ago

**2. Identify Stale Work**

Find:
- Progress files for closed issues
- Update directories for completed work
- Orphaned task files (epic deleted)
- Empty directories

**3. Show Cleanup Plan**

```
🧹 Cleanup Plan

Completed Epics to Archive:
  {epic_name} - Completed {days} days ago

Stale Progress to Remove:
  {count} progress files for closed issues

Empty Directories:
  {list_of_empty_dirs}

Space to Recover: ~{size}KB

{If --dry-run}: This is a dry run. No changes made.
{Otherwise}: Proceed with cleanup? (yes/no)
```

**4. Execute Cleanup**

If user confirms:

- Archive Epics:
```bash
mkdir -p .claude/epics/.archived
mv .claude/epics/{completed_epic} .claude/epics/.archived/
```

- Remove Stale Files:
  - Delete progress files for closed issues > 30 days
  - Remove empty update directories
  - Clean up orphaned files

- Create Archive Log at `.claude/epics/.archived/archive-log.md`

**5. Output**

```
✅ Cleanup Complete

Archived:
  {count} completed epics

Removed:
  {count} stale files
  {count} empty directories

Space recovered: {size}KB
```
</process>

<success_criteria>
- Cleanup plan shown to user before execution
- Completed epics archived to .archived directory
- Stale files removed
- Archive log updated
- No deletion of PRDs or incomplete work
</success_criteria>
