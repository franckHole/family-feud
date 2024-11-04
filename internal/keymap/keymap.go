package keymap

import "github.com/charmbracelet/bubbles/key"

// KeyMap defines key bindings for each user action.
type KeyMap struct {
	Quit        key.Binding
	BlueFamily  key.Binding
	RedFamily   key.Binding
	WinRound    key.Binding
	GoodChoice  key.Binding
	WrongChoice key.Binding
}

// DefaultKeyMap defines the default keybindings.
func DefaultKeyMap() KeyMap {
	return KeyMap{
		Quit:        key.NewBinding(key.WithKeys("ctrl+c"), key.WithHelp("ctrl+c", "quit")),
		WinRound:    key.NewBinding(key.WithKeys("w"), key.WithHelp("w", "Win round")),
		BlueFamily:  key.NewBinding(key.WithKeys("b"), key.WithHelp("b", "Blue family")),
		RedFamily:   key.NewBinding(key.WithKeys("r"), key.WithHelp("r", "Red family")),
		GoodChoice:  key.NewBinding(key.WithKeys("1", "2", "3", "4", "5", "6", "7", "8"), key.WithHelp("1-2-3-4-5-6-7-8", "Good choice")),
		WrongChoice: key.NewBinding(key.WithKeys("x", "0"), key.WithHelp("0/x", "failed choice")),
	}
}
