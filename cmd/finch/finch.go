package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	ui "github.com/stepbrobd/finch/ui"
	"os"
)

func main() {
	fmt.Println("Finch")

	p := tea.NewProgram(ui.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error occurred when trying to start UI: %v", err)
		os.Exit(1)
	}
}
