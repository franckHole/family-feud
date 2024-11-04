package score

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/franciscolkdo/family-feud/internal/keymap"
	"github.com/franciscolkdo/family-feud/internal/style"
)

type Model struct {
	score  int
	keyMap keymap.KeyMap
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case WinRound:
		cmd = OnWinRoundScore(m.score)
	case ScoreMsg:
		m.score += msg.value
	}
	return m, cmd
}

func (m Model) View() string {
	score := style.RootStyle.Border(lipgloss.DoubleBorder()).Align(lipgloss.Center).Padding(0, 1)
	return score.Render(fmt.Sprint(m.score))

}

func New() tea.Model {
	return Model{
		score:  0,
		keyMap: keymap.DefaultKeyMap(),
	}
}
