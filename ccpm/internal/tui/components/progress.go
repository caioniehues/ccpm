package components

import (
	"fmt"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/lipgloss"
)

type ProgressBar struct {
	Progress progress.Model
	Value    float64
	Label    string
	Width    int
}

func NewProgressBar(width int) ProgressBar {
	p := progress.New(
		progress.WithScaledGradient(string(Electric), string(Volt)),
		progress.WithWidth(width-4),
	)
	return ProgressBar{
		Progress: p,
		Width:    width,
	}
}

func (p ProgressBar) View() string {
	bar := p.Progress.ViewAs(p.Value)
	pct := fmt.Sprintf("%.0f%%", p.Value*100)

	if p.Label != "" {
		return lipgloss.JoinVertical(
			lipgloss.Left,
			LabelStyle.Render(p.Label),
			bar+" "+MutedStyle.Render(pct),
		)
	}
	return bar + " " + MutedStyle.Render(pct)
}

func (p ProgressBar) SetValue(v float64) ProgressBar {
	p.Value = v
	return p
}

func (p ProgressBar) SetLabel(l string) ProgressBar {
	p.Label = l
	return p
}
