package cli

import (
	"fmt"

	"github.com/M4elstr0m/gophoner/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("%s %s\nauthor: %s\ncommit: %s\nbuilt:  %s\n",
			version.APP_NAME_STYLED,
			version.Version,
			version.AUTHOR_STYLED,
			version.Commit,
			version.Date,
		)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
