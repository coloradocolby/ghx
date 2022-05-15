package utils

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Up            key.Binding
	Down          key.Binding
	PageDown      key.Binding
	PageUp        key.Binding
	NextTab       key.Binding
	PrevTab       key.Binding
	TogglePreview key.Binding
	OpenGithub    key.Binding
	SwitchView    key.Binding
	Refresh       key.Binding
	Help          key.Binding
	Quit          key.Binding
	Search        key.Binding
	FocusInput    key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down},
		{k.PrevTab, k.NextTab},
		{k.TogglePreview, k.OpenGithub},
		{k.Help, k.Quit},
	}
}

var Keys = KeyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	PrevTab: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("/h", "previous tab"),
	),
	NextTab: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("/l", "next tab"),
	),
	PageUp: key.NewBinding(
		key.WithKeys("ctrl+u"),
		key.WithHelp("ctrl+u", "preview page up"),
	),
	PageDown: key.NewBinding(
		key.WithKeys("ctrl+d"),
		key.WithHelp("ctrl+d", "preview page down"),
	),
	TogglePreview: key.NewBinding(
		key.WithKeys("p"),
		key.WithHelp("p", "open in Preview"),
	),
	OpenGithub: key.NewBinding(
		key.WithKeys("o"),
		key.WithHelp("o", "open in GitHub"),
	),
	Refresh: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "refresh"),
	),
	SwitchView: key.NewBinding(
		key.WithKeys("s"),
		key.WithHelp("s", "switch view"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		// key.WithKeys("q", "esc", "ctrl+c"),
		key.WithKeys("esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
	Search: key.NewBinding(
		// key.WithKeys("q", "esc", "ctrl+c"),
		key.WithKeys("enter"),
		key.WithHelp("enter", "search"),
	),
	FocusInput: key.NewBinding(
		// key.WithKeys("q", "esc", "ctrl+c"),
		key.WithKeys("/"),
		key.WithHelp("/", "focus input"),
	),
}
