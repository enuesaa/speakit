package cmd

import (
	"fmt"
	"os"

	"github.com/enuesaa/speakit/lib/prot"
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

			app := prot.GenerateFrom(&prot.RSSGenerator{
				Feed: feed,
			})
			app.Skipper(&prot.UniqueSkipper{
				UniqueField: "link",
			})
			app.Transform(&prot.AIVoiceTransformer{
				OpenAIKey: openaiKey,
			})
			app.Speak(&prot.BeepSpeaker{})

			return app.Run()
		},
	}
	return cmd
}
