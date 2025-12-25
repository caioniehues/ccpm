package components

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ToastType int

const (
	ToastSuccess ToastType = iota
	ToastError
	ToastWarning
	ToastInfo
)

type Toast struct {
	ID        int
	Message   string
	Type      ToastType
	CreatedAt time.Time
	Duration  time.Duration
}

type ToastModel struct {
	Toasts    []Toast
	MaxToasts int
	nextID    int
	Width     int
}

type ToastDismissMsg struct {
	ID int
}

func NewToastModel() ToastModel {
	return ToastModel{
		MaxToasts: 5,
		Width:     40,
	}
}

func (t *ToastModel) Add(msg string, toastType ToastType) tea.Cmd {
	duration := 3 * time.Second
	switch toastType {
	case ToastError:
		duration = 0
	case ToastWarning:
		duration = 5 * time.Second
	}

	t.nextID++
	toast := Toast{
		ID:        t.nextID,
		Message:   msg,
		Type:      toastType,
		CreatedAt: time.Now(),
		Duration:  duration,
	}

	t.Toasts = append(t.Toasts, toast)

	if len(t.Toasts) > t.MaxToasts {
		t.Toasts = t.Toasts[1:]
	}

	if duration > 0 {
		id := toast.ID
		return tea.Tick(duration, func(time.Time) tea.Msg {
			return ToastDismissMsg{ID: id}
		})
	}
	return nil
}

func (t *ToastModel) Dismiss(id int) {
	for i, toast := range t.Toasts {
		if toast.ID == id {
			t.Toasts = append(t.Toasts[:i], t.Toasts[i+1:]...)
			return
		}
	}
}

func (t *ToastModel) DismissAll() {
	t.Toasts = nil
}

func (t ToastModel) HasToasts() bool {
	return len(t.Toasts) > 0
}

func (t ToastModel) View() string {
	if len(t.Toasts) == 0 {
		return ""
	}

	var toastViews []string
	for _, toast := range t.Toasts {
		toastViews = append(toastViews, t.renderToast(toast))
	}

	return lipgloss.JoinVertical(lipgloss.Right, toastViews...)
}

func (t ToastModel) renderToast(toast Toast) string {
	icon := "ℹ"
	bgColor := Charcoal
	fgColor := Lavender
	borderColor := Lavender

	switch toast.Type {
	case ToastSuccess:
		icon = "✓"
		fgColor = Volt
		borderColor = Volt
	case ToastError:
		icon = "⊗"
		fgColor = Plasma
		borderColor = Plasma
	case ToastWarning:
		icon = "⚠"
		fgColor = Amber
		borderColor = Amber
	}

	dismissHint := ""
	if toast.Duration == 0 {
		dismissHint = "  [x]"
	}

	content := fmt.Sprintf("%s %s%s", icon, toast.Message, dismissHint)

	style := lipgloss.NewStyle().
		Background(bgColor).
		Foreground(fgColor).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(borderColor).
		Padding(0, 1).
		MaxWidth(t.Width)

	return style.Render(content)
}

func (t *ToastModel) Update(msg tea.Msg) tea.Cmd {
	if dismiss, ok := msg.(ToastDismissMsg); ok {
		t.Dismiss(dismiss.ID)
	}
	return nil
}
