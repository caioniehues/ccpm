package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	HelpKeyStyle   = lipgloss.NewStyle().Foreground(Electric).Bold(true)
	HelpDescStyle  = lipgloss.NewStyle().Foreground(Silver)
	StatusBarStyle = lipgloss.NewStyle().Background(Charcoal).Foreground(Silver).Padding(0, 1)
)

type KeyBinding struct {
	Key  string
	Desc string
}

type Footer struct {
	Bindings []KeyBinding
	Width    int
}

func NewFooter(width int) Footer {
	return Footer{
		Width: width,
		Bindings: []KeyBinding{
			{Key: "e", Desc: "Epic"},
			{Key: "t", Desc: "Tasks"},
			{Key: "p", Desc: "PRD"},
			{Key: "s", Desc: "Sync"},
			{Key: "w", Desc: "Wizard"},
			{Key: "?", Desc: "Help"},
			{Key: "q", Desc: "Quit"},
		},
	}
}

func (f Footer) View() string {
	var keys []string
	for _, b := range f.Bindings {
		keys = append(keys, HelpKeyStyle.Render("["+b.Key+"]")+HelpDescStyle.Render(b.Desc))
	}
	content := strings.Join(keys, "  ")
	return StatusBarStyle.Width(f.Width).Render(content)
}

func (f Footer) SetBindings(bindings []KeyBinding) Footer {
	f.Bindings = bindings
	return f
}
