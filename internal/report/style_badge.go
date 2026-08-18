package report

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	useCaps = false // nerd font feature

	badgeCapLeft  = "\uE0B6"
	badgeCapRight = "\uE0B4"
)

var badgeStyle = lipgloss.NewStyle().
	Bold(true).
	Padding(0, 1).
	Align(lipgloss.Center).
	Foreground(lipgloss.Color(textLightForeground))

func renderBadge(style lipgloss.Style, label string) string {
	if useCaps {
		capStyle := lipgloss.NewStyle().Foreground(style.GetBackground())
		return capStyle.Render(badgeCapLeft) + style.Render(label) + capStyle.Render(badgeCapRight)
	} else {
		return style.Render(label)
	}
}
