package tui

import "github.com/charmbracelet/lipgloss"

type GameMode struct {
	mode string
}

var (
	NavMode      = GameMode{mode: "Nav"}
	ConvMode     = GameMode{mode: "Conv"}
	BattleMode   = GameMode{mode: "Battle"}
	InteractMode = GameMode{mode: "Interact"}

	WindowHeight      = 20
	WindowWidth       = 40
	BorderStyleNormal = lipgloss.NewStyle().
				Width(WindowWidth).
				Height(WindowHeight).
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("62")).
				Padding(4)
	BorderStyleThick = lipgloss.NewStyle().
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("69")).
				Padding(4)
	DefaultStyle = lipgloss.NewStyle().BorderForeground(lipgloss.Color("36")).BorderStyle(lipgloss.NormalBorder()).Padding(1).Width(WindowWidth)
)
