package parser

import (
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

type TaskFrontmatter struct {
	ID        string   `yaml:"id"`
	Name      string   `yaml:"name"`
	Status    string   `yaml:"status"`
	Epic      string   `yaml:"epic"`
	DependsOn []string `yaml:"depends_on"`
	Effort    string   `yaml:"effort"`
	Created   string   `yaml:"created"`
	Started   string   `yaml:"started_at"`
	Completed string   `yaml:"completed_at"`
}

type Task struct {
	ID          string
	Name        string
	Status      string
	Epic        string
	Description string
	Progress    float64
	DependsOn   string
	BlockedBy   string
	StartedAt   string
	CompletedAt string
	Effort      string
}

func ParseTask(path string) (*Task, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	frontmatter, body := SplitFrontmatter(string(content))

	var fm TaskFrontmatter
	if err := yaml.Unmarshal([]byte(frontmatter), &fm); err != nil {
		return nil, err
	}

	progress := CalculateCheckboxProgress(body)

	dependsOn := ""
	if len(fm.DependsOn) > 0 {
		dependsOn = strings.Join(fm.DependsOn, ", ")
	}

	return &Task{
		ID:          fm.ID,
		Name:        fm.Name,
		Status:      fm.Status,
		Epic:        fm.Epic,
		Description: body,
		Progress:    progress,
		DependsOn:   dependsOn,
		StartedAt:   fm.Started,
		CompletedAt: fm.Completed,
		Effort:      fm.Effort,
	}, nil
}

func ParseTasksInDir(epicDir string) ([]Task, error) {
	entries, err := os.ReadDir(epicDir)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if name == "epic.md" || !strings.HasSuffix(name, ".md") {
			continue
		}

		task, err := ParseTask(filepath.Join(epicDir, name))
		if err != nil {
			continue
		}
		tasks = append(tasks, *task)
	}

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].ID < tasks[j].ID
	})

	return tasks, nil
}

func CalculateCheckboxProgress(content string) float64 {
	checked := strings.Count(content, "- [x]") + strings.Count(content, "- [X]")
	unchecked := strings.Count(content, "- [ ]")
	total := checked + unchecked

	if total == 0 {
		return 0
	}
	return float64(checked) / float64(total)
}
