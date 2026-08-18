package report_test

import (
	"testing"

	"github.com/M4elstr0m/gophoner/internal/recon"
	"github.com/M4elstr0m/gophoner/internal/report"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"

	_ "github.com/M4elstr0m/gophoner/internal/cli"
)

func init() {
	lipgloss.SetColorProfile(termenv.TrueColor)
}

func TestPrintPreview(t *testing.T) {
	results := []recon.Output{
		{Indicator: recon.Registered},
		{Indicator: recon.NotRegistered},
		{Indicator: recon.LimitReached},
		{Indicator: recon.ModuleError},
	}

	report.Print("+15551234567", results)
}
