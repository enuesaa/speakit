package cmd

import (
	"github.com/enuesaa/speakit/lib/prot"
	"github.com/spf13/cobra"
)

func NewProtCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "prot",
		Short: "prot",
		RunE: func(cmd *cobra.Command, args []string) error {
			app := prot.New()
			// Generate().
			// Transform()
			return nil
		},
	}
	return cmd
}
