package family

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/franciscolkdo/family-feud/internal/style"
)

var _ tea.Model = FailModel{}

type FailModel struct {
	count     int
	maxFailed int
}

func (m FailModel) Init() tea.Cmd {
	return nil
}

func (m FailModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case OnFamilyFailMsg:
		if m.count < m.maxFailed {
			m.count++
		}
	}
	return m, nil
}

func (m FailModel) View() string {
	score := style.RootStyle.Border(lipgloss.DoubleBorder()).Align(lipgloss.Center).Padding(0, 1)

	return score.Render(strings.Repeat(" X ", m.count) + strings.Repeat("   ", m.maxFailed-m.count))
}

func newFail() tea.Model {
	return FailModel{
		count:     0,
		maxFailed: 3,
	}
}
