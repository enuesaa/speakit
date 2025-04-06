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
		fmt.Printf("res: %+v\n", res)
		resbody, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		fmt.Printf("resbody: %s\n", string(resbody))

		for i := range 3 {
			receiverHost := sonos.GetReceiverHost()

			url := fmt.Sprintf("http://%s/storage/%d.mp3", receiverHost, i)
			fmt.Println(url)
			res, err := sonos.AddURIToQueue(url)
			if err != nil {
				return err
			}
			fmt.Printf("res: %+v\n", res)

			if i == 0 {
				res, err = sonos.Play()
				if err != nil {
					return err
				}
				fmt.Printf("res: %+v\n", res)
			}
		}

		// i := 0
		// sonos.OnMediaControl = func() {
		// 	receiverHost := sonos.GetReceiverHost()

		// 	url := fmt.Sprintf("http://%s/storage/%d.mp3", receiverHost, i)
		// 	res, err := sonos.SetUri(url)
		// 	if err != nil {
		// 		fmt.Printf("Error: %s\n", err.Error())
		// 		return
		// 	}
		// 	fmt.Printf("res: %+v\n", res)

		// 	res, err = sonos.Play()
		// 	if err != nil {
		// 		fmt.Printf("Error: %s\n", err.Error())
		// 		return
		// 	}
		// 	fmt.Printf("res: %+v\n", res)
		// 	i++
		// }

		// sonos.OnMediaControl()

		time.Sleep(100 * time.Second)

		return nil
	},
}
