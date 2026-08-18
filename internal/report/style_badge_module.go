package report

import (
	"strconv"
	"strings"

	"github.com/M4elstr0m/gophoner/internal/modules"
	"github.com/charmbracelet/lipgloss"
)

var moduleBadgeWidth = func() int {
	max := 0
	for m := modules.Module(0); m < modules.LENGTH; m++ {
		if w := len(m.String()); w > max {
			max = w
		}
	}
	return max + 2
}()

func ModuleBadge(module modules.Module) string {
	background := module.Color()

	style := badgeStyle.
		Width(moduleBadgeWidth).
		Background(lipgloss.Color(background)).
		Foreground(lipgloss.Color(moduleBadgeForeground(background)))

	return renderBadge(style, strings.ToUpper(module.String()))
}

// moduleBadgeForeground picks a light or dark foreground based on the
// perceived brightness of the badge's background color, so module badges
// stay readable regardless of how light or dark a module's color is.
func moduleBadgeForeground(hexColor string) string {
	r, g, b, ok := parseHexColor(hexColor)
	if !ok {
		return textLightForeground
	}

	luma := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
	if luma > 150 {
		return textDarkForeground
	}
	return textLightForeground
}

func parseHexColor(hex string) (r, g, b int, ok bool) {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) != 6 {
		return 0, 0, 0, false
	}

	rVal, errR := strconv.ParseInt(hex[0:2], 16, 32)
	gVal, errG := strconv.ParseInt(hex[2:4], 16, 32)
	bVal, errB := strconv.ParseInt(hex[4:6], 16, 32)
	if errR != nil || errG != nil || errB != nil {
		return 0, 0, 0, false
	}

	return int(rVal), int(gVal), int(bVal), true
}
