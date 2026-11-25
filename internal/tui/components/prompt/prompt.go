package prompt

import (
	"charm.land/bubbles/v2/textarea"
	tea "charm.land/bubbletea/v2"
)

type Model struct {
	textarea textarea.Model
}

func New() Model {
	ta := textarea.New()
	ta.KeyMap.InsertNewline.SetKeys("ctrl+j")
	ta.MaxHeight = 15
	ta.DynamicHeight = true
	ta.Focus()
	return Model{
		textarea: ta,
	}
}

func (m Model) Init() tea.Msg {
	return m.textarea.Focus()
}

func (m Model) Height() int {
	return m.textarea.Height()
}

func (m *Model) SetHeight(h int) {
	m.textarea.SetHeight(h)
}

func (m Model) Width() int {
	return m.textarea.Width()
}

func (m *Model) SetWidth(w int) {
	m.textarea.SetWidth(w)
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	m.textarea, cmd = m.textarea.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return m.textarea.View()
}
