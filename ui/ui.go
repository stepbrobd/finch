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
	Msg   AlgoMsg
}

func InitialModel(algo *genetic.Algo) Model {
	return Model{
		Model: *algo,
		Watch: stopwatch.NewWithInterval(time.Second),
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.Generation(),
		m.Watch.Init(),
		tea.ClearScreen,
		tea.EnterAltScreen,
	)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case AlgoMsg:
		m.Msg = msg
		return m, m.Generation()
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c":
			return m, tea.Quit
		}
	default:
		var cmd tea.Cmd
		m.Watch, cmd = m.Watch.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m Model) View() string {
	w, _, err := term.GetSize(0)
	if err != nil {
		log.Fatalf("failed to get terminal size: %v", err)
	}

	view := lipgloss.NewStyle().
		Bold(true).
		Underline(true).
		Render("Finch") + "\n\n"

	view += "Elapsed: " + m.Watch.View() +
		fmt.Sprintf(" | Generation: %d", (m.Msg.Generation)) +
		fmt.Sprintf(" | Error: %.25f", m.Msg.ErrorRate) +
		"\n\n\n"

	view += "Weights" + "\n\n"
	for i := range m.Msg.Weights {
		view += MkView(m.Model.GetWeights(), i) + "\n"
	}
	view += "\n\n"

	view += "Biases" + "\n\n"
	for i := range m.Msg.Biases {
		view += MkView(m.Model.GetBiases(), i) + "\n"
	}
	view += "\n\n"

	view += "Outputs" + "\n\n"
	for i := range m.Msg.Outputs {
		view += MkView(m.Model.GetOutputs(), i) + "\n"
	}

	return lipgloss.PlaceHorizontal(w, lipgloss.Center, view)
}
