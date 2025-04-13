package main

import (
	"fmt"
	"io"
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

		if err := aictl.Run(openaiToken, rssfeed); err != nil {
			return err
		}

		sonos, err := sonosctl.New()
		if err != nil {
			return err
		}
		go sonos.StartReceiver()

		time.Sleep(2 * time.Second)

		res, err := sonos.SubscribeMediaControl()
		if err != nil {
			return err
		}
		resbody, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		fmt.Printf("resbody: %s\n", string(resbody))

		receiverHost := sonos.GetReceiverHost()

		url := fmt.Sprintf("http://%s/storage/0.mp3", receiverHost)
		if _, err := sonos.SetUri(url); err != nil {
			return err
		}

		res, err = sonos.Play()
		if err != nil {
			return err
		}
		fmt.Printf("res: %+v\n", res)

		time.Sleep(10 * time.Second)

		for i := range 3 {
			url = fmt.Sprintf("http://%s/storage/%d.mp3", receiverHost, i + 1)
			fmt.Println(url)
			res, err = sonos.SetNextURI(url)
			if err != nil {
				return err
			}
			fmt.Printf("res: %+v\n", res)
			time.Sleep(10 * time.Second)
		}

		time.Sleep(100 * time.Second)

		return nil
	},
}
