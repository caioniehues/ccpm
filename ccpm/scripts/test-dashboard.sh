#!/bin/bash
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"

echo "=== CCPM TUI Dashboard Visual Test ==="
echo ""

cd "$PROJECT_ROOT"

if [ ! -f "./ccpm-tui" ]; then
    echo "Building ccpm-tui..."
    go build -o ccpm-tui ./cmd/ccpm-tui
fi

TEST_DIR=".claude/epics/visual-test"
mkdir -p "$TEST_DIR"

cat > "$TEST_DIR/epic.md" << 'EOF'
---
name: visual-test
status: approved
created: 2025-12-23T10:00:00Z
---

# Visual Test Epic

Testing dashboard visualization with sample data.
EOF

cat > "$TEST_DIR/001.md" << 'EOF'
---
id: "001"
name: Completed Task
status: completed
---

A task that is complete.

- [x] First step
- [x] Second step
EOF

cat > "$TEST_DIR/002.md" << 'EOF'
---
id: "002"
name: In Progress Task
status: in-progress
---

A task currently in progress.

- [x] Started
- [ ] Working on it
- [ ] Not done yet
EOF

cat > "$TEST_DIR/003.md" << 'EOF'
---
id: "003"
name: Pending Task
status: pending
depends_on: ["002"]
---

A task waiting to be started.

- [ ] Step one
- [ ] Step two
EOF

cat > "$TEST_DIR/004.md" << 'EOF'
---
id: "004"
name: Blocked Task
status: blocked
---

A blocked task.
EOF

echo "Test data created in $TEST_DIR"
echo ""
echo "Starting dashboard..."
echo "Press 'q' to quit when done testing."
echo ""

./ccpm-tui

rm -rf "$TEST_DIR"
echo ""
echo "Test data cleaned up."
