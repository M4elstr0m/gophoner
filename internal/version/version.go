package version

import "github.com/charmbracelet/lipgloss"

const (
	AUTHOR   = "M4elstr0m"
	APP_NAME = "gophoner"
)

var (
	APP_NAME_STYLED string = lipgloss.NewStyle().Bold(true).Render(APP_NAME)
	AUTHOR_STYLED   string = lipgloss.NewStyle().Bold(true).Render(AUTHOR)

	Version string = "v?.?.?"
	Commit  string = "none"
	Date    string = "unknown"
)
