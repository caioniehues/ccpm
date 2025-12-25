package parser

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCalculateCheckboxProgress(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected float64
	}{
		{
			name:     "half done",
			content:  "- [x] Done\n- [ ] Not done",
			expected: 0.5,
		},
		{
			name:     "all done",
			content:  "- [x] Done\n- [x] Also done",
			expected: 1.0,
		},
		{
			name:     "none done",
			content:  "- [ ] Not done\n- [ ] Also not done",
			expected: 0.0,
		},
		{
			name:     "no checkboxes",
			content:  "No checkboxes here",
			expected: 0.0,
		},
		{
			name:     "uppercase X",
			content:  "- [X] Case insensitive",
			expected: 1.0,
		},
		{
			name:     "mixed case",
			content:  "- [x] Lower\n- [X] Upper\n- [ ] Unchecked",
			expected: 2.0 / 3.0,
		},
		{
			name:     "many items",
			content:  "- [x] 1\n- [x] 2\n- [x] 3\n- [ ] 4\n- [ ] 5",
			expected: 0.6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalculateCheckboxProgress(tt.content)
			if result != tt.expected {
				t.Errorf("Expected %f, got %f", tt.expected, result)
			}
		})
	}
}

func TestParseTask(t *testing.T) {
	tmpDir := t.TempDir()
	taskPath := filepath.Join(tmpDir, "001.md")

	taskContent := `---
id: "001"
name: Test Task
status: in-progress
epic: test-epic
depends_on: ["000"]
effort: M
---

## Description

Test task content.

- [x] Step 1
- [ ] Step 2
`
	os.WriteFile(taskPath, []byte(taskContent), 0644)

	task, err := ParseTask(taskPath)
	if err != nil {
		t.Fatalf("ParseTask failed: %v", err)
	}

	if task.ID != "001" {
		t.Errorf("Expected ID '001', got '%s'", task.ID)
	}

	if task.Name != "Test Task" {
		t.Errorf("Expected name 'Test Task', got '%s'", task.Name)
	}

	if task.Status != "in-progress" {
		t.Errorf("Expected status 'in-progress', got '%s'", task.Status)
	}

	if task.DependsOn != "000" {
		t.Errorf("Expected depends_on '000', got '%s'", task.DependsOn)
	}

	if task.Progress != 0.5 {
		t.Errorf("Expected progress 0.5, got %f", task.Progress)
	}

	if task.Effort != "M" {
		t.Errorf("Expected effort 'M', got '%s'", task.Effort)
	}
}

func TestParseTasksInDir(t *testing.T) {
	tmpDir := t.TempDir()

	os.WriteFile(filepath.Join(tmpDir, "epic.md"), []byte(`---
name: test
---
`), 0644)

	os.WriteFile(filepath.Join(tmpDir, "002.md"), []byte(`---
id: "002"
name: Second
status: pending
---
`), 0644)

	os.WriteFile(filepath.Join(tmpDir, "001.md"), []byte(`---
id: "001"
name: First
status: completed
---
`), 0644)

	os.WriteFile(filepath.Join(tmpDir, "003.md"), []byte(`---
id: "003"
name: Third
status: in-progress
---
`), 0644)

	tasks, err := ParseTasksInDir(tmpDir)
	if err != nil {
		t.Fatalf("ParseTasksInDir failed: %v", err)
	}

	if len(tasks) != 3 {
		t.Fatalf("Expected 3 tasks, got %d", len(tasks))
	}

	if tasks[0].ID != "001" {
		t.Errorf("Expected first task ID '001', got '%s'", tasks[0].ID)
	}
	if tasks[1].ID != "002" {
		t.Errorf("Expected second task ID '002', got '%s'", tasks[1].ID)
	}
	if tasks[2].ID != "003" {
		t.Errorf("Expected third task ID '003', got '%s'", tasks[2].ID)
	}
}

func TestParseTasksInDirEmpty(t *testing.T) {
	tmpDir := t.TempDir()

	os.WriteFile(filepath.Join(tmpDir, "epic.md"), []byte(`---
name: test
---
`), 0644)

	tasks, err := ParseTasksInDir(tmpDir)
	if err != nil {
		t.Fatalf("ParseTasksInDir failed: %v", err)
	}

	if len(tasks) != 0 {
		t.Errorf("Expected 0 tasks, got %d", len(tasks))
	}
}

func TestParseTaskNoDependencies(t *testing.T) {
	tmpDir := t.TempDir()
	taskPath := filepath.Join(tmpDir, "standalone.md")

	taskContent := `---
id: "standalone"
name: Standalone Task
status: pending
---

No dependencies.
`
	os.WriteFile(taskPath, []byte(taskContent), 0644)

	task, err := ParseTask(taskPath)
	if err != nil {
		t.Fatalf("ParseTask failed: %v", err)
	}

	if task.DependsOn != "" {
		t.Errorf("Expected empty depends_on, got '%s'", task.DependsOn)
	}
}
