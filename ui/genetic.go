package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type AlgoMsg struct {
	Generation int
	Err        float64
	Biases     [][]float32
}

func (m Model) Start() tea.Cmd {
	model := m.Model
	cmd := func() tea.Msg {
		er := model.RunGens(1)
		gn := model.GetNumGens()
		bs := model.GetBiases()
		return AlgoMsg{Generation: gn, Err: er, Biases: bs}
	}
	m.Model = model
	return cmd
}

func (m Model) ModelUpdate() (Model, tea.Cmd) {
	model := m.Model
	cmd := func() tea.Msg {
		er := model.RunGens(1)
		gn := model.GetNumGens()
		bs := model.GetBiases()
		return AlgoMsg{Generation: gn, Err: er, Biases: bs}
	}
	m.Model = model
	return m, cmd
}
