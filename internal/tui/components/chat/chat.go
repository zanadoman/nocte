package chat

import "charm.land/bubbles/v2/viewport"

type Model struct {
	viewport viewport.Model
}

func New() Model {
	vp := viewport.New()
	vp.SetContent("hello, world")
	return Model{
		viewport: vp,
	}
}

func (m Model) Height() int {
	return m.viewport.Height()
}

func (m *Model) SetHeight(h int) {
	m.viewport.SetHeight(h)
}

func (m Model) Width() int {
	return m.viewport.Width()
}

func (m *Model) SetWidth(w int) {
	m.viewport.SetWidth(w)
}

func (m Model) View() string {
	return m.viewport.View()
}
