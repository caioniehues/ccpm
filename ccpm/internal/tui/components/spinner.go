package components

import (
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/lipgloss"
)

func BrailleSpinner() spinner.Model {
	s := spinner.New()
	s.Spinner = spinner.Spinner{
		Frames: []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
		FPS:    time.Second / 12,
	}
	s.Style = lipgloss.NewStyle().Foreground(Electric)
	return s
}

func SyncSpinner() spinner.Model {
	s := spinner.New()
	s.Spinner = spinner.Spinner{
		Frames: []string{"⟳", "↻", "⟳", "↺"},
		FPS:    time.Second / 6,
	}
	s.Style = lipgloss.NewStyle().Foreground(Electric)
	return s
}

func PulseSpinner() spinner.Model {
	s := spinner.New()
	s.Spinner = spinner.Spinner{
		Frames: []string{"●", "◐", "○", "◑"},
		FPS:    time.Second / 4,
	}
	s.Style = lipgloss.NewStyle().Foreground(Amber)
	return s
}

func DotsSpinner() spinner.Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(Electric)
	return s
}
