package ui

import (
	"fmt"
	"log"
	"time"

	"github.com/stepbrobd/finch/genetic"

	stopwatch "github.com/charmbracelet/bubbles/stopwatch"
	tea "github.com/charmbracelet/bubbletea"
	lipgloss "github.com/charmbracelet/lipgloss"
	term "golang.org/x/term"
)

type Model struct {
	Model genetic.Algo
	Watch stopwatch.Model
}

func InitialModel(algo *genetic.Algo) Model {
	return Model{
		Model: *algo,
		Watch: stopwatch.NewWithInterval(time.Second),
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.Watch.Init(),
		tea.EnterAltScreen,
	)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// esc, ctrl+c
		case "esc", "ctrl+c":
			return m, tea.Quit
		}
	}
	var cmd tea.Cmd
	m.Watch, cmd = m.Watch.Update(msg)

	return m, cmd
}

func (m Model) View() string {
	w, _, err := term.GetSize(0)
	if err != nil {
		log.Fatalf("failed to get terminal size: %v", err)
	}

	view := lipgloss.NewStyle().
		Bold(true).
		Underline(true).
		Align(lipgloss.Center).
		Render("Finch") + "\n\n"
	view += "Elapsed: " + m.Watch.View() + fmt.Sprintf(" Generation: %d", (m.Model.GetNumGens())) + "\n\n"

	return lipgloss.PlaceHorizontal(w, lipgloss.Center, view)
}
