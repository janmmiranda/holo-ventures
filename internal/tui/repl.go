package tui

import tea "github.com/charmbracelet/bubbletea"

func StartRepl() (tea.Model, error) {
	p := tea.NewProgram(InitialMenuModel(), tea.WithAltScreen())
	return p.Run()
}
