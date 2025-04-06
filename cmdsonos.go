package main

import (
	"fmt"
	"io"
	// "os"
	"time"

	// "github.com/enuesaa/speakit/internal/aictl"
	"github.com/enuesaa/speakit/internal/sonosctl"
	"github.com/spf13/cobra"
)

var sonosCmd = &cobra.Command{
	Use:   "sonos",
	Short: "sonos",
	RunE: func(cmd *cobra.Command, args []string) error {
		// openaiToken := os.Getenv("OPENAI_TOKEN")
		// rssfeed := os.Getenv("RSS")

		// if err := aictl.Run(openaiToken, rssfeed); err != nil {
		// 	return err
		// }

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

		// receiverHost := sonos.GetReceiverHost()

		// url := fmt.Sprintf("http://%s/storage/1.mp3", receiverHost)
		// res, err := sonos.SetUri(url)
		// if err != nil {
		// 	return err
		// }
		// fmt.Printf("res: %+v\n", res)

		// res, err = sonos.Play()
		// if err != nil {
		// 	return err
		// }
		// fmt.Printf("res: %+v\n", res)

		time.Sleep(100 * time.Second)

		return nil
	},
}
