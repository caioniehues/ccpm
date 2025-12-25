package parser

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSplitFrontmatter(t *testing.T) {
	tests := []struct {
		name         string
		content      string
		expectedFM   string
		expectedBody string
		shouldHaveFM bool
	}{
		{
			name: "valid frontmatter",
			content: `---
name: test
status: pending
---

# Body content

More content here.`,
			expectedFM:   "name: test\nstatus: pending",
			expectedBody: "# Body content\n\nMore content here.",
			shouldHaveFM: true,
		},
		{
			name:         "no frontmatter",
			content:      "# Just a body\n\nNo frontmatter here.",
			expectedFM:   "",
			expectedBody: "# Just a body\n\nNo frontmatter here.",
			shouldHaveFM: false,
		},
		{
			name:         "empty content",
			content:      "",
			expectedFM:   "",
			expectedBody: "",
			shouldHaveFM: false,
		},
		{
			name: "incomplete frontmatter",
			content: `---
name: test
no closing delimiter`,
			expectedFM:   "",
			expectedBody: "---\nname: test\nno closing delimiter",
			shouldHaveFM: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fm, body := SplitFrontmatter(tt.content)

			if tt.shouldHaveFM && fm == "" {
				t.Error("Expected frontmatter but got empty")
			}
			if !tt.shouldHaveFM && fm != "" {
				t.Errorf("Expected no frontmatter but got: %s", fm)
			}
			if tt.shouldHaveFM && fm != tt.expectedFM {
				t.Errorf("Frontmatter mismatch:\nExpected: %q\nGot: %q", tt.expectedFM, fm)
			}
			_ = body
		})
	}
}

func TestExtractDescription(t *testing.T) {
	tests := []struct {
		name     string
		body     string
		expected string
	}{
		{
			name:     "with heading and description",
			body:     "# Title\n\nThis is the description.",
			expected: "This is the description.",
		},
		{
			name:     "description first",
			body:     "This is the description.\n\n# Title",
			expected: "This is the description.",
		},
		{
			name:     "only heading",
			body:     "# Just a heading",
			expected: "",
		},
		{
			name:     "empty body",
			body:     "",
			expected: "",
		},
		{
			name:     "multiple paragraphs",
			body:     "# Title\n\nFirst paragraph.\n\nSecond paragraph.",
			expected: "First paragraph.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExtractDescription(tt.body)
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestParseEpic(t *testing.T) {
	tmpDir := t.TempDir()
	epicDir := filepath.Join(tmpDir, "test-epic")
	os.MkdirAll(epicDir, 0755)

	epicContent := `---
name: test-epic
status: approved
created: 2025-12-23T10:00:00Z
prd: test-prd
---

# Test Epic

This is a test epic for unit testing.
`
	os.WriteFile(filepath.Join(epicDir, "epic.md"), []byte(epicContent), 0644)

	taskContent := `---
id: "001"
name: Test Task
status: completed
---

## Task Description

Test task content.

- [x] Acceptance criteria 1
- [ ] Acceptance criteria 2
`
	os.WriteFile(filepath.Join(epicDir, "001.md"), []byte(taskContent), 0644)

	epic, err := ParseEpic(epicDir)
	if err != nil {
		t.Fatalf("ParseEpic failed: %v", err)
	}

	if epic.Name != "test-epic" {
		t.Errorf("Expected name 'test-epic', got '%s'", epic.Name)
	}

	if epic.Status != "approved" {
		t.Errorf("Expected status 'approved', got '%s'", epic.Status)
	}

	if epic.TaskCount != 1 {
		t.Errorf("Expected 1 task, got %d", epic.TaskCount)
	}

	if epic.DoneCount != 1 {
		t.Errorf("Expected 1 done, got %d", epic.DoneCount)
	}

	if epic.Progress != 1.0 {
		t.Errorf("Expected progress 1.0, got %f", epic.Progress)
	}
}

func TestParseEpicNoTasks(t *testing.T) {
	tmpDir := t.TempDir()
	epicDir := filepath.Join(tmpDir, "empty-epic")
	os.MkdirAll(epicDir, 0755)

	epicContent := `---
name: empty-epic
status: pending
---

# Empty Epic
`
	os.WriteFile(filepath.Join(epicDir, "epic.md"), []byte(epicContent), 0644)

	epic, err := ParseEpic(epicDir)
	if err != nil {
		t.Fatalf("ParseEpic failed: %v", err)
	}

	if epic.TaskCount != 0 {
		t.Errorf("Expected 0 tasks, got %d", epic.TaskCount)
	}

	if epic.Progress != 0.0 {
		t.Errorf("Expected progress 0.0, got %f", epic.Progress)
	}
}

func TestLoadAllEpics(t *testing.T) {
	tmpDir := t.TempDir()
	claudeDir := filepath.Join(tmpDir, ".claude", "epics")
	os.MkdirAll(claudeDir, 0755)

	epic1Dir := filepath.Join(claudeDir, "epic-one")
	os.MkdirAll(epic1Dir, 0755)
	os.WriteFile(filepath.Join(epic1Dir, "epic.md"), []byte(`---
name: epic-one
status: approved
---
# Epic One
`), 0644)

	epic2Dir := filepath.Join(claudeDir, "epic-two")
	os.MkdirAll(epic2Dir, 0755)
	os.WriteFile(filepath.Join(epic2Dir, "epic.md"), []byte(`---
name: epic-two
status: pending
---
# Epic Two
`), 0644)

	epics, err := LoadAllEpics(tmpDir)
	if err != nil {
		t.Fatalf("LoadAllEpics failed: %v", err)
	}

	if len(epics) != 2 {
		t.Errorf("Expected 2 epics, got %d", len(epics))
	}
}

func TestLoadAllEpicsEmpty(t *testing.T) {
	tmpDir := t.TempDir()

	epics, err := LoadAllEpics(tmpDir)
	if err != nil {
		t.Fatalf("LoadAllEpics failed: %v", err)
	}

	if len(epics) != 0 {
		t.Errorf("Expected 0 epics, got %d", len(epics))
	}
}
