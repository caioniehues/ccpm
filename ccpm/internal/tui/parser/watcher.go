package parser

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fsnotify/fsnotify"
)

type FileChangedMsg struct {
	Path      string
	Operation string
	Time      time.Time
}

type WatcherErrorMsg struct {
	Err error
}

type Watcher struct {
	watcher   *fsnotify.Watcher
	baseDir   string
	debounce  time.Duration
	lastEvent map[string]time.Time
}

func NewWatcher(baseDir string) (*Watcher, error) {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	return &Watcher{
		watcher:   w,
		baseDir:   baseDir,
		debounce:  100 * time.Millisecond,
		lastEvent: make(map[string]time.Time),
	}, nil
}

func (w *Watcher) Start(p *tea.Program) error {
	epicsDir := filepath.Join(w.baseDir, ".claude", "epics")
	if err := w.watchRecursive(epicsDir); err != nil {
		// Directory might not exist yet, that's OK
	}

	prdsDir := filepath.Join(w.baseDir, ".claude", "prds")
	if err := w.watcher.Add(prdsDir); err != nil {
		// Directory might not exist yet, that's OK
	}

	go w.processEvents(p)

	return nil
}

func (w *Watcher) watchRecursive(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.IsDir() {
			return w.watcher.Add(path)
		}
		return nil
	})
}

func (w *Watcher) processEvents(p *tea.Program) {
	for {
		select {
		case event, ok := <-w.watcher.Events:
			if !ok {
				return
			}

			if !strings.HasSuffix(event.Name, ".md") {
				continue
			}

			if last, exists := w.lastEvent[event.Name]; exists {
				if time.Since(last) < w.debounce {
					continue
				}
			}
			w.lastEvent[event.Name] = time.Now()

			op := "write"
			if event.Has(fsnotify.Create) {
				op = "create"
			} else if event.Has(fsnotify.Remove) {
				op = "remove"
			}

			p.Send(FileChangedMsg{
				Path:      event.Name,
				Operation: op,
				Time:      time.Now(),
			})

		case err, ok := <-w.watcher.Errors:
			if !ok {
				return
			}
			p.Send(WatcherErrorMsg{Err: err})
		}
	}
}

func (w *Watcher) Close() error {
	return w.watcher.Close()
}
