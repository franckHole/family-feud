package family

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/franciscolkdo/family-feud/internal/style"
)

var _ tea.Model = FailModel{}

type FailModel struct {
	count int
}

func (m FailModel) Init() tea.Cmd {
	return nil
}

func (m FailModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case OnFamilyFailMsg:
		m.count++
	}
	return m, nil
}

func (m FailModel) View() string {
	var s strings.Builder

	score := style.RootStyle.Border(lipgloss.DoubleBorder()).Align(lipgloss.Center).Padding(1, 1)
	s.WriteString("[" + score.Render(strings.Repeat(" X ", m.count)) + "]")

	return s.String()
}

func newFail() tea.Model {
	return FailModel{
		count: 0,
	}
}
