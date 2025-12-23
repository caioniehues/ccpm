package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type Header struct {
	Title      string
	Branch     string
	Commit     string
	Width      int
	ShowBranch bool
}

func NewHeader(width int) Header {
	return Header{
		Title:      "CCPM Dashboard",
		Width:      width,
		ShowBranch: true,
	}
}

func (h Header) View() string {
	title := AccentStyle.Render(h.Title)

	var rightSide string
	if h.ShowBranch && h.Branch != "" {
		rightSide = MutedStyle.Render("Branch: " + h.Branch)
		if h.Commit != "" {
			rightSide += MutedStyle.Render(" @ " + h.Commit[:7])
		}
	}

	padding := h.Width - lipgloss.Width(title) - lipgloss.Width(rightSide)
	if padding < 0 {
		padding = 0
	}

	return title + strings.Repeat(" ", padding) + rightSide
}

func (h Header) SetBranch(branch string) Header {
	h.Branch = branch
	return h
}

func (h Header) SetCommit(commit string) Header {
	h.Commit = commit
	return h
}
