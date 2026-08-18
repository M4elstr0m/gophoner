package cli

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

func init() {
	rootCmd.PersistentFlags().Bool("no-color", false, "disable colored/styled output (for redirecting to a file)")
}

func runNoColor(noColor bool) {
	if noColor {
		lipgloss.SetColorProfile(termenv.Ascii)
	}
}
