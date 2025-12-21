---
description: Test the task reference update logic used in epic-sync
allowed-tools: Bash, Read, Write
---

<objective>
Test the task reference update logic that updates task IDs after GitHub issue creation.
</objective>

<process>
**Usage**: `/pm:test-reference-update`

**1. Create Test Files**

Create test task files with references:
```bash
mkdir -p /tmp/test-refs
cd /tmp/test-refs

# Create task 001 with conflicts_with references
cat > 001.md << 'EOF'
---
name: Task One
status: open
depends_on: []
conflicts_with: [002, 003]
---
# Task One
EOF

# Create task 002 with depends_on reference
cat > 002.md << 'EOF'
---
name: Task Two
status: open
depends_on: [001]
conflicts_with: [003]
---
# Task Two
EOF

# Create task 003 with multiple depends_on
cat > 003.md << 'EOF'
---
name: Task Three
status: open
depends_on: [001, 002]
conflicts_with: []
---
# Task Three
EOF
```

**2. Create Mappings**

Simulate the issue creation mappings:
```bash
# Simulate task -> issue number mapping
cat > /tmp/task-mapping.txt << 'EOF'
001.md:42
002.md:43
003.md:44
EOF

# Create old -> new ID mapping
> /tmp/id-mapping.txt
while IFS=: read -r task_file task_number; do
  old_num=$(basename "$task_file" .md)
  echo "$old_num:$task_number" >> /tmp/id-mapping.txt
done < /tmp/task-mapping.txt
```

**3. Update References**

Process each file and update references using sed:
```bash
while IFS=: read -r task_file task_number; do
  content=$(cat "$task_file")
  while IFS=: read -r old_num new_num; do
    content=$(echo "$content" | sed "s/\b$old_num\b/$new_num/g")
  done < /tmp/id-mapping.txt
  new_name="${task_number}.md"
  echo "$content" > "$new_name"
done < /tmp/task-mapping.txt
```

**4. Verify Results**

Check that references were updated correctly:
```bash
echo "=== Final Results ==="
for file in 42.md 43.md 44.md; do
  echo "File: $file"
  grep -E "name:|depends_on:|conflicts_with:" "$file"
done
```

Expected output:
- 42.md should have conflicts_with: [43, 44]
- 43.md should have depends_on: [42] and conflicts_with: [44]
- 44.md should have depends_on: [42, 43]

**5. Cleanup**

```bash
cd -
rm -rf /tmp/test-refs
rm -f /tmp/task-mapping.txt /tmp/id-mapping.txt
echo "✅ Test complete and cleaned up"
```
</process>

<success_criteria>
- Test files created with cross-references
- ID mapping generated correctly
- References updated from local IDs to GitHub issue numbers
- Cleanup completed without errors
</success_criteria>
