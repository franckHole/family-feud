package table

import (
	"math"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/franciscolkdo/family-feud/internal/keymap"
	"github.com/franciscolkdo/family-feud/internal/style"
)

var _ tea.Model = Model{}

type Model struct {
	boxes  []tea.Model
	keyMap keymap.KeyMap
}

func (m Model) Init() tea.Cmd {
	var cmds []tea.Cmd
	for _, box := range m.boxes {
		cmds = append(cmds, box.Init())
	}
	return tea.Batch(cmds...)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, m.keyMap.GoodChoice) {
			i, err := strconv.Atoi(msg.String())
			if err != nil {
				// TODO: log it
				return m, nil
			}
			if i > 0 && i < len(m.boxes)+1 {
				return m, startAnimation(i - 1)
			}
		} else if key.Matches(msg, m.keyMap.WrongChoice) {
			return m, onResult(Failed, 0)
		}
	default:
		var cmds []tea.Cmd
		for i := range m.boxes {
			var cmd tea.Cmd
			m.boxes[i], cmd = m.boxes[i].Update(msg)
			cmds = append(cmds, cmd)
		}
		return m, tea.Batch(cmds...)
	}
	return m, nil
}

func (m Model) View() string {
	var s strings.Builder

	// Split boxes in 2 columns
	mid := int(math.Ceil(float64(len(m.boxes)) / 2))
	var l, r strings.Builder

	for i, box := range m.boxes {
		buf := &l
		if i+1 > mid {
			buf = &r
		}
		buf.WriteString(box.View())
		// Add new line except last box in each column
		if i != mid-1 && i < len(m.boxes)-1 {
			buf.WriteString("\n")
		}
	}

	right := lipgloss.Place(lipgloss.Width(l.String()), lipgloss.Height(l.String()), lipgloss.Left, lipgloss.Top, r.String(), lipgloss.WithWhitespaceBackground(style.RootStyle.GetBackground()))
	body := lipgloss.JoinHorizontal(lipgloss.Top, l.String(), right)

	table := style.RootStyle.Border(lipgloss.DoubleBorder()).Align(lipgloss.Center).Padding(0, 1)
	s.WriteString(table.Render(body))
	return s.String()
}

func New(cfg Config) tea.Model {
	var boxes []tea.Model
	for i, answer := range cfg.Boxes {
		boxes = append(boxes, newBox(answer, i))
	}
	return Model{
		boxes:  boxes,
		keyMap: keymap.DefaultKeyMap(),
	}
}
