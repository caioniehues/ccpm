package components

import (
	"fmt"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/lipgloss"
)

var (
	Void     = lipgloss.Color("#0D0D0D")
	Charcoal = lipgloss.Color("#1A1A2E")
	Graphite = lipgloss.Color("#2D2D44")
	Slate    = lipgloss.Color("#4A4A6A")
	Silver   = lipgloss.Color("#8888AA")
	Pearl    = lipgloss.Color("#E8E8F0")
	Electric = lipgloss.Color("#00D4FF")
	Plasma   = lipgloss.Color("#FF006E")
	Volt     = lipgloss.Color("#ADFF02")
	Amber    = lipgloss.Color("#FFB800")
	Lavender = lipgloss.Color("#B388FF")

	AccentStyle  = lipgloss.NewStyle().Foreground(Electric).Bold(true)
	TitleStyle   = lipgloss.NewStyle().Foreground(Pearl).Bold(true)
	BodyStyle    = lipgloss.NewStyle().Foreground(Pearl)
	MutedStyle   = lipgloss.NewStyle().Foreground(Slate)
	SuccessStyle = lipgloss.NewStyle().Foreground(Volt).Bold(true)
	ElevatedCard = lipgloss.NewStyle().Border(lipgloss.DoubleBorder()).BorderForeground(Electric).Padding(1, 2)
)

type EpicCard struct {
	Name         string
	Description  string
	Progress     float64
	TasksDone    int
	TasksTotal   int
	PRDApproved  bool
	EpicApproved bool
	TasksReady   bool
	Synced       bool
	Width        int
}

func (e EpicCard) View() string {
	title := AccentStyle.Render("◆ ACTIVE EPIC")
	name := TitleStyle.Render(e.Name)
	desc := BodyStyle.Render(e.Description)

	prog := progress.New(progress.WithScaledGradient(string(Electric), string(Volt)))
	progBar := prog.ViewAs(e.Progress)
	progText := MutedStyle.Render(fmt.Sprintf("%d/%d tasks", e.TasksDone, e.TasksTotal))

	phases := e.renderPhases()

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		name,
		desc,
		"",
		progBar,
		progText,
		"",
		phases,
	)

	return ElevatedCard.Width(e.Width).Render(content)
}

func (e EpicCard) renderPhases() string {
	prd := MutedStyle.Render("○ PRD")
	if e.PRDApproved {
		prd = SuccessStyle.Render("✓ PRD")
	}

	epic := MutedStyle.Render("○ Epic")
	if e.EpicApproved {
		epic = SuccessStyle.Render("✓ Epic")
	}

	tasks := MutedStyle.Render("○ Tasks")
	if e.TasksReady {
		tasks = lipgloss.NewStyle().Foreground(Amber).Render("◐ Tasks")
	}
	if e.TasksDone == e.TasksTotal && e.TasksTotal > 0 {
		tasks = SuccessStyle.Render("✓ Tasks")
	}

	sync := MutedStyle.Render("○ Sync")
	if e.Synced {
		sync = SuccessStyle.Render("✓ Sync")
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, prd, "  ", epic, "  ", tasks, "  ", sync)
}
