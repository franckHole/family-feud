package table

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/franciscolkdo/family-feud/internal/style"
)

const defaultWidth = 30

type animationMsg struct {
	Id int // Id is the position on table (1 to 8)
}

func startAnimation(id int) tea.Cmd {
	return func() tea.Msg {
		return animationMsg{Id: id}
	}
}

type Box struct {
	Id           int // Id is the position on table (1 to 8)
	points       int
	content      string
	frames       []string
	currentFrame int
	width        int

	ticker *time.Ticker
}

func (m Box) nextFrame() tea.Cmd {
	return func() tea.Msg {
		<-m.ticker.C
		return animationMsg{Id: m.Id}
	}
}

func (m Box) showContent() string {
	var s strings.Builder
	s.WriteString(fmt.Sprint(m.points))

	s.WriteString(" | ")
	if len(m.content) > m.width {
		s.WriteString(m.content[:m.width])
	} else {
		s.WriteString(m.content + strings.Repeat(" ", m.width-len(m.content+s.String())))
	}
	return s.String()
}

func (m Box) setStates() Box {
	m.frames = []string{}
	for _, frame := range []struct {
		border  lipgloss.Border
		content string
	}{
		{lipgloss.RoundedBorder(), strings.Repeat(" ", m.width)},
		{lipgloss.HiddenBorder(), strings.Repeat("═", m.width)},
		{lipgloss.HiddenBorder(), strings.Repeat("─", m.width)},
		{lipgloss.HiddenBorder(), strings.Repeat("═", m.width)},
		{lipgloss.RoundedBorder(), m.showContent()},
	} {
		m.frames = append(m.frames, box(frame.border, frame.content))
	}

	return m
}

func (m Box) Init() tea.Cmd {
	return nil
}

func (m Box) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case animationMsg:
		if msg.Id == m.Id {
			if m.currentFrame < len(m.frames)-1 {
				m.currentFrame++
				cmd = m.nextFrame()
			} else {
				m.ticker.Stop()
				cmd = onResult(Success, m.points)
			}
		}
	}
	return m, cmd
}

func (m Box) View() string {
	return m.frames[m.currentFrame]
}

func newBox(cfg BoxConfig, id int) tea.Model {
	m := Box{
		Id:           id,
		content:      cfg.Answer,
		points:       cfg.Points,
		frames:       []string{""},
		currentFrame: 0,
		width:        defaultWidth,
		ticker:       time.NewTicker(500 * time.Millisecond),
	}
	m = m.setStates()
	return m
}

func box(border lipgloss.Border, content string) string {
	var s strings.Builder
	contentStyle := style.RootStyle.Border(border).Align(lipgloss.Center).Padding(0, 0).Margin(0, 1)

	s.WriteString(contentStyle.Render(content))

	return style.RootStyle.Render(s.String())
}
