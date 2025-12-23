package parser

import (
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type PRDFrontmatter struct {
	Name       string `yaml:"name"`
	Status     string `yaml:"status"`
	Created    string `yaml:"created"`
	ApprovedAt string `yaml:"approved_at"`
}

type PRD struct {
	Name       string
	Status     string
	Content    string
	CreatedAt  string
	ApprovedAt string
}

func ParsePRD(path string) (*PRD, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	frontmatter, body := SplitFrontmatter(string(content))

	var fm PRDFrontmatter
	if err := yaml.Unmarshal([]byte(frontmatter), &fm); err != nil {
		return nil, err
	}

	return &PRD{
		Name:       fm.Name,
		Status:     fm.Status,
		Content:    body,
		CreatedAt:  fm.Created,
		ApprovedAt: fm.ApprovedAt,
	}, nil
}

func LoadAllPRDs(baseDir string) ([]PRD, error) {
	prdsDir := filepath.Join(baseDir, ".claude", "prds")

	entries, err := os.ReadDir(prdsDir)
	if err != nil {
		return []PRD{}, nil
	}

	var prds []PRD
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}

		prd, err := ParsePRD(filepath.Join(prdsDir, entry.Name()))
		if err != nil {
			continue
		}
		prds = append(prds, *prd)
	}

	return prds, nil
}
