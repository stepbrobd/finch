package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type AlgoMsg struct {
	ErrorRate  float64
	Generation int
	Biases     [][]float32
}

func (m Model) Generation() tea.Cmd {
	return func() tea.Msg {
		return AlgoMsg{
			ErrorRate:  m.Model.RunGens(1),
			Generation: m.Model.GetNumGens(),
			Biases:     m.Model.GetBiases(),
		}
	}
}
