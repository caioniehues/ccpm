package views

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func RenderTooSmall(width, height, minWidth, minHeight int) string {
	box := `
╭─────────────────────────────────╮
│                                 │
│   Terminal too small            │
│                                 │
│   Minimum: %3d × %2d             │
│   Current: %3d × %2d             │
│                                 │
│   Please resize your terminal   │
│                                 │
╰─────────────────────────────────╯`

	message := fmt.Sprintf(box, minWidth, minHeight, width, height)

	style := lipgloss.NewStyle().
		Foreground(Plasma).
		Bold(true)

	return lipgloss.Place(
		width, height,
		lipgloss.Center, lipgloss.Center,
		style.Render(message),
	)
}

func RenderResizeHint(width, height int) string {
	hint := fmt.Sprintf("↔ %d×%d (need 80×24)", width, height)
	return lipgloss.NewStyle().
		Foreground(Plasma).
		Render(hint)
}
