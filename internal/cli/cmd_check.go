package cli

import (
	"strings"

	"github.com/M4elstr0m/gophoner/internal/modules"
	"github.com/M4elstr0m/gophoner/internal/recon"
	"github.com/M4elstr0m/gophoner/internal/report"
	"github.com/spf13/cobra"
)

var checkCmdFlags struct {
	Target string

	CountryCode string
	PhoneNumber string

	ModuleSlice modules.ModuleSlice
	AllModules  bool
}

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check if a phone number is registered on a service",
	RunE: func(cmd *cobra.Command, args []string) error {
		runCheck()
		return nil
	},
}

func init() {
	checkCmd.Flags().StringVarP(&checkCmdFlags.Target, "target", "t", "", "full phone number including country code, e.g. +15551234567")
	checkCmd.Flags().StringVarP(&checkCmdFlags.CountryCode, "country-code", "c", "", "country code, e.g. +1")
	checkCmd.Flags().StringVarP(&checkCmdFlags.PhoneNumber, "phone", "p", "", "national phone number, e.g. 5551234567")
	checkCmd.Flags().VarP(&checkCmdFlags.ModuleSlice, "module", "m", "module(s) to check, comma-separated (e.g. amazon,uber)")
	checkCmd.Flags().BoolVarP(&checkCmdFlags.AllModules, "all", "A", false, "check against every available module")

	checkCmd.MarkFlagsRequiredTogether("country-code", "phone")
	checkCmd.MarkFlagsMutuallyExclusive("target", "country-code")
	checkCmd.MarkFlagsMutuallyExclusive("target", "phone")
	checkCmd.MarkFlagsOneRequired("target", "country-code")

	checkCmd.MarkFlagsMutuallyExclusive("module", "all")
	checkCmd.MarkFlagsOneRequired("module", "all")

	rootCmd.AddCommand(checkCmd)
}

func runCheck() {
	phoneNumber := checkCmdFlags.Target
	if phoneNumber == "" {
		if !strings.Contains(checkCmdFlags.CountryCode, "+") {
			checkCmdFlags.CountryCode = "+" + checkCmdFlags.CountryCode
		}
		phoneNumber = checkCmdFlags.CountryCode + checkCmdFlags.PhoneNumber
	}

	moduleSlice := checkCmdFlags.ModuleSlice
	if checkCmdFlags.AllModules {
		moduleSlice = modules.All()
	}

	out := recon.Run(&recon.Input{
		PhoneNumber: phoneNumber,
		ModuleSlice: moduleSlice,
	})

	report.Print(phoneNumber, out)
}
