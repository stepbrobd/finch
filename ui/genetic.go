package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type AlgoMsg struct {
	ErrorRate  float64
	Generation int
	Weights    [][]float32
	Biases     [][]float32
	Outputs    [][]float32
}

func (m Model) Generation() tea.Cmd {
	return func() tea.Msg {
		return AlgoMsg{
			ErrorRate:  m.Model.RunGens(1),
			Generation: m.Model.GetNumGens(),
			Weights:    m.Model.GetWeights(),
			Biases:     m.Model.GetBiases(),
			Outputs:    m.Model.GetOutputs(),
		}
	}
}
