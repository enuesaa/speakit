package main

import (
	"fmt"
	"os"

	"github.com/enuesaa/speakit/kitprot"
	"github.com/spf13/cobra"
)

func NewProtCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "prot",
		Short: "prot",
		RunE: func(cmd *cobra.Command, args []string) error {
			feed := os.Getenv("RSS")
			openaiKey := os.Getenv("OPENAI_TOKEN")
			if feed == "" || openaiKey == "" {
				return fmt.Errorf("err")
			}

			app := kitprot.GenerateFrom(&kitprot.RSSGenerator{
				Feed: feed,
			})
			app.Skip(&kitprot.UniqueSkipper{})
			app.Transform(&kitprot.AIVoiceTransformer{
				OpenAIKey: openaiKey,
			})
			app.Speak(&kitprot.BeepSpeaker{})

			return app.RunE()
		},
	}
	return cmd
}
