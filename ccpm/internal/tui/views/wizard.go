package views

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type WizardData struct {
	Step        int
	EpicName    string
	PRDContent  string
	EpicContent string
	TaskCount   int
	Revising    bool
	Feedback    string
	Width       int
	Height      int
}

func RenderWizard(d WizardData) string {
	header := RenderWizardHeader(d.Step)
	stepper := RenderWizardStepper(d.Step)

	var content string
	switch d.Step {
	case 0:
		content = RenderWizardPRD(d)
	case 1:
		content = RenderWizardEpic(d)
	case 2:
		content = RenderWizardTasks(d)
	case 3:
		content = RenderWizardSync(d)
	}

	footer := MutedStyle.Render("[↑/↓]navigate  [⏎]select  [esc]cancel wizard")

	frame := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		BorderForeground(Graphite).
		Padding(1, 2).
		Width(d.Width - 2)

	inner := lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		stepper,
		"",
		content,
	)

	return frame.Render(inner) + "\n" + footer
}

func RenderWizardHeader(step int) string {
	stepText := fmt.Sprintf("Step %d/4", step+1)

	titles := []string{"PRD Review", "Epic Review", "Tasks Review", "GitHub Sync"}
	title := titles[step]

	card := lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder()).
		BorderForeground(Electric).
		Padding(1, 2)

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		TitleStyle.Render("EPIC WIZARD"),
		AccentStyle.Render(strings.Repeat("═", 50)),
		"",
		LabelStyle.Render(stepText+" - "+title),
	)

	return card.Render(content)
}

func RenderWizardStepper(currentStep int) string {
	steps := []string{"PRD", "Epic", "Tasks", "Sync"}
	states := []string{"○", "○", "○", "○"}
	labels := []string{"Pending", "Pending", "Pending", "Pending"}

	for i := 0; i < currentStep; i++ {
		states[i] = "■"
		labels[i] = "✓ Done"
	}
	states[currentStep] = "◐"
	labels[currentStep] = "Review"

	stepper := fmt.Sprintf(
		"    [%s]────────[%s]────────[%s]────────[%s]",
		SuccessStyle.Render(states[0]),
		SuccessStyle.Render(states[1]),
		WarningStyle.Render(states[2]),
		MutedStyle.Render(states[3]),
	)

	labelLine := fmt.Sprintf(
		"     %s          %s         %s         %s",
		steps[0], steps[1], steps[2], steps[3],
	)

	statusLine := fmt.Sprintf(
		"    %s       %s      %s     %s",
		labels[0], labels[1], labels[2], labels[3],
	)

	return lipgloss.JoinVertical(lipgloss.Left, stepper, labelLine, MutedStyle.Render(statusLine))
}

func RenderWizardPRD(d WizardData) string {
	title := LabelStyle.Render("REVIEW: PRD DOCUMENT")

	docCard := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(Graphite).
		Padding(1, 2).
		Width(d.Width - 16).
		Height(10)

	doc := docCard.Render(lipgloss.JoinVertical(lipgloss.Left, title, "", d.PRDContent))

	if d.Revising {
		return lipgloss.JoinVertical(lipgloss.Left, doc, "", RenderRevisionInput(d.Feedback))
	}

	approval := RenderApprovalPrompt("Do you approve this PRD?")

	return lipgloss.JoinVertical(lipgloss.Left, doc, "", approval)
}

func RenderWizardEpic(d WizardData) string {
	title := LabelStyle.Render("REVIEW: EPIC DOCUMENT")

	docCard := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(Graphite).
		Padding(1, 2).
		Width(d.Width - 16).
		Height(10)

	doc := docCard.Render(lipgloss.JoinVertical(lipgloss.Left, title, "", d.EpicContent))

	approval := RenderApprovalPrompt("Do you approve this Epic?")

	return lipgloss.JoinVertical(lipgloss.Left, doc, "", approval)
}

func RenderWizardTasks(d WizardData) string {
	title := LabelStyle.Render(fmt.Sprintf("REVIEW: %d TASKS CREATED", d.TaskCount))

	taskPreview := `
  001.md: Task one (effort: M)
  002.md: Task two (effort: S)
  003.md: Task three (effort: L)
  ...`

	docCard := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(Graphite).
		Padding(1, 2).
		Width(d.Width - 16)

	doc := docCard.Render(lipgloss.JoinVertical(lipgloss.Left, title, taskPreview))

	approval := RenderApprovalPrompt("Do you approve these tasks?")

	return lipgloss.JoinVertical(lipgloss.Left, doc, "", approval)
}

func RenderWizardSync(d WizardData) string {
	summary := fmt.Sprintf(`
◆ %s

PRD:    .claude/prds/%s.md          ✅ APPROVED
Epic:   .claude/epics/%s/epic.md    ✅ APPROVED
Tasks:  %d tasks ready for execution  ✅ APPROVED

Ready to begin work!`, d.EpicName, d.EpicName, d.EpicName, d.TaskCount)

	summaryCard := lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder()).
		BorderForeground(Electric).
		Padding(1, 2).
		Width(d.Width - 16)

	choice := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		BorderForeground(Plasma).
		Background(Charcoal).
		Padding(1, 2)

	choiceContent := `? Sync epic to GitHub?

  › ⟳ Yes, create GitHub issues now
    ○ No, work locally only
    ⏭ Skip for now, I'll sync later`

	tip := MutedStyle.Render(`
Tip: Working locally is great for offline development.
     Sync to GitHub anytime with [s] from the dashboard.`)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		summaryCard.Render(SuccessStyle.Render(summary)),
		"",
		choice.Render(choiceContent),
		"",
		tip,
	)
}

func RenderApprovalPrompt(question string) string {
	card := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		BorderForeground(Plasma).
		Background(Charcoal).
		Padding(1, 2)

	content := fmt.Sprintf(`? %s

  › ✓ Yes, proceed
    ✗ No, I need to make changes`, question)

	return card.Render(content)
}

func RenderRevisionInput(feedback string) string {
	card := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		BorderForeground(Plasma).
		Background(Charcoal).
		Padding(1, 2)

	content := fmt.Sprintf(`What changes are needed?

┌───────────────────────────────────────────────────────┐
│  %s█                                                  │
└───────────────────────────────────────────────────────┘

Characters: %d/500

Tip: Be specific about what needs to change.`, feedback, len(feedback))

	return card.Render(content)
}
