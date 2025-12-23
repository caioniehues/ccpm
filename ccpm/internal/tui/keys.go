package tui

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Up       key.Binding
	Down     key.Binding
	Left     key.Binding
	Right    key.Binding
	Enter    key.Binding
	Back     key.Binding
	Epic     key.Binding
	Tasks    key.Binding
	PRD      key.Binding
	Sync     key.Binding
	Wizard   key.Binding
	Help     key.Binding
	Search   key.Binding
	Refresh  key.Binding
	Complete key.Binding
	Block    key.Binding
	Next     key.Binding
	Quit     key.Binding
}

var Keys = KeyMap{
	Up:       key.NewBinding(key.WithKeys("up", "k"), key.WithHelp("↑/k", "up")),
	Down:     key.NewBinding(key.WithKeys("down", "j"), key.WithHelp("↓/j", "down")),
	Left:     key.NewBinding(key.WithKeys("left", "h"), key.WithHelp("←/h", "back")),
	Right:    key.NewBinding(key.WithKeys("right", "l"), key.WithHelp("→/l", "forward")),
	Enter:    key.NewBinding(key.WithKeys("enter"), key.WithHelp("⏎", "select")),
	Back:     key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "back")),
	Epic:     key.NewBinding(key.WithKeys("e"), key.WithHelp("e", "epic")),
	Tasks:    key.NewBinding(key.WithKeys("t"), key.WithHelp("t", "tasks")),
	PRD:      key.NewBinding(key.WithKeys("p"), key.WithHelp("p", "prd")),
	Sync:     key.NewBinding(key.WithKeys("s"), key.WithHelp("s", "sync")),
	Wizard:   key.NewBinding(key.WithKeys("w"), key.WithHelp("w", "wizard")),
	Help:     key.NewBinding(key.WithKeys("?"), key.WithHelp("?", "help")),
	Search:   key.NewBinding(key.WithKeys("/"), key.WithHelp("/", "search")),
	Refresh:  key.NewBinding(key.WithKeys("r"), key.WithHelp("r", "refresh")),
	Complete: key.NewBinding(key.WithKeys("c"), key.WithHelp("c", "complete")),
	Block:    key.NewBinding(key.WithKeys("b"), key.WithHelp("b", "block")),
	Next:     key.NewBinding(key.WithKeys("n"), key.WithHelp("n", "next")),
	Quit:     key.NewBinding(key.WithKeys("q", "ctrl+c"), key.WithHelp("q", "quit")),
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Epic, k.Tasks, k.PRD, k.Wizard, k.Sync, k.Help, k.Quit}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Left, k.Right, k.Enter, k.Back},
		{k.Epic, k.Tasks, k.PRD, k.Wizard, k.Sync},
		{k.Complete, k.Block, k.Next, k.Search, k.Refresh},
		{k.Help, k.Quit},
	}
}
