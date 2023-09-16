package ui

import (
	lipgloss "github.com/charmbracelet/lipgloss"
)

// https://tailwindcss.com/docs/customizing-colors
const (
	Purple400 = lipgloss.Color("#c084fc")
	Purple500 = lipgloss.Color("#a855f7")
	Purple600 = lipgloss.Color("#9333ea")
	Purple700 = lipgloss.Color("#7e22ce")
	Purple800 = lipgloss.Color("#6b21a8")
)

// activation levels
var (
	AL1 = lipgloss.NewStyle().Foreground(Purple400)
	AL2 = lipgloss.NewStyle().Foreground(Purple500)
	AL3 = lipgloss.NewStyle().Foreground(Purple600)
	AL4 = lipgloss.NewStyle().Foreground(Purple700)
	AL5 = lipgloss.NewStyle().Foreground(Purple800)
)
