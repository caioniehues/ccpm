package tui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
)

type ViewMode int

const (
	ViewDashboard ViewMode = iota
	ViewEpicDetail
	ViewTaskDetail
	ViewPRDDetail
	ViewWizard
	ViewEpicSelector
	ViewHelp
	ViewSettings
	ViewSearch
)

type Model struct {
	Epics       []Epic
	ActiveEpic  *Epic
	Tasks       []Task
	ActiveTask  int
	PRDs        []PRD
	ActivityLog []ActivityEntry

	CurrentView ViewMode
	Loading     bool
	LastError   error
	Ready       bool

	WizardStep int
	WizardName string

	TaskList list.Model
	Progress progress.Model
	Viewport viewport.Model
	Spinner  spinner.Model
	Help     help.Model

	Width  int
	Height int
}

type Epic struct {
	Name        string
	Status      string
	Progress    float64
	TaskCount   int
	DoneCount   int
	Description string
	PRDName     string
	Branch      string
	CreatedAt   string
	ApprovedAt  string
}

type Task struct {
	ID          string
	Name        string
	Status      string
	Epic        string
	Description string
	Progress    float64
}

type PRD struct {
	Name       string
	Status     string
	Content    string
	CreatedAt  string
	ApprovedAt string
}

type ActivityEntry struct {
	Time    string
	Message string
	Type    string
}

func NewModel() Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = AccentStyle

	p := progress.New(
		progress.WithScaledGradient(string(Electric), string(Volt)),
	)

	return Model{
		CurrentView: ViewDashboard,
		Spinner:     s,
		Help:        help.New(),
		Progress:    p,
	}
}
