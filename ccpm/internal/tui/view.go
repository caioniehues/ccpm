package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	if !m.Ready {
		return m.Spinner.View() + " Initializing..."
	}

	var content string
	switch m.CurrentView {
	case ViewDashboard:
		content = m.viewDashboard()
	case ViewEpicDetail:
		content = m.viewEpicDetail()
	case ViewTaskDetail:
		content = m.viewTaskDetail()
	case ViewPRDDetail:
		content = m.viewPRDDetail()
	case ViewWizard:
		content = m.viewWizard()
	case ViewEpicSelector:
		content = m.viewEpicSelector()
	case ViewHelp:
		content = m.viewHelp()
	case ViewSettings:
		content = m.viewSettings()
	case ViewSearch:
		content = m.viewSearch()
	default:
		content = m.viewDashboard()
	}

	return AppFrame.Width(m.Width - 2).Height(m.Height - 2).Render(content)
}

func (m Model) viewDashboard() string {
	var sections []string

	sections = append(sections, m.renderHeader())
	sections = append(sections, "")
	sections = append(sections, m.renderEpicCard())
	sections = append(sections, "")
	sections = append(sections, m.renderTaskList())
	sections = append(sections, "")
	sections = append(sections, m.renderActivityLog())
	sections = append(sections, "")
	sections = append(sections, m.renderFooter())

	return lipgloss.JoinVertical(lipgloss.Left, sections...)
}

func (m Model) renderHeader() string {
	title := AccentStyle.Render("CCPM Dashboard")
	branch := ""
	if m.ActiveEpic != nil && m.ActiveEpic.Branch != "" {
		branch = MutedStyle.Render("Branch: " + m.ActiveEpic.Branch)
	}

	width := m.Width - 6
	padding := width - lipgloss.Width(title) - lipgloss.Width(branch)
	if padding < 0 {
		padding = 0
	}

	return title + strings.Repeat(" ", padding) + branch
}

func (m Model) renderEpicCard() string {
	if m.ActiveEpic == nil {
		return SubtleCard.Render(MutedStyle.Render("No active epic. Run /pm:epic-wizard to create one."))
	}

	e := m.ActiveEpic
	title := SectionHeader.Render("◆ ACTIVE EPIC")
	name := TitleStyle.Render(e.Name)
	desc := BodyStyle.Render(e.Description)

	progBar := m.Progress.ViewAs(e.Progress)
	progText := MutedStyle.Render(fmt.Sprintf("%d/%d tasks", e.DoneCount, e.TaskCount))

	phases := m.renderPhases(e)

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

	return ElevatedCard.Width(m.Width - 10).Render(content)
}

func (m Model) renderPhases(e *Epic) string {
	prd := MutedStyle.Render("○ PRD")
	epic := MutedStyle.Render("○ Epic")
	tasks := MutedStyle.Render("○ Tasks")
	sync := MutedStyle.Render("○ Sync")

	if e.Status == "approved" || e.Status == "approved-for-work" || e.Status == "completed" {
		prd = SuccessStyle.Render("✓ PRD")
		epic = SuccessStyle.Render("✓ Epic")
	}
	if e.Status == "approved-for-work" || e.Status == "completed" {
		tasks = WarningStyle.Render("◐ Tasks")
	}
	if e.Status == "completed" {
		tasks = SuccessStyle.Render("✓ Tasks")
		sync = SuccessStyle.Render("✓ Sync")
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, prd, "  ", epic, "  ", tasks, "  ", sync)
}

