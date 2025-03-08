package main

import (
	"log"

	"github.com/spf13/cobra"
)

func main() {
	app := &cobra.Command{
		Use:     "speakit",
		Short:   "Toy app to read aloud rss feed",
		Version: "0.0.3",
	}
	app.AddCommand(serveCmd)
	app.AddCommand(emitOpenapiCmd)
	app.AddCommand(speechCmd)

	// disable default
	app.SetHelpCommand(&cobra.Command{Hidden: true})
	app.CompletionOptions.DisableDefaultCmd = true
	app.SilenceUsage = true
	app.SilenceErrors = true
	app.PersistentFlags().SortFlags = false
	app.PersistentFlags().BoolP("help", "", false, "Show help information")
	app.PersistentFlags().BoolP("version", "", false, "Show version")

	if err := app.Execute(); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
