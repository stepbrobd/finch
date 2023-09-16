package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	gen "github.com/stepbrobd/finch/genetic"
	ui "github.com/stepbrobd/finch/ui"
)

func main() {
	gen.Test()
	p := tea.NewProgram(ui.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("error occurred when trying to start UI: %v", err)
		os.Exit(1)
	}
}
