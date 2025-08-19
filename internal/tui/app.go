package tui

import (
	"swagger_to_test/internal/tui/views"

	tea "github.com/charmbracelet/bubbletea"
)

type App struct {
	program *tea.Program
}

func NewApp() (*App, error) {
	homeView := views.InitalHomeView()
	return &App{
		program: tea.NewProgram(homeView),
	}, nil
}

func (a *App) Run() (tea.Model, error) {
	return a.program.Run()
}
