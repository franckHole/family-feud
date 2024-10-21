package family

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/franciscolkdo/family-feud/internal/style"
)

var _ tea.Model = ScoreModel{}

type ScoreModel struct {
	value int
}

func (m ScoreModel) Init() tea.Cmd {
	return nil
}

func (m ScoreModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case OnFamilyWinMsg:
		m.value = msg.Score
	}
	return m, nil
}

func (m ScoreModel) View() string {
	score := style.RootStyle.Border(lipgloss.DoubleBorder()).Align(lipgloss.Center).Padding(0, 1)
	return score.Render(fmt.Sprint(m.value))
}

func newScore() tea.Model {
	return ScoreModel{
		value: 0,
	}
}
