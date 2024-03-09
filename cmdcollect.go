package main

import (
	"bytes"
	"fmt"
	"time"

	"github.com/enuesaa/speakit/pkg/repository"
	"github.com/enuesaa/speakit/pkg/service"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"github.com/spf13/cobra"
)

func init() {
	collectCmd.Flags().StringVar(&voicevoxBaseUrl, "voicevox", "http://localhost:50021", "voicevox host")
}

var collectCmd = &cobra.Command{
	Use:   "collect",
	Short: "collect <url>",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("please specify rss url to fetch")
		}
		url := args[0]
		fmt.Printf("found url: %s\n", url)

		env := repository.Env{
			VOICEVOX_BASE_URL: voicevoxBaseUrl,
		}
		repos := repository.NewRepos(env)
		feedSrv := service.NewFeedSevice(repos)
		programSrv := service.NewProgramService(repos)

		realfeed, err := feedSrv.TryFetch(url)
		if err != nil {
			return err
		}
		fmt.Printf("fetched %d items\n", len(realfeed.Items))

		for i, realfeeditem := range realfeed.Items {
			content, err := programSrv.TryConvert(realfeeditem.Title)
			if err != nil {
				return err
			}
			fmt.Printf("converted %d\n", i)

			streamer, format, err := wav.Decode(bytes.NewReader([]byte(content)))
			if err != nil {
				return err
			}
			defer streamer.Close()
			fmt.Printf("decoded %d\n", i)

			err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
			if err != nil {
				return err
			}

			fast := beep.ResampleRatio(4, 1.3, streamer)
	
			done := make(chan bool)
			speaker.Play(beep.Seq(fast, beep.Callback(func() {
				done <- true
			})))
			<-done

			if i > 1 {
				break
			}
		}

		return nil
	},
}
