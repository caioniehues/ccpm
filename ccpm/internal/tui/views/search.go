package views

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type SearchData struct {
	Query        string
	TypeFilters  []string
	StatusFilter string
	Results      []SearchResult
	Selected     int
	Width        int
	Height       int
}

type SearchResult struct {
	Type        string
	Name        string
	Description string
	Icon        string
}

func RenderSearch(d SearchData) string {
	card := lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder()).
		BorderForeground(Electric).
		Padding(1, 2).
		Width(60).
		Background(Charcoal)

	title := AccentStyle.Render("🔍 SEARCH")
	divider := AccentStyle.Render(strings.Repeat("═", 56))

	input := RenderSearchInput(d.Query)
	filters := RenderSearchFilters(d.TypeFilters, d.StatusFilter)
	results := RenderSearchResults(d.Results, d.Selected)

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		divider,
		"",
		input,
		"",
		filters,
		"",
		results,
	)

	footer := MutedStyle.Render("[↑/↓]navigate  [⏎]select  [tab]toggle filter  [esc]close")

	modal := lipgloss.JoinVertical(
		lipgloss.Center,
		card.Render(content),
		"",
		footer,
	)

	return lipgloss.Place(
		d.Width, d.Height,
		lipgloss.Center, lipgloss.Center,
		modal,
		lipgloss.WithWhitespaceChars("░"),
		lipgloss.WithWhitespaceForeground(Graphite),
	)
}

func RenderSearchInput(query string) string {
	input := fmt.Sprintf("┌%s┐\n│ %s█%s │\n└%s┘",
		strings.Repeat("─", 52),
		query,
		strings.Repeat(" ", 50-len(query)),
		strings.Repeat("─", 52),
	)

	return lipgloss.NewStyle().Foreground(Pearl).Render(input)
}

func RenderSearchFilters(types []string, status string) string {
	typeLabel := LabelStyle.Render("Type: ")
	typeFilters := ""
	allTypes := []string{"Epics", "Tasks", "PRDs"}
	for _, t := range allTypes {
		active := false
		for _, enabled := range types {
			if enabled == t {
				active = true
				break
			}
		}
		if active {
			typeFilters += AccentStyle.Render("["+t+"]") + " "
		} else {
			typeFilters += MutedStyle.Render("["+t+"]") + " "
		}
	}

	statusLabel := LabelStyle.Render("Status: ")
	allStatuses := []string{"All", "Pending", "In Progress", "Done"}
	statusFilters := ""
	for _, s := range allStatuses {
		if s == status {
			statusFilters += AccentStyle.Render("["+s+"]") + " "
		} else {
			statusFilters += MutedStyle.Render("["+s+"]") + " "
		}
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		typeLabel+typeFilters,
		statusLabel+statusFilters,
	)
}

func RenderSearchResults(results []SearchResult, selected int) string {
	if len(results) == 0 {
		return MutedStyle.Render("No results found")
	}

	title := LabelStyle.Render(fmt.Sprintf("Results (%d)", len(results)))

	var lines []string
	lines = append(lines, title)
	lines = append(lines, MutedStyle.Render(strings.Repeat("─", 52)))

	for i, r := range results {
		cursor := "  "
		nameStyle := lipgloss.NewStyle().Foreground(Pearl)
		if i == selected {
			cursor = AccentStyle.Render("❯ ")
			nameStyle = AccentStyle
		}

		typeIcon := "◆"
		switch r.Type {
		case "Epic":
			typeIcon = "◆"
		case "Task":
			typeIcon = "▢"
		case "PRD":
			typeIcon = "◈"
		}

		line := fmt.Sprintf("%s%s %s %s",
			cursor,
			MutedStyle.Render(typeIcon),
			nameStyle.Render(r.Name),
			MutedStyle.Render("- "+r.Description),
		)
		lines = append(lines, line)
	}

	return lipgloss.JoinVertical(lipgloss.Left, lines...)
}
