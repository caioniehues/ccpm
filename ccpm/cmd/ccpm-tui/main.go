package main

import (
	"fmt"
	"os"

	"github.com/automazeio/ccpm/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--version", "-v", "version":
			printVersion()
			return
		case "--help", "-h", "help":
			printHelp()
			return
		}
	}

	p := tea.NewProgram(
		tui.NewModel(),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println("ccpm-tui - CCPM Terminal Dashboard")
	fmt.Println()
	fmt.Println("Usage: ccpm-tui [options]")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  --version, -v    Show version information")
	fmt.Println("  --help, -h       Show this help message")
	fmt.Println()
	fmt.Println("Keyboard shortcuts:")
	fmt.Println("  j/k              Navigate up/down")
	fmt.Println("  Enter            Select/expand")
	fmt.Println("  e                Epic detail view")
	fmt.Println("  t                Tasks view")
	fmt.Println("  w                Launch wizard")
	fmt.Println("  ?                Help")
	fmt.Println("  q                Quit")
}
