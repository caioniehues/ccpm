package views

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type EpicDetailData struct {
	Name        string
	Description string
	Status      string
	CreatedAt   string
	Branch      string
	PRDName     string
	Content     string
	Progress    float64
	TasksDone   int
	TasksTotal  int
	Width       int
	Height      int
}

func RenderEpicDetail(d EpicDetailData) string {
	breadcrumb := MutedStyle.Render("← Dashboard › Epic Details")

	title := AccentStyle.Render("◆ " + d.Name)
	divider := AccentStyle.Render(strings.Repeat("═", d.Width-8))

	meta := lipgloss.JoinVertical(
		lipgloss.Left,
		LabelStyle.Render("STATUS")+"        "+d.Status,
		LabelStyle.Render("CREATED")+"       "+d.CreatedAt,
		LabelStyle.Render("BRANCH")+"        "+d.Branch,
		LabelStyle.Render("SOURCE PRD")+"    ◈ "+d.PRDName,
	)

	docTitle := LabelStyle.Render("EPIC DOCUMENT")
	docContent := lipgloss.NewStyle().Foreground(Pearl).Render(d.Content)

	docCard := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(Graphite).
		Padding(1, 2).
		Width(d.Width - 12)

	doc := docCard.Render(lipgloss.JoinVertical(lipgloss.Left, docTitle, "", docContent))

	phases := RenderPhaseStatus(d.Status, d.TasksDone, d.TasksTotal)

	footer := RenderEpicDetailFooter(d.Width)

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		breadcrumb,
		"",
		title,
		divider,
		"",
		meta,
		"",
		doc,
		"",
		phases,
	)

	frame := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		BorderForeground(Graphite).
		Padding(1, 2).
		Width(d.Width - 2).
		Height(d.Height - 4)

	return frame.Render(content) + "\n" + footer
}

func RenderPhaseStatus(status string, done, total int) string {
	prdState := "■"
	epicState := "■"
	taskState := "◐"

	prdLabel := "✓ Approved"
	epicLabel := "✓ Approved"
	taskLabel := fmt.Sprintf("%d/%d done", done, total)

	if status == "completed" {
		taskState = "■"
	}

	box := fmt.Sprintf(`
┌──────────────────────────────────────────────────────────┐
│                                                          │
│  [%s] PRD Created   ──▶  [%s] Epic Approved  ──▶  [%s] Tasks│
│                                                          │
│       %s            %s          %s │
│                                                          │
└──────────────────────────────────────────────────────────┘`,
		prdState, epicState, taskState, prdLabel, epicLabel, taskLabel)

	return MutedStyle.Render(box)
}

func RenderEpicDetailFooter(width int) string {
	keys := []string{
		HelpKeyStyle.Render("[←]") + HelpDescStyle.Render("back"),
		HelpKeyStyle.Render("[t]") + HelpDescStyle.Render("asks"),
		HelpKeyStyle.Render("[p]") + HelpDescStyle.Render("rd"),
		HelpKeyStyle.Render("[s]") + HelpDescStyle.Render("ync"),
		HelpKeyStyle.Render("[m]") + HelpDescStyle.Render("ark complete"),
	}

	content := strings.Join(keys, "  ")

	bar := lipgloss.NewStyle().
		Background(Charcoal).
		Foreground(Silver).
		Padding(0, 1).
		Width(width)

	return bar.Render(content)
}
