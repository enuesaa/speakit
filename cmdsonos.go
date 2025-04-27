package main

import (
	"fmt"
	"os"
	"time"

	"github.com/enuesaa/speakit/internal/aictl"
	"github.com/enuesaa/speakit/internal/sonosctl"
	"github.com/spf13/cobra"
)

var sonosCmd = &cobra.Command{
	Use:   "sonos",
	Short: "sonos",
	RunE: func(cmd *cobra.Command, args []string) error {
		openaiToken := os.Getenv("OPENAI_TOKEN")
		rssfeed := os.Getenv("RSS")

		sonos, err := sonosctl.New()
		if err != nil {
			return err
		}
		go sonos.StartReceiver()

		if _, err := sonos.SubscribeMediaControl(); err != nil {
			return err
		}

		if err := aictl.Run(openaiToken, rssfeed); err != nil {
			return err
		}

		for i := range 3 {
			url := fmt.Sprintf("http://%s/storage/%d.mp3", sonos.GetReceiverHost(), i)
			if _, err = sonos.SetNextURI(url); err != nil {
				return err
			}
			if i == 0 {
				if _, err = sonos.Play(); err != nil {
					return err
				}
			}
			time.Sleep(10 * time.Second)
		}
		time.Sleep(100 * time.Second)

		return nil
	},
}
