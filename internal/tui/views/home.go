package views

import (
	"fmt"
	"swagger_to_test/pkg/logger"

	tea "github.com/charmbracelet/bubbletea"
)

type HomeView struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func InitalHomeView() HomeView {
	return HomeView{
		choices:  []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},
		selected: make(map[int]struct{}),
	}
}

func (h HomeView) Init() tea.Cmd {
	return nil
}

func (h HomeView) View() string {
	s := "Welcome to swagger to test tool\n"
	for i, choice := range h.choices {
		cursor := " "
		if h.cursor == i {
			cursor = ">"
		}
		checked := ""
		if _, ok := h.selected[i]; ok {
			checked = "x"
		}
		logger.Info("%s", checked)
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}
	s += "\nPress q or ctl+c to quit.\n"
	return s
}

func (m HomeView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}
