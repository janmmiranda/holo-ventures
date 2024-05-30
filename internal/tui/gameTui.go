package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/janmmiranda/holo-ventures/internal/game/entities"
)

type gameModel struct {
	width      int
	height     int
	input      textinput.Model
	inputFocus bool
	history    HistoryModel
	character  entities.Entity
	textView   []string
	mode       GameMode
}

func InitialGameModel() *gameModel {
	ti := textinput.New()
	ti.Placeholder = "Enter argument"
	ti.Focus()
	ti.CharLimit = 256
	ti.Width = WindowWidth

	h := NewHistory()

	m := &gameModel{
		input:      ti,
		inputFocus: true,
		history:    *h,
		character:  *entities.NewCharacter(),
		textView:   make([]string, 0),
		mode:       NavMode,
	}
	return m
}

func (m *gameModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *gameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var inputCmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			m.textView = make([]string, 0)
			argVal := m.input.Value()
			m.history.History = append(m.history.History, argVal)
			m.textView = append(m.textView, argVal)
			m.input.Reset()
		}
	}
	m.input, inputCmd = m.input.Update(msg)
	cmds = append(cmds, inputCmd)
	return m, tea.Batch(cmds...)
}

func (m *gameModel) View() string {
	var hsb strings.Builder
	hsb.WriteString("History\n")
	for i := 0; i < len(m.history.History); i++ {
		fmt.Fprintf(&hsb, " - %s\n", m.history.History[i])
	}
	var tsb strings.Builder
	tsb.WriteString("Input\n")
	for i := 0; i < len(m.textView); i++ {
		fmt.Fprintf(&tsb, "%s\n", m.textView[i])
	}

	return lipgloss.Place(
		80,
		40,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			BorderStyleNormal.Render(hsb.String()),
			BorderStyleNormal.Render(
				lipgloss.JoinVertical(
					lipgloss.Left,
					tsb.String(),
					m.input.View(),
				),
			),
		),
	)
}
