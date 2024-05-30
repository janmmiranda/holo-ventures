package tui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type menuModel struct {
	choiceList list.Model
	width      int
	height     int
}

type itemChoice struct {
	choice string
}

func (i itemChoice) Title() string       { return i.choice }
func (i itemChoice) Description() string { return "" }
func (i itemChoice) FilterValue() string { return i.choice }

func InitialMenuModel() menuModel {
	menuItems := []itemChoice{itemChoice{choice: "Play"}, itemChoice{choice: "Settings"}, itemChoice{choice: "Quit"}}
	menuList := list.New(convertToItems(menuItems), list.NewDefaultDelegate(), WindowWidth, WindowHeight)
	menuList.Title = "Menu"
	return menuModel{
		choiceList: menuList,
	}
}

func (m menuModel) Init() tea.Cmd {
	return nil
}

func (m menuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter", " ":
			c, ok := m.choiceList.SelectedItem().(itemChoice)
			if ok {
				switch c.choice {
				case "Play":
					return InitialGameModel(), nil
				case "Settings":

				case "Quit":
					return m, tea.Quit
				}
			}
		}
	}

	m.choiceList, cmd = m.choiceList.Update(msg)
	return m, cmd
}

func (m menuModel) View() string {
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Left,
			BorderStyleNormal.Render(m.choiceList.View()),
		),
	)
}
