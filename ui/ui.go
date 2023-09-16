package ui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func InitialModel() Model {
	return Model{
		choices:  []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},
		selected: make(map[int]struct{}),
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

		// k, arrow up
		case "k", "up":
			if m.cursor > 0 {
				m.cursor--
			}

		// j, arrow down
		case "j", "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// space, enter
		case " ", "enter":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	return m, nil
}

func (m Model) View() string {
	view := "What should we buy at the market?\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		view += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	view += "\nPress q to quit.\n"

	return view
}
