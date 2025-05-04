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
			app.Transform(&prot.AITransformer{
				OpenAIKey: openaiKey,
				PromptTmpl: "次の文章を podcast で読み上げるので文面を調整ください. 50文字程度に. ハルシネーションしないで. \ntext: {{.text}}",
			})
			app.Transform(&prot.TTSTransformer{
				OpenAIKey: openaiKey,
			})
			app.Speak(&prot.BeepSpeaker{})

			return app.Run()
		},
	}
	return cmd
}
