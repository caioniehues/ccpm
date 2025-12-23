package components

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
)

var (
	LabelStyle   = lipgloss.NewStyle().Foreground(Silver).Bold(true)
	InfoStyle    = lipgloss.NewStyle().Foreground(Lavender)
	WarningStyle = lipgloss.NewStyle().Foreground(Amber).Bold(true)
	ErrorStyle   = lipgloss.NewStyle().Foreground(Plasma).Bold(true)
	SubtleCard   = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(Graphite).Padding(1, 2)
)

type ActivityEntry struct {
	Time    string
	Message string
	Type    string
}

type ActivityLog struct {
	Entries  []ActivityEntry
	Viewport viewport.Model
	Width    int
	Height   int
}

func NewActivityLog(width, height int) ActivityLog {
	vp := viewport.New(width-4, height-2)
	return ActivityLog{
		Viewport: vp,
		Width:    width,
		Height:   height,
	}
}

func (a *ActivityLog) SetEntries(entries []ActivityEntry) {
	a.Entries = entries
	a.Viewport.SetContent(a.renderEntries())
	a.Viewport.GotoBottom()
}

func (a *ActivityLog) AddEntry(entry ActivityEntry) {
	a.Entries = append(a.Entries, entry)
	a.Viewport.SetContent(a.renderEntries())
	a.Viewport.GotoBottom()
}

func (a ActivityLog) renderEntries() string {
	var lines []string
	for _, entry := range a.Entries {
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

		line := fmt.Sprintf("%s  %s %s",
			MutedStyle.Render(entry.Time),
			style.Render(icon),
			entry.Message,
		)
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
}

func (a ActivityLog) View() string {
	title := LabelStyle.Render("ACTIVITY")
	scroll := MutedStyle.Render("↑↓")
	header := lipgloss.JoinHorizontal(lipgloss.Top, title, "  ", scroll)

	content := lipgloss.JoinVertical(lipgloss.Left, header, a.Viewport.View())
	return SubtleCard.Width(a.Width).Render(content)
}
