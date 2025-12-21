# Datetime Handling

<overview>
CCPM uses ISO 8601 format (UTC) for all timestamps. Always use real datetimes from system commands, never placeholders.
</overview>

<critical_rule>
**NEVER use placeholders like `{current_datetime}` or `{timestamp}` in actual files.**

Always execute the datetime command to get the real current time.
</critical_rule>

<getting_current_datetime>
**Get Current Datetime**

**Standard command:**
```bash
date -u +"%Y-%m-%dT%H:%M:%SZ"
```

**Example output:**
```
2024-01-15T10:30:45Z
```

**In scripts:**
```bash
#!/bin/bash
datetime=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
echo "Current time: $datetime"

# Use in sed
sed -i "s/^updated: .*/updated: $datetime/" file.md
```
</getting_current_datetime>

<iso_8601_format>
**ISO 8601 Format**

**Structure:**
```
YYYY-MM-DDTHH:MM:SSZ
```

**Components:**
- `YYYY`: 4-digit year
- `MM`: 2-digit month (01-12)
- `DD`: 2-digit day (01-31)
- `T`: Separator between date and time
- `HH`: 2-digit hour (00-23) in UTC
- `MM`: 2-digit minute (00-59)
- `SS`: 2-digit second (00-59)
- `Z`: UTC timezone indicator

**Examples:**
- `2024-01-15T10:30:00Z` - Valid
- `2024-1-15T10:30:00Z` - Invalid (month not zero-padded)
- `2024-01-15 10:30:00` - Invalid (missing T and Z)
- `{current_datetime}` - Invalid (placeholder, not real time)
</iso_8601_format>

<usage_patterns>
**Common Usage Patterns**

**Update task file:**
```bash
# Get current time
datetime=$(date -u +"%Y-%m-%dT%H:%M:%SZ")

# Update frontmatter
sed -i "s/^updated: .*/updated: $datetime/" .claude/epics/epic-name/task-001.md
```

**Create progress file:**
```bash
# Get current time
datetime=$(date -u +"%Y-%m-%dT%H:%M:%SZ")

# Create file with frontmatter
cat > progress.md << EOF
---
issue: 123
started: $datetime
last_sync: $datetime
completion: 0
---

# Progress Tracking

Started work at $datetime
EOF
```

**Add timestamp to commit:**
```bash
datetime=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
git commit -m "Issue #123: Complete feature

Completed at: $datetime"
```
</usage_patterns>

<validation>
**Validate Datetime Format**

**Check if datetime is valid ISO 8601:**
```bash
datetime="2024-01-15T10:30:00Z"

# Validate format
if [[ $datetime =~ ^[0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}Z$ ]]; then
  echo "Valid ISO 8601 datetime"
else
  echo "Invalid datetime format"
fi
```

**Detect placeholders:**
```bash
# Check for placeholder patterns
if grep -q '{.*datetime.*}' file.md; then
  echo "❌ ERROR: File contains datetime placeholders!"
  echo "Run: date -u +'%Y-%m-%dT%H:%M:%SZ' to get real time"
  exit 1
fi
```
</validation>

<timezone_handling>
**Timezone Handling**

**Always use UTC:**
CCPM standardizes on UTC (indicated by `Z` suffix) to avoid timezone confusion.

**Convert local time to UTC:**
```bash
# Get UTC time (recommended)
date -u +"%Y-%m-%dT%H:%M:%SZ"

# If you have local time and need UTC
TZ=UTC date +"%Y-%m-%dT%H:%M:%SZ"
```

**Display local time (for user output only):**
```bash
# Get UTC time for storage
utc_time=$(date -u +"%Y-%m-%dT%H:%M:%SZ")

# Display in local timezone (for user)
local_time=$(date +"%Y-%m-%d %H:%M:%S %Z")

echo "Synced at: $local_time (stored as: $utc_time)"
```
</timezone_handling>

<common_mistakes>
**Common Mistakes**

**❌ Using placeholders:**
```bash
# WRONG
sed -i 's/^updated: .*/updated: {current_datetime}/' file.md
```

**✅ Using real datetime:**
```bash
# CORRECT
datetime=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
sed -i "s/^updated: .*/updated: $datetime/" file.md
```

**❌ Forgetting to execute command:**
```yaml
---
updated: $(date -u +"%Y-%m-%dT%H:%M:%SZ")
---
```

**✅ Execute and substitute:**
```bash
datetime=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
cat > file.md << EOF
---
updated: $datetime
---
EOF
```

**❌ Using local timezone:**
```bash
# WRONG
date +"%Y-%m-%dT%H:%M:%S"  # Missing Z, local time
```

**✅ Using UTC:**
```bash
# CORRECT
date -u +"%Y-%m-%dT%H:%M:%SZ"  # UTC with Z
```
</common_mistakes>

<success_criteria>
Datetime handling is correct when:
- All timestamps use ISO 8601 format with Z suffix
- No placeholder strings like `{datetime}` in files
- All times are UTC (not local timezone)
- Timestamps generated from `date -u` command
- Validation confirms format before writing to files
</success_criteria>
