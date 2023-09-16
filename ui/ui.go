package ui

import (
	"fmt"

	textinput "github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	// TODO: change definition to the one defined in genetics/genetics.go
	model  [][]float32       // the model
	inputs []textinput.Model // field count is equal to the neurons in the input layer
}

func InitialModel() Model {
	return Model{
		model: [][]float32{
			{0.0, 0.0, 0.0},
			{0.0, 0.0, 0.0, 0.0, 0.0},
			{0.0, 0.0},
		},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// q, ctrl+c
		case "q", "ctrl+c":
			return m, tea.Quit

		}
	}

	return m, nil
}

func (m Model) View() string {
	view := StyleBlue500.Render("Finch") + "\n\n"

	// print vertically with padding
	for _, row := range m.model {
		for _, col := range row {
			view += fmt.Sprintf("%f ", col)
		}
		view += "\n"
	}

	return view
}
