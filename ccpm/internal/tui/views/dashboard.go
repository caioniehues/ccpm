package views

import (
	"fmt"
	"strings"

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
)

var (
	TitleStyle    = lipgloss.NewStyle().Foreground(Pearl).Bold(true)
	AccentStyle   = lipgloss.NewStyle().Foreground(Electric).Bold(true)
	MutedStyle    = lipgloss.NewStyle().Foreground(Slate)
	LabelStyle    = lipgloss.NewStyle().Foreground(Silver).Bold(true)
	SuccessStyle  = lipgloss.NewStyle().Foreground(Volt).Bold(true)
	WarningStyle  = lipgloss.NewStyle().Foreground(Amber).Bold(true)
	ErrorStyle    = lipgloss.NewStyle().Foreground(Plasma).Bold(true)
	HelpKeyStyle  = lipgloss.NewStyle().Foreground(Electric).Bold(true)
	HelpDescStyle = lipgloss.NewStyle().Foreground(Silver)
)

const Logo = `
 ██████╗ ██████╗██████╗ ███╗   ███╗
██╔════╝██╔════╝██╔══██╗████╗ ████║
██║     ██║     ██████╔╝██╔████╔██║
██║     ██║     ██╔═══╝ ██║╚██╔╝██║
╚██████╗╚██████╗██║     ██║ ╚═╝ ██║
 ╚═════╝ ╚═════╝╚═╝     ╚═╝     ╚═╝`

const EmptyState = `
╭─────────────────────────────────────╮
│                                     │
│        ◇ No epics yet               │
│                                     │
│     Start by creating your first    │
│     PRD with the wizard             │
│                                     │
│     Press [w] to launch wizard      │
│                                     │
╰─────────────────────────────────────╯`

type DashboardData struct {
	EpicName     string
	EpicDesc     string
	EpicProgress float64
	TasksDone    int
	TasksTotal   int
	EpicStatus   string
	Branch       string
	Tasks        []TaskData
	Activity     []ActivityData
	Width        int
	Height       int
	ActiveTask   int
	HasEpic      bool
}

type TaskData struct {
	ID       string
	Name     string
	Status   string
	Progress float64
}

type ActivityData struct {
	Time    string
	Message string
	Type    string
}

func RenderDashboard(d DashboardData) string {
	if !d.HasEpic {
		return RenderEmptyState(d.Width, d.Height)
	}

	var sections []string

	sections = append(sections, RenderHeader(d.Branch, d.Width))
	sections = append(sections, "")
	sections = append(sections, RenderEpicCard(d))
	sections = append(sections, "")
	sections = append(sections, RenderTaskList(d.Tasks, d.ActiveTask, d.Width))
	sections = append(sections, "")
	sections = append(sections, RenderActivityLog(d.Activity, d.Width))
	sections = append(sections, "")
	sections = append(sections, RenderFooter(d.Width))

	return lipgloss.JoinVertical(lipgloss.Left, sections...)
}

func RenderHeader(branch string, width int) string {
	logo := AccentStyle.Render(Logo)
	branchText := ""
	if branch != "" {
		branchText = MutedStyle.Render("Branch: " + branch)
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, logo, "    ", branchText)
}

func RenderEpicCard(d DashboardData) string {
	title := AccentStyle.Render("◆ ACTIVE EPIC")
	name := TitleStyle.Render(d.EpicName)
	desc := lipgloss.NewStyle().Foreground(Pearl).Render(d.EpicDesc)

	progWidth := d.Width - 20
	if progWidth < 20 {
		progWidth = 20
	}
	progBar := RenderProgressBar(d.EpicProgress, progWidth)
	progText := MutedStyle.Render(fmt.Sprintf("%d/%d tasks", d.TasksDone, d.TasksTotal))

	phases := RenderPhases(d.EpicStatus)

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

	card := lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder()).
		BorderForeground(Electric).
		Padding(1, 2).
		Width(d.Width - 10)

	return card.Render(content)
}

func RenderProgressBar(progress float64, width int) string {
	filled := int(progress * float64(width))
	empty := width - filled

	bar := AccentStyle.Render(strings.Repeat("█", filled)) +
		MutedStyle.Render(strings.Repeat("░", empty))

	return bar
}

