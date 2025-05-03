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

			app := prot.GenerateFrom(&prot.RSSFeedGenerator{
				Feed: feed,
			})
			app.Transform(&prot.TTSTransformer{
				OpenAIKey: openaiKey,
			})
			app.Speak(&prot.BeepSpeaker{})

			app.Controller(&prot.SampleController{})

			return app.Run()
		},
	}
	return cmd
}
