package tui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type FileChangedMsg struct{ Path string }
type EpicsLoadedMsg struct{ Epics []Epic }
type TasksLoadedMsg struct{ Tasks []Task }
type ErrorMsg struct{ Err error }

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.Spinner.Tick,
		loadEpics(),
	)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.handleKeyPress(msg)

	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		m.Ready = true
		m.Help.Width = msg.Width
		return m, nil

	case FileChangedMsg:
		return m, loadEpics()

	case EpicsLoadedMsg:
		m.Epics = msg.Epics
		if len(m.Epics) > 0 {
			m.ActiveEpic = &m.Epics[0]
		}
		m.Loading = false
		return m, nil

	case TasksLoadedMsg:
		m.Tasks = msg.Tasks
		return m, nil

	case ErrorMsg:
		m.LastError = msg.Err
		m.Loading = false
		return m, nil

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.Spinner, cmd = m.Spinner.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) handleKeyPress(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(msg, Keys.Quit):
		return m, tea.Quit

	case key.Matches(msg, Keys.Help):
		if m.CurrentView == ViewHelp {
			m.CurrentView = ViewDashboard
		} else {
			m.CurrentView = ViewHelp
		}
		return m, nil

	case key.Matches(msg, Keys.Back):
		m.CurrentView = ViewDashboard
		return m, nil

	case key.Matches(msg, Keys.Epic):
		m.CurrentView = ViewEpicDetail
		return m, nil

	case key.Matches(msg, Keys.Tasks):
		m.CurrentView = ViewTaskDetail
		return m, nil

	case key.Matches(msg, Keys.PRD):
		m.CurrentView = ViewPRDDetail
		return m, nil

	case key.Matches(msg, Keys.Wizard):
		m.CurrentView = ViewWizard
		return m, nil

	case key.Matches(msg, Keys.Search):
		m.CurrentView = ViewSearch
		return m, nil
	}

	return m.updateCurrentView(msg)
}

func (m Model) updateCurrentView(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch m.CurrentView {
	case ViewDashboard:
		return m.updateDashboard(msg)
	case ViewEpicDetail:
		return m.updateEpicDetail(msg)
	case ViewTaskDetail:
		return m.updateTaskDetail(msg)
	}
	return m, nil
}

func (m Model) updateDashboard(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(msg, Keys.Up):
		if m.ActiveTask > 0 {
			m.ActiveTask--
		}
	case key.Matches(msg, Keys.Down):
		if m.ActiveTask < len(m.Tasks)-1 {
			m.ActiveTask++
		}
	case key.Matches(msg, Keys.Enter):
		m.CurrentView = ViewTaskDetail
	}
	return m, nil
}

func (m Model) updateEpicDetail(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model) updateTaskDetail(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	return m, nil
}

func loadEpics() tea.Cmd {
	return func() tea.Msg {
		return EpicsLoadedMsg{Epics: []Epic{}}
	}
}
