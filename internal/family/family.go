package family

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/franciscolkdo/family-feud/internal/style"
)

var _ tea.Model = Model{}

type model int

type FamilyName int

const (
	Blue FamilyName = iota
	Red
	None
)

type Model struct {
	isCurrent bool
	Id        FamilyName
	Name      string
	models    map[model]tea.Model

	score int

	failcount int
	maxFails  int

	Width  int
	Height int
	color  lipgloss.Color
}

func (m Model) Init() tea.Cmd {
	var cmds []tea.Cmd
	for _, model := range m.models {
		cmds = append(cmds, model.Init())
	}
	return tea.Batch(cmds...)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
	case winValue:
		if !m.isCurrent {
			return m, cmd
		}
		if msg < 0 {
			if m.failcount < m.maxFails {
				m.failcount++
			}
		} else {
			m.score += int(msg)
		}
	case FamilyName:
		m.isCurrent = (msg == m.Id)
	}
	return m, cmd
}

func (m Model) View() string {
	var s strings.Builder
	sty := style.RootStyle.Height(m.Height / 2)
	if m.isCurrent {
		sty = sty.Background(m.color).BorderBackground(style.DarkGray).MarginBackground(style.DarkGray)
	}

	score := style.RootStyle.Height(m.Height/6).Width(m.Width/10).Border(lipgloss.DoubleBorder()).Align(lipgloss.Center).Padding(0, 1).Render(fmt.Sprint(m.score))
	fails := style.RootStyle.Height(m.Height/6).Width(m.Width/10).Border(lipgloss.DoubleBorder()).Align(lipgloss.Center).Padding(0, 1).Render(strings.Repeat(" X ", m.failcount) + strings.Repeat("   ", m.maxFails-m.failcount))

	sc := lipgloss.Place(lipgloss.Width(fails), lipgloss.Height(score), lipgloss.Center, lipgloss.Center, score, lipgloss.WithWhitespaceBackground(sty.GetBackground()))
	body := lipgloss.JoinVertical(lipgloss.Center, m.Name, fails, sc)

	s.WriteString(sty.Border(lipgloss.DoubleBorder()).Align(lipgloss.Center).Padding(0, 1).Render(body))
	return s.String()
}

func New(cfg Config, id FamilyName, color lipgloss.Color) tea.Model {
	return Model{
		Name:      cfg.Name,
		Id:        id,
		color:     color,
		score:     0,
		maxFails:  3,
		failcount: 0,
	}
}

func NewBlueFamily(cfg Config) tea.Model {
	return New(cfg, Blue, style.RoyalBlue)
}

func NewRedFamily(cfg Config) tea.Model {
	return New(cfg, Red, style.DarkRed)
}
