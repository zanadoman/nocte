package tui

import tea "charm.land/bubbletea/v2"

type Model struct{}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyPressMsg); ok && msg.String() == "q" {
		return m, tea.Quit
	}
	return m, nil
}

func (m Model) View() tea.View {
	v := tea.NewView("hello, world")
	v.AltScreen = true
	return v
}
