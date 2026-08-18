package tui

import (
	"github.com/M4elstr0m/gophoner/internal/recon"
	tea "github.com/charmbracelet/bubbletea"
)

func Run() (*recon.Input, error) {
	a := newApp()

	finalModel, err := tea.NewProgram(a, tea.WithAltScreen()).Run()
	if err != nil {
		return nil, err
	}

	return finalModel.(*app).result, nil
}
