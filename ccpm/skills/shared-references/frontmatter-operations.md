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

<validation_rules>
- Frontmatter MUST start on line 1 with `---`
- Closing `---` must be present
- YAML must be valid (proper indentation, quoting)
- Required fields must be present
- Dates must follow ISO 8601 format
</validation_rules>
