package main

import (
	"github.com/spf13/cobra"
)

func main() {
	app := &cobra.Command{
		Use:     "speakit",
		Short:   "Toy app to read aloud rss feed",
		Version: "0.0.2",
	}
	app.AddCommand(emitOpenapiCmd)
	app.AddCommand(tryFetchCmd)
	app.AddCommand(serveCmd)

	// disable default
	app.SetHelpCommand(&cobra.Command{Hidden: true})
	app.CompletionOptions.DisableDefaultCmd = true
	app.SilenceUsage = true
	app.PersistentFlags().SortFlags = false
	app.PersistentFlags().BoolP("help", "", false, "Show help information")
	app.PersistentFlags().BoolP("version", "", false, "Show version")
	app.Execute()
}
