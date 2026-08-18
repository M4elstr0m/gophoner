package cli

func init() {
	rootCmd.PersistentFlags().Bool("no-log", false, "discard all logging output (incognito mode)")
}
