package parser

import (
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type EpicFrontmatter struct {
	Name    string `yaml:"name"`
	Status  string `yaml:"status"`
	Created string `yaml:"created"`
	Updated string `yaml:"updated"`
	PRD     string `yaml:"prd"`
	GitHub  string `yaml:"github"`
}

type Epic struct {
	Name        string
	Status      string
	Description string
	Content     string
	PRDName     string
	Branch      string
	CreatedAt   string
	UpdatedAt   string
	Progress    float64
	TaskCount   int
	DoneCount   int
	Tasks       []Task
}

func ParseEpic(epicDir string) (*Epic, error) {
	epicPath := filepath.Join(epicDir, "epic.md")

	content, err := os.ReadFile(epicPath)
	if err != nil {
		return nil, err
	}

	frontmatter, body := SplitFrontmatter(string(content))

	var fm EpicFrontmatter
	if err := yaml.Unmarshal([]byte(frontmatter), &fm); err != nil {
		return nil, err
	}

	tasks, err := ParseTasksInDir(epicDir)
	if err != nil {
		tasks = []Task{}
	}

	doneCount := 0
	for _, t := range tasks {
		if t.Status == "completed" {
			doneCount++
		}
	}

	progress := 0.0
	if len(tasks) > 0 {
		progress = float64(doneCount) / float64(len(tasks))
	}

	return &Epic{
		Name:        fm.Name,
		Status:      fm.Status,
		Description: ExtractDescription(body),
		Content:     body,
		PRDName:     fm.PRD,
		CreatedAt:   fm.Created,
		UpdatedAt:   fm.Updated,
		Progress:    progress,
		TaskCount:   len(tasks),
		DoneCount:   doneCount,
		Tasks:       tasks,
	}, nil
}

func LoadAllEpics(baseDir string) ([]Epic, error) {
	epicsDir := filepath.Join(baseDir, ".claude", "epics")

	entries, err := os.ReadDir(epicsDir)
	if err != nil {
		return []Epic{}, nil
	}

	var epics []Epic
	for _, entry := range entries {
		if entry.IsDir() {
			epic, err := ParseEpic(filepath.Join(epicsDir, entry.Name()))
			if err != nil {
				continue
			}
			epics = append(epics, *epic)
		}
	}

	return epics, nil
}

func SplitFrontmatter(content string) (string, string) {
	if !strings.HasPrefix(content, "---") {
		return "", content
	}

	parts := strings.SplitN(content[3:], "---", 2)
	if len(parts) != 2 {
		return "", content
	}

	return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
}

func ExtractDescription(body string) string {
	lines := strings.Split(body, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" && !strings.HasPrefix(line, "#") {
			return line
		}
	}
	return ""
}
