package report

import (
	"fmt"

	"github.com/M4elstr0m/gophoner/internal/recon"
	"github.com/charmbracelet/lipgloss"
)

const (
	badgeLabelRegistered      string = "REGISTERED"
	badgeBackgroundRegistered string = "#009427"

	badgeLabelNotRegistered      string = "NOT REGISTERED"
	badgeBackgroundNotRegistered string = "#cc0000"

	badgeLabelRateLimit      string = "RATE LIMIT"
	badgeBackgroundRateLimit string = "#d9b500"

	badgeLabelError      string = "ERROR"
	badgeBackgroundError string = "#b800cd"

	indicatorBadgeWidth int = len(badgeLabelNotRegistered) + 2

	badgeBackgroundNeutral string = "#757575"
)

func Print(phoneNumber string, results []recon.Output) {
	fmt.Printf("%s Target: %s\n\n",
		renderBadge(
			badgeStyle.
				Background(lipgloss.Color(badgeBackgroundNeutral)),
			"RESULTS",
		),
		lipgloss.NewStyle().
			Bold(true).
			Render(phoneNumber),
	)

	var registeredCount uint = 0
	for _, result := range results {
		print_result(result)
		if result.Indicator == recon.Registered {
			registeredCount++
		}
	}

	fmt.Printf("\n%s Registered on %v service(s)\n",
		renderBadge(
			badgeStyle.
				Background(lipgloss.Color(badgeBackgroundNeutral)),
			"OVERVIEW",
		),
		registeredCount,
	)
}

func print_result(result recon.Output) {
	var indicatorBadge string

	var adjustedIndicatorBadgeStyle = badgeStyle.Width(indicatorBadgeWidth)

	switch result.Indicator {
	case recon.Registered:
		indicatorBadge = renderBadge(
			adjustedIndicatorBadgeStyle.
				Background(lipgloss.Color(badgeBackgroundRegistered)),
			badgeLabelRegistered,
		)
	case recon.NotRegistered:
		indicatorBadge = renderBadge(
			adjustedIndicatorBadgeStyle.
				Background(lipgloss.Color(badgeBackgroundNotRegistered)),
			badgeLabelNotRegistered,
		)
	case recon.LimitReached:
		indicatorBadge = renderBadge(
			adjustedIndicatorBadgeStyle.
				Background(lipgloss.Color(badgeBackgroundRateLimit)).
				Foreground(lipgloss.Color(textDarkForeground)),
			badgeLabelRateLimit,
		)
	case recon.ModuleError:
		indicatorBadge = renderBadge(
			adjustedIndicatorBadgeStyle.
				Background(lipgloss.Color(badgeBackgroundError)),
			badgeLabelError,
		)
	}

	// Additionnal Infos would go here

	fmt.Println(ModuleBadge(result.Module), indicatorBadge)
}
