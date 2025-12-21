# Frontmatter Operations

<overview>
CCPM uses YAML frontmatter to store task metadata. All frontmatter operations must preserve structure and use proper YAML syntax.
</overview>

<frontmatter_structure>
## Task File Frontmatter

```yaml
---
name: Task Title
status: open|in-progress|blocked|closed
created: 2024-01-15T10:30:00Z
updated: 2024-01-15T14:45:00Z
github: https://github.com/org/repo/issues/123
priority: high|medium|low
size: small|medium|large
---
```

**Required fields:**
- `name`: Task title (string)
- `status`: Current status (enum)
- `created`: Creation timestamp (ISO 8601)
- `updated`: Last update timestamp (ISO 8601)
- `github`: GitHub issue URL (string)

**Optional fields:**
- `priority`: Task priority
- `size`: Estimated size
</frontmatter_structure>

<progress_file_frontmatter>
## Progress File Frontmatter

```yaml
---
issue: 123
started: 2024-01-15T10:30:00Z
last_sync: 2024-01-15T14:45:00Z
completion: 65
---
```

**Required fields:**
- `issue`: Issue number (integer)
- `started`: When work began (ISO 8601)
- `last_sync`: Last GitHub sync (ISO 8601)
- `completion`: Progress percentage (0-100)
</progress_file_frontmatter>

<epic_frontmatter>
## Epic Frontmatter

```yaml
---
name: Epic Name
status: planning|in-progress|completed
created: 2024-01-15T10:30:00Z
progress: 45
prd: .claude/epics/epic-name/PRD.md
github: https://github.com/org/repo/issues/100
---
```

**Required fields:**
- `name`: Epic name (string)
- `status`: Epic status (enum)
- `created`: Creation timestamp (ISO 8601)
- `progress`: Completion percentage (0-100)
- `prd`: Path to PRD file (string)
- `github`: GitHub epic issue URL (string)
</epic_frontmatter>

<update_patterns>
## Safe Update Patterns

**Read existing frontmatter:**
```bash
# Extract frontmatter from file
sed -n '/^---$/,/^---$/p' file.md | sed '1d;$d'
```

**Update single field:**
```bash
# Update status field
sed -i 's/^status: .*/status: closed/' file.md
```

**Update with timestamp:**
```bash
# Get current datetime (see datetime-handling.md)
datetime=$(date -u +"%Y-%m-%dT%H:%M:%SZ")

# Update updated field
sed -i "s/^updated: .*/updated: $datetime/" file.md
```

**Preserve existing values:**
Always read current values before updating:
```bash
# Read current status
current_status=$(grep '^status:' file.md | cut -d' ' -f2)

# Only update if needed
if [ "$current_status" != "closed" ]; then
  sed -i 's/^status: .*/status: closed/' file.md
fi
```
</update_patterns>

<validation_rules>
## Validation Rules

**Before any update:**
1. File must exist
2. Frontmatter must be valid YAML
3. Required fields must be present
4. Field values must match expected types

**After any update:**
1. Frontmatter still valid YAML
2. All required fields present
3. No duplicate fields
4. Timestamps in ISO 8601 format
</validation_rules>

<common_operations>
## Common Operations

**Mark task in progress:**
```bash
datetime=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
sed -i "s/^status: .*/status: in-progress/" task.md
sed -i "s/^updated: .*/updated: $datetime/" task.md
```

**Mark task complete:**
```bash
datetime=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
sed -i "s/^status: .*/status: closed/" task.md
sed -i "s/^updated: .*/updated: $datetime/" task.md
```

**Update progress percentage:**
```bash
datetime=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
sed -i "s/^completion: .*/completion: $new_percentage/" progress.md
sed -i "s/^last_sync: .*/last_sync: $datetime/" progress.md
```
</common_operations>

<anti_patterns>
**Never:**
- Delete required fields
- Use placeholders like `{current_datetime}` - always use real timestamps
- Modify frontmatter without validating YAML syntax
- Update GitHub before updating local files
- Skip the `updated` field when making changes
</anti_patterns>
