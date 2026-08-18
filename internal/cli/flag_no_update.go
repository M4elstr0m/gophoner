package cli

func init() {
	rootCmd.PersistentFlags().Bool("no-update", false, "disable the startup update check")
}
