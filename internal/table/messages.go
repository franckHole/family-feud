package table

import (
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
)

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

type Choice int

func OnWrongChoice() tea.Cmd {
	return func() tea.Msg {
		return Choice(0)
	}
}

func OnGoodChoice(v string) tea.Cmd {
	return func() tea.Msg {
		i, err := strconv.Atoi(v)
		if err != nil {
			// TODO: log it
			return -1
		}
		return Choice(i)
	}
}
