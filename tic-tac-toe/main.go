package main

import (
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tttUI struct {
	markers [9]string
	current int
}

const (
	z = " "
	x = "X"
	o = "O"
)

func NewTTTUI() tttUI {
	t := tttUI{}
	for i := range 9 {
		t.markers[i] = z
	}
	return t
}

func (t tttUI) Get(i int) string {
	if i == t.current {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("#000000")).Background(lipgloss.Color("#FFFFFF")).Render(t.markers[i])
	}
	return t.markers[i]
}

// Init implements tea.Model.
func (t tttUI) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model.
func (t tttUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return t, tea.Quit
		case "k":
			t.current = (t.current + 6) % 9
		case "j":
			t.current = (t.current + 3) % 9
		case "l":
			t.current = (t.current + 1) % 9
		case "h":
			t.current = (t.current - 1) % 9
		case "x":
			t.markers[t.current] = x
		case "o":
			t.markers[t.current] = o
		}
	}
	return t, nil
}

// View implements tea.Model.
func (t tttUI) View() string {
	return "Tic Tac Toe\n" +
		" " + t.Get(0) + " | " + t.Get(1) + " | " + t.Get(2) + "\n" +
		strings.Repeat("-", 11) + "\n" +
		" " + t.Get(3) + " | " + t.Get(4) + " | " + t.Get(5) + "\n" +
		strings.Repeat("-", 11) + "\n" +
		" " + t.Get(6) + " | " + t.Get(7) + " | " + t.Get(8) + "\n"
}

var _ tea.Model = tttUI{}

func main() {
	_, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatal(err)
	}
	p := tea.NewProgram(NewTTTUI())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
