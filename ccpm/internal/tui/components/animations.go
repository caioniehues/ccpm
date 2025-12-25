package components

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	FlashDuration   = 300 * time.Millisecond
	ShakeDuration   = 200 * time.Millisecond
	PulseDuration   = 1000 * time.Millisecond
	CursorBlinkRate = 530 * time.Millisecond
)

type SuccessFlashMsg struct {
	Target string
}

type SuccessFlashDoneMsg struct {
	Target string
}

type ErrorShakeMsg struct {
	Target string
}

type ErrorShakeDoneMsg struct {
	Target string
}

type PulseTickMsg struct{}

type CursorBlinkMsg struct{}

type AnimationState struct {
	FlashTargets  map[string]bool
	ShakeTargets  map[string]bool
	PulseActive   bool
	CursorVisible bool
}

func NewAnimationState() AnimationState {
	return AnimationState{
		FlashTargets:  make(map[string]bool),
		ShakeTargets:  make(map[string]bool),
		CursorVisible: true,
	}
}

func (a AnimationState) IsFlashing(target string) bool {
	return a.FlashTargets[target]
}

func (a AnimationState) IsShaking(target string) bool {
	return a.ShakeTargets[target]
}

func SuccessFlash(target string) tea.Cmd {
	return func() tea.Msg {
		return SuccessFlashMsg{Target: target}
	}
}

func SuccessFlashEnd(target string) tea.Cmd {
	return tea.Tick(FlashDuration, func(time.Time) tea.Msg {
		return SuccessFlashDoneMsg{Target: target}
	})
}

func ErrorShake(target string) tea.Cmd {
	return func() tea.Msg {
		return ErrorShakeMsg{Target: target}
	}
}

func ErrorShakeEnd(target string) tea.Cmd {
	return tea.Tick(ShakeDuration, func(time.Time) tea.Msg {
		return ErrorShakeDoneMsg{Target: target}
	})
}

func StartPulse() tea.Cmd {
	return tea.Tick(PulseDuration, func(time.Time) tea.Msg {
		return PulseTickMsg{}
	})
}

func StartCursorBlink() tea.Cmd {
	return tea.Tick(CursorBlinkRate, func(time.Time) tea.Msg {
		return CursorBlinkMsg{}
	})
}

func (a *AnimationState) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case SuccessFlashMsg:
		a.FlashTargets[msg.Target] = true
		return SuccessFlashEnd(msg.Target)

	case SuccessFlashDoneMsg:
		delete(a.FlashTargets, msg.Target)

	case ErrorShakeMsg:
		a.ShakeTargets[msg.Target] = true
		return ErrorShakeEnd(msg.Target)

	case ErrorShakeDoneMsg:
		delete(a.ShakeTargets, msg.Target)

	case PulseTickMsg:
		a.PulseActive = !a.PulseActive
		return StartPulse()

	case CursorBlinkMsg:
		a.CursorVisible = !a.CursorVisible
		return StartCursorBlink()
	}

	return nil
}

func ShakeOffset(shaking bool, tick int) string {
	if !shaking {
		return ""
	}
	if tick%2 == 0 {
		return "  "
	}
	return ""
}
