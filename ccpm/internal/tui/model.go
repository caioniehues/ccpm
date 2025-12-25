package tui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"

	"github.com/automazeio/ccpm/internal/tui/components"
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

	TaskList    list.Model
	Progress    progress.Model
	Viewport    viewport.Model
	Spinner     spinner.Model
	SyncSpinner spinner.Model
	Help        help.Model

	Toasts     components.ToastModel
	Animations components.AnimationState
	Layout     LayoutDimensions

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
	s := components.BrailleSpinner()
	syncS := components.SyncSpinner()

	p := progress.New(
		progress.WithScaledGradient(string(Electric), string(Volt)),
	)

	return Model{
		CurrentView: ViewDashboard,
		Spinner:     s,
		SyncSpinner: syncS,
		Help:        help.New(),
		Progress:    p,
		Toasts:      components.NewToastModel(),
		Animations:  components.NewAnimationState(),
	}
}

func (m *Model) ShowSuccess(msg string) {
	m.Toasts.Add(msg, components.ToastSuccess)
}

func (m *Model) ShowError(msg string) {
	m.Toasts.Add(msg, components.ToastError)
}

func (m *Model) ShowWarning(msg string) {
	m.Toasts.Add(msg, components.ToastWarning)
}

func (m *Model) ShowInfo(msg string) {
	m.Toasts.Add(msg, components.ToastInfo)
}
