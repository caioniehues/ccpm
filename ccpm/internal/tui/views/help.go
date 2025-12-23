package views

import (
	"github.com/charmbracelet/lipgloss"
)

func RenderHelp(width, height int) string {
	helpContent := `
KEYBOARD SHORTCUTS
═══════════════════════════════════════════════════════════

NAVIGATION                       ACTIONS
──────────                       ───────
↑ k      Move up                 ⏎       Select/Confirm
↓ j      Move down               esc     Back/Cancel
← h      Previous section        /       Filter/Search
→ l      Next section            ?       Toggle help
g        Go to top
G        Go to bottom            VIEWS
tab      Next panel              ─────
                                 e       Epic details
WORKFLOW                         t       Task details
────────                         p       View PRD
w        Launch wizard           a       Activity log
s        Sync to GitHub
c        Mark complete           SYSTEM
b        Mark blocked            ──────
n        Next task               r       Refresh
                                 q       Quit
`

	card := lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder()).
		BorderForeground(Electric).
		Padding(1, 2).
		Background(Charcoal)

	dismissHint := MutedStyle.Render("Press any key to close")

	modal := lipgloss.JoinVertical(
		lipgloss.Center,
		card.Render(AccentStyle.Render(helpContent)),
		"",
		dismissHint,
	)

	return lipgloss.Place(
		width, height,
		lipgloss.Center, lipgloss.Center,
		modal,
		lipgloss.WithWhitespaceChars("░"),
		lipgloss.WithWhitespaceForeground(Graphite),
	)
}
