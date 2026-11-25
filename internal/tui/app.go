package tui

import (
	"nocte/internal/tui/components/chat"
	"nocte/internal/tui/components/prompt"

	tea "charm.land/bubbletea/v2"
)

type Model struct {
	height int
	chat   chat.Model
	prompt prompt.Model
}

func New() Model {
	return Model{
		chat:   chat.New(),
		prompt: prompt.New(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		if msg.String() == "ctrl+q" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.chat.SetWidth(msg.Width)
		m.prompt.SetWidth(msg.Width)
	}
	var cmd tea.Cmd
	m.prompt, cmd = m.prompt.Update(msg)
	m.chat.SetHeight(m.height - m.prompt.Height())
	return m, cmd
}

func (m Model) View() tea.View {
	v := tea.NewView(m.chat.View() + "\n" + m.prompt.View())
	v.AltScreen = true
	return v
}
