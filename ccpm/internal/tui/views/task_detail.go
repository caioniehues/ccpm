package views

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type TaskDetailData struct {
	ID          string
	Name        string
	Status      string
	Progress    float64
	Description string
	DependsOn   string
	BlockedBy   string
	StartedAt   string
	EpicName    string
	Width       int
	Height      int
}

func RenderTaskDetail(d TaskDetailData) string {
	breadcrumb := MutedStyle.Render("← Dashboard › Epic › Task Details")

	icon := "▶"
	iconStyle := WarningStyle
	switch d.Status {
	case "completed":
		icon = "▣"
		iconStyle = SuccessStyle
	case "blocked":
		icon = "⊘"
		iconStyle = ErrorStyle
	case "pending":
		icon = "▢"
		iconStyle = MutedStyle
	}

	title := iconStyle.Render(icon) + " " + AccentStyle.Render(d.ID+": "+d.Name)
	divider := AccentStyle.Render(strings.Repeat("═", d.Width-8))

	progBar := RenderProgressBar(d.Progress, d.Width-30)
	progPct := fmt.Sprintf(" %d%% %s", int(d.Progress*100), d.Status)

	statusLine := lipgloss.JoinHorizontal(
		lipgloss.Top,
		LabelStyle.Render("STATUS")+"        ",
		progBar,
		MutedStyle.Render(progPct),
	)

	meta := lipgloss.JoinVertical(
		lipgloss.Left,
		statusLine,
		LabelStyle.Render("STARTED")+"       "+d.StartedAt,
		LabelStyle.Render("DEPENDS ON")+"    "+d.DependsOn,
		LabelStyle.Render("BLOCKED BY")+"    "+d.BlockedBy,
	)

	docTitle := LabelStyle.Render("TASK DESCRIPTION")
	docContent := lipgloss.NewStyle().Foreground(Pearl).Render(d.Description)

	docCard := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(Graphite).
		Padding(1, 2).
		Width(d.Width - 12)

	doc := docCard.Render(lipgloss.JoinVertical(lipgloss.Left, docTitle, "", docContent))

	trace := RenderRequirementsTrace(d.Width)

	footer := RenderTaskDetailFooter(d.Width)

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
		trace,
	)

	frame := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		BorderForeground(Graphite).
		Padding(1, 2).
		Width(d.Width - 2).
		Height(d.Height - 4)

	return frame.Render(content) + "\n" + footer
}

func RenderRequirementsTrace(width int) string {
	content := `REQUIREMENTS TRACE

Linked to:
  ◈ PRD → Requirements section
  ◆ Epic → Implementation goals

Leverages:
  • Existing codebase patterns
  • Project conventions`

	card := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(Graphite).
		Padding(1, 2).
		Width(width - 12)

	return card.Render(MutedStyle.Render(content))
}

func RenderTaskDetailFooter(width int) string {
	keys := []string{
		HelpKeyStyle.Render("[←]") + HelpDescStyle.Render("back"),
		HelpKeyStyle.Render("[c]") + HelpDescStyle.Render("omplete"),
		HelpKeyStyle.Render("[b]") + HelpDescStyle.Render("lock"),
		HelpKeyStyle.Render("[n]") + HelpDescStyle.Render("ext task"),
		HelpKeyStyle.Render("[g]") + HelpDescStyle.Render("ithub"),
	}

	content := strings.Join(keys, "  ")

	bar := lipgloss.NewStyle().
		Background(Charcoal).
		Foreground(Silver).
		Padding(0, 1).
		Width(width)

	return bar.Render(content)
}