func RenderPhases(status string) string {
	prd := MutedStyle.Render("○ PRD")
	epic := MutedStyle.Render("○ Epic")
	tasks := MutedStyle.Render("○ Tasks")
	sync := MutedStyle.Render("○ Sync")

	if status == "approved" || status == "approved-for-work" || status == "completed" {
		prd = SuccessStyle.Render("✓ PRD")
		epic = SuccessStyle.Render("✓ Epic")
	}
	if status == "approved-for-work" || status == "completed" {
		tasks = WarningStyle.Render("◐ Tasks")
	}
	if status == "completed" {
		tasks = SuccessStyle.Render("✓ Tasks")
		sync = SuccessStyle.Render("✓ Sync")
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, prd, "  ", epic, "  ", tasks, "  ", sync)
}

func RenderTaskList(tasks []TaskData, active int, width int) string {
	title := LabelStyle.Render("TASKS")
	nav := MutedStyle.Render("[j/k nav]")
	header := title + "  " + nav

	var lines []string
	lines = append(lines, header)
	lines = append(lines, MutedStyle.Render(strings.Repeat("─", width-14)))

	if len(tasks) == 0 {
		lines = append(lines, MutedStyle.Render("  No tasks yet"))
	}

	for i, task := range tasks {
		icon := "▢"
		iconStyle := MutedStyle
		nameStyle := lipgloss.NewStyle().Foreground(Pearl)

		switch task.Status {
		case "completed":
			icon = "▣"
			iconStyle = SuccessStyle
		case "in-progress":
			icon = "▶"
			iconStyle = WarningStyle
		case "blocked":
			icon = "⊘"
			iconStyle = ErrorStyle
		}

		cursor := "  "
		if i == active {
			cursor = AccentStyle.Render("❯ ")
			nameStyle = AccentStyle
		}

		line := fmt.Sprintf("%s%s  %s: %s",
			cursor,
			iconStyle.Render(icon),
			MutedStyle.Render(task.ID),
			nameStyle.Render(task.Name),
		)
		lines = append(lines, line)
	}

	content := lipgloss.JoinVertical(lipgloss.Left, lines...)

	card := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(Graphite).
		Padding(1, 2).
		Width(width - 10)

	return card.Render(content)
}

func RenderActivityLog(activity []ActivityData, width int) string {
	title := LabelStyle.Render("ACTIVITY")
	scroll := MutedStyle.Render("↑↓")
	header := title + "  " + scroll

	var lines []string
	lines = append(lines, header)
	lines = append(lines, MutedStyle.Render(strings.Repeat("─", width-14)))

	if len(activity) == 0 {
		lines = append(lines, MutedStyle.Render("  No activity yet"))
	}

	for _, entry := range activity {
		icon := "ℹ"
		style := lipgloss.NewStyle().Foreground(Lavender)
		switch entry.Type {
		case "success":
			icon = "✓"
			style = SuccessStyle
		case "warning":
			icon = "⚠"
			style = WarningStyle
		case "error":
			icon = "✗"
			style = ErrorStyle
		}

		line := fmt.Sprintf("  %s  %s %s",
			MutedStyle.Render(entry.Time),
			style.Render(icon),
			entry.Message,
		)
		lines = append(lines, line)
	}

	content := lipgloss.JoinVertical(lipgloss.Left, lines...)

	card := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(Graphite).
		Padding(1, 2).
		Width(width - 10)

	return card.Render(content)
}

func RenderFooter(width int) string {
	keys := []string{
		HelpKeyStyle.Render("[e]") + HelpDescStyle.Render("Epic"),
		HelpKeyStyle.Render("[t]") + HelpDescStyle.Render("Tasks"),
		HelpKeyStyle.Render("[p]") + HelpDescStyle.Render("PRD"),
		HelpKeyStyle.Render("[s]") + HelpDescStyle.Render("Sync"),
		HelpKeyStyle.Render("[w]") + HelpDescStyle.Render("Wizard"),
		HelpKeyStyle.Render("[?]") + HelpDescStyle.Render("Help"),
		HelpKeyStyle.Render("[q]") + HelpDescStyle.Render("Quit"),
	}

	content := strings.Join(keys, "  ")

	bar := lipgloss.NewStyle().
		Background(Charcoal).
		Foreground(Silver).
		Padding(0, 1).
		Width(width)

	return bar.Render(content)
}

func RenderEmptyState(width, height int) string {
	message := MutedStyle.Render(EmptyState)

	return lipgloss.Place(
		width, height,
		lipgloss.Center, lipgloss.Center,
		message,
	)
}
