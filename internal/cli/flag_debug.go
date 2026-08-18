package cli

import "github.com/charmbracelet/log"

func init() {
	rootCmd.PersistentFlags().Bool("debug", false, "set log level to debug")
}

func runDebug(debug bool) {
	if debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}
