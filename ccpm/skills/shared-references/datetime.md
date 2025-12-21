<overview>
Standard patterns for getting and formatting datetime values in CCPM operations. All datetime operations MUST use real system time, never estimates or placeholders.
</overview>

<get_current_datetime>
Use the `date` command to get the current ISO 8601 formatted datetime:

```bash
# Primary method (Linux/Mac)
date -u +"%Y-%m-%dT%H:%M:%SZ"

# Alternative if primary fails
date --iso-8601=seconds

# For Windows (if using PowerShell)
Get-Date -Format "yyyy-MM-ddTHH:mm:ssZ"

# Cross-platform fallback chain
date -u +"%Y-%m-%dT%H:%M:%SZ" 2>/dev/null || \
date +"%Y-%m-%dT%H:%M:%SZ" 2>/dev/null || \
python3 -c "from datetime import datetime; print(datetime.utcnow().strftime('%Y-%m-%dT%H:%M:%SZ'))" 2>/dev/null || \
python -c "from datetime import datetime; print(datetime.utcnow().strftime('%Y-%m-%dT%H:%M:%SZ'))"
```

Store the output before using in files:
```bash
CURRENT_DATE=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
```
</get_current_datetime>

<required_format>
All dates in CCPM MUST use ISO 8601 format with UTC timezone:

- **Format**: `YYYY-MM-DDTHH:MM:SSZ`
- **Example**: `2024-01-15T14:30:45Z`
- **Timezone**: Always UTC (the `Z` suffix)
</required_format>

<frontmatter_usage>
**Creating new files:**
- Set both `created` and `updated` to current datetime
- PRDs, Epics, Tasks all follow this pattern

**Updating existing files:**
- Preserve original `created` field
- Update only the `updated` field with current datetime
- For sync operations, update `last_sync` with real datetime

```yaml
---
name: example
created: 2024-01-10T09:15:30Z  # Keep original
updated: 2024-01-15T14:30:45Z  # Use current time
---
```
</frontmatter_usage>

<critical_rules>
- **Never use placeholders** like `[Current ISO date/time]` or `YYYY-MM-DD`
- **Never estimate dates** - always get actual system time
- **Always use UTC** for consistency across timezones
- **Run date command first** before writing any file with timestamps
</critical_rules>

<affected_operations>
This rule applies to all operations that write timestamps:
- PRD creation and updates
- Epic creation, sync, and progress tracking
- Task/Issue creation and status changes
- Progress file updates
- Any log or timestamp entry
</affected_operations>
