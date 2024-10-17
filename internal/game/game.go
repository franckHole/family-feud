package game

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/franciscolkdo/family-feud/internal/keymap"
	"github.com/franciscolkdo/family-feud/internal/table"
)

var _ tea.Model = Model{}

type Model struct {
	families   []tea.Model
	table      tea.Model
	totalScore tea.Model

	keyMap keymap.KeyMap
}

func (m Model) Init() tea.Cmd {
	var cmds []tea.Cmd
	var models []tea.Model

	models = append(models, m.totalScore, m.table)
	models = append(models, m.families...)
	for _, model := range models {
		cmds = append(cmds, model.Init())
	}
	return tea.Batch(cmds...)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, m.keyMap.Quit) {
			return m, tea.Quit
		}
		if key.Matches(msg, m.keyMap.GoodChoice, m.keyMap.WrongChoice) {
			m.table, cmd = m.table.Update(msg)
		}
	case table.ResultMsg:

	}
	return m, cmd
}

func (m Model) View() string {
	var s strings.Builder

	return s.String()
}

// func New(cfg Config) tea.Model {
// 	var boxes []tea.Model
// 	for i, answer := range cfg.Boxes {
// 		boxes = append(boxes, newBox(answer, i))
// 	}
// 	return Model{
// 		boxes:  boxes,
// 		keyMap: keymap.DefaultKeyMap(),
// 	}
// }
