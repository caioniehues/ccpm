package tui

import "github.com/charmbracelet/lipgloss"

var (
	Void     = lipgloss.Color("#0D0D0D")
	Charcoal = lipgloss.Color("#1A1A2E")
	Graphite = lipgloss.Color("#2D2D44")
	Slate    = lipgloss.Color("#4A4A6A")
	Silver   = lipgloss.Color("#8888AA")
	Pearl    = lipgloss.Color("#E8E8F0")

	Electric = lipgloss.Color("#00D4FF")
	Plasma   = lipgloss.Color("#FF006E")
	Volt     = lipgloss.Color("#ADFF02")
	Amber    = lipgloss.Color("#FFB800")
	Lavender = lipgloss.Color("#B388FF")
)

var (
	BaseStyle = lipgloss.NewStyle().
			Background(Void).
			Foreground(Pearl)

	MutedStyle = lipgloss.NewStyle().
			Foreground(Slate)

	AccentStyle = lipgloss.NewStyle().
			Foreground(Electric).
			Bold(true)
)

var (
	AppFrame = lipgloss.NewStyle().
			Border(lipgloss.ThickBorder()).
			BorderForeground(Graphite).
			Padding(1, 2)

	ElevatedCard = lipgloss.NewStyle().
			Border(lipgloss.DoubleBorder()).
			BorderForeground(Electric).
			Padding(1, 2).
			MarginBottom(1)

	SubtleCard = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(Graphite).
			Padding(1, 2)

	AlertCard = lipgloss.NewStyle().
			Background(Charcoal).
			Border(lipgloss.ThickBorder()).
			BorderForeground(Plasma).
			Padding(1, 2)
)

var (
	TitleStyle = lipgloss.NewStyle().
			Foreground(Pearl).
			Bold(true).
			MarginBottom(1)

	SectionHeader = lipgloss.NewStyle().
			Foreground(Electric).
			Bold(true).
			Underline(true)

	LabelStyle = lipgloss.NewStyle().
			Foreground(Silver).
			Bold(true)

	BodyStyle = lipgloss.NewStyle().
			Foreground(Pearl)

	DimStyle = lipgloss.NewStyle().
			Foreground(Slate).
			Italic(true)
)

var (
	SuccessStyle = lipgloss.NewStyle().
			Foreground(Volt).
			Bold(true)

	WarningStyle = lipgloss.NewStyle().
			Foreground(Amber).
			Bold(true)

	ErrorStyle = lipgloss.NewStyle().
			Foreground(Plasma).
			Bold(true)

	InfoStyle = lipgloss.NewStyle().
			Foreground(Lavender)
)

var (
	SelectedItem = lipgloss.NewStyle().
			Border(lipgloss.DoubleBorder()).
			BorderForeground(Electric).
			Padding(0, 1)

	FocusedInput = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(Electric).
			Padding(0, 1)

	HelpKeyStyle = lipgloss.NewStyle().
			Foreground(Electric).
			Bold(true)

	HelpDescStyle = lipgloss.NewStyle().
			Foreground(Silver)
)

var (
	ProgressFilled = lipgloss.NewStyle().
			Foreground(Electric)

	ProgressEmpty = lipgloss.NewStyle().
			Foreground(Graphite)

	ProgressComplete = lipgloss.NewStyle().
				Foreground(Volt)
)

var (
	StatusBar = lipgloss.NewStyle().
		Background(Charcoal).
		Foreground(Silver).
		Padding(0, 1)
)

var (
	TaskPending    = lipgloss.NewStyle().Foreground(Slate)
	TaskInProgress = lipgloss.NewStyle().Foreground(Amber).Bold(true)
	TaskCompleted  = lipgloss.NewStyle().Foreground(Volt).Bold(true)
	TaskBlocked    = lipgloss.NewStyle().Foreground(Plasma).Bold(true)
)

var (
	IconPending    = "▢"
	IconInProgress = "▶"
	IconCompleted  = "▣"
	IconBlocked    = "⊘"
	IconFailed     = "⊗"
	IconCursor     = "❯"
	IconBreadcrumb = "›"
)
