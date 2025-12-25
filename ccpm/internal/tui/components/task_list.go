package components

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TaskItem struct {
	ID       string
	Name     string
	Status   string
	Progress float64
}

func (t TaskItem) Title() string       { return t.Name }
func (t TaskItem) Description() string { return t.ID }
func (t TaskItem) FilterValue() string { return t.Name }

type TaskDelegate struct {
	Width int
}

func (d TaskDelegate) Height() int                             { return 1 }
func (d TaskDelegate) Spacing() int                            { return 0 }
func (d TaskDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }

func (d TaskDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	task, ok := item.(TaskItem)
	if !ok {
		return
	}

	icon := "▢"
	iconStyle := MutedStyle
	switch task.Status {
	case "completed":
		icon = "▣"
		iconStyle = SuccessStyle
	case "in-progress":
		icon = "▶"
		iconStyle = lipgloss.NewStyle().Foreground(Amber).Bold(true)
	case "blocked":
		icon = "⊘"
		iconStyle = lipgloss.NewStyle().Foreground(Plasma).Bold(true)
	}

	cursor := "  "
	nameStyle := BodyStyle
	if index == m.Index() {
		cursor = AccentStyle.Render("❯ ")
		nameStyle = AccentStyle
	}

	line := fmt.Sprintf("%s%s  %s: %s",
		cursor,
		iconStyle.Render(icon),
		MutedStyle.Render(task.ID),
		nameStyle.Render(task.Name),
	)

	fmt.Fprint(w, line)
}

func NewTaskList(tasks []TaskItem, width, height int) list.Model {
	items := make([]list.Item, len(tasks))
	for i, t := range tasks {
		items[i] = t
	}

	delegate := TaskDelegate{Width: width}
	l := list.New(items, delegate, width, height)
	l.SetShowTitle(false)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)

	return l
}