func (m Model) renderTaskList() string {
	title := LabelStyle.Render("TASKS")
	nav := MutedStyle.Render("[j/k nav]")
	header := title + "  " + nav

	var lines []string
	lines = append(lines, header)
	lines = append(lines, MutedStyle.Render(strings.Repeat("─", m.Width-14)))

	if len(m.Tasks) == 0 {
		lines = append(lines, MutedStyle.Render("  No tasks yet"))
	}

	for i, task := range m.Tasks {
		icon := IconPending
		iconStyle := TaskPending
		nameStyle := BodyStyle

		switch task.Status {
		case "completed":
			icon = IconCompleted
			iconStyle = TaskCompleted
		case "in-progress":
			icon = IconInProgress
			iconStyle = TaskInProgress
		case "blocked":
			icon = IconBlocked
			iconStyle = TaskBlocked
		}

		cursor := "  "
		if i == m.ActiveTask {
			cursor = AccentStyle.Render(IconCursor + " ")
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
	return SubtleCard.Width(m.Width - 10).Render(content)
}

func (m Model) renderActivityLog() string {
	title := LabelStyle.Render("ACTIVITY")
	scroll := MutedStyle.Render("↑↓")
	header := title + "  " + scroll

	var lines []string
	lines = append(lines, header)
	lines = append(lines, MutedStyle.Render(strings.Repeat("─", m.Width-14)))

	if len(m.ActivityLog) == 0 {
		lines = append(lines, MutedStyle.Render("  No activity yet"))
	}

	for _, entry := range m.ActivityLog {
		icon := "ℹ"
		style := InfoStyle
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
	return SubtleCard.Width(m.Width - 10).Render(content)
}

func (m Model) renderFooter() string {
	keys := []string{
		HelpKeyStyle.Render("[e]") + HelpDescStyle.Render("Epic"),
		HelpKeyStyle.Render("[t]") + HelpDescStyle.Render("Tasks"),
		HelpKeyStyle.Render("[p]") + HelpDescStyle.Render("PRD"),
		HelpKeyStyle.Render("[s]") + HelpDescStyle.Render("Sync"),
		HelpKeyStyle.Render("[w]") + HelpDescStyle.Render("Wizard"),
		HelpKeyStyle.Render("[?]") + HelpDescStyle.Render("Help"),
		HelpKeyStyle.Render("[q]") + HelpDescStyle.Render("Quit"),
	}
	return StatusBar.Render(strings.Join(keys, "  "))
}

func (m Model) viewEpicDetail() string {
	if m.ActiveEpic == nil {
		return "No epic selected"
	}
	return ElevatedCard.Render(
		TitleStyle.Render("Epic: "+m.ActiveEpic.Name) + "\n\n" +
			BodyStyle.Render(m.ActiveEpic.Description) + "\n\n" +
			MutedStyle.Render("Press ESC to go back"),
	)
}

func (m Model) viewTaskDetail() string {
	if len(m.Tasks) == 0 || m.ActiveTask >= len(m.Tasks) {
		return "No task selected"
	}
	task := m.Tasks[m.ActiveTask]
	return ElevatedCard.Render(
		TitleStyle.Render("Task: "+task.Name) + "\n\n" +
			LabelStyle.Render("Status: ") + task.Status + "\n" +
			LabelStyle.Render("Epic: ") + task.Epic + "\n\n" +
			BodyStyle.Render(task.Description) + "\n\n" +
			MutedStyle.Render("Press ESC to go back"),
	)
}

func (m Model) viewPRDDetail() string {
	return ElevatedCard.Render(
		TitleStyle.Render("PRD Detail") + "\n\n" +
			MutedStyle.Render("Press ESC to go back"),
	)
}

func (m Model) viewWizard() string {
	steps := []string{"PRD", "Epic", "Tasks", "Sync"}
	var stepLine string
	for i, step := range steps {
		if i == m.WizardStep {
			stepLine += AccentStyle.Render("● " + step)
		} else if i < m.WizardStep {
			stepLine += SuccessStyle.Render("✓ " + step)
		} else {
			stepLine += MutedStyle.Render("○ " + step)
		}
		if i < len(steps)-1 {
			stepLine += MutedStyle.Render(" ─── ")
		}
	}

	return ElevatedCard.Render(
		TitleStyle.Render("Epic Wizard") + "\n\n" +
			stepLine + "\n\n" +
			MutedStyle.Render("Press ESC to exit wizard"),
	)
}

func (m Model) viewEpicSelector() string {
	var lines []string
	lines = append(lines, TitleStyle.Render("Select Epic"))
	lines = append(lines, "")

	for i, epic := range m.Epics {
		cursor := "  "
		style := BodyStyle
		if m.ActiveEpic != nil && epic.Name == m.ActiveEpic.Name {
			cursor = AccentStyle.Render(IconCursor + " ")
			style = AccentStyle
		}
		lines = append(lines, cursor+style.Render(epic.Name))
		_ = i
	}

	return ElevatedCard.Render(lipgloss.JoinVertical(lipgloss.Left, lines...))
}

func (m Model) viewHelp() string {
	return ElevatedCard.Render(
		TitleStyle.Render("Keyboard Shortcuts") + "\n\n" +
			m.Help.View(Keys),
	)
}

func (m Model) viewSettings() string {
	return ElevatedCard.Render(
		TitleStyle.Render("Settings") + "\n\n" +
			MutedStyle.Render("Press ESC to go back"),
	)
}

func (m Model) viewSearch() string {
	return ElevatedCard.Render(
		TitleStyle.Render("Search") + "\n\n" +
			MutedStyle.Render("Type to search... (Press ESC to cancel)"),
	)
}
