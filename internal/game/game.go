package game

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/franciscolkdo/family-feud/internal/family"
	"github.com/franciscolkdo/family-feud/internal/keymap"
	"github.com/franciscolkdo/family-feud/internal/score"
	"github.com/franciscolkdo/family-feud/internal/style"
	"github.com/franciscolkdo/family-feud/internal/table"
)

var _ tea.Model = Model{}

type Model struct {
	families      []tea.Model
	currentFamily int
	table         tea.Model
	totalScore    tea.Model

	keyMap keymap.KeyMap
}

func (m Model) getAllModels() []tea.Model {
	var models []tea.Model

	models = append(models, m.totalScore, m.table)
	models = append(models, m.families...)
	return models
}

func (m Model) Init() tea.Cmd {
	var cmds []tea.Cmd
	for _, model := range m.getAllModels() {
		cmds = append(cmds, model.Init())
	}
	return tea.Batch(cmds...)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, m.keyMap.Quit) {
			return m, tea.Quit
		}
		if key.Matches(msg, m.keyMap.GoodChoice, m.keyMap.WrongChoice) {
			var cmd tea.Cmd
			m.table, cmd = m.table.Update(msg)
			cmds = append(cmds, cmd)
		}
		if key.Matches(msg, m.keyMap.SwitchFamily) {
			m.currentFamily = 1 - m.currentFamily
		}
		if key.Matches(msg, m.keyMap.WinRound) {
			var cmd tea.Cmd
			m.totalScore, cmd = m.totalScore.Update(msg)
			cmds = append(cmds, cmd)
		}
	case score.WinRoundScoreMsg:
		cmds = append(cmds, family.OnFamilyWin(m.currentFamily, msg.Value))
	case table.ResultMsg:
		var cmd tea.Cmd
		if msg.Status == table.Success {
			cmd = score.OnScoreMsg(msg.Points)
		} else {
			cmd = family.OnFamilyFail(m.currentFamily)
		}
		cmds = append(cmds, cmd)
	default:
		var cmd tea.Cmd
		m.totalScore, cmd = m.totalScore.Update(msg)
		cmds = append(cmds, cmd)
		m.table, cmd = m.table.Update(msg)
		cmds = append(cmds, cmd)
		for i := range m.families {
			m.families[i], cmd = m.families[i].Update(msg)
			cmds = append(cmds, cmd)
		}
	}
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	var s strings.Builder

	left := lipgloss.Place(lipgloss.Width(m.families[0].View()), lipgloss.Height(m.table.View()), lipgloss.Left, lipgloss.Top, m.families[0].View(), lipgloss.WithWhitespaceBackground(style.RootStyle.GetBackground()))
	center := lipgloss.Place(lipgloss.Width(m.table.View()), lipgloss.Height(m.table.View()), lipgloss.Left, lipgloss.Top, m.table.View(), lipgloss.WithWhitespaceBackground(style.RootStyle.GetBackground()))
	right := lipgloss.Place(lipgloss.Width(m.families[1].View()), lipgloss.Height(m.table.View()), lipgloss.Left, lipgloss.Top, m.families[1].View(), lipgloss.WithWhitespaceBackground(style.RootStyle.GetBackground()))
	body := lipgloss.JoinHorizontal(lipgloss.Top, left, center, right)

	points := lipgloss.Place(lipgloss.Width(body), lipgloss.Height(m.totalScore.View()), lipgloss.Center, lipgloss.Center, m.totalScore.View(), lipgloss.WithWhitespaceBackground(style.RootStyle.GetBackground()))

	s.WriteString(lipgloss.JoinVertical(lipgloss.Top, points, body))
	return s.String()
}

func New(cfg Config) tea.Model {
	var families []tea.Model
	for i, f := range cfg.Families {
		families = append(families, family.New(f, i))
	}
	return Model{
		families:   families,
		totalScore: score.New(),
		table:      table.New(cfg.Table),
		keyMap:     keymap.DefaultKeyMap(),
	}
}
