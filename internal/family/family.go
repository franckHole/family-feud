package family

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/franciscolkdo/family-feud/internal/style"
)

var _ tea.Model = Model{}

type Model struct {
	Id    int
	Name  string
	Score tea.Model
	Fails tea.Model
}

func (m Model) Init() tea.Cmd {
	var cmds []tea.Cmd
	cmds = append(cmds, m.Fails.Init(), m.Score.Init())
	return tea.Batch(cmds...)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case OnFamilyFailMsg:
		if m.Id == msg.Id {
			m.Fails, cmd = m.Fails.Update(msg)
		}
	case OnFamilyWinMsg:
		if m.Id == msg.Id {
			m.Score, cmd = m.Score.Update(msg)
		}
	}
	return m, cmd
}

func (m Model) View() string {
	var s strings.Builder

	score := lipgloss.Place(lipgloss.Width(m.Fails.View()), lipgloss.Height(m.Score.View()), lipgloss.Center, lipgloss.Center, m.Score.View(), lipgloss.WithWhitespaceBackground(style.RootStyle.GetBackground()))
	body := lipgloss.JoinVertical(lipgloss.Center, m.Name, m.Fails.View(), score)

	s.WriteString(style.RootStyle.Border(lipgloss.DoubleBorder()).Align(lipgloss.Center).Padding(0, 1).Render(body))
	return s.String()
}

func New(cfg Config, id int) tea.Model {
	return Model{
		Name:  cfg.Name,
		Id:    id,
		Score: newScore(),
		Fails: newFail(),
	}
}
