package cmd

import "github.com/spf13/cobra"

func New() *cobra.Command {
	app := &cobra.Command{
		Use:     "speakit",
		Short:   "Toy app to read aloud rss feed",
	}
	app.AddCommand(NewSonosCmd())
	app.AddCommand(NewProtCmd())

	// disable default
	app.SetHelpCommand(&cobra.Command{Hidden: true})
	app.CompletionOptions.DisableDefaultCmd = true
	app.SilenceUsage = true
	app.SilenceErrors = true
	app.PersistentFlags().SortFlags = false
	app.PersistentFlags().BoolP("help", "", false, "Show help information")
	app.PersistentFlags().BoolP("version", "", false, "Show version")

	return app
}
