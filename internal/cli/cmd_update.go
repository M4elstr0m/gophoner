package cli

import (
	"fmt"

	"github.com/M4elstr0m/gophoner/internal/update"
	"github.com/M4elstr0m/gophoner/internal/version"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Check whether a new version of gophoner is available",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		noUpdate, _ := cmd.Flags().GetBool("no-update")
		var hasPrinted bool = update.Notify(noUpdate)
		if !hasPrinted {
			fmt.Printf("Your %s seems up-to-date!\n", version.APP_NAME_STYLED)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
