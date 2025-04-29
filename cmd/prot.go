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
			if feed == "" {
				return fmt.Errorf("err")
			}

			app := prot.New(&prot.RSSFeedGenerator{
				Feed: feed,
			})
			app.Transform(&prot.CustomTransformer{})

			return app.Speak(&prot.SonosSpeaker{})
		},
	}
	return cmd
}
