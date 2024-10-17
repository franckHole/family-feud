package family

import tea "github.com/charmbracelet/bubbletea"

type OnFamilyFailMsg struct {
	Id int
}

func OnFamilyFail(id int) tea.Cmd {
	return func() tea.Msg {
		return OnFamilyFailMsg{Id: id}
	}
}

type OnFamilyWinMsg struct {
	Id    int
	Score int
}

func OnFamilyWin(id int, score int) tea.Cmd {
	return func() tea.Msg {
		return OnFamilyWinMsg{Id: id, Score: score}
	}
}
