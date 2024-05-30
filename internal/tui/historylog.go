package tui

import (
	"github.com/charmbracelet/lipgloss"
)

type HistoryModel struct {
	Width       int
	Height      int
	History     []string
	InFocus     bool
	BorderColor lipgloss.Style
}

func NewHistory() *HistoryModel {
	return &HistoryModel{
		Width:       WindowWidth,
		Height:      WindowHeight,
		History:     make([]string, 0),
		InFocus:     false,
		BorderColor: DefaultStyle,
	}
}
