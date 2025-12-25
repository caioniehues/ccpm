package parser

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParsePRD(t *testing.T) {
	tmpDir := t.TempDir()
	prdPath := filepath.Join(tmpDir, "test-prd.md")

	prdContent := `---
name: test-feature
status: approved
created: 2025-12-23T10:00:00Z
approved_at: 2025-12-23T12:00:00Z
---

# Test Feature PRD

## Overview

This is a test PRD for unit testing.
`
	os.WriteFile(prdPath, []byte(prdContent), 0644)

	prd, err := ParsePRD(prdPath)
	if err != nil {
		t.Fatalf("ParsePRD failed: %v", err)
	}

	if prd.Name != "test-feature" {
		t.Errorf("Expected name 'test-feature', got '%s'", prd.Name)
	}

	if prd.Status != "approved" {
		t.Errorf("Expected status 'approved', got '%s'", prd.Status)
	}

	if prd.CreatedAt != "2025-12-23T10:00:00Z" {
		t.Errorf("Expected created '2025-12-23T10:00:00Z', got '%s'", prd.CreatedAt)
	}

	if prd.ApprovedAt != "2025-12-23T12:00:00Z" {
		t.Errorf("Expected approved_at '2025-12-23T12:00:00Z', got '%s'", prd.ApprovedAt)
	}
}

func TestParsePRDNotFound(t *testing.T) {
	_, err := ParsePRD("/nonexistent/path/prd.md")
	if err == nil {
		t.Error("Expected error for nonexistent file")
	}
}

func TestParsePRDInvalidYAML(t *testing.T) {
	tmpDir := t.TempDir()
	prdPath := filepath.Join(tmpDir, "invalid.md")

	invalidContent := `---
name: [invalid yaml
---

# Content
`
	os.WriteFile(prdPath, []byte(invalidContent), 0644)

	_, err := ParsePRD(prdPath)
	if err == nil {
		t.Error("Expected error for invalid YAML")
	}
}

func TestLoadAllPRDs(t *testing.T) {
	tmpDir := t.TempDir()
	prdsDir := filepath.Join(tmpDir, ".claude", "prds")
	os.MkdirAll(prdsDir, 0755)

	os.WriteFile(filepath.Join(prdsDir, "prd-one.md"), []byte(`---
name: prd-one
status: approved
---
# PRD One
`), 0644)

	os.WriteFile(filepath.Join(prdsDir, "prd-two.md"), []byte(`---
name: prd-two
status: draft
---
# PRD Two
`), 0644)

	os.WriteFile(filepath.Join(prdsDir, "not-a-prd.txt"), []byte("ignored"), 0644)

	prds, err := LoadAllPRDs(tmpDir)
	if err != nil {
		t.Fatalf("LoadAllPRDs failed: %v", err)
	}

	if len(prds) != 2 {
		t.Errorf("Expected 2 PRDs, got %d", len(prds))
	}
}

func TestLoadAllPRDsEmpty(t *testing.T) {
	tmpDir := t.TempDir()

	prds, err := LoadAllPRDs(tmpDir)
	if err != nil {
		t.Fatalf("LoadAllPRDs failed: %v", err)
	}

	if len(prds) != 0 {
		t.Errorf("Expected 0 PRDs, got %d", len(prds))
	}
}

func TestLoadAllPRDsSkipsInvalid(t *testing.T) {
	tmpDir := t.TempDir()
	prdsDir := filepath.Join(tmpDir, ".claude", "prds")
	os.MkdirAll(prdsDir, 0755)

	os.WriteFile(filepath.Join(prdsDir, "valid.md"), []byte(`---
name: valid
status: approved
---
# Valid
`), 0644)

	os.WriteFile(filepath.Join(prdsDir, "invalid.md"), []byte(`---
name: [broken yaml
---
`), 0644)

	prds, err := LoadAllPRDs(tmpDir)
	if err != nil {
		t.Fatalf("LoadAllPRDs failed: %v", err)
	}

	if len(prds) != 1 {
		t.Errorf("Expected 1 valid PRD, got %d", len(prds))
	}
}

func TestLoadAllPRDsSkipsDirectories(t *testing.T) {
	tmpDir := t.TempDir()
	prdsDir := filepath.Join(tmpDir, ".claude", "prds")
	os.MkdirAll(prdsDir, 0755)

	os.WriteFile(filepath.Join(prdsDir, "valid.md"), []byte(`---
name: valid
status: approved
---
# Valid
`), 0644)

	os.MkdirAll(filepath.Join(prdsDir, "subdir"), 0755)

	prds, err := LoadAllPRDs(tmpDir)
	if err != nil {
		t.Fatalf("LoadAllPRDs failed: %v", err)
	}

	if len(prds) != 1 {
		t.Errorf("Expected 1 PRD, got %d", len(prds))
	}
}
