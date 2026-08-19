package cli

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"

	"github.com/M4elstr0m/gophoner/internal/logs"
	"github.com/M4elstr0m/gophoner/internal/update"
	"github.com/M4elstr0m/gophoner/internal/version"

	_ "github.com/M4elstr0m/gophoner/internal/modules/amazon"
	_ "github.com/M4elstr0m/gophoner/internal/modules/facebook"
	_ "github.com/M4elstr0m/gophoner/internal/modules/microsoft"
)

var rootCmd = &cobra.Command{
	Use: "gophoner",
	Short: fmt.Sprintf("%s %s\nby %s\n\nCheck simultaneously if a phone number is registered on popular apps & websites",
		version.APP_NAME_STYLED, version.Version, version.AUTHOR_STYLED,
	),
	Version: version.Version,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		debug, _ := cmd.Flags().GetBool("debug")
		runDebug(debug)

		noColor, _ := cmd.Flags().GetBool("no-color")
		runNoColor(noColor)

		return nil
	},
}

func Execute() {
	if err := logs.Init(parseNoLogFlag()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if shouldNotifyUpdateUpfront() {
		update.Notify(parseNoUpdateFlag())
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		log.Error(version.APP_NAME+" exited", "error", err)
		os.Exit(1)
	}

	log.Info(version.APP_NAME + " exited")
}

func parseNoLogFlag() bool {
	for _, arg := range os.Args[1:] {
		switch arg {
		case "--no-log", "--no-log=true":
			return true
		case "--no-log=false":
			return false
		}
	}
	return false
}

func parseNoUpdateFlag() bool {
	for _, arg := range os.Args[1:] {
		switch arg {
		case "--no-update", "--no-update=true":
			return true
		case "--no-update=false":
			return false
		}
	}
	return false
}

func shouldNotifyUpdateUpfront() bool {
	if len(os.Args) < 2 {
		return true
	}
	switch os.Args[1] {
	case "interactive", "update":
		return false
	default:
		return true
	}
}
