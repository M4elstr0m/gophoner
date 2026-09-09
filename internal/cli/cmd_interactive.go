package cli

import (
	"github.com/M4elstr0m/gophoner/internal/report"
	"github.com/M4elstr0m/gophoner/internal/tui"
	"github.com/M4elstr0m/gophoner/internal/update"
	"github.com/spf13/cobra"
)

var interactiveCmd = &cobra.Command{
	Use:   "interactive",
	Short: "Launch the interactive TUI",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		noUpdate, _ := cmd.Flags().GetBool("no-update")
		return runInteractive(noUpdate)
	},
}

func init() {
	rootCmd.AddCommand(interactiveCmd)
}

func runInteractive(noUpdate bool) error {
	input, err := tui.Run()
	if err != nil {
		return err
	} else if input == nil {
		return nil
	}

	out, err := report.Run(input)
	if err != nil {
		return err
	}

	report.Print(input.PhoneNumber, out)

	update.Notify(noUpdate)

	return nil
}
