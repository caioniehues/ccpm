package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/automazeio/ccpm/internal/tui/views"
)

func (m Model) View() string {
	if !m.Ready {
		return m.Spinner.View() + " Initializing..."
	}

	if m.Layout.IsTooSmall() {
		return views.RenderTooSmall(m.Width, m.Height, MinWidth, MinHeight)
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

	mainView := AppFrame.Width(m.Width - 2).Height(m.Height - 2).Render(content)

	if m.Toasts.HasToasts() {
		toastView := m.Toasts.View()
		return m.overlayToasts(mainView, toastView)
	}

	return mainView
}

func (m Model) overlayToasts(main, toasts string) string {
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		main,
		lipgloss.Place(0, m.Height, lipgloss.Right, lipgloss.Top, toasts),
	)
}

func (m Model) viewDashboard() string {
	var sections []string

	sections = append(sections, m.renderHeader())
	sections = append(sections, "")
	sections = append(sections, m.renderEpicCard())
	sections = append(sections, "")

	if m.Layout.SideBySide {
		taskPanel := m.renderTaskList()
		activityPanel := m.renderActivityLog()
		sideBySide := lipgloss.JoinHorizontal(lipgloss.Top, taskPanel, "  ", activityPanel)
		sections = append(sections, sideBySide)
	} else {
		sections = append(sections, m.renderTaskList())
		if m.Layout.ActivityVisible {
			sections = append(sections, "")
			sections = append(sections, m.renderActivityLog())
		}
	}

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

	width := m.Layout.EpicCardWidth
	if width == 0 {
		width = m.Width - 10
	}

	e := m.ActiveEpic
	title := SectionHeader.Render("◆ ACTIVE EPIC")
	name := TitleStyle.Render(e.Name)
	desc := BodyStyle.Render(e.Description)

	m.Progress.Width = m.Layout.ProgressBarLen
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

	return ElevatedCard.Width(width).Render(content)
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
	label := "TASKS"
	if m.Layout.UseAbbreviatedLabels {
		label = "TSK"
	}
	title := LabelStyle.Render(label)
	nav := MutedStyle.Render("[j/k]")
	header := title + "  " + nav

	width := m.Layout.TaskListWidth
	if width == 0 {
		width = m.Width - 10
	}

	var lines []string
	lines = append(lines, header)
	lines = append(lines, MutedStyle.Render(strings.Repeat("─", width-6)))

	if len(m.Tasks) == 0 {
		lines = append(lines, MutedStyle.Render("  No tasks yet"))
	}

	maxItems := m.Layout.TaskItemCount
	if maxItems == 0 {
		maxItems = len(m.Tasks)
	}

	displayTasks := m.Tasks
	if len(displayTasks) > maxItems {
		displayTasks = displayTasks[:maxItems]
	}

	for i, task := range displayTasks {
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
			if m.Animations.PulseActive {
				iconStyle = TaskInProgress.Blink(true)
			}
		case "blocked":
			icon = IconBlocked
			iconStyle = TaskBlocked
		}

		cursor := "  "
		if i == m.ActiveTask {
			cursor = AccentStyle.Render(IconCursor + " ")
			nameStyle = AccentStyle
		}

		if m.Animations.IsFlashing(task.ID) {
			nameStyle = lipgloss.NewStyle().Background(Volt).Foreground(Void)
		}

		offset := ""
		if m.Animations.IsShaking(task.ID) {
			offset = "  "
		}

		line := fmt.Sprintf("%s%s%s  %s: %s",
			offset,
			cursor,
			iconStyle.Render(icon),
			MutedStyle.Render(task.ID),
			nameStyle.Render(task.Name),
		)
		lines = append(lines, line)
	}

	if len(m.Tasks) > maxItems {
		more := fmt.Sprintf("  ... +%d more", len(m.Tasks)-maxItems)
		lines = append(lines, MutedStyle.Render(more))
	}

	content := lipgloss.JoinVertical(lipgloss.Left, lines...)
	return SubtleCard.Width(width).Render(content)
}

func (m Model) renderActivityLog() string {
	label := "ACTIVITY"
	if m.Layout.UseAbbreviatedLabels {
		label = "LOG"
	}
	title := LabelStyle.Render(label)
	scroll := MutedStyle.Render("↑↓")
	header := title + "  " + scroll

	width := m.Layout.ActivityWidth
	if width == 0 {
		width = m.Width - 10
	}

	var lines []string
	lines = append(lines, header)
	lines = append(lines, MutedStyle.Render(strings.Repeat("─", width-6)))

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
	return SubtleCard.Width(width).Render(content)
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
