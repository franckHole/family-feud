package table

import tea "github.com/charmbracelet/bubbletea"

type Status int

const (
	Success Status = iota
	Failed
)

type ResultMsg struct {
	Status Status
	Points int
}

func onResult(status Status, points int) tea.Cmd {
	return func() tea.Msg {
		return ResultMsg{Status: status, Points: points}
	}
}
