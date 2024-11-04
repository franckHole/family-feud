package score

import tea "github.com/charmbracelet/bubbletea"

type ScoreMsg struct {
	value int
}

func OnScore(value int) tea.Cmd {
	return func() tea.Msg {
		return ScoreMsg{value: value}
	}
}

type WinRoundScoreMsg struct {
	Value int
}

func OnWinRoundScore(value int) tea.Cmd {
	return func() tea.Msg {
		return WinRoundScoreMsg{Value: value}
	}
}

type WinRound struct{}

func OnWinRound() tea.Cmd {
	return func() tea.Msg {
		return WinRound{}
	}
}
