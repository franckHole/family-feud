package family

import tea "github.com/charmbracelet/bubbletea"

type winValue int

func OnFamilyFail() tea.Cmd {
	return func() tea.Msg {
		return winValue(-1)
	}
}

func OnFamilyWin(value int) tea.Cmd {
	return func() tea.Msg {
		return winValue(value)
	}
}

func OnFamilySelection(id FamilyName) tea.Cmd {
	return func() tea.Msg {
		return id
	}
}
