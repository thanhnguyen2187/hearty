package main

import (
	_ "embed"
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fatih/color"
	"hearty/text"
)

type Model struct {
	heartStates [][]string
	index       int
}

func NewModel() Model {
	return Model{
		heartStates: text.HeartStates(),
		index:       0,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "right", "enter":
			m.index = (m.index + 1) % len(m.heartStates)
		}
	}
	return m, nil
}

func (m Model) View() string {
	if m.index%2 == 0 {
		color.Set(color.FgHiCyan)
	} else {
		color.Set(color.FgHiMagenta)
	}
	return strings.Join(m.heartStates[m.index], "\n")
}

func main() {
	p := tea.NewProgram(NewModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("main error %v", err)
		return
	}
}
