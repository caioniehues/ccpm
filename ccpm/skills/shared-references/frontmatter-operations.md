<overview>
Standard patterns for reading, writing, and updating YAML frontmatter in markdown files. Frontmatter is the metadata block between `---` markers at the start of files.
</overview>

<reading_frontmatter>
Extract frontmatter from markdown files:

1. Look for content between `---` markers at start of file
2. Parse as YAML
3. If invalid or missing, use sensible defaults

```bash
# Extract frontmatter with sed
sed -n '/^---$/,/^---$/p' file.md | tail -n +2 | head -n -1
```
</reading_frontmatter>

<updating_frontmatter>
When modifying existing files:

1. **Preserve all existing fields** - never delete fields you don't understand
2. **Only update specified fields** - don't touch unrelated metadata
3. **Always update `updated` field** - see [datetime.md](datetime.md) for format
4. **Never change `created` field** - this is immutable after file creation
</updating_frontmatter>

<standard_fields>
**All CCPM files:**
```yaml
---
name: {identifier}           # Required: unique identifier
created: {ISO datetime}      # Required: never change after creation
updated: {ISO datetime}      # Required: update on any modification
---
```

**Status values by file type:**
- **PRDs**: `backlog`, `in-progress`, `complete`
- **Epics**: `backlog`, `in-progress`, `completed`
- **Tasks**: `open`, `in-progress`, `closed`

**Progress tracking:**
```yaml
progress: {0-100}%           # For epics
completion: {0-100}%         # For progress files
```
</standard_fields>

<creating_new_files>
Always include frontmatter when creating markdown files:

```yaml
---
name: {from_arguments_or_context}
status: {initial_status}
created: {current_datetime}
updated: {current_datetime}
---
```

The `created` and `updated` fields should be identical on initial creation.
</creating_new_files>

<wizard_approval_fields>
**Wizard approval tracking (used by /pm:epic-wizard):**

PRD approval fields:
```yaml
---
name: feature-name
status: approved              # pending | approved | revision-needed
created: 2025-12-23T09:00:00Z
updated: 2025-12-23T09:15:00Z
approved_at: 2025-12-23T09:15:00Z    # Set when approved
revision_count: 0             # Incremented on each revision
---
```

Epic approval fields:
```yaml
---
name: feature-name
status: approved-for-work     # pending | approved | approved-for-work
created: 2025-12-23T09:15:00Z
updated: 2025-12-23T09:30:00Z
approved_at: 2025-12-23T09:30:00Z
revision_count: 0
prd: feature-name             # Reference to source PRD
---
```

Task status fields:
```yaml
---
id: "001"
name: task-name
status: pending               # pending | in-progress | completed | blocked
created: 2025-12-23T09:30:00Z
updated: 2025-12-23T09:30:00Z
epic: feature-name            # Reference to parent epic
effort: S                     # S | M | L | XL
---
```

**Status transitions:**
- PRD: `pending` → `approved` (or `revision-needed` → `approved`)
- Epic: `pending` → `approved` → `approved-for-work` (after tasks approved)
- Task: `pending` → `in-progress` → `completed` (or `blocked`)
</wizard_approval_fields>

<validation_rules>
- Frontmatter MUST start on line 1 with `---`
- Closing `---` must be present
- YAML must be valid (proper indentation, quoting)
- Required fields must be present
- Dates must follow ISO 8601 format
</validation_rules>
