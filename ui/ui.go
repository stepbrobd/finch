package ui

import (
	"github.com/stepbrobd/finch/genetic"

	textinput "github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Model  genetic.Net
	Inputs []textinput.Model // field count is equal to the neurons in the input layer
}

func InitialModel() Model {
	return Model{}
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

	return view
}
