package tui

import (
	"github.com/M4elstr0m/gophoner/internal/recon"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type View interface {
	Init() tea.Cmd
	Update(tea.Msg) (View, tea.Cmd)
	View() string
}

var (
	labelStyle = lipgloss.NewStyle().Bold(true).MarginBottom(1)

	buttonStyle = lipgloss.NewStyle().
			Padding(0, 2).
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(lipgloss.Color("#444444"))

	buttonFocusedStyle = buttonStyle.
				Background(lipgloss.Color("#7D56F4"))

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#cc0000"))
)

func padForEvenCentering(content string, termWidth int) string {
	width := lipgloss.Width(content)
	if (termWidth-width)%2 == 0 {
		return content
	}
	return lipgloss.NewStyle().Width(width + 1).Render(content)
}

type app struct {
	current View

	countryCode string
	phoneNumber string

	result *recon.Input
}

func newApp() *app {
	return &app{current: newPhoneView()}
}

func (a *app) Init() tea.Cmd {
	return a.current.Init()
}

func (a *app) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return a, tea.Quit
		}

	case cancelledMsg:
		return a, tea.Quit

	case backMsg:
		a.current = newPhoneViewWithValues(a.countryCode, a.phoneNumber)
		return a, a.current.Init()

	case phoneSubmittedMsg:
		a.countryCode = msg.CountryCode
		a.phoneNumber = msg.PhoneNumber
		a.current = newModuleView()
		return a, a.current.Init()

	case modulesSubmittedMsg:
		a.result = &recon.Input{
			PhoneNumber: a.countryCode + a.phoneNumber,
			ModuleSlice: msg.Modules,
		}
		return a, tea.Quit
	}

	var cmd tea.Cmd
	a.current, cmd = a.current.Update(msg)
	return a, cmd
}

func (a *app) View() string {
	return a.current.View()
}
